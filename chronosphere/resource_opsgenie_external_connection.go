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

func OpsgenieExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.OpsgenieExternalConnection, error) {
	return opsgenieExternalConnectionConverter{}.fromModel(m)
}

func resourceOpsgenieExternalConnection() *schema.Resource {
	r := newGenericResource(
		"opsgenie_external_connection",
		opsgenieExternalConnectionConverter{},
		generatedExternalConnection{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.OpsgenieExternalConnection,
		CustomizeDiff: r.ValidateDryRun(&OpsGenieExternalConnectionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// OpsGenieExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var OpsGenieExternalConnectionDryRunCount atomic.Int64

type opsgenieExternalConnectionConverter struct{}

func (opsgenieExternalConnectionConverter) toModel(
	n *intschema.OpsgenieExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		OpsGenie: &models.Configv1ExternalConnectionOpsGenieConfig{
			APIKey: n.ApiKey,
			APIURL: n.ApiUrl,
			HTTPConfig: externalConnectionHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				bearerToken:           n.BearerToken,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
			}.toModel(),
		},
	}, nil
}

func (opsgenieExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.OpsgenieExternalConnection, error) {
	o := m.OpsGenie
	if o == nil {
		return &intschema.OpsgenieExternalConnection{
			Name: m.Name + " (not an opsgenie external connection)",
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.OpsgenieExternalConnection{
		Name:   m.Name,
		Slug:   m.Slug,
		ApiKey: o.APIKey,
		ApiUrl: o.APIURL,
	}
	if o.HTTPConfig != nil {
		c := externalConnectionHTTPConfigFromModel(o.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.BearerToken = c.bearerToken
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
	}
	return n, nil
}

func (opsgenieExternalConnectionConverter) normalize(config, server *intschema.OpsgenieExternalConnection) {
	server.ApiKey = config.ApiKey
	server.BasicAuthPassword = config.BasicAuthPassword
	server.BearerToken = config.BearerToken
}
