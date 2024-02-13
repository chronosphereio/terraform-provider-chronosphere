package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Enum defines the parameters of an enum field in a Terraform schema.
type Enum struct {
	Value    enum.Enum[string, string]
	Required bool
	Optional bool
	ForceNew bool
}

// Schema returns the Terraform of the enum.
func (e Enum) Schema() *schema.Schema {
	return withDiffSuppress(e, &schema.Schema{
		Type:             schema.TypeString,
		Required:         e.Required,
		Optional:         e.Optional,
		ForceNew:         e.ForceNew,
		ValidateDiagFunc: e.Value.Validate,
	})
}

// Normalize implements typeset.Normalizer.
func (e Enum) Normalize(v any) any {
	return e.Value.V1(v.(string))
}
