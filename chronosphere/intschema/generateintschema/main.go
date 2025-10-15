// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// generateintschema maps our Terraform schemas into internal struct
// representations which can freely convert to and from Terraform ResourceData.
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/multierr"
	"golang.org/x/exp/maps"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/intschematag"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/registry"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// Add shared schema references here to generate shared types.
// Insert in sorted order.
var sharedSchemaTypeNames = map[*schema.Schema]string{
	tfschema.KeyValueLogParserSchema:                "KeyValueParser",
	tfschema.LogAllocationConfigSchema:              "LogAllocationConfigSchema",
	tfschema.LogParserSchema:                        "LogParser",
	tfschema.LogPrioritiesSchema:                    "LogPrioritiesSchema",
	tfschema.LogSearchFilterSchema:                  "LogSearchFilterSchema",
	tfschema.MatcherListSchema:                      "Matcher",
	tfschema.MonitorSeriesConditionSchema:           "MonitorSeriesCondition",
	tfschema.RegexLogParserSchema:                   "RegexParser",
	tfschema.ResourcePoolAllocationSchema:           "ResourcePoolAllocationSchema",
	tfschema.ResourcePoolPrioritiesSchema:           "ResourcePoolPrioritiesSchema",
	tfschema.ResourcePoolAllocationThresholdsSchema: "ResourcePoolAllocationThresholdsSchema",
	tfschema.ResourcePoolAllocationThresholdSchema:  "ResourcePoolAllocationThresholdSchema",
	tfschema.TraceBoolFilterSchema:                  "TraceBoolFilter",
	tfschema.TraceDurationFilterSchema:              "TraceDurationFilter",
	tfschema.TraceFilterSchema:                      "TraceFilter",
	tfschema.TraceNumericFilterSchema:               "TraceNumericFilter",
	tfschema.TraceSearchFilterSchema:                "TraceSearchFilter",
	tfschema.TraceSpanCountFilterSchema:             "TraceSpanCountFilter",
	tfschema.TraceSpanFilterListSchema:              "TraceSpanFilter",
	tfschema.TraceStringFilterSchema:                "TraceStringFilter",
	tfschema.TraceTagFilterSchema:                   "TraceTagFilter",
	tfschema.ValueMappingsSchema:                    "ValueMappings",
	tfschema.SLOAdditionalPromQLFilters:             "SLOAdditionalPromQLFilters",
	tfschema.SignalGrouping:                         "SignalGrouping",
	tfschema.PartitionFilterSchema:                  "PartitionFilter",
	// Field normalization schemas are handled via sharedElemTypeNames below
	// Log control schemas
	tfschema.LogFieldPathSchema: "LogControlConfigFieldPath",
}

// Add shared element references here to generate shared types. Usually we
// shouldn't be consolidating on an element type, we should be consolidating on
// the container type. This is only for extremely rare cases where the
// containers diverge without affecting the generated intschema code.
//
// TODO: To keep things simple, we don't validate that there are no
// duplicate struct defs for shared element resources since this is so rare. If
// it becomes more common, we should add validation.
var sharedElemTypeNames = map[*schema.Resource]string{
	tfschema.NotificationRouteSchema.Elem.(*schema.Resource): "NotificationRoute",
	tfschema.ResourcePoolElemSchema:                          "ResourcePoolsConfigPool",
	// Log field normalization resources need to be here to prevent the generator from creating
	// duplicate types when it processes nested resources within NamedStringNormalizationResource
	tfschema.StringNormalizationResource:      "LogIngestConfigStringNormalization",
	tfschema.NamedStringNormalizationResource: "LogIngestConfigNamedStringNormalization",
	tfschema.TimestampNormalizationResource:   "LogIngestConfigTimestampNormalization",
	tfschema.LogFieldSelectorResource:         "LogFieldPath",
}

// Exhaustive list of all "xxx_id" fields which identifies whether the field
// should be generated as a tfid.ID or not (i.e. does the field refer to a
// registered Terraform resource or not).
var fieldIsTFID = map[string]bool{
	"bucket_id":                true,
	"collection_id":            true,
	"dashboard_id":             true,
	"notification_policy_id":   true,
	"team_id":                  true,
	"execution_group":          true,
	"notification_policy_data": true,
	"action_ids":               true,
	"dataset_id":               true,
	"consumption_config_id":    true,

	"callback_id": false,
	"external_id": false,
	"project_id":  false,
	"client_id":   false,
	"tenant_id":   false,
}

var fieldAddFile = map[string]bool{
	"dashboard_json": true,
}

