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

package enum

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type swaggerEnum interface {
	~string
}

// Enum defines a set of enum values, where a single logical value is
// has multiple accepted values: v1, an alias, and an optional legacy alias.
//
// Note that Enum will simply parrot back unknown values to ensure forwards
// compatibilty, for example if an unknown value was read from the server.
// However, we still rely on schema validation to prevent users from configuring
// unknown values in their Terraform files.
type Enum[V1 swaggerEnum] interface {
	// Name returns the type name of the enum.
	Name() string

	// V1 resolves into the enum's V1 value. If it is not
	// registered, it is simply propagated as an L type.
	V1(s string) V1

	// Validate implements schema.ValidateDiagFunc.
	Validate(v interface{}, _ cty.Path) diag.Diagnostics
}

// value defines the possible representations of a single enum value.
type value[V1 swaggerEnum] struct {
	// v1 is the configv1pb value. Required if isDefault=false.
	v1 V1

	// legacyAlias is a value for backwards compatibility.
	// Only use for types that predated the configv1 API.
	legacyAlias string

	// alias is the user-facing Terraform alias for the value. Required if
	// isDefault=false.
	alias string

	// isDefault indicates that this value is the default, in which case an
	// alias is optional (e.g. no alias needed for INVALID values). Must only be
	// set once per enum type.
	isDefault bool
}

type enum[V1 swaggerEnum] struct {
	name           string
	values         map[string]value[V1]
	displayAliases []string
}

func newEnum[V1 swaggerEnum](name string, values []value[V1]) enum[V1] {
	var displayAliases []string
	m := make(map[string]value[V1])
	defaultFound := false

	register := func(s string, v value[V1], requireUnique bool) {
		if s == "" && !v.isDefault {
			panic(fmt.Errorf("%q: missing value in %+v", name, v))
		}
		if vv, ok := m[s]; ok {
			if requireUnique {
				panic(fmt.Errorf("%q: duplicate registration: %q", name, s))
			} else if vv != v {
				panic(fmt.Errorf("%q: duplicate value across indexes: %q", name, s))
			}
		}
		m[s] = v
	}

	for _, v := range values {
		if v.isDefault {
			if defaultFound {
				panic(fmt.Errorf("%q: cannot configured two defaults", name))
			}
			if v.alias != "" {
				panic(fmt.Errorf("%q: default must use empty alias", name))
			}
			defaultFound = true
		}

		register(string(v.v1), v, true /* requireUnique */)
		if v.legacyAlias != "" {
			register(v.legacyAlias, v, true /* requireUnique */)
		}
		register(v.alias, v, false /* requireUnique */)

		if v.alias != "" {
			// Don't display empty aliases.
			displayAliases = append(displayAliases, v.alias)
		}
	}

	return enum[V1]{name, m, displayAliases}
}

func (e enum[V1]) Name() string {
	return e.name
}

func (e enum[V1]) V1(s string) V1 {
	v, ok := e.values[s]
	if !ok {
		return V1(s)
	}
	return v.v1
}

func (e enum[V1]) Validate(v interface{}, _ cty.Path) diag.Diagnostics {
	s := v.(string)
	if s == "" {
		return nil
	}
	if _, ok := e.values[s]; ok {
		return nil
	}
	var quotedAliases []string
	for _, a := range e.displayAliases {
		quotedAliases = append(quotedAliases, `"`+a+`"`)
	}
	return diag.Errorf(
		"%q is not a valid %s value; valid values: %s",
		s, e.name, strings.Join(quotedAliases, ", "))
}

func (e enum[V1]) ToStrings() Enum[string] {
	return stringAdaptor[V1]{e}
}

// stringAdaptor wraps any enum as a Enum[string, string] so that it can be used
// in non-generic situations.
type stringAdaptor[V1 swaggerEnum] struct {
	enum[V1]
}

func (a stringAdaptor[V1]) V1(s string) string {
	return string(a.enum.V1(s))
}

func (a stringAdaptor[V1]) Validate(v interface{}, p cty.Path) diag.Diagnostics {
	return a.enum.Validate(v, p)
}
