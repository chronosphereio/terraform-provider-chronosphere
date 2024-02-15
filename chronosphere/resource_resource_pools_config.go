// Copyright 2023 Chronosphere Inc.
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
	"context"
	stderrors "errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/resource_pools"
	apimodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// ResourcePoolsConfigID is the static ID of the global resource pools singleton.
const ResourcePoolsConfigID = "resource_pool_singleton"

// ResourcePoolsConfigFromModel maps an API model into an intschema model.
func ResourcePoolsConfigFromModel(m *apimodels.Configv1ResourcePools) *intschema.ResourcePoolsConfig {
	return &intschema.ResourcePoolsConfig{
		DefaultPool: &intschema.ResourcePoolsConfigDefaultPool{
			Allocation: expandAllocation(m.DefaultPool.Allocation),
			Priorities: expandPriorities(m.DefaultPool.Priorities),
		},
		Pool: expandPools(m.Pools),
	}
}

func resourceResourcePoolsConfig() *schema.Resource {
	return &schema.Resource{
		Schema:        tfschema.ResourcePoolsConfig,
		SchemaVersion: 1,
		Description:   "Shared admin config controlling quota usage in Chronosphere",
		CreateContext: resourceResourcePoolsConfigCreate,
		ReadContext:   resourceResourcePoolsConfigRead,
		UpdateContext: resourceResourcePoolsConfigUpdate,
		DeleteContext: resourceResourcePoolsConfigDelete,
		CustomizeDiff: resourceResourcePoolCustomizeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceResourcePoolsConfigCreate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "resource_pools_config")
	cli := getConfigClient(meta)

	poolsConfig, err := buildResourcePoolsConfig(d)
	if err != nil {
		return diag.Errorf("could not build resource pools config: %v", err)
	}

	req := &resource_pools.CreateResourcePoolsParams{
		Body: &apimodels.Configv1CreateResourcePoolsRequest{
			ResourcePools: poolsConfig,
		},
		Context: ctx,
	}
	if _, err = cli.ResourcePools.CreateResourcePools(req); err != nil {
		return diag.Errorf("could not create resource pool config: %v", err)
	}

	d.SetId(ResourcePoolsConfigID)

	return nil
}

func resourceResourcePoolsConfigRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "resource_pools_config")
	cli := getConfigClient(meta)

	resp, err := cli.ResourcePools.ReadResourcePools(&resource_pools.ReadResourcePoolsParams{Context: ctx})
	if clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to read resource pools config: %v", clienterror.Wrap(err))
	}

	serverCfg := resp.Payload.ResourcePools
	schemaConfig := ResourcePoolsConfigFromModel(serverCfg)

	// The Terraform schema has equivalent "pool" and deprecated "pools" fields that are the same on the server.
	// Take the server config and read it into whichever field (pool or pools) that the user has decided to set.
	// This avoids a meaningless diff being shown during "terraform plan".
	tfConfig := &intschema.ResourcePoolsConfig{}
	if err := tfConfig.FromResourceData(d); err != nil {
		return diag.Errorf("reading existing config: %v", err)
	}

	// pick the tf pool or pools to use.
	tfPools := tfConfig.Pool
	if len(tfConfig.Pools) > 0 {
		tfPools = tfConfig.Pools
	}

	// index tfConfig pools by name since we may be adding or removing items
	// in the new schema. We cannot assumt len(tfConfig.Pools) == len(schemaConfig.Pools)
	tfPoolsByName := map[string]intschema.ResourcePoolsConfigPool{}
	for _, p := range tfPools {
		tfPoolsByName[p.Name] = p
	}

	for i, pool := range schemaConfig.Pool {
		if tfPool, ok := tfPoolsByName[pool.Name]; ok {
			if tfPool.MatchRule != "" && len(pool.MatchRules) == 1 {
				pool.MatchRule = pool.MatchRules[0]
				pool.MatchRules = nil
				schemaConfig.Pool[i] = pool
			}
		}
	}

	if len(tfConfig.Pools) > 0 {
		schemaConfig.Pools = schemaConfig.Pool
		schemaConfig.Pool = nil
	}

	return schemaConfig.ToResourceData(d)
}

