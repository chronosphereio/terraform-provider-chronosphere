package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var LogSearchSchema = map[string]*schema.Schema{
	"query": {
		Type:     schema.TypeString,
		Required: true,
	},
}
