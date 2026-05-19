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

package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var actionFields = []string{
	"email_action",
	"humio_action",
	"ops_genie_action",
	"pager_duty_action",
	"slack_action",
	"slack_post_message_action",
	"victor_ops_action",
	"webhook_action",
	"upload_file_action",
}

var LogscaleAction = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the LogScale action. Generated from `name` if omitted. Immutable after creation.",
	},
	"repository": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Name of the LogScale repository the action belongs to. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the LogScale action.",
	},
	"email_action":              LogscaleEmailActionSchema,
	"humio_action":              LogscaleHumioRepoActionSchema,
	"ops_genie_action":          LogscaleOpsGenieActionSchema,
	"pager_duty_action":         LogscalePagerDutyActionSchema,
	"slack_action":              LogscaleSlackActionSchema,
	"slack_post_message_action": LogscaleSlackPostMessageActionSchema,
	"victor_ops_action":         LogscaleVictorOpsActionSchema,
	"webhook_action":            LogscaleWebhookActionSchema,
	"upload_file_action":        LogscaleUploadFileActionSchema,
}

var LogscaleEmailActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Send email when the alert triggers. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"recipients": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MinItems:    1,
				Required:    true,
				Description: "List of email addresses to send the message to.",
			},
			"subject_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Subject of the email. May be templated with values from the query result.",
			},
			"body_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Body of the email. May be templated with values from the query result.",
			},
			"attach_csv": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, attaches the query result set as a CSV file.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}

var LogscaleHumioRepoActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Forward results to another LogScale (Humio) repository via its ingest token. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ingest_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Ingest token for the target repository.",
			},
		},
	},
}

var LogscaleOpsGenieActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Send an OpsGenie alert when the alert triggers. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "OpsGenie webhook URL to send the request to.",
			},
			"ops_genie_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Key used to authenticate with OpsGenie.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}

var LogscalePagerDutyActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Send a PagerDuty event when the alert triggers. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"severity": Enum{
				Value:       enum.LogScalePagerDutyActionSeverity.ToStrings(),
				Required:    true,
				Description: "Severity attached to the PagerDuty event.",
			}.Schema(),
			"routing_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Routing key used to authenticate with PagerDuty.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}

var LogscaleSlackActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Post a message to a Slack incoming webhook when the alert triggers. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Slack incoming webhook URL to send the request to.",
			},
			"fields": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Fields to include in the Slack message. Values may be templated with the query result.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}

var LogscaleSlackPostMessageActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Post a message to Slack channels using the Slack `chat.postMessage` API. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Slack API token used to authenticate the request.",
			},
			"channels": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
				Required:    true,
				Description: "List of Slack channels to post the message to.",
			},
			"fields": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Fields to include in the Slack message. Values may be templated with the query result.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}

var LogscaleVictorOpsActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Send a VictorOps (Splunk On-Call) event when the alert triggers. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"message_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the VictorOps message to send (for example, `CRITICAL`, `WARNING`, `INFO`).",
			},
			"notify_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VictorOps webhook URL to send the request to.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}

var LogscaleUploadFileActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Upload the query result as a file in LogScale. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name to use for the uploaded file.",
			},
		},
	},
}

var LogscaleWebhookActionSchema = &schema.Schema{
	Type:         schema.TypeList,
	Optional:     true,
	MinItems:     0,
	MaxItems:     1,
	ExactlyOneOf: actionFields,
	Description:  "Send an HTTP or HTTPS webhook when the alert triggers. Exactly one action type must be set.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method": Enum{
				Value:       enum.LogScaleWebhookActionHTTPMethod.ToStrings(),
				Required:    true,
				Description: "HTTP method used for the webhook request.",
			}.Schema(),
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL to send the HTTP or HTTPS request to.",
			},
			"headers": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Headers to include on the HTTP or HTTPS request.",
			},
			"body_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Body of the request. May be templated with values from the query result.",
			},
			"ignore_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, skips SSL certificate verification for the request.",
			},
			"use_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If `true`, sends the request through the configured outbound proxy.",
			},
		},
	},
}
