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

// ConsumptionBudgetResource is an enum.
var ConsumptionBudgetResource = newEnum("ConsumptionBudgetResource", []value[configv1.Configv1ConsumptionBudgetResource]{
	{
		v1:    configv1.Configv1ConsumptionBudgetResourceLOGPERSISTEDBYTES,
		alias: "LOG_PERSISTED_BYTES",
	},
})

// ConsumptionBudgetThresholdType is an enum.
var ConsumptionBudgetThresholdType = newEnum("ConsumptionBudgetThresholdType", []value[configv1.ConsumptionBudgetThresholdType]{
	{
		v1:    configv1.ConsumptionBudgetThresholdTypeMONTHLYVOLUME,
		alias: "MONTHLY_VOLUME",
	},
	{
		v1:    configv1.ConsumptionBudgetThresholdTypeWEEKLYVOLUME,
		alias: "WEEKLY_VOLUME",
	},
	{
		v1:    configv1.ConsumptionBudgetThresholdTypeDAILYVOLUME,
		alias: "DAILY_VOLUME",
	},
	{
		v1:    configv1.ConsumptionBudgetThresholdTypeINSTANTRATE,
		alias: "INSTANT_RATE",
	},
})

// ConsumptionBudgetThresholdAction is an enum.
var ConsumptionBudgetThresholdAction = newEnum("ConsumptionBudgetThresholdAction", []value[configv1.ConsumptionBudgetThresholdAction]{
	{
		v1:    configv1.ConsumptionBudgetThresholdActionALERTWARN,
		alias: "ALERT_WARN",
	},
	{
		v1:    configv1.ConsumptionBudgetThresholdActionALERTCRITICAL,
		alias: "ALERT_CRITICAL",
	},
	{
		v1:    configv1.ConsumptionBudgetThresholdActionDROP,
		alias: "DROP",
	},
})
