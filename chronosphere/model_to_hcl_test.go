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

package chronosphere_test

import (
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal/hcltest"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitorToHCL(t *testing.T) {
	tests := []struct {
		m    *models.Configv1Monitor
		want string
	}{
		{
			m: &models.Configv1Monitor{
				Name: "empty",
			},
			want: `
resource "chronosphere_monitor" "" {
  name = "empty"

  query {
  }

  series_conditions {
  }
}
`,
		},
		{
			m: &models.Configv1Monitor{
				Name:            "basic monitor",
				BucketSlug:      "b",
				PrometheusQuery: `up{cluster="foo"}`,
				SeriesConditions: &models.MonitorSeriesConditions{
					Defaults: &models.SeriesConditionsSeverityConditions{
						Warn: &models.SeriesConditionsConditions{
							Conditions: []*models.Configv1MonitorCondition{{
								Op:    models.ConditionOpEQ,
								Value: 10,
							}},
						},
					},
				},
			},
			want: `
resource "chronosphere_monitor" "" {
  name      = "basic monitor"
  bucket_id = "b"

  query {
    prometheus_expr = "up{cluster=\"foo\"}"
  }

  series_conditions {
    condition {
      op       = "EQ"
      severity = "warn"
      value    = 10
    }
  }
}
`,
		},
		{
			m: &models.Configv1Monitor{
				Name:            "empty condition",
				BucketSlug:      "b",
				PrometheusQuery: `up{cluster="foo"}`,
				SeriesConditions: &models.MonitorSeriesConditions{
					Defaults: &models.SeriesConditionsSeverityConditions{},
				},
			},
			want: `
resource "chronosphere_monitor" "" {
  name      = "empty condition"
  bucket_id = "b"

  query {
    prometheus_expr = "up{cluster=\"foo\"}"
  }

  series_conditions {
  }
}
`,
		},
		{
			m: &models.Configv1Monitor{
				Name:            "empty override",
				BucketSlug:      "b",
				PrometheusQuery: `up{cluster="foo"}`,
				SeriesConditions: &models.MonitorSeriesConditions{
					Defaults: &models.SeriesConditionsSeverityConditions{
						Warn: &models.SeriesConditionsConditions{
							Conditions: []*models.Configv1MonitorCondition{{
								Op:    models.ConditionOpEQ,
								Value: 10,
							}},
						},
					},
					Overrides: []*models.MonitorSeriesConditionsOverride{{}},
				},
			},
			want: `
resource "chronosphere_monitor" "" {
  name      = "empty override"
  bucket_id = "b"

  query {
    prometheus_expr = "up{cluster=\"foo\"}"
  }

  series_conditions {
    condition {
      op       = "EQ"
      severity = "warn"
      value    = 10
    }

    override {
    }
  }
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.m.Name, func(t *testing.T) {
			s, err := chronosphere.MonitorFromModel(tt.m)
			require.NoError(t, err)

			got := hcltest.MustMarshalString(t, s)
			assert.Equal(t, tt.want, got)
		})
	}
}
