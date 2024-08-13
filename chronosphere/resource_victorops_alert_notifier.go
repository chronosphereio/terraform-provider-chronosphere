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

func VictoropsAlertNotifierFromModel(
	m *models.Configv1Notifier,
) (*intschema.VictoropsAlertNotifier, error) {
	return victoropsAlertNotifierConverter{}.fromModel(m)
}

func resourceVictorOpsAlertNotifier() *schema.Resource {
	r := newGenericResource(
		"victorops_alert_notifier",
		victoropsAlertNotifierConverter{},
		generatedNotifier{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.VictoropsAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&VictorOpsAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// VictorOpsAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var VictorOpsAlertNotifierDryRunCount atomic.Int64

type victoropsAlertNotifierConverter struct{}

func (victoropsAlertNotifierConverter) toModel(
	n *intschema.VictoropsAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Name:         n.Name,
		Slug:         n.Slug,
		SkipResolved: !n.SendResolved,
		VictorOps: &models.NotifierVictorOpsConfig{
			APIKey:            n.ApiKey,
			APIURL:            n.ApiUrl,
			CustomFields:      n.CustomFields,
			EntityDisplayName: n.EntityDisplayName,
			HTTPConfig: notifierHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
				bearerToken:           n.BearerToken,
				proxyURL:              n.ProxyUrl,
			}.toModel(),
			MessageType:    n.MessageType,
			MonitoringTool: n.MonitoringTool,
			RoutingKey:     n.RoutingKey,
			StateMessage:   n.StateMessage,
		},
	}, nil
}

func (victoropsAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.VictoropsAlertNotifier, error) {
	v := m.VictorOps
	if v == nil {
		return &intschema.VictoropsAlertNotifier{
			Name: m.Name + notifierTypeChangedName("victorops"),
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.VictoropsAlertNotifier{
		Name:              m.Name,
		Slug:              m.Slug,
		SendResolved:      !m.SkipResolved,
		ApiKey:            v.APIKey,
		ApiUrl:            v.APIURL,
		CustomFields:      v.CustomFields,
		EntityDisplayName: v.EntityDisplayName,
		MessageType:       v.MessageType,
		MonitoringTool:    v.MonitoringTool,
		RoutingKey:        v.RoutingKey,
		StateMessage:      v.StateMessage,
	}
	if v.HTTPConfig != nil {
		c := notifierHTTPConfigFromModel(v.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
		n.BearerToken = c.bearerToken
		n.ProxyUrl = c.proxyURL
	}
	return n, nil
}
