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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceCommandCenterGroup() *schema.Resource {
	r := newGenericResource(
		"command_center_group",
		commandCenterGroupConverter{},
		generatedUnstableCommandCenterGroup{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description:   "A named group of signals tracked in the command center. " + unstableAPIWarning,
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
) (*models.ConfigunstableCommandCenterGroup, error) {
	m := &models.ConfigunstableCommandCenterGroup{
		Name:                 g.Name,
		Slug:                 g.Slug,
		RelatedSLOReferences: []*models.ConfigunstableSLOReference{},
	}
	primary := g.PrimarySloReference
	if primary == nil && g.GroupSloReference != nil {
		primary = &intschema.CommandCenterGroupPrimarySloReference{
			Slug: g.GroupSloReference.Slug,
		}
	}
	if primary != nil {
		m.PrimarySLOReference = &models.ConfigunstableSLOReference{
			Slug: primary.Slug,
		}
	}
	if len(g.RelatedSloReferences) > 0 {
		m.RelatedSLOReferences = sliceutil.Map(
			g.RelatedSloReferences,
			func(r intschema.CommandCenterGroupRelatedSloReferences) *models.ConfigunstableSLOReference {
				return &models.ConfigunstableSLOReference{Slug: r.Slug}
			},
		)
	}
	return m, nil
}

func (commandCenterGroupConverter) fromModel(
	m *models.ConfigunstableCommandCenterGroup,
) (*intschema.CommandCenterGroup, error) {
	g := &intschema.CommandCenterGroup{
		Name: m.Name,
		Slug: m.Slug,
	}
	primary := m.PrimarySLOReference
	if primary == nil {
		primary = m.GroupSLOReference
	}
	if primary != nil {
		g.PrimarySloReference = &intschema.CommandCenterGroupPrimarySloReference{
			Slug: primary.Slug,
		}
	}
	if len(m.RelatedSLOReferences) > 0 {
		g.RelatedSloReferences = sliceutil.Map(
			m.RelatedSLOReferences,
			func(r *models.ConfigunstableSLOReference) intschema.CommandCenterGroupRelatedSloReferences {
				return intschema.CommandCenterGroupRelatedSloReferences{Slug: r.Slug}
			},
		)
	}
	return g, nil
}

// The Terraform schema has an equivalent "primary_slo_reference" and a
// deprecated "group_slo_reference" that are the same field on the server.
// Take the server value and fold it into whichever block the user has
// decided to set, avoiding a meaningless diff during "terraform plan".
func (commandCenterGroupConverter) normalize(config, server *intschema.CommandCenterGroup) {
	if config.GroupSloReference != nil && config.PrimarySloReference == nil {
		if server.PrimarySloReference != nil {
			server.GroupSloReference = &intschema.CommandCenterGroupGroupSloReference{
				Slug: server.PrimarySloReference.Slug,
			}
		}
		server.PrimarySloReference = nil
	}
}
