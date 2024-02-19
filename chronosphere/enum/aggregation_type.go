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

package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// AggregationType is an enum.
var AggregationType = newEnum("AggregationType", []value[
	string,
	configv1.Configv1AggregationType,
]{
	{
		legacy:    "UNKNOWN_AT",
		isDefault: true,
	},
	{
		legacy: "LAST",
		v1:     configv1.Configv1AggregationTypeLAST,
		alias:  "LAST",
	},
	{
		legacy: "MIN",
		v1:     configv1.Configv1AggregationTypeMIN,
		alias:  "MIN",
	},
	{
		legacy: "MAX",
		v1:     configv1.Configv1AggregationTypeMAX,
		alias:  "MAX",
	},
	{
		legacy: "MEAN",
		v1:     configv1.Configv1AggregationTypeMEAN,
		alias:  "MEAN",
	},
	{
		legacy: "MEDIAN",
		v1:     configv1.Configv1AggregationTypeMEDIAN,
		alias:  "MEDIAN",
	},
	{
		legacy: "COUNT",
		v1:     configv1.Configv1AggregationTypeCOUNT,
		alias:  "COUNT",
	},
	{
		v1:    configv1.Configv1AggregationTypeCOUNTSAMPLES,
		alias: "COUNT_SAMPLES",
	},
	{
		legacy: "SUM",
		v1:     configv1.Configv1AggregationTypeSUM,
		alias:  "SUM",
	},
	{
		legacy: "SUMSQ",
		v1:     configv1.Configv1AggregationTypeSUMSQ,
		alias:  "SUMSQ",
	},
	{
		legacy: "STDEV",
		v1:     configv1.Configv1AggregationTypeSTDEV,
		alias:  "STDEV",
	},
	{
		legacy: "P10",
		v1:     configv1.Configv1AggregationTypeP10,
		alias:  "P10",
	},
	{
		legacy: "P20",
		v1:     configv1.Configv1AggregationTypeP20,
		alias:  "P20",
	},
	{
		legacy: "P25",
		v1:     configv1.Configv1AggregationTypeP25,
		alias:  "P25",
	},
	{
		legacy: "P30",
		v1:     configv1.Configv1AggregationTypeP30,
		alias:  "P30",
	},
	{
		legacy: "P40",
		v1:     configv1.Configv1AggregationTypeP40,
		alias:  "P40",
	},
	{
		legacy: "P50",
		v1:     configv1.Configv1AggregationTypeP50,
		alias:  "P50",
	},
	{
		legacy: "P60",
		v1:     configv1.Configv1AggregationTypeP60,
		alias:  "P60",
	},
	{
		legacy: "P70",
		v1:     configv1.Configv1AggregationTypeP70,
		alias:  "P70",
	},
	{
		legacy: "P75",
		v1:     configv1.Configv1AggregationTypeP75,
		alias:  "P75",
	},
	{
		legacy: "P80",
		v1:     configv1.Configv1AggregationTypeP80,
		alias:  "P80",
	},
	{
		legacy: "P90",
		v1:     configv1.Configv1AggregationTypeP90,
		alias:  "P90",
	},
	{
		legacy: "P95",
		v1:     configv1.Configv1AggregationTypeP95,
		alias:  "P95",
	},
	{
		legacy: "P99",
		v1:     configv1.Configv1AggregationTypeP99,
		alias:  "P99",
	},
	{
		legacy: "P999",
		v1:     configv1.Configv1AggregationTypeP999,
		alias:  "P999",
	},
	{
		legacy: "P9999",
		v1:     configv1.Configv1AggregationTypeP9999,
		alias:  "P9999",
	},
})
