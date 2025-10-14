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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// ConsumptionBudgetFromModel maps an API model to an intschema model.
func ConsumptionBudgetFromModel(m *models.Configv1ConsumptionBudget) (*intschema.ConsumptionBudget, error) {
	return consumptionBudgetConverter{}.fromModel(m)
}

func resourceConsumptionBudget() *schema.Resource {
	r := newGenericResource(
		"consumption_budget",
		consumptionBudgetConverter{},
		generatedConsumptionBudget{})

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.ConsumptionBudget,
		CustomizeDiff: r.ValidateDryRunOptions(&ConsumptionBudgetDryRunCount, ValidateDryRunOpts[*models.Configv1ConsumptionBudget]{
			SetUnknownReferencesSkip: []string{
				// Because this isn't a real API reference.
				"consumption_config_id",
				// Because dry run doesn't support tfids in lists (an artificial constraint).
				"priority.[].filter.[].dataset_id",
			},
			ModifyAPIModel: func(cfg *models.Configv1ConsumptionBudget) {
				for _, p := range cfg.Priorities {
					for _, f := range p.Filters {
						// NOTE(codyg): There's no way to tell if the user
						// actually set a dataset_id or not, and since it has
						// one-of relationship with log_filter, we can't blindly
						// set it, but we also can't always leave it empty,
						// because at least one field must be set. So we do the
						// slightly wrong thing of assuming if the user didn't
						// set log_filter, then the probably set dataset_id.
						// This will cause invalid configs to erroneously pass
						// dry run validation if the user actually provided
						// an empty filter object.
						if f.LogFilter == nil {
							f.DatasetSlug = dryRunUnknownRef.Slug()
						}
					}
				}
			},
		}),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// ConsumptionBudgetDryRunCount tracks how many times dry run is run during
// validation for testing.
var ConsumptionBudgetDryRunCount atomic.Int64

type consumptionBudgetConverter struct{}

func (c consumptionBudgetConverter) toModel(
	s *intschema.ConsumptionBudget,
) (*models.Configv1ConsumptionBudget, error) {
	switch s.ConsumptionConfigId.Slug() {
	case "", ConsumptionConfigID:
		// Valid. Can be empty if the ConsumptionConfig has been created yet.
	default:
		return nil, fmt.Errorf(
			"invalid consumption_config_id %q: must point at chronosphere_consumption_config resource",
			s.ConsumptionConfigId.Slug())
	}
	m := &models.Configv1ConsumptionBudget{
		Name:              s.Name,
		Slug:              s.Slug,
		Resource:          models.Configv1ConsumptionBudgetResource(s.Resource),
		PartitionSlugPath: s.PartitionSlugPath,
		Priorities: sliceutil.Map(s.Priority, func(p intschema.ConsumptionBudgetPriority) *models.ConsumptionBudgetPriority {
			return &models.ConsumptionBudgetPriority{
				Filters:  sliceutil.Map(p.Filter, consumptionBudgetPriorityFilterToModel),
				Priority: int32(p.Priority),
			}
		}),
		Thresholds:             sliceutil.Map(s.Threshold, consumptionBudgetThresholdToModel),
		DefaultPriority:        int32(s.DefaultPriority),
		NotificationPolicySlug: s.NotificationPolicyId.Slug(),
		AlertActionConfig:      consumptionBudgetAlertActionConfigToModel(s.AlertActionConfig),
	}
	return m, nil
}

func (c consumptionBudgetConverter) fromModel(
	m *models.Configv1ConsumptionBudget,
) (*intschema.ConsumptionBudget, error) {
	thresholds, err := sliceutil.MapErr(m.Thresholds, consumptionBudgetThresholdFromModel)
	if err != nil {
		return nil, err
	}
	return &intschema.ConsumptionBudget{
		ConsumptionConfigId: tfid.Slug(ConsumptionConfigID),
		Name:                m.Name,
		Slug:                m.Slug,
		Resource:            string(m.Resource),
		PartitionSlugPath:   m.PartitionSlugPath,
		Priority: sliceutil.Map(m.Priorities, func(p *models.ConsumptionBudgetPriority) intschema.ConsumptionBudgetPriority {
			return intschema.ConsumptionBudgetPriority{
				Filter:   sliceutil.Map(p.Filters, consumptionBudgetPriorityFilterFromModel),
				Priority: int64(p.Priority),
			}
		}),
		Threshold:            thresholds,
		DefaultPriority:      int64(m.DefaultPriority),
		NotificationPolicyId: tfid.Slug(m.NotificationPolicySlug),
		AlertActionConfig:    consumptionBudgetAlertActionConfigFromModel(m.AlertActionConfig),
	}, nil
}

func consumptionBudgetAlertActionConfigToModel(a *intschema.ConsumptionBudgetAlertActionConfig) *models.ConsumptionBudgetAlertActionConfig {
	if a == nil {
		return nil
	}
	return &models.ConsumptionBudgetAlertActionConfig{
		Annotations:            a.Annotations,
		Labels:                 a.Labels,
		InstantRateSustainSecs: int32(a.InstantRateSustainSecs),
	}
}

func consumptionBudgetAlertActionConfigFromModel(a *models.ConsumptionBudgetAlertActionConfig) *intschema.ConsumptionBudgetAlertActionConfig {
	if a == nil {
		return nil
	}
	return &intschema.ConsumptionBudgetAlertActionConfig{
		Annotations:            a.Annotations,
		Labels:                 a.Labels,
		InstantRateSustainSecs: int64(a.InstantRateSustainSecs),
	}
}

func consumptionBudgetPriorityFilterToModel(f intschema.ConsumptionBudgetPriorityFilter) *models.ConsumptionBudgetPriorityFilter {
	result := &models.ConsumptionBudgetPriorityFilter{
		DatasetSlug: f.DatasetId.Slug(),
	}
	if f.LogFilter != nil {
		result.LogFilter = &models.Configv1LogSearchFilter{
			Query: f.LogFilter.Query,
		}
	}
	return result
}

func consumptionBudgetPriorityFilterFromModel(f *models.ConsumptionBudgetPriorityFilter) intschema.ConsumptionBudgetPriorityFilter {
	result := intschema.ConsumptionBudgetPriorityFilter{
		DatasetId: tfid.Slug(f.DatasetSlug),
	}
	if f.LogFilter != nil {
		result.LogFilter = &intschema.ConsumptionBudgetPriorityFilterLogFilter{
			Query: f.LogFilter.Query,
		}
	}
	return result
}

func consumptionBudgetThresholdToModel(b intschema.ConsumptionBudgetThreshold) *models.Configv1ConsumptionBudgetThreshold {
	var instantRate *models.ThresholdInstantRate
	if b.InstantRate != nil {
		instantRate = &models.ThresholdInstantRate{
			FixedValuePerSec: fmt.Sprint(b.InstantRate.FixedValuePerSec),
		}
	}

	var volume *models.ThresholdVolume
	if b.Volume != nil {
		volume = &models.ThresholdVolume{
			FixedValue: fmt.Sprint(b.Volume.FixedValue),
		}
	}

	return &models.Configv1ConsumptionBudgetThreshold{
		Action:      models.ConsumptionBudgetThresholdAction(b.Action),
		Type:        models.ConsumptionBudgetThresholdType(b.Type),
		InstantRate: instantRate,
		Volume:      volume,
	}
}

func consumptionBudgetThresholdFromModel(b *models.Configv1ConsumptionBudgetThreshold) (intschema.ConsumptionBudgetThreshold, error) {
	behavior := intschema.ConsumptionBudgetThreshold{
		Action: string(b.Action),
		Type:   string(b.Type),
	}

	if b.InstantRate != nil {
		fixedValuePerSec, err := parseStringToInt64(
			b.InstantRate.FixedValuePerSec, "instant_rate.fixed_value_per_sec")
		if err != nil {
			return intschema.ConsumptionBudgetThreshold{}, err
		}
		behavior.InstantRate = &intschema.ConsumptionBudgetThresholdInstantRate{
			FixedValuePerSec: fixedValuePerSec,
		}
	}

	if b.Volume != nil {
		fixedValue, err := parseStringToInt64(
			b.Volume.FixedValue, "volume.fixed_value")
		if err != nil {
			return intschema.ConsumptionBudgetThreshold{}, err
		}
		behavior.Volume = &intschema.ConsumptionBudgetThresholdVolume{
			FixedValue: fixedValue,
		}
	}

	return behavior, nil
}
