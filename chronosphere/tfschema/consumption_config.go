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
				"dataset_filter": DatasetFilterSchema,
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

var DatasetFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"operator": Enum{
				Value:    enum.DatasetFilterOperator.ToStrings(),
				Optional: true,
			}.Schema(),
			"dataset": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	},
}
