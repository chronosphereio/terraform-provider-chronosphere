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
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"repository": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"recipients": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MinItems: 1,
				Required: true,
			},
			"subject_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"body_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attach_csv": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ingest_token": {
				Type:     schema.TypeString,
				Required: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ops_genie_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"severity": Enum{
				Value:    enum.LogScalePagerDutyActionSeverity.ToStrings(),
				Required: true,
			}.Schema(),
			"routing_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"channels": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"message_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notify_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_name": {
				Type:     schema.TypeString,
				Required: true,
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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method": Enum{
				Value:    enum.LogScaleWebhookActionHTTPMethod.ToStrings(),
				Required: true,
			}.Schema(),
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"headers": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"body_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	},
}
