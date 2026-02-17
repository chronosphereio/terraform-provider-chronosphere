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
	// instead of "global/gateway/dev"), which would be super unreadable
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
		Type:     schema.TypeString,
		Optional: true,
	},
	"priority": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: consumptionBudgetPrioritySchema,
		},
	},
	"threshold": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: consumptionBudgetThresholdSchema,
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
	"alert_action_config": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: consumptionBudgetAlertActionConfigSchema,
		},
	},
}

var consumptionBudgetAlertActionConfigSchema = map[string]*schema.Schema{
	"annotations": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"instant_rate_sustain_secs": {
		Type:     schema.TypeInt,
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

var consumptionBudgetThresholdSchema = map[string]*schema.Schema{
	"action": Enum{
		Value:    enum.ConsumptionBudgetThresholdAction.ToStrings(),
		Optional: true,
	}.Schema(),
	"type": Enum{
		Value:    enum.ConsumptionBudgetThresholdType.ToStrings(),
		Optional: true,
	}.Schema(),
	"instant_rate": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: consumptionBudgetInstantRateSchema,
		},
	},
	"volume": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: consumptionBudgetVolumeSchema,
		},
	},
	// TODO: add sku_group and unit enums. the action field is a good
	// example to follow.
	// hint:
	// 1. add to enum package
	// 2. add bindings here in schema
	// 3. `make generate-no-swagger` to generate the intschema
	// 4. map to/from API models in resource_consumption_budget.go
}

var consumptionBudgetInstantRateSchema = map[string]*schema.Schema{
	"fixed_value_per_sec": {
		Type:     schema.TypeInt,
		Optional: true,
	},
}

var consumptionBudgetVolumeSchema = map[string]*schema.Schema{
	"fixed_value": {
		Type:     schema.TypeInt,
		Optional: true,
	},
}
