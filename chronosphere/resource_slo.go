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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	v1models "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// SLOFromModel maps an API model to the intschema model.
func SLOFromModel(m *models.ConfigunstableSLO) (*intschema.Slo, error) {
	return sloConverter{}.fromModel(m)
}

func resourceSLO() *schema.Resource {
	r := newGenericResource(
		"slo",
		sloConverter{},
		generatedUnstableSLO{},
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

func (sloConverter) toModel(s *intschema.Slo) (*models.ConfigunstableSLO, error) {
	var (
		customIndicator      *models.SLICustomIndicatorConfig
		endpointAvailability *models.SLIEndpointAvailabilityConfig
		endpointLatency      *models.SLIEndpointLatencyConfig
	)

	if s.Sli.CustomIndicator != nil {
		customIndicator = &models.SLICustomIndicatorConfig{
			GoodQueryTemplate:  s.Sli.CustomIndicator.GoodQueryTemplate,
			BadQueryTemplate:   s.Sli.CustomIndicator.BadQueryTemplate,
			TotalQueryTemplate: s.Sli.CustomIndicator.TotalQueryTemplate,
		}
	}

	if s.Sli.EndpointAvailability != nil {
		endpointAvailability = &models.SLIEndpointAvailabilityConfig{
			ErrorCodes:              s.Sli.EndpointAvailability.ErrorCodes,
			AdditionalPromqlFilters: promFiltersToModel(s.Sli.EndpointAvailability.AdditionalPromqlFilters),
			EndpointsMonitored:      s.Sli.EndpointAvailability.EndpointsMonitored,
		}
	}

	if s.Sli.EndpointLatency != nil {
		endpointLatency = &models.SLIEndpointLatencyConfig{
			LatencyBucket:           s.Sli.EndpointLatency.LatencyBucket,
			AdditionalPromqlFilters: promFiltersToModel(s.Sli.EndpointLatency.AdditionalPromqlFilters),
			EndpointsMonitored:      s.Sli.EndpointLatency.EndpointsMonitored,
		}
	}

	collSlug, v1CollRef := collectionRefFromID(s.CollectionId.Slug())
	if v1CollRef == nil {
		v1CollRef = &v1models.Configv1CollectionReference{
			Type: v1models.Configv1CollectionReferenceTypeSIMPLE,
			Slug: collSlug,
		}
	}

	return &models.ConfigunstableSLO{
		Name:                   s.Name,
		Slug:                   s.Slug,
		Description:            s.Description,
		CollectionRef:          v1CollectionRefToUnstable(v1CollRef),
		NotificationPolicySlug: s.NotificationPolicyId.Slug(),
		Definition: &models.SLODefinition{
			Objective:              s.Definition.Objective,
			ReportingWindows:       reportingWindowsToModel(s.Definition.ReportingWindows),
			BurnRateAlertingConfig: burnRateDefinitionToModel(s.Definition.BurnRateAlertingConfig),
		},
		Sli: &models.ConfigunstableSLI{
			LensTemplateIndicator:   s.Sli.LensTemplateIndicator,
			CustomIndicator:         customIndicator,
			EndpointAvailability:    endpointAvailability,
			EndpointLatency:         endpointLatency,
			CustomDimensionLabels:   s.Sli.CustomDimensionLabels,
			AdditionalPromqlFilters: promFiltersToModel(s.Sli.AdditionalPromqlFilters),
		},
		SignalGrouping: unstableMonitorSignalGroupingToModel(s.SignalGrouping),
		Annotations:    s.Annotations,
		Labels:         s.Labels,
	}, nil
}

func (sloConverter) fromModel(
	s *models.ConfigunstableSLO,
) (*intschema.Slo, error) {
	var (
		endpointAvailability *intschema.SloSliEndpointAvailability
		endpointLatency      *intschema.SloSliEndpointLatency
		customIndicator      *intschema.SloSliCustomIndicator
	)

	if s.Sli.EndpointAvailability != nil {
		endpointAvailability = &intschema.SloSliEndpointAvailability{
			ErrorCodes:              s.Sli.EndpointAvailability.ErrorCodes,
			AdditionalPromqlFilters: promFiltersFromModel(s.Sli.EndpointAvailability.AdditionalPromqlFilters),
			EndpointsMonitored:      s.Sli.EndpointAvailability.EndpointsMonitored,
		}
	}

	if s.Sli.EndpointLatency != nil {
		endpointLatency = &intschema.SloSliEndpointLatency{
			LatencyBucket:           s.Sli.EndpointLatency.LatencyBucket,
			AdditionalPromqlFilters: promFiltersFromModel(s.Sli.EndpointLatency.AdditionalPromqlFilters),
			EndpointsMonitored:      s.Sli.EndpointLatency.EndpointsMonitored,
		}
	}

	if s.Sli.CustomIndicator != nil {
		customIndicator = &intschema.SloSliCustomIndicator{
			GoodQueryTemplate:  s.Sli.CustomIndicator.GoodQueryTemplate,
			BadQueryTemplate:   s.Sli.CustomIndicator.BadQueryTemplate,
			TotalQueryTemplate: s.Sli.CustomIndicator.TotalQueryTemplate,
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
			ReportingWindows:       reportingWindowsFromModel(s.Definition.ReportingWindows),
			BurnRateAlertingConfig: burnRateDefinitionFromModel(s.Definition.BurnRateAlertingConfig),
		},
		Sli: intschema.SloSli{
			LensTemplateIndicator:   s.Sli.LensTemplateIndicator,
			CustomIndicator:         customIndicator,
			EndpointAvailability:    endpointAvailability,
			EndpointLatency:         endpointLatency,
			CustomDimensionLabels:   s.Sli.CustomDimensionLabels,
			AdditionalPromqlFilters: promFiltersFromModel(s.Sli.AdditionalPromqlFilters),
		},
		SignalGrouping: unstableMonitorSignalGroupingFromModel(s.SignalGrouping),
		Annotations:    s.Annotations,
		Labels:         s.Labels,
	}, nil
}

func unstableRefToV1(ref *models.Configv1CollectionReference) *v1models.Configv1CollectionReference {
	return &v1models.Configv1CollectionReference{
		Type: v1models.Configv1CollectionReferenceType(ref.Type),
		Slug: ref.Slug,
	}
}

func v1CollectionRefToUnstable(ref *v1models.Configv1CollectionReference) *models.Configv1CollectionReference {
	return &models.Configv1CollectionReference{
		Type: models.Configv1CollectionReferenceType(ref.Type),
		Slug: ref.Slug,
	}
}

func promFiltersToModel(filters []intschema.SLOAdditionalPromQLFilters) []*models.ConfigunstablePromQLMatcher {
	return sliceutil.Map(filters, func(f intschema.SLOAdditionalPromQLFilters) *models.ConfigunstablePromQLMatcher {
		return &models.ConfigunstablePromQLMatcher{
			Name:  f.Name,
			Type:  models.ConfigunstablePromQLMatcherType(f.Type),
			Value: f.Value,
		}
	})
}

func promFiltersFromModel(filters []*models.ConfigunstablePromQLMatcher) []intschema.SLOAdditionalPromQLFilters {
	return sliceutil.Map(filters, func(f *models.ConfigunstablePromQLMatcher) intschema.SLOAdditionalPromQLFilters {
		return intschema.SLOAdditionalPromQLFilters{
			Name:  f.Name,
			Type:  string(f.Type),
			Value: f.Value,
		}
	})
}

func reportingWindowsToModel(windows []intschema.SloDefinitionReportingWindows) []*models.DefinitionTimeWindow {
	return sliceutil.Map(windows, func(w intschema.SloDefinitionReportingWindows) *models.DefinitionTimeWindow {
		return &models.DefinitionTimeWindow{Duration: w.Duration}
	})
}

func reportingWindowsFromModel(windows []*models.DefinitionTimeWindow) []intschema.SloDefinitionReportingWindows {
	return sliceutil.Map(windows, func(w *models.DefinitionTimeWindow) intschema.SloDefinitionReportingWindows {
		return intschema.SloDefinitionReportingWindows{Duration: w.Duration}
	})
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

// TODO: once SLOs have been promoted to v1 this can be removed and monitorSignalGroupingToModel can be used.
func unstableMonitorSignalGroupingToModel(
	g *intschema.SignalGrouping,
) *models.MonitorSignalGrouping {
	if g == nil {
		return nil
	}
	return &models.MonitorSignalGrouping{
		LabelNames:      g.LabelNames,
		SignalPerSeries: g.SignalPerSeries,
	}
}

// TODO: once SLOs have been promoted to v1 this can be removed and monitorSignalGroupingFromModel can be used.
func unstableMonitorSignalGroupingFromModel(
	g *models.MonitorSignalGrouping,
) *intschema.SignalGrouping {
	if g == nil {
		return nil
	}
	return &intschema.SignalGrouping{
		LabelNames:      g.LabelNames,
		SignalPerSeries: g.SignalPerSeries,
	}
}
