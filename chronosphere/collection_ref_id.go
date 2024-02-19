// Copyright 2024 Chronosphere Inc.
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
