// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
