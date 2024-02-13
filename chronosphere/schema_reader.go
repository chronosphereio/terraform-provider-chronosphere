package chronosphere

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var (
	_ ResourceGetter = (*schema.ResourceData)(nil)
	_ ResourceGetter = (*schema.ResourceDiff)(nil)
)

// ResourceGetter is a subset of the read-only interface to read resource data
// from schema.ResourceData and schema.ResourceDiff.
type ResourceGetter interface {
	Get(key string) any
	Id() string
	GetOk(key string) (any, bool)
}
