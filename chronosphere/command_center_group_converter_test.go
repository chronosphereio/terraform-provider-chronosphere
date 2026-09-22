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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommandCenterGroupToModel(t *testing.T) {
	converter := commandCenterGroupConverter{}

	t.Run("primary only", func(t *testing.T) {
		g := &intschema.CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
				Slug: "checkout-availability",
			},
		}

		m, err := converter.toModel(g)
		require.NoError(t, err)

		require.NotNil(t, m.PrimarySLOReference)
		assert.Equal(t, "checkout-availability", m.PrimarySLOReference.Slug)
		assert.Nil(t, m.GroupSLOReference)
		assert.NotNil(t, m.RelatedSLOReferences)
		assert.Empty(t, m.RelatedSLOReferences)
	})

	t.Run("deprecated group block only", func(t *testing.T) {
		g := &intschema.CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			GroupSloReference: &intschema.CommandCenterGroupGroupSloReference{
				Slug: "checkout-availability",
			},
		}

		m, err := converter.toModel(g)
		require.NoError(t, err)

		require.NotNil(t, m.PrimarySLOReference)
		assert.Equal(t, "checkout-availability", m.PrimarySLOReference.Slug)
		assert.Nil(t, m.GroupSLOReference)
		assert.NotNil(t, m.RelatedSLOReferences)
		assert.Empty(t, m.RelatedSLOReferences)
	})

	t.Run("primary plus two related", func(t *testing.T) {
		g := &intschema.CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
				Slug: "checkout-availability",
			},
			RelatedSloReferences: []intschema.CommandCenterGroupRelatedSloReferences{
				{Slug: "checkout-latency"},
				{Slug: "checkout-payment-success"},
			},
		}

		m, err := converter.toModel(g)
		require.NoError(t, err)

		require.NotNil(t, m.PrimarySLOReference)
		assert.Equal(t, "checkout-availability", m.PrimarySLOReference.Slug)
		require.Len(t, m.RelatedSLOReferences, 2)
		assert.Equal(t, "checkout-latency", m.RelatedSLOReferences[0].Slug)
		assert.Equal(t, "checkout-payment-success", m.RelatedSLOReferences[1].Slug)
	})

	t.Run("neither", func(t *testing.T) {
		g := &intschema.CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
		}

		m, err := converter.toModel(g)
		require.NoError(t, err)

		assert.Nil(t, m.PrimarySLOReference)
		assert.Nil(t, m.GroupSLOReference)
		assert.NotNil(t, m.RelatedSLOReferences)
		assert.Empty(t, m.RelatedSLOReferences)
	})
}

func TestCommandCenterGroupFromModel(t *testing.T) {
	converter := commandCenterGroupConverter{}

	t.Run("primary and related populate the new fields", func(t *testing.T) {
		m := &models.ConfigunstableCommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			PrimarySLOReference: &models.ConfigunstableSLOReference{
				Slug: "checkout-availability",
			},
			RelatedSLOReferences: []*models.ConfigunstableSLOReference{
				{Slug: "checkout-latency"},
				{Slug: "checkout-payment-success"},
			},
		}

		g, err := converter.fromModel(m)
		require.NoError(t, err)

		require.NotNil(t, g.PrimarySloReference)
		assert.Equal(t, "checkout-availability", g.PrimarySloReference.Slug)
		require.Len(t, g.RelatedSloReferences, 2)
		assert.Equal(t, "checkout-latency", g.RelatedSloReferences[0].Slug)
		assert.Equal(t, "checkout-payment-success", g.RelatedSloReferences[1].Slug)
	})

	t.Run("only the deprecated field populates the primary block", func(t *testing.T) {
		m := &models.ConfigunstableCommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			GroupSLOReference: &models.ConfigunstableSLOReference{
				Slug: "checkout-availability",
			},
		}

		g, err := converter.fromModel(m)
		require.NoError(t, err)

		require.NotNil(t, g.PrimarySloReference)
		assert.Equal(t, "checkout-availability", g.PrimarySloReference.Slug)
	})

	t.Run("no related refs leaves the list nil", func(t *testing.T) {
		m := &models.ConfigunstableCommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
		}

		g, err := converter.fromModel(m)
		require.NoError(t, err)

		assert.Nil(t, g.RelatedSloReferences)
	})
}

func TestCommandCenterGroupNormalize(t *testing.T) {
	converter := commandCenterGroupConverter{}

	t.Run("deprecated block in config moves the server value back", func(t *testing.T) {
		config := &intschema.CommandCenterGroup{
			GroupSloReference: &intschema.CommandCenterGroupGroupSloReference{
				Slug: "checkout-availability",
			},
		}
		server := &intschema.CommandCenterGroup{
			PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
				Slug: "checkout-availability",
			},
		}

		converter.normalize(config, server)

		assert.Nil(t, server.PrimarySloReference)
		require.NotNil(t, server.GroupSloReference)
		assert.Equal(t, "checkout-availability", server.GroupSloReference.Slug)
	})

	t.Run("new block in config leaves the server value untouched", func(t *testing.T) {
		config := &intschema.CommandCenterGroup{
			PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
				Slug: "checkout-availability",
			},
		}
		server := &intschema.CommandCenterGroup{
			PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
				Slug: "checkout-availability",
			},
		}

		converter.normalize(config, server)

		require.NotNil(t, server.PrimarySloReference)
		assert.Equal(t, "checkout-availability", server.PrimarySloReference.Slug)
		assert.Nil(t, server.GroupSloReference)
	})

	t.Run("no reference in config leaves the server value untouched", func(t *testing.T) {
		config := &intschema.CommandCenterGroup{}
		server := &intschema.CommandCenterGroup{
			PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
				Slug: "checkout-availability",
			},
		}

		converter.normalize(config, server)

		require.NotNil(t, server.PrimarySloReference)
		assert.Equal(t, "checkout-availability", server.PrimarySloReference.Slug)
		assert.Nil(t, server.GroupSloReference)
	})
}

func TestCommandCenterGroupRoundTrip(t *testing.T) {
	converter := commandCenterGroupConverter{}

	g := &intschema.CommandCenterGroup{
		Name: "checkout",
		Slug: "checkout",
		PrimarySloReference: &intschema.CommandCenterGroupPrimarySloReference{
			Slug: "checkout-availability",
		},
		RelatedSloReferences: []intschema.CommandCenterGroupRelatedSloReferences{
			{Slug: "checkout-latency"},
			{Slug: "checkout-payment-success"},
		},
	}

	m, err := converter.toModel(g)
	require.NoError(t, err)

	result, err := converter.fromModel(m)
	require.NoError(t, err)

	require.NotNil(t, result.PrimarySloReference)
	assert.Equal(t, g.PrimarySloReference.Slug, result.PrimarySloReference.Slug)
	require.Len(t, result.RelatedSloReferences, len(g.RelatedSloReferences))
	for i := range g.RelatedSloReferences {
		assert.Equal(t, g.RelatedSloReferences[i].Slug, result.RelatedSloReferences[i].Slug)
	}
}
