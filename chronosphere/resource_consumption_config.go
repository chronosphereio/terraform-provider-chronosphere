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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// ConsumptionConfigConverter is a converter for ConsumptionConfig
func ConsumptionConfigFromModel(m *models.ConfigunstableConsumptionConfig) (*intschema.ConsumptionConfig, error) {
	return ConsumptionConfigConverter{}.fromModel(m)
}

func resourceConsumptionConfig() *schema.Resource {
	r := newGenericResource(
		"consumption_config",
		ConsumptionConfigConverter{},
		generatedUnstableConsumptionConfig{},
	)
	return &schema.Resource{
		Schema:        tfschema.ConsumptionConfig,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		// TODO(codyg): dry run not implemented until setUnknown supports recursive types.
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

type ConsumptionConfigConverter struct{}

func (ConsumptionConfigConverter) toModel(
	m *intschema.ConsumptionConfig,
) (*models.ConfigunstableConsumptionConfig, error) {
	return &models.ConfigunstableConsumptionConfig{
		Partitions: sliceutil.Map(m.Partition, partitionToModel),
	}, nil
}

func partitionToModel(p intschema.ConsumptionConfigPartition) *models.ConsumptionConfigPartition {
	return &models.ConsumptionConfigPartition{
		Name:           p.Name,
		DatasetFilters: sliceutil.Map(p.DatasetFilter, datasetFilterToModel),
		Partitions:     sliceutil.Map(p.Partition, partitionToModel),
	}
}

func datasetFilterToModel(df intschema.DatasetFilter) *models.ConfigunstableDatasetFilter {
	return &models.ConfigunstableDatasetFilter{
		Operator: models.DatasetFilterOperator(df.Operator),
		Datasets: sliceutil.Map(df.Dataset, func(d intschema.DatasetFilterDataset) *models.DatasetFilterDataset {
			return &models.DatasetFilterDataset{
				DatasetSlug: d.DatasetId.Slug(),
			}
		}),
	}
}

func (ConsumptionConfigConverter) fromModel(
	m *models.ConfigunstableConsumptionConfig,
) (*intschema.ConsumptionConfig, error) {
	return &intschema.ConsumptionConfig{
		Partition: sliceutil.Map(m.Partitions, partitionFromModel),
	}, nil
}

func partitionFromModel(p *models.ConsumptionConfigPartition) intschema.ConsumptionConfigPartition {
	return intschema.ConsumptionConfigPartition{
		Name:          p.Name,
		DatasetFilter: sliceutil.Map(p.DatasetFilters, datasetFilterFromModel),
		Partition:     sliceutil.Map(p.Partitions, partitionFromModel),
	}
}

func datasetFilterFromModel(df *models.ConfigunstableDatasetFilter) intschema.DatasetFilter {
	return intschema.DatasetFilter{
		Operator: string(df.Operator),
		Dataset: sliceutil.Map(df.Datasets, func(d *models.DatasetFilterDataset) intschema.DatasetFilterDataset {
			return intschema.DatasetFilterDataset{
				DatasetId: tfid.Slug(d.DatasetSlug),
			}
		}),
	}
}
