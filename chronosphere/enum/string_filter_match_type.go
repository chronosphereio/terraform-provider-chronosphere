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
	configunstable "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// StringFilterMatchType is an enum.
var StringFilterMatchType = newEnum("StringFilterMatchType", []value[
	configunstable.StringFilterStringFilterMatchType,
	configv1.StringFilterStringFilterMatchType,
]{
	{
		legacy: configunstable.StringFilterStringFilterMatchTypeEXACT,
		v1:     configv1.StringFilterStringFilterMatchTypeEXACT,
		alias:  "EXACT",
	},
	{
		legacy: configunstable.StringFilterStringFilterMatchTypeREGEX,
		v1:     configv1.StringFilterStringFilterMatchTypeREGEX,
		alias:  "REGEX",
	},
})