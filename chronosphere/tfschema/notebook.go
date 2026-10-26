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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
)

var Notebook = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the notebook. Can be changed after creation.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the notebook. Generated from `name` if omitted. Immutable after creation.",
	},
	"collection_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "ID of the collection that owns this notebook. A collection's `id`, or `USER:<slug>` for a user's personal collection.",
	},
	"notebook_json": {
		Type:                  schema.TypeString,
		Required:              true,
		DiffSuppressFunc:      notebookJSONDiffSuppress,
		DiffSuppressOnRefresh: true,
		Description:           "JSON payload describing the notebook's content. Must be a JSON object. Wrap with `jsonencode({...})` in HCL. The provider compares canonical JSON, so key ordering and whitespace do not cause spurious plans.",
	},
	"labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Key/value labels attached to the notebook for organization and filtering.",
	},
}

// notebookJSONDiffSuppress treats two notebook JSON payloads as equal when
// they encode the same document, whatever their key order or whitespace.
// Unlike dashboards, the provider knows nothing about a notebook's content,
// so it drops no fields. See xjson.Canonical for why the SDK's
// structure.SuppressJsonDiff is not used.
func notebookJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	canonicalOld, err := xjson.Canonical(old)
	if err != nil {
		return false
	}

	canonicalNew, err := xjson.Canonical(new)
	if err != nil {
		return false
	}

	return canonicalOld == canonicalNew
}
