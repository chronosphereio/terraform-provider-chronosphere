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

// Package localid is used to create fake IDs used for Terraform resources not managed on the server.
package localid

import (
	"fmt"
	"math/rand"
	"regexp"
	"sync"
	"time"
)

var (
	// localPolicyIDRegex is the pattern for Terraform state IDs that were generated by this provider and not the server.
	// There are in the form of an INT, local-INT, or imported-INT: "12345", "local-12345", "imported-12345".
	// NB: while the latest provider uses "local-INT" format, older versions used just "INT".
	localPolicyIDRegex = regexp.MustCompile(`^((local|imported)-)?\d+$`)

	mu         sync.Mutex
	randSource = rand.NewSource(time.Now().UnixNano())
)

// IsLocalID returns whether an ID was generated by this package.
func IsLocalID(id string) bool {
	return localPolicyIDRegex.MatchString(id)
}

// NewLocalID generates a random local id.
func NewLocalID() string {
	return prefixedID("local")
}

// NewImportedID generates a random local id.
func NewImportedID() string {
	return prefixedID("imported")
}

func prefixedID(prefix string) string {
	mu.Lock()
	defer mu.Unlock()

	return fmt.Sprintf("%s-%d", prefix, randSource.Int63())
}