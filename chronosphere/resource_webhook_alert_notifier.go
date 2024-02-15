// Copyright 2023 Chronosphere Inc.
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

func WebhookAlertNotifierFromModel(
	m *models.Configv1Notifier,
) (*intschema.WebhookAlertNotifier, error) {
	return webhookAlertNotifierConverter{}.fromModel(m)
}

func resourceWebhookAlertNotifier() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Notifier,
		intschema.WebhookAlertNotifier,
		*intschema.WebhookAlertNotifier,
	](
		"webhook_alert_notifier",
		webhookAlertNotifierConverter{},
		generatedNotifier{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.WebhookAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&WebhookAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// WebhookAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var WebhookAlertNotifierDryRunCount atomic.Int64

type webhookAlertNotifierConverter struct{}

func (webhookAlertNotifierConverter) toModel(
	n *intschema.WebhookAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Name:         n.Name,
		Slug:         n.Slug,
		SkipResolved: !n.SendResolved,
		Webhook: &models.NotifierWebhookConfig{
			HTTPConfig: notifierHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
				bearerToken:           n.BearerToken,
				proxyURL:              n.ProxyUrl,
			}.toModel(),
			URL: n.Url,
		},
	}, nil
}

func (webhookAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.WebhookAlertNotifier, error) {
	w := m.Webhook
	if w == nil {
		return &intschema.WebhookAlertNotifier{
			Name: m.Name + notifierTypeChangedName("webhook"),
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.WebhookAlertNotifier{
		Name:         m.Name,
		Slug:         m.Slug,
		SendResolved: !m.SkipResolved,
		Url:          w.URL,
	}
	if w.HTTPConfig != nil {
		c := notifierHTTPConfigFromModel(w.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
		n.BearerToken = c.bearerToken
		n.ProxyUrl = c.proxyURL
	}
	return n, nil
}
