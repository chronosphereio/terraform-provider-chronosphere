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

// BehaviorThresholdType is an enum.
var BehaviorThresholdType = newEnum("BehaviorThresholdType", []value[configv1.BehaviorThresholdType]{
	{
		v1:    configv1.BehaviorThresholdTypeMONTHLYVOLUME,
		alias: "MONTHLY_VOLUME",
	},
	{
		v1:    configv1.BehaviorThresholdTypeWEEKLYVOLUME,
		alias: "WEEKLY_VOLUME",
	},
	{
		v1:    configv1.BehaviorThresholdTypeDAILYVOLUME,
		alias: "DAILY_VOLUME",
	},
	{
		v1:    configv1.BehaviorThresholdTypeINSTANTRATE,
		alias: "INSTANT_RATE",
	},
})
