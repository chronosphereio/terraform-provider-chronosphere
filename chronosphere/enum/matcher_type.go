// Copyright 2023 Chronosphere Inc.
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

// MatcherType is an enum.
var MatcherType = newEnum("MatcherType", []value[
	string,
	configv1.Configv1LabelMatcherMatcherType,
]{
	{
		legacy:    "INVALID_MATCHER_TYPE",
		isDefault: true,
	},
	{
		legacy: "EXACT_MATCHER_TYPE",
		v1:     configv1.Configv1LabelMatcherMatcherTypeEXACT,
		alias:  "EXACT",
	},
	{
		legacy: "REGEXP_MATCHER_TYPE",
		v1:     configv1.Configv1LabelMatcherMatcherTypeREGEX,
		alias:  "REGEX",
	},
})
