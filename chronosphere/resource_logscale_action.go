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
	"errors"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"
)

// LogscaleActionFromModel maps an API model to the intschema model.
func LogscaleActionFromModel(m *models.ConfigunstableLogScaleAction) (*intschema.LogscaleAction, error) {
	return (logscaleActionConverter{}).fromModel(m)
}

func resourceLogscaleAction() *schema.Resource {
	r := newGenericResource[
		*models.ConfigunstableLogScaleAction,
		intschema.LogscaleAction,
		*intschema.LogscaleAction,
	](
		"logscale_action",
		logscaleActionConverter{},
		generatedUnstableLogScaleAction{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.LogscaleAction,
		CustomizeDiff: r.ValidateDryRun(&LogscaleActionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// LogscaleActionDryRunCount tracks how many times dry run is run during validation for testing.
var LogscaleActionDryRunCount atomic.Int64

type logscaleActionConverter struct{}

func (logscaleActionConverter) toModel(
	c *intschema.LogscaleAction,
) (*models.ConfigunstableLogScaleAction, error) {
	if c == nil {
		return nil, nil
	}
	actionType, err := actionTypeFromIntSchema(c)
	if err != nil {
		return nil, err
	}
	return &models.ConfigunstableLogScaleAction{
		Name:                   c.Name,
		Slug:                   c.Slug,
		Repository:             c.Repository,
		ActionType:             actionType,
		EmailAction:            emailActionToModel(c.EmailAction),
		HumioAction:            humioActionToModel(c.HumioAction),
		OpsGenieAction:         opsGenieActionToModel(c.OpsGenieAction),
		PagerDutyAction:        pagerDutyActionToModel(c.PagerDutyAction),
		SlackAction:            slackActionToModel(c.SlackAction),
		SlackPostMessageAction: slackPostMessageActionToModel(c.SlackPostMessageAction),
		UploadFileAction:       uploadFileActionToModel(c.UploadFileAction),
		VictorOpsAction:        victorOpsActionToModel(c.VictorOpsAction),
		WebhookAction:          webhookActionToModel(c.WebhookAction),
	}, nil
}

func (logscaleActionConverter) fromModel(
	m *models.ConfigunstableLogScaleAction,
) (*intschema.LogscaleAction, error) {
	return &intschema.LogscaleAction{
		Name:                   m.Name,
		Slug:                   m.Slug,
		Repository:             m.Repository,
		EmailAction:            emailActionFromModel(m.EmailAction),
		HumioAction:            humioActionFromModel(m.HumioAction),
		OpsGenieAction:         opsGenieActionFromModel(m.OpsGenieAction),
		PagerDutyAction:        pagerDutyActionFromModel(m.PagerDutyAction),
		SlackAction:            slackActionFromModel(m.SlackAction),
		SlackPostMessageAction: slackPostMessageActionFromModel(m.SlackPostMessageAction),
		UploadFileAction:       uploadFileActionFromModel(m.UploadFileAction),
		VictorOpsAction:        victorOpsActionFromModel(m.VictorOpsAction),
		WebhookAction:          webhookActionFromModel(m.WebhookAction),
	}, nil
}

func actionTypeFromIntSchema(c *intschema.LogscaleAction) (models.LogScaleActionActionType, error) {
	switch {
	case c.EmailAction != nil:
		return models.LogScaleActionActionTypeEMAIL, nil
	case c.HumioAction != nil:
		return models.LogScaleActionActionTypeHUMIOREPO, nil
	case c.OpsGenieAction != nil:
		return models.LogScaleActionActionTypeOPSGENIE, nil
	case c.PagerDutyAction != nil:
		return models.LogScaleActionActionTypePAGERDUTY, nil
	case c.SlackAction != nil:
		return models.LogScaleActionActionTypeSLACK, nil
	case c.SlackPostMessageAction != nil:
		return models.LogScaleActionActionTypeSLACKPOSTMESSAGE, nil
	case c.VictorOpsAction != nil:
		return models.LogScaleActionActionTypeVICTOROPS, nil
	case c.UploadFileAction != nil:
		return models.LogScaleActionActionTypeUPLOADFILE, nil
	case c.WebhookAction != nil:
		return models.LogScaleActionActionTypeWEBHOOK, nil
	}

	return "", errors.New("at least one type of action needs to be set")
}

func emailActionToModel(c *intschema.LogscaleActionEmailAction) *models.LogScaleActionEmailAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionEmailAction{
		AttachCsv:       c.AttachCsv,
		BodyTemplate:    c.BodyTemplate,
		Recipients:      c.Recipients,
		SubjectTemplate: c.SubjectTemplate,
		UseProxy:        c.UseProxy,
	}
}

func emailActionFromModel(c *models.LogScaleActionEmailAction) *intschema.LogscaleActionEmailAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionEmailAction{
		AttachCsv:       c.AttachCsv,
		BodyTemplate:    c.BodyTemplate,
		Recipients:      c.Recipients,
		SubjectTemplate: c.SubjectTemplate,
		UseProxy:        c.UseProxy,
	}
}

func humioActionToModel(c *intschema.LogscaleActionHumioAction) *models.LogScaleActionHumioRepoAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionHumioRepoAction{
		IngestToken: c.IngestToken,
	}
}

func humioActionFromModel(c *models.LogScaleActionHumioRepoAction) *intschema.LogscaleActionHumioAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionHumioAction{
		IngestToken: c.IngestToken,
	}
}

