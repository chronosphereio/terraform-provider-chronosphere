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

package chronosphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/prettyenum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceTraceMetricsRule() *schema.Resource {
	r := newGenericResource[
		*models.Configv1TraceMetricsRule,
		intschema.TraceMetricsRule,
		*intschema.TraceMetricsRule,
	](
		"trace_metrics_rule",
		traceMetricsRuleConverter{},
		generatedTraceMetricsRule{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.TraceMetricsRule,
		CustomizeDiff: r.ValidateDryRun(&TraceMetricsRuleDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// TraceMetricsRuleDryRunCount tracks how many times dry run is run during validation for testing.
var TraceMetricsRuleDryRunCount atomic.Int64

type traceMetricsRuleConverter struct{}

//nolint:unused
func (traceMetricsRuleConverter) toModel(r *intschema.TraceMetricsRule) (*models.Configv1TraceMetricsRule, error) {
	return &models.Configv1TraceMetricsRule{
		Slug:                 r.Slug,
		Name:                 r.Name,
		MetricName:           r.MetricName,
		TraceFilter:          traceMetricsRuleSearchFilterToModel(r.TraceFilter),
		HistogramBucketsSecs: r.HistogramBucketsSeconds,
		MetricLabels:         r.MetricLabels,
		GroupBy:              expandTraceMetricsRuleGroupBy(r.GroupBy),
	}, nil
}

func (traceMetricsRuleConverter) fromModel(r *models.Configv1TraceMetricsRule) (*intschema.TraceMetricsRule, error) {
	return &intschema.TraceMetricsRule{
		Name:                    r.Name,
		Slug:                    r.Slug,
		MetricName:              r.MetricName,
		TraceFilter:             traceMetricsRuleSearchFilterFromModel(r.TraceFilter),
		HistogramBucketsSeconds: r.HistogramBucketsSecs,
		MetricLabels:            r.MetricLabels,
		GroupBy:                 sliceutil.Map(r.GroupBy, traceMetricsRuleGroupByFromModel),
	}, nil
}

func TraceMetricsRuleFromModel(r *models.Configv1TraceMetricsRule) (*intschema.TraceMetricsRule, error) {
	return traceMetricsRuleConverter{}.fromModel(r)
}

func traceMetricsRuleSearchFilterFromModel(f *models.Configv1TraceSearchFilter) intschema.TraceMetricsRuleTraceFilter {
	if f == nil {
		return intschema.TraceMetricsRuleTraceFilter{}
	}

	return intschema.TraceMetricsRuleTraceFilter{
		Trace: traceMetricsRuleFilterFromModel(f.Trace),
		Span:  traceMetricsRuleFilterSpanFromModel(f.Span),
	}
}

func traceMetricsRuleFilterFromModel(f *models.TraceSearchFilterTraceFilter) *intschema.TraceMetricsRuleTraceFilterTrace {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsRuleTraceFilterTrace{
		Duration: traceMetricsRuleFilterDurationFromModel(f.Duration),
		Error:    traceMetricsRuleFilterBoolFromModel(f.Error),
	}
}

func traceMetricsRuleFilterBoolFromModel(f *models.TraceSearchFilterBoolFilter) *intschema.TraceMetricsBoolFilter {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsBoolFilter{
		Value: f.Value,
	}
}

func traceMetricsRuleFilterDurationFromModel(f *models.TraceSearchFilterDurationFilter) *intschema.TraceMetricsDurationFilter {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsDurationFilter{
		MaxSeconds: float64(f.MaxSecs),
		MinSeconds: float64(f.MinSecs),
	}
}

func traceMetricsRuleFilterSpanFromModel(f []*models.TraceSearchFilterSpanFilter) []intschema.TraceMetricsRuleTraceFilterSpan {
	if len(f) == 0 {
		return nil
	}

	result := make([]intschema.TraceMetricsRuleTraceFilterSpan, 0, len(f))
	for _, s := range f {
		span := intschema.TraceMetricsRuleTraceFilterSpan{
			MatchType:       string(prettyenum.SpanFilterMatchTypeFromModel(s.MatchType)),
			Service:         traceMetricsRuleStringFilterFromModel(s.Service),
			Operation:       traceMetricsRuleStringFilterFromModel(s.Operation),
			ParentService:   traceMetricsRuleStringFilterFromModel(s.ParentService),
			ParentOperation: traceMetricsRuleStringFilterFromModel(s.ParentOperation),
			Duration:        traceMetricsRulesDurationFilterFromModel(s.Duration),
			Error:           traceMetricsRulesBoolFilterFromModel(s.Error),
			Tag:             traceMetricsRulesTagFiltersFromModel(s.Tags),
			SpanCount:       traceMetricsRulesCountFilterFromModel(s.SpanCount),
		}
		result = append(result, span)
	}
	return result
}

func traceMetricsRulesDurationFilterFromModel(f *models.TraceSearchFilterDurationFilter) *intschema.TraceMetricsDurationFilter {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsDurationFilter{
		MaxSeconds: float64(f.MaxSecs),
		MinSeconds: float64(f.MinSecs),
	}
}

func traceMetricsRulesBoolFilterFromModel(f *models.TraceSearchFilterBoolFilter) *intschema.TraceMetricsBoolFilter {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsBoolFilter{
		Value: f.Value,
	}
}

func traceMetricsRulesTagFiltersFromModel(f []*models.TraceSearchFilterTagFilter) []intschema.TraceMetricsRuleTraceFilterSpanTag {
	if len(f) == 0 {
		return nil
	}

	result := make([]intschema.TraceMetricsRuleTraceFilterSpanTag, 0, len(f))
	for _, t := range f {
		tag := intschema.TraceMetricsRuleTraceFilterSpanTag{
			Key:          t.Key,
			NumericValue: traceMetricsRuleNumericFilterSchemaFromModel(t.NumericValue),
			Value:        traceMetricsRuleStringFilterFromModel(t.Value),
		}
		result = append(result, tag)
	}
	return result
}

func traceMetricsRuleNumericFilterSchemaFromModel(f *models.TraceSearchFilterNumericFilter) *intschema.TraceMetricsNumericFilterSchema {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsNumericFilterSchema{
		Comparison: string(f.Comparison),
		Value:      f.Value,
	}
}

func traceMetricsRuleStringFilterFromModel(f *models.TraceSearchFilterStringFilter) *intschema.TraceMetricsStringFilter {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsStringFilter{
		Match: string(prettyenum.StringFilterMatchTypeFromModel(f.Match)),
		Value: f.Value,
	}
}

func traceMetricsRulesCountFilterFromModel(f *models.TraceSearchFilterCountFilter) *intschema.TraceMetricsRuleTraceFilterSpanSpanCount {
	if f == nil {
		return nil
	}

	return &intschema.TraceMetricsRuleTraceFilterSpanSpanCount{
		Max: int64(f.Max),
		Min: int64(f.Min),
	}
}

func traceMetricsRuleSearchFilterToModel(f intschema.TraceMetricsRuleTraceFilter) *models.Configv1TraceSearchFilter {
	return &models.Configv1TraceSearchFilter{
		Trace: traceMetricsRuleFilterToModel(f.Trace),
		Span:  spanFiltersToModel(f.Span),
	}
}

func traceMetricsRuleFilterToModel(f *intschema.TraceMetricsRuleTraceFilterTrace) *models.TraceSearchFilterTraceFilter {
	if f == nil {
		return nil
	}

	return &models.TraceSearchFilterTraceFilter{
		Duration: traceMetricsRuleDurationFilterToModel(f.Duration),
		Error:    traceMetricsRuleBoolFilterToModel(f.Error),
	}
}

func traceMetricsRuleDurationFilterToModel(f *intschema.TraceMetricsDurationFilter) *models.TraceSearchFilterDurationFilter {
	if f == nil {
		return nil
	}

	return &models.TraceSearchFilterDurationFilter{
		MaxSecs: f.MaxSeconds,
		MinSecs: f.MinSeconds,
	}
}

func traceMetricsRuleBoolFilterToModel(f *intschema.TraceMetricsBoolFilter) *models.TraceSearchFilterBoolFilter {
	if f == nil {
		return nil
	}

	return &models.TraceSearchFilterBoolFilter{
		Value: f.Value,
	}
}

func spanFiltersToModel(f []intschema.TraceMetricsRuleTraceFilterSpan) []*models.TraceSearchFilterSpanFilter {
	if len(f) == 0 {
		return nil
	}

	result := make([]*models.TraceSearchFilterSpanFilter, 0, len(f))
	for _, s := range f {
		matchType, err := prettyenum.NewSpanFilterMatchType(s.MatchType)
		if err != nil {
			panic(err)
		}

		span := &models.TraceSearchFilterSpanFilter{
			MatchType:       matchType.Model(),
			Service:         traceMetricsRuleStringFilterToModel(s.Service),
			Operation:       traceMetricsRuleStringFilterToModel(s.Operation),
			ParentService:   traceMetricsRuleStringFilterToModel(s.ParentService),
			ParentOperation: traceMetricsRuleStringFilterToModel(s.ParentOperation),
			Duration:        traceMetricsRuleDurationFilterToModel(s.Duration),
			Error:           traceMetricsRuleBoolFilterToModel(s.Error),
			Tags:            traceMetricsRuleTagFiltersToModel(s.Tag),
			SpanCount:       traceMetricsRuleCountFilterToModel(s.SpanCount),
		}
		result = append(result, span)
	}
	return result
}

func traceMetricsRuleStringFilterToModel(f *intschema.TraceMetricsStringFilter) *models.TraceSearchFilterStringFilter {
	if f == nil {
		return nil
	}

	match, err := prettyenum.NewStringFilterMatchType(f.Match)
	if err != nil {
		panic(err)
	}

	return &models.TraceSearchFilterStringFilter{
		Match: match.Model(),
		Value: f.Value,
	}
}

func traceMetricsRuleTagFiltersToModel(f []intschema.TraceMetricsRuleTraceFilterSpanTag) []*models.TraceSearchFilterTagFilter {
	if len(f) == 0 {
		return nil
	}

	result := make([]*models.TraceSearchFilterTagFilter, 0, len(f))
	for _, t := range f {
		tag := &models.TraceSearchFilterTagFilter{
			Key:          t.Key,
			NumericValue: traceMetricsRuleNumericFilterSchemaToModel(t.NumericValue),
			Value:        traceMetricsRuleStringFilterSchemaToModel(t.Value),
		}
		result = append(result, tag)
	}
	return result
}

func traceMetricsRuleNumericFilterSchemaToModel(f *intschema.TraceMetricsNumericFilterSchema) *models.TraceSearchFilterNumericFilter {
	if f == nil {
		return nil
	}

	return &models.TraceSearchFilterNumericFilter{
		Comparison: prettyenum.NumericFilterComparisonType(f.Comparison).Model(),
		Value:      f.Value,
	}
}

func traceMetricsRuleStringFilterSchemaToModel(f *intschema.TraceMetricsStringFilter) *models.TraceSearchFilterStringFilter {
	if f == nil {
		return nil
	}

	match, err := prettyenum.NewStringFilterMatchType(f.Match)
	if err != nil {
		panic(err)
	}

	return &models.TraceSearchFilterStringFilter{
		Match: match.Model(),
		Value: f.Value,
	}
}

func traceMetricsRuleCountFilterToModel(f *intschema.TraceMetricsRuleTraceFilterSpanSpanCount) *models.TraceSearchFilterCountFilter {
	if f == nil {
		return nil
	}

	return &models.TraceSearchFilterCountFilter{
		Max: int32(f.Max),
		Min: int32(f.Min),
	}
}

func expandTraceMetricsRuleGroupBy(d []intschema.TraceMetricsRuleGroupBy) []*models.TraceMetricsRuleGroupBy {
	result := make([]*models.TraceMetricsRuleGroupBy, 0, len(d))
	for _, obj := range d {
		result = append(result, &models.TraceMetricsRuleGroupBy{
			Key: &models.GroupByGroupByKey{
				NamedKey: obj.Key.NamedKey,
				Type:     models.GroupByKeyGroupByKeyType(obj.Key.Type),
			},
			Label: obj.Label,
		})
	}
	return result
}

func traceMetricsRuleGroupByFromModel(m *models.TraceMetricsRuleGroupBy) intschema.TraceMetricsRuleGroupBy {
	if m == nil || m.Key == nil {
		return intschema.TraceMetricsRuleGroupBy{}
	}

	return intschema.TraceMetricsRuleGroupBy{
		Key: intschema.TraceMetricsRuleGroupByKey{
			NamedKey: m.Key.NamedKey,
			Type:     string(m.Key.Type),
		},
		Label: m.Label,
	}
}
