// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	r := newGenericResource(
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
