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
		Description:  "ID of the bucket the recording rule belongs to. At least one of `bucket_id` or `execution_group` must be set; if both are set their values must match.",
	},
	"execution_group": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: executionGroupFields,
		Description:  "Slug of the execution group in which the rule is evaluated. Rules in the same group run sequentially at the configured interval; all rules in a group must finish before the next iteration starts. At least one of `bucket_id` or `execution_group` must be set.",
	},
	"execution_mode": Enum{
		Value:       enum.RecordingRuleExecutionModeType.ToStrings(),
		Optional:    true,
		Description: "Execution mode controlling whether the recording rule is active.",
	}.Schema(),
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the recording rule. Can be changed after creation.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the recording rule. Generated from `name` if omitted. Immutable after creation.",
	},
	"labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Key/value labels added to every series produced by this recording rule.",
	},
	"interval": Duration{
		Optional:    true,
		Description: "Evaluation interval (e.g. `30s`, `1m`). Defaults to `60s` when unset.",
	}.Schema(),
	"expr": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "PromQL expression evaluated at each interval. The result is written to a new series named by `metric_name` (or `name` if unset).",
	},
	"metric_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the output time series produced by `expr`. Must be a valid metric name. Defaults to `name` if omitted.",
	},
}
