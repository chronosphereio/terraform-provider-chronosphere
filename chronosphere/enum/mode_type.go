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

// RollupModeType is an enum.
var RollupModeType = newEnum("RollupModeType", []value[configv1.Configv1RollupRuleMode]{
	{
		v1:        configv1.Configv1RollupRuleModeENABLED,
		isDefault: true,
	},
	{
		v1:    configv1.Configv1RollupRuleModePREVIEW,
		alias: "PREVIEW",
	},
})

// MappingModeType is an enum.
var MappingModeType = newEnum("MappingModeType", []value[configv1.Configv1MappingRuleMode]{
	{
		v1:        configv1.Configv1MappingRuleModeENABLED,
		isDefault: true,
	},
	{
		v1:    configv1.Configv1MappingRuleModePREVIEW,
		alias: "PREVIEW",
	},
})

// DropRuleModeType is an enum.
var DropRuleModeType = newEnum("DropRuleModeType", []value[configv1.Configv1DropRuleMode]{
	{
		v1:    configv1.Configv1DropRuleModeENABLED,
		alias: "ENABLED",
	},
	{
		v1:    configv1.Configv1DropRuleModeDISABLED,
		alias: "DISABLED",
	},
	{
		v1:    configv1.Configv1DropRuleModePREVIEW,
		alias: "PREVIEW",
	},
})

// RecordingRuleExecutionModeType is an enum.
var RecordingRuleExecutionModeType = newEnum("RecordingRuleExecutionModeType", []value[configv1.RecordingRuleExecutionMode]{
	{
		v1:    configv1.RecordingRuleExecutionModeEXECUTIONMODEDEFAULT,
		alias: "DEFAULT",
	},
	{
		v1:    configv1.RecordingRuleExecutionModeEXECUTIONMODESYNCHRONIZED,
		alias: "SYNCHRONIZED",
	},
})