// Exhaustive list of all set/list fields which identifies whether the field
// should be generated as a []tfid.ID.
// Unlike fieldIsTFID, it only matches set/list fields.
var listFieldIsTFID = map[*schema.Schema]bool{
	mustLookup(tfschema.NotificationRouteSchema.Elem.(*schema.Resource).Schema, "notifiers"): true,
}

// List of fields which are associated with recursive container types.
// Terraform does not support recursive schemas, but our API does, so the
// Terraform schema is implemented with a max depth (see makeRecursiveResource
// in tfschema), and then mapped to a proper recursive intschema type to make
// TF<->API mapping easy. The server is responsible for enforcing <= max depth
// that the Terraform provider uses, else import state will fail.
var recursiveFields = map[string]*schema.Schema{
	"partition": tfschema.ConsumptionConfig["partition"],
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	resourceSchemas := make(map[string]map[string]*schema.Schema)
	for _, r := range registry.Resources {
		resourceSchemas[r.Name] = r.Schema
	}
	resourceSchemas["test_resource"] = tfschema.TestResource

	files := []*fileData{
		newSharedFile(),
	}
	for _, name := range sortedKeys(resourceSchemas) {
		files = append(files, newFileData(schemaTypeResource, name, resourceSchemas[name]))
	}
	files = append(files, newFileData(schemaTypeData, "bucket", tfschema.DataBucket))
	files = append(files, newFileData(schemaTypeData, "collection", tfschema.DataCollection))
	files = append(files, newFileData(schemaTypeData, "service", tfschema.DataService))

	if err := validateNoDuplicateStructDefs(files); err != nil {
		return err
	}
	for _, f := range files {
		if err := f.writeFile(); err != nil {
			return err
		}
	}
	return nil
}

func mustLookup(m map[string]*schema.Schema, k string) *schema.Schema {
	s, ok := m[k]
	if !ok {
		panic(fmt.Errorf("key %q not found", k))
	}
	return s
}

func validateNoDuplicateStructDefs(files []*fileData) error {
	refCounts := make(map[*schema.Schema]int)
	for _, f := range files {
		for _, d := range f.Structs {
			if d.container == nil {
				continue
			}
			refCounts[d.container]++
		}
	}
	var errs error
	for ref, n := range refCounts {
		_, ok := sharedSchemaTypeNames[ref]
		if n > 1 && !ok {
			errs = multierr.Append(errs, fmt.Errorf(
				"duplicate schema ref found in %d places: %v (to share types, must add to sharedSchemaTypeNames)",
				n, sortedKeys(ref.Elem.(*schema.Resource).Schema)))
		}
	}
	return errs
}

func newSharedFile() *fileData {
	f := &fileData{
		filename: "shared_schemas.go",
	}
	for s, typeName := range sharedSchemaTypeNames {
		f.newRootStructDef(typeName, s.Elem.(*schema.Resource).Schema)
	}
	for e, typeName := range sharedElemTypeNames {
		f.newRootStructDef(typeName, e.Schema)
	}

	// Above map iterations are non-deterministic and we can't use sortedKeys
	// as the map keys are not strings.
	// Instead, sort the generated structs to make output deterministic.
	sort.Slice(f.Structs, func(i, j int) bool {
		return f.Structs[i].TypeName < f.Structs[j].TypeName
	})
	return f
}

type schemaType string

const (
	schemaTypeResource = "resource"
	schemaTypeData     = "data"
)

func (st schemaType) filename(name string) string {
	if st.Datasource() {
		return "data_" + name + ".go"
	}
	return name + ".go"
}

func (st schemaType) structName(name string) string {
	n := upperCamelCase(name)
	if st.Datasource() {
		return "Data" + n
	}
	return n
}

// Datasource returns if this is a schema for a datasource.
func (st schemaType) Datasource() bool {
	return st == schemaTypeData
}

// MarshalAddFuncName returns the marshal function to call for this schema type.
func (st schemaType) MarshalAddFuncName() string {
	if st == schemaTypeData {
		return "AddData"
	}
	return "AddResource"
}

