package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ConsumptionConfig = map[string]*schema.Schema{
	"partition": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Top-level partitions of the consumption tree. Partitions are non-overlapping; incoming requests are evaluated in order and assigned to the first matching partition. Requests matching no partition fall into an implicit `default` partition.",
		Elem:        ConsumptionConfigPartitionResource,
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
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Display name of the partition. Must be unique within its parent partition. Can be changed after creation.",
				},
				"slug": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Stable identifier of the partition. Must be unique within its parent partition. Immutable after creation.",
				},
				"filter": PartitionFilterSchema,
			},
		}
		if child != nil {
			r.Schema["partition"] = &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Child partitions of this partition. Evaluated in order; requests not matching any child fall into an implicit `default` child partition.",
				Elem:        child,
			}
		}
		return r
	})

var PartitionFilterSchema = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: "Filters identifying which data belongs to this partition. Filters are AND-ed together: a request must match every filter to be assigned to the partition. At most one `IN` filter and one `NOT_IN` filter can be specified.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"operator": Enum{
				Value:       enum.PartitionFilterOperator.ToStrings(),
				Optional:    true,
				Description: "Match operator (e.g. `IN`, `NOT_IN`) applied to the filter conditions.",
			}.Schema(),
			"condition": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Conditions evaluated by the filter. Each condition matches by dataset, logs, metrics, or trace data; exactly one of `log_filter`, `metric_filter`, or `dataset_id` must be set per condition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Deprecated: use `log_filter`, `metric_filter`, or trace filters instead. Slug of the dataset to match.",
						},
						"log_filter": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Log search filter matching log data for this condition.",
							Elem: &schema.Resource{
								Schema: LogSearchSchema,
							},
						},
						"metric_filter": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Metric label filters matched against incoming metric data. Multiple filters are AND-ed together; values support glob patterns including `service:{svc1,svc2}` style alternations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Label name to match.",
									},
									"value_glob": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Glob pattern matched against the label's value.",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
