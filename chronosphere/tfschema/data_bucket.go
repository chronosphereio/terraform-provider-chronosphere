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

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var dataBucketOneOfAddressFields = []string{"name", "slug"} // fields used to address the bucket.

var DataBucket = map[string]*schema.Schema{
	"slug": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: dataBucketOneOfAddressFields,
		Description:  "Slug of the bucket to look up. Exactly one of `slug` or `name` must be set.",
	},
	"name": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: dataBucketOneOfAddressFields,
		Description:  "Name of the bucket to look up. Exactly one of `slug` or `name` must be set.",
	},
	"description": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Read-only: free-form description of the bucket.",
	},
	"labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Read-only: key/value labels attached to the bucket.",
	},
}
