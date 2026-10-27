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

func TestCloudflareExternalConnectionConverter(t *testing.T) {
	model, err := (cloudflareExternalConnectionConverter{}).toModel(&intschema.CloudflareExternalConnection{
		Name:       "Cloudflare production",
		Slug:       "cloudflare-production",
		ApiTokenWo: "cfat_secret",
	})
	require.NoError(t, err)
	assert.Equal(t, "cfat_secret", model.Cloudflare.APIToken)

	connection, err := (cloudflareExternalConnectionConverter{}).fromModel(&models.Configv1ExternalConnection{
		Name:       "Cloudflare production",
		Slug:       "cloudflare-production",
		Cloudflare: &models.ExternalConnectionCloudflareConfig{APIToken: "**REDACTED**"},
	})
	require.NoError(t, err)
	assert.Equal(t, &intschema.CloudflareExternalConnection{
		Name: "Cloudflare production",
		Slug: "cloudflare-production",
	}, connection)
}

func TestCloudflareExternalConnectionNormalizePreservesTokenVersion(t *testing.T) {
	config := &intschema.CloudflareExternalConnection{ApiTokenWoVersion: 3}
	server := &intschema.CloudflareExternalConnection{}

	cloudflareExternalConnectionConverter{}.normalize(config, server)

	assert.Equal(t, int64(3), server.ApiTokenWoVersion)
}

func TestCloudflareExternalConnectionEmptyTokenFailsValidation(t *testing.T) {
	_, errs := tfschema.CloudflareExternalConnection["api_token_wo"].ValidateFunc("", "api_token_wo")

	require.Len(t, errs, 1)
	assert.ErrorContains(t, errs[0], `expected "api_token_wo" to not be an empty string`)
}

func TestCloudflareExternalConnectionUnknownTokenSkipsDryRun(t *testing.T) {
	config := terraform.NewResourceConfigRaw(map[string]any{
		"name":         "Cloudflare generated",
		"api_token_wo": terraformUnknownValue,
	})
	before := CloudflareExternalConnectionDryRunCount.Load()

	_, err := resourceCloudflareExternalConnection().Diff(context.Background(), nil, config, nil)

	require.NoError(t, err)
	assert.Equal(t, before, CloudflareExternalConnectionDryRunCount.Load())
}

func TestCloudflareExternalConnectionKnownInvalidTokenFailsDryRun(t *testing.T) {
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "/api/v1/config/external-connections", req.URL.Path)
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Status:     "400 Bad Request",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body: io.NopCloser(strings.NewReader(
				`{"code":400000,"message":"cloudflare token is not valid"}`,
			)),
			Request: req,
		}, nil
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	client, err := configv1.NewClient(transport.ComponentTerraformProvider, "test", "test", "")
	require.NoError(t, err)

	_, err = resourceCloudflareExternalConnection().Diff(
		context.Background(),
		nil,
		terraform.NewResourceConfigRaw(map[string]any{
			"name":         "Cloudflare",
			"api_token_wo": "not-valid",
		}),
		apiclients.Clients{ConfigV1: client},
	)

	require.ErrorContains(t, err, "dry run validation failed")
	assert.ErrorContains(t, err, "cloudflare token is not valid")
}

func TestCloudflareExternalConnectionTokenVersionPlansUpdate(t *testing.T) {
	data := schema.TestResourceDataRaw(t, tfschema.CloudflareExternalConnection, map[string]any{
		"name":                 "Cloudflare",
		"slug":                 "cloudflare",
		"api_token_wo_version": 1,
	})
	data.SetId("cloudflare")
	state := data.State()
	assert.NotContains(t, state.Attributes, "api_token_wo")

	resource := &schema.Resource{Schema: tfschema.CloudflareExternalConnection}
	diff, err := resource.Diff(
		context.Background(),
		state,
		terraform.NewResourceConfigRaw(map[string]any{
			"name":                 "Cloudflare",
			"slug":                 "cloudflare",
			"api_token_wo":         "cfat_rotated",
			"api_token_wo_version": 2,
		}),
		nil,
	)
	require.NoError(t, err)

	versionDiff := diff.Attributes["api_token_wo_version"]
	require.NotNil(t, versionDiff)
	assert.Equal(t, "1", versionDiff.Old)
	assert.Equal(t, "2", versionDiff.New)
}

// Terraform's legacy test API recognizes unknown values through this unexported sentinel.
const terraformUnknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
