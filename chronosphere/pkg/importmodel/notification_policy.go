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
