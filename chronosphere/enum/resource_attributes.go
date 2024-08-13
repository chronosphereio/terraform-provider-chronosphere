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

package enum

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// ResourceAttributesFlattenMode is an enum.
var ResourceAttributesFlattenMode = newEnum("ResourceAttributesFlattenMode", []value[models.ResourceAttributesFlattenMode]{
	{
		v1:    models.ResourceAttributesFlattenModeMERGE,
		alias: "MERGE",
	},
	{
		v1:    models.ResourceAttributesFlattenModeOVERWRITE,
		alias: "OVERWRITE",
	},
	{
		v1:    models.ResourceAttributesFlattenModeIGNORE,
		alias: "IGNORE",
	},
})

// ResourceAttributesFilterMode is an enum.
var ResourceAttributesFilterMode = newEnum("ResourceAttributesFilterMode", []value[models.ResourceAttributesFilterMode]{
	{
		v1:    models.ResourceAttributesFilterModeAPPENDDEFAULTEXCLUDEKEYS,
		alias: "APPEND_DEFAULT_EXCLUDE_KEYS",
	},
	{
		v1:    models.ResourceAttributesFilterModeCUSTOMEXCLUDEKEYS,
		alias: "CUSTOM_EXCLUDE_KEYS",
	},
})
