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
	configunstable "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// ResourceAttributesFlattenMode is an enum.
var ResourceAttributesFlattenMode = newV1OnlyEnum("ResourceAttributesFlattenMode", []v1OnlyValue[configunstable.ResourceAttributesFlattenMode]{
	{
		v1:    configunstable.ResourceAttributesFlattenModeMERGE,
		alias: "MERGE",
	},
	{
		v1:    configunstable.ResourceAttributesFlattenModeOVERWRITE,
		alias: "OVERWRITE",
	},
	{
		v1:    configunstable.ResourceAttributesFlattenModeIGNORE,
		alias: "IGNORE",
	},
})

// ResourceAttributesFilterMode is an enum.
var ResourceAttributesFilterMode = newV1OnlyEnum("ResourceAttributesFilterMode", []v1OnlyValue[configunstable.ResourceAttributesFilterMode]{
	{
		v1:    configunstable.ResourceAttributesFilterModeAPPENDDEFAULTEXCLUDEKEYS,
		alias: "APPEND_DEFAULT_EXCLUDE_KEYS",
	},
	{
		v1:    configunstable.ResourceAttributesFilterModeCUSTOMEXCLUDEKEYS,
		alias: "CUSTOM_EXCLUDE_KEYS",
	},
})
