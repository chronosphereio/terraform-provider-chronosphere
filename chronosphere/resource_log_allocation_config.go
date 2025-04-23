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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func LogAllocationConfigFromModel(m *models.Configv1LogAllocationConfig) (*intschema.LogAllocationConfig, error) {
	return LogAllocationConfigConverter{}.fromModel(m)
}

func resourceLogAllocationConfig() *schema.Resource {
	r := newGenericResource[
		*models.Configv1LogAllocationConfig,
		intschema.LogAllocationConfig,
		*intschema.LogAllocationConfig,
	](
		"log_allocation_config",
		LogAllocationConfigConverter{},
		generatedLogAllocationConfig{},
	)

	return &schema.Resource{
		Schema:        tfschema.LogAllocationConfig,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRunOptions(
			&LogAllocationConfigDryRunCount,
			ValidateDryRunOpts[*models.Configv1LogAllocationConfig]{
				// Note: We skip dataset_id as it's within a list which is not supported
				// by setUnknownReferences.
				// We instead handle them explicitly in ModifyAPIModel.
				SetUnknownReferencesSkip: []string{"dataset_allocation.[].dataset_id"},
				ModifyAPIModel: func(cfg *models.Configv1LogAllocationConfig) {
					for _, alloc := range cfg.DatasetAllocations {
						if alloc.DatasetSlug == "" {
							alloc.DatasetSlug = dryRunUnknownRef.Slug()
						}
					}
				},
			}),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// LogAllocationConfigDryRunCount tracks how many times dry run is run during validation for testing.
var LogAllocationConfigDryRunCount atomic.Int64

type LogAllocationConfigConverter struct{}

func (LogAllocationConfigConverter) toModel(
	m *intschema.LogAllocationConfig,
) (*models.Configv1LogAllocationConfig, error) {
	return &models.Configv1LogAllocationConfig{
		DefaultDataset: &models.LogAllocationConfigDefaultDataset{
			Allocation: allocationToModel(m.DefaultDataset.Allocation),
			Priorities: prioritiesToModel(m.DefaultDataset.Priorities),
		},
		DatasetAllocations: sliceutil.Map(m.DatasetAllocation, datasetAllocationToModel),
	}, nil
}

func datasetAllocationToModel(
	datasetAllocation intschema.LogAllocationConfigDatasetAllocation,
) *models.LogAllocationConfigDatasetAllocation {
	return &models.LogAllocationConfigDatasetAllocation{
		Allocation:  allocationToModel(datasetAllocation.Allocation),
		DatasetSlug: datasetAllocation.DatasetId.Slug(),
		Priorities:  prioritiesToModel(datasetAllocation.Priorities),
	}
}

func allocationToModel(a *intschema.LogAllocationConfigSchema) *models.Configv1LogAllocationConfigAllocation {
	return &models.Configv1LogAllocationConfigAllocation{
		PercentOfLicense: a.PercentOfLicense,
	}
}

func prioritiesToModel(a *intschema.LogPrioritiesSchema) *models.LogAllocationConfigHighLowPriorities {
	if a == nil {
		return nil
	}
	return &models.LogAllocationConfigHighLowPriorities{
		HighPriorityFilters: sliceutil.Map(a.HighPriorityFilter, searchFilterToModel),
		LowPriorityFilters:  sliceutil.Map(a.LowPriorityFilter, searchFilterToModel),
	}
}

func searchFilterToModel(p intschema.LogSearchFilterSchema,
) *models.Configv1LogSearchFilter {
	return &models.Configv1LogSearchFilter{Query: p.Query}
}

func (LogAllocationConfigConverter) fromModel(
	m *models.Configv1LogAllocationConfig,
) (*intschema.LogAllocationConfig, error) {
	return &intschema.LogAllocationConfig{
		DefaultDataset: &intschema.LogAllocationConfigDefaultDataset{
			Allocation: allocationFromModel(m.DefaultDataset.Allocation),
			Priorities: prioritiesFromModel(m.DefaultDataset.Priorities),
		},
		DatasetAllocation: sliceutil.Map(m.DatasetAllocations, datasetAllocationFromModel),
	}, nil
}

func datasetAllocationFromModel(
	datasetAllocation *models.LogAllocationConfigDatasetAllocation,
) intschema.LogAllocationConfigDatasetAllocation {
	return intschema.LogAllocationConfigDatasetAllocation{
		Allocation: allocationFromModel(datasetAllocation.Allocation),
		DatasetId:  tfid.Slug(datasetAllocation.DatasetSlug),
		Priorities: prioritiesFromModel(datasetAllocation.Priorities),
	}
}

func allocationFromModel(a *models.Configv1LogAllocationConfigAllocation) *intschema.LogAllocationConfigSchema {
	return &intschema.LogAllocationConfigSchema{
		PercentOfLicense: a.PercentOfLicense,
	}
}

func prioritiesFromModel(a *models.LogAllocationConfigHighLowPriorities) *intschema.LogPrioritiesSchema {
	if a == nil {
		return nil
	}
	return &intschema.LogPrioritiesSchema{
		HighPriorityFilter: sliceutil.Map(a.HighPriorityFilters, searchFilterFromModel),
		LowPriorityFilter:  sliceutil.Map(a.LowPriorityFilters, searchFilterFromModel),
	}
}

func searchFilterFromModel(p *models.Configv1LogSearchFilter,
) intschema.LogSearchFilterSchema {
	return intschema.LogSearchFilterSchema{Query: p.Query}
}
