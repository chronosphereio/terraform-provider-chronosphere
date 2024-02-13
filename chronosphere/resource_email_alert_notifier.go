package chronosphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func EmailAlertNotifierFromModel(
	m *models.Configv1Notifier,
) (*intschema.EmailAlertNotifier, error) {
	return emailAlertNotifierConverter{}.fromModel(m)
}

func resourceEmailAlertNotifier() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Notifier,
		intschema.EmailAlertNotifier,
		*intschema.EmailAlertNotifier,
	](
		"email_alert_notifier",
		emailAlertNotifierConverter{},
		generatedNotifier{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&EmailAlertNotifierDryRunCount),
		Schema:        tfschema.EmailAlertNotifier,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// EmailAlertNotifierDryRunCount tracks how many times dry run is run during validation for testing.
var EmailAlertNotifierDryRunCount atomic.Int64

type emailAlertNotifierConverter struct{}

func (emailAlertNotifierConverter) toModel(
	n *intschema.EmailAlertNotifier,
) (*models.Configv1Notifier, error) {
	return &models.Configv1Notifier{
		Slug:         n.Slug,
		Name:         n.Name,
		SkipResolved: !n.SendResolved,
		Email: &models.NotifierEmailConfig{
			HTML: n.Html,
			Text: n.Text,
			To:   n.To,
		},
	}, nil
}

func (emailAlertNotifierConverter) fromModel(
	m *models.Configv1Notifier,
) (*intschema.EmailAlertNotifier, error) {
	if m.Email == nil {
		return &intschema.EmailAlertNotifier{
			Name: m.Name + notifierTypeChangedName("email"),
			Slug: m.Slug,
		}, nil
	}

	return &intschema.EmailAlertNotifier{
		Slug:         m.Slug,
		Name:         m.Name,
		SendResolved: !m.SkipResolved,
		Html:         m.Email.HTML,
		Text:         m.Email.Text,
		To:           m.Email.To,
	}, nil
}
