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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// LogscaleAlertFromModel maps an API model to the intschema model.
func LogscaleAlertFromModel(m *models.Configv1LogScaleAlert) (*intschema.LogscaleAlert, error) {
	return (logscaleAlertConverter{}).fromModel(m)
}

func resourceLogscaleAlert() *schema.Resource {
	r := newGenericResource[
		*models.Configv1LogScaleAlert,
		intschema.LogscaleAlert,
		*intschema.LogscaleAlert,
	](
		"logscale_alert",
		logscaleAlertConverter{},
		generatedLogScaleAlert{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.LogscaleAlert,
		CustomizeDiff: r.ValidateDryRunOptions(
			&LogscaleAlertDryRunCount,
			// ignore the action_ids field for unknown references check
			ValidateDryRunOpts[*models.Configv1LogScaleAlert]{
				SetUnknownReferencesSkip: []string{"action_ids.[]"},
			}),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// LogscaleAlertDryRunCount tracks how many times dry run is run during validation for testing.
var LogscaleAlertDryRunCount atomic.Int64

type logscaleAlertConverter struct{}

func (logscaleAlertConverter) toModel(
	c *intschema.LogscaleAlert,
) (*models.Configv1LogScaleAlert, error) {
	if c == nil {
		return nil, nil
	}

	throttleSecs, err := durationToSecs(c.ThrottleDuration)
	if err != nil {
		return nil, err
	}
	timeWindowSecs, err := durationToSecs(c.TimeWindow)
	if err != nil {
		return nil, err
	}
	return &models.Configv1LogScaleAlert{
		Name:                c.Name,
		Slug:                c.Slug,
		Repository:          c.Repository,
		AlertType:           models.LogScaleAlertAlertType(c.AlertType),
		Description:         c.Description,
		Disabled:            c.Disabled,
		LogScaleActionSlugs: sliceutil.Map(c.ActionIds, tfid.ID.Slug),
		LogScaleQuery:       c.Query,
		Tags:                c.Tags,
		ThrottleSecs:        throttleSecs,
		TimeWindowSecs:      timeWindowSecs,
		RunAsUser:           c.RunAsUser,
		ThrottleField:       c.ThrottleField,
	}, nil
}

func (logscaleAlertConverter) fromModel(
	m *models.Configv1LogScaleAlert,
) (*intschema.LogscaleAlert, error) {
	return &intschema.LogscaleAlert{
		Name:             m.Name,
		Slug:             m.Slug,
		Repository:       m.Repository,
		AlertType:        string(m.AlertType),
		Description:      m.Description,
		Disabled:         m.Disabled,
		ActionIds:        sliceutil.Map(m.LogScaleActionSlugs, tfid.Slug),
		Query:            m.LogScaleQuery,
		Tags:             m.Tags,
		ThrottleDuration: durationFromSecs(m.ThrottleSecs),
		TimeWindow:       durationFromSecs(m.TimeWindowSecs),
		RunAsUser:        m.RunAsUser,
		ThrottleField:    m.ThrottleField,
	}, nil
}
