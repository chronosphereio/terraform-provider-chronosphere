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
	"github.com/pkg/errors"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func TraceTailSamplingRulesFromModel(m *models.Configv1TraceTailSamplingRules) (*intschema.TraceTailSamplingRules, error) {
	return traceTailSamplingRulesConverter{}.fromModel(m)
}

func resourceTraceTailSamplingRules() *schema.Resource {
	r := newGenericResource(
		"trace_tail_sampling_rules",
		traceTailSamplingRulesConverter{},
		generatedTraceTailSamplingRules{},
	)
	return &schema.Resource{
		Schema:        tfschema.TraceTailSamplingRules,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&TraceTailSamplingRulesDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// TraceTailSamplingRulesDryRunCount tracks how many times dry run is run during validation for testing.
var TraceTailSamplingRulesDryRunCount atomic.Int64

type traceTailSamplingRulesConverter struct{}

// -----
// toModel and related helpers.
// -----

func (traceTailSamplingRulesConverter) toModel(
	m *intschema.TraceTailSamplingRules,
) (*models.Configv1TraceTailSamplingRules, error) {
	spanFilters, err := sliceutil.MapErr(m.Rules, ruleToModel)
	if err != nil {
		return nil, err
	}
	return &models.Configv1TraceTailSamplingRules{
		DefaultSampleRate: defaultSampleRateToModel(m.DefaultSampleRate),
		Rules:             spanFilters,
	}, nil
}

func defaultSampleRateToModel(
	t *intschema.TraceTailSamplingRulesDefaultSampleRate,
) *models.Configv1DefaultSampleRate {
	if t == nil {
		return nil
	}
	return &models.Configv1DefaultSampleRate{
		Enabled:    t.Enabled,
		SampleRate: t.SampleRate,
	}
}

func ruleToModel(r intschema.TraceTailSamplingRulesRules) (*models.Configv1TraceTailSamplingRule, error) {
	filter, err := traceSearchFilterToModel(r.Filter)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &models.Configv1TraceTailSamplingRule{
		Name:       r.Name,
		SystemName: r.SystemName,
		Filter:     filter,
		SampleRate: r.SampleRate,
	}, nil
}

// -----
// fromModel and related helpers.
// -----

func (traceTailSamplingRulesConverter) fromModel(
	m *models.Configv1TraceTailSamplingRules,
) (*intschema.TraceTailSamplingRules, error) {
	return &intschema.TraceTailSamplingRules{
		DefaultSampleRate: defaultSampleRateFromModel(m.DefaultSampleRate),
		Rules:             sliceutil.Map(m.Rules, ruleFromModel),
	}, nil
}

func defaultSampleRateFromModel(
	r *models.Configv1DefaultSampleRate,
) *intschema.TraceTailSamplingRulesDefaultSampleRate {
	if r == nil {
		return nil
	}
	return &intschema.TraceTailSamplingRulesDefaultSampleRate{
		Enabled:    r.Enabled,
		SampleRate: r.SampleRate,
	}
}

func ruleFromModel(
	r *models.Configv1TraceTailSamplingRule,
) intschema.TraceTailSamplingRulesRules {
	return intschema.TraceTailSamplingRulesRules{
		Filter:     traceSearchFilterFromModel(r.Filter),
		SampleRate: r.SampleRate,
		SystemName: r.SystemName,
		Name:       r.Name,
	}
}
