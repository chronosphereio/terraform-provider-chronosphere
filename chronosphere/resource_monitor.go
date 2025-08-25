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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"
	"go.uber.org/multierr"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// MonitorFromModel maps an API model to an intschema model.
func MonitorFromModel(m *models.Configv1Monitor) (*intschema.Monitor, error) {
	return monitorConverter{}.fromModel(m)
}

func resourceMonitor() *schema.Resource {
	r := newGenericResource(
		"monitor",
		monitorConverter{},
		generatedMonitor{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRunOptions(&MonitorDryRunCount, ValidateDryRunOpts[*models.Configv1Monitor]{
			DryRunDefaults: map[string]any{
				"query.[0].prometheus_expr": "dry_run_unknown_query",
			},
		}),
		Schema: tfschema.Monitor,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

const (
	warn     = "warn"
	critical = "critical"
)

var validSeverities = map[string]struct{}{
	warn:     {},
	critical: {},
}

func checkSeverity(sev string) error {
	if _, ok := validSeverities[sev]; !ok {
		return fmt.Errorf(
			"%q is not a valid severity; must be one of %v",
			sev, sortedKeys(validSeverities))
	}
	return nil
}

const (
	monday    = "monday"
	tuesday   = "tuesday"
	wednesday = "wednesday"
	thursday  = "thursday"
	friday    = "friday"
	saturday  = "saturday"
	sunday    = "sunday"

	minRangeStart = "00:00"
	maxRangeEnd   = "24:00"
)

var validDays = map[string]struct{}{
	monday:    {},
	tuesday:   {},
	wednesday: {},
	thursday:  {},
	friday:    {},
	saturday:  {},
	sunday:    {},
}

// MonitorDryRunCount tracks how many times dry run is run during validation for testing.
var MonitorDryRunCount atomic.Int64

type monitorConverter struct{}

func (monitorConverter) toModel(
	m *intschema.Monitor,
) (*models.Configv1Monitor, error) {
	intervalSecs, err := durationToSecs(m.Interval)
	if err != nil {
		return nil, err
	}
	seriesConditions, err := monitorSeriesConditionsToModel(m.SeriesConditions)
	if err != nil {
		return nil, err
	}
	schedule, err := monitorScheduleToModel(m.Schedule)
	if err != nil {
		return nil, err
	}
	collSlug, collRef := collectionRefFromID(m.CollectionId.Slug())
	return &models.Configv1Monitor{
		Annotations:            m.Annotations,
		BucketSlug:             m.BucketId.Slug(),
		CollectionSlug:         collSlug,
		Collection:             collRef,
		GraphiteQuery:          m.Query.GraphiteExpr,
		IntervalSecs:           intervalSecs,
		Labels:                 m.Labels,
		Name:                   m.Name,
		NotificationPolicySlug: m.NotificationPolicyId.Slug(),
		PrometheusQuery:        m.Query.PrometheusExpr,
		Schedule:               schedule,
		SeriesConditions:       seriesConditions,
		SignalGrouping:         monitorSignalGroupingToModel(m.SignalGrouping),
		Slug:                   m.Slug,
		LoggingQuery:           m.Query.LoggingExpr,
	}, nil
}

func (monitorConverter) fromModel(
	m *models.Configv1Monitor,
) (*intschema.Monitor, error) {
	schedule, err := monitorScheduleFromModel(m.Schedule)
	if err != nil {
		return nil, err
	}
	return &intschema.Monitor{
		Name:                 m.Name,
		Slug:                 m.Slug,
		BucketId:             tfid.Slug(m.BucketSlug),
		CollectionId:         tfid.Slug(collectionIDFromRef(m.CollectionSlug, m.Collection)),
		NotificationPolicyId: tfid.Slug(m.NotificationPolicySlug),
		Query: intschema.MonitorQuery{
			GraphiteExpr:   m.GraphiteQuery,
			PrometheusExpr: m.PrometheusQuery,
			LoggingExpr:    m.LoggingQuery,
		},
		SeriesConditions: monitorSeriesConditionsFromModel(m.SeriesConditions),
		Annotations:      m.Annotations,
		Interval:         durationFromSecs(m.IntervalSecs),
		Labels:           m.Labels,
		Schedule:         schedule,
		SignalGrouping:   monitorSignalGroupingFromModel(m.SignalGrouping),
	}, nil
}

func monitorScheduleToModel(s *intschema.MonitorSchedule) (*models.MonitorSchedule, error) {
	if s == nil {
		return nil, nil
	}
	w, err := monitorWeeklyScheduleToModel(s.Range)
	if err != nil {
		return nil, err
	}
	return &models.MonitorSchedule{
		Timezone:       s.Timezone,
		WeeklySchedule: w,
	}, nil
}

func monitorScheduleFromModel(s *models.MonitorSchedule) (*intschema.MonitorSchedule, error) {
	if s == nil {
		return nil, nil
	}
	ranges, err := monitorWeeklyScheduleFromModel(s.WeeklySchedule)
	if err != nil {
		return nil, err
	}
	return &intschema.MonitorSchedule{
		Timezone: s.Timezone,
		Range:    ranges,
	}, nil
}

func monitorWeeklyScheduleToModel(
	ranges []intschema.MonitorScheduleRange,
) (*models.ScheduleWeeklySchedule, error) {
	byDay := make(map[string][]*models.ScheduleDayTimeRange)
	for _, r := range ranges {
		day := strings.ToLower(r.Day)
		if _, ok := validDays[day]; !ok {
			return nil, fmt.Errorf(
				"unknown schedule day %q, must be one of %v",
				r.Day, sortedKeys(validDays))
		}
		byDay[day] = append(byDay[day], &models.ScheduleDayTimeRange{
			StartHhMm: r.Start,
			EndHhMm:   r.End,
		})
	}

	load := func(day string) *models.ScheduleScheduleDay {
		rs := byDay[day]
		if len(rs) == 0 {
			return &models.ScheduleScheduleDay{
				Active: models.ScheduleDayActiveNEVER,
			}
		}
		if len(rs) == 1 && rs[0].StartHhMm == minRangeStart && rs[0].EndHhMm == maxRangeEnd {
			return &models.ScheduleScheduleDay{
				Active: models.ScheduleDayActiveALLDAY,
			}
		}
		return &models.ScheduleScheduleDay{
			Active: models.ScheduleDayActiveONLYDURINGRANGES,
			Ranges: rs,
		}
	}
	return &models.ScheduleWeeklySchedule{
		Monday:    load(monday),
		Tuesday:   load(tuesday),
		Wednesday: load(wednesday),
		Thursday:  load(thursday),
		Friday:    load(friday),
		Saturday:  load(saturday),
		Sunday:    load(sunday),
	}, nil
}

func monitorWeeklyScheduleFromModel(
	w *models.ScheduleWeeklySchedule,
) ([]intschema.MonitorScheduleRange, error) {
	if w == nil {
		return nil, nil
	}

	var out []intschema.MonitorScheduleRange
	var errs error
	load := func(day string, f *models.ScheduleScheduleDay) {
		switch f.Active {
		case models.ScheduleDayActiveNEVER:
			// noop
		case models.ScheduleDayActiveALLDAY:
			out = append(out, intschema.MonitorScheduleRange{
				Day:   day,
				Start: minRangeStart,
				End:   maxRangeEnd,
			})
		case models.ScheduleDayActiveONLYDURINGRANGES:
			for _, r := range f.Ranges {
				out = append(out, intschema.MonitorScheduleRange{
					Day:   day,
					Start: r.StartHhMm,
					End:   r.EndHhMm,
				})
			}
		default:
			errs = multierr.Append(errs, fmt.Errorf(
				"unknown schedule active value on %q: %v", day, f.Active))
		}
	}
	load(monday, w.Monday)
	load(tuesday, w.Tuesday)
	load(wednesday, w.Wednesday)
	load(thursday, w.Thursday)
	load(friday, w.Friday)
	load(saturday, w.Saturday)
	load(sunday, w.Sunday)

	if errs != nil {
		return nil, errs
	}
	return out, nil
}

func monitorSeriesConditionsToModel(
	c intschema.MonitorSeriesConditions,
) (*models.MonitorSeriesConditions, error) {
	defaults, err := monitorConditionsToModel(c.Condition)
	if err != nil {
		return nil, err
	}
	overrides, err := monitorOverridesToModel(c.Override)
	if err != nil {
		return nil, err
	}
	return &models.MonitorSeriesConditions{
		Defaults:  defaults,
		Overrides: overrides,
	}, nil
}

func monitorSeriesConditionsFromModel(
	c *models.MonitorSeriesConditions,
) intschema.MonitorSeriesConditions {
	if c == nil {
		return intschema.MonitorSeriesConditions{}
	}
	return intschema.MonitorSeriesConditions{
		Condition: monitorConditionsFromModel(c.Defaults),
		Override:  monitorOverridesFromModel(c.Overrides),
	}
}

func monitorOverridesToModel(
	overrides []intschema.MonitorSeriesConditionsOverride,
) ([]*models.MonitorSeriesConditionsOverride, error) {
	var out []*models.MonitorSeriesConditionsOverride
	for _, o := range overrides {
		conds, err := monitorConditionsToModel(o.Condition)
		if err != nil {
			return nil, err
		}
		out = append(out, &models.MonitorSeriesConditionsOverride{
			LabelMatchers:      matchersToModel(o.LabelMatcher),
			SeverityConditions: conds,
		})
	}
	return out, nil
}

func monitorOverridesFromModel(
	overrides []*models.MonitorSeriesConditionsOverride,
) []intschema.MonitorSeriesConditionsOverride {
	var out []intschema.MonitorSeriesConditionsOverride
	for _, o := range overrides {
		out = append(out, intschema.MonitorSeriesConditionsOverride{
			Condition:    monitorConditionsFromModel(o.SeverityConditions),
			LabelMatcher: matchersFromModel(o.LabelMatchers),
		})
	}
	return out
}

func monitorConditionsToModel(
	conds []intschema.MonitorSeriesCondition,
) (*models.SeriesConditionsSeverityConditions, error) {
	if len(conds) == 0 {
		return nil, nil
	}

	bySev := make(map[string][]*models.Configv1MonitorCondition)
	for _, c := range conds {
		if err := checkSeverity(c.Severity); err != nil {
			return nil, err
		}
		sustainSecs, err := durationToSecs(c.Sustain)
		if err != nil {
			return nil, err
		}
		resolveSustainSecs, err := durationToSecs(c.ResolveSustain)
		if err != nil {
			return nil, err
		}
		bySev[c.Severity] = append(bySev[c.Severity], &models.Configv1MonitorCondition{
			Op:                 enum.ConditionOp.V1(c.Op),
			SustainSecs:        sustainSecs,
			ResolveSustainSecs: resolveSustainSecs,
			Value:              c.Value,
		})
	}

	load := func(sev string) *models.SeriesConditionsConditions {
		cs := bySev[sev]
		if len(cs) == 0 {
			return nil
		}
		return &models.SeriesConditionsConditions{
			Conditions: cs,
		}
	}

	return &models.SeriesConditionsSeverityConditions{
		Critical: load(critical),
		Warn:     load(warn),
	}, nil
}

func monitorConditionsFromModel(
	m *models.SeriesConditionsSeverityConditions,
) []intschema.MonitorSeriesCondition {
	if m == nil {
		return nil
	}

	var out []intschema.MonitorSeriesCondition
	load := func(sev string, f *models.SeriesConditionsConditions) {
		if f == nil {
			return
		}
		for _, c := range f.Conditions {
			out = append(out, intschema.MonitorSeriesCondition{
				Severity:       sev,
				Op:             string(c.Op),
				Sustain:        durationFromSecs(c.SustainSecs),
				ResolveSustain: durationFromSecs(c.ResolveSustainSecs),
				Value:          c.Value,
			})
		}
	}
	load(warn, m.Warn)
	load(critical, m.Critical)
	return out
}

func monitorSignalGroupingToModel(
	g *intschema.SignalGrouping,
) *models.MonitorSignalGrouping {
	if g == nil {
		return nil
	}
	return &models.MonitorSignalGrouping{
		LabelNames:      g.LabelNames,
		SignalPerSeries: g.SignalPerSeries,
	}
}

func monitorSignalGroupingFromModel(
	g *models.MonitorSignalGrouping,
) *intschema.SignalGrouping {
	if g == nil {
		return nil
	}
	return &intschema.SignalGrouping{
		LabelNames:      g.LabelNames,
		SignalPerSeries: g.SignalPerSeries,
	}
}
