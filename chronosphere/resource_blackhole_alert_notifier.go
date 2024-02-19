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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func BlackholeAlertNotifierFromModel(
	n *models.Configv1Notifier,
) (*intschema.BlackholeAlertNotifier, error) {
	return blackholeAlertNotifierConverter{}.fromModel(n)
}

func resourceBlackHoleAlertNotifier() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Notifier,
		intschema.BlackholeAlertNotifier,
		*intschema.BlackholeAlertNotifier,
	](
		"blackhole_alert_notifier",
		blackholeAlertNotifierConverter{},
		generatedNotifier{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.BlackholeAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&BlackHoleAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// BlackHoleAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var BlackHoleAlertNotifierDryRunCount atomic.Int64

type blackholeAlertNotifierConverter struct{}

func (blackholeAlertNotifierConverter) toModel(
	n *intschema.BlackholeAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Name:    n.Name,
		Slug:    n.Slug,
		Discard: true,
	}, nil
}

func (blackholeAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.BlackholeAlertNotifier, error) {
	if !m.Discard {
		return &intschema.BlackholeAlertNotifier{
			Name: m.Name + notifierTypeChangedName("blackhole"),
			Slug: m.Slug,
		}, nil
	}
	return &intschema.BlackholeAlertNotifier{
		Name: m.Name,
		Slug: m.Slug,
	}, nil
}
