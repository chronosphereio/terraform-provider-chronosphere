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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func PagerdutyExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.PagerdutyExternalConnection, error) {
	return pagerdutyExternalConnectionConverter{}.fromModel(m)
}

func resourcePagerdutyExternalConnection() *schema.Resource {
	r := newGenericResource(
		"pagerduty_external_connection",
		pagerdutyExternalConnectionConverter{},
		generatedExternalConnection{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.PagerdutyExternalConnection,
		CustomizeDiff: r.ValidateDryRun(&PagerdutyExternalConnectionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// PagerdutyExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var PagerdutyExternalConnectionDryRunCount atomic.Int64

type pagerdutyExternalConnectionConverter struct{}

func (pagerdutyExternalConnectionConverter) toModel(
	n *intschema.PagerdutyExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		Pagerduty: &models.Configv1ExternalConnectionPagerdutyConfig{
			Events: &models.PagerdutyConfigPagerdutyEventsConfig{
				APIKey:  n.PagerdutyApiKey,
				Version: models.Configv1PagerdutyEventsVersion(n.PagerdutyEventsVersion),
			},
		},
	}, nil
}

func (pagerdutyExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.PagerdutyExternalConnection, error) {
	p := m.Pagerduty
	if p == nil {
		return &intschema.PagerdutyExternalConnection{
			Name: m.Name + " (not a pagerduty external connection)",
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.PagerdutyExternalConnection{
		Name: m.Name,
		Slug: m.Slug,
	}
	if p.Events != nil {
		n.PagerdutyApiKey = clearRedactedSecret(p.Events.APIKey)
		n.PagerdutyEventsVersion = string(p.Events.Version)
	}
	return n, nil
}

func (pagerdutyExternalConnectionConverter) normalize(config, server *intschema.PagerdutyExternalConnection) {
	server.PagerdutyApiKey = config.PagerdutyApiKey
}
