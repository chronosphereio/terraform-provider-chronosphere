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
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_tail_sampling_rules"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceTraceTailSamplingRules() *schema.Resource {
	return &schema.Resource{
		Schema:        tfschema.TraceTailSamplingRules,
		CreateContext: resourceTraceTailSamplingRulesCreate,
		ReadContext:   resourceTraceTailSamplingRulesRead,
		UpdateContext: resourceTraceTailSamplingRulesUpdate,
		DeleteContext: resourceTraceTailSamplingRulesDelete,
		CustomizeDiff: resourceTraceTailSamplingRulesCustomizeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// -----
// toModel and related helpers.
// -----

func toModel(
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
	return &models.Configv1TraceTailSamplingRule{
		Name:       r.Name,
		SystemName: r.SystemName,
		Filter:     traceSearchFilterToModel(r.Filter),
		SampleRate: r.SampleRate,
	}, nil
}

// -----
// fromModel and related helpers.
// -----

func fromModel(
	m *models.Configv1TraceTailSamplingRules,
) *intschema.TraceTailSamplingRules {
	return &intschema.TraceTailSamplingRules{
		DefaultSampleRate: defaultSampleRateFromModel(m.DefaultSampleRate),
		Rules:             sliceutil.Map(m.Rules, ruleFromModel),
	}
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

func resourceTraceTailSamplingRulesCreate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigClient(meta)

	rules, err := buildTraceTailSamplingRules(d)
	if err != nil {
		return diag.Errorf("could not build trace tail sampling rules: %v", err)
	}
	req := &trace_tail_sampling_rules.CreateTraceTailSamplingRulesParams{
		Body: &models.Configv1CreateTraceTailSamplingRulesRequest{
			TraceTailSamplingRules: rules,
		},
		Context: ctx,
	}

	if _, err := cli.TraceTailSamplingRules.CreateTraceTailSamplingRules(req); err != nil {
		return diag.Errorf("could not create trace tail sampling rules: %v", err)
	}

	d.SetId(TraceTailSamplingRulesID)

	return nil
}

func resourceTraceTailSamplingRulesRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigClient(meta)

	resp, err := cli.TraceTailSamplingRules.ReadTraceTailSamplingRules(&trace_tail_sampling_rules.ReadTraceTailSamplingRulesParams{Context: ctx})
	if clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to read trace tail sampling rules: %v", clienterror.Wrap(err))
	}

	rules := resp.Payload.TraceTailSamplingRules
	schemaRules := fromModel(rules)

	if err := schemaRules.ToResourceData(d); err != nil {
		return err
	}
	d.SetId(TraceTailSamplingRulesID)
	return nil
}

func resourceTraceTailSamplingRulesUpdate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigClient(meta)

	rules, err := buildTraceTailSamplingRules(d)
	if err != nil {
		return diag.Errorf("could not build trace tail sampling rules: %v", err)
	}
	req := &trace_tail_sampling_rules.UpdateTraceTailSamplingRulesParams{
		Context: ctx,
		Body: &models.Configv1UpdateTraceTailSamplingRulesRequest{
			TraceTailSamplingRules: rules,
		},
	}
	if _, err := cli.TraceTailSamplingRules.UpdateTraceTailSamplingRules(req); err != nil {
		return diag.Errorf("unable to update trace tail sampling rules: %v", err)
	}
	return nil
}

func resourceTraceTailSamplingRulesDelete(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "trace_tail_sampling_rules")
	cli := getConfigClient(meta)

	req := &trace_tail_sampling_rules.DeleteTraceTailSamplingRulesParams{Context: ctx}
	if _, err := cli.TraceTailSamplingRules.DeleteTraceTailSamplingRules(req); clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to delete trace tail sampling rules: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTraceTailSamplingRulesCustomizeDiff(
	_ context.Context, d *schema.ResourceDiff, meta any,
) error {
	rules, err := buildTraceTailSamplingRules(d)
	if err != nil {
		return fmt.Errorf("unable to build trace tail sampling rules: %w", err)
	}
	return validateTraceTailSamplingRules(rules)
}

func validateTraceTailSamplingRules(rules *models.Configv1TraceTailSamplingRules) error {
	for _, r := range rules.Rules {
		if r.SampleRate < 0 || r.SampleRate > 1.0 {
			return fmt.Errorf("expected sample rate to be a float from 0 to 1.0 inclusive, got %f", r.SampleRate)
		}
	}

	if rules.DefaultSampleRate.SampleRate < 0 || rules.DefaultSampleRate.SampleRate > 1.0 {
		return fmt.Errorf("expected sample rate to be a float from 0 to 1.0 inclusive, got %f", rules.DefaultSampleRate.SampleRate)
	}

	return nil
}

func buildTraceTailSamplingRules(d ResourceGetter) (*models.Configv1TraceTailSamplingRules, error) {
	rules := &intschema.TraceTailSamplingRules{}
	if err := rules.FromResourceData(d); err != nil {
		return nil, err
	}

	return toModel(rules)
}
