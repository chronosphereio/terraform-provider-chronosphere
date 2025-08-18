// Copyright 2024 Chronosphere Inc.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// AzureMetricsIntegrationFromModel maps an API model to the intschema model.
func AzureMetricsIntegrationFromModel(m *models.ConfigunstableAzureMetricsIntegration) (*intschema.AzureMetricsIntegration, error) {
	return (azureMetricsIntegrationConverter{}).fromModel(m)
}

func resourceAzureMetricsIntegration() *schema.Resource {
	r := newGenericResource(
		"azure_metrics_integration",
		azureMetricsIntegrationConverter{},
		generatedUnstableAzureMetricsIntegration{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.AzureMetricsIntegration,
		CustomizeDiff: r.ValidateDryRun(&AzureMetricsIntegrationDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// AzureMetricsIntegrationDryRunCount tracks how many times dry run is run during validation for testing.
var AzureMetricsIntegrationDryRunCount atomic.Int64

type azureMetricsIntegrationConverter struct{}

func (azureMetricsIntegrationConverter) toModel(
	g *intschema.AzureMetricsIntegration,
) (*models.ConfigunstableAzureMetricsIntegration, error) {
	return &models.ConfigunstableAzureMetricsIntegration{
		Name:                g.Name,
		Slug:                g.Slug,
		CountMetricsEnabled: g.CountMetricsEnabled,
		UsageMetricsEnabled: g.UsageMetricsEnabled,
		PropagateTags:       g.PropagateTags,
		Principal: &models.AzureMetricsIntegrationAzurePrincipal{
			TenantID: g.Principal.TenantId,
			ClientID: g.Principal.ClientId,
		},
		ScrapeConfig: resourceToModelScrapeConfig(g.ScrapeConfig),
	}, nil
}

func (azureMetricsIntegrationConverter) fromModel(
	m *models.ConfigunstableAzureMetricsIntegration,
) (*intschema.AzureMetricsIntegration, error) {
	return &intschema.AzureMetricsIntegration{
		Name:                m.Name,
		Slug:                m.Slug,
		CountMetricsEnabled: m.CountMetricsEnabled,
		UsageMetricsEnabled: m.UsageMetricsEnabled,
		PropagateTags:       m.PropagateTags,
		Principal: &intschema.AzureMetricsIntegrationPrincipal{
			TenantId: m.Principal.TenantID,
			ClientId: m.Principal.ClientID,
		},
		ScrapeConfig: resourceFromModelScrapeConfig(m.ScrapeConfig),
	}, nil
}

func resourceToModelScrapeConfig(
	scrapeConfig *intschema.AzureMetricsIntegrationScrapeConfig,
) *models.AzureMetricsIntegrationAzureScrapeConfig {
	if scrapeConfig == nil {
		return nil
	}

	modelResourceTypes := make([]*models.AzureMetricsIntegrationAzureResourceType, len(scrapeConfig.ResourceTypes))
	for i, rt := range scrapeConfig.ResourceTypes {
		modelResourceTypes[i] = &models.AzureMetricsIntegrationAzureResourceType{
			Name:        rt.Name,
			MetricNames: rt.MetricNames,
		}
	}

	return &models.AzureMetricsIntegrationAzureScrapeConfig{
		SubscriptionIds: scrapeConfig.SubscriptionIds,
		Locations:       scrapeConfig.Locations,
		ResourceTypes:   modelResourceTypes,
	}
}

func resourceFromModelScrapeConfig(
	modelScrapeConfig *models.AzureMetricsIntegrationAzureScrapeConfig,
) *intschema.AzureMetricsIntegrationScrapeConfig {
	if modelScrapeConfig == nil {
		return nil
	}

	resourceTypes := make([]intschema.AzureMetricsIntegrationScrapeConfigResourceTypes, len(modelScrapeConfig.ResourceTypes))
	for i, rt := range modelScrapeConfig.ResourceTypes {
		resourceTypes[i] = intschema.AzureMetricsIntegrationScrapeConfigResourceTypes{
			Name:        rt.Name,
			MetricNames: rt.MetricNames,
		}
	}

	return &intschema.AzureMetricsIntegrationScrapeConfig{
		SubscriptionIds: modelScrapeConfig.SubscriptionIds,
		Locations:       modelScrapeConfig.Locations,
		ResourceTypes:   resourceTypes,
	}
}
