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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	configv1models "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

const testNotebookJSON = `{"cells":[{"kind":"markdown","text":"hi"}],"version":1}`

func TestNotebookConverterToModel(t *testing.T) {
	got, err := notebookConverter{}.toModel(&intschema.Notebook{
		Name:         "My Notebook",
		Slug:         "my-notebook",
		CollectionId: tfid.Slug("platform"),
		NotebookJson: testNotebookJSON,
	})
	require.NoError(t, err)
	assert.Equal(t, &configmodels.ConfigunstableNotebook{
		Name:           "My Notebook",
		Slug:           "my-notebook",
		CollectionSlug: "platform",
		NotebookJSON:   testNotebookJSON,
	}, got)
}

func TestNotebookConverterToModelErrors(t *testing.T) {
	tests := []struct {
		name     string
		in       *intschema.Notebook
		errorMsg string
	}{
		{
			name: "missing notebook_json",
			in: &intschema.Notebook{
				CollectionId: tfid.Slug("platform"),
			},
			errorMsg: "notebook_json is required",
		},
		{
			name: "invalid notebook_json",
			in: &intschema.Notebook{
				CollectionId: tfid.Slug("platform"),
				NotebookJson: `{"cells":`,
			},
			errorMsg: "invalid notebook_json",
		},
		{
			// The unstable Notebook model has no typed collection reference, so
			// a SERVICE: ID (from the chronosphere_service data source) cannot
			// be represented. It must fail loudly rather than bind the notebook
			// to whatever plain collection shares that slug.
			name: "typed collection reference",
			in: &intschema.Notebook{
				CollectionId: tfid.Slug(CollectionTypeSlugToID(
					configv1models.Configv1CollectionReferenceTypeSERVICE, "checkout")),
				NotebookJson: testNotebookJSON,
			},
			errorMsg: `invalid collection_id "SERVICE:checkout"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := notebookConverter{}.toModel(tt.in)
			assert.ErrorContains(t, err, tt.errorMsg)
		})
	}
}

func TestNotebookConverterFromModel(t *testing.T) {
	got, err := NotebookFromModel(&configmodels.ConfigunstableNotebook{
		Name:           "My Notebook",
		Slug:           "my-notebook",
		CollectionSlug: "platform",
		NotebookJSON:   testNotebookJSON,
	})
	require.NoError(t, err)
	assert.Equal(t, &intschema.Notebook{
		Name:         "My Notebook",
		Slug:         "my-notebook",
		CollectionId: tfid.Slug("platform"),
		NotebookJson: testNotebookJSON,
	}, got)
}

func TestNotebookConverterRoundTrip(t *testing.T) {
	want := &intschema.Notebook{
		Name:         "My Notebook",
		Slug:         "my-notebook",
		CollectionId: tfid.Slug("platform"),
		NotebookJson: testNotebookJSON,
	}

	m, err := notebookConverter{}.toModel(want)
	require.NoError(t, err)
	got, err := notebookConverter{}.fromModel(m)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}
