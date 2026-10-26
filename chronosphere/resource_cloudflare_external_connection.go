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

func CloudflareExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.CloudflareExternalConnection, error) {
	return cloudflareExternalConnectionConverter{}.fromModel(m)
}

func resourceCloudflareExternalConnection() *schema.Resource {
	r := newGenericResource(
		"cloudflare_external_connection",
		cloudflareExternalConnectionConverter{},
		generatedExternalConnection{},
	)
	validateDryRun := r.ValidateDryRun(&CloudflareExternalConnectionDryRunCount)
	validateDryRunWithTokenPlaceholder := r.ValidateDryRunOptions(&CloudflareExternalConnectionDryRunCount, ValidateDryRunOpts[*models.Configv1ExternalConnection]{
		DryRunDefaults: map[string]any{
			"api_token": "dry_run_unknown_token",
		},
	})

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description:   "Cloudflare API token that a `chronosphere_cloud_integration` with `provider_type = \"cloudflare\"` references through `external_connection_slug` to collect Cloudflare metrics.",
		Schema:        tfschema.CloudflareExternalConnection,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
			// A token produced by another resource in the same apply is unknown at plan
			// time, and the API rejects a connection without one.
			if !diff.NewValueKnown("api_token") {
				return validateDryRunWithTokenPlaceholder(ctx, diff, meta)
			}
			return validateDryRun(ctx, diff, meta)
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CloudflareExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var CloudflareExternalConnectionDryRunCount atomic.Int64

type cloudflareExternalConnectionConverter struct{}

func (cloudflareExternalConnectionConverter) toModel(
	n *intschema.CloudflareExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		Cloudflare: &models.ExternalConnectionCloudflareConfig{
			APIToken: n.ApiToken,
		},
	}, nil
}

func (cloudflareExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.CloudflareExternalConnection, error) {
	c := m.Cloudflare
	if c == nil {
		return &intschema.CloudflareExternalConnection{
			Name: m.Name + " (not a cloudflare external connection)",
			Slug: m.Slug,
		}, nil
	}
	return &intschema.CloudflareExternalConnection{
		Name:     m.Name,
		Slug:     m.Slug,
		ApiToken: c.APIToken,
	}, nil
}

func (cloudflareExternalConnectionConverter) normalize(config, server *intschema.CloudflareExternalConnection) {
	preserveRedactedSecret(&server.ApiToken, config.ApiToken)
}
