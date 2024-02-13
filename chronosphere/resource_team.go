package chronosphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// TeamFromModel maps an API model to the intschema model.
func TeamFromModel(m *models.Configv1Team) (*intschema.Team, error) {
	return teamConverter{}.fromModel(m)
}

func resourceTeam() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Team,
		intschema.Team,
		*intschema.Team,
	](
		"team",
		teamConverter{},
		generatedTeam{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.Team,
		CustomizeDiff: r.ValidateDryRun(&TeamDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// TeamDryRunCount tracks how many times dry run is run during validation for testing.
var TeamDryRunCount atomic.Int64

type teamConverter struct{}

func (teamConverter) toModel(
	t *intschema.Team,
) (*models.Configv1Team, error) {
	return &models.Configv1Team{
		Name:        t.Name,
		Slug:        t.Slug,
		Description: t.Description,
		UserEmails:  t.UserEmails,
	}, nil
}

func (teamConverter) fromModel(
	t *models.Configv1Team,
) (*intschema.Team, error) {
	return &intschema.Team{
		Name:        t.Name,
		Slug:        t.Slug,
		Description: t.Description,
		UserEmails:  t.UserEmails,
	}, nil
}
