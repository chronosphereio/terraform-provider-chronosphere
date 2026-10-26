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

package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// Canonical re-encodes a JSON document with sorted object keys and no
// insignificant whitespace, so two encodings of the same document compare
// equal as strings. Numbers keep their original literal, so integers above
// 2^53 survive the round trip; this is why it is used in place of
// terraform-plugin-sdk's structure.SuppressJsonDiff, which decodes through
// float64. The whole input must be one JSON value: trailing data is an error.
func Canonical(data string) (string, error) {
	dec := json.NewDecoder(bytes.NewReader([]byte(data)))
	dec.UseNumber()

	var v any
	if err := dec.Decode(&v); err != nil {
		return "", err
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return "", errors.New("trailing data after JSON value")
	}

	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
