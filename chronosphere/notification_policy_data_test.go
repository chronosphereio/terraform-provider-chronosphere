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
