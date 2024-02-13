package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var BlackholeAlertNotifier = map[string]*schema.Schema{
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
}
