package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var LogSearchFilterSchema = map[string]*schema.Schema{
	"query": {
		Type:     schema.TypeString,
		Required: true,
	},
}
