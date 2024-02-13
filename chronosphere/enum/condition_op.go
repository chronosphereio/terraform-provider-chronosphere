package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// ConditionOp is an enum.
var ConditionOp = newEnum("ConditionOp", []value[
	string,
	configv1.ConditionOp,
]{
	{
		legacy:    "INVALID",
		isDefault: true,
	},
	{
		legacy: "GEQ",
		v1:     configv1.ConditionOpGEQ,
		alias:  "GEQ",
	},
	{
		legacy: "GT",
		v1:     configv1.ConditionOpGT,
		alias:  "GT",
	},
	{
		legacy: "LEQ",
		v1:     configv1.ConditionOpLEQ,
		alias:  "LEQ",
	},
	{
		legacy: "LT",
		v1:     configv1.ConditionOpLT,
		alias:  "LT",
	},
	{
		legacy: "EQ",
		v1:     configv1.ConditionOpEQ,
		alias:  "EQ",
	},
	{
		legacy: "NEQ",
		v1:     configv1.ConditionOpNEQ,
		alias:  "NEQ",
	},
	{
		legacy: "EXISTS",
		v1:     configv1.ConditionOpEXISTS,
		alias:  "EXISTS",
	},
	{
		legacy: "NOT_EXISTS",
		v1:     configv1.ConditionOpNOTEXISTS,
		alias:  "NOT_EXISTS",
	},
})
