package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const maxAllocations = 10

var LogAllocationConfig = map[string]*schema.Schema{
	"default_dataset": {
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allocation": LogAllocationConfigSchema,
			},
			SchemaVersion: 1,
		},
		Required: true,
		MaxItems: 1,
	},
	"dataset_allocation": {
		Type:     schema.TypeList,
		Elem:     LogDatasetAllocationSchema,
		Optional: true,
		MaxItems: maxAllocations,
	},
}

var LogAllocationConfigSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"percent_of_license": {
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	},
	MaxItems: 1,
	Required: true,
}

var LogDatasetAllocationSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"dataset_slug": {
			Type:     schema.TypeString,
			Required: true,
		},
		"allocation": LogAllocationConfigSchema,
	},
}
