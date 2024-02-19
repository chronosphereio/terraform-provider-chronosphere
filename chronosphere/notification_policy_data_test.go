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

import (
	"encoding/json"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/ptr"
	"github.com/stretchr/testify/require"
)

func TestNotificationPolicyData(t *testing.T) {
	npd := &NotificationPolicyData{
		Routes: &Routes{
			Defaults: map[string]RoutesNotifierList{
				"critical": {
					Notifiers: []*RoutesNotifierListNotifier{
						{
							Name: "test-notifier",
							Slug: "test-notifier-slug",
						},
					},
				},
				"warn": {
					Notifiers: []*RoutesNotifierListNotifier{
						{
							Name: "test-other-notifier",
							Slug: "test-other-notifier-slug",
						},
					},
				},
			},
			Overrides: []*RoutesOverride{
				{
					AlertLabelMatchers: []*AlertLabelMatcher{
						{
							Name:  "test-matcher",
							Type:  ptr.Wrap(ExactMatcherType),
							Value: "test-matcher-value",
						},
					},
					Notifiers: map[string]RoutesNotifierList{
						"critical": {
							Notifiers: []*RoutesNotifierListNotifier{
								{
									Name: "test-third-notifier",
									Slug: "test-third-notifier-slug",
								},
							},
						},
					},
				},
			},
		},
	}

	npdJSON, err := json.Marshal(npd)
	require.NoError(t, err)

	unmarshalledNpd, err := unmarshalNotificationPolicyData(string(npdJSON))
	require.NoError(t, err)

	require.Equal(t, npd, unmarshalledNpd)
}
