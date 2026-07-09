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
		Description:   "A named group of signals tracked in the command center. This resource is backed by Chronosphere's unstable config API and is subject to breaking change without notice.",
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
		Name: g.Name,
		Slug: g.Slug,
	}
	if g.GroupSloReference != nil {
		m.GroupSLOReference = &models.ConfigunstableSLOReference{
			Slug: g.GroupSloReference.Slug,
		}
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
	if m.GroupSLOReference != nil {
		g.GroupSloReference = &intschema.CommandCenterGroupGroupSloReference{
			Slug: m.GroupSLOReference.Slug,
		}
	}
	return g, nil
}
