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
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// Permission represents a mapping between the external API permissions and the
// v1 version
var Permission = newEnum("Permission", []value[configv1.MetricsRestrictionPermission]{
	{
		isDefault:   true,
		legacyAlias: "UNKNOWN_PERMISSION",
	},
	{
		v1:          configv1.MetricsRestrictionPermissionREAD,
		legacyAlias: "READ_PERMISSION",
		alias:       "READ_ONLY",
	},
	{
		v1:          configv1.MetricsRestrictionPermissionWRITE,
		legacyAlias: "WRITE_PERMISSION",
		alias:       "WRITE_ONLY",
	},
	{
		v1:          configv1.MetricsRestrictionPermissionREADWRITE,
		legacyAlias: "READWRITE_PERMISSION",
		alias:       "READ_AND_WRITE",
	},
})
