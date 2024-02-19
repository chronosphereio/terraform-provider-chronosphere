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
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/prometheus/common/model"
)

// ParseDuration parses a time.Duration string based on the prom model.Duration
// type. This matches the backend parsing logic.
func ParseDuration(v string) (time.Duration, error) {
	result, err := model.ParseDuration(v)
	if err == nil {
		return time.Duration(result), nil
	}

	// If it fails then try to parse using time.ParseDuration
	return time.ParseDuration(v)
}

// FormatDuration formats a time.Duration based on the prom model.Duration type.
// This matches the backend parsing logic.
func FormatDuration(d time.Duration) string {
	return model.Duration(d).String()
}

// Duration defines the parameters of a duration field in a Terraform schema.
type Duration struct {
	Required bool
	Optional bool
}

// Schema returns the Terraform schema of the duration.
func (d Duration) Schema() *schema.Schema {
	return withDiffSuppress(d, &schema.Schema{
		Type:             schema.TypeString,
		ValidateDiagFunc: d.validate,
		Required:         d.Required,
		Optional:         d.Optional,
	})
}

// Normalize implements typeset.Normalizer.
func (d Duration) Normalize(v any) any {
	s := v.(string)

	if s == "" {
		return FormatDuration(0)
	}

	parsed, err := ParseDuration(s)
	if err != nil {
		return s
	}

	return FormatDuration(parsed)
}

func (d Duration) validate(v any, _ cty.Path) diag.Diagnostics {
	s, ok := v.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	if s == "" {
		if d.Required {
			return diag.Errorf("The argument must be a valid duration and not an empty string")
		}
		return nil
	}
	if _, err := ParseDuration(s); err != nil {
		return diag.Errorf("%q is not a valid duration: %v", v, err)
	}
	return nil
}
