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

type TraceMetricsRule struct {
	Name                    string                      `intschema:"name"`
	Slug                    string                      `intschema:"slug,optional,computed"`
	MetricName              string                      `intschema:"metric_name"`
	TraceFilter             TraceMetricsRuleTraceFilter `intschema:"trace_filter,list_encoded_object"`
	GroupBy                 []TraceMetricsRuleGroupBy   `intschema:"group_by,optional"`
	HistogramBucketsSeconds []float64                   `intschema:"histogram_buckets_seconds,optional"`
	MetricLabels            map[string]string           `intschema:"metric_labels,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *TraceMetricsRule) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.TraceMetricsRule, d, o)
}

func (o *TraceMetricsRule) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *TraceMetricsRule) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_trace_metrics_rule", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *TraceMetricsRule) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_trace_metrics_rule",
		ID:   o.HCLID,
	}.AsID()
}

type TraceMetricsRuleTraceFilter struct {
	Span  []TraceMetricsRuleTraceFilterSpan `intschema:"span,optional"`
	Trace *TraceMetricsRuleTraceFilterTrace `intschema:"trace,optional,list_encoded_object"`
}

type TraceMetricsRuleTraceFilterTrace struct {
	Duration *TraceMetricsDurationFilter `intschema:"duration,optional,list_encoded_object"`
	Error    *TraceMetricsBoolFilter     `intschema:"error,optional,list_encoded_object"`
}

type TraceMetricsRuleTraceFilterSpan struct {
	Duration        *TraceMetricsDurationFilter               `intschema:"duration,optional,list_encoded_object"`
	Error           *TraceMetricsBoolFilter                   `intschema:"error,optional,list_encoded_object"`
	MatchType       string                                    `intschema:"match_type,optional,default:include"`
	Operation       *TraceMetricsStringFilter                 `intschema:"operation,optional,list_encoded_object"`
	ParentOperation *TraceMetricsStringFilter                 `intschema:"parent_operation,optional,list_encoded_object"`
	ParentService   *TraceMetricsStringFilter                 `intschema:"parent_service,optional,list_encoded_object"`
	Service         *TraceMetricsStringFilter                 `intschema:"service,optional,list_encoded_object"`
	SpanCount       *TraceMetricsRuleTraceFilterSpanSpanCount `intschema:"span_count,optional,list_encoded_object"`
	Tag             []TraceMetricsRuleTraceFilterSpanTag      `intschema:"tag,optional"`
}

type TraceMetricsRuleTraceFilterSpanTag struct {
	Key          string                           `intschema:"key"`
	NumericValue *TraceMetricsNumericFilterSchema `intschema:"numeric_value,optional,list_encoded_object"`
	Value        *TraceMetricsStringFilter        `intschema:"value,optional,list_encoded_object"`
}

type TraceMetricsRuleTraceFilterSpanSpanCount struct {
	Max int64 `intschema:"max,optional,default:0"`
	Min int64 `intschema:"min,optional,default:0"`
}

type TraceMetricsRuleGroupBy struct {
	Key   TraceMetricsRuleGroupByKey `intschema:"key,list_encoded_object"`
	Label string                     `intschema:"label"`
}

type TraceMetricsRuleGroupByKey struct {
	Type     string `intschema:"type"`
	NamedKey string `intschema:"named_key,optional"`
}
