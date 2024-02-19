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

package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Enum defines the parameters of an enum field in a Terraform schema.
type Enum struct {
	Value    enum.Enum[string, string]
	Required bool
	Optional bool
	ForceNew bool
}

// Schema returns the Terraform of the enum.
func (e Enum) Schema() *schema.Schema {
	return withDiffSuppress(e, &schema.Schema{
		Type:             schema.TypeString,
		Required:         e.Required,
		Optional:         e.Optional,
		ForceNew:         e.ForceNew,
		ValidateDiagFunc: e.Value.Validate,
	})
}

// Normalize implements typeset.Normalizer.
func (e Enum) Normalize(v any) any {
	return e.Value.V1(v.(string))
}
