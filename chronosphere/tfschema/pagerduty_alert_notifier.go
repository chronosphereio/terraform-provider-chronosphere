package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var PagerdutyAlertNotifier = map[string]*schema.Schema{
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
	"class": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"client": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"client_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"component": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"details": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"group": {
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
	"image": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"alt": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"href": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"src": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
	"link": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"href": {
					Type:     schema.TypeString,
					Required: true,
				},
				"text": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	},
	"routing_key": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"service_key": {
		Type:      schema.TypeString,
		Optional:  true,
		Sensitive: true,
	},
	"severity": {
		Type:     schema.TypeString,
		Required: true,
	},
	"url": {
		Type:     schema.TypeString,
		Required: true,
	},
}
