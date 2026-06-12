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
		Description:   "Singleton tree of consumption partitions that classify metric, log, and trace data into hierarchical groups so that consumption_budget resources can enforce quotas against them.",
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
	partitions, err := sliceutil.MapErr(m.Partition, partitionToModel)
	if err != nil {
		return nil, err
	}
	return &models.Configv1ConsumptionConfig{
		Partitions: partitions,
	}, nil
}

func partitionToModel(p intschema.ConsumptionConfigPartition) (*models.ConsumptionConfigPartition, error) {
	filters, err := sliceutil.MapErr(p.Filter, filterToModel)
	if err != nil {
		return nil, err
	}
	partitions, err := sliceutil.MapErr(p.Partition, partitionToModel)
	if err != nil {
		return nil, err
	}
	return &models.ConsumptionConfigPartition{
		Name:       p.Name,
		Slug:       p.Slug,
		Filters:    filters,
		Partitions: partitions,
	}, nil
}

func filterToModel(f intschema.PartitionFilter) (*models.ConsumptionConfigPartitionFilter, error) {
	conditions, err := sliceutil.MapErr(f.Condition, conditionToModel)
	if err != nil {
		return nil, err
	}
	return &models.ConsumptionConfigPartitionFilter{
		Operator:   models.FilterOperator(f.Operator),
		Conditions: conditions,
	}, nil
}

func conditionToModel(c intschema.PartitionFilterCondition) (*models.PartitionFilterCondition, error) {
	result := &models.PartitionFilterCondition{
		DatasetSlug:   c.DatasetId.Slug(),
		MetricFilters: metricFiltersToModel(c.MetricFilter),
	}
	if c.LogFilter != nil {
		result.LogFilter = &models.Configv1LogSearchFilter{
			Query: c.LogFilter.Query,
		}
	}
	if c.TraceFilter != nil {
		result.TraceFilter = &models.Configv1ConsumptionTraceFilter{
			SpanFilters: sliceutil.Map(c.TraceFilter.SpanFilter, consumptionSpanFilterToModel),
		}
	}
	return result, nil
}

func metricFiltersToModel(filters []intschema.PartitionFilterConditionMetricFilter) []*models.Configv1LabelFilter {
	return sliceutil.Map(filters, func(f intschema.PartitionFilterConditionMetricFilter) *models.Configv1LabelFilter {
		return &models.Configv1LabelFilter{
			Name:      f.Name,
			ValueGlob: f.ValueGlob,
		}
	})
}

func consumptionSpanFilterToModel(
	s intschema.PartitionFilterConditionTraceFilterSpanFilter,
) *models.ConsumptionTraceFilterConsumptionSpanFilter {
	return &models.ConsumptionTraceFilterConsumptionSpanFilter{
		Service:         stringFilterToModel(s.Service),
		Operation:       stringFilterToModel(s.Operation),
		ParentService:   stringFilterToModel(s.ParentService),
		ParentOperation: stringFilterToModel(s.ParentOperation),
		Duration:        durationFilterToModel(s.Duration),
		Error:           boolFilterToModel(s.Error),
		Tags:            sliceutil.Map(s.Tag, tagFilterToModel),
		IsRootSpan:      boolFilterToModel(s.IsRootSpan),
	}
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
		DatasetId:    tfid.Slug(c.DatasetSlug),
		MetricFilter: metricFiltersFromModel(c.MetricFilters),
	}
	if c.LogFilter != nil {
		result.LogFilter = &intschema.PartitionFilterConditionLogFilter{
			Query: c.LogFilter.Query,
		}
	}
	if c.TraceFilter != nil {
		result.TraceFilter = &intschema.PartitionFilterConditionTraceFilter{
			SpanFilter: sliceutil.Map(c.TraceFilter.SpanFilters, consumptionSpanFilterFromModel),
		}
	}
	return result
}

func metricFiltersFromModel(filters []*models.Configv1LabelFilter) []intschema.PartitionFilterConditionMetricFilter {
	return sliceutil.Map(filters, func(f *models.Configv1LabelFilter) intschema.PartitionFilterConditionMetricFilter {
		return intschema.PartitionFilterConditionMetricFilter{
			Name:      f.Name,
			ValueGlob: f.ValueGlob,
		}
	})
}

func consumptionSpanFilterFromModel(
	s *models.ConsumptionTraceFilterConsumptionSpanFilter,
) intschema.PartitionFilterConditionTraceFilterSpanFilter {
	return intschema.PartitionFilterConditionTraceFilterSpanFilter{
		Service:         stringFilterFromModel(s.Service),
		Operation:       stringFilterFromModel(s.Operation),
		ParentService:   stringFilterFromModel(s.ParentService),
		ParentOperation: stringFilterFromModel(s.ParentOperation),
		Duration:        durationFilterFromModel(s.Duration),
		Error:           boolFilterFromModel(s.Error),
		Tag:             sliceutil.Map(s.Tags, tagFilterFromModel),
		IsRootSpan:      boolFilterFromModel(s.IsRootSpan),
	}
}
