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
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
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
	filter, err := traceSearchFilterToModel(r.TraceFilter)
	if err != nil {
		return nil, fmt.Errorf("couldn't convert trace filter to model: %w", err)
	}
	return &models.Configv1TraceMetricsRule{
		Slug:                 r.Slug,
		Name:                 r.Name,
		MetricName:           r.MetricName,
		TraceFilter:          filter,
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
		TraceFilter:             traceSearchFilterFromModel(r.TraceFilter),
		HistogramBucketsSeconds: r.HistogramBucketsSecs,
		MetricLabels:            r.MetricLabels,
		GroupBy:                 sliceutil.Map(r.GroupBy, traceMetricsRuleGroupByFromModel),
	}, nil
}

func TraceMetricsRuleFromModel(r *models.Configv1TraceMetricsRule) (*intschema.TraceMetricsRule, error) {
	return traceMetricsRuleConverter{}.fromModel(r)
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
