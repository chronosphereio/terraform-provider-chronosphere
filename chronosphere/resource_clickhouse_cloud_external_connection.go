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

func ClickhouseCloudExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.ClickhouseCloudExternalConnection, error) {
	return clickhouseCloudExternalConnectionConverter{}.fromModel(m)
}

func resourceClickhouseCloudExternalConnection() *schema.Resource {
	r := newGenericResource(
		"clickhouse_cloud_external_connection",
		clickhouseCloudExternalConnectionConverter{},
		generatedExternalConnection{},
	)
	validateDryRun := r.ValidateDryRun(&ClickhouseCloudExternalConnectionDryRunCount)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description: "ClickHouse Cloud API key that a `chronosphere_cloud_integration` with " +
			"`provider_type = \"clickhouse_cloud\"` references through `external_connection_slug` to collect " +
			"ClickHouse Cloud metrics.",
		Schema: tfschema.ClickhouseCloudExternalConnection,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
			// The API rejects a connection without both halves of the key, so a key
			// produced by another resource in the same apply cannot be validated at
			// plan time.
			if !diff.NewValueKnown("key_id") || !diff.NewValueKnown("key_secret") {
				return nil
			}
			return validateDryRun(ctx, diff, meta)
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// ClickhouseCloudExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var ClickhouseCloudExternalConnectionDryRunCount atomic.Int64

type clickhouseCloudExternalConnectionConverter struct{}

func (clickhouseCloudExternalConnectionConverter) toModel(
	n *intschema.ClickhouseCloudExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		ClickhouseCloud: &models.ExternalConnectionClickHouseCloudConfig{
			KeyID:     n.KeyId,
			KeySecret: n.KeySecret,
		},
	}, nil
}

func (clickhouseCloudExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.ClickhouseCloudExternalConnection, error) {
	c := m.ClickhouseCloud
	if c == nil {
		return &intschema.ClickhouseCloudExternalConnection{
			Name: m.Name + " (not a clickhouse_cloud external connection)",
			Slug: m.Slug,
		}, nil
	}
	return &intschema.ClickhouseCloudExternalConnection{
		Name:      m.Name,
		Slug:      m.Slug,
		KeyId:     c.KeyID,
		KeySecret: c.KeySecret,
	}, nil
}

func (clickhouseCloudExternalConnectionConverter) normalize(
	config, server *intschema.ClickhouseCloudExternalConnection,
) {
	server.KeySecret = config.KeySecret
}
