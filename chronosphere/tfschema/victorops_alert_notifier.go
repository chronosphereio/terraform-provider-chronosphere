package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var VictoropsAlertNotifier = map[string]*schema.Schema{
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
	"api_key": {
		Type:      schema.TypeString,
		Required:  true,
		Sensitive: true,
	},
	"api_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"routing_key": {
		Type:     schema.TypeString,
		Required: true,
	},
	"state_message": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"message_type": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"custom_fields": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"monitoring_tool": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"entity_display_name": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"basic_auth_username": {
		Type:          schema.TypeString,
		Optional:      true,
		RequiredWith:  []string{"basic_auth_password"},
		ConflictsWith: []string{"bearer_token"},
	},
	"basic_auth_password": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"basic_auth_password"},
		Sensitive:    true,
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
	},
	"proxy_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"tls_insecure_skip_verify": {
		Type:     schema.TypeBool,
		Optional: true,
	},
}
