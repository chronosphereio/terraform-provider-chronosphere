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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/transport"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func TestCloudflareExternalConnectionNormalizeKeepsConfiguredToken(t *testing.T) {
	config := &intschema.CloudflareExternalConnection{ApiToken: "cfat_configured"}
	server := &intschema.CloudflareExternalConnection{ApiToken: "**REDACTED**"}

	cloudflareExternalConnectionConverter{}.normalize(config, server)

	assert.Equal(t, "cfat_configured", server.ApiToken)
}

func TestCloudflareExternalConnectionDryRunSendsRotatedToken(t *testing.T) {
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
	data := schema.TestResourceDataRaw(t, tfschema.CloudflareExternalConnection, map[string]any{
		"name":      "Cloudflare",
		"slug":      "cloudflare",
		"api_token": "cfat_original",
	})
	data.SetId("cloudflare")

	_, err = resourceCloudflareExternalConnection().Diff(
		context.Background(),
		data.State(),
		terraform.NewResourceConfigRaw(map[string]any{
			"name":      "Cloudflare",
			"slug":      "cloudflare",
			"api_token": "cfat_rotated",
		}),
		apiclients.Clients{ConfigV1: client},
	)

	require.NoError(t, err)
	assert.Equal(t, []string{"PUT /api/v1/config/external-connections/cloudflare"}, requests)
	require.Len(t, bodies, 1)
	assert.Contains(t, bodies[0], `"api_token":"cfat_rotated"`)
}

func TestCloudflareExternalConnectionUnknownTokenSkipsDryRun(t *testing.T) {
	config := terraform.NewResourceConfigRaw(map[string]any{
		"name":      "Cloudflare generated",
		"api_token": terraformUnknownValue,
	})
	before := CloudflareExternalConnectionDryRunCount.Load()

	_, err := resourceCloudflareExternalConnection().Diff(context.Background(), nil, config, nil)

	require.NoError(t, err)
	assert.Equal(t, before, CloudflareExternalConnectionDryRunCount.Load())
}

// Terraform's legacy test API recognizes unknown values through this unexported sentinel.
const terraformUnknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
