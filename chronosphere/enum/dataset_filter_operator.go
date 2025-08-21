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

// PartitionFilterOperator is an enum.
var PartitionFilterOperator = newEnum("PartitionFilterOperator", []value[configunstable.PartitionFilterOperator]{
	{
		v1:    configunstable.PartitionFilterOperatorIN,
		alias: "IN",
	},
	{
		v1:    configunstable.PartitionFilterOperatorNOTIN,
		alias: "NOT_IN",
	},
})
