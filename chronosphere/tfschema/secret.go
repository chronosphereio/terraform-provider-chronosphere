// Copyright 2026 Chronosphere Inc.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WriteOnlySecret describes the attribute named Name+"_wo", holding a
// credential that the API stores as ciphertext and never returns. Pair it with
// the SecretVersionSchema attribute named Name+"_wo_version".
type WriteOnlySecret struct {
	Name        string
	Description string
	Required    bool
}

// Schema returns the Terraform schema of the write-only secret.
func (s WriteOnlySecret) Schema() *schema.Schema {
	return &schema.Schema{
		Type:      schema.TypeString,
		Required:  s.Required,
		Optional:  !s.Required,
		WriteOnly: true,
		Sensitive: true,
		Description: s.Description + " Write-only: never stored in Terraform state, and " +
			"never returned by the API. Terraform cannot see it change, so editing it " +
			"alone produces no plan: bump `" + s.Name + "_wo_version` to roll the new " +
			"value out.",
	}
}

// SecretVersionSchema builds the companion attribute that rolls out a change to
// the write-only credential named secret+"_wo".
func SecretVersionSchema(secret string) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Description: "Increment this whenever `" + secret + "_wo` changes. Terraform cannot " +
			"detect a change to a write-only attribute on its own, so bumping this is " +
			"what turns an edit to `" + secret + "_wo` into a plan; whatever the " +
			"attribute holds at apply time is then sent. The API does not store the " +
			"version, so an imported test starts at 0 and its first plan shows a diff here.",
	}
}
