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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// ConsumptionConfigConverter is a converter for ConsumptionConfig
func ConsumptionConfigFromModel(m *models.Configv1ConsumptionConfig) (*intschema.ConsumptionConfig, error) {
	return ConsumptionConfigConverter{}.fromModel(m)
}

func resourceConsumptionConfig() *schema.Resource {
	r := newGenericResource(
		"consumption_config",
		ConsumptionConfigConverter{},
		generatedConsumptionConfig{},
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
) (*models.Configv1ConsumptionConfig, error) {
	return &models.Configv1ConsumptionConfig{
		Partitions: sliceutil.Map(m.Partition, partitionToModel),
	}, nil
}

func partitionToModel(p intschema.ConsumptionConfigPartition) *models.ConsumptionConfigPartition {
	return &models.ConsumptionConfigPartition{
		Name:       p.Name,
		Slug:       p.Slug,
		Filters:    sliceutil.Map(p.Filter, filterToModel),
		Partitions: sliceutil.Map(p.Partition, partitionToModel),
	}
}

func filterToModel(f intschema.PartitionFilter) *models.ConsumptionConfigPartitionFilter {
	return &models.ConsumptionConfigPartitionFilter{
		Operator:   models.FilterOperator(f.Operator),
		Conditions: sliceutil.Map(f.Condition, conditionToModel),
	}
}

func conditionToModel(c intschema.PartitionFilterCondition) *models.PartitionFilterCondition {
	result := &models.PartitionFilterCondition{
		DatasetSlug: c.DatasetId.Slug(),
	}
	if c.LogFilter != nil {
		result.LogFilter = &models.Configv1LogSearchFilter{
			Query: c.LogFilter.Query,
		}
	}
	return result
}

func (ConsumptionConfigConverter) fromModel(
	m *models.Configv1ConsumptionConfig,
) (*intschema.ConsumptionConfig, error) {
	return &intschema.ConsumptionConfig{
		Partition: sliceutil.Map(m.Partitions, partitionFromModel),
	}, nil
}

func partitionFromModel(p *models.ConsumptionConfigPartition) intschema.ConsumptionConfigPartition {
	return intschema.ConsumptionConfigPartition{
		Name:      p.Name,
		Slug:      p.Slug,
		Filter:    sliceutil.Map(p.Filters, filterFromModel),
		Partition: sliceutil.Map(p.Partitions, partitionFromModel),
	}
}

func filterFromModel(f *models.ConsumptionConfigPartitionFilter) intschema.PartitionFilter {
	return intschema.PartitionFilter{
		Operator:  string(f.Operator),
		Condition: sliceutil.Map(f.Conditions, conditionFromModel),
	}
}

func conditionFromModel(c *models.PartitionFilterCondition) intschema.PartitionFilterCondition {
	result := intschema.PartitionFilterCondition{
		DatasetId: tfid.Slug(c.DatasetSlug),
	}
	if c.LogFilter != nil {
		result.LogFilter = &intschema.PartitionFilterConditionLogFilter{
			Query: c.LogFilter.Query,
		}
	}
	return result
}
