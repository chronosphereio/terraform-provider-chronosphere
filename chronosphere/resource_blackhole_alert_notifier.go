package chronosphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func BlackholeAlertNotifierFromModel(
	n *models.Configv1Notifier,
) (*intschema.BlackholeAlertNotifier, error) {
	return blackholeAlertNotifierConverter{}.fromModel(n)
}

func resourceBlackHoleAlertNotifier() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Notifier,
		intschema.BlackholeAlertNotifier,
		*intschema.BlackholeAlertNotifier,
	](
		"blackhole_alert_notifier",
		blackholeAlertNotifierConverter{},
		generatedNotifier{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.BlackholeAlertNotifier,
		CustomizeDiff: r.ValidateDryRun(&BlackHoleAlertNotifierDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// BlackHoleAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var BlackHoleAlertNotifierDryRunCount atomic.Int64

type blackholeAlertNotifierConverter struct{}

func (blackholeAlertNotifierConverter) toModel(
	n *intschema.BlackholeAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Name:    n.Name,
		Slug:    n.Slug,
		Discard: true,
	}, nil
}

func (blackholeAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.BlackholeAlertNotifier, error) {
	if !m.Discard {
		return &intschema.BlackholeAlertNotifier{
			Name: m.Name + notifierTypeChangedName("blackhole"),
			Slug: m.Slug,
		}, nil
	}
	return &intschema.BlackholeAlertNotifier{
		Name: m.Name,
		Slug: m.Slug,
	}, nil
}
