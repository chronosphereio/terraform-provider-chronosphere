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

// LogScalePagerDutyActionSeverity is an enum.
var LogScalePagerDutyActionSeverity = newEnum("LogScaleActionPagerDutyActionSeverity", []value[models.PagerDutyActionSeverity]{
	{
		v1:        "INVALID",
		isDefault: true,
	},
	{
		v1:    models.PagerDutyActionSeverityCRITICAL,
		alias: "CRITICAL",
	},
	{
		v1:    models.PagerDutyActionSeverityERROR,
		alias: "ERROR",
	},
	{
		v1:    models.PagerDutyActionSeverityWARNING,
		alias: "WARNING",
	},
	{
		v1:    models.PagerDutyActionSeverityINFO,
		alias: "INFO",
	},
})
