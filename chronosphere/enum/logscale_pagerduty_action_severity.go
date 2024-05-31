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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// LogScalePagerDutyActionSeverity is an enum.
var LogScalePagerDutyActionSeverity = newEnum("LogScaleActionPagerDutyActionSeverity", []value[
	string,
	models.LogScaleActionPagerDutyActionSeverity,
]{
	{
		legacy:    "INVALID",
		v1:        "INVALID",
		isDefault: true,
	},
	{
		v1:     models.LogScaleActionPagerDutyActionSeverityCRITICAL,
		legacy: "CRITICAL",
		alias:  "CRITICAL",
	},
	{
		v1:     models.LogScaleActionPagerDutyActionSeverityERROR,
		legacy: "ERROR",
		alias:  "ERROR",
	},
	{
		v1:     models.LogScaleActionPagerDutyActionSeverityWARNING,
		legacy: "WARNING",
		alias:  "WARNING",
	},
	{
		v1:     models.LogScaleActionPagerDutyActionSeverityINFO,
		legacy: "INFO",
		alias:  "INFO",
	},
})
