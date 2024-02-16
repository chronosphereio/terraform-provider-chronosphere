// Copyright 2023 Chronosphere Inc.
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
	"context"
	"fmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/trace_tail_sampling_rules"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TraceTailSamplingRulesID is the static ID of the global trace tail sampling rules singleton.
const TraceTailSamplingRulesID = "trace_tail_sampling_singleton"

func resourceTraceTailSamplingRules() *schema.Resource {
	return &schema.Resource{
		Schema:        tfschema.TraceTailSamplingRules,
		CreateContext: resourceTraceTailSamplingRulesCreate,
		ReadContext:   resourceTraceTailSamplingRulesRead,
		UpdateContext: resourceTraceTailSamplingRulesUpdate,
		DeleteContext: resourceTraceTailSamplingRulesDelete,
		CustomizeDiff: resourceTraceTailSamplingRulesCustomizeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// -----
// toModel and related helpers.
// -----

func toModel(
	m *intschema.TraceTailSamplingRules,
) (*models.ConfigunstableTraceTailSamplingRules, error) {
	return &models.ConfigunstableTraceTailSamplingRules{
		DefaultSampleRate: defaultSampleRateToModel(m.DefaultSampleRate),
		Rules:             sliceutil.Map(m.Rules, ruleToModel),
	}, nil
}

func defaultSampleRateToModel(
	t *intschema.TraceTailSamplingRulesDefaultSampleRate,
) *models.ConfigunstableDefaultSampleRate {
	if t == nil {
		return nil
	}
	return &models.ConfigunstableDefaultSampleRate{
		Enabled:    t.Enabled,
		SampleRate: t.SampleRate,
	}
}

func ruleToModel(
	r intschema.TraceTailSamplingRulesRules,
) *models.ConfigunstableTraceTailSamplingRule {
	return &models.ConfigunstableTraceTailSamplingRule{
		Name:       r.Name,
		SystemName: r.SystemName,
		Filter:     traceSearchFilterToModel(r.Filter),
		SampleRate: r.SampleRate,
	}
}

func traceSearchFilterToModel(
	f *intschema.TraceTailSamplingRulesRulesFilter,
) *models.Configv1TraceSearchFilter {
	if f == nil {
		return nil
	}
	return &models.Configv1TraceSearchFilter{
		Span:  sliceutil.Map(f.Span, spanFilterToModel),
		Trace: traceFilterToModel(f.Trace),
	}
}

func spanFilterToModel(
	s intschema.TraceTailSamplingRulesRulesFilterSpan,
) *models.TraceSearchFilterSpanFilter {
	return &models.TraceSearchFilterSpanFilter{
		Duration:        durationFilterToModel(s.Duration),
		Error:           boolFilterToModel(s.Error),
		MatchType:       enum.SpanFilterMatchType.Legacy(s.MatchType),
		Operation:       stringFilterToModel(s.Operation),
		ParentOperation: stringFilterToModel(s.ParentOperation),
		ParentService:   stringFilterToModel(s.ParentService),
		Service:         stringFilterToModel(s.Service),
		SpanCount:       countFilterToModel(s.SpanCount),
		Tags:            sliceutil.Map(s.Tags, tagToModel),
	}
}

func durationFilterToModel(
	d *intschema.TraceTailSamplingDurationFilterSchema,
) *models.TraceSearchFilterDurationFilter {
	if d == nil {
		return nil
	}
	return &models.TraceSearchFilterDurationFilter{
		MaxSecs: d.MaxSecs,
		MinSecs: d.MinSecs,
	}
}

func boolFilterToModel(
	b *intschema.TraceTailSamplingBoolFilterSchema,
) *models.TraceSearchFilterBoolFilter {
	if b == nil {
		return nil
	}
	return &models.TraceSearchFilterBoolFilter{
		Value: b.Value,
	}
}

func stringFilterToModel(
	s *intschema.TraceTailSamplingStringFilterSchema,
) *models.TraceSearchFilterStringFilter {
	if s == nil {
		return nil
	}
	return &models.TraceSearchFilterStringFilter{
		Match: enum.StringFilterMatchType.Legacy(s.Match),
		Value: s.Value,
	}
}

func numericFilterToModel(
	s *intschema.TraceTailSamplingNumericFilterSchema,
) *models.TraceSearchFilterNumericFilter {
	if s == nil {
		return nil
	}
	return &models.TraceSearchFilterNumericFilter{
		Comparison: enum.NumericFilterComparisonType.Legacy(s.Comparison),
		Value:      s.Value,
	}
}

func countFilterToModel(
	c *intschema.TraceTailSamplingRulesRulesFilterSpanSpanCount,
) *models.TraceSearchFilterCountFilter {
	if c == nil {
		return nil
	}
	return &models.TraceSearchFilterCountFilter{
		Max: int32(c.Max),
		Min: int32(c.Min),
	}
}

func tagToModel(
	t intschema.TraceTailSamplingRulesRulesFilterSpanTags,
) *models.TraceSearchFilterTagFilter {
	return &models.TraceSearchFilterTagFilter{
		Key:          t.Key,
		Value:        stringFilterToModel(t.Value),
		NumericValue: numericFilterToModel(t.NumericValue),
	}
}

func traceFilterToModel(
	t *intschema.TraceTailSamplingRulesRulesFilterTrace,
) *models.TraceSearchFilterTraceFilter {
	if t == nil {
		return nil
	}
	return &models.TraceSearchFilterTraceFilter{
		Duration: durationFilterToModel(t.Duration),
		Error:    boolFilterToModel(t.Error),
	}
}

// -----
// fromModel and related helpers.
// -----

func fromModel(
	m *models.ConfigunstableTraceTailSamplingRules,
) *intschema.TraceTailSamplingRules {
	return &intschema.TraceTailSamplingRules{
		DefaultSampleRate: defaultSampleRateFromModel(m.DefaultSampleRate),
		Rules:             sliceutil.Map(m.Rules, ruleFromModel),
	}
}

func defaultSampleRateFromModel(
	r *models.ConfigunstableDefaultSampleRate,
) *intschema.TraceTailSamplingRulesDefaultSampleRate {
	if r == nil {
		return nil
	}
	return &intschema.TraceTailSamplingRulesDefaultSampleRate{
		Enabled:    r.Enabled,
		SampleRate: r.SampleRate,
	}
}

func ruleFromModel(
	r *models.ConfigunstableTraceTailSamplingRule,
) intschema.TraceTailSamplingRulesRules {
	return intschema.TraceTailSamplingRulesRules{
		Filter:     traceSearchFilterFromModel(r.Filter),
		SampleRate: r.SampleRate,
		SystemName: r.SystemName,
		Name:       r.Name,
	}
}

func traceSearchFilterFromModel(
	f *models.Configv1TraceSearchFilter,
) *intschema.TraceTailSamplingRulesRulesFilter {
	if f == nil {
		return nil
	}
	return &intschema.TraceTailSamplingRulesRulesFilter{
		Span:  sliceutil.Map(f.Span, spanFilterFromModel),
		Trace: traceFilterFromModel(f.Trace),
	}
}

func spanFilterFromModel(
	s *models.TraceSearchFilterSpanFilter,
) intschema.TraceTailSamplingRulesRulesFilterSpan {
	return intschema.TraceTailSamplingRulesRulesFilterSpan{
		Duration:        durationFilterFromModel(s.Duration),
		Error:           boolFilterFromModel(s.Error),
		MatchType:       string(s.MatchType),
		Operation:       stringFilterFromModel(s.Operation),
		ParentOperation: stringFilterFromModel(s.ParentOperation),
		ParentService:   stringFilterFromModel(s.ParentService),
		Service:         stringFilterFromModel(s.Service),
		SpanCount:       countFilterFromModel(s.SpanCount),
		Tags:            sliceutil.Map(s.Tags, tagFromModel),
	}
}

func durationFilterFromModel(
	d *models.TraceSearchFilterDurationFilter,
) *intschema.TraceTailSamplingDurationFilterSchema {
	if d == nil {
		return nil
	}
	return &intschema.TraceTailSamplingDurationFilterSchema{
		MaxSecs: d.MaxSecs,
		MinSecs: d.MinSecs,
	}
}

func boolFilterFromModel(
	b *models.TraceSearchFilterBoolFilter,
) *intschema.TraceTailSamplingBoolFilterSchema {
	if b == nil {
		return nil
	}
	return &intschema.TraceTailSamplingBoolFilterSchema{
		Value: b.Value,
	}
}

func stringFilterFromModel(
	s *models.TraceSearchFilterStringFilter,
) *intschema.TraceTailSamplingStringFilterSchema {
	if s == nil {
		return nil
	}
	return &intschema.TraceTailSamplingStringFilterSchema{
		Match: string(s.Match),
		Value: s.Value,
	}
}

func numericFilterFromModel(
	s *models.TraceSearchFilterNumericFilter,
) *intschema.TraceTailSamplingNumericFilterSchema {
	if s == nil {
		return nil
	}
	return &intschema.TraceTailSamplingNumericFilterSchema{
		Comparison: string(s.Comparison),
		Value:      s.Value,
	}
}

func countFilterFromModel(
	c *models.TraceSearchFilterCountFilter,
) *intschema.TraceTailSamplingRulesRulesFilterSpanSpanCount {
	if c == nil {
		return nil
	}
	return &intschema.TraceTailSamplingRulesRulesFilterSpanSpanCount{
		Max: int64(c.Max),
		Min: int64(c.Min),
	}
}

func tagFromModel(
	t *models.TraceSearchFilterTagFilter,
) intschema.TraceTailSamplingRulesRulesFilterSpanTags {
	return intschema.TraceTailSamplingRulesRulesFilterSpanTags{
		Key:          t.Key,
		Value:        stringFilterFromModel(t.Value),
		NumericValue: numericFilterFromModel(t.NumericValue),
	}
}

func traceFilterFromModel(
	t *models.TraceSearchFilterTraceFilter,
) *intschema.TraceTailSamplingRulesRulesFilterTrace {
	if t == nil {
		return nil
	}
	return &intschema.TraceTailSamplingRulesRulesFilterTrace{
		Duration: durationFilterFromModel(t.Duration),
		Error:    boolFilterFromModel(t.Error),
	}
}

func resourceTraceTailSamplingRulesCreate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigUnstableClient(meta)

	rules, err := buildTraceTailSamplingRules(d)
	if err != nil {
		return diag.Errorf("could not build trace tail sampling rules: %v", err)
	}
	req := &trace_tail_sampling_rules.CreateTraceTailSamplingRulesParams{
		Body: &models.ConfigunstableCreateTraceTailSamplingRulesRequest{
			TraceTailSamplingRules: rules,
		},
		Context: ctx,
	}

	if _, err := cli.TraceTailSamplingRules.CreateTraceTailSamplingRules(req); err != nil {
		return diag.Errorf("could not create trace tail sampling rules: %v", err)
	}

	d.SetId(TraceTailSamplingRulesID)

	return nil
}

func resourceTraceTailSamplingRulesRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigUnstableClient(meta)

	resp, err := cli.TraceTailSamplingRules.ReadTraceTailSamplingRules(&trace_tail_sampling_rules.ReadTraceTailSamplingRulesParams{Context: ctx})
	if clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to read trace tail sampling rules: %v", clienterror.Wrap(err))
	}

	rules := resp.Payload.TraceTailSamplingRules
	schemaRules := fromModel(rules)

	if err := schemaRules.ToResourceData(d); err != nil {
		return err
	}
	d.SetId(TraceTailSamplingRulesID)
	return nil
}

