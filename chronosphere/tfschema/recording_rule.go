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

var executionGroupFields = []string{"bucket_id", "execution_group"}

var RecordingRule = map[string]*schema.Schema{
	"bucket_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: executionGroupFields,
	},
	"execution_group": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: executionGroupFields,
	},
	"execution_mode": Enum{
		Value:    enum.RecordingRuleExecutionModeType.ToStrings(),
		Optional: true,
	}.Schema(),
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"interval": Duration{
		Optional: true,
	}.Schema(),
	"expr": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_name": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
