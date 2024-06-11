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
	"fmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/notification_policy"
	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/ptr"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
)

// LabelMatcherType represents matcher types for labels.
type LabelMatcherType string

// EntitiesMatcherType enum values are as follows:
const (
	InvalidMatcherType LabelMatcherType = "INVALID_MATCHER_TYPE"
	ExactMatcherType   LabelMatcherType = "EXACT_MATCHER_TYPE"
	RegexpMatcherType  LabelMatcherType = "REGEXP_MATCHER_TYPE"
)

// NotificationPolicyData stores policy information for inline policies, which is used by the bucket
// resource to create and manage the policy.
// It matches the previously marshalled values generated from legacy API models.
type NotificationPolicyData struct {
	Routes *Routes `json:"routes,omitempty"`
}

// Routes represents the routes of a notification policy
type Routes struct {
	Defaults  map[string]RoutesNotifierList `json:"defaults,omitempty"`
	Overrides []*RoutesOverride             `json:"overrides"`
}

// RoutesNotifierList is a list of notifiers along with configuration for
// when those notifiers should repeat alerting and whether they should be grouped
// by any specific values.
type RoutesNotifierList struct {
	Notifiers      []*RoutesNotifierListNotifier `json:"notifiers"`
	RepeatInterval string                        `json:"repeat_interval,omitempty"`
	GroupBy        *RoutesNotifierListGroupBy    `json:"group_by,omitempty"`
}

// RoutesOverride is a set of rules that override the default routes
type RoutesOverride struct {
	AlertLabelMatchers []*AlertLabelMatcher          `json:"alert_label_matchers"`
	Notifiers          map[string]RoutesNotifierList `json:"notifiers,omitempty"`
}

// RoutesNotifierListNotifier is a notifier in a list of notifiers
type RoutesNotifierListNotifier struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// RoutesNotifierListGroupBy is a list of label names by which notifications are grouped.
type RoutesNotifierListGroupBy struct {
	LabelNames []string `json:"label_names,omitempty"`
}

// AlertLabelMatcher represents a matcher for entities
type AlertLabelMatcher struct {
	Name  string            `json:"name,omitempty"`
	Type  *LabelMatcherType `json:"type,omitempty"`
	Value string            `json:"value,omitempty"`
}

func unmarshalNotificationPolicyData(data string) (*NotificationPolicyData, error) {
	policy := &NotificationPolicyData{}
	if err := json.Unmarshal([]byte(data), &policy); err != nil {
		return nil, fmt.Errorf("failed to unmarshal notification policy data: %w", err)
	}
	return policy, nil
}

func (n *NotificationPolicyData) ToModel() (*configmodels.Configv1NotificationPolicy, error) {
	routes, err := routesToModel(n.Routes)
	if err != nil {
		return nil, err
	}
	return &configmodels.Configv1NotificationPolicy{
		Routes: routes,
	}, nil
}