func newFileData(t schemaType, name string, objSchema map[string]*schema.Schema) *fileData {
	f := &fileData{
		Imports: []string{
			"io",
			"github.com/hashicorp/terraform-plugin-sdk/v2/diag",
			"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema",
			"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/convertintschema",
			"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal",
			"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema",
		},
		Type:     t,
		filename: t.filename(name),
	}
	f.newRootStructDef(t.structName(name), objSchema)

	// The newStructDef recursion will always push the main resource struct to
	// the last element, so we reverse to ensure it shows up first for
	// readability.
	reverse(f.Structs)
	r := f.Structs[0]
	r.IsConverter = true
	r.ResourceName = "chronosphere_" + name

	r.Fields = append(r.Fields, &field{
		Name:     intschematag.StateIDField,
		TypeName: "string",
		Tag:      intschematag.InternalFieldTag(),
		Comment: []string{
			"Internal identifier used in the .state file, i.e. ResourceData.Id().",
			"Cannot be set, else ToResourceData will panic.",
		},
	}, &field{
		Name:     intschematag.HCLIDField,
		TypeName: "string",
		Tag:      intschematag.InternalFieldTag(),
		Comment: []string{
			"HCL-level identifier used in the .tf file. FromResourceData will always",
			"leave this empty, and ToResourceData will panic if set.",
		},
	})

	return f
}

type fileData struct {
	Imports []string
	Structs []*structDef

	Type     schemaType
	filename string
}

func (f *fileData) writeFile() error {
	b := &bytes.Buffer{}
	if err := fileTemplate.Execute(b, f); err != nil {
		return fmt.Errorf("template execution failed: %v", err)
	}
	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("gofmt failed: %v", err)
	}
	fmt.Printf("writing file %s\n", f.filename)
	return os.WriteFile(f.filename, formatted, 0o644)
}

func (f *fileData) newRootStructDef(typeName string, objSchema map[string]*schema.Schema) {
	f.newStructDef(typeName, objSchema, nil /* container */, "" /* inRecursiveField */)
}

func (f *fileData) newStructDef(
	typeName string, objSchema map[string]*schema.Schema, container *schema.Schema,
	inRecursiveField string,
) {
	d := &structDef{
		TypeName:  typeName,
		container: container,
	}
	// TODO: Order keys by:
	// - ID / Name / Slug
	// - Parent refs
	// - Required fields
	// - Optional fields
	for _, name := range sortedKeys(objSchema) {
		fieldName := upperCamelCase(name)
		typeInfo := f.loadType(typeName, fieldName, name, objSchema[name], nil /* container */, inRecursiveField)
		tag := intschematag.Tag{
			TFName:            name,
			Optional:          objSchema[name].Optional || typeInfo.optionalListEncodedObject,
			Computed:          objSchema[name].Computed,
			ListEncodedObject: typeInfo.listEncodedObject,
		}
		if d := objSchema[name].Default; d != nil {
			tag.Default = fmt.Sprint(d)
		}
		d.Fields = append(d.Fields, &field{
			Name:     fieldName,
			TypeName: typeInfo.name,
			Tag:      tag,
		})

		if addFileField(name, objSchema[name]) {
			d.Fields = append(d.Fields, &field{
				Name:     "HCLFile" + fieldName,
				TypeName: "string",
				Tag: intschematag.Tag{
					TFName: name,
					File:   true,
				},
			})
		}
	}
	sort.SliceStable(d.Fields, func(i, j int) bool {
		return d.Fields[i].sortScore() < d.Fields[j].sortScore()
	})
	f.Structs = append(f.Structs, d)
}

type typeInfo struct {
	name                      string
	listEncodedObject         bool
	optionalListEncodedObject bool
}

