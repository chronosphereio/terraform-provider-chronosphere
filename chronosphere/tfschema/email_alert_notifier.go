package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var EmailAlertNotifier = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"send_resolved": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},
	"to": {
		Type:     schema.TypeString,
		Required: true,
	},
	"html": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"text": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
