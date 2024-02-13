package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// TestResource is used exclusively for unit testing.
var TestResource = map[string]*schema.Schema{
	"some_string": {
		Type: schema.TypeString,
	},
	"some_bool": {
		Type: schema.TypeBool,
	},
	"some_float": {
		Type: schema.TypeFloat,
	},
	"some_int": {
		Type: schema.TypeInt,
	},
	"some_string_list": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"some_object_set": {
		Type: schema.TypeSet,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inner_string": {
					Type: schema.TypeString,
				},
				"inner_bool": {
					Type: schema.TypeBool,
				},
			},
		},
	},
	"some_string_map": {
		Type: schema.TypeMap,
		Elem: &schema.Schema{Type: schema.TypeString},
	},
	"some_object": {
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inner_string": {
					Type: schema.TypeString,
				},
				"inner_bool": {
					Type: schema.TypeBool,
				},
			},
		},
	},
	"optional_object": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inner_string_list": {
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	"optional_string_list": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
	},
	"collection_id": {
		Type: schema.TypeString,
	},
	// Intentionally match the exact notifiers schema since it's the only
	// list of TF IDs, and we want to ensure it works correctly.
	"notifiers": NotificationRouteSchema.Elem.(*schema.Resource).Schema["notifiers"],

	// Matches send_resolved in real resources.
	"optional_bool_with_default": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},
	"computed_and_not_optional": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"computed_and_optional": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"dashboard_json": {
		Type: schema.TypeString,
	},
}
