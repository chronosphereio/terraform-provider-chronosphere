package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// RollupModeType is an enum.
var RollupModeType = newV1OnlyEnum("RollupModeType", []v1OnlyValue[configv1.Configv1RollupRuleMode]{
	{
		v1:        configv1.Configv1RollupRuleModeENABLED,
		isDefault: true,
	},
	{
		v1:    configv1.Configv1RollupRuleModePREVIEW,
		alias: "PREVIEW",
	},
})

// MappingModeType is an enum.
var MappingModeType = newV1OnlyEnum("MappingModeType", []v1OnlyValue[configv1.Configv1MappingRuleMode]{
	{
		v1:        configv1.Configv1MappingRuleModeENABLED,
		isDefault: true,
	},
	{
		v1:    configv1.Configv1MappingRuleModePREVIEW,
		alias: "PREVIEW",
	},
})
