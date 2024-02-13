package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var MappingRule = map[string]*schema.Schema{
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
	"bucket_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"filter": Filter{
		KVDelimiter: aggregationfilter.MappingRuleDelimiter,
	}.Schema(),
	"aggregations": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: Enum{
			Value: enum.AggregationType.ToStrings(),
		}.Schema(),
	},
	// Storage policies to apply to the mapped metrics.
	"storage_policy": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"resolution": Duration{
					Required: true,
				}.Schema(),
				"retention": Duration{
					Required: true,
				}.Schema(),
			},
		},
	},
	"drop": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	// Whether or not to drop the timestamp.
	"drop_timestamp": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"interval": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"mode": Enum{
		Value:    enum.MappingModeType.ToStrings(),
		Optional: true,
	}.Schema(),
}
