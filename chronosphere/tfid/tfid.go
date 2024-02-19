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

package tfid

import "regexp"

type Type int

const (
	// TypeEmpty indicates an uninitialized ID.
	TypeEmpty Type = iota

	// TypeSlug indicates a slug ID.
	TypeSlug

	// TypeLocalRef indicates a local reference ID.
	TypeLocalRef
)

// TF names support unicode characters as well:
// https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md#identifiers
// However, Chronosphere slugs do not.
var (
	_tfNameReplaceRegex = regexp.MustCompile("[^a-zA-Z0-9_]+")
	_tfStartNotSafe     = regexp.MustCompile("^[^a-zA-Z_]")
)

// ID represents a Terraform ID field. It is a mutually exclusive one-of type of
// either slug, e.g. "my-bucket", or local Terraform reference, e.g.
// "chronosphere_bucket.my_bucket.id".
//
// The core Terraform resource providers should always use slugs. Use-cases
// which generate HCL (e.g. scenario tests and export config) may use local
// references.
type ID struct {
	t        Type
	slug     string
	localRef Ref
}

// Slug unwraps id as a slug, else panics.
func (id ID) Slug() string {
	switch id.t {
	case TypeEmpty:
		return ""
	case TypeSlug:
		return id.slug
	default:
		panic("ID is not a slug")
	}
}

// LocalRef unwraps id as a local reference, else panics.
func (id ID) LocalRef() Ref {
	switch id.t {
	case TypeEmpty:
		return Ref{}
	case TypeLocalRef:
		return id.localRef
	default:
		panic("ID is not a local ref")
	}
}

// Type returns the type of the id.
func (id ID) Type() Type {
	return id.t
}

// Slug returns a slug ID.
func Slug(slug string) ID {
	if slug == "" {
		return ID{}
	}
	return ID{
		t:    TypeSlug,
		slug: slug,
	}
}

// LocalRef returns a local reference ID.
func LocalRef(r Ref) ID {
	if r == (Ref{}) {
		return ID{}
	}
	return ID{
		t:        TypeLocalRef,
		localRef: r,
	}
}

// Ref contains metadata of a local reference.
type Ref struct {
	// True if the address is for a datasource and not a resource, i.e. should
	// there be a "data" prefix.
	Datasource bool

	// Resource type, required. E.g. chronosphere_bucket
	Type string

	// Resource id, required. E.g. "foo" for a chronosphere_bucket.foo address.
	ID string

	// Field name, optional, defaults to "id". E.g. "bar" for
	// chronosphere_bucket.things.bar
	Field string
}

// AsID is an ergonomic helper for LocalRef.
func (r Ref) AsID() ID {
	return LocalRef(r)
}

// SafeID converts a slug into a terraform-safe name.
// This is a stateless approach that replaces unsafe characters, and can result
// in duplicate terraform IDs if there's multiple similar slugs.
// E.g., foo-bar and foo_bar would both become foo_bar.
func SafeID(slug string) string {
	// Terraform naming restrictions:
	// > A name must start with a letter or underscore and may contain only letters, digits,
	// > underscores, and dashes
	// Terraform best practices (https://www.terraform-best-practices.com/naming) suggest
	// using `_` instead of `-` in all names, so we replace any non-alphanumeric characters
	// with `_`.
	name := _tfNameReplaceRegex.ReplaceAllString(slug, "_")

	// Some characters that are valid in the name are not valid at the start (e.g., numbers).
	// Add an _ prefix for those names here.
	if _tfStartNotSafe.MatchString(name) {
		name = "_" + name
	}
	return name
}
