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
	"github.com/pkg/errors"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/prettyenum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
)

func stringFilterToModel(
	f *intschema.TraceStringFilter,
) *models.TraceSearchFilterStringFilter {
	if f == nil {
		return nil
	}

	match, err := prettyenum.NewStringFilterMatchType(f.Match)
	if err != nil {
		panic(err)
	}

	return &models.TraceSearchFilterStringFilter{
		Match:    match.Model(),
		Value:    f.Value,
		InValues: f.InValues,
	}
}

func stringFilterFromModel(
	f *models.TraceSearchFilterStringFilter,
) *intschema.TraceStringFilter {
	if f == nil {
		return nil
	}

	return &intschema.TraceStringFilter{
		Match:    string(prettyenum.StringFilterMatchTypeFromModel(f.Match)),
		Value:    f.Value,
		InValues: f.InValues,
	}
}

func countFilterFromModel(
	c *models.TraceSearchFilterCountFilter,
) *intschema.TraceSpanCountFilter {
	if c == nil {
		return nil
	}
	return &intschema.TraceSpanCountFilter{
		Max: int64(c.Max),
		Min: int64(c.Min),
	}
}

func countFilterToModel(
	c *intschema.TraceSpanCountFilter,
) *models.TraceSearchFilterCountFilter {
	if c == nil {
		return nil
	}
	return &models.TraceSearchFilterCountFilter{
		Max: int32(c.Max),
		Min: int32(c.Min),
	}
}

func boolFilterFromModel(
	b *models.TraceSearchFilterBoolFilter,
) *intschema.TraceBoolFilter {
	if b == nil {
		return nil
	}
	return &intschema.TraceBoolFilter{
		Value: b.Value,
	}
}

func boolFilterToModel(
	b *intschema.TraceBoolFilter,
) *models.TraceSearchFilterBoolFilter {
	if b == nil {
		return nil
	}
	return &models.TraceSearchFilterBoolFilter{
		Value: b.Value,
	}
}

func numericFilterFromModel(
	s *models.TraceSearchFilterNumericFilter,
) *intschema.TraceNumericFilter {
	if s == nil {
		return nil
	}
	return &intschema.TraceNumericFilter{
		Comparison: string(s.Comparison),
		Value:      s.Value,
	}
}

func numericFilterToModel(
	s *intschema.TraceNumericFilter,
) *models.TraceSearchFilterNumericFilter {
	if s == nil {
		return nil
	}
	return &models.TraceSearchFilterNumericFilter{
		Comparison: enum.NumericFilterComparisonType.V1(s.Comparison),
		Value:      s.Value,
	}
}

func durationFilterFromModel(
	d *models.TraceSearchFilterDurationFilter,
) *intschema.TraceDurationFilter {
	if d == nil {
		return nil
	}
	return &intschema.TraceDurationFilter{
		MinSecs: d.MinSecs, // Read into preferred min_secs field (not min_seconds).
		MaxSecs: d.MaxSecs, // Read into preferred max_secs field (not max_seconds).
	}
}

func durationFilterToModel(
	d *intschema.TraceDurationFilter,
) (*models.TraceSearchFilterDurationFilter, error) {
	if d == nil {
		return nil, nil
	}

	if d.MaxSeconds != 0 || d.MinSeconds != 0 {
		return nil, errors.New("replace all usage of \"min_seconds\" and \"max_seconds\" with \"min_secs\" and \"max_secs\"")
	}

	return &models.TraceSearchFilterDurationFilter{
		MaxSecs: d.MaxSecs,
		MinSecs: d.MinSecs,
	}, nil
}

func tagFilterFromModel(
	t *models.TraceSearchFilterTagFilter,
) intschema.TraceTagFilter {
	return intschema.TraceTagFilter{
		Key:          t.Key,
		Value:        stringFilterFromModel(t.Value),
		NumericValue: numericFilterFromModel(t.NumericValue),
	}
}

func tagFilterToModel(
	t intschema.TraceTagFilter,
) *models.TraceSearchFilterTagFilter {
	return &models.TraceSearchFilterTagFilter{
		Key:          t.Key,
		Value:        stringFilterToModel(t.Value),
		NumericValue: numericFilterToModel(t.NumericValue),
	}
}

func traceSearchFilterFromModel(
	f *models.Configv1TraceSearchFilter,
) intschema.TraceSearchFilter {
	return intschema.TraceSearchFilter{
		Span:  sliceutil.Map(f.Span, spanFilterFromModel),
		Trace: traceFilterFromModel(f.Trace),
	}
}

func traceSearchFilterToModel(f intschema.TraceSearchFilter) (*models.Configv1TraceSearchFilter, error) {
	traceFilter, err := traceFilterToModel(f.Trace)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	spanFilters, err := sliceutil.MapErr(f.Span, spanFilterToModel)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &models.Configv1TraceSearchFilter{
		Span:  spanFilters,
		Trace: traceFilter,
	}, nil
}

func spanFilterFromModel(
	s *models.TraceSearchFilterSpanFilter,
) intschema.TraceSpanFilter {
	return intschema.TraceSpanFilter{
		Duration:        durationFilterFromModel(s.Duration),
		Error:           boolFilterFromModel(s.Error),
		MatchType:       string(s.MatchType),
		Operation:       stringFilterFromModel(s.Operation),
		ParentOperation: stringFilterFromModel(s.ParentOperation),
		ParentService:   stringFilterFromModel(s.ParentService),
		Service:         stringFilterFromModel(s.Service),
		SpanCount:       countFilterFromModel(s.SpanCount),
		Tag:             sliceutil.Map(s.Tags, tagFilterFromModel),
	}
}

func spanFilterToModel(s intschema.TraceSpanFilter) (*models.TraceSearchFilterSpanFilter, error) {
	matchType, err := prettyenum.NewSpanFilterMatchType(s.MatchType)
	if err != nil {
		return nil, err
	}
	durFilter, err := durationFilterToModel(s.Duration)
	if err != nil {
		return nil, err
	}
	if len(s.Tags) > 0 {
		return nil, errors.New("replace all usage of \"tags\" with \"tag\"")
	}
	return &models.TraceSearchFilterSpanFilter{
		Duration:        durFilter,
		Error:           boolFilterToModel(s.Error),
		MatchType:       matchType.Model(),
		Operation:       stringFilterToModel(s.Operation),
		ParentOperation: stringFilterToModel(s.ParentOperation),
		ParentService:   stringFilterToModel(s.ParentService),
		Service:         stringFilterToModel(s.Service),
		SpanCount:       countFilterToModel(s.SpanCount),
		Tags:            sliceutil.Map(s.Tag, tagFilterToModel),
	}, nil
}

func traceFilterFromModel(
	t *models.TraceSearchFilterTraceFilter,
) *intschema.TraceFilter {
	if t == nil {
		return nil
	}
	return &intschema.TraceFilter{
		Duration: durationFilterFromModel(t.Duration),
		Error:    boolFilterFromModel(t.Error),
	}
}

func traceFilterToModel(
	t *intschema.TraceFilter,
) (*models.TraceSearchFilterTraceFilter, error) {
	if t == nil {
		return nil, nil
	}
	duration, err := durationFilterToModel(t.Duration)
	if err != nil {
		return nil, err
	}
	return &models.TraceSearchFilterTraceFilter{
		Duration: duration,
		Error:    boolFilterToModel(t.Error),
	}, nil
}
