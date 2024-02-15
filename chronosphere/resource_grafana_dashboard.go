// Copyright 2023 Chronosphere Inc.
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
	"fmt"

	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/grafana"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGrafanaDashboard() *schema.Resource {
	r := newGenericResource[
		*models.Configv1GrafanaDashboard,
		intschema.GrafanaDashboard,
		*intschema.GrafanaDashboard,
	](
		"grafana_dashboard",
		grafanaDashboardConverter{},
		generatedGrafanaDashboard{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: r.ValidateDryRun(&GrafanaDashboardDryRunCount),
		Schema:        tfschema.GrafanaDashboard,
	}
}

// GrafanaDashboardDryRunCount tracks how many times dry run is run during validation for testing.
var GrafanaDashboardDryRunCount atomic.Int64

type grafanaDashboardConverter struct{}

func (grafanaDashboardConverter) toModel(
	d *intschema.GrafanaDashboard,
) (*models.Configv1GrafanaDashboard, error) {
	dashMeta, dashboardJSON, err := GrafanaDashboardExtract(d.DashboardJson)
	if err != nil {
		return nil, err
	}

	// If the provided uid during create is empty, then the server
	// will generate one (as uid mirrors slug). However, unlike
	// slug, uid is embedded in dashboard_json and dashboard_json is
	// not a computed field. To workaround this, the computed
	// sub-fields of dashboard_json are manually stripped via a
	// DiffSuppressFunc, which unfortunately means the server-side
	// generated uid never makes it into the state file. Thus, if
	// uid is missing, we must fallback to using the state file ID
	// as a source-of-truth for slug.
	if dashMeta.Slug == "" {
		dashMeta.Slug = d.StateID
	}

	collSlug, collRef := collectionRefFromID(d.CollectionId.Slug())
	return &models.Configv1GrafanaDashboard{
		Name:           dashMeta.Name,
		Slug:           dashMeta.Slug,
		BucketSlug:     d.BucketId.Slug(),
		CollectionSlug: collSlug,
		Collection:     collRef,
		DashboardJSON:  dashboardJSON,
	}, nil
}

func (grafanaDashboardConverter) fromModel(
	m *models.Configv1GrafanaDashboard,
) (*intschema.GrafanaDashboard, error) {
	dashboardJSON, err := grafana.SanitizedDashboardJSON(string(m.DashboardJSON))
	if err != nil {
		return nil, err
	}
	return &intschema.GrafanaDashboard{
		DashboardJson: dashboardJSON,
		BucketId:      tfid.Slug(m.BucketSlug),
		CollectionId:  tfid.Slug(collectionIDFromRef(m.CollectionSlug, m.Collection)),
	}, nil
}

// GrafanaDashboardFromModel maps an API model to an intschema model.
func GrafanaDashboardFromModel(m *models.Configv1GrafanaDashboard) (*intschema.GrafanaDashboard, error) {
	return grafanaDashboardConverter{}.fromModel(m)
}

// GrafanaDashboardExtract sanitizes the dashboard JSON and extracts DashboardMetadata.
func GrafanaDashboardExtract(dashboardJSON string) (DashboardMetadata, string, error) {
	var dashMeta DashboardMetadata
	sanitizedJSON, err := grafana.SanitizedDashboardJSON(
		dashboardJSON,
		func(dashboard map[string]any) error {
			dashMeta.Name = getJSONFieldStrOrEmpty(dashboard, "title")
			dashMeta.Slug = getJSONFieldStrOrEmpty(dashboard, "uid")

			// title acts as the required name field for the resource.
			if dashMeta.Name == "" {
				return fmt.Errorf("invalid dashboard_json: title required as string")
			}

			return nil
		})
	if err != nil {
		return DashboardMetadata{}, "", err
	}
	return dashMeta, sanitizedJSON, nil
}

// getJSONFieldStrOrEmpty tries to extract the given key from the map as a string.
func getJSONFieldStrOrEmpty(m map[string]any, k string) string {
	v, ok := m[k]
	if !ok {
		return ""
	}

	vStr, ok := v.(string)
	if !ok {
		return ""
	}

	return vStr
}
