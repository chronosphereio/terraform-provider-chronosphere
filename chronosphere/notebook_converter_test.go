package chronosphere

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

const testNotebookJSON = `{"version":1,"document":{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"hello"}]}]}}`

func TestNotebookConverterToModel(t *testing.T) {
	tests := []struct {
		name    string
		in      *intschema.Notebook
		want    *models.Configv1Notebook
		wantErr string
	}{
		{
			name: "simple collection by slug",
			in: &intschema.Notebook{
				Name:         "On-call runbook",
				Slug:         "on-call-runbook",
				CollectionId: tfid.Slug("platform"),
				NotebookJson: testNotebookJSON,
				Labels:       map[string]string{"team": "platform"},
			},
			want: &models.Configv1Notebook{
				Name:           "On-call runbook",
				Slug:           "on-call-runbook",
				CollectionSlug: "platform",
				NotebookJSON:   testNotebookJSON,
				Labels:         map[string]string{"team": "platform"},
			},
		},
		{
			name: "user collection by typed id",
			in: &intschema.Notebook{
				Name:         "Scratch",
				CollectionId: tfid.Slug("USER:alice"),
				NotebookJson: testNotebookJSON,
			},
			want: &models.Configv1Notebook{
				Name: "Scratch",
				Collection: &models.Configv1CollectionReference{
					Type: models.Configv1CollectionReferenceTypeUSER,
					Slug: "alice",
				},
				NotebookJSON: testNotebookJSON,
			},
		},
		{
			name: "missing notebook_json",
			in: &intschema.Notebook{
				Name:         "Empty",
				CollectionId: tfid.Slug("platform"),
			},
			wantErr: "notebook_json is required",
		},
		{
			name: "notebook_json is not an object",
			in: &intschema.Notebook{
				Name:         "Array",
				CollectionId: tfid.Slug("platform"),
				NotebookJson: `[1,2,3]`,
			},
			wantErr: "invalid notebook_json",
		},
		{
			name: "notebook_json is null",
			in: &intschema.Notebook{
				Name:         "Null",
				CollectionId: tfid.Slug("platform"),
				NotebookJson: `null`,
			},
			wantErr: "invalid notebook_json: must be a JSON object",
		},
		{
			name: "notebook_json is not JSON",
			in: &intschema.Notebook{
				Name:         "Broken",
				CollectionId: tfid.Slug("platform"),
				NotebookJson: `{"version":`,
			},
			wantErr: "invalid notebook_json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := notebookConverter{}.toModel(tt.in)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNotebookConverterFromModel(t *testing.T) {
	tests := []struct {
		name string
		in   *models.Configv1Notebook
		want *intschema.Notebook
	}{
		{
			name: "simple collection by slug",
			in: &models.Configv1Notebook{
				Name:           "On-call runbook",
				Slug:           "on-call-runbook",
				CollectionSlug: "platform",
				NotebookJSON:   testNotebookJSON,
				Labels:         map[string]string{"team": "platform"},
			},
			want: &intschema.Notebook{
				Name:         "On-call runbook",
				Slug:         "on-call-runbook",
				CollectionId: tfid.Slug("platform"),
				NotebookJson: testNotebookJSON,
				Labels:       map[string]string{"team": "platform"},
			},
		},
		{
			name: "user collection reference",
			in: &models.Configv1Notebook{
				Name: "Scratch",
				Slug: "scratch",
				Collection: &models.Configv1CollectionReference{
					Type: models.Configv1CollectionReferenceTypeUSER,
					Slug: "alice",
				},
				NotebookJSON: testNotebookJSON,
			},
			want: &intschema.Notebook{
				Name:         "Scratch",
				Slug:         "scratch",
				CollectionId: tfid.Slug("USER:alice"),
				NotebookJson: testNotebookJSON,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NotebookFromModel(tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNotebookConverterRoundTrip(t *testing.T) {
	in := &intschema.Notebook{
		Name:         "On-call runbook",
		Slug:         "on-call-runbook",
		CollectionId: tfid.Slug("USER:alice"),
		NotebookJson: testNotebookJSON,
		Labels:       map[string]string{"team": "platform"},
	}

	m, err := notebookConverter{}.toModel(in)
	require.NoError(t, err)

	out, err := notebookConverter{}.fromModel(m)
	require.NoError(t, err)
	assert.Equal(t, in, out)
}
