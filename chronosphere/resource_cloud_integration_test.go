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
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/apiclients"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/transport"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func cloudIntegrationConfigWithoutState() map[string]any {
	return map[string]any{
		"name":            "gcp metrics",
		"provider_type":   "gcp",
		"provider_config": `{"service_account":{"client_email":"sa@example.iam.gserviceaccount.com"}}`,
	}
}

func TestCloudIntegrationOmittedStateIsSentAsEnabled(t *testing.T) {
	d := schema.TestResourceDataRaw(t, tfschema.CloudIntegration, cloudIntegrationConfigWithoutState())
	in := &intschema.CloudIntegration{}
	require.NoError(t, in.FromResourceData(d))

	m, err := cloudIntegrationConverter{}.toModel(in)
	require.NoError(t, err)

	assert.Equal(t, models.Configv1CloudIntegrationStateSTATEENABLED, m.State)
}

// The API stores ENABLED for an omitted state, so reading it back must not plan
// a change to null.
func TestCloudIntegrationOmittedStatePlansClean(t *testing.T) {
	raw := cloudIntegrationConfigWithoutState()
	d := schema.TestResourceDataRaw(t, tfschema.CloudIntegration, raw)
	d.SetId("gcp-metrics")
	server := &intschema.CloudIntegration{
		Name:           "gcp metrics",
		Slug:           "gcp-metrics",
		State:          "ENABLED",
		ProviderType:   "gcp",
		ProviderConfig: `{"service_account":{"client_email":"sa@example.iam.gserviceaccount.com"}}`,
	}
	require.False(t, server.ToResourceData(d).HasError())

	r := &schema.Resource{Schema: tfschema.CloudIntegration}
	diff, err := r.Diff(context.Background(), d.State(), terraform.NewResourceConfigRaw(raw), nil)
	require.NoError(t, err)

	assert.True(t, diff == nil || diff.Attributes["state"] == nil, "omitted state should not plan a change")
}

// The Hub exports a form with no state as an API model with an empty state.
func TestCloudIntegrationExportLeavesOutUnsetState(t *testing.T) {
	ci, err := CloudIntegrationFromModel(&models.Configv1CloudIntegration{
		Name:         "gcp metrics",
		Slug:         "gcp-metrics",
		ProviderType: "gcp",
		ProviderConfig: map[string]any{
			"service_account": map[string]any{"client_email": "sa@example.iam.gserviceaccount.com"},
		},
	})
	require.NoError(t, err)
	ci.HCLID = "gcp_metrics"

	var hcl bytes.Buffer
	require.NoError(t, ci.MarshalHCL(&hcl))

	assert.Equal(t, `
resource "chronosphere_cloud_integration" "gcp_metrics" {
  name            = "gcp metrics"
  slug            = "gcp-metrics"
  provider_config = "{\"service_account\":{\"client_email\":\"sa@example.iam.gserviceaccount.com\"}}"
  provider_type   = "gcp"
}
`, hcl.String())
}

func clickhouseCloudIntegrationConfig(connectionSlug string) map[string]any {
	return map[string]any{
		"name":                     "ClickHouse Cloud",
		"provider_type":            "clickhouse_cloud",
		"provider_config":          `{"organization_id":"0c6f2b8e-3d1a-4e7b-9f5c-8a2d4e6b1c73"}`,
		"external_connection_slug": connectionSlug,
	}
}

func TestCloudIntegrationUnknownConnectionSkipsDryRun(t *testing.T) {
	before := CloudIntegrationDryRunCount.Load()

	_, err := resourceCloudIntegration().Diff(context.Background(), nil,
		terraform.NewResourceConfigRaw(clickhouseCloudIntegrationConfig(terraformUnknownValue)), nil)

	require.NoError(t, err)
	assert.Equal(t, before, CloudIntegrationDryRunCount.Load())
}

func TestCloudIntegrationKnownConnectionRunsDryRun(t *testing.T) {
	var bodies []string
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		b, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		bodies = append(bodies, string(b))
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{}`)),
			Request:    req,
		}, nil
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })
	client, err := configv1.NewClient(transport.ComponentTerraformProvider, "test", "test", "")
	require.NoError(t, err)
	before := CloudIntegrationDryRunCount.Load()

	_, err = resourceCloudIntegration().Diff(context.Background(), nil,
		terraform.NewResourceConfigRaw(clickhouseCloudIntegrationConfig("clickhouse-cloud")),
		apiclients.Clients{ConfigV1: client})

	require.NoError(t, err)
	assert.Greater(t, CloudIntegrationDryRunCount.Load(), before)
	require.NotEmpty(t, bodies)
	for _, b := range bodies {
		assert.Contains(t, b, `"external_connection_slug":"clickhouse-cloud"`)
		assert.Contains(t, b, `"dry_run":true`)
	}
}
