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
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// MappingRuleFromModel maps an API model to an intschema model.
func MappingRuleFromModel(m *models.Configv1MappingRule) (*intschema.MappingRule, error) {
	return mappingRuleConverter{}.fromModel(m)
}

func resourceMappingRule() *schema.Resource {
	r := newGenericResource[
		*models.Configv1MappingRule,
		intschema.MappingRule,
		*intschema.MappingRule,
	](
		"mapping_rule",
		mappingRuleConverter{},
		generatedMappingRule{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&MappingRuleDryRunCount),
		Schema:        tfschema.MappingRule,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// MappingRuleDryRunCount tracks how many times dry run is run during validation for testing.
var MappingRuleDryRunCount atomic.Int64

type mappingRuleConverter struct{}

func (mappingRuleConverter) toModel(
	r *intschema.MappingRule,
) (*models.Configv1MappingRule, error) {
	if err := validateMappingRule(r); err != nil {
		return nil, fmt.Errorf("invalid mapping rule: %v", err)
	}
	filter, err := aggregationfilter.StringToModel(r.Filter, aggregationfilter.MappingRuleDelimiter)
	if err != nil {
		return nil, err
	}
	return &models.Configv1MappingRule{
		Slug:              r.Slug,
		Name:              r.Name,
		Drop:              r.Drop,
		Filters:           filter,
		BucketSlug:        r.BucketId.Slug(),
		AggregationPolicy: aggregationPolicyToModel(r),
		Mode:              enum.MappingModeType.V1(r.Mode),
	}, nil
}

func (mappingRuleConverter) fromModel(
	m *models.Configv1MappingRule,
) (*intschema.MappingRule, error) {
	r := &intschema.MappingRule{
		Slug:     m.Slug,
		Name:     m.Name,
		Drop:     m.Drop,
		Filter:   aggregationfilter.StringFromModel(m.Filters, aggregationfilter.MappingRuleDelimiter),
		BucketId: tfid.Slug(m.BucketSlug),
		Mode:     string(m.Mode),
	}
	if m.AggregationPolicy != nil {
		r.Aggregations = aggregationFromModel(m.AggregationPolicy.Aggregation)
		r.StoragePolicy = mappingStoragePolicyFromModel(m.AggregationPolicy.StoragePolicy)
		r.DropTimestamp = m.AggregationPolicy.DropTimestamp
		r.Interval = m.AggregationPolicy.Interval
	}
	return r, nil
}

func aggregationPolicyToModel(
	r *intschema.MappingRule,
) *models.MappingRuleAggregationPolicy {
	a := models.MappingRuleAggregationPolicy{
		Aggregation:   aggregationToModel(r.Aggregations),
		DropTimestamp: r.DropTimestamp,
		StoragePolicy: mappingStoragePolicyToModel(r.StoragePolicy),
		Interval:      r.Interval,
	}
	if a == (models.MappingRuleAggregationPolicy{}) {
		return nil
	}
	return &a
}

func aggregationToModel(aggs []string) models.Configv1AggregationType {
	if len(aggs) == 0 {
		return ""
	}
	// "aggregations" uses MaxItems=1, so we can assume there are never multiple
	// elements.
	return enum.AggregationType.V1(aggs[0])
}

func aggregationFromModel(m models.Configv1AggregationType) []string {
	if m == "" {
		return nil
	}
	return []string{string(m)}
}

func mappingStoragePolicyToModel(
	p *intschema.MappingRuleStoragePolicy,
) *models.Configv1MappingRuleStoragePolicy {
	if p == nil {
		return nil
	}
	return &models.Configv1MappingRuleStoragePolicy{
		Resolution: p.Resolution,
		Retention:  p.Retention,
	}
}

func mappingStoragePolicyFromModel(
	p *models.Configv1MappingRuleStoragePolicy,
) *intschema.MappingRuleStoragePolicy {
	if p == nil {
		return nil
	}
	return &intschema.MappingRuleStoragePolicy{
		Resolution: p.Resolution,
		Retention:  p.Retention,
	}
}

func validateMappingRule(rule *intschema.MappingRule) error {
	if rule.Drop {
		if len(rule.Aggregations) > 0 {
			return fmt.Errorf("cannot set aggregations when drop is true")
		}
		if rule.StoragePolicy != nil {
			return fmt.Errorf("cannot set storage_policy when drop is true")
		}
		return nil
	}

	if rule.StoragePolicy == nil && len(rule.Aggregations) == 0 {
		return fmt.Errorf("must set aggregations and storage_policy when drop is not set")
	}

	// Aggregations and StoragePolicies must be set together.
	if rule.StoragePolicy != nil && len(rule.Aggregations) == 0 {
		return fmt.Errorf("must set aggregations")
	}

	return nil
}
