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
		Required:    true,
		MaxItems:    1,
		Description: "Allocation and priority configuration for the default dataset, which receives any logs not matched by a `dataset_allocation` entry.",
	},
	"dataset_allocation": {
		Type:        schema.TypeList,
		Elem:        LogDatasetAllocationSchema,
		Optional:    true,
		MaxItems:    maxAllocations,
		Description: "Per-dataset allocation and priority overrides. Datasets are evaluated in order; the first match wins.",
	},
}

var LogDatasetAllocationSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"dataset_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Slug of the dataset this allocation applies to.",
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
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Percentage of the tenant's log license to allocate to this dataset, expressed as a number between 0 and 100.",
			},
		},
	},
	MaxItems:    1,
	Required:    true,
	Description: "Resource allocation for the dataset, expressed as a share of the overall log license.",
}

var LogPrioritiesSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high_priority_filter": LogSearchFilterSchema,
			"low_priority_filter":  LogSearchFilterSchema,
		},
	},
	MaxItems:    1,
	Optional:    true,
	Description: "Defines high and low priority match criteria. Low priority logs are dropped first when the allocation is exhausted, then default priority, with high priority dropped last.",
}

var LogSearchFilterSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: LogSearchSchema,
	},
	MinItems:    1,
	Optional:    true,
	Description: "List of log search filters. Filters are combined as OR statements so only one filter needs to match.",
}
