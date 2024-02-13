package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var Collection = map[string]*schema.Schema{
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"team_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"notification_policy_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
