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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
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
