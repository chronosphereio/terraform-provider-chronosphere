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

// LogControlRuleType is an enum.
var LogControlRuleType = newEnum("LogControlRuleType", []value[configv1.Configv1LogControlRuleType]{
	{
		v1:    configv1.Configv1LogControlRuleTypeDROP,
		alias: "DROP",
	},
	{
		v1:    configv1.Configv1LogControlRuleTypeSAMPLE,
		alias: "SAMPLE",
	},
	{
		v1:    configv1.Configv1LogControlRuleTypeDROPFIELD,
		alias: "DROP_FIELD",
	},
})
