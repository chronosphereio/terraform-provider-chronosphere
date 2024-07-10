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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceLogAllocationConfig() *schema.Resource {
	r := newGenericResource[
		*models.ConfigunstableLogAllocationConfig,
		intschema.LogAllocationConfig,
		*intschema.LogAllocationConfig,
	](
		"log_allocation_config",
		LogAllocationConfigConverter{},
		generatedUnstableLogAllocationConfig{},
	)

	return &schema.Resource{
		Schema:        tfschema.LogAllocationConfig,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&LogAllocationConfigDryRunCount),
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
) (*models.ConfigunstableLogAllocationConfig, error) {
	return &models.ConfigunstableLogAllocationConfig{
		DefaultDataset: &models.LogAllocationConfigDefaultDataset{
			Allocation: allocationToModel(m.DefaultDataset.Allocation),
		},
		DatasetAllocations: sliceutil.Map(m.DatasetAllocation, datasetAllocationToModel),
	}, nil
}

func datasetAllocationToModel(
	datasetAllocation intschema.LogAllocationConfigDatasetAllocation,
) *models.LogAllocationConfigDatasetAllocation {
	return &models.LogAllocationConfigDatasetAllocation{
		Allocation:  allocationToModel(datasetAllocation.Allocation),
		DatasetSlug: datasetAllocation.DatasetSlug,
	}
}

func allocationToModel(a *intschema.LogAllocationConfigSchema) *models.LogAllocationConfigAllocation {
	return &models.LogAllocationConfigAllocation{
		PercentOfLicense: a.PercentOfLicense,
	}
}

func (LogAllocationConfigConverter) fromModel(
	m *models.ConfigunstableLogAllocationConfig,
) (*intschema.LogAllocationConfig, error) {
	return &intschema.LogAllocationConfig{
		DefaultDataset: &intschema.LogAllocationConfigDefaultDataset{
			Allocation: allocationFromModel(m.DefaultDataset.Allocation),
		},
		DatasetAllocation: sliceutil.Map(m.DatasetAllocations, datasetAllocationFromModel),
	}, nil
}

func datasetAllocationFromModel(
	datasetAllocation *models.LogAllocationConfigDatasetAllocation,
) intschema.LogAllocationConfigDatasetAllocation {
	return intschema.LogAllocationConfigDatasetAllocation{
		Allocation:  allocationFromModel(datasetAllocation.Allocation),
		DatasetSlug: datasetAllocation.DatasetSlug,
	}
}

func allocationFromModel(a *models.LogAllocationConfigAllocation) *intschema.LogAllocationConfigSchema {
	return &intschema.LogAllocationConfigSchema{
		PercentOfLicense: a.PercentOfLicense,
	}
}
