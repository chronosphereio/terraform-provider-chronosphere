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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var MappingRule = map[string]*schema.Schema{
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
	"bucket_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"filter": Filter{
		KVDelimiter: aggregationfilter.MappingRuleDelimiter,
	}.Schema(),
	"aggregations": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: Enum{
			Value: enum.AggregationType.ToStrings(),
		}.Schema(),
	},
	// Storage policies to apply to the mapped metrics.
	"storage_policy": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"resolution": Duration{
					Required: true,
				}.Schema(),
				"retention": Duration{
					Required: true,
				}.Schema(),
			},
		},
		Deprecated: "use `interval` instead",
	},
	"drop": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	// Whether or not to drop the timestamp.
	"drop_timestamp": {
		Type:       schema.TypeBool,
		Optional:   true,
		Default:    false,
		Deprecated: "drop timestamp is no longer supported",
	},
	"interval": {
		Type:     schema.TypeString,
		Optional: true,
		// When no interval is specified, a server-side default is used.
		Computed:      true,
		ConflictsWith: []string{"storage_policy"},
	},
	"mode": Enum{
		Value:    enum.MappingModeType.ToStrings(),
		Optional: true,
	}.Schema(),
}
