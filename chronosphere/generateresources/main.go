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

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/registry"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	v1 := api{
		Package:       "configv1",
		SwaggerPrefix: "Configv1",
		GoPrefix:      "generated",
		Client:        "ConfigV1",
	}
	unstable := api{
		Package:       "configunstable",
		SwaggerPrefix: "Configunstable",
		GoPrefix:      "generatedUnstable",
		Client:        "ConfigUnstable",
	}

	var entityTypes []entityType
	for _, e := range registry.StandardEntities(registry.V1) {
		entityTypes = append(entityTypes, newEntityType(v1, e))
	}

	includesUnstable := false
	for _, e := range registry.StandardEntities(registry.Unstable) {
		entityTypes = append(entityTypes, newEntityType(unstable, e))
		includesUnstable = true
	}

	b := &bytes.Buffer{}
	if err := fileTemplate.Execute(b, data{
		EntityTypes:      entityTypes,
		IncludesUnstable: includesUnstable,
	}); err != nil {
		return err
	}
	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	return os.WriteFile("generated_resources.gen.go", formatted, 0o644)
}

type api struct {
	Package       string
	SwaggerPrefix string
	GoPrefix      string
	Client        string
}

type entityType struct {
	API         api
	GoType      string
	SwaggerType string
	// SwaggerModel represents the underlying model to use. Will normally be the
	// same as SwaggerType
	SwaggerModel         string
	FieldName            string
	SwaggerClient        string
	SwaggerClientPackage string
	DryRun               bool
	UpdateUnsupported    bool
}

func newEntityType(a api, r registry.Resource) entityType {
	et := entityType{
		API:                  a,
		GoType:               fmt.Sprintf("%s%s", a.GoPrefix, r.Entity),
		SwaggerType:          r.Entity,
		SwaggerModel:         r.Entity,
		SwaggerClient:        fmt.Sprintf("%s.%s", a.Client, r.Entity),
		SwaggerClientPackage: strcase.ToSnake(r.Entity),
		DryRun:               r.DryRun,
		UpdateUnsupported:    r.UpdateUnsupported,
	}

	// special case for classic dashboards
	if r.Name == "classic_dashboard" {
		et.GoType = fmt.Sprintf("%s%s", a.GoPrefix, "ClassicDashboard")
	}
	return et
}

type data struct {
	EntityTypes      []entityType
	IncludesUnstable bool
}

var fileTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package chronosphere

import (
	"context"
	{{ range .EntityTypes }}
	{{ if not .DryRun }}
	"fmt"
	{{ break }}
	{{ end }}
	{{ end }}

	configv1models "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	{{- if .IncludesUnstable }}
	configunstablemodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	{{- end }}
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/apiclients"
	{{- range .EntityTypes }}
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/{{.API.Package}}/client/{{.SwaggerClientPackage}}"
	{{- end }}
)

{{ range .EntityTypes }}

	type {{.GoType}} struct{}

func ({{.GoType}}) slugOf(m *{{.API.Package}}models.{{.API.SwaggerPrefix}}{{.SwaggerModel}}) string {
	return m.Slug
}

func ({{.GoType}}) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *{{.API.Package}}models.{{.API.SwaggerPrefix}}{{.SwaggerModel}},
    dryRun bool,
) (string, error) {
	{{ if not .DryRun -}}
	if dryRun {
		return "", fmt.Errorf("dry run not supported for this entity type")
	}
	{{ end -}}
	req := &{{.SwaggerClientPackage}}.Create{{.SwaggerType}}Params{
		Context: ctx,
		Body: &{{.API.Package}}models.{{.API.SwaggerPrefix}}Create{{.SwaggerType}}Request{
			{{.SwaggerType}}: m,
			{{ if .DryRun }} DryRun: dryRun, {{ end }}
		},
	}
	resp, err := clients.{{.SwaggerClient}}.Create{{.SwaggerType}}(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.{{.SwaggerType}}
	if e == nil {
	  return "", nil
	}
	return e.Slug, nil
}

func ({{.GoType}}) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*{{.API.Package}}models.{{.API.SwaggerPrefix}}{{.SwaggerModel}}, error) {
	req := &{{.SwaggerClientPackage}}.Read{{.SwaggerType}}Params{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.{{.SwaggerClient}}.Read{{.SwaggerType}}(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.{{.SwaggerType}}, nil
}

{{ if not .UpdateUnsupported -}}
func ({{.GoType}}) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *{{.API.Package}}models.{{.API.SwaggerPrefix}}{{.SwaggerModel}},
	params updateParams,
) error {
	{{ if not .DryRun -}}
	if params.dryRun {
		return fmt.Errorf("dry run not supported for this entity type")
	}
	{{ end -}}
	req := &{{.SwaggerClientPackage}}.Update{{.SwaggerType}}Params{
		Context: ctx,
		Slug:    m.Slug,
		Body: {{.SwaggerClientPackage}}.Update{{.SwaggerType}}Body{
			{{.SwaggerType}}: m,
			CreateIfMissing: params.createIfMissing,
			{{ if .DryRun }} DryRun: params.dryRun, {{ end }}
		},
	}
	_, err := clients.{{.SwaggerClient}}.Update{{.SwaggerType}}(req)
	return err
}
{{ end -}}

func ({{.GoType}}) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &{{.SwaggerClientPackage}}.Delete{{.SwaggerType}}Params{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.{{.SwaggerClient}}.Delete{{.SwaggerType}}(req)
	return err
}
{{ end -}}
`))
