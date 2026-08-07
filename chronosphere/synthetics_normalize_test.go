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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// A read rebuilds the auth block out of what the API returned, which is missing
// the secret version; left alone the version comes back as 0 and diffs against
// config on every plan.
func TestSyntheticTestNormalizeSecretVersion(t *testing.T) {
	const (
		versionPath  = "http_test.0.authentication.0.basic_auth.0.password_wo_version"
		usernamePath = "http_test.0.authentication.0.basic_auth.0.username"
	)

	newData := func(t *testing.T) *schema.ResourceData {
		d := schema.TestResourceDataRaw(t, tfschema.SyntheticTest, map[string]any{
			"name":          "checkout availability",
			"collection_id": "web",
			"test_type":     "HTTP",
			"locations":     []any{"GCP_US_OREGON"},
			"http_test": []any{map[string]any{
				"url":    "https://example.com/checkout",
				"method": "GET",
				"authentication": []any{map[string]any{
					"basic_auth": []any{map[string]any{
						"username":            "probe",
						"password_wo_version": 2,
					}},
				}},
			}},
		})
		d.SetId("checkout-availability")
		return d
	}

	newFromServer := func() *intschema.SyntheticTest {
		return &intschema.SyntheticTest{
			Name:         "checkout availability",
			CollectionId: tfid.Slug("web"),
			TestType:     "HTTP",
			Locations:    []string{"GCP_US_OREGON"},
			HttpTest: &intschema.SyntheticTestHttpTest{
				Url:    "https://example.com/checkout",
				Method: "GET",
				Authentication: &intschema.SyntheticTestHttpTestAuthentication{
					BasicAuth: &intschema.SyntheticTestHttpTestAuthenticationBasicAuth{
						Username: "probe",
					},
				},
			},
		}
	}

	d := newData(t)
	require.False(t, newFromServer().ToResourceData(d).HasError())
	assert.Equal(t, "0", d.State().Attributes[versionPath])

	d = newData(t)
	config := &intschema.SyntheticTest{}
	require.NoError(t, config.FromResourceData(d))
	server := newFromServer()
	syntheticTestConverter{}.normalize(config, server)
	require.False(t, server.ToResourceData(d).HasError())

	assert.Equal(t, "2", d.State().Attributes[versionPath])
	assert.Equal(t, "probe", d.State().Attributes[usernamePath])
}
