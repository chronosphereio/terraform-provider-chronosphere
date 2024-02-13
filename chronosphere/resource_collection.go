package chronosphere

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/localid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// CollectionFromModel maps an API model to the intschema model.
func CollectionFromModel(m *models.Configv1Collection) (*intschema.Collection, error) {
	return (collectionConverter{}).fromModel(m)
}

func resourceCollection() *schema.Resource {
	r := newGenericResource[
		*models.Configv1Collection,
		intschema.Collection,
		*intschema.Collection,
	](
		"collection",
		collectionConverter{},
		generatedCollection{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.Collection,
		CustomizeDiff: r.ValidateDryRun(&CollectionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CollectionDryRunCount tracks how many times dry run is run during validation for testing.
var CollectionDryRunCount atomic.Int64

type collectionConverter struct{}

func (collectionConverter) toModel(
	c *intschema.Collection,
) (*models.Configv1Collection, error) {
	if localid.IsLocalID(c.NotificationPolicyId.Slug()) {
		return nil, fmt.Errorf("notification_policy_id must reference a notification policy with name")
	}

	return &models.Configv1Collection{
		Slug:                   c.Slug,
		Name:                   c.Name,
		NotificationPolicySlug: c.NotificationPolicyId.Slug(),
		TeamSlug:               c.TeamId.Slug(),
		Description:            c.Description,
	}, nil
}

func (collectionConverter) fromModel(
	m *models.Configv1Collection,
) (*intschema.Collection, error) {
	return &intschema.Collection{
		Slug:                 m.Slug,
		Name:                 m.Name,
		NotificationPolicyId: tfid.Slug(m.NotificationPolicySlug),
		Description:          m.Description,
		TeamId:               tfid.Slug(m.TeamSlug),
	}, nil
}
