package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var OpsgenieAlertNotifier = map[string]*schema.Schema{
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
	"message": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"source": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"details": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"responder": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"username": {
					Type:     schema.TypeString,
					Optional: true,
				},

				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
	"tags": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
	},
	"note": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"priority": {
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
