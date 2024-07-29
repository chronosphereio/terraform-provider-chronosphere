package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const maxAllocations = 10

var LogAllocationConfig = map[string]*schema.Schema{
	"default_dataset": {
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allocation": LogAllocationConfigSchema,
				"priorities": LogPrioritiesSchema,
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

var LogDatasetAllocationSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"dataset_slug": {
			Type:     schema.TypeString,
			Required: true,
		},
		"allocation": LogAllocationConfigSchema,
		"priorities": LogPrioritiesSchema,
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

var LogPrioritiesSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high_priority_filters": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: LogSearchFilterSchema},
				MinItems: 1,
				Optional: true,
			},
			"low_priority_filters": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: LogSearchFilterSchema},
				MinItems: 1,
				Optional: true,
			},
		},
	},
	MaxItems: 1,
	Optional: true,
}
