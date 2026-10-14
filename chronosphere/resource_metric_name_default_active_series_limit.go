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
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// MetricNameDefaultActiveSeriesLimitFromModel maps an API model to the intschema model.
func MetricNameDefaultActiveSeriesLimitFromModel(
	m *models.Configv1MetricNameDefaultActiveSeriesLimit,
) (*intschema.MetricNameDefaultActiveSeriesLimit, error) {
	return (metricNameDefaultActiveSeriesLimitConverter{}).fromModel(m)
}

func resourceMetricNameDefaultActiveSeriesLimit() *schema.Resource {
	r := newGenericResource(
		"metric_name_default_active_series_limit",
		metricNameDefaultActiveSeriesLimitConverter{},
		generatedMetricNameDefaultActiveSeriesLimit{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description: "Singleton per-tenant default cap on the number of active time series accepted " +
			"for each metric name over the active time series rolling window.",
		Schema:        tfschema.MetricNameDefaultActiveSeriesLimit,
		CustomizeDiff: r.ValidateDryRun(&MetricNameDefaultActiveSeriesLimitDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// MetricNameDefaultActiveSeriesLimitDryRunCount tracks how many times dry run is run during validation for testing.
var MetricNameDefaultActiveSeriesLimitDryRunCount atomic.Int64

type metricNameDefaultActiveSeriesLimitConverter struct{}

func (metricNameDefaultActiveSeriesLimitConverter) toModel(
	l *intschema.MetricNameDefaultActiveSeriesLimit,
) (*models.Configv1MetricNameDefaultActiveSeriesLimit, error) {
	return &models.Configv1MetricNameDefaultActiveSeriesLimit{
		DefaultMaxActiveStandardSeries:  formatOptionalInt64(l.DefaultMaxActiveStandardSeries),
		DefaultMaxActiveHistogramSeries: formatOptionalInt64(l.DefaultMaxActiveHistogramSeries),
	}, nil
}

func (metricNameDefaultActiveSeriesLimitConverter) fromModel(
	m *models.Configv1MetricNameDefaultActiveSeriesLimit,
) (*intschema.MetricNameDefaultActiveSeriesLimit, error) {
	standard, err := parseStringToInt64(m.DefaultMaxActiveStandardSeries, "default_max_active_standard_series")
	if err != nil {
		return nil, err
	}
	histogram, err := parseStringToInt64(m.DefaultMaxActiveHistogramSeries, "default_max_active_histogram_series")
	if err != nil {
		return nil, err
	}
	return &intschema.MetricNameDefaultActiveSeriesLimit{
		DefaultMaxActiveStandardSeries:  standard,
		DefaultMaxActiveHistogramSeries: histogram,
	}, nil
}

// formatOptionalInt64 encodes an int64 for the wire, where it is a string.
// Zero means "use the managed default", so it is omitted rather than sent.
func formatOptionalInt64(v int64) string {
	if v == 0 {
		return ""
	}
	return strconv.FormatInt(v, 10)
}
