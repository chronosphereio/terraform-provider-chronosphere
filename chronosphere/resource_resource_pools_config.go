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
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	apimodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// ResourcePoolsConfigFromModel maps an API model into an intschema model.
func ResourcePoolsConfigFromModel(m *apimodels.Configv1ResourcePools) (*intschema.ResourcePoolsConfig, error) {
	return (resourcePoolsConfigConverter{}).fromModel(m)
}

func resourceResourcePoolsConfig() *schema.Resource {
	r := newGenericResource[
		*models.Configv1ResourcePools,
		intschema.ResourcePoolsConfig,
		*intschema.ResourcePoolsConfig,
	](
		"resource_pools_config",
		resourcePoolsConfigConverter{},
		generatedResourcePools{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.ResourcePoolsConfig,
		CustomizeDiff: r.ValidateDryRun(&ResourcePoolsConfigDryRunCount),
		SchemaVersion: 1,
		Description:   "Shared admin config controlling quota usage in Chronosphere",
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// ResourcePoolsConfigDryRunCount tracks how many times dry run is run during validation for testing.
var ResourcePoolsConfigDryRunCount atomic.Int64

type resourcePoolsConfigConverter struct{}

func (resourcePoolsConfigConverter) toModel(
	r *intschema.ResourcePoolsConfig,
) (*models.Configv1ResourcePools, error) {
	if len(r.Pool) > 0 && len(r.Pools) > 0 {
		return nil, fmt.Errorf("cannot set both pool and pools")
	}

	// The resource has equivalent "pool" and deprecated "pools" lists.
	// Only one is set at any given time and the server does not distinguish the different lists.
	resourcePools := r.Pools
	if len(r.Pool) > 0 {
		resourcePools = r.Pool
	}

	pools, err := sliceutil.MapErr(resourcePools, buildPool)
	if err != nil {
		return nil, err
	}
	defaultPool, err := buildDefaultPool(r.DefaultPool)
	if err != nil {
		return nil, err
	}
	return &apimodels.Configv1ResourcePools{
		DefaultPool: defaultPool,
		Pools:       pools,
	}, nil
}

func (resourcePoolsConfigConverter) fromModel(
	m *models.Configv1ResourcePools,
) (*intschema.ResourcePoolsConfig, error) {
	pools, err := expandPools(m.Pools)
	if err != nil {
		return nil, err
	}
	allocation, err := expandAllocation(m.DefaultPool.Allocation)
	if err != nil {
		return nil, err
	}
	return &intschema.ResourcePoolsConfig{
		DefaultPool: &intschema.ResourcePoolsConfigDefaultPool{
			Allocation: allocation,
			Priorities: expandPriorities(m.DefaultPool.Priorities),
		},
		Pool: pools,
	}, nil
}

type resourcePoolsResourceWrapper struct {
	resource genericResource[
		*models.Configv1ResourcePools,
		intschema.ResourcePoolsConfig,
		*intschema.ResourcePoolsConfig,
	]
}

// The Terraform schema has equivalent "pool" and deprecated "pools" fields that are the same on the server.
// Take the server config and read it into whichever field (pool or pools) that the user has decided to set.
// This avoids a meaningless diff being shown during "terraform plan".
func (resourcePoolsConfigConverter) normalize(config, server *intschema.ResourcePoolsConfig) {
	// pick the tf pool or pools to use.
	tfPools := config.Pool
	if len(config.Pools) > 0 {
		tfPools = config.Pools
	}

	// index config pools by name since we may be adding or removing items
	// in the new schema. We cannot assumt len(config.Pools) == len(schemaConfig.Pools)
	tfPoolsByName := map[string]intschema.ResourcePoolsConfigPool{}
	for _, p := range tfPools {
		tfPoolsByName[p.Name] = p
	}

	for i, pool := range server.Pool {
		if tfPool, ok := tfPoolsByName[pool.Name]; ok {
			if tfPool.MatchRule != "" && len(pool.MatchRules) == 1 {
				pool.MatchRule = pool.MatchRules[0]
				pool.MatchRules = nil
				server.Pool[i] = pool
			}
		}
	}

	if len(config.Pools) > 0 {
		server.Pools = server.Pool
		server.Pool = nil
	}
}

func expandAllocation(allocation *apimodels.ResourcePoolsAllocation) (*intschema.ResourcePoolAllocationSchema, error) {
	if allocation == nil {
		return nil, nil
	}

	fv, err := expandAllocationFixedValues(allocation.FixedValues)
	if err != nil {
		return nil, err
	}

	return &intschema.ResourcePoolAllocationSchema{
		PercentOfLicense: allocation.PercentOfLicense,
		FixedValues:      fv,
	}, nil
}

func expandAllocationFixedValues(
	fixedValues []*apimodels.AllocationFixedValue,
) ([]intschema.ResourcePoolAllocationSchemaFixedValues, error) {
	if len(fixedValues) == 0 {
		return nil, nil
	}
	return sliceutil.MapErr(fixedValues, func(f *apimodels.AllocationFixedValue) (intschema.ResourcePoolAllocationSchemaFixedValues, error) {
		var (
			v   int64
			err error
		)
		if f.Value != "" {
			// Value of zero treated as empty, so only parse if not empty.
			v, err = strconv.ParseInt(f.Value, 10, 64)
			if err != nil {
				return intschema.ResourcePoolAllocationSchemaFixedValues{}, err
			}
		}
		return intschema.ResourcePoolAllocationSchemaFixedValues{
			License: string(f.License),
			Value:   v,
		}, nil
	})
}

func expandPriorities(priorities *apimodels.ResourcePoolsPriorities) *intschema.ResourcePoolPrioritiesSchema {
	if priorities == nil {
		return nil
	}

	return &intschema.ResourcePoolPrioritiesSchema{
		HighPriorityMatchRules: aggregationfilter.ListFromModel(priorities.HighPriorityFilters, aggregationfilter.ResourcePoolsDelimiter),
		LowPriorityMatchRules:  aggregationfilter.ListFromModel(priorities.LowPriorityFilters, aggregationfilter.ResourcePoolsDelimiter),
	}
}

func expandPools(pools []*apimodels.ResourcePoolsPool) ([]intschema.ResourcePoolsConfigPool, error) {
	return sliceutil.MapErr(pools, func(pool *apimodels.ResourcePoolsPool) (intschema.ResourcePoolsConfigPool, error) {
		rules := aggregationfilter.ListFromModel(pool.Filters, aggregationfilter.ResourcePoolsDelimiter)
		allocation, err := expandAllocation(pool.Allocation)
		if err != nil {
			return intschema.ResourcePoolsConfigPool{}, err
		}
		return intschema.ResourcePoolsConfigPool{
			Name:       pool.Name,
			MatchRules: rules,
			Allocation: allocation,
			Priorities: expandPriorities(pool.Priorities),
		}, nil
	})
}

func buildDefaultPool(defaultPool *intschema.ResourcePoolsConfigDefaultPool) (*apimodels.ResourcePoolsDefaultPool, error) {
	if defaultPool == nil {
		return nil, nil
	}
	priorities, err := buildPriorities(defaultPool.Priorities)
	if err != nil {
		return nil, err
	}
	return &apimodels.ResourcePoolsDefaultPool{
		Allocation: buildAllocation(defaultPool.Allocation),
		Priorities: priorities,
	}, nil
}

func buildPool(pool intschema.ResourcePoolsConfigPool) (*apimodels.ResourcePoolsPool, error) {
	var (
		filters []*apimodels.Configv1LabelFilter
		err     error
	)
	if pool.MatchRule != "" {
		filters, err = aggregationfilter.ListToModel([]string{pool.MatchRule}, aggregationfilter.ResourcePoolsDelimiter)
		if err != nil {
			return nil, err
		}
	} else {
		filters, err = aggregationfilter.ListToModel(pool.MatchRules, aggregationfilter.ResourcePoolsDelimiter)
		if err != nil {
			return nil, err
		}
	}
	priorities, err := buildPriorities(pool.Priorities)
	if err != nil {
		return nil, err
	}
	return &apimodels.ResourcePoolsPool{
		Allocation: buildAllocation(pool.Allocation),
		Priorities: priorities,
		Filters:    filters,
		Name:       pool.Name,
	}, nil
}

func buildAllocation(allocation *intschema.ResourcePoolAllocationSchema) *apimodels.ResourcePoolsAllocation {
	if allocation == nil {
		return nil
	}

	return &apimodels.ResourcePoolsAllocation{
		PercentOfLicense: allocation.PercentOfLicense,
		FixedValues:      buildFixedValues(allocation.FixedValues),
	}
}

func buildFixedValues(fixedValues []intschema.ResourcePoolAllocationSchemaFixedValues) []*apimodels.AllocationFixedValue {
	if len(fixedValues) == 0 {
		return nil
	}

	return sliceutil.Map(fixedValues, func(f intschema.ResourcePoolAllocationSchemaFixedValues) *apimodels.AllocationFixedValue {
		return &apimodels.AllocationFixedValue{
			License: apimodels.ResourcePoolsLicense(f.License),
			Value:   fmt.Sprint(f.Value),
		}
	})
}

func buildPriorities(priorities *intschema.ResourcePoolPrioritiesSchema) (*apimodels.ResourcePoolsPriorities, error) {
	if priorities == nil {
		return nil, nil
	}

	high, err := aggregationfilter.ListToModel(priorities.HighPriorityMatchRules, aggregationfilter.ResourcePoolsDelimiter)
	if err != nil {
		return nil, err
	}

	low, err := aggregationfilter.ListToModel(priorities.LowPriorityMatchRules, aggregationfilter.ResourcePoolsDelimiter)
	if err != nil {
		return nil, err
	}

	return &apimodels.ResourcePoolsPriorities{
		HighPriorityFilters: high,
		LowPriorityFilters:  low,
	}, nil
}
