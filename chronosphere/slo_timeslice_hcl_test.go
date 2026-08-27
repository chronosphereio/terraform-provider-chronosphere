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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal/hcltest"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"

	"github.com/stretchr/testify/assert"
)

func TestSLOTimesliceIndicatorToHCL(t *testing.T) {
	tests := []struct {
		name string
		slo  *intschema.Slo
		want string
	}{
		{
			name: "basic timeslice indicator",
			slo: &intschema.Slo{
				HCLID:        "test",
				Name:         "test timeslice slo",
				CollectionId: tfid.Slug("test-collection"),
				Definition: intschema.SloDefinition{
					Objective: 99.95,
				},
				Sli: intschema.SloSli{
					CustomTimesliceIndicator: &intschema.SloSliCustomTimesliceIndicator{
						QueryTemplate: "sum(rate(http_requests_total{status!~\"5..\"}[{{.TimeSlice}}])) / sum(rate(http_requests_total[{{.TimeSlice}}]))",
						TimesliceSize: "ONE_MINUTE",
						Condition: intschema.SloSliCustomTimesliceIndicatorCondition{
							Op:    "GEQ",
							Value: 0.99,
						},
					},
				},
			},
			want: `
resource "chronosphere_slo" "test" {
  name          = "test timeslice slo"
  collection_id = "test-collection"

  definition {
    objective = 99.95
  }

  sli {
    custom_timeslice_indicator {
      condition {
        op    = "GEQ"
        value = 0.99
      }

      query_template = "sum(rate(http_requests_total{status!~\"5..\"}[{{.TimeSlice}}])) / sum(rate(http_requests_total[{{.TimeSlice}}]))"
      timeslice_size = "ONE_MINUTE"
    }
  }
}
`,
		},
		{
			name: "timeslice indicator with five minute window and LEQ condition",
			slo: &intschema.Slo{
				HCLID:        "latency",
				Name:         "latency timeslice slo",
				CollectionId: tfid.Slug("test-collection"),
				Definition: intschema.SloDefinition{
					Objective: 95.0,
				},
				Sli: intschema.SloSli{
					CustomTimesliceIndicator: &intschema.SloSliCustomTimesliceIndicator{
						QueryTemplate: "histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[{{.TimeSlice}}])) by (le))",
						TimesliceSize: "FIVE_MINUTES",
						Condition: intschema.SloSliCustomTimesliceIndicatorCondition{
							Op:    "LEQ",
							Value: 0.5,
						},
					},
				},
			},
			want: `
resource "chronosphere_slo" "latency" {
  name          = "latency timeslice slo"
  collection_id = "test-collection"

  definition {
    objective = 95
  }

  sli {
    custom_timeslice_indicator {
      condition {
        op    = "LEQ"
        value = 0.5
      }

      query_template = "histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[{{.TimeSlice}}])) by (le))"
      timeslice_size = "FIVE_MINUTES"
    }
  }
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hcltest.MustMarshalString(t, tt.slo)
			assert.Equal(t, tt.want, got)
		})
	}
}