func resourceTraceTailSamplingRulesUpdate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigUnstableClient(meta)

	rules, err := buildTraceTailSamplingRules(d)
	if err != nil {
		return diag.Errorf("could not build trace tail sampling rules: %v", err)
	}
	req := &trace_tail_sampling_rules.UpdateTraceTailSamplingRulesParams{
		Context: ctx,
		Body: &models.ConfigunstableUpdateTraceTailSamplingRulesRequest{
			TraceTailSamplingRules: rules,
		},
	}
	if _, err := cli.TraceTailSamplingRules.UpdateTraceTailSamplingRules(req); err != nil {
		return diag.Errorf("unable to update trace tail sampling rules: %v", err)
	}
	return nil
}

func resourceTraceTailSamplingRulesDelete(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigUnstableClient(meta)

	req := &trace_tail_sampling_rules.DeleteTraceTailSamplingRulesParams{Context: ctx}
	if _, err := cli.TraceTailSamplingRules.DeleteTraceTailSamplingRules(req); clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to delete trace tail sampling rules: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTraceTailSamplingRulesCustomizeDiff(
	_ context.Context, d *schema.ResourceDiff, meta any,
) error {
	rules, err := buildTraceTailSamplingRules(d)
	if err != nil {
		return fmt.Errorf("unable to build trace tail sampling rules: %w", err)
	}
	return validateTraceTailSamplingRules(rules)
}

func validateTraceTailSamplingRules(rules *models.ConfigunstableTraceTailSamplingRules) error {
	for _, r := range rules.Rules {
		if r.SampleRate < 0 || r.SampleRate > 1.0 {
			return fmt.Errorf("expected sample rate to be a float from 0 to 1.0 inclusive, got %f", r.SampleRate)
		}
	}

	if rules.DefaultSampleRate.SampleRate < 0 || rules.DefaultSampleRate.SampleRate > 1.0 {
		return fmt.Errorf("expected sample rate to be a float from 0 to 1.0 inclusive, got %f", rules.DefaultSampleRate.SampleRate)
	}

	return nil
}

func buildTraceTailSamplingRules(d ResourceGetter) (*models.ConfigunstableTraceTailSamplingRules, error) {
	rules := &intschema.TraceTailSamplingRules{}
	if err := rules.FromResourceData(d); err != nil {
		return nil, err
	}

	return toModel(rules)
}
