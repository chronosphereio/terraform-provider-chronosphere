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
)

var TraceTailSamplingRules = map[string]*schema.Schema{
	"default_sample_rate": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"enabled": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"sample_rate": SampleRateSchema,
			},
		},
	},
	"rules": rulesSchema,
}

var rulesSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true, // this can be set to required: true once all callers are upgraded
			},
			"system_name": {
				Type:     schema.TypeString,
				Optional: true, // this can be set to required: true once all callers are upgraded
			},
			"sample_rate": SampleRateSchema,
			"filter":      TraceSearchFilterSchema,
		},
	},
}

var SampleRateSchema = &schema.Schema{
	Type:             schema.TypeFloat,
	Required:         true,
	ValidateDiagFunc: float64RangeValidator(0, 1.0),
}
