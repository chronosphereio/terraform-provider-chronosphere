package chronosphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceNotificationPolicy creates a schema for a notification policy data source.
func dataSourceNotificationPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceInlineNotificationPolicyRead,

		Schema: resourceNotificationPolicy().Schema,
	}
}
