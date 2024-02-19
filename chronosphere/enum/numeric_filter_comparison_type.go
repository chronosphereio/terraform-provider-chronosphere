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
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// NumericFilterComparisonType is an enum.
var NumericFilterComparisonType = newEnum("NumericFilterComparisonType", []value[
	configunstable.NumericFilterComparisonType,
	configv1.NumericFilterComparisonType,
]{
	{
		legacy: configunstable.NumericFilterComparisonTypeEQUAL,
		v1:     configv1.NumericFilterComparisonTypeEQUAL,
		alias:  "EQUAL",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeNOTEQUAL,
		v1:     configv1.NumericFilterComparisonTypeNOTEQUAL,
		alias:  "NOT_EQUAL",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeGREATERTHAN,
		v1:     configv1.NumericFilterComparisonTypeGREATERTHAN,
		alias:  "GREATER_THAN",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeGREATERTHANOREQUAL,
		v1:     configv1.NumericFilterComparisonTypeGREATERTHANOREQUAL,
		alias:  "GREATER_THAN_OR_EQUAL",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeLESSTHAN,
		v1:     configv1.NumericFilterComparisonTypeLESSTHAN,
		alias:  "LESS_THAN",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeLESSTHANOREQUAL,
		v1:     configv1.NumericFilterComparisonTypeLESSTHANOREQUAL,
		alias:  "LESS_THAN_OR_EQUAL",
	},
})
