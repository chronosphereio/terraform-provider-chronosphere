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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// DropRuleFromModel maps an API model to an intschema model.
func DropRuleFromModel(m *models.Configv1DropRule) (*intschema.DropRule, error) {
	return dropRuleConverter{}.fromModel(m)
}

func resourceDropRule() *schema.Resource {
	r := newGenericResource(
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

	// Mode favored over deprecated active field.
	if !r.Active {
		// Purposeful breaking change that will be encountered whenever the legacy `active` field was set to false.
		// This is in order to transition to the use of `mode` field.
		// Active is now entirely ignored and treated as always `true` (its default value). Mode is the field that actually drives behavior.
		return nil, errors.New("must set `mode` instead of `active`")
	}

	return &models.Configv1DropRule{
		Name:                     r.Name,
		Slug:                     r.Slug,
		Mode:                     enum.DropRuleModeType.V1(r.Mode),
		Filters:                  filter,
		ConditionalRateBasedDrop: conditionalRateBasedDrop,
		DropNanValue:             r.DropNanValue,
		ValueBasedDrop:           valueBaseDropToModel(r.ValueBasedDrop),
	}, nil
}

func (dropRuleConverter) fromModel(
	m *models.Configv1DropRule,
) (*intschema.DropRule, error) {
	r := &intschema.DropRule{
		Name: m.Name,
		Slug: m.Slug,
		// Active should already be treated as true to avoid diffs when mode is omitted.
		// The actual source of truth is the mode field.
		Active:         true,
		Mode:           string(m.Mode),
		Query:          aggregationfilter.ListFromModel(m.Filters, aggregationfilter.DropRuleDelimiter),
		DropNanValue:   m.DropNanValue,
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
