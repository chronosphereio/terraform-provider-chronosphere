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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceMetricNameActiveSeriesLimit() *schema.Resource {
	r := newGenericResource(
		"metric_name_active_series_limit",
		metricNameActiveSeriesLimitConverter{},
		generatedUnstableMetricNameActiveSeriesLimit{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description: "A per-metric-name cap on the number of active time series accepted " +
			"over the active time series rolling window. " + unstableAPIWarning,
		Schema:        tfschema.MetricNameActiveSeriesLimit,
		CustomizeDiff: r.ValidateDryRun(&MetricNameActiveSeriesLimitDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// MetricNameActiveSeriesLimitDryRunCount tracks how many times dry run is run during validation for testing.
var MetricNameActiveSeriesLimitDryRunCount atomic.Int64

type metricNameActiveSeriesLimitConverter struct{}

func (metricNameActiveSeriesLimitConverter) toModel(
	l *intschema.MetricNameActiveSeriesLimit,
) (*models.ConfigunstableMetricNameActiveSeriesLimit, error) {
	return &models.ConfigunstableMetricNameActiveSeriesLimit{
		Name:       l.Name,
		Slug:       l.Slug,
		MetricName: l.MetricName,
		// int64 is encoded as a string on the wire. Always send it explicitly:
		// the API rejects zero, so omitempty would never be correct here.
		MaxActiveSeries: strconv.FormatInt(l.MaxActiveSeries, 10),
	}, nil
}

func (metricNameActiveSeriesLimitConverter) fromModel(
	m *models.ConfigunstableMetricNameActiveSeriesLimit,
) (*intschema.MetricNameActiveSeriesLimit, error) {
	maxActiveSeries, err := parseStringToInt64(m.MaxActiveSeries, "max_active_series")
	if err != nil {
		return nil, err
	}
	return &intschema.MetricNameActiveSeriesLimit{
		Name:            m.Name,
		Slug:            m.Slug,
		MetricName:      m.MetricName,
		MaxActiveSeries: maxActiveSeries,
	}, nil
}
