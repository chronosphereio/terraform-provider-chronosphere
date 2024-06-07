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
	"fmt"
	"strings"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func newIndependentNotificationPolicy() genericResource[
	*models.Configv1NotificationPolicy,
	intschema.NotificationPolicy,
	*intschema.NotificationPolicy,
] {
	return newGenericResource[
		*models.Configv1NotificationPolicy,
		intschema.NotificationPolicy,
		*intschema.NotificationPolicy,
	](
		"notification_policy",
		independentNotificationPolicyConverter{},
		generatedNotificationPolicy{})
}

// NotificationPolicyFromModel maps an API model to an intschema model.
func NotificationPolicyFromModel(m *models.Configv1NotificationPolicy) (*intschema.NotificationPolicy, error) {
	return independentNotificationPolicyConverter{}.fromModel(m)
}

type independentNotificationPolicyConverter struct{}

func (independentNotificationPolicyConverter) toModel(
	p *intschema.NotificationPolicy,
) (*models.Configv1NotificationPolicy, error) {
	defaults, err := notificationRoutesToModel(p.Route)
	if err != nil {
		return nil, err
	}
	overrides, err := sliceutil.MapErr(p.Override, notificationOverrideToModel)
	if err != nil {
		return nil, err
	}
	return &models.Configv1NotificationPolicy{
		Name: p.Name,
		Routes: &models.NotificationPolicyRoutes{
			Defaults:  defaults,
			Overrides: overrides,
		},
		Slug:     p.Slug,
		TeamSlug: p.TeamId.Slug(),
	}, nil
}

func (independentNotificationPolicyConverter) fromModel(
	m *models.Configv1NotificationPolicy,
) (*intschema.NotificationPolicy, error) {
	var (
		overrides []intschema.NotificationPolicyOverride
		routes    []intschema.NotificationRoute
	)
	if m.Routes != nil {
		overrides = sliceutil.Map(m.Routes.Overrides, notificationOverrideFromModel)
		routes = notificationRoutesFromModel(m.Routes.Defaults)
	}

	// Note: BucketSlug being set should cause a change to be detected by TF.
	// Since a policy is either bucket-owned or team-owned, a bucket owned policy
	// will have no team, so the team will mismatch.
	return &intschema.NotificationPolicy{
		Name:                   m.Name,
		Slug:                   m.Slug,
		TeamId:                 tfid.Slug(m.TeamSlug),
		Override:               overrides,
		Route:                  routes,
		NotificationPolicyData: tfid.Slug(tfschema.IndependentNotificationPolicyData),
		IsIndependent:          true,
	}, nil
}

func notificationRoutesToModel(
	routes []intschema.NotificationRoute,
) (*models.RoutesSeverityNotifiers, error) {
	if len(routes) == 0 {
		return nil, nil
	}

	bySev := make(map[string]*models.RoutesNotifierList)
	for _, r := range routes {
		sev := strings.ToLower(r.Severity)
		if err := checkSeverity(sev); err != nil {
			return nil, err
		}
		if _, ok := bySev[sev]; ok {
			return nil, fmt.Errorf("duplicate route with severity=%v", sev)
		}

		intervalSecs, err := durationToSecs(r.RepeatInterval)
		if err != nil {
			return nil, err
		}
		bySev[sev] = &models.RoutesNotifierList{
			NotifierSlugs:      sliceutil.Map(r.Notifiers, (tfid.ID).Slug),
			RepeatIntervalSecs: intervalSecs,
			GroupBy: &models.NotificationPolicyRoutesGroupBy{
				LabelNames: r.GroupBy,
			},
		}
	}
	return &models.RoutesSeverityNotifiers{
		Warn:     bySev[warn],
		Critical: bySev[critical],
	}, nil
}

func notificationRoutesFromModel(
	m *models.RoutesSeverityNotifiers,
) []intschema.NotificationRoute {
	if m == nil {
		return nil
	}
	var out []intschema.NotificationRoute
	load := func(sev string, f *models.RoutesNotifierList) {
		if f == nil {
			return
		}
		var groupBy []string
		if f.GroupBy != nil {
			groupBy = f.GroupBy.LabelNames
		}
		out = append(out, intschema.NotificationRoute{
			Severity:       sev,
			Notifiers:      sliceutil.Map(f.NotifierSlugs, tfid.Slug),
			RepeatInterval: durationFromSecs(f.RepeatIntervalSecs),
			GroupBy:        groupBy,
		})
	}
	load(warn, m.Warn)
	load(critical, m.Critical)
	return out
}

func notificationOverrideToModel(
	o intschema.NotificationPolicyOverride,
) (*models.NotificationPolicyRoutesOverride, error) {
	routes, err := notificationRoutesToModel(o.Route)
	if err != nil {
		return nil, err
	}
	return &models.NotificationPolicyRoutesOverride{
		AlertLabelMatchers: matchersToModel(o.AlertLabelMatcher),
		Notifiers:          routes,
	}, nil
}

func notificationOverrideFromModel(
	o *models.NotificationPolicyRoutesOverride,
) intschema.NotificationPolicyOverride {
	return intschema.NotificationPolicyOverride{
		AlertLabelMatcher: matchersFromModel(o.AlertLabelMatchers),
		Route:             notificationRoutesFromModel(o.Notifiers),
	}
}
