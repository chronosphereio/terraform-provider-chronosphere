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

func WebhookExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.WebhookExternalConnection, error) {
	return webhookExternalConnectionConverter{}.fromModel(m)
}

func resourceWebhookExternalConnection() *schema.Resource {
	r := newGenericResource(
		"webhook_external_connection",
		webhookExternalConnectionConverter{},
		generatedExternalConnection{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.WebhookExternalConnection,
		CustomizeDiff: r.ValidateDryRun(&WebhookExternalConnectionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// WebhookExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var WebhookExternalConnectionDryRunCount atomic.Int64

type webhookExternalConnectionConverter struct{}

func (webhookExternalConnectionConverter) toModel(
	n *intschema.WebhookExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		Webhook: &models.Configv1ExternalConnectionWebhookConfig{
			URL: n.Url,
			HTTPConfig: externalConnectionHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				bearerToken:           n.BearerToken,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
			}.toModel(),
		},
	}, nil
}

func (webhookExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.WebhookExternalConnection, error) {
	w := m.Webhook
	if w == nil {
		return &intschema.WebhookExternalConnection{
			Name: m.Name + " (not a webhook external connection)",
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.WebhookExternalConnection{
		Name: m.Name,
		Slug: m.Slug,
		Url:  w.URL,
	}
	if w.HTTPConfig != nil {
		c := externalConnectionHTTPConfigFromModel(w.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.BearerToken = c.bearerToken
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
	}
	return n, nil
}

func (webhookExternalConnectionConverter) normalize(config, server *intschema.WebhookExternalConnection) {
	server.BasicAuthPassword = config.BasicAuthPassword
	server.BearerToken = config.BearerToken
}

// redactedSecretSentinel is the placeholder the Chronosphere API returns for
// secret fields on read. We strip it on fromModel so that export-config
// produces TF without dummy secret values that would fail on create.
const redactedSecretSentinel = "**REDACTED**"

func clearRedactedSecret(s string) string {
	if s == redactedSecretSentinel {
		return ""
	}
	return s
}

// externalConnectionHTTPConfig is a helper for mapping external connection HTTP config fields.
type externalConnectionHTTPConfig struct {
	basicAuthUsername     string
	basicAuthPassword     string
	bearerToken           string
	tlsInsecureSkipVerify bool
}

func (c externalConnectionHTTPConfig) toModel() *models.Configv1ExternalConnectionHTTPConfig {
	m := &models.Configv1ExternalConnectionHTTPConfig{}
	if c.basicAuthUsername != "" || c.basicAuthPassword != "" {
		m.BasicAuth = &models.Configv1ExternalConnectionHTTPConfigBasicAuth{
			Username: c.basicAuthUsername,
			Password: c.basicAuthPassword,
		}
	}
	if c.bearerToken != "" {
		m.BearerToken = c.bearerToken
	}
	if c.tlsInsecureSkipVerify {
		m.TLSConfig = &models.Configv1ExternalConnectionHTTPConfigTLSConfig{
			InsecureSkipVerify: c.tlsInsecureSkipVerify,
		}
	}
	return m
}

func externalConnectionHTTPConfigFromModel(m *models.Configv1ExternalConnectionHTTPConfig) externalConnectionHTTPConfig {
	c := externalConnectionHTTPConfig{
		bearerToken: clearRedactedSecret(m.BearerToken),
	}
	if m.BasicAuth != nil {
		c.basicAuthUsername = m.BasicAuth.Username
		c.basicAuthPassword = clearRedactedSecret(m.BasicAuth.Password)
	}
	if m.TLSConfig != nil {
		c.tlsInsecureSkipVerify = m.TLSConfig.InsecureSkipVerify
	}
	return c
}
