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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// CommandCenterGroupFromModel maps an API model to the intschema model.
func CommandCenterGroupFromModel(m *models.Configv1CommandCenterGroup) (*intschema.CommandCenterGroup, error) {
	return (commandCenterGroupConverter{}).fromModel(m)
}

func resourceCommandCenterGroup() *schema.Resource {
	r := newGenericResource(
		"command_center_group",
		commandCenterGroupConverter{},
		generatedCommandCenterGroup{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description:   "A named group of signals tracked in the command center.",
		Schema:        tfschema.CommandCenterGroup,
		CustomizeDiff: r.ValidateDryRun(&CommandCenterGroupDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CommandCenterGroupDryRunCount tracks how many times dry run is run during validation for testing.
var CommandCenterGroupDryRunCount atomic.Int64

type commandCenterGroupConverter struct{}

func (commandCenterGroupConverter) toModel(
	g *intschema.CommandCenterGroup,
) (*models.Configv1CommandCenterGroup, error) {
	m := &models.Configv1CommandCenterGroup{
		Name:                 g.Name,
		Slug:                 g.Slug,
		RelatedSLOReferences: []*models.Configv1SLOReference{},
	}
	if g.PrimarySloReference != nil {
		m.PrimarySLOReference = &models.Configv1SLOReference{
			Slug: g.PrimarySloReference.Slug,
		}
	}
	if len(g.RelatedSloReferences) > 0 {
		m.RelatedSLOReferences = sliceutil.Map(
			g.RelatedSloReferences,
			func(r intschema.CommandCenterGroupRelatedSloReferences) *models.Configv1SLOReference {
				return &models.Configv1SLOReference{Slug: r.Slug}
			},
		)
	}
	return m, nil
}

func (commandCenterGroupConverter) fromModel(
	m *models.Configv1CommandCenterGroup,
) (*intschema.CommandCenterGroup, error) {
	g := &intschema.CommandCenterGroup{
		Name: m.Name,
		Slug: m.Slug,
	}
	if m.PrimarySLOReference != nil {
		g.PrimarySloReference = &intschema.CommandCenterGroupPrimarySloReference{
			Slug: m.PrimarySLOReference.Slug,
		}
	}
	if len(m.RelatedSLOReferences) > 0 {
		g.RelatedSloReferences = sliceutil.Map(
			m.RelatedSLOReferences,
			func(r *models.Configv1SLOReference) intschema.CommandCenterGroupRelatedSloReferences {
				return intschema.CommandCenterGroupRelatedSloReferences{Slug: r.Slug}
			},
		)
	}
	return g, nil
}
