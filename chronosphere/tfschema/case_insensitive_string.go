package tfschema

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CaseInsensitiveString defines the parameters of a case-insensitive string
// field in a Terraform schema.
type CaseInsensitiveString struct {
	Required bool
}

// Schema returns the Terraform schema of the string.
func (s CaseInsensitiveString) Schema() *schema.Schema {
	return withDiffSuppress(s, &schema.Schema{
		Type:     schema.TypeString,
		Required: s.Required,
		Optional: !s.Required,
	})
}

// Normalize implements typeset.Normalizer.
func (s CaseInsensitiveString) Normalize(v any) any {
	return strings.ToLower(v.(string))
}
