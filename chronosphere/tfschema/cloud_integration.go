// Copyright 2026 Chronosphere Inc.
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
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

// CloudIntegration is the schema for a cloud integration.
//
// provider_config is deliberately an opaque JSON string rather than typed
// per-provider blocks: the set of supported providers is a server-side
// concern, so the provider schema never enumerates it and adding a new cloud
// provider requires no Terraform provider change or release.
var CloudIntegration = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the cloud integration.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the integration. Generated from `name` if omitted. Immutable after creation.",
	},
	"state": Enum{
		Value:       enum.CloudIntegrationState.ToStrings(),
		Optional:    true,
		Description: "Operational state of the integration: ENABLED, DISABLED, or PREVIEW.",
	}.Schema(),
	"external_connection_slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Slug of the ExternalConnection used for credentials. Optional depending on the provider.",
	},
	"metric_labels": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Labels applied to all metrics emitted from this integration.",
	},
	"provider_type": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Identifies the cloud provider (e.g. `gcp`). The supported set is defined by the Chronosphere API.",
	},
	"provider_config": {
		Type:             schema.TypeString,
		Required:         true,
		DiffSuppressFunc: JSONCloudIntegrationProviderConfigDiffSuppress,
		ValidateFunc:     ValidateCloudIntegrationProviderConfig,
		Description:      "Provider-specific configuration as a JSON object (use `jsonencode`). Structure depends on `provider_type`.",
	},
}

// JSONCloudIntegrationProviderConfigDiffSuppress returns true if the diff
// between the old and new provider config JSON values should be suppressed,
// i.e. the deserialized objects are equal.
func JSONCloudIntegrationProviderConfigDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	if old == new {
		return true
	}

	if old == "" || new == "" {
		return false
	}

	var oldConfig map[string]any
	if err := json.Unmarshal([]byte(old), &oldConfig); err != nil {
		return false
	}

	var newConfig map[string]any
	if err := json.Unmarshal([]byte(new), &newConfig); err != nil {
		return false
	}

	return reflect.DeepEqual(oldConfig, newConfig)
}

// ValidateCloudIntegrationProviderConfig is a SchemaValidateFunc which tests
// that the provided value is a JSON object. Provider-specific structure is
// validated server-side (surfaced at plan time via dry-run).
func ValidateCloudIntegrationProviderConfig(i any, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return warnings, errors
	}

	var config map[string]any
	if err := json.Unmarshal([]byte(v), &config); err != nil {
		errors = append(errors, fmt.Errorf("%s must be a JSON object: %w", k, err))
	}

	return warnings, errors
}