func expandAllocation(allocation *apimodels.ResourcePoolsAllocation) *intschema.ResourcePoolAllocationSchema {
	if allocation == nil {
		return nil
	}

	return &intschema.ResourcePoolAllocationSchema{PercentOfLicense: allocation.PercentOfLicense}
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

func expandPools(pools []*apimodels.ResourcePoolsPool) []intschema.ResourcePoolsConfigPool {
	return sliceutil.Map(pools, func(pool *apimodels.ResourcePoolsPool) intschema.ResourcePoolsConfigPool {
		rules := aggregationfilter.ListFromModel(pool.Filters, aggregationfilter.ResourcePoolsDelimiter)
		return intschema.ResourcePoolsConfigPool{
			Name:       pool.Name,
			MatchRules: rules,
			Allocation: expandAllocation(pool.Allocation),
			Priorities: expandPriorities(pool.Priorities),
		}
	})
}

func resourceResourcePoolsConfigUpdate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "resource_pools_config")
	cli := getConfigClient(meta)

	poolsConfig, err := buildResourcePoolsConfig(d)
	if err != nil {
		return diag.Errorf("could not build resource pools config: %v", err)
	}
	req := &resource_pools.UpdateResourcePoolsParams{
		Context: ctx,
		Body: &apimodels.Configv1UpdateResourcePoolsRequest{
			ResourcePools: poolsConfig,
		},
	}
	_, err = cli.ResourcePools.UpdateResourcePools(req)
	if err != nil {
		return diag.Errorf("unable to update resource pools config: %v", err)
	}
	return nil
}

func resourceResourcePoolsConfigDelete(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "resource_pools_config")
	cli := getConfigClient(meta)

	req := &resource_pools.DeleteResourcePoolsParams{Context: ctx}
	if _, err := cli.ResourcePools.DeleteResourcePools(req); clienterror.IsNotFound(err) {
		// Ignore 404s since it means the resource is already deleted
	} else if err != nil {
		return diag.Errorf("unable to delete resource pool config: %v", err)
	}
	d.SetId("")

	return nil
}

func resourceResourcePoolCustomizeDiff(
	_ context.Context, d *schema.ResourceDiff, meta any,
) error {
	cfg, err := buildResourcePoolsConfig(d)
	if err != nil {
		return fmt.Errorf("unable to build resource pools config: %w", err)
	}

	tfConfig := &intschema.ResourcePoolsConfig{}
	if err := tfConfig.FromResourceData(d); err != nil {
		return fmt.Errorf("reading existing config: %v", err)
	}

	for i, pool := range tfConfig.Pools {
		if (pool.MatchRule != "") == (len(pool.MatchRules) > 0) {
			return fmt.Errorf("pool %d must set exactly one match_rules or match_rule", i)
		}
	}

	return validateResourcePoolsConfig(cfg)
}

func validateResourcePoolsConfig(cfg *apimodels.Configv1ResourcePools) error {
	sum := cfg.DefaultPool.Allocation.PercentOfLicense
	for _, pool := range cfg.Pools {
		sum += pool.Allocation.PercentOfLicense

		if err := validateResourcePoolPriorities(pool.Priorities); err != nil {
			return err
		}
	}
	if sum != 100 {
		return stderrors.New("total allocation must sum to 100%")
	}

	if err := validateResourcePoolPriorities(cfg.DefaultPool.Priorities); err != nil {
		return err
	}

	return nil
}

func validateResourcePoolPriorities(priorities *apimodels.ResourcePoolsPriorities) error {
	if priorities == nil {
		return nil
	}
	if len(priorities.HighPriorityFilters) == 0 && len(priorities.LowPriorityFilters) == 0 {
		return stderrors.New("priorities have at least one of high_priority_match_rules or low_priority_match_rules")
	}
	return nil
}

func buildResourcePoolsConfig(d ResourceGetter) (*apimodels.Configv1ResourcePools, error) {
	config := &intschema.ResourcePoolsConfig{}
	if err := config.FromResourceData(d); err != nil {
		return nil, err
	}

	if len(config.Pool) > 0 && len(config.Pools) > 0 {
		return nil, fmt.Errorf("cannot set both pool and pools")
	}

	// The resource has equivalent "pool" and deprecated "pools" lists.
	// Only one is set at any given time and the server does not distinguish the different lists.
	resourcePools := config.Pool
	if len(config.Pools) > 0 {
		resourcePools = config.Pools
	}

	pools, err := sliceutil.MapErr(resourcePools, func(pool intschema.ResourcePoolsConfigPool) (*apimodels.ResourcePoolsPool, error) {
		return buildPool(pool)
	})
	if err != nil {
		return nil, err
	}
	defaultPool, err := buildDefaultPool(config.DefaultPool)
	if err != nil {
		return nil, err
	}
	return &apimodels.Configv1ResourcePools{
		DefaultPool: defaultPool,
		Pools:       pools,
	}, nil
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
	}
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
