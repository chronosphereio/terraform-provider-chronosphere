// Copyright 2026 Chronosphere Inc.
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
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// QueryResourcePoolsFromModel maps an API model into an intschema model.
func QueryResourcePoolsFromModel(m *models.ConfigunstableQueryResourcePools) (*intschema.QueryResourcePools, error) {
	return (queryResourcePoolsConverter{}).fromModel(m)
}

func resourceQueryResourcePools() *schema.Resource {
	r := newGenericResource(
		"query_resource_pools",
		queryResourcePoolsConverter{},
		generatedUnstableQueryResourcePools{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.QueryResourcePools,
		CustomizeDiff: r.ValidateDryRun(&QueryResourcePoolsDryRunCount),
		Description:   "Singleton grouping of automated metric query sources (monitors, recording rules, SLOs, and service accounts) into named pools, each with a data read limit that rejects the pool's queries when sustained. " + unstableAPIWarning,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// QueryResourcePoolsDryRunCount tracks how many times dry run is run during validation for testing.
var QueryResourcePoolsDryRunCount atomic.Int64

type queryResourcePoolsConverter struct{}

func (queryResourcePoolsConverter) toModel(
	r *intschema.QueryResourcePools,
) (*models.ConfigunstableQueryResourcePools, error) {
	m := &models.ConfigunstableQueryResourcePools{
		Pools: sliceutil.Map(r.Pool, queryPoolToModel),
	}
	if r.DefaultPool != nil {
		m.DefaultPool = &models.QueryResourcePoolsDefaultPool{
			DataReadLimit: queryPoolLimitToModel(r.DefaultPool.DataReadLimit),
		}
	}
	return m, nil
}

func (queryResourcePoolsConverter) fromModel(
	m *models.ConfigunstableQueryResourcePools,
) (*intschema.QueryResourcePools, error) {
	pools, err := sliceutil.MapErr(m.Pools, queryPoolFromModel)
	if err != nil {
		return nil, err
	}
	r := &intschema.QueryResourcePools{
		Pool: pools,
	}
	if m.DefaultPool != nil {
		limit, err := queryPoolLimitFromModel(m.DefaultPool.DataReadLimit)
		if err != nil {
			return nil, err
		}
		r.DefaultPool = &intschema.QueryResourcePoolsDefaultPool{
			DataReadLimit: limit,
		}
	}
	return r, nil
}

func queryPoolToModel(pool intschema.QueryResourcePoolsPool) *models.QueryResourcePoolsPool {
	return &models.QueryResourcePoolsPool{
		Name: pool.Name,
		SourceTypes: sliceutil.Map(pool.SourceTypes, func(t string) models.QueryResourcePoolsSourceType {
			return models.QueryResourcePoolsSourceType(t)
		}),
		SourceNames:   pool.SourceNames,
		DataReadLimit: queryPoolLimitToModel(pool.DataReadLimit),
	}
}

func queryPoolFromModel(pool *models.QueryResourcePoolsPool) (intschema.QueryResourcePoolsPool, error) {
	limit, err := queryPoolLimitFromModel(pool.DataReadLimit)
	if err != nil {
		return intschema.QueryResourcePoolsPool{}, err
	}
	return intschema.QueryResourcePoolsPool{
		Name: pool.Name,
		SourceTypes: sliceutil.Map(pool.SourceTypes, func(t models.QueryResourcePoolsSourceType) string {
			return string(t)
		}),
		SourceNames:   pool.SourceNames,
		DataReadLimit: limit,
	}, nil
}

func queryPoolLimitToModel(limit *intschema.QueryResourcePoolDataReadLimitSchema) *models.QueryResourcePoolsDataReadLimit {
	if limit == nil {
		return nil
	}
	return &models.QueryResourcePoolsDataReadLimit{
		// Always send the max explicitly: the API requires it whenever a
		// data read limit is set, and "0" means block all queries.
		MaxDatapointsReadPerSecond: strconv.FormatInt(limit.MaxDatapointsReadPerSecond, 10),
		SustainSecs:                int32(limit.SustainSecs),
	}
}

func queryPoolLimitFromModel(limit *models.QueryResourcePoolsDataReadLimit) (*intschema.QueryResourcePoolDataReadLimitSchema, error) {
	if limit == nil {
		return nil, nil
	}
	var (
		maxDatapoints int64
		err           error
	)
	if limit.MaxDatapointsReadPerSecond != "" {
		maxDatapoints, err = strconv.ParseInt(limit.MaxDatapointsReadPerSecond, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return &intschema.QueryResourcePoolDataReadLimitSchema{
		MaxDatapointsReadPerSecond: maxDatapoints,
		SustainSecs:                int64(limit.SustainSecs),
	}, nil
}