func routesToModel(routes *Routes) (*configmodels.NotificationPolicyRoutes, error) {
	defaultRoutes, err := notifierListToModel(routes.Defaults)
	if err != nil {
		return nil, err
	}
	overrides, err := sliceutil.MapErr(routes.Overrides, func(override *RoutesOverride) (*configmodels.NotificationPolicyRoutesOverride, error) {
		notifiers, err := notifierListToModel(override.Notifiers)
		if err != nil {
			return nil, err
		}
		alertLabelMatchers, err := sliceutil.MapErr(override.AlertLabelMatchers, alertLabelMatcherToModel)
		if err != nil {
			return nil, err
		}
		return &configmodels.NotificationPolicyRoutesOverride{
			AlertLabelMatchers: alertLabelMatchers,
			Notifiers:          notifiers,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return &configmodels.NotificationPolicyRoutes{
		Defaults:  defaultRoutes,
		Overrides: overrides,
	}, nil
}

func alertLabelMatcherToModel(m *AlertLabelMatcher) (*configmodels.Configv1LabelMatcher, error) {
	var matcherType configmodels.Configv1LabelMatcherMatcherType
	switch ptr.Unwrap(m.Type) {
	case ExactMatcherType:
		matcherType = configmodels.Configv1LabelMatcherMatcherTypeEXACT
	case RegexpMatcherType:
		matcherType = configmodels.Configv1LabelMatcherMatcherTypeREGEX
	default:
		return nil, fmt.Errorf("invalid alert label matcher type: %v", m.Type)
	}
	return &configmodels.Configv1LabelMatcher{
		Name:  m.Name,
		Type:  matcherType,
		Value: m.Value,
	}, nil
}

func notifierListToModel(
	notifierList map[string]RoutesNotifierList,
) (*configmodels.RoutesSeverityNotifiers, error) {
	// The API allows (and some use cases expect) that the notifiers for a particular
	// severity can be empty. However, other parts of the code (alerts v2 pathways)
	// look for the existence of a route to determine if an alert can be moved to a
	// bucket with a given notification policy. To avoid breaking those use cases,
	// we are very intentional here. If the notifier list we are given has a key
	// for the severity, then we add a route for that severity, with a possibly empty
	// notifier list. If the key does not exist, then we must leave it nil.
	// In Short:
	//   - user didn't specify severity --> nil
	//   - user did specify severity but nothing inside --> empty notifiers
	//   - user did specify severity + notifier --> notifiers
	notifiers := &configmodels.RoutesSeverityNotifiers{}
	if list, ok := notifierList["critical"]; ok {
		critList, err := notifierListMapToModel(list)
		if err != nil {
			return nil, err
		}
		notifiers.Critical = critList
	}

	if list, ok := notifierList["warn"]; ok {
		warnList, err := notifierListMapToModel(list)
		if err != nil {
			return nil, err
		}
		notifiers.Warn = warnList
	}
	return notifiers, nil
}

func notifierListMapToModel(notifierList RoutesNotifierList) (*configmodels.RoutesNotifierList, error) {
	durationInSecs, err := durationToSecs(notifierList.RepeatInterval)
	if err != nil {
		return nil, err
	}

	return &configmodels.RoutesNotifierList{
		NotifierSlugs:      sliceutil.Map(notifierList.Notifiers, func(n *RoutesNotifierListNotifier) string { return n.Slug }),
		RepeatIntervalSecs: durationInSecs,
		GroupBy:            notifierListToGroupByModel(notifierList),
	}, nil
}

func notifierListToGroupByModel(notifierList RoutesNotifierList) *configmodels.NotificationPolicyRoutesGroupBy {
	if notifierList.GroupBy == nil {
		return nil
	}
	routesGroupBy := &configmodels.NotificationPolicyRoutesGroupBy{}
	if notifierList.GroupBy.LabelNames != nil {
		routesGroupBy.LabelNames = notifierList.GroupBy.LabelNames
	}
	return routesGroupBy
}

func routesFromModel(m *configmodels.NotificationPolicyRoutes) (*Routes, error) {
	overrides, err := sliceutil.MapErr(m.Overrides, routesOverrideFromModel)
	if err != nil {
		return nil, err
	}
	return &Routes{
		Defaults:  notifierListMapFromModel(m.Defaults),
		Overrides: overrides,
	}, nil
}

func routesOverrideFromModel(o *configmodels.NotificationPolicyRoutesOverride) (*RoutesOverride, error) {
	labelMatchers, err := sliceutil.MapErr(o.AlertLabelMatchers, alertLabelsMatcherFromModel)
	if err != nil {
		return nil, err
	}
	return &RoutesOverride{
		AlertLabelMatchers: labelMatchers,
		Notifiers:          notifierListMapFromModel(o.Notifiers),
	}, nil
}

func alertLabelsMatcherFromModel(m *configmodels.Configv1LabelMatcher) (*AlertLabelMatcher, error) {
	if m == nil {
		return nil, nil
	}

	var matcherType LabelMatcherType
	switch m.Type {
	case configmodels.Configv1LabelMatcherMatcherTypeEXACT:
		matcherType = ExactMatcherType
	case configmodels.Configv1LabelMatcherMatcherTypeREGEX:
		matcherType = RegexpMatcherType
	default:
		return nil, fmt.Errorf("invalid alert label matcher type: %s", m.Type)
	}
	return &AlertLabelMatcher{
		Name:  m.Name,
		Type:  ptr.Wrap(matcherType),
		Value: m.Value,
	}, nil
}

func notifierListMapFromModel(m *configmodels.RoutesSeverityNotifiers) map[string]RoutesNotifierList {
	if m == nil {
		return nil
	}
	notifiers := map[string]RoutesNotifierList{}
	if m.Critical != nil {
		notifiers["critical"] = routesNotifierListFromModel(m.Critical)
	}
	if m.Warn != nil {
		notifiers["warn"] = routesNotifierListFromModel(m.Warn)
	}
	return notifiers
}

func routesNotifierListFromModel(m *configmodels.RoutesNotifierList) RoutesNotifierList {
	if m == nil {
		return RoutesNotifierList{}
	}
	return RoutesNotifierList{
		Notifiers:      sliceutil.Map(m.NotifierSlugs, routesNotifierListNotifierFromSlug),
		RepeatInterval: durationFromSecs(m.RepeatIntervalSecs),
		GroupBy:        routesNotifierListGroupByFromModel(m.GroupBy),
	}
}

func routesNotifierListNotifierFromSlug(slug string) *RoutesNotifierListNotifier {
	return &RoutesNotifierListNotifier{
		Slug: slug,
	}
}

func routesNotifierListGroupByFromModel(m *configmodels.NotificationPolicyRoutesGroupBy) *RoutesNotifierListGroupBy {
	if m == nil {
		return nil
	}

	routesGroupBy := &RoutesNotifierListGroupBy{}
	if m.LabelNames != nil {
		routesGroupBy.LabelNames = m.LabelNames
	}

	return routesGroupBy
}

func notificationPolicyResponseToInlineData(res *notification_policy.ReadNotificationPolicyOK) (string, error) {
	policy := res.GetPayload().NotificationPolicy
	if policy == nil {
		return "", fmt.Errorf("missing policy in response: %v", res)
	}

	routes, err := routesFromModel(policy.Routes)
	if err != nil {
		return "", err
	}
	npModel := &NotificationPolicyData{
		Routes: routes,
	}

	data, err := json.Marshal(&npModel)
	if err != nil {
		return "", fmt.Errorf("unable to marshal notification policy to JSON: %v", err)
	}

	return string(data), nil
}
