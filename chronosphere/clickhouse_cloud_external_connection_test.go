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

func TestClickhouseCloudExternalConnectionModelRoundTrip(t *testing.T) {
	in := &intschema.ClickhouseCloudExternalConnection{
		Name:      "ClickHouse Cloud",
		Slug:      "clickhouse-cloud",
		KeyId:     "example-key-id",
		KeySecret: "example-key-secret",
	}

	m, err := clickhouseCloudExternalConnectionConverter{}.toModel(in)
	require.NoError(t, err)
	require.NotNil(t, m.ClickhouseCloud)
	assert.Equal(t, "example-key-id", m.ClickhouseCloud.KeyID)
	assert.Equal(t, "example-key-secret", m.ClickhouseCloud.KeySecret)

	out, err := clickhouseCloudExternalConnectionConverter{}.fromModel(m)
	require.NoError(t, err)
	assert.Equal(t, in, out)
}

func TestClickhouseCloudExternalConnectionFromOtherType(t *testing.T) {
	out, err := clickhouseCloudExternalConnectionConverter{}.fromModel(&models.Configv1ExternalConnection{
		Name:       "Cloudflare",
		Slug:       "cloudflare",
		Cloudflare: &models.ExternalConnectionCloudflareConfig{APIToken: "**REDACTED**"},
	})

	require.NoError(t, err)
	assert.Equal(t, "Cloudflare (not a clickhouse_cloud external connection)", out.Name)
	assert.Empty(t, out.KeyId)
	assert.Empty(t, out.KeySecret)
}

// The API returns the key ID as stored, so only the secret comes from config.
func TestClickhouseCloudExternalConnectionNormalizeKeepsConfiguredSecret(t *testing.T) {
	config := &intschema.ClickhouseCloudExternalConnection{KeyId: "configured-key-id", KeySecret: "configured-secret"}
	server := &intschema.ClickhouseCloudExternalConnection{KeyId: "server-key-id", KeySecret: "**REDACTED**"}

	clickhouseCloudExternalConnectionConverter{}.normalize(config, server)

	assert.Equal(t, "server-key-id", server.KeyId)
	assert.Equal(t, "configured-secret", server.KeySecret)
}

func TestClickhouseCloudExternalConnectionDryRunSendsRotatedKey(t *testing.T) {
	var requests, bodies []string
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		b, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		requests = append(requests, req.Method+" "+req.URL.Path)
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
	data := schema.TestResourceDataRaw(t, tfschema.ClickhouseCloudExternalConnection, map[string]any{
		"name":       "ClickHouse Cloud",
		"slug":       "clickhouse-cloud",
		"key_id":     "original-key-id",
		"key_secret": "original-secret",
	})
	data.SetId("clickhouse-cloud")

	_, err = resourceClickhouseCloudExternalConnection().Diff(
		context.Background(),
		data.State(),
		terraform.NewResourceConfigRaw(map[string]any{
			"name":       "ClickHouse Cloud",
			"slug":       "clickhouse-cloud",
			"key_id":     "rotated-key-id",
			"key_secret": "rotated-secret",
		}),
		apiclients.Clients{ConfigV1: client},
	)

	require.NoError(t, err)
	assert.Equal(t, []string{"PUT /api/v1/config/external-connections/clickhouse-cloud"}, requests)
	require.Len(t, bodies, 1)
	assert.Contains(t, bodies[0], `"clickhouse_cloud":{"key_id":"rotated-key-id","key_secret":"rotated-secret"}`)
}

func TestClickhouseCloudExternalConnectionUnknownKeySkipsDryRun(t *testing.T) {
	tests := []struct {
		name   string
		config map[string]any
	}{
		{
			name: "unknown key_id",
			config: map[string]any{
				"name":       "ClickHouse Cloud generated",
				"key_id":     terraformUnknownValue,
				"key_secret": "known-secret",
			},
		},
		{
			name: "unknown key_secret",
			config: map[string]any{
				"name":       "ClickHouse Cloud generated",
				"key_id":     "known-key-id",
				"key_secret": terraformUnknownValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := ClickhouseCloudExternalConnectionDryRunCount.Load()

			_, err := resourceClickhouseCloudExternalConnection().Diff(
				context.Background(), nil, terraform.NewResourceConfigRaw(tt.config), nil)

			require.NoError(t, err)
			assert.Equal(t, before, ClickhouseCloudExternalConnectionDryRunCount.Load())
		})
	}
}
