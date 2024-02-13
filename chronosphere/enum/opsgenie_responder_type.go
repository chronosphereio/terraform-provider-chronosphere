package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

var OpsgenieResponderType = newEnum("OpsgenieResponderType", []value[
	string,
	configv1.ResponderResponderType,
]{
	{
		legacy:    "UNKNOWN_RESPONSE_TYPE",
		isDefault: true,
	},
	{
		legacy: "TEAM",
		v1:     configv1.ResponderResponderTypeTEAM,
		alias:  "TEAM",
	},
	{
		legacy: "USER",
		v1:     configv1.ResponderResponderTypeUSER,
		alias:  "USER",
	},
	{
		legacy: "ESCALATION",
		v1:     configv1.ResponderResponderTypeESCALATION,
		alias:  "ESCALATION",
	},
	{
		legacy: "SCHEDULE",
		v1:     configv1.ResponderResponderTypeSCHEDULE,
		alias:  "SCHEDULE",
	},
})
