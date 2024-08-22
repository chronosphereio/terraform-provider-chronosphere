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
		"dataset_id": {
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
			"high_priority_filter": LogSearchFilterSchema,
			"low_priority_filter":  LogSearchFilterSchema,
		},
	},
	MaxItems: 1,
	Optional: true,
}

var LogSearchFilterSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: LogSearchSchema,
	},
	MinItems: 1,
	Optional: true,
}
