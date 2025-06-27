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
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSLOTimesliceIndicatorConversion(t *testing.T) {
	// Test conversion from intschema to model
	intSlo := &intschema.Slo{
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
	}

	converter := sloConverter{}

	// Convert to model
	model, err := converter.toModel(intSlo)
	require.NoError(t, err)

	// Verify the model fields
	assert.NotNil(t, model.Sli.CustomTimesliceIndicator)
	assert.Equal(t, "sum(rate(http_requests_total{status!~\"5..\"}[{{.TimeSlice}}])) / sum(rate(http_requests_total[{{.TimeSlice}}]))", model.Sli.CustomTimesliceIndicator.QueryTemplate)
	assert.Equal(t, models.SLITimeSliceSizeTIMESLICESIZEONEMINUTE, model.Sli.CustomTimesliceIndicator.TimesliceSize)
	assert.NotNil(t, model.Sli.CustomTimesliceIndicator.Condition)
	assert.Equal(t, models.ConditionOpGEQ, model.Sli.CustomTimesliceIndicator.Condition.Op)
	assert.Equal(t, 0.99, model.Sli.CustomTimesliceIndicator.Condition.Value)

	// Convert back to intschema
	intSloResult, err := converter.fromModel(model)
	require.NoError(t, err)

	// Verify round-trip conversion
	assert.NotNil(t, intSloResult.Sli.CustomTimesliceIndicator)
	assert.Equal(t, intSlo.Sli.CustomTimesliceIndicator.QueryTemplate, intSloResult.Sli.CustomTimesliceIndicator.QueryTemplate)
	assert.Equal(t, intSlo.Sli.CustomTimesliceIndicator.TimesliceSize, intSloResult.Sli.CustomTimesliceIndicator.TimesliceSize)
	assert.Equal(t, intSlo.Sli.CustomTimesliceIndicator.Condition.Op, intSloResult.Sli.CustomTimesliceIndicator.Condition.Op)
	assert.Equal(t, intSlo.Sli.CustomTimesliceIndicator.Condition.Value, intSloResult.Sli.CustomTimesliceIndicator.Condition.Value)
}
