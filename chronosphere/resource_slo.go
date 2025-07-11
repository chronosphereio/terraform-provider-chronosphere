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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// SLOFromModel maps an API model to the intschema model.
func SLOFromModel(m *models.Configv1SLO) (*intschema.Slo, error) {
	return sloConverter{}.fromModel(m)
}

func resourceSLO() *schema.Resource {
	r := newGenericResource(
		"slo",
		sloConverter{},
		generatedSLO{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.Slo,
		CustomizeDiff: r.ValidateDryRun(&SLODryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// SLODryRunCount tracks how many times dry run is run during validation for testing.
var SLODryRunCount atomic.Int64

type sloConverter struct{}

func (sloConverter) toModel(s *intschema.Slo) (*models.Configv1SLO, error) {
	var customIndicator *models.SLICustomIndicatorConfig
	var customTimesliceIndicator *models.SLICustomTimeSliceIndicatorConfig

	if s.Sli.CustomIndicator != nil {
		customIndicator = &models.SLICustomIndicatorConfig{
			GoodQueryTemplate:  s.Sli.CustomIndicator.GoodQueryTemplate,
			BadQueryTemplate:   s.Sli.CustomIndicator.BadQueryTemplate,
			TotalQueryTemplate: s.Sli.CustomIndicator.TotalQueryTemplate,
		}
	}

	if s.Sli.CustomTimesliceIndicator != nil {
		customTimesliceIndicator = &models.SLICustomTimeSliceIndicatorConfig{
			QueryTemplate: s.Sli.CustomTimesliceIndicator.QueryTemplate,
			TimesliceSize: enum.SLITimeSliceSize.V1(s.Sli.CustomTimesliceIndicator.TimesliceSize),
			Condition: &models.SLITimeSliceCondition{
				Op:    enum.ConditionOp.V1(s.Sli.CustomTimesliceIndicator.Condition.Op),
				Value: s.Sli.CustomTimesliceIndicator.Condition.Value,
			},
		}
	}

	collSlug, v1CollRef := collectionRefFromID(s.CollectionId.Slug())
	if v1CollRef == nil {
		v1CollRef = &models.Configv1CollectionReference{
			Type: models.Configv1CollectionReferenceTypeSIMPLE,
			Slug: collSlug,
		}
	}

	return &models.Configv1SLO{
		Name:                   s.Name,
		Slug:                   s.Slug,
		Description:            s.Description,
		CollectionRef:          v1CollectionRefToUnstable(v1CollRef),
		NotificationPolicySlug: s.NotificationPolicyId.Slug(),
		Definition: &models.SLODefinition{
			Objective:              s.Definition.Objective,
			BurnRateAlertingConfig: burnRateDefinitionToModel(s.Definition.BurnRateAlertingConfig),
			TimeWindow:             timeWindowToModel(s.Definition.TimeWindow),
			EnableBurnRateAlerting: s.Definition.EnableBurnRateAlerting,
		},
		Sli: &models.Configv1SLI{
			CustomIndicator:          customIndicator,
			CustomTimesliceIndicator: customTimesliceIndicator,
			CustomDimensionLabels:    s.Sli.CustomDimensionLabels,
			AdditionalPromqlFilters:  promFiltersToModel(s.Sli.AdditionalPromqlFilters),
		},
		SignalGrouping: monitorSignalGroupingToModel(s.SignalGrouping),
		Annotations:    s.Annotations,
		Labels:         s.Labels,
	}, nil
}

func (sloConverter) fromModel(
	s *models.Configv1SLO,
) (*intschema.Slo, error) {
	var customIndicator *intschema.SloSliCustomIndicator
	var customTimesliceIndicator *intschema.SloSliCustomTimesliceIndicator

	if s.Sli.CustomIndicator != nil {
		customIndicator = &intschema.SloSliCustomIndicator{
			GoodQueryTemplate:  s.Sli.CustomIndicator.GoodQueryTemplate,
			BadQueryTemplate:   s.Sli.CustomIndicator.BadQueryTemplate,
			TotalQueryTemplate: s.Sli.CustomIndicator.TotalQueryTemplate,
		}
	}

	if s.Sli.CustomTimesliceIndicator != nil {
		// Convert the enum values back to their user-facing aliases
		timesliceSize := "ONE_MINUTE" // default
		if s.Sli.CustomTimesliceIndicator.TimesliceSize == models.SLITimeSliceSizeTIMESLICESIZEFIVEMINUTES {
			timesliceSize = "FIVE_MINUTES"
		}

		customTimesliceIndicator = &intschema.SloSliCustomTimesliceIndicator{
			QueryTemplate: s.Sli.CustomTimesliceIndicator.QueryTemplate,
			TimesliceSize: timesliceSize,
			Condition: intschema.SloSliCustomTimesliceIndicatorCondition{
				Op:    string(s.Sli.CustomTimesliceIndicator.Condition.Op),
				Value: s.Sli.CustomTimesliceIndicator.Condition.Value,
			},
		}
	}

	var collectionID tfid.ID
	if s.CollectionRef.Type == models.Configv1CollectionReferenceTypeSIMPLE {
		collectionID = tfid.Slug(s.CollectionRef.Slug)
	} else {
		collectionID = tfid.Slug(collectionIDFromRef("", unstableRefToV1(s.CollectionRef)))
	}

	return &intschema.Slo{
		Name:                 s.Name,
		Slug:                 s.Slug,
		CollectionId:         collectionID,
		NotificationPolicyId: tfid.Slug(s.NotificationPolicySlug),
		Description:          s.Description,
		Definition: intschema.SloDefinition{
			Objective:              s.Definition.Objective,
			TimeWindow:             timeWindowFromModel(s.Definition.TimeWindow),
			BurnRateAlertingConfig: burnRateDefinitionFromModel(s.Definition.BurnRateAlertingConfig),
			EnableBurnRateAlerting: s.Definition.EnableBurnRateAlerting,
		},
		Sli: intschema.SloSli{
			CustomIndicator:          customIndicator,
			CustomTimesliceIndicator: customTimesliceIndicator,
			CustomDimensionLabels:    s.Sli.CustomDimensionLabels,
			AdditionalPromqlFilters:  promFiltersFromModel(s.Sli.AdditionalPromqlFilters),
		},
		SignalGrouping: monitorSignalGroupingFromModel(s.SignalGrouping),
		Annotations:    s.Annotations,
		Labels:         s.Labels,
	}, nil
}

func unstableRefToV1(ref *models.Configv1CollectionReference) *models.Configv1CollectionReference {
	return &models.Configv1CollectionReference{
		Type: ref.Type,
		Slug: ref.Slug,
	}
}

func v1CollectionRefToUnstable(ref *models.Configv1CollectionReference) *models.Configv1CollectionReference {
	return &models.Configv1CollectionReference{
		Type: models.Configv1CollectionReferenceType(ref.Type),
		Slug: ref.Slug,
	}
}

func promFiltersToModel(filters []intschema.SLOAdditionalPromQLFilters) []*models.CommonPromQLMatcher {
	return sliceutil.Map(filters, func(f intschema.SLOAdditionalPromQLFilters) *models.CommonPromQLMatcher {
		return &models.CommonPromQLMatcher{
			Name:  f.Name,
			Type:  models.CommonPromQLMatcherType(f.Type),
			Value: f.Value,
		}
	})
}

func promFiltersFromModel(filters []*models.CommonPromQLMatcher) []intschema.SLOAdditionalPromQLFilters {
	return sliceutil.Map(filters, func(f *models.CommonPromQLMatcher) intschema.SLOAdditionalPromQLFilters {
		return intschema.SLOAdditionalPromQLFilters{
			Name:  f.Name,
			Type:  string(f.Type),
			Value: f.Value,
		}
	})
}

func timeWindowToModel(window *intschema.SloDefinitionTimeWindow) *models.DefinitionTimeWindow {
	if window == nil {
		return nil
	}
	return &models.DefinitionTimeWindow{Duration: window.Duration}
}

func timeWindowFromModel(window *models.DefinitionTimeWindow) *intschema.SloDefinitionTimeWindow {
	if window == nil {
		return nil
	}
	return &intschema.SloDefinitionTimeWindow{Duration: window.Duration}
}

func burnRateDefinitionToModel(defs []intschema.SloDefinitionBurnRateAlertingConfig) []*models.DefinitionBurnRateDefinition {
	return sliceutil.Map(defs, func(w intschema.SloDefinitionBurnRateAlertingConfig) *models.DefinitionBurnRateDefinition {
		return &models.DefinitionBurnRateDefinition{
			Window:   w.Window,
			Budget:   w.Budget,
			Severity: w.Severity,
			Labels:   w.Labels,
		}
	})
}

func burnRateDefinitionFromModel(defs []*models.DefinitionBurnRateDefinition) []intschema.SloDefinitionBurnRateAlertingConfig {
	return sliceutil.Map(defs, func(w *models.DefinitionBurnRateDefinition) intschema.SloDefinitionBurnRateAlertingConfig {
		return intschema.SloDefinitionBurnRateAlertingConfig{
			Window:   w.Window,
			Budget:   w.Budget,
			Severity: w.Severity,
			Labels:   w.Labels,
		}
	})
}
