// Copyright 2025 Chronosphere Inc.
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

// PlaintextParserMode is an enum.
var PlaintextParserMode = newEnum("PlaintextParserMode", []value[configv1.Configv1PlaintextParserMode]{
	{
		v1:        configv1.Configv1PlaintextParserModeENABLED,
		isDefault: true,
	},
	{
		v1:    configv1.Configv1PlaintextParserModeDISABLED,
		alias: "DISABLED",
	},
})

// LogFieldParserMode is an enum.
var LogFieldParserMode = newEnum("LogFieldParserMode", []value[configv1.Configv1LogFieldParserMode]{
	{
		v1:        configv1.Configv1LogFieldParserModeENABLED,
		isDefault: true,
	},
	{
		v1:    configv1.Configv1LogFieldParserModeDISABLED,
		alias: "DISABLED",
	},
})

// LogParserType is an enum.
var LogParserType = newEnum("LogParserType", []value[configv1.LogParserParserType]{
	{
		v1:    configv1.LogParserParserTypeJSON,
		alias: "JSON",
	},
	{
		v1:    configv1.LogParserParserTypeREGEX,
		alias: "REGEX",
	},
	{
		v1:    configv1.LogParserParserTypeKEYVALUE,
		alias: "KEY_VALUE",
	},
})
