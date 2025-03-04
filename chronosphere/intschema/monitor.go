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

type Monitor struct {
	Name                 string                  `intschema:"name"`
	Slug                 string                  `intschema:"slug,optional,computed"`
	BucketId             tfid.ID                 `intschema:"bucket_id,optional"`
	CollectionId         tfid.ID                 `intschema:"collection_id,optional"`
	NotificationPolicyId tfid.ID                 `intschema:"notification_policy_id,optional"`
	Query                MonitorQuery            `intschema:"query,list_encoded_object"`
	SeriesConditions     MonitorSeriesConditions `intschema:"series_conditions,list_encoded_object"`
	Annotations          map[string]string       `intschema:"annotations,optional"`
	Interval             string                  `intschema:"interval,optional"`
	Labels               map[string]string       `intschema:"labels,optional"`
	Schedule             *MonitorSchedule        `intschema:"schedule,optional,list_encoded_object"`
	SignalGrouping       *SignalGrouping         `intschema:"signal_grouping,optional,list_encoded_object"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *Monitor) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.Monitor, d, o)
}

func (o *Monitor) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *Monitor) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_monitor", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *Monitor) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_monitor",
		ID:   o.HCLID,
	}.AsID()
}

type MonitorSeriesConditions struct {
	Condition []MonitorSeriesCondition          `intschema:"condition"`
	Override  []MonitorSeriesConditionsOverride `intschema:"override,optional"`
}

type MonitorSeriesConditionsOverride struct {
	Condition    []MonitorSeriesCondition `intschema:"condition"`
	LabelMatcher []Matcher                `intschema:"label_matcher"`
}

type MonitorSchedule struct {
	Timezone string                 `intschema:"timezone"`
	Range    []MonitorScheduleRange `intschema:"range,optional"`
}

type MonitorScheduleRange struct {
	Day   string `intschema:"day"`
	End   string `intschema:"end"`
	Start string `intschema:"start"`
}

type MonitorQuery struct {
	GraphiteExpr   string `intschema:"graphite_expr,optional"`
	LoggingExpr    string `intschema:"logging_expr,optional"`
	PrometheusExpr string `intschema:"prometheus_expr,optional"`
}
