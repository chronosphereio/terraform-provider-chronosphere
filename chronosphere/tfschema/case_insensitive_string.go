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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CaseInsensitiveString defines the parameters of a case-insensitive string
// field in a Terraform schema.
type CaseInsensitiveString struct {
	Required bool
}

// Schema returns the Terraform schema of the string.
func (s CaseInsensitiveString) Schema() *schema.Schema {
	return withDiffSuppress(s, &schema.Schema{
		Type:     schema.TypeString,
		Required: s.Required,
		Optional: !s.Required,
	})
}

// Normalize implements typeset.Normalizer.
func (s CaseInsensitiveString) Normalize(v any) any {
	return strings.ToLower(v.(string))
}
