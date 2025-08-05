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
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/maps"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

const (
	// severityLabelName is the name of the label that is used to indicate the severity
	// level of an alert.
	severityLabelName = "severity"
)

// setKey sets a key on the given ResourceData. If that set fails, an error is appended to the
// returned diagnostics set.
func setKey(errors diag.Diagnostics, rd *schema.ResourceData, key string, value any) diag.Diagnostics {
	if err := rd.Set(key, value); err != nil {
		return append(errors, diag.FromErr(err)[0])
	}
	return errors
}

func diagError(diags diag.Diagnostics) error {
	return tfschema.DiagError(diags)
}

func setResourceNotFound(d *schema.ResourceData) {
	// From https://learn.hashicorp.com/tutorials/terraform/provider-setup:
	// If the resource isn't available, the function should set the ID to an empty string so Terraform "destroys" the resource in state.
	d.SetId("")
}

func BoolPtr(b bool) *bool {
	return &b
}

// ParseOptionalDuration is similar to ParseDuration, but treats empty
// string as 0.
func ParseOptionalDuration(v string) (time.Duration, error) {
	if v == "" {
		return 0, nil
	}
	return tfschema.ParseDuration(v)
}

// isRawAttributeNull returns a bool indicating whether the given attribute is set in the
// raw given resource data.  This is helpful in particular in dry run validation to differentiate
// between
// a) the user is setting the value to a local resource id which doesn't yet exist
// b) the user is not setting the value at all
func isRawAttributeNull(diff *schema.ResourceDiff, attr string) bool {
	return diff.GetRawConfig().GetAttr(attr).IsNull()
}

func durationToSecs(rawDur string) (int32, error) {
	var d time.Duration
	if rawDur != "" {
		var err error
		d, err = tfschema.ParseDuration(rawDur)
		if err != nil {
			return 0, err
		}
	}
	if d.Round(time.Second) != d {
		return 0, fmt.Errorf("invalid duration %q: must use seconds granularity", rawDur)
	}
	return int32(d.Seconds()), nil
}

func durationFromSecs(secs int32) string {
	if secs == 0 {
		return ""
	}
	return tfschema.FormatDuration(time.Duration(secs) * time.Second)
}

func matchersToModel(matchers []intschema.Matcher) []*models.Configv1LabelMatcher {
	var out []*models.Configv1LabelMatcher
	for _, m := range matchers {
		out = append(out, &models.Configv1LabelMatcher{
			Name:  m.Name,
			Type:  enum.MatcherType.V1(m.Type),
			Value: m.Value,
		})
	}
	return out
}

func matchersFromModel(matchers []*models.Configv1LabelMatcher) []intschema.Matcher {
	var out []intschema.Matcher
	for _, m := range matchers {
		out = append(out, intschema.Matcher{
			Name:  m.Name,
			Type:  string(m.Type),
			Value: m.Value,
		})
	}
	return out
}

func sortedKeys[V any](m map[string]V) []string {
	keys := maps.Keys(m)
	sort.Strings(keys)
	return keys
}

func skipDryRun(diff *schema.ResourceDiff) bool {
	// Get all updated attributes, including nested attributes.
	updatedKeys := diff.GetChangedKeysPrefix("")

	// Empty means unchanged diff (i.e. no-op).
	return len(updatedKeys) == 0
}

// notifierTypeChangedName appends a suffix to append to a name to indicate that a notifier
// type has changed on the server. This will indicate to the user that Terraform needs to
// update the entity, which will fix the incorrect type. See internal ticket 28252 for
// more details
func notifierTypeChangedName(expectedNotifierType string) string {
	return " << Notifier type is no longer " + expectedNotifierType + ", apply to fix >>"
}

func collectionRefFromID(id string) (slug string, ref *models.Configv1CollectionReference) {
	collType, slug, ok := CollectionTypeSlugFromID(id)
	if !ok {
		return id, nil
	}

	return "", &models.Configv1CollectionReference{
		Type: collType,
		Slug: slug,
	}
}

func collectionIDFromRef(slug string, ref *models.Configv1CollectionReference) string {
	if ref == nil {
		return slug
	}

	return CollectionTypeSlugToID(ref.Type, ref.Slug)
}

// parseStringToInt64 parses a string to int64 with proper error handling.
// Returns 0 if the string is empty, otherwise parses the string as base 10.
func parseStringToInt64(s string, fieldName string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s %q: %w", fieldName, s, err)
	}

	return value, nil
}
