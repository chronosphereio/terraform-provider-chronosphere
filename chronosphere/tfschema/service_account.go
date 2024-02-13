package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var serviceAccountPermOneOfFields = []string{"restriction", "unrestricted"}

var ServiceAccount = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"email": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"token": {
		Type:      schema.TypeString,
		Computed:  true,
		Sensitive: true,
	},
	"unrestricted": {
		Type:         schema.TypeBool,
		Optional:     true,
		ForceNew:     true,
		ExactlyOneOf: serviceAccountPermOneOfFields,
	},
	"restriction": {
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"permission": Enum{
					Value:    enum.Permission.ToStrings(),
					Required: true,
				}.Schema(),
				"labels": {
					Type:     schema.TypeMap,
					Elem:     &schema.Schema{Type: schema.TypeString},
					Optional: true,
				},
			},
		},
		ExactlyOneOf: serviceAccountPermOneOfFields,
	},
}
