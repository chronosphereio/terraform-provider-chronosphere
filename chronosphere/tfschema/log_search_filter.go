package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var LogSearchSchema = map[string]*schema.Schema{
	"query": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Log search query that selects matching logs. Supports only top-level operations; nested clauses are not allowed and only one type of `AND` or `OR` operator may be used.",
	},
}
