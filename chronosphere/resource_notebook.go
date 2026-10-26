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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
)

// NotebookFromModel maps an API model to an intschema model.
func NotebookFromModel(m *models.Configv1Notebook) (*intschema.Notebook, error) {
	return notebookConverter{}.fromModel(m)
}

func resourceNotebook() *schema.Resource {
	resource := newGenericResource(
		"notebook",
		notebookConverter{},
		generatedNotebook{})

	return &schema.Resource{
		CreateContext: resource.CreateContext,
		ReadContext:   resource.ReadContext,
		UpdateContext: resource.UpdateContext,
		DeleteContext: resource.DeleteContext,
		Description:   "A Chronosphere notebook, a document of text, queries, and panels defined by `notebook_json` and owned by a collection. The API rejects edits from the Chronosphere app to a notebook that Terraform manages.",
		Schema:        tfschema.Notebook,
		CustomizeDiff: resource.ValidateDryRunOptions(&NotebookDryRunCount, ValidateDryRunOpts[*models.Configv1Notebook]{
			// notebook_json is unknown at plan time when it interpolates an attribute of a
			// resource that does not exist yet; dry-run against a placeholder document rather
			// than failing the plan with a misleading "notebook_json is required".
			DryRunDefaults: map[string]any{
				"notebook_json": "{}",
			},
		}),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// NotebookDryRunCount tracks how many times dry run is run during validation for testing.
var NotebookDryRunCount atomic.Int64

type notebookConverter struct{}

func (notebookConverter) toModel(
	n *intschema.Notebook,
) (*models.Configv1Notebook, error) {
	if n.NotebookJson == "" {
		return nil, errors.New("notebook_json is required")
	}

	// The API stores notebook_json as an opaque blob and only checks that it
	// is a JSON object, so check the same thing here for a plan-time error.
	var notebook map[string]any
	if err := xjson.Unmarshal([]byte(n.NotebookJson), &notebook); err != nil {
		return nil, fmt.Errorf("invalid notebook_json: %s", err)
	}
	// A JSON null unmarshals into a nil map without error.
	if notebook == nil {
		return nil, errors.New("invalid notebook_json: must be a JSON object")
	}

	collSlug, collRef := collectionRefFromID(n.CollectionId.Slug())
	return &models.Configv1Notebook{
		Name:           n.Name,
		Slug:           n.Slug,
		CollectionSlug: collSlug,
		Collection:     collRef,
		NotebookJSON:   n.NotebookJson,
		Labels:         n.Labels,
	}, nil
}

func (notebookConverter) fromModel(
	m *models.Configv1Notebook,
) (*intschema.Notebook, error) {
	return &intschema.Notebook{
		Name:         m.Name,
		Slug:         m.Slug,
		NotebookJson: m.NotebookJSON,
		CollectionId: tfid.Slug(collectionIDFromRef(m.CollectionSlug, m.Collection)),
		Labels:       m.Labels,
	}, nil
}
