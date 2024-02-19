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

func SlackAlertNotifierFromModel(
	m *models.Configv1Notifier,
) (*intschema.SlackAlertNotifier, error) {
	return slackAlertNotifierConverter{}.fromModel(m)
}

func resourceSlackAlertNotifier() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Notifier,
		intschema.SlackAlertNotifier,
		*intschema.SlackAlertNotifier,
	](
		"slack_alert_notifier",
		slackAlertNotifierConverter{},
		generatedNotifier{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.SlackAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&SlackAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// SlackAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var SlackAlertNotifierDryRunCount atomic.Int64

type slackAlertNotifierConverter struct{}

func (slackAlertNotifierConverter) toModel(
	n *intschema.SlackAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Name:         n.Name,
		Slug:         n.Slug,
		SkipResolved: !n.SendResolved,
		Slack: &models.NotifierSlackConfig{
			Actions:    slackActionsToModel(n.Action),
			APIURL:     n.ApiUrl,
			CallbackID: n.CallbackId,
			Channel:    n.Channel,
			Color:      n.Color,
			Fallback:   n.Fallback,
			Fields:     slackFieldsToModel(n.Fields),
			Footer:     n.Footer,
			HTTPConfig: notifierHTTPConfig{
				basicAuthUsername:     n.BasicAuthUsername,
				basicAuthPassword:     n.BasicAuthPassword,
				tlsInsecureSkipVerify: n.TlsInsecureSkipVerify,
				bearerToken:           n.BearerToken,
				proxyURL:              n.ProxyUrl,
			}.toModel(),
			IconEmoji:   n.IconEmoji,
			IconURL:     n.IconUrl,
			ImageURL:    n.ImageUrl,
			LinkNames:   n.LinkNames,
			MrkdwnIn:    n.MrkdwnIn,
			Pretext:     n.Pretext,
			ShortFields: n.ShortFields,
			Text:        n.Text,
			ThumbURL:    n.ThumbUrl,
			Title:       n.Title,
			TitleLink:   n.TitleLink,
			Username:    n.Username,
		},
	}, nil
}

func (slackAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.SlackAlertNotifier, error) {
	s := m.Slack
	if s == nil {
		return &intschema.SlackAlertNotifier{
			Name: m.Name + notifierTypeChangedName("slack"),
			Slug: m.Slug,
		}, nil
	}
	n := &intschema.SlackAlertNotifier{
		Name:         m.Name,
		Slug:         m.Slug,
		ApiUrl:       s.APIURL,
		Channel:      s.Channel,
		Action:       slackActionsFromModel(s.Actions),
		CallbackId:   s.CallbackID,
		Color:        s.Color,
		Fallback:     s.Fallback,
		Fields:       slackFieldsFromModel(s.Fields),
		Footer:       s.Footer,
		IconEmoji:    s.IconEmoji,
		IconUrl:      s.IconURL,
		ImageUrl:     s.ImageURL,
		LinkNames:    s.LinkNames,
		MrkdwnIn:     s.MrkdwnIn,
		Pretext:      s.Pretext,
		SendResolved: !m.SkipResolved,
		ShortFields:  s.ShortFields,
		Text:         s.Text,
		ThumbUrl:     s.ThumbURL,
		Title:        s.Title,
		TitleLink:    s.TitleLink,
		Username:     s.Username,
	}
	if s.HTTPConfig != nil {
		c := notifierHTTPConfigFromModel(s.HTTPConfig)
		n.BasicAuthUsername = c.basicAuthUsername
		n.BasicAuthPassword = c.basicAuthPassword
		n.TlsInsecureSkipVerify = c.tlsInsecureSkipVerify
		n.BearerToken = c.bearerToken
		n.ProxyUrl = c.proxyURL
	}
	return n, nil
}

func slackActionsToModel(
	actions []intschema.SlackAlertNotifierAction,
) []*models.NotifierSlackConfigAction {
	var out []*models.NotifierSlackConfigAction
	for _, a := range actions {
		out = append(out, &models.NotifierSlackConfigAction{
			ConfirmField: &models.SlackConfigConfirmationField{
				DismissText: a.ActionConfirmDismissText,
				OkText:      a.ActionConfirmOkText,
				Text:        a.ActionConfirmText,
				Title:       a.ActionConfirmTile,
			},
			Name:  a.Name,
			Style: a.Style,
			Text:  a.Text,
			Type:  a.Type,
			URL:   a.Url,
			Value: a.Value,
		})
	}
	return out
}

func slackActionsFromModel(
	actions []*models.NotifierSlackConfigAction,
) []intschema.SlackAlertNotifierAction {
	var out []intschema.SlackAlertNotifierAction
	for _, a := range actions {
		o := intschema.SlackAlertNotifierAction{
			Name:  a.Name,
			Style: a.Style,
			Text:  a.Text,
			Type:  a.Type,
			Url:   a.URL,
			Value: a.Value,
		}
		if a.ConfirmField != nil {
			o.ActionConfirmDismissText = a.ConfirmField.DismissText
			o.ActionConfirmOkText = a.ConfirmField.OkText
			o.ActionConfirmText = a.ConfirmField.Text
			o.ActionConfirmTile = a.ConfirmField.Title
		}
		out = append(out, o)
	}
	return out
}

func slackFieldsToModel(
	fields []intschema.SlackAlertNotifierFields,
) []*models.SlackConfigField {
	var out []*models.SlackConfigField
	for _, f := range fields {
		out = append(out, &models.SlackConfigField{
			Short: f.Short,
			Title: f.Title,
			Value: f.Value,
		})
	}
	return out
}

func slackFieldsFromModel(
	fields []*models.SlackConfigField,
) []intschema.SlackAlertNotifierFields {
	var out []intschema.SlackAlertNotifierFields
	for _, f := range fields {
		out = append(out, intschema.SlackAlertNotifierFields{
			Short: f.Short,
			Title: f.Title,
			Value: f.Value,
		})
	}
	return out
}
