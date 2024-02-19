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

// OptionalValue is used for fields that are marked as optional, where Terraform differentiates
// between null and a default value, but the server does not.
type OptionalValue interface {
	Value() any
	IsEmpty() bool
}

// OptionalString implements OptionalValue for strings.
type OptionalString string

func (s OptionalString) Value() any    { return s }
func (s OptionalString) IsEmpty() bool { return s == "" }

// OptionalMap implements OptionalValue for maps.
type OptionalMap[T comparable, V any] map[T]V

func (m OptionalMap[T, V]) Value() any    { return m }
func (m OptionalMap[T, V]) IsEmpty() bool { return len(m) == 0 }
