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
)

// BehaviorAction is an enum.
var BehaviorAction = newEnum("BehaviorAction", []value[configunstable.BehaviorAction]{
	{
		v1:    configunstable.BehaviorActionALERTWARN,
		alias: "ALERT_WARN",
	},
	{
		v1:    configunstable.BehaviorActionALERTCRITICAL,
		alias: "ALERT_CRITICAL",
	},
	{
		v1:    configunstable.BehaviorActionDROP,
		alias: "DROP",
	},
})
