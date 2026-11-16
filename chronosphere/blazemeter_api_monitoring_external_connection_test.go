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

func TestBlazemeterApiMonitoringExternalConnectionModelRoundTrip(t *testing.T) {
	in := &intschema.BlazemeterApiMonitoringExternalConnection{
		Name:        "BlazeMeter API Monitoring",
		Slug:        "blazemeter-api-monitoring",
		AccessToken: "example-token",
	}

	m, err := blazemeterApiMonitoringExternalConnectionConverter{}.toModel(in)
	require.NoError(t, err)
	require.NotNil(t, m.BlazemeterAPIMonitoring)
	assert.Equal(t, "example-token", m.BlazemeterAPIMonitoring.AccessToken)

	out, err := blazemeterApiMonitoringExternalConnectionConverter{}.fromModel(m)
	require.NoError(t, err)
	assert.Equal(t, in, out)
}

func TestBlazemeterApiMonitoringExternalConnectionFromOtherType(t *testing.T) {
	out, err := blazemeterApiMonitoringExternalConnectionConverter{}.fromModel(&models.Configv1ExternalConnection{
		Name:       "Cloudflare",
		Slug:       "cloudflare",
		Cloudflare: &models.ExternalConnectionCloudflareConfig{APIToken: "**REDACTED**"},
	})

	require.NoError(t, err)
	assert.Equal(t, "Cloudflare (not a blazemeter_api_monitoring external connection)", out.Name)
	assert.Empty(t, out.AccessToken)
}

func TestBlazemeterApiMonitoringExternalConnectionNormalizeKeepsConfiguredToken(t *testing.T) {
	config := &intschema.BlazemeterApiMonitoringExternalConnection{AccessToken: "configured-token"}
	server := &intschema.BlazemeterApiMonitoringExternalConnection{AccessToken: "**REDACTED**"}

	blazemeterApiMonitoringExternalConnectionConverter{}.normalize(config, server)

	assert.Equal(t, "configured-token", server.AccessToken)
}

func TestBlazemeterApiMonitoringExternalConnectionDryRunSendsRotatedToken(t *testing.T) {
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
	data := schema.TestResourceDataRaw(t, tfschema.BlazemeterApiMonitoringExternalConnection, map[string]any{
		"name":         "BlazeMeter API Monitoring",
		"slug":         "blazemeter-api-monitoring",
		"access_token": "original-token",
	})
	data.SetId("blazemeter-api-monitoring")

	_, err = resourceBlazemeterApiMonitoringExternalConnection().Diff(
		context.Background(),
		data.State(),
		terraform.NewResourceConfigRaw(map[string]any{
			"name":         "BlazeMeter API Monitoring",
			"slug":         "blazemeter-api-monitoring",
			"access_token": "rotated-token",
		}),
		apiclients.Clients{ConfigV1: client},
	)

	require.NoError(t, err)
	assert.Equal(t, []string{"PUT /api/v1/config/external-connections/blazemeter-api-monitoring"}, requests)
	require.Len(t, bodies, 1)
	assert.Contains(t, bodies[0], `"blazemeter_api_monitoring":{"access_token":"rotated-token"}`)
}

func TestBlazemeterApiMonitoringExternalConnectionUnknownTokenSkipsDryRun(t *testing.T) {
	config := terraform.NewResourceConfigRaw(map[string]any{
		"name":         "BlazeMeter API Monitoring generated",
		"access_token": terraformUnknownValue,
	})
	before := BlazemeterApiMonitoringExternalConnectionDryRunCount.Load()

	_, err := resourceBlazemeterApiMonitoringExternalConnection().Diff(context.Background(), nil, config, nil)

	require.NoError(t, err)
	assert.Equal(t, before, BlazemeterApiMonitoringExternalConnectionDryRunCount.Load())
}
