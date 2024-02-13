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

type DerivedMetric struct {
	Name        string                 `intschema:"name"`
	Slug        string                 `intschema:"slug,optional,computed"`
	MetricName  string                 `intschema:"metric_name"`
	Queries     []DerivedMetricQueries `intschema:"queries"`
	Description string                 `intschema:"description,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *DerivedMetric) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.DerivedMetric, d, o)
}

func (o *DerivedMetric) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *DerivedMetric) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_derived_metric", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *DerivedMetric) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_derived_metric",
		ID:   o.HCLID,
	}.AsID()
}

type DerivedMetricQueries struct {
	Query    DerivedMetricQueriesQuery     `intschema:"query,list_encoded_object"`
	Selector *DerivedMetricQueriesSelector `intschema:"selector,optional,list_encoded_object"`
}

type DerivedMetricQueriesSelector struct {
	Labels map[string]string `intschema:"labels,optional"`
}

type DerivedMetricQueriesQuery struct {
	Expr      string                               `intschema:"expr"`
	Variables []DerivedMetricQueriesQueryVariables `intschema:"variables,optional"`
}

type DerivedMetricQueriesQueryVariables struct {
	Name            string `intschema:"name"`
	DefaultSelector string `intschema:"default_selector"`
}
