package chronosphere

import (
	"strings"

	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// CollectionTypeSlugToID encodes a collection type and slug into a single string
// for referencing collections of different types.
func CollectionTypeSlugToID(collType configmodels.Configv1CollectionReferenceType, slug string) string {
	return string(collType) + ":" + slug
}

// CollectionTypeSlugFromID converts an encoded ID containing type/slug into its' components.
// TODO: Should we validate type here, or rely on the swagger client validation? Client validation may not happen on Plan.
func CollectionTypeSlugFromID(id string) (collType configmodels.Configv1CollectionReferenceType, slug string, ok bool) {
	if t, s, ok := strings.Cut(id, ":"); ok {
		return configmodels.Configv1CollectionReferenceType(t), s, true
	}

	return "", id, false
}
