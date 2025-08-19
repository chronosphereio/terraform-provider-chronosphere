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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var LogControlConfig = map[string]*schema.Schema{
	"rules": {
		Type:     schema.TypeList,
		Elem:     logControlRuleResource,
		Optional: true,
	},
}

var logControlRuleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"mode": Enum{
			Value:    enum.LogControlRuleMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"filter": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Log query syntax to select logs. Only matching logs will have control action applied.",
		},
		"type": Enum{
			Value:    enum.LogControlRuleType.ToStrings(),
			Required: true,
		}.Schema(),
		"sample": {
			Type:     schema.TypeList,
			Elem:     logControlRuleSampleResource,
			Optional: true,
			MaxItems: 1,
		},
		"drop_field": {
			Type:     schema.TypeList,
			Elem:     logControlRuleDropFieldResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleSampleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"rate": {
			Type:        schema.TypeFloat,
			Required:    true,
			Description: "Percentage of matching logs to keep. Must be in the range (0, 1].",
		},
	},
}

var logControlRuleDropFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"field_regex": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Regular expression to match the field name(s) to drop.",
		},
		"parent_path": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logFieldPathResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"selector": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "LogQL Selector to indicate field path. Use 'parent[child]' syntax to indicate nesting.",
		},
	},
}
