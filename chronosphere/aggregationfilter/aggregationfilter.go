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

package aggregationfilter

import (
	"fmt"
	"strings"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
)

const (
	DropRuleDelimiter      = "="
	MappingRuleDelimiter   = ":"
	RollupRuleDelimiter    = ":"
	ResourcePoolsDelimiter = ":"
)

// ListToModel converts a filter list to an API model.
func ListToModel(
	filter []string,
	kvDelimiter string,
) ([]*models.Configv1LabelFilter, error) {
	return sliceutil.MapErr(
		filter,
		func(s string) (*models.Configv1LabelFilter, error) {
			key, value, ok := strings.Cut(s, kvDelimiter)
			if !ok {
				return nil, fmt.Errorf("invalid filter %q: expected key%svalue", s, kvDelimiter)
			}
			return &models.Configv1LabelFilter{
				Name:      key,
				ValueGlob: value,
			}, nil
		})
}

// ListFromModel converts an API model to a filter list.
func ListFromModel(
	filter []*models.Configv1LabelFilter,
	kvDelimiter string,
) []string {
	return sliceutil.Map(
		filter,
		func(m *models.Configv1LabelFilter) string {
			return m.Name + kvDelimiter + m.ValueGlob
		})
}

// StringToModel converts a filter string to an API model.
func StringToModel(
	filter string,
	kvDelimiter string,
) ([]*models.Configv1LabelFilter, error) {
	if filter == "" {
		return nil, nil
	}
	return ListToModel(strings.Fields(filter), kvDelimiter)
}

// StringFromModel converts an API model to a filter string.
func StringFromModel(
	filter []*models.Configv1LabelFilter,
	kvDelimiter string,
) string {
	return strings.Join(ListFromModel(filter, kvDelimiter), " ")
}
