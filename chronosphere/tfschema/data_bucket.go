package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var dataBucketOneOfAddressFields = []string{"name", "slug"} // fields used to address the bucket.

var DataBucket = map[string]*schema.Schema{
	"slug": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: dataBucketOneOfAddressFields,
	},
	"name": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: dataBucketOneOfAddressFields,
	},
	"description": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
}
