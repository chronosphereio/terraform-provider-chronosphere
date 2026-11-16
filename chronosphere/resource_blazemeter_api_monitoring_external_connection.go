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
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func BlazemeterApiMonitoringExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.BlazemeterApiMonitoringExternalConnection, error) {
	return blazemeterApiMonitoringExternalConnectionConverter{}.fromModel(m)
}

func resourceBlazemeterApiMonitoringExternalConnection() *schema.Resource {
	r := newGenericResource(
		"blazemeter_api_monitoring_external_connection",
		blazemeterApiMonitoringExternalConnectionConverter{},
		generatedExternalConnection{},
	)
	validateDryRun := r.ValidateDryRun(&BlazemeterApiMonitoringExternalConnectionDryRunCount)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description: "BlazeMeter API Monitoring access token that a `chronosphere_cloud_integration` with " +
			"`provider_type = \"blazemeter_api_monitoring\"` references through `external_connection_slug` to collect " +
			"BlazeMeter API Monitoring test results.",
		Schema: tfschema.BlazemeterApiMonitoringExternalConnection,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
			// The API rejects a connection without a token, so a token produced by
			// another resource in the same apply cannot be validated at plan time.
			if !diff.NewValueKnown("access_token") {
				return nil
			}
			return validateDryRun(ctx, diff, meta)
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// BlazemeterApiMonitoringExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var BlazemeterApiMonitoringExternalConnectionDryRunCount atomic.Int64

type blazemeterApiMonitoringExternalConnectionConverter struct{}

func (blazemeterApiMonitoringExternalConnectionConverter) toModel(
	n *intschema.BlazemeterApiMonitoringExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		BlazemeterAPIMonitoring: &models.ExternalConnectionBlazemeterAPIMonitoringConfig{
			AccessToken: n.AccessToken,
		},
	}, nil
}

func (blazemeterApiMonitoringExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.BlazemeterApiMonitoringExternalConnection, error) {
	c := m.BlazemeterAPIMonitoring
	if c == nil {
		return &intschema.BlazemeterApiMonitoringExternalConnection{
			Name: m.Name + " (not a blazemeter_api_monitoring external connection)",
			Slug: m.Slug,
		}, nil
	}
	return &intschema.BlazemeterApiMonitoringExternalConnection{
		Name:        m.Name,
		Slug:        m.Slug,
		AccessToken: c.AccessToken,
	}, nil
}

func (blazemeterApiMonitoringExternalConnectionConverter) normalize(
	config, server *intschema.BlazemeterApiMonitoringExternalConnection,
) {
	server.AccessToken = config.AccessToken
}
