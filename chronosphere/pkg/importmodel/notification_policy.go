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

package importmodel

import (
	"encoding/base64"
	"encoding/json"
)

// Bucket encapsulates all the data needed to import a bucket using `terraform import`
type Bucket struct {
	Slug                   string `json:"slug"`
	NotificationPolicySlug string `json:"notification_policy_slug"`
}

// ParseBucket parses a serialized bucket passed to `terraform import`.
func ParseBucket(importID string) (Bucket, error) {
	var bucket Bucket
	b, err := base64.StdEncoding.DecodeString(importID)
	if err != nil {
		return Bucket{}, err
	}
	if err := json.Unmarshal(b, &bucket); err != nil {
		return Bucket{}, err
	}
	return bucket, nil
}
