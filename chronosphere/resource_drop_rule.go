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
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DropRuleFromModel maps an API model to an intschema model.
func DropRuleFromModel(m *models.Configv1DropRule) (*intschema.DropRule, error) {
	return dropRuleConverter{}.fromModel(m)
}

func resourceDropRule() *schema.Resource {
	r := newGenericResource[
		*models.Configv1DropRule,
		intschema.DropRule,
		*intschema.DropRule,
	](
		"drop_rule",
		dropRuleConverter{},
		generatedDropRule{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&DropRuleDryRunCount),
		Schema:        tfschema.DropRule,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// DropRuleDryRunCount tracks how many times dry run is run during validation for testing.
var DropRuleDryRunCount atomic.Int64

type dropRuleConverter struct{}

func (dropRuleConverter) toModel(
	r *intschema.DropRule,
) (*models.Configv1DropRule, error) {
	conditionalRateBasedDrop, err := conditionalRateBasedDropToModel(r)
	if err != nil {
		return nil, err
	}
	filter, err := aggregationfilter.ListToModel(r.Query, aggregationfilter.DropRuleDelimiter)
	if err != nil {
		return nil, err
	}
	return &models.Configv1DropRule{
		Name:                     r.Name,
		Slug:                     r.Slug,
		Mode:                     dropRuleModeToModel(r.Active),
		Filters:                  filter,
		ConditionalRateBasedDrop: conditionalRateBasedDrop,
		ValueBasedDrop:           valueBaseDropToModel(r.ValueBasedDrop),
	}, nil
}

func (dropRuleConverter) fromModel(
	m *models.Configv1DropRule,
) (*intschema.DropRule, error) {
	r := &intschema.DropRule{
		Name:           m.Name,
		Slug:           m.Slug,
		Active:         m.Mode == models.Configv1DropRuleModeENABLED,
		Query:          aggregationfilter.ListFromModel(m.Filters, aggregationfilter.DropRuleDelimiter),
		ValueBasedDrop: valueBasedDropFromModel(m.ValueBasedDrop),
	}
	if m.ConditionalRateBasedDrop != nil {
		r.ActivatedDropDuration = durationFromSecs(
			m.ConditionalRateBasedDrop.ActivatedDropDurationSecs)
		r.ConditionalDrop = m.ConditionalRateBasedDrop.Enabled
		r.RateLimitThreshold = m.ConditionalRateBasedDrop.RateLimitThreshold
	}
	return r, nil
}

func conditionalRateBasedDropToModel(
	r *intschema.DropRule,
) (*models.DropRuleConditionalRateBasedDrop, error) {
	activatedDropDurationSecs, err := durationToSecs(r.ActivatedDropDuration)
	if err != nil {
		return nil, err
	}
	d := models.DropRuleConditionalRateBasedDrop{
		ActivatedDropDurationSecs: activatedDropDurationSecs,
		Enabled:                   r.ConditionalDrop,
		RateLimitThreshold:        r.RateLimitThreshold,
	}
	if d == (models.DropRuleConditionalRateBasedDrop{}) {
		return nil, nil
	}
	return &d, nil
}

func valueBasedDropFromModel(
	p *models.DropRuleValueBasedDrop,
) *intschema.DropRuleValueBasedDrop {
	if p == nil {
		return nil
	}
	if !p.Enabled {
		return nil
	}
	return &intschema.DropRuleValueBasedDrop{
		TargetDropValue: p.TargetDropValue,
	}
}

func valueBaseDropToModel(
	p *intschema.DropRuleValueBasedDrop,
) *models.DropRuleValueBasedDrop {
	if p == nil {
		return nil
	}
	return &models.DropRuleValueBasedDrop{
		Enabled:         true,
		TargetDropValue: p.TargetDropValue,
	}
}

func dropRuleModeToModel(active bool) models.Configv1DropRuleMode {
	if active {
		return models.Configv1DropRuleModeENABLED
	}
	return models.Configv1DropRuleModeDISABLED
}
