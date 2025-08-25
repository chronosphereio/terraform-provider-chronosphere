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
	"strings"

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
		PartitionSlugPath: strings.Join(s.PartitionSlugPath, "/"),
		Priorities: sliceutil.Map(s.Priority, func(p intschema.ConsumptionBudgetPriority) *models.ConsumptionBudgetPriority {
			return &models.ConsumptionBudgetPriority{
				Filters:  sliceutil.Map(p.Filter, consumptionBudgetPriorityFilterToModel),
				Priority: int32(p.Priority),
			}
		}),
		Behaviors:              sliceutil.Map(s.Behavior, consumptionBudgetBehaviorToModel),
		DefaultPriority:        int32(s.DefaultPriority),
		NotificationPolicySlug: s.NotificationPolicyId.Slug(),
	}
	return m, nil
}

func (c consumptionBudgetConverter) fromModel(
	m *models.Configv1ConsumptionBudget,
) (*intschema.ConsumptionBudget, error) {
	behaviors, err := sliceutil.MapErr(m.Behaviors, consumptionBudgetBehaviorFromModel)
	if err != nil {
		return nil, err
	}
	return &intschema.ConsumptionBudget{
		ConsumptionConfigId: tfid.Slug(ConsumptionConfigID),
		Name:                m.Name,
		Slug:                m.Slug,
		Resource:            string(m.Resource),
		PartitionSlugPath:   strings.Split(m.PartitionSlugPath, "/"),
		Priority: sliceutil.Map(m.Priorities, func(p *models.ConsumptionBudgetPriority) intschema.ConsumptionBudgetPriority {
			return intschema.ConsumptionBudgetPriority{
				Filter:   sliceutil.Map(p.Filters, consumptionBudgetPriorityFilterFromModel),
				Priority: int64(p.Priority),
			}
		}),
		Behavior:             behaviors,
		DefaultPriority:      int64(m.DefaultPriority),
		NotificationPolicyId: tfid.Slug(m.NotificationPolicySlug),
	}, nil
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

func consumptionBudgetBehaviorToModel(b intschema.ConsumptionBudgetBehavior) *models.ConsumptionBudgetBehavior {
	var instantRateThreshold *models.BehaviorInstantRateThreshold
	if b.InstantRateThreshold != nil {
		instantRateThreshold = &models.BehaviorInstantRateThreshold{
			FixedValuePerSec: fmt.Sprint(b.InstantRateThreshold.FixedValuePerSec),
		}
	}

	var volumeThreshold *models.BehaviorVolumeThreshold
	if b.VolumeThreshold != nil {
		volumeThreshold = &models.BehaviorVolumeThreshold{
			FixedValue: fmt.Sprint(b.VolumeThreshold.FixedValue),
		}
	}

	return &models.ConsumptionBudgetBehavior{
		Action:               models.ConsumptionBudgetBehaviorAction(b.Action),
		ThresholdType:        models.BehaviorThresholdType(b.ThresholdType),
		InstantRateThreshold: instantRateThreshold,
		VolumeThreshold:      volumeThreshold,
	}
}

func consumptionBudgetBehaviorFromModel(b *models.ConsumptionBudgetBehavior) (intschema.ConsumptionBudgetBehavior, error) {
	behavior := intschema.ConsumptionBudgetBehavior{
		Action:        string(b.Action),
		ThresholdType: string(b.ThresholdType),
	}

	if b.InstantRateThreshold != nil {
		fixedValuePerSec, err := parseStringToInt64(
			b.InstantRateThreshold.FixedValuePerSec, "instant_rate_threshold.fixed_value_per_sec")
		if err != nil {
			return intschema.ConsumptionBudgetBehavior{}, err
		}
		behavior.InstantRateThreshold = &intschema.ConsumptionBudgetBehaviorInstantRateThreshold{
			FixedValuePerSec: fixedValuePerSec,
		}
	}

	if b.VolumeThreshold != nil {
		fixedValue, err := parseStringToInt64(
			b.VolumeThreshold.FixedValue, "volume_threshold.fixed_value")
		if err != nil {
			return intschema.ConsumptionBudgetBehavior{}, err
		}
		behavior.VolumeThreshold = &intschema.ConsumptionBudgetBehaviorVolumeThreshold{
			FixedValue: fixedValue,
		}
	}

	return behavior, nil
}
