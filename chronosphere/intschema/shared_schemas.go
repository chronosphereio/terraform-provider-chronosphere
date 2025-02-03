// Code generated by go generate; DO NOT EDIT.
package intschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

var _ tfid.ID // Always use tfid for simplified import generation.

type LogAllocationConfigSchema struct {
	PercentOfLicense float64 `intschema:"percent_of_license"`
}

type LogPrioritiesSchema struct {
	HighPriorityFilter []LogSearchFilterSchema `intschema:"high_priority_filter,optional"`
	LowPriorityFilter  []LogSearchFilterSchema `intschema:"low_priority_filter,optional"`
}

type LogSearchFilterSchema struct {
	Query string `intschema:"query"`
}

type Matcher struct {
	Name  string `intschema:"name"`
	Type  string `intschema:"type"`
	Value string `intschema:"value"`
}

type MonitorSeriesCondition struct {
	Op             string  `intschema:"op"`
	Severity       string  `intschema:"severity"`
	ResolveSustain string  `intschema:"resolve_sustain,optional"`
	Sustain        string  `intschema:"sustain,optional"`
	Value          float64 `intschema:"value,optional,default:0"`
}

type NotificationRoute struct {
	Severity       string                    `intschema:"severity"`
	GroupBy        *NotificationRouteGroupBy `intschema:"group_by,optional,list_encoded_object"`
	Notifiers      []tfid.ID                 `intschema:"notifiers,optional"`
	RepeatInterval string                    `intschema:"repeat_interval,optional"`
}

type NotificationRouteGroupBy struct {
	LabelNames []string `intschema:"label_names,optional"`
}

type ResourcePoolAllocationSchema struct {
	FixedValue       []ResourcePoolAllocationSchemaFixedValue `intschema:"fixed_value,optional"`
	PercentOfLicense float64                                  `intschema:"percent_of_license,optional"`
}

type ResourcePoolAllocationSchemaFixedValue struct {
	License string `intschema:"license"`
	Value   int64  `intschema:"value"`
}

type ResourcePoolPrioritiesSchema struct {
	HighPriorityMatchRules []string `intschema:"high_priority_match_rules,optional"`
	LowPriorityMatchRules  []string `intschema:"low_priority_match_rules,optional"`
}

type ResourcePoolsConfigPool struct {
	Name       string                        `intschema:"name"`
	Allocation *ResourcePoolAllocationSchema `intschema:"allocation,optional,list_encoded_object"`
	MatchRule  string                        `intschema:"match_rule,optional"`
	MatchRules []string                      `intschema:"match_rules,optional"`
	Priorities *ResourcePoolPrioritiesSchema `intschema:"priorities,optional,list_encoded_object"`
}

type SLOAdditionalPromQLFilters struct {
	Name  string `intschema:"name"`
	Type  string `intschema:"type"`
	Value string `intschema:"value"`
}

type SignalGrouping struct {
	LabelNames      []string `intschema:"label_names,optional"`
	SignalPerSeries bool     `intschema:"signal_per_series,optional"`
}

type TraceBoolFilter struct {
	Value bool `intschema:"value"`
}

type TraceDurationFilter struct {
	MaxSecs float64 `intschema:"max_secs,optional"`
	MinSecs float64 `intschema:"min_secs,optional,default:0"`
}

type TraceFilter struct {
	Duration *TraceDurationFilter `intschema:"duration,optional,list_encoded_object"`
	Error    *TraceBoolFilter     `intschema:"error,optional,list_encoded_object"`
}

type TraceNumericFilter struct {
	Comparison string  `intschema:"comparison"`
	Value      float64 `intschema:"value"`
}

type TraceSearchFilter struct {
	Span  []TraceSpanFilter `intschema:"span,optional"`
	Trace *TraceFilter      `intschema:"trace,optional,list_encoded_object"`
}

type TraceSpanCountFilter struct {
	Max int64 `intschema:"max,optional,default:0"`
	Min int64 `intschema:"min,optional,default:0"`
}

type TraceSpanFilter struct {
	Duration        *TraceDurationFilter  `intschema:"duration,optional,list_encoded_object"`
	Error           *TraceBoolFilter      `intschema:"error,optional,list_encoded_object"`
	IsRootSpan      *TraceBoolFilter      `intschema:"is_root_span,optional,list_encoded_object"`
	MatchType       string                `intschema:"match_type,optional,default:include"`
	Operation       *TraceStringFilter    `intschema:"operation,optional,list_encoded_object"`
	ParentOperation *TraceStringFilter    `intschema:"parent_operation,optional,list_encoded_object"`
	ParentService   *TraceStringFilter    `intschema:"parent_service,optional,list_encoded_object"`
	Service         *TraceStringFilter    `intschema:"service,optional,list_encoded_object"`
	SpanCount       *TraceSpanCountFilter `intschema:"span_count,optional,list_encoded_object"`
	Tag             []TraceTagFilter      `intschema:"tag,optional"`
}

type TraceStringFilter struct {
	InValues []string `intschema:"in_values,optional"`
	Match    string   `intschema:"match,optional,default:exact"`
	Value    string   `intschema:"value,optional"`
}

type TraceTagFilter struct {
	Key          string              `intschema:"key,optional"`
	NumericValue *TraceNumericFilter `intschema:"numeric_value,optional,list_encoded_object"`
	Value        *TraceStringFilter  `intschema:"value,optional,list_encoded_object"`
}

type ValueMappings struct {
	SourceValueGlobs []string `intschema:"source_value_globs"`
	TargetValue      string   `intschema:"target_value"`
}
