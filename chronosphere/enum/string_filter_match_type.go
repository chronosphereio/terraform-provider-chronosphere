package enum

import (
	configunstable "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// StringFilterMatchType is an enum.
var StringFilterMatchType = newEnum("StringFilterMatchType", []value[
	configunstable.StringFilterStringFilterMatchType,
	configv1.StringFilterStringFilterMatchType,
]{
	{
		legacy: configunstable.StringFilterStringFilterMatchTypeEXACT,
		v1:     configv1.StringFilterStringFilterMatchTypeEXACT,
		alias:  "EXACT",
	},
	{
		legacy: configunstable.StringFilterStringFilterMatchTypeREGEX,
		v1:     configv1.StringFilterStringFilterMatchTypeREGEX,
		alias:  "REGEX",
	},
})
