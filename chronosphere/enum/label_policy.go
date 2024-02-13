package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// LabelPolicy is an enum.
var LabelPolicy = newEnum("LabelPolicy", []value[
	configv1.Configv1DerivedLabelLabelPolicy,
	configv1.Configv1DerivedLabelLabelPolicy,
]{
	{
		legacy: "KEEP",
		v1:     configv1.Configv1DerivedLabelLabelPolicyKEEP,
		alias:  "KEEP",
	},
	{
		legacy: "OVERRIDE",
		v1:     configv1.Configv1DerivedLabelLabelPolicyOVERRIDE,
		alias:  "OVERRIDE",
	},
})
