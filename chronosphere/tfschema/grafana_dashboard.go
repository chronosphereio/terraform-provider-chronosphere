package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/grafana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var GrafanaDashboard = map[string]*schema.Schema{
	"bucket_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
	},
	"collection_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
	},
	"dashboard_json": {
		Type:             schema.TypeString,
		Required:         true,
		DiffSuppressFunc: grafanaDashboardJSONDiffSuppress,
	},
}

// grafanaDashboardJSONDiffSuppress sanitizes and then diffs two dashboard JSON payloads.
func grafanaDashboardJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	sanitizedOld, err := grafana.SanitizedDashboardJSON(old, grafana.WithUID(""))
	if err != nil {
		return false
	}

	sanitizedNew, err := grafana.SanitizedDashboardJSON(new, grafana.WithUID(""))
	if err != nil {
		return false
	}

	return sanitizedOld == sanitizedNew
}
