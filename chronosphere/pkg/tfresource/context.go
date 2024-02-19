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

package tfresource

import (
	"context"
)

type tfResourceContextKey struct{}

// NewContext creates a new context with the TF resource name set
func NewContext(ctx context.Context, resourceName string) context.Context {
	return context.WithValue(ctx, tfResourceContextKey{}, resourceName)
}

// FromContext returns the resource name from the context
func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(tfResourceContextKey{}).(string)
	if !ok {
		return "", false
	}
	return v, true
}
