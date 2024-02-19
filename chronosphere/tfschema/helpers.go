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

package tfschema

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema/typeset"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/multierr"
)

// IsListEncodedObject returns whether s defines a list encoded object.
func IsListEncodedObject(s *schema.Schema) bool {
	_, elemIsObj := s.Elem.(*schema.Resource)
	return s.MaxItems == 1 && elemIsObj
}

// JSONNotificationPolicyDiffSuppress returns true if the diff between old and
// new notification policy JSON values should be suppressed, i.e. the resources
// are not considered different.
func JSONNotificationPolicyDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if old == new {
		return true
	}

	if old == "" || new == "" {
		return false
	}

	var oldPolicy map[string]any
	if err := json.Unmarshal([]byte(old), &oldPolicy); err != nil {
		return false
	}

	var newPolicy map[string]any
	if err := json.Unmarshal([]byte(new), &newPolicy); err != nil {
		return false
	}

	return reflect.DeepEqual(oldPolicy, newPolicy)
}

// ValidateNotificationPolicyData is a SchemaValidateFunc which tests if the
// provided value is of type string and is valid JSON. It does not validate
// that it is a serialization of a notification policy but we control that
// in the provider so we know what we're passing in is safe. If for some
// reason some valid but incorrect JSON does get set then this will
// fail on apply
//
// TODO: fix the code so that valid, but incorrect JSON for a notification policy fails on plan
func ValidateNotificationPolicyData(i any, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return warnings, errors
	}
	if v == IndependentNotificationPolicyData {
		return nil, nil
	}

	// make sure it is valid json
	var policy map[string]any
	if err := json.Unmarshal([]byte(v), &policy); err != nil {
		errors = append(errors, err)
	}

	return warnings, errors
}

func withDiffSuppress(n typeset.Normalizer, s *schema.Schema) *schema.Schema {
	if s.DiffSuppressFunc != nil {
		panic("DiffSuppressFunc is already set")
	}
	s.DiffSuppressFunc = func(_, old, new string, _ *schema.ResourceData) bool {
		return n.Normalize(old) == n.Normalize(new)
	}
	s.DiffSuppressOnRefresh = true
	return s
}

func float64RangeValidator(low, high float64) schema.SchemaValidateDiagFunc {
	return func(i any, _ cty.Path) diag.Diagnostics {
		f, ok := i.(float64)
		if !ok {
			return diag.Errorf("expected type to be float")
		}
		if f > high {
			return diag.Errorf("value must be within range [%f, %f]", low, high)
		}
		if f < low {
			return diag.Errorf("value must be within range [%f, %f]", low, high)
		}
		return nil
	}
}

func DiagError(diags diag.Diagnostics) error {
	if !diags.HasError() {
		return nil
	}

	errs := make([]error, 0, len(diags))
	for _, d := range diags {
		if d.Severity == diag.Error {
			errs = append(errs, fmt.Errorf("%s: %s", d.Summary, d.Detail))
		}
	}

	return multierr.Combine(errs...)
}
