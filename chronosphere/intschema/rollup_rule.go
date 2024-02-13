// Code generated by go generate; DO NOT EDIT.
package intschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/convertintschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
)

var _ tfid.ID // Always use tfid for simplified import generation.

type RollupRule struct {
	Name            string                     `intschema:"name"`
	Slug            string                     `intschema:"slug"`
	BucketId        tfid.ID                    `intschema:"bucket_id,optional"`
	Filter          string                     `intschema:"filter"`
	MetricType      string                     `intschema:"metric_type"`
	Aggregation     string                     `intschema:"aggregation,optional"`
	DropRaw         bool                       `intschema:"drop_raw,optional,default:false"`
	ExcludeBy       []string                   `intschema:"exclude_by,optional"`
	GroupBy         []string                   `intschema:"group_by,optional"`
	Interval        string                     `intschema:"interval,optional"`
	MetricTypeTag   bool                       `intschema:"metric_type_tag,optional,default:false"`
	Mode            string                     `intschema:"mode,optional"`
	NewMetric       string                     `intschema:"new_metric,optional"`
	Permissive      bool                       `intschema:"permissive,optional,default:false"`
	StoragePolicies *RollupRuleStoragePolicies `intschema:"storage_policies,optional,computed,list_encoded_object"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *RollupRule) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.RollupRule, d, o)
}

func (o *RollupRule) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *RollupRule) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_rollup_rule", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *RollupRule) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_rollup_rule",
		ID:   o.HCLID,
	}.AsID()
}

type RollupRuleStoragePolicies struct {
	Resolution string `intschema:"resolution"`
	Retention  string `intschema:"retention"`
}
