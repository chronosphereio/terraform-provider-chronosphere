package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/dashboard"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Dashboard = map[string]*schema.Schema{
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"collection_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"dashboard_json": {
		Type:                  schema.TypeString,
		Required:              true,
		DiffSuppressFunc:      dashboardJSONDiffSuppress,
		DiffSuppressOnRefresh: true,
	},
}

// dashboardJSONDiffSuppress sanitizes and then diffs two dashboard JSON payloads.
func dashboardJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	sanitizedOld, err := dashboard.SanitizedDashboardJSON(old)
	if err != nil {
		return false
	}

	sanitizedNew, err := dashboard.SanitizedDashboardJSON(new)
	if err != nil {
		return false
	}

	return sanitizedOld == sanitizedNew
}
