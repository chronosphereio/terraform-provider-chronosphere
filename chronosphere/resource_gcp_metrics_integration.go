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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// GcpMetricsIntegrationFromModel maps an API model to the intschema model.
func GcpMetricsIntegrationFromModel(m *models.Configv1GcpMetricsIntegration) (*intschema.GcpMetricsIntegration, error) {
	return (gcpMetricsIntegrationConverter{}).fromModel(m)
}

func resourceGcpMetricsIntegration() *schema.Resource {
	r := newGenericResource(
		"gcp_metrics_integration",
		gcpMetricsIntegrationConverter{},
		generatedGcpMetricsIntegration{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.GcpMetricsIntegration,
		CustomizeDiff: r.ValidateDryRun(&GcpMetricsIntegrationDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// GcpMetricsIntegrationDryRunCount tracks how many times dry run is run during validation for testing.
var GcpMetricsIntegrationDryRunCount atomic.Int64

type gcpMetricsIntegrationConverter struct{}

func (gcpMetricsIntegrationConverter) toModel(
	g *intschema.GcpMetricsIntegration,
) (*models.Configv1GcpMetricsIntegration, error) {
	return &models.Configv1GcpMetricsIntegration{
		Name:         g.Name,
		Slug:         g.Slug,
		MetricGroups: resourceToModelMetricGroups(g.MetricGroups),
		ServiceAccount: &models.Configv1GcpMetricsIntegrationServiceAccount{
			ClientEmail: g.ServiceAccount.ClientEmail,
		},
	}, nil
}

func (gcpMetricsIntegrationConverter) fromModel(
	m *models.Configv1GcpMetricsIntegration,
) (*intschema.GcpMetricsIntegration, error) {
	return &intschema.GcpMetricsIntegration{
		Name:         m.Name,
		Slug:         m.Slug,
		MetricGroups: resourceFromModelMetricGroups(m.MetricGroups),
		ServiceAccount: &intschema.GcpMetricsIntegrationServiceAccount{
			ClientEmail: m.ServiceAccount.ClientEmail,
		},
	}, nil
}

func resourceToModelMetricGroups(
	mg []intschema.GcpMetricsIntegrationMetricGroups,
) []*models.GcpMetricsIntegrationMetricGroup {
	metricGroups := make([]*models.GcpMetricsIntegrationMetricGroup, len(mg))
	for i, g := range mg {
		metricGroups[i] = &models.GcpMetricsIntegrationMetricGroup{
			Prefixes:    g.Prefixes,
			ProjectID:   g.ProjectId,
			Filters:     resourceToModelFilters(g.Filters),
			RollupRules: resourceToModelRollupRules(g.RollupRules),
		}
	}
	return metricGroups
}

func resourceFromModelMetricGroups(
	mg []*models.GcpMetricsIntegrationMetricGroup,
) []intschema.GcpMetricsIntegrationMetricGroups {
	metricGroups := make([]intschema.GcpMetricsIntegrationMetricGroups, len(mg))
	for i, g := range mg {
		metricGroups[i] = intschema.GcpMetricsIntegrationMetricGroups{
			Prefixes:    g.Prefixes,
			ProjectId:   g.ProjectID,
			Filters:     resourceFromModelFilters(g.Filters),
			RollupRules: resourceFromModelRollupRules(g.RollupRules),
		}
	}
	return metricGroups
}

func resourceToModelFilters(
	filters []intschema.GcpMetricsIntegrationMetricGroupsFilters,
) []*models.Configv1GcpMetricsIntegrationFilter {
	result := make([]*models.Configv1GcpMetricsIntegrationFilter, len(filters))
	for i, f := range filters {
		result[i] = &models.Configv1GcpMetricsIntegrationFilter{
			Name:      f.Name,
			ValueGlob: f.ValueGlob,
			Context:   models.GcpMetricsIntegrationLabelContext(f.Context),
		}
	}
	return result
}

func resourceFromModelFilters(
	filters []*models.Configv1GcpMetricsIntegrationFilter,
) []intschema.GcpMetricsIntegrationMetricGroupsFilters {
	result := make([]intschema.GcpMetricsIntegrationMetricGroupsFilters, len(filters))
	for i, f := range filters {
		result[i] = intschema.GcpMetricsIntegrationMetricGroupsFilters{
			Name:      f.Name,
			ValueGlob: f.ValueGlob,
			Context:   string(f.Context),
		}
	}
	return result
}

func resourceToModelRollupRules(
	rules []intschema.GcpMetricsIntegrationMetricGroupsRollupRules,
) []*models.Configv1GcpMetricsIntegrationRollupRule {
	result := make([]*models.Configv1GcpMetricsIntegrationRollupRule, len(rules))
	for i, r := range rules {
		result[i] = &models.Configv1GcpMetricsIntegrationRollupRule{
			MetricName:  r.MetricName,
			Aggregation: models.Configv1GcpMetricsIntegrationAggregation(r.Aggregation),
			LabelPolicy: resourceToModelLabelPolicy(r.LabelPolicy),
		}
	}
	return result
}

func resourceFromModelRollupRules(
	rules []*models.Configv1GcpMetricsIntegrationRollupRule,
) []intschema.GcpMetricsIntegrationMetricGroupsRollupRules {
	result := make([]intschema.GcpMetricsIntegrationMetricGroupsRollupRules, len(rules))
	for i, r := range rules {
		result[i] = intschema.GcpMetricsIntegrationMetricGroupsRollupRules{
			MetricName:  r.MetricName,
			Aggregation: string(r.Aggregation),
			LabelPolicy: resourceFromModelLabelPolicy(r.LabelPolicy),
		}
	}
	return result
}

func resourceToModelLabelPolicy(
	policy *intschema.GcpMetricsIntegrationMetricGroupsRollupRulesLabelPolicy,
) *models.GcpMetricsIntegrationRollupRuleLabelPolicy {
	if policy == nil {
		return nil
	}
	keep := make([]*models.GcpMetricsIntegrationRollupRuleLabel, len(policy.Keep))
	for i, k := range policy.Keep {
		keep[i] = &models.GcpMetricsIntegrationRollupRuleLabel{
			Name:    k.Name,
			Context: models.GcpMetricsIntegrationLabelContext(k.Context),
		}
	}
	return &models.GcpMetricsIntegrationRollupRuleLabelPolicy{
		Keep: keep,
	}
}

func resourceFromModelLabelPolicy(
	policy *models.GcpMetricsIntegrationRollupRuleLabelPolicy,
) *intschema.GcpMetricsIntegrationMetricGroupsRollupRulesLabelPolicy {
	if policy == nil {
		return nil
	}
	keep := make([]intschema.GcpMetricsIntegrationMetricGroupsRollupRulesLabelPolicyKeep, len(policy.Keep))
	for i, k := range policy.Keep {
		keep[i] = intschema.GcpMetricsIntegrationMetricGroupsRollupRulesLabelPolicyKeep{
			Name:    k.Name,
			Context: string(k.Context),
		}
	}
	return &intschema.GcpMetricsIntegrationMetricGroupsRollupRulesLabelPolicy{
		Keep: keep,
	}
}
