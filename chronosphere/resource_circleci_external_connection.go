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

func CircleciExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.CircleciExternalConnection, error) {
	return circleciExternalConnectionConverter{}.fromModel(m)
}

func resourceCircleciExternalConnection() *schema.Resource {
	r := newGenericResource(
		"circleci_external_connection",
		circleciExternalConnectionConverter{},
		generatedExternalConnection{},
	)
	validateDryRun := r.ValidateDryRun(&CircleciExternalConnectionDryRunCount)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description: "CircleCI API token that a `chronosphere_cloud_integration` with `provider_type = \"circleci\"` " +
			"references through `external_connection_slug` to collect CircleCI metrics.",
		Schema: tfschema.CircleciExternalConnection,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
			// The API rejects a connection without a token, so a token produced by
			// another resource in the same apply cannot be validated at plan time.
			if !diff.NewValueKnown("api_token") {
				return nil
			}
			return validateDryRun(ctx, diff, meta)
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CircleciExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var CircleciExternalConnectionDryRunCount atomic.Int64

type circleciExternalConnectionConverter struct{}

func (circleciExternalConnectionConverter) toModel(
	n *intschema.CircleciExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		Circleci: &models.ExternalConnectionCircleCiConfig{
			APIToken: n.ApiToken,
		},
	}, nil
}

func (circleciExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.CircleciExternalConnection, error) {
	c := m.Circleci
	if c == nil {
		return &intschema.CircleciExternalConnection{
			Name: m.Name + " (not a circleci external connection)",
			Slug: m.Slug,
		}, nil
	}
	return &intschema.CircleciExternalConnection{
		Name:     m.Name,
		Slug:     m.Slug,
		ApiToken: c.APIToken,
	}, nil
}

func (circleciExternalConnectionConverter) normalize(
	config, server *intschema.CircleciExternalConnection,
) {
	server.ApiToken = config.ApiToken
}
