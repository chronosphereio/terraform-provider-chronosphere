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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"

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

	t.Run("related only", func(t *testing.T) {
		g := &intschema.CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			RelatedSloReferences: []intschema.CommandCenterGroupRelatedSloReferences{
				{Slug: "checkout-latency"},
			},
		}

		m, err := converter.toModel(g)
		require.NoError(t, err)

		assert.Nil(t, m.PrimarySLOReference)
		require.Len(t, m.RelatedSLOReferences, 1)
		assert.Equal(t, "checkout-latency", m.RelatedSLOReferences[0].Slug)
	})

	t.Run("no references sends an empty related list", func(t *testing.T) {
		g := &intschema.CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
		}

		m, err := converter.toModel(g)
		require.NoError(t, err)

		assert.Nil(t, m.PrimarySLOReference)
		assert.NotNil(t, m.RelatedSLOReferences)
		assert.Empty(t, m.RelatedSLOReferences)
	})
}

func TestCommandCenterGroupFromModel(t *testing.T) {
	converter := commandCenterGroupConverter{}

	t.Run("primary and related", func(t *testing.T) {
		m := &models.Configv1CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
			PrimarySLOReference: &models.Configv1SLOReference{
				Slug: "checkout-availability",
			},
			RelatedSLOReferences: []*models.Configv1SLOReference{
				{Slug: "checkout-latency"},
				{Slug: "checkout-payment-success"},
			},
		}

		g, err := converter.fromModel(m)
		require.NoError(t, err)

		assert.Equal(t, "checkout", g.Name)
		assert.Equal(t, "checkout", g.Slug)
		require.NotNil(t, g.PrimarySloReference)
		assert.Equal(t, "checkout-availability", g.PrimarySloReference.Slug)
		require.Len(t, g.RelatedSloReferences, 2)
		assert.Equal(t, "checkout-latency", g.RelatedSloReferences[0].Slug)
		assert.Equal(t, "checkout-payment-success", g.RelatedSloReferences[1].Slug)
	})

	t.Run("no references leaves both unset", func(t *testing.T) {
		m := &models.Configv1CommandCenterGroup{
			Name: "checkout",
			Slug: "checkout",
		}

		g, err := converter.fromModel(m)
		require.NoError(t, err)

		assert.Nil(t, g.PrimarySloReference)
		assert.Nil(t, g.RelatedSloReferences)
	})

	t.Run("empty related list leaves the list nil", func(t *testing.T) {
		m := &models.Configv1CommandCenterGroup{
			Name:                 "checkout",
			Slug:                 "checkout",
			RelatedSLOReferences: []*models.Configv1SLOReference{},
		}

		g, err := converter.fromModel(m)
		require.NoError(t, err)

		assert.Nil(t, g.RelatedSloReferences)
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

	assert.Equal(t, g, result)
}
