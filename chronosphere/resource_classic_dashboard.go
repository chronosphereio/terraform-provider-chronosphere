package chronosphere

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"
)

// ClassicDashboardDryRunCount tracks how many times dry run is run during validation for testing.
var ClassicDashboardDryRunCount atomic.Int64

func resourceClassicDashboard() *schema.Resource {
	r := newGenericResource[
		*models.Configv1GrafanaDashboard,
		intschema.GrafanaDashboard,
		*intschema.GrafanaDashboard,
	](
		"classic_dashboard",
		grafanaDashboardConverter{},
		generatedClassicDashboard{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: r.ValidateDryRun(&ClassicDashboardDryRunCount),
		Schema:        tfschema.GrafanaDashboard,
	}
}
