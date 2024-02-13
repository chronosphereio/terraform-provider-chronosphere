package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// MetricType is an enum.
var MetricType = newEnum("MetricType", []value[
	string,
	configv1.RollupRuleMetricType,
]{
	{
		legacy:    "UNKNOWN_MT",
		isDefault: true,
	},
	{
		legacy: "COUNTER",
		v1:     configv1.RollupRuleMetricTypeCOUNTER,
		alias:  "COUNTER",
	},
	{
		legacy: "GAUGE",
		v1:     configv1.RollupRuleMetricTypeGAUGE,
		alias:  "GAUGE",
	},
	{
		legacy: "DELTA",
		v1:     configv1.RollupRuleMetricTypeDELTA,
		alias:  "DELTA",
	},
	{
		legacy: "DISTRIBUTION",
		v1:     configv1.RollupRuleMetricTypeDISTRIBUTION,
		alias:  "DISTRIBUTION",
	},
	{
		v1:    configv1.RollupRuleMetricTypeCUMULATIVEEXPONENTIALHISTOGRAM,
		alias: "CUMULATIVE_EXPONENTIAL_HISTOGRAM",
	},
	{
		v1:    configv1.RollupRuleMetricTypeMEASUREMENT,
		alias: "MEASUREMENT",
	},
})
