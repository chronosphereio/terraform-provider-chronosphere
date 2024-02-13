package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const maxResourcePools = 128

var ResourcePoolsConfig = map[string]*schema.Schema{
	"default_pool": {
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allocation": ResourcePoolAllocationSchema,
				"priorities": ResourcePoolPrioritiesSchema,
			},
			SchemaVersion: 1,
		},
		Required: true,
		MaxItems: 1,
	},
	"pools": {
		Type:          schema.TypeList,
		Elem:          ResourcePoolElemSchema,
		Optional:      true,
		ConflictsWith: []string{"pool"},
		Deprecated:    "Use pool instead of pools",
		MaxItems:      maxResourcePools,
	},
	"pool": {
		Type:          schema.TypeList,
		Elem:          ResourcePoolElemSchema,
		ConflictsWith: []string{"pools"},
		Optional:      true,
		MaxItems:      maxResourcePools,
	},
}

var ResourcePoolAllocationSchema = &schema.Schema{
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

var ResourcePoolElemSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"allocation": ResourcePoolAllocationSchema,
		"match_rule": {
			Type:       schema.TypeString,
			Optional:   true,
			Deprecated: "use match_rules",
		},
		"match_rules": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			MinItems: 1,
			Optional: true,
		},
		"priorities": ResourcePoolPrioritiesSchema,
	},
}

var ResourcePoolPrioritiesSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high_priority_match_rules": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				Optional: true,
			},
			"low_priority_match_rules": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				Optional: true,
			},
		},
	},
	MaxItems: 1,
	Optional: true,
}
