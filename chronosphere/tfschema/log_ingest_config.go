package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const maxLogParsers = 10

var LogIngestConfig = map[string]*schema.Schema{
	"parser": {
		Type:     schema.TypeList,
		Elem:     LogParserSchema,
		Optional: true,
		MaxItems: maxLogParsers,
	},
}

var LogParserSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"regex": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}
