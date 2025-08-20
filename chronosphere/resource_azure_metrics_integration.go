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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
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
	m := &models.ConfigunstableAzureMetricsIntegration{
		Name:                g.Name,
		Slug:                g.Slug,
		CountMetricsEnabled: g.CountMetricsEnabled,
		UsageMetricsEnabled: g.UsageMetricsEnabled,
		PropagateTags:       g.PropagateTags,
	}

	if g.Principal != nil {
		m.Principal = &models.AzureMetricsIntegrationAzurePrincipal{
			TenantID: g.Principal.TenantId,
			ClientID: g.Principal.ClientId,
		}
	}
	if g.ScrapeConfig != nil {
		m.ScrapeConfig = resourceToModelScrapeConfig(g.ScrapeConfig)
	}

	return m, nil
}

func (azureMetricsIntegrationConverter) fromModel(
	m *models.ConfigunstableAzureMetricsIntegration,
) (*intschema.AzureMetricsIntegration, error) {
	i := &intschema.AzureMetricsIntegration{
		Name:                m.Name,
		Slug:                m.Slug,
		CountMetricsEnabled: m.CountMetricsEnabled,
		UsageMetricsEnabled: m.UsageMetricsEnabled,
		PropagateTags:       m.PropagateTags,
	}

	if m.Principal != nil {
		i.Principal = &intschema.AzureMetricsIntegrationPrincipal{
			TenantId: m.Principal.TenantID,
			ClientId: m.Principal.ClientID,
		}
	}

	if m.ScrapeConfig != nil {
		i.ScrapeConfig = resourceFromModelScrapeConfig(m.ScrapeConfig)
	}

	return i, nil
}

func resourceToModelScrapeConfig(
	scrapeConfig *intschema.AzureMetricsIntegrationScrapeConfig,
) *models.AzureMetricsIntegrationAzureScrapeConfig {
	if scrapeConfig == nil {
		return nil
	}

	return &models.AzureMetricsIntegrationAzureScrapeConfig{
		SubscriptionIds: scrapeConfig.SubscriptionId,
		Locations:       scrapeConfig.Location,
		ResourceTypes:   sliceutil.Map(scrapeConfig.ResourceType, resourceToModelResourceType),
	}
}

func resourceFromModelScrapeConfig(
	modelScrapeConfig *models.AzureMetricsIntegrationAzureScrapeConfig,
) *intschema.AzureMetricsIntegrationScrapeConfig {
	if modelScrapeConfig == nil {
		return nil
	}

	return &intschema.AzureMetricsIntegrationScrapeConfig{
		SubscriptionId: modelScrapeConfig.SubscriptionIds,
		Location:       modelScrapeConfig.Locations,
		ResourceType:   sliceutil.Map(modelScrapeConfig.ResourceTypes, modelToResourceResourceType),
	}
}

func resourceToModelResourceType(
	resourceType intschema.AzureMetricsIntegrationScrapeConfigResourceType,
) *models.AzureMetricsIntegrationAzureResourceType {
	return &models.AzureMetricsIntegrationAzureResourceType{
		Name:        resourceType.Name,
		MetricNames: resourceType.MetricName,
	}
}

func modelToResourceResourceType(
	modelResourceType *models.AzureMetricsIntegrationAzureResourceType,
) intschema.AzureMetricsIntegrationScrapeConfigResourceType {
	return intschema.AzureMetricsIntegrationScrapeConfigResourceType{
		Name:       modelResourceType.Name,
		MetricName: modelResourceType.MetricNames,
	}
}
