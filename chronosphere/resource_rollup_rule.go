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
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// RollupRuleDryRunCount tracks how many times dry run is run during validation for testing.
var RollupRuleDryRunCount atomic.Int64

// RollupRuleFromModel maps an API model to an intschema model.
func RollupRuleFromModel(m *models.Configv1RollupRule) (*intschema.RollupRule, error) {
	return rollupRuleConverter{}.fromModel(m)
}

func resourceRollupRule() *schema.Resource {
	r := newGenericResource[
		*models.Configv1RollupRule,
		intschema.RollupRule,
		*intschema.RollupRule,
	](
		"rollup_rule",
		rollupRuleConverter{},
		generatedRollupRule{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			rr := &intschema.RollupRule{}
			if err := rr.FromResourceData(d); err != nil {
				return err
			}
			return r.ValidateDryRun(&RollupRuleDryRunCount)(ctx, d, m)
		},
		Schema: tfschema.RollupRule,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

type rollupRuleConverter struct{}

func (rollupRuleConverter) toModel(
	r *intschema.RollupRule,
) (*models.Configv1RollupRule, error) {
	// The schema doesn't allow interval and policies to be set together.
	// However, we receive both because storage_policies is an optional computed field
	// which will send the last server-returned value even if the user doesn't have it set
	// in the configuration. Hence we clear it manually to workaround sc-81013.
	if r.Interval != "" && r.StoragePolicies != nil {
		r.StoragePolicies = nil
	}

	filter, err := aggregationfilter.StringToModel(r.Filter, aggregationfilter.RollupRuleDelimiter)
	if err != nil {
		return nil, err
	}
	return &models.Configv1RollupRule{
		BucketSlug:         r.BucketId.Slug(),
		Slug:               r.Slug,
		Name:               r.Name,
		Filters:            filter,
		MetricName:         r.NewMetric,
		StoragePolicy:      rollupStoragePolicyToModel(r.StoragePolicies),
		AddMetricTypeLabel: r.MetricTypeTag,
		MetricType:         enum.MetricType.V1(r.MetricType),
		Aggregation:        enum.AggregationType.V1(r.Aggregation),
		DropRaw:            r.DropRaw,
		ExpansiveMatch:     r.Permissive,
		Mode:               enum.RollupModeType.V1(r.Mode),
		LabelPolicy: &models.Configv1RollupRuleLabelPolicy{
			Keep:    r.GroupBy,
			Discard: r.ExcludeBy,
		},
		Interval:            r.Interval,
		GraphiteLabelPolicy: rollupGraphiteLabelPolicyToModel(r.GraphiteLabelPolicy),
	}, nil
}

func (rollupRuleConverter) fromModel(
	m *models.Configv1RollupRule,
) (*intschema.RollupRule, error) {
	r := &intschema.RollupRule{
		BucketId:            tfid.Slug(m.BucketSlug),
		Name:                m.Name,
		Slug:                m.Slug,
		Filter:              aggregationfilter.StringFromModel(m.Filters, aggregationfilter.RollupRuleDelimiter),
		MetricType:          string(m.MetricType),
		Aggregation:         string(m.Aggregation),
		DropRaw:             m.DropRaw,
		MetricTypeTag:       m.AddMetricTypeLabel,
		NewMetric:           m.MetricName,
		Permissive:          m.ExpansiveMatch,
		StoragePolicies:     rollupStoragePolicyFromModel(m.StoragePolicy),
		Mode:                string(m.Mode),
		Interval:            m.Interval,
		GraphiteLabelPolicy: rollupGraphiteLabelPolicyFromModel(m.GraphiteLabelPolicy),
	}
	if m.LabelPolicy != nil {
		r.GroupBy = m.LabelPolicy.Keep
		r.ExcludeBy = m.LabelPolicy.Discard
	}
	return r, nil
}

func rollupStoragePolicyToModel(
	p *intschema.RollupRuleStoragePolicies,
) *models.Configv1RollupRuleStoragePolicy {
	if p == nil {
		return nil
	}
	return &models.Configv1RollupRuleStoragePolicy{
		Resolution: p.Resolution,
		Retention:  p.Retention,
	}
}

func rollupStoragePolicyFromModel(
	p *models.Configv1RollupRuleStoragePolicy,
) *intschema.RollupRuleStoragePolicies {
	if p == nil {
		return nil
	}
	return &intschema.RollupRuleStoragePolicies{
		Resolution: p.Resolution,
		Retention:  p.Retention,
	}
}

func rollupGraphiteLabelPolicyToModel(
	p *intschema.RollupRuleGraphiteLabelPolicy,
) *models.RollupRuleGraphiteLabelPolicy {
	if p == nil {
		return nil
	}
	return &models.RollupRuleGraphiteLabelPolicy{
		Replace: sliceutil.Map(
			p.Replace,
			func(r intschema.RollupRuleGraphiteLabelPolicyReplace) *models.GraphiteLabelPolicyReplace {
				return &models.GraphiteLabelPolicyReplace{
					Name:     r.Name,
					NewValue: r.NewValue,
				}
			}),
	}
}

func rollupGraphiteLabelPolicyFromModel(
	p *models.RollupRuleGraphiteLabelPolicy,
) *intschema.RollupRuleGraphiteLabelPolicy {
	if p == nil {
		return nil
	}
	return &intschema.RollupRuleGraphiteLabelPolicy{
		Replace: sliceutil.Map(
			p.Replace,
			func(r *models.GraphiteLabelPolicyReplace) intschema.RollupRuleGraphiteLabelPolicyReplace {
				return intschema.RollupRuleGraphiteLabelPolicyReplace{
					Name:     r.Name,
					NewValue: r.NewValue,
				}
			}),
	}
}
