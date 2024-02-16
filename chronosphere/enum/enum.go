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

package enum

import (
	"fmt"
	"strings"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type swaggerEnum interface {
	~string
}

// Enum defines a set of enum values, where a single logical value is
// represented by multiple acceptable strings (legacy, v1, and an alias).
//
// Note that Enum will simply parrot back unknown values to ensure forwards
// compatibilty, for example if an unknown value was read from the server.
// However, we still rely on schema validation to prevent users from configuring
// unknown values in their Terraform files.
type Enum[L, V1 swaggerEnum] interface {
	// Name returns the type name of the enum.
	Name() string

	// Legacy resolves into the enum's legacy value. If it is not registered
	// it is simply propagated as an L type.
	Legacy(s string) L

	// V1 resolves into the enum's V1 value. If it is not
	// registered, it is simply propagated as an L type.
	V1(s string) V1

	// Validate implements schema.ValidateDiagFunc.
	Validate(v interface{}, _ cty.Path) diag.Diagnostics
}

// value defines the possible representations of a single enum value.
//
// In general, you should NOT be using when introducing new enums. Legacy
// mappings are only needed for backwards compatibility for old enums.
type value[L, V1 swaggerEnum] struct {
	// legacy is the apipb value. Required if isDefault=false.
	legacy L

	// v1 is the configv1pb value. Required if isDefault=false.
	v1 V1

	// alias is the user-facing Terraform alias for the value. Required if
	// isDefault=false.
	alias string

	// isDefault indicates that this value is the default, in which case an
	// alias is optional (e.g. no alias needed for INVALID values). Must only be
	// set once per enum type.
	isDefault bool
}

// v1OnlyValue is a subset of value which configures enums which don't require
// legacy support.
//
// In general, you should be using this when introducing new enums, since
// Terraform doesn't use the legacy API anymore.
type v1OnlyValue[V1 swaggerEnum] struct {
	v1        V1
	alias     string
	isDefault bool
}

type enum[L, V1 swaggerEnum] struct {
	name           string
	values         map[string]value[L, V1]
	displayAliases []string
}

func newV1OnlyEnum[V1 swaggerEnum](name string, v1Values []v1OnlyValue[V1]) enum[V1, V1] {
	return newEnum(name, sliceutil.Map(v1Values, func(v v1OnlyValue[V1]) value[V1, V1] {
		return value[V1, V1]{
			legacy:    v.v1,
			v1:        v.v1,
			alias:     v.alias,
			isDefault: v.isDefault,
		}
	}))
}

func newEnum[L, V1 swaggerEnum](name string, values []value[L, V1]) enum[L, V1] {
	var displayAliases []string
	m := make(map[string]value[L, V1])
	defaultFound := false

	register := func(s string, v value[L, V1]) {
		if s == "" && !v.isDefault {
			panic(fmt.Errorf("%q: missing value in %+v", name, v))
		}
		if vv, ok := m[s]; ok && vv != v {
			panic(fmt.Errorf("%q: duplicate value across indexes: %q", name, s))
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

		// legacy is optional now that we use public api everywhere.
		if v.legacy != "" {
			register(string(v.legacy), v)
		}
		register(string(v.v1), v)
		register(v.alias, v)

		if v.alias != "" {
			// Don't display empty aliases.
			displayAliases = append(displayAliases, v.alias)
		}
	}

	return enum[L, V1]{name, m, displayAliases}
}

func (e enum[L, V1]) Name() string {
	return e.name
}

func (e enum[L, V1]) Legacy(s string) L {
	v, ok := e.values[s]
	if !ok {
		return L(s)
	}
	if v.legacy == "" {
		return L(s)
	}
	return v.legacy
}

func (e enum[L, V1]) V1(s string) V1 {
	v, ok := e.values[s]
	if !ok {
		return V1(s)
	}
	return v.v1
}

func (e enum[L, V1]) Validate(v interface{}, _ cty.Path) diag.Diagnostics {
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

func (e enum[L, V1]) ToStrings() Enum[string, string] {
	return stringAdaptor[L, V1]{e}
}

// stringAdaptor wraps any enum as a Enum[string, string] so that it can be used
// in non-generic situations.
type stringAdaptor[L, V1 swaggerEnum] struct {
	enum[L, V1]
}

func (a stringAdaptor[L, V1]) Legacy(s string) string {
	return string(a.enum.Legacy(s))
}

func (a stringAdaptor[L, V1]) V1(s string) string {
	return string(a.enum.V1(s))
}

func (a stringAdaptor[L, V1]) Validate(v interface{}, p cty.Path) diag.Diagnostics {
	return a.enum.Validate(v, p)
}