func opsGenieActionToModel(c *intschema.LogscaleActionOpsGenieAction) *models.LogScaleActionOpsGenieAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionOpsGenieAction{
		APIURL:      c.ApiUrl,
		OpsGenieKey: c.OpsGenieKey,
		UseProxy:    c.UseProxy,
	}
}

func opsGenieActionFromModel(c *models.LogScaleActionOpsGenieAction) *intschema.LogscaleActionOpsGenieAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionOpsGenieAction{
		ApiUrl:      c.APIURL,
		OpsGenieKey: c.OpsGenieKey,
		UseProxy:    c.UseProxy,
	}
}

func pagerDutyActionToModel(c *intschema.LogscaleActionPagerDutyAction) *models.LogScaleActionPagerDutyAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionPagerDutyAction{
		RoutingKey: c.RoutingKey,
		Severity:   models.LogScaleActionPagerDutyActionSeverity(c.Severity),
		UseProxy:   c.UseProxy,
	}
}

func pagerDutyActionFromModel(c *models.LogScaleActionPagerDutyAction) *intschema.LogscaleActionPagerDutyAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionPagerDutyAction{
		RoutingKey: c.RoutingKey,
		Severity:   string(c.Severity),
		UseProxy:   c.UseProxy,
	}
}

func slackActionToModel(c *intschema.LogscaleActionSlackAction) *models.LogScaleActionSlackAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionSlackAction{
		Fields:   c.Fields,
		URL:      c.Url,
		UseProxy: c.UseProxy,
	}
}

func slackActionFromModel(c *models.LogScaleActionSlackAction) *intschema.LogscaleActionSlackAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionSlackAction{
		Fields:   c.Fields,
		Url:      c.URL,
		UseProxy: c.UseProxy,
	}
}

func slackPostMessageActionToModel(c *intschema.LogscaleActionSlackPostMessageAction) *models.LogScaleActionSlackPostMessageAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionSlackPostMessageAction{
		APIToken: c.ApiToken,
		Channels: c.Channels,
		Fields:   c.Fields,
		UseProxy: c.UseProxy,
	}
}

func slackPostMessageActionFromModel(c *models.LogScaleActionSlackPostMessageAction) *intschema.LogscaleActionSlackPostMessageAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionSlackPostMessageAction{
		ApiToken: c.APIToken,
		Channels: c.Channels,
		Fields:   c.Fields,
		UseProxy: c.UseProxy,
	}
}

func victorOpsActionToModel(c *intschema.LogscaleActionVictorOpsAction) *models.LogScaleActionVictorOpsAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionVictorOpsAction{
		MessageType: c.MessageType,
		NotifyURL:   c.NotifyUrl,
		UseProxy:    c.UseProxy,
	}
}

func victorOpsActionFromModel(c *models.LogScaleActionVictorOpsAction) *intschema.LogscaleActionVictorOpsAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionVictorOpsAction{
		MessageType: c.MessageType,
		NotifyUrl:   c.NotifyURL,
		UseProxy:    c.UseProxy,
	}
}

func uploadFileActionToModel(c *intschema.LogscaleActionUploadFileAction) *models.LogScaleActionUploadFileAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionUploadFileAction{
		FileName: c.FileName,
	}
}

func uploadFileActionFromModel(c *models.LogScaleActionUploadFileAction) *intschema.LogscaleActionUploadFileAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionUploadFileAction{
		FileName: c.FileName,
	}
}

func webhookActionToModel(c *intschema.LogscaleActionWebhookAction) *models.LogScaleActionWebhookAction {
	if c == nil {
		return nil
	}
	return &models.LogScaleActionWebhookAction{
		BodyTemplate: c.BodyTemplate,
		Headers:      c.Headers,
		IgnoreSsl:    c.IgnoreSsl,
		Method:       models.WebhookActionHTTPMethod(c.Method),
		URL:          c.Url,
		UseProxy:     c.UseProxy,
	}
}

func webhookActionFromModel(c *models.LogScaleActionWebhookAction) *intschema.LogscaleActionWebhookAction {
	if c == nil {
		return nil
	}
	return &intschema.LogscaleActionWebhookAction{
		BodyTemplate: c.BodyTemplate,
		Headers:      c.Headers,
		IgnoreSsl:    c.IgnoreSsl,
		Method:       string(c.Method),
		Url:          c.URL,
		UseProxy:     c.UseProxy,
	}
}
