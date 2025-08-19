package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ConsumptionConfig = map[string]*schema.Schema{
	"partition": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     ConsumptionConfigPartitionResource,
	},
}

func makeRecursiveResource(
	depth int,
	spawn func(child *schema.Resource) *schema.Resource,
) *schema.Resource {
	if depth == 0 {
		return spawn(nil)
	}
	return spawn(makeRecursiveResource(depth-1, spawn))
}

var ConsumptionConfigPartitionResource = makeRecursiveResource(
	5, /* depth */
	func(child *schema.Resource) *schema.Resource {
		r := &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"slug": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"filter": PartitionFilterSchema,
			},
		}
		if child != nil {
			r.Schema["partition"] = &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     child,
			}
		}
		return r
	})

var PartitionFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"operator": Enum{
				Value:    enum.PartitionFilterOperator.ToStrings(),
				Optional: true,
			}.Schema(),
			"condition": {
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
		},
	},
}
