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
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// RecordingRuleFromModel maps an API model to an intschema model.
func RecordingRuleFromModel(m *models.Configv1RecordingRule) (*intschema.RecordingRule, error) {
	return recordingRuleConverter{}.fromModel(m)
}

func resourceRecordingRule() *schema.Resource {
	r := newGenericResource(
		"recording_rule",
		recordingRuleConverter{},
		generatedRecordingRule{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.RecordingRule,
		CustomizeDiff: r.ValidateDryRunOptions(&RecordingRuleDryRunCount, ValidateDryRunOpts[*models.Configv1RecordingRule]{
			ModifyAPIModel: func(rr *models.Configv1RecordingRule) {
				// If bucket-slug and execution group are set, server validation will require them to match.
				// However, BucketSlug may not be known at dry-run time, and fallback to the unknown ref value.
				// This won't match ExecutionGroup, causing dry-run to fail, so set it to match execution group.
				if rr.BucketSlug == dryRunUnknownRef.Slug() && rr.ExecutionGroup != "" {
					rr.BucketSlug = rr.ExecutionGroup
				}
			},
		}),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// RecordingRuleDryRunCount tracks how many times dry run is run during validation for testing.
var RecordingRuleDryRunCount atomic.Int64

type recordingRuleConverter struct{}

func (recordingRuleConverter) toModel(
	r *intschema.RecordingRule,
) (*models.Configv1RecordingRule, error) {
	intervalSecs, err := durationToSecs(r.Interval)
	if err != nil {
		return nil, err
	}
	return &models.Configv1RecordingRule{
		Name:           r.Name,
		Slug:           r.Slug,
		BucketSlug:     r.BucketId.Slug(),
		ExecutionGroup: r.ExecutionGroup.Slug(),
		PrometheusExpr: r.Expr,
		IntervalSecs:   intervalSecs,
		LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
			Add: r.Labels,
		},
		MetricName: r.MetricName,
	}, nil
}

func (recordingRuleConverter) fromModel(
	m *models.Configv1RecordingRule,
) (*intschema.RecordingRule, error) {
	r := &intschema.RecordingRule{
		Name:           m.Name,
		Slug:           m.Slug,
		BucketId:       tfid.Slug(m.BucketSlug),
		ExecutionGroup: tfid.Slug(m.ExecutionGroup),
		Expr:           m.PrometheusExpr,
		Interval:       durationFromSecs(m.IntervalSecs),
		MetricName:     m.MetricName,
	}
	if m.LabelPolicy != nil {
		r.Labels = m.LabelPolicy.Add
	}
	return r, nil
}
