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

type RecordingRule struct {
	Name           string            `intschema:"name"`
	Slug           string            `intschema:"slug,optional,computed"`
	BucketId       tfid.ID           `intschema:"bucket_id,optional"`
	ExecutionGroup tfid.ID           `intschema:"execution_group,optional"`
	Expr           string            `intschema:"expr"`
	Interval       string            `intschema:"interval,optional"`
	Labels         map[string]string `intschema:"labels,optional"`
	MetricName     string            `intschema:"metric_name,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *RecordingRule) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.RecordingRule, d, o)
}

func (o *RecordingRule) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *RecordingRule) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_recording_rule", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *RecordingRule) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_recording_rule",
		ID:   o.HCLID,
	}.AsID()
}
