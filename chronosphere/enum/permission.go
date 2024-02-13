package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// Permission represents a mapping between the external API permissions and the
// v1 version
var Permission = newEnum("Permission", []value[
	string,
	configv1.MetricsRestrictionPermission,
]{
	{
		legacy:    "UNKNOWN_PERMISSION",
		isDefault: true,
	},
	{
		legacy: "READ_PERMISSION",
		v1:     configv1.MetricsRestrictionPermissionREAD,
		alias:  "READ_ONLY",
	},
	{
		legacy: "WRITE_PERMISSION",
		v1:     configv1.MetricsRestrictionPermissionWRITE,
		alias:  "WRITE_ONLY",
	},
	{
		legacy: "READWRITE_PERMISSION",
		v1:     configv1.MetricsRestrictionPermissionREADWRITE,
		alias:  "READ_AND_WRITE",
	},
})
