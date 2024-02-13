package enum

import (
	configunstable "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// NumericFilterComparisonType is an enum.
var NumericFilterComparisonType = newEnum("NumericFilterComparisonType", []value[
	configunstable.NumericFilterComparisonType,
	configv1.NumericFilterComparisonType,
]{
	{
		legacy: configunstable.NumericFilterComparisonTypeEQUAL,
		v1:     configv1.NumericFilterComparisonTypeEQUAL,
		alias:  "EQUAL",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeNOTEQUAL,
		v1:     configv1.NumericFilterComparisonTypeNOTEQUAL,
		alias:  "NOT_EQUAL",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeGREATERTHAN,
		v1:     configv1.NumericFilterComparisonTypeGREATERTHAN,
		alias:  "GREATER_THAN",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeGREATERTHANOREQUAL,
		v1:     configv1.NumericFilterComparisonTypeGREATERTHANOREQUAL,
		alias:  "GREATER_THAN_OR_EQUAL",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeLESSTHAN,
		v1:     configv1.NumericFilterComparisonTypeLESSTHAN,
		alias:  "LESS_THAN",
	},
	{
		legacy: configunstable.NumericFilterComparisonTypeLESSTHANOREQUAL,
		v1:     configv1.NumericFilterComparisonTypeLESSTHANOREQUAL,
		alias:  "LESS_THAN_OR_EQUAL",
	},
})
