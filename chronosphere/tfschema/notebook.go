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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/notebook"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Notebook = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Optional:    true,
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
		Description: "ID of the collection that owns this notebook.",
	},
	"notebook_json": {
		Type:                  schema.TypeString,
		Required:              true,
		DiffSuppressFunc:      notebookJSONDiffSuppress,
		DiffSuppressOnRefresh: true,
		Description:           "JSON payload describing the notebook's cells and content. Wrap with `jsonencode({...})` in HCL. The provider compares this value semantically, so key ordering and whitespace do not cause spurious plans.",
	},
}

// notebookJSONDiffSuppress reports whether two notebook JSON payloads are
// semantically equal, ignoring key ordering and insignificant whitespace.
//
// Unlike dashboards, the API stores and returns notebook_json verbatim and
// populates no fields inside it, so there is nothing to strip before
// comparing. Canonicalizing is enough. If the API later starts writing into
// the payload, this is the place to add a sanitizer.
func notebookJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	canonicalOld, err := notebook.CanonicalNotebookJSON(old)
	if err != nil {
		return false
	}

	canonicalNew, err := notebook.CanonicalNotebookJSON(new)
	if err != nil {
		return false
	}

	return canonicalOld == canonicalNew
}
