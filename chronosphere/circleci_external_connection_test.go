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

func TestCircleciExternalConnectionModelRoundTrip(t *testing.T) {
	in := &intschema.CircleciExternalConnection{
		Name:     "CircleCI",
		Slug:     "circleci",
		ApiToken: "example-token",
	}

	m, err := circleciExternalConnectionConverter{}.toModel(in)
	require.NoError(t, err)
	require.NotNil(t, m.Circleci)
	assert.Equal(t, "example-token", m.Circleci.APIToken)

	out, err := circleciExternalConnectionConverter{}.fromModel(m)
	require.NoError(t, err)
	assert.Equal(t, in, out)
}

func TestCircleciExternalConnectionFromOtherType(t *testing.T) {
	out, err := circleciExternalConnectionConverter{}.fromModel(&models.Configv1ExternalConnection{
		Name:       "Cloudflare",
		Slug:       "cloudflare",
		Cloudflare: &models.ExternalConnectionCloudflareConfig{APIToken: "**REDACTED**"},
	})

	require.NoError(t, err)
	assert.Equal(t, "Cloudflare (not a circleci external connection)", out.Name)
	assert.Empty(t, out.ApiToken)
}

func TestCircleciExternalConnectionNormalizeKeepsConfiguredToken(t *testing.T) {
	config := &intschema.CircleciExternalConnection{ApiToken: "configured-token"}
	server := &intschema.CircleciExternalConnection{ApiToken: "**REDACTED**"}

	circleciExternalConnectionConverter{}.normalize(config, server)

	assert.Equal(t, "configured-token", server.ApiToken)
}

func TestCircleciExternalConnectionDryRunSendsRotatedToken(t *testing.T) {
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
	data := schema.TestResourceDataRaw(t, tfschema.CircleciExternalConnection, map[string]any{
		"name":      "CircleCI",
		"slug":      "circleci",
		"api_token": "original-token",
	})
	data.SetId("circleci")

	_, err = resourceCircleciExternalConnection().Diff(
		context.Background(),
		data.State(),
		terraform.NewResourceConfigRaw(map[string]any{
			"name":      "CircleCI",
			"slug":      "circleci",
			"api_token": "rotated-token",
		}),
		apiclients.Clients{ConfigV1: client},
	)

	require.NoError(t, err)
	assert.Equal(t, []string{"PUT /api/v1/config/external-connections/circleci"}, requests)
	require.Len(t, bodies, 1)
	assert.Contains(t, bodies[0], `"circleci":{"api_token":"rotated-token"}`)
}

func TestCircleciExternalConnectionUnknownTokenSkipsDryRun(t *testing.T) {
	config := terraform.NewResourceConfigRaw(map[string]any{
		"name":      "CircleCI generated",
		"api_token": terraformUnknownValue,
	})
	before := CircleciExternalConnectionDryRunCount.Load()

	_, err := resourceCircleciExternalConnection().Diff(context.Background(), nil, config, nil)

	require.NoError(t, err)
	assert.Equal(t, before, CircleciExternalConnectionDryRunCount.Load())
}
