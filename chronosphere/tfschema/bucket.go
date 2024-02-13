package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var Bucket = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	// notification_policy_slug is an internal field used to track the slug of inline policies
	// set via notification_policy_data.
	// Users who want to specify a default policy should use notification_policy_id.
	"notification_policy_slug": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"notification_policy_id": {
		Type:          schema.TypeString,
		ConflictsWith: []string{"notification_policy_data"},
		Optional:      true,
	},
	"team_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"notification_policy_data": {
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: JSONNotificationPolicyDiffSuppress,
		ValidateFunc:     ValidateNotificationPolicyData,
	},
}
