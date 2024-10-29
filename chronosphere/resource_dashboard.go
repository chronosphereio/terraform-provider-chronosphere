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
	"errors"
	"fmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"
)

// DashboardFromModel maps an API model to an intschema model.
func DashboardFromModel(m *models.Configv1Dashboard) (*intschema.Dashboard, error) {
	return dashboardConverter{}.fromModel(m)
}

func resourceDashboard() *schema.Resource {
	resource := newGenericResource(
		"dashboard",
		dashboardConverter{},
		generatedDashboard{})

	return &schema.Resource{
		CreateContext: resource.CreateContext,
		ReadContext:   resource.ReadContext,
		UpdateContext: resource.UpdateContext,
		DeleteContext: resource.DeleteContext,
		Schema:        tfschema.Dashboard,
		CustomizeDiff: resource.ValidateDryRun(&DashboardDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// DashboardDryRunCount tracks how many times dry run is run during validation for testing.
var DashboardDryRunCount atomic.Int64

type dashboardConverter struct{}

func (dashboardConverter) toModel(
	d *intschema.Dashboard,
) (*models.Configv1Dashboard, error) {
	if d.DashboardJson == "" {
		return nil, errors.New("dashboard_json is required")
	}

	var dashboard map[string]any
	if err := xjson.Unmarshal([]byte(d.DashboardJson), &dashboard); err != nil {
		return nil, fmt.Errorf("invalid dashboard_json: %s", err)
	}

	collSlug, collRef := collectionRefFromID(d.CollectionId.Slug())
	return &models.Configv1Dashboard{
		Name:           d.Name,
		Slug:           d.Slug,
		CollectionSlug: collSlug,
		Collection:     collRef,
		DashboardJSON:  d.DashboardJson,
		Labels:         d.Labels,
	}, nil
}

func (dashboardConverter) fromModel(
	m *models.Configv1Dashboard,
) (*intschema.Dashboard, error) {
	return &intschema.Dashboard{
		Name:          m.Name,
		Slug:          m.Slug,
		DashboardJson: m.DashboardJSON,
		CollectionId:  tfid.Slug(collectionIDFromRef(m.CollectionSlug, m.Collection)),
		Labels:        m.Labels,
	}, nil
}