func (f *fileData) loadType(
	parentTypeName, fieldName, tfName string, fieldMeta any, container *schema.Schema,
	inRecursiveField string,
) (result typeInfo) {
	if typeName, ok := sharedSchemaTypeNames[container]; ok {
		return typeInfo{name: typeName}
	}
	switch t := fieldMeta.(type) {
	case *schema.Schema:
		switch t.Type {
		case schema.TypeList, schema.TypeSet:
			if tfschema.IsListEncodedObject(t) {
				typeName := f.loadType(
					parentTypeName, fieldName, tfName, t.Elem, t, inRecursiveField).name
				optional := false
				if t.Optional || t.MinItems == 0 {
					// Optional nested objects are wrapped in pointers to ensure
					// we can still distinguish an empty list from a non-empty
					// list containing an empty value.
					typeName = "*" + typeName
					optional = true
				}
				return typeInfo{
					name:                      typeName,
					listEncodedObject:         true,
					optionalListEncodedObject: optional,
				}
			} else if listFieldIsTFID[t] {
				return typeInfo{name: "[]tfid.ID"}
			} else {
				return typeInfo{
					name: "[]" + f.loadType(
						parentTypeName, fieldName, tfName, t.Elem, t, inRecursiveField).name,
				}
			}
		case schema.TypeMap:
			return typeInfo{
				name: "map[string]" + f.loadType(
					parentTypeName, fieldName, tfName, t.Elem, t, inRecursiveField).name,
			}
		case schema.TypeString:
			typeName := "string"
			if isTFID(tfName) {
				typeName = "tfid.ID"
			}
			return typeInfo{name: typeName}
		case schema.TypeBool:
			return typeInfo{name: "bool"}
		case schema.TypeFloat:
			return typeInfo{name: "float64"}
		case schema.TypeInt:
			return typeInfo{name: "int64"}
		default:
			panic(fmt.Sprintf("unhandled type: %s", t.Type))
		}
	case *schema.Resource:
		if typeName, ok := sharedElemTypeNames[t]; ok {
			return typeInfo{name: typeName}
		}
		if startContainer, ok := recursiveFields[tfName]; ok {
			if inRecursiveField == "" {
				if startContainer == container {
					// Begin recursive type.
					inRecursiveField = tfName
				}
			} else {
				if tfName == inRecursiveField {
					// End recursive type.
					return typeInfo{name: parentTypeName}
				}
				// TODO: while building inRecursiveField, we encountered a new
				// recursive field (tfName). This is an artificial limitation of
				// this code which can be removed.
				panic("recursive types cannot contain other recursive types")
			}
		}
		typeName := parentTypeName + fieldName
		f.newStructDef(typeName, t.Schema, container, inRecursiveField)
		return typeInfo{name: typeName}
	default:
		panic(fmt.Sprintf("unhandled meta type: %T", fieldMeta))
	}
}

type structDef struct {
	IsConverter  bool
	ResourceName string
	TypeName     string
	Fields       []*field

	container *schema.Schema
}

type field struct {
	Name     string
	TypeName string
	Tag      intschematag.Tag

	Comment []string // Optional.
}

func (f *field) sortScore() int {
	switch f.Tag.TFName {
	case "name":
		return 0
	case "slug":
		return 1
	}
	if isTFID(f.Tag.TFName) {
		return 2
	}
	if !f.Tag.Optional {
		return 3
	}

	// This is a hack to ensure that "override" fields always
	// come last in the serialized HCL output.
	if f.Tag.TFName == "override" {
		return 5
	}

	return 4
}

func isTFID(tfName string) bool {
	isTFID, ok := fieldIsTFID[tfName]
	if !ok && strings.HasSuffix(tfName, "_id") {
		panic(fmt.Sprintf(
			"field %q looks like tfid but is not configured in fieldIsTFID",
			tfName))
	}
	return isTFID
}

func addFileField(tfName string, s *schema.Schema) bool {
	return fieldAddFile[tfName] && s.Type == schema.TypeString
}

func upperCamelCase(snakeCase string) string {
	b := &strings.Builder{}
	startToken := true
	for _, c := range snakeCase {
		if startToken {
			b.WriteRune(unicode.ToUpper(c))
			startToken = false
			continue
		}
		if c == '_' {
			startToken = true
			continue
		}
		b.WriteRune(c)
	}
	return b.String()
}

func reverse(defs []*structDef) {
	for i := 0; i < len(defs)/2; i++ {
		j := len(defs) - 1 - i
		defs[i], defs[j] = defs[j], defs[i]
	}
}

func sortedKeys[V any](m map[string]V) []string {
	keys := maps.Keys(m)
	sort.Strings(keys)
	return keys
}

var fileTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package intschema

import (
{{range .Imports -}}
{{printf "%q" .}}
{{end -}}
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

var _ tfid.ID // Always use tfid for simplified import generation.

{{ $schemaType := .Type }}
{{range .Structs}}
type {{.TypeName}} struct {
	{{range .Fields -}}
	{{if .Comment}}
	{{end -}}
	{{range .Comment -}}
	// {{.}}
	{{end -}}
	{{.Name}} {{.TypeName}} ` + "`" + `{{.Tag.Marshal}}` + "`" + `
	{{end}}
}
{{if .IsConverter}}
func (o *{{.TypeName}}) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.{{.TypeName}}, d, o)
}

func (o *{{.TypeName}}) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *{{.TypeName}}) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.{{ $schemaType.MarshalAddFuncName }}("{{.ResourceName}}", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *{{.TypeName}}) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		{{ if $schemaType.Datasource -}}
		Datasource: true,
		{{ end -}}
		Type: "{{.ResourceName}}",
		ID: o.HCLID,
	}.AsID()
}
{{end}}
{{end}}
`))
