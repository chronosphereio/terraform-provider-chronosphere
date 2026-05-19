// Copyright 2025 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
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

var LogRetentionConfig = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the log retention config.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the log retention config. Generated from `name` if omitted. Immutable after creation.",
	},
	"mode": Enum{
		Value:       enum.LogRetentionConfigMode.ToStrings(),
		Required:    true,
		Description: "Mode that determines how matching logs are retained in long-term (Iceberg) storage.",
	}.Schema(),
	"filter": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Log query filter. The retention policy applies only to logs that match.",
	},
	"retention_days": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Number of days to retain matching logs in long-term (Iceberg) storage after they are exported. When multiple configs overlap, the longest retention wins.",
	},
}
