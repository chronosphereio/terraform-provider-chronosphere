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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// TraceMetricsRuleGroupByType is an enum.
var TraceMetricsRuleGroupByType = newEnum("TraceMetricsRuleGroupByType", []value[
	string,
	models.GroupByKeyGroupByKeyType,
]{
	{
		v1:     models.GroupByKeyGroupByKeyTypeTAG,
		legacy: "TAG",
		alias:  "TAG",
	},
	{
		v1:     models.GroupByKeyGroupByKeyTypeOPERATION,
		legacy: "OPERATION",
		alias:  "OPERATION",
	},
	{
		v1:     models.GroupByKeyGroupByKeyTypeSERVICE,
		legacy: "SERVICE",
		alias:  "SERVICE",
	},
})
