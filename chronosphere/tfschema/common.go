package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var MatcherListSchema = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	MinItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": Enum{
				Value:    enum.MatcherType.ToStrings(),
				Required: true,
			}.Schema(),
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	},
}
