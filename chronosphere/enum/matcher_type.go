package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// MatcherType is an enum.
var MatcherType = newEnum("MatcherType", []value[
	string,
	configv1.Configv1LabelMatcherMatcherType,
]{
	{
		legacy:    "INVALID_MATCHER_TYPE",
		isDefault: true,
	},
	{
		legacy: "EXACT_MATCHER_TYPE",
		v1:     configv1.Configv1LabelMatcherMatcherTypeEXACT,
		alias:  "EXACT",
	},
	{
		legacy: "REGEXP_MATCHER_TYPE",
		v1:     configv1.Configv1LabelMatcherMatcherTypeREGEX,
		alias:  "REGEX",
	},
})
