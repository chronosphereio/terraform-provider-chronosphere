package enum

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// TraceMetricsRuleGroupByType is an enum.
var TraceMetricsRuleGroupByType = newEnum("TraceMetricsRuleGroupByType", []value[
	string,
	models.GroupByKeyGroupByKeyType,
]{
	{
		v1:     models.GroupByKeyGroupByKeyTypeTAG,
		legacy: "TAG",
		alias:  "TAG",
	},
	{
		v1:     models.GroupByKeyGroupByKeyTypeOPERATION,
		legacy: "OPERATION",
		alias:  "OPERATION",
	},
	{
		v1:     models.GroupByKeyGroupByKeyTypeSERVICE,
		legacy: "SERVICE",
		alias:  "SERVICE",
	},
})
