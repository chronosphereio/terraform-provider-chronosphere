package chronosphere

import (
	"testing"

	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/stretchr/testify/assert"
)

func TestCollectionTypeSlugToID(t *testing.T) {
	assert.Equal(t, "SIMPLE:bar",
		CollectionTypeSlugToID(configmodels.Configv1CollectionReferenceTypeSIMPLE, "bar"))
	assert.Equal(t, "SERVICE:foo",
		CollectionTypeSlugToID(configmodels.Configv1CollectionReferenceTypeSERVICE, "foo"))
}

func TestCollectionTypeSlugFromID(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		wantType configmodels.Configv1CollectionReferenceType
		wantSlug string
		wantOK   bool
	}{
		{
			name:     "only slug",
			id:       "foo",
			wantSlug: "foo",
			wantOK:   false,
		},
		{
			name:     "simple collection",
			id:       "SIMPLE:foo",
			wantType: configmodels.Configv1CollectionReferenceTypeSIMPLE,
			wantSlug: "foo",
			wantOK:   true,
		},
		{
			name:     "service collection",
			id:       "SERVICE:foo",
			wantType: configmodels.Configv1CollectionReferenceTypeSERVICE,
			wantSlug: "foo",
			wantOK:   true,
		},
		{
			name:     "unknown type",
			id:       "BAZ:foo",
			wantType: "BAZ", // validation is left to the swagger client.
			wantSlug: "foo",
			wantOK:   true,
		},
		{
			name:     "multiple separators",
			id:       "SIMPLE:foo:bar",
			wantType: configmodels.Configv1CollectionReferenceTypeSIMPLE,
			wantSlug: "foo:bar",
			wantOK:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotSlug, gotOK := CollectionTypeSlugFromID(tt.id)
			assert.Equal(t, tt.wantOK, gotOK, "ok mismatch")
			assert.Equal(t, tt.wantType, gotType, "type mismatch")
			assert.Equal(t, tt.wantSlug, gotSlug, "slug mismatch")
		})
	}
}
