package enum

import (
	configunstable "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// SpanFilterMatchType is an enum.
var SpanFilterMatchType = newEnum("SpanFilterMatchType", []value[
	configunstable.SpanFilterSpanFilterMatchType,
	configv1.SpanFilterSpanFilterMatchType,
]{
	{
		legacy: configunstable.SpanFilterSpanFilterMatchTypeINCLUDE,
		v1:     configv1.SpanFilterSpanFilterMatchTypeINCLUDE,
		alias:  "INCLUDE",
	},
	{
		legacy: configunstable.SpanFilterSpanFilterMatchTypeEXCLUDE,
		v1:     configv1.SpanFilterSpanFilterMatchTypeEXCLUDE,
		alias:  "EXCLUDE",
	},
})
