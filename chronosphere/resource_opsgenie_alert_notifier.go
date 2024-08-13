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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func OpsgenieAlertNotifierFromModel(
	m *models.Configv1Notifier,
) (*intschema.OpsgenieAlertNotifier, error) {
	return opsgenieAlertNotifierConverter{}.fromModel(m)
}

func resourceOpsGenieAlertNotifier() *schema.Resource {
	r := newGenericResource(
		"opsgenie_alert_notifier",
		opsgenieAlertNotifierConverter{},
		generatedNotifier{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.OpsgenieAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&OpsGenieAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// OpsGenieAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var OpsGenieAlertNotifierDryRunCount atomic.Int64

type opsgenieAlertNotifierConverter struct{}

func (opsgenieAlertNotifierConverter) toModel(
	n *intschema.OpsgenieAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Slug:         n.Slug,
		Name:         n.Name,
		SkipResolved: !n.SendResolved,
		OpsGenie: &models.NotifierOpsGenieConfig{
			APIKey:      n.ApiKey,
			APIURL:      n.ApiUrl,
			Description: n.Description,
			Details:     n.Details,
			HTTPConfig: notifierHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
				bearerToken:           n.BearerToken,
				proxyURL:              n.ProxyUrl,
			}.toModel(),
			Message:    n.Message,
			Note:       n.Note,
			Priority:   n.Priority,
			Responders: sliceutil.Map(n.Responder, opsgenieResponderToModel),
			Source:     n.Source,
			Tags:       strings.Join(n.Tags, ","),
		},
	}, nil
}

func (opsgenieAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.OpsgenieAlertNotifier, error) {
	o := m.OpsGenie
	if o == nil {
		return &intschema.OpsgenieAlertNotifier{
			Name: m.Name + notifierTypeChangedName("opsgenie"),
			Slug: m.Slug,
		}, nil
	}

	n := &intschema.OpsgenieAlertNotifier{
		Name:         m.Name,
		Slug:         m.Slug,
		ApiKey:       o.APIKey,
		ApiUrl:       o.APIURL,
		Description:  o.Description,
		Details:      o.Details,
		Message:      o.Message,
		Note:         o.Note,
		Priority:     o.Priority,
		Responder:    sliceutil.Map(o.Responders, opsgenieResponderFromModel),
		Source:       o.Source,
		Tags:         opsgenieTagsFromModel(o.Tags),
		SendResolved: !m.SkipResolved,
	}
	if o.HTTPConfig != nil {
		c := notifierHTTPConfigFromModel(o.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
		n.BearerToken = c.bearerToken
		n.ProxyUrl = c.proxyURL
	}
	return n, nil
}

func opsgenieResponderToModel(
	r intschema.OpsgenieAlertNotifierResponder,
) *models.OpsGenieConfigResponder {
	return &models.OpsGenieConfigResponder{
		ID:            r.Id,
		Name:          r.Name,
		ResponderType: enum.OpsgenieResponderType.V1(r.Type),
		Username:      r.Username,
	}
}

func opsgenieResponderFromModel(
	r *models.OpsGenieConfigResponder,
) intschema.OpsgenieAlertNotifierResponder {
	return intschema.OpsgenieAlertNotifierResponder{
		Id:       r.ID,
		Name:     r.Name,
		Type:     string(r.ResponderType),
		Username: r.Username,
	}
}

func opsgenieTagsFromModel(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
