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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/prettyenum"
)

var Dataset = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the dataset. Generated from `name` if omitted. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the dataset. Can be changed after creation.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Free-form description of the dataset.",
	},
	"configuration": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Description: "Configuration block selecting the dataset type and its match criteria.",
		Elem: &schema.Resource{
			Schema: datasetConfigurationSchema,
		},
	},
}

var datasetConfigurationSchema = map[string]*schema.Schema{
	"type": {
		Type:             schema.TypeString,
		Required:         true,
		DiffSuppressFunc: diffSuppressDatasetType,
		Description:      "Dataset type. Determines which of `trace_dataset` or `log_dataset` must be set.",
	},
	"trace_dataset": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Trace-specific dataset configuration. Set only when `type` is a trace type.",
		Elem: &schema.Resource{
			Schema: traceDatasetConfigurationSchema,
		},
	},
	"log_dataset": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Log-specific dataset configuration. Set only when `type` is a log type.",
		Elem: &schema.Resource{
			Schema: logDatasetConfigurationSchema,
		},
	},
}

var traceDatasetConfigurationSchema = map[string]*schema.Schema{
	"match_criteria": TraceSearchFilterSchema,
}

var logDatasetConfigurationSchema = map[string]*schema.Schema{
	"match_criteria": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Log search filter that defines which logs are included in this dataset.",
		Elem: &schema.Resource{
			Schema: LogSearchSchema,
		},
	},
}

// diffSuppressDatasetType sanitizes and then diffs two span filter match type payloads.
func diffSuppressDatasetType(_, old, new string, _ *schema.ResourceData) bool {
	if old == new {
		return true
	}
	mtOld, err := prettyenum.NewDatasetDatasetType(old)
	if err != nil {
		return false
	}
	mtNew, err := prettyenum.NewDatasetDatasetType(new)
	if err != nil {
		return false
	}
	return mtOld.Model() == mtNew.Model()
}
