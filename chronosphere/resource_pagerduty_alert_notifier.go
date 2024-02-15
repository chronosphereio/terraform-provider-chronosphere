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

func PagerdutyAlertNotifierFromModel(
	m *models.Configv1Notifier,
) (*intschema.PagerdutyAlertNotifier, error) {
	return pagerdutyAlertNotifierConverter{}.fromModel(m)
}

func resourcePagerdutyAlertNotifier() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Notifier,
		intschema.PagerdutyAlertNotifier,
		*intschema.PagerdutyAlertNotifier,
	](
		"pagerduty_alert_notifier",
		pagerdutyAlertNotifierConverter{},
		generatedNotifier{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.PagerdutyAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&PagerDutyAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// PagerDutyAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var PagerDutyAlertNotifierDryRunCount atomic.Int64

type pagerdutyAlertNotifierConverter struct{}

func (pagerdutyAlertNotifierConverter) toModel(
	n *intschema.PagerdutyAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Name:         n.Name,
		Slug:         n.Slug,
		SkipResolved: !n.SendResolved,
		Pagerduty: &models.NotifierPagerdutyConfig{
			Class:       n.Class,
			Client:      n.Client,
			ClientURL:   n.ClientUrl,
			Component:   n.Component,
			Description: n.Description,
			Details:     n.Details,
			Group:       n.Group,
			HTTPConfig: notifierHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
				bearerToken:           n.BearerToken,
				proxyURL:              n.ProxyUrl,
			}.toModel(),
			Images:     pagerdutyImagesToModel(n.Image),
			Links:      pagerdutyLinksToModel(n.Link),
			RoutingKey: n.RoutingKey,
			ServiceKey: n.ServiceKey,
			Severity:   n.Severity,
			URL:        n.Url,
		},
	}, nil
}

func (pagerdutyAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.PagerdutyAlertNotifier, error) {
	p := m.Pagerduty
	if p == nil {
		return &intschema.PagerdutyAlertNotifier{
			Name: m.Name + notifierTypeChangedName("pagerduty"),
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.PagerdutyAlertNotifier{
		Name:         m.Name,
		Slug:         m.Slug,
		Severity:     p.Severity,
		Url:          p.URL,
		Class:        p.Class,
		Client:       p.Client,
		ClientUrl:    p.ClientURL,
		Component:    p.Component,
		Description:  p.Description,
		Details:      p.Details,
		Group:        p.Group,
		Image:        pagerdutyImagesFromModel(p.Images),
		Link:         pagerdutyLinksFromModel(p.Links),
		RoutingKey:   p.RoutingKey,
		SendResolved: !m.SkipResolved,
		ServiceKey:   p.ServiceKey,
	}
	if p.HTTPConfig != nil {
		c := notifierHTTPConfigFromModel(p.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
		n.BearerToken = c.bearerToken
		n.ProxyUrl = c.proxyURL
	}
	return n, nil
}

func pagerdutyImagesToModel(
	images []intschema.PagerdutyAlertNotifierImage,
) []*models.PagerdutyConfigImage {
	var out []*models.PagerdutyConfigImage
	for _, i := range images {
		out = append(out, &models.PagerdutyConfigImage{
			Alt:  i.Alt,
			Href: i.Href,
			Src:  i.Src,
		})
	}
	return out
}

func pagerdutyImagesFromModel(
	images []*models.PagerdutyConfigImage,
) []intschema.PagerdutyAlertNotifierImage {
	var out []intschema.PagerdutyAlertNotifierImage
	for _, i := range images {
		out = append(out, intschema.PagerdutyAlertNotifierImage{
			Alt:  i.Alt,
			Href: i.Href,
			Src:  i.Src,
		})
	}
	return out
}

func pagerdutyLinksToModel(
	links []intschema.PagerdutyAlertNotifierLink,
) []*models.PagerdutyConfigLink {
	var out []*models.PagerdutyConfigLink
	for _, l := range links {
		out = append(out, &models.PagerdutyConfigLink{
			Href: l.Href,
			Text: l.Text,
		})
	}
	return out
}

func pagerdutyLinksFromModel(
	links []*models.PagerdutyConfigLink,
) []intschema.PagerdutyAlertNotifierLink {
	var out []intschema.PagerdutyAlertNotifierLink
	for _, l := range links {
		out = append(out, intschema.PagerdutyAlertNotifierLink{
			Href: l.Href,
			Text: l.Text,
		})
	}
	return out
}
