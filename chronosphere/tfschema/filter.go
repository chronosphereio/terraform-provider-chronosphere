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

package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Filter represents a raw string glob label filter, where label-value pairs are
// delimited by the given KVDelimiter. E.g. "__name__:foo instace:service*"
// (where KVDelimiter=":").
//
// This requires a normalized schema because the original whitespace is not
// persisted in the server-side database (e.g. a filter with multiple
// consecutive spaces would always produce a diff).
type Filter struct {
	KVDelimiter string
}

func (f Filter) Schema() *schema.Schema {
	if f.KVDelimiter == "" {
		panic("KVDelimiter must be set")
	}
	return withDiffSuppress(f, &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	})
}

func (f Filter) Normalize(v any) any {
	s := v.(string)
	if s == "" {
		return ""
	}
	// Convert to the model and back to normalize any whitespace.
	m, err := aggregationfilter.StringToModel(s, f.KVDelimiter)
	if err != nil {
		return s
	}
	return aggregationfilter.StringFromModel(m, f.KVDelimiter)
}
