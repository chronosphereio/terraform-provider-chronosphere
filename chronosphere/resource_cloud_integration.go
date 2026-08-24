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
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// CloudIntegrationFromModel maps an API model to the intschema model.
func CloudIntegrationFromModel(m *models.Configv1CloudIntegration) (*intschema.CloudIntegration, error) {
	return (cloudIntegrationConverter{}).fromModel(m)
}

func resourceCloudIntegration() *schema.Resource {
	r := newGenericResource(
		"cloud_integration",
		cloudIntegrationConverter{},
		generatedCloudIntegration{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.CloudIntegration,
		Description:   "Ingests metrics from a cloud provider. The provider-specific configuration is an opaque JSON object whose structure is defined by `provider_type`.",
		CustomizeDiff: r.ValidateDryRun(&CloudIntegrationDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CloudIntegrationDryRunCount tracks how many times dry run is run during validation for testing.
var CloudIntegrationDryRunCount atomic.Int64

type cloudIntegrationConverter struct{}

func (cloudIntegrationConverter) toModel(
	c *intschema.CloudIntegration,
) (*models.Configv1CloudIntegration, error) {
	var providerConfig any
	if c.ProviderConfig != "" {
		if err := json.Unmarshal([]byte(c.ProviderConfig), &providerConfig); err != nil {
			return nil, fmt.Errorf("provider_config is not valid JSON: %w", err)
		}
	}

	return &models.Configv1CloudIntegration{
		Name:                   c.Name,
		Slug:                   c.Slug,
		State:                  enum.CloudIntegrationState.V1(c.State),
		ExternalConnectionSlug: c.ExternalConnectionSlug,
		MetricLabels:           c.MetricLabels,
		ProviderType:           c.ProviderType,
		ProviderConfig:         providerConfig,
	}, nil
}

func (cloudIntegrationConverter) fromModel(
	m *models.Configv1CloudIntegration,
) (*intschema.CloudIntegration, error) {
	providerConfig := ""
	if m.ProviderConfig != nil {
		configJSON, err := json.Marshal(m.ProviderConfig)
		if err != nil {
			return nil, fmt.Errorf("marshaling provider_config: %w", err)
		}
		providerConfig = string(configJSON)
	}

	return &intschema.CloudIntegration{
		Name:                   m.Name,
		Slug:                   m.Slug,
		State:                  enum.CloudIntegrationState.Alias(m.State),
		ExternalConnectionSlug: m.ExternalConnectionSlug,
		MetricLabels:           m.MetricLabels,
		ProviderType:           m.ProviderType,
		ProviderConfig:         providerConfig,
	}, nil
}
