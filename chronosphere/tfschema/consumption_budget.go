package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var ConsumptionBudget = map[string]*schema.Schema{
	// This isn't an actual field in the API; the true reference is the
	// partition_name_path, and even though Terraform supports resource
	// attribute references, we choose not to use that feature because the
	// partition references would be by index (e.g. partitions[1].partitions[2]
	// instead of ["global", "gateway", "dev"]), which would be super unreadable
	// & fragile.
	//
	// So instead, we opt for a coarse reference which acts as a "depends_on"
	// pointer. We validate that the value is the singleton ID of the
	// ConsumptionConfig resource.
	"consumption_config_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"resource": Enum{
		Value:    enum.ConsumptionBudgetResource.ToStrings(),
		Optional: true,
	}.Schema(),
	"partition_slug_path": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"priority": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: consumptionBudgetPrioritySchema,
		},
	},
	"behavior": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: consumptionBudgetBehaviorSchema,
		},
	},
	"default_priority": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"notification_policy_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

var consumptionBudgetPrioritySchema = map[string]*schema.Schema{
	"filter": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dataset_id": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"log_filter": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: LogSearchSchema,
					},
				},
			},
		},
	},
	"priority": {
		Type:     schema.TypeInt,
		Optional: true,
	},
}

var consumptionBudgetBehaviorSchema = map[string]*schema.Schema{
	"action": Enum{
		Value:    enum.BehaviorAction.ToStrings(),
		Optional: true,
	}.Schema(),
	"threshold_type": Enum{
		Value:    enum.BehaviorThresholdType.ToStrings(),
		Optional: true,
	}.Schema(),
	"instant_rate_threshold": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: consumptionBudgetInstantRateThresholdSchema,
		},
	},
	"volume_threshold": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: consumptionBudgetVolumeThresholdSchema,
		},
	},
}

var consumptionBudgetInstantRateThresholdSchema = map[string]*schema.Schema{
	"fixed_value_per_sec": {
		Type:     schema.TypeInt,
		Optional: true,
	},
}

var consumptionBudgetVolumeThresholdSchema = map[string]*schema.Schema{
	"fixed_value": {
		Type:     schema.TypeInt,
		Optional: true,
	},
}
