// Copyright 2023 Chronosphere Inc.
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

package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func newClassicDashboard(a api) entityType {
	const entity = "ClassicDashboard"
	return entityType{
		API:                  a,
		GoType:               fmt.Sprintf("%s%s", a.GoPrefix, entity),
		SwaggerType:          entity,
		SwaggerModel:         "GrafanaDashboard",
		SwaggerClient:        fmt.Sprintf("%s.%s", a.Client, entity),
		SwaggerClientPackage: strcase.ToSnake(entity),
		DryRun:               true,
		UpdateUnsupported:    false,
	}
}
