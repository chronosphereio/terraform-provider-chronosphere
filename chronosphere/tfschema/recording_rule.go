package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var executionGroupFields = []string{"bucket_id", "execution_group"}

var RecordingRule = map[string]*schema.Schema{
	"bucket_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: executionGroupFields,
	},
	"execution_group": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: executionGroupFields,
	},
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
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"interval": Duration{
		Optional: true,
	}.Schema(),
	"expr": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_name": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
