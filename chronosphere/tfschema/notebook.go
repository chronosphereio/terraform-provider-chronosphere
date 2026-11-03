package tfschema

import (
	"encoding/json"

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
		Description:           "JSON payload describing the notebook's content. Must be a JSON object. Wrap with `jsonencode({...})` in HCL. The provider normalizes the JSON before diffing, so cosmetic differences (key ordering, whitespace) do not cause spurious plans.",
	},
	"labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Key/value labels attached to the notebook for organization and filtering.",
	},
}

// notebookJSONDiffSuppress normalizes and then diffs two notebook JSON payloads,
// as dashboardJSONDiffSuppress does for dashboards. Unlike dashboards, the
// provider knows nothing about a notebook's content, so it drops no fields.
func notebookJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	normalizedOld, err := normalizedNotebookJSON(old)
	if err != nil {
		return false
	}

	normalizedNew, err := normalizedNotebookJSON(new)
	if err != nil {
		return false
	}

	return normalizedOld == normalizedNew
}

// normalizedNotebookJSON re-encodes a notebook JSON object with sorted keys and
// no insignificant whitespace, so two encodings of the same document compare
// equal as strings.
func normalizedNotebookJSON(data string) (string, error) {
	var notebook map[string]any
	if err := xjson.Unmarshal([]byte(data), &notebook); err != nil {
		return "", err
	}

	normalized, err := json.Marshal(notebook)
	if err != nil {
		return "", err
	}

	return string(normalized), nil
}
