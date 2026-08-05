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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
)

// NotebookFromModel maps an API model to an intschema model.
func NotebookFromModel(m *models.ConfigunstableNotebook) (*intschema.Notebook, error) {
	return notebookConverter{}.fromModel(m)
}

func resourceNotebook() *schema.Resource {
	r := newGenericResource(
		"notebook",
		notebookConverter{},
		generatedUnstableNotebook{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description:   "A Chronosphere notebook, whose cells and content are defined by `notebook_json`. This resource is backed by Chronosphere's unstable config API and is subject to breaking change without notice.",
		Schema:        tfschema.Notebook,
		// No CustomizeDiff: the notebook API has no dry_run, so the registry
		// entry sets DisableDryRun and there is no plan-time validation to run.
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

type notebookConverter struct{}

func (notebookConverter) toModel(
	n *intschema.Notebook,
) (*models.ConfigunstableNotebook, error) {
	if n.NotebookJson == "" {
		return nil, errors.New("notebook_json is required")
	}

	var notebook any
	if err := xjson.Unmarshal([]byte(n.NotebookJson), &notebook); err != nil {
		return nil, fmt.Errorf("invalid notebook_json: %s", err)
	}

	// The unstable Notebook model carries a bare collection slug and has no
	// embedded collection reference like the v1 Dashboard does, so a typed ID
	// (e.g. the SERVICE:<slug> produced by the chronosphere_service data
	// source) cannot be represented. Reject it rather than silently dropping
	// the type and binding the notebook to an unrelated collection.
	collID := n.CollectionId.Slug()
	if collType, _, typed := CollectionTypeSlugFromID(collID); typed {
		return nil, fmt.Errorf(
			"invalid collection_id %q: notebooks only support plain collection slugs, not %s references",
			collID, collType)
	}

	return &models.ConfigunstableNotebook{
		Name:           n.Name,
		Slug:           n.Slug,
		CollectionSlug: collID,
		NotebookJSON:   n.NotebookJson,
	}, nil
}

func (notebookConverter) fromModel(
	m *models.ConfigunstableNotebook,
) (*intschema.Notebook, error) {
	return &intschema.Notebook{
		Name:         m.Name,
		Slug:         m.Slug,
		NotebookJson: m.NotebookJSON,
		// Always a bare slug: the model has no typed collection reference, and
		// toModel rejects typed IDs on the way in.
		CollectionId: tfid.Slug(m.CollectionSlug),
	}, nil
}
