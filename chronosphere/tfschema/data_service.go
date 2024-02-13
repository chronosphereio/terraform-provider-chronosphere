package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var DataService = map[string]*schema.Schema{
	"slug": {
		Type:     schema.TypeString,
		Required: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"description": {
		Type:     schema.TypeString,
		Computed: true,
	},
}
