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
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DerivedLabelFromModel maps an API model to an intschema model.
func DerivedLabelFromModel(m *models.Configv1DerivedLabel) (*intschema.DerivedLabel, error) {
	return derivedLabelConverter{}.fromModel(m)
}

func resourceDerivedLabel() *schema.Resource {
	r := newGenericResource(
		"derived_label",
		derivedLabelConverter{},
		generatedDerivedLabel{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.DerivedLabel,
		CustomizeDiff: r.ValidateDryRun(&DerivedLabelDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// DerivedLabelDryRunCount tracks how many times dry run is run during validation for testing.
var DerivedLabelDryRunCount atomic.Int64

type derivedLabelConverter struct{}

func (derivedLabelConverter) toModel(
	m *intschema.DerivedLabel,
) (*models.Configv1DerivedLabel, error) {
	return &models.Configv1DerivedLabel{
		Name:                m.Name,
		Slug:                m.Slug,
		LabelName:           m.LabelName,
		MetricLabel:         derivedLabelMetricLabelToModel(m.MetricLabel),
		ExistingLabelPolicy: enum.LabelPolicy.V1(m.ExistingLabelPolicy),
		Description:         m.Description,
	}, nil
}

func (derivedLabelConverter) fromModel(
	m *models.Configv1DerivedLabel,
) (*intschema.DerivedLabel, error) {
	return &intschema.DerivedLabel{
		Name:                m.Name,
		Slug:                m.Slug,
		LabelName:           m.LabelName,
		MetricLabel:         derivedLabelMetricLabelFromModel(m.MetricLabel),
		ExistingLabelPolicy: string(m.ExistingLabelPolicy),
		Description:         m.Description,
	}, nil
}

func derivedLabelMetricLabelFromModel(
	d *models.DerivedLabelMetricLabel,
) *intschema.DerivedLabelMetricLabel {
	if d == nil {
		return nil
	}
	return &intschema.DerivedLabelMetricLabel{
		ConstructedLabel: constructedLabelFromModel(d.ConstructedLabel),
		MappingLabel:     mappingLabelFromModel(d.MappingLabel),
	}
}

func constructedLabelFromModel(
	c *models.MetricLabelConstructedLabel,
) *intschema.DerivedLabelMetricLabelConstructedLabel {
	if c == nil {
		return nil
	}
	return &intschema.DerivedLabelMetricLabelConstructedLabel{
		ValueDefinitions: sliceutil.Map(c.ValueDefinitions, valueDefinitionFromModel),
	}
}

func valueDefinitionFromModel(
	v *models.ConstructedLabelValueDefinition,
) intschema.DerivedLabelMetricLabelConstructedLabelValueDefinitions {
	return intschema.DerivedLabelMetricLabelConstructedLabelValueDefinitions{
		Filters: sliceutil.Map(v.Filters, constructedLabelFilterFromModel),
		Value:   v.Value,
	}
}

func constructedLabelFilterFromModel(
	f *models.Configv1LabelFilter,
) intschema.DerivedLabelMetricLabelConstructedLabelValueDefinitionsFilters {
	return intschema.DerivedLabelMetricLabelConstructedLabelValueDefinitionsFilters{
		Name:      f.Name,
		ValueGlob: f.ValueGlob,
	}
}

func mappingLabelFromModel(
	c *models.MetricLabelMappingLabel,
) *intschema.DerivedLabelMetricLabelMappingLabel {
	if c == nil {
		return nil
	}
	return &intschema.DerivedLabelMetricLabelMappingLabel{
		NameMappings:  sliceutil.Map(c.NameMappings, nameMappingFromModel),
		ValueMappings: sliceutil.Map(c.ValueMappings, valueMappingFromModel),
	}
}

func nameMappingFromModel(
	v *models.MappingLabelNameMapping,
) intschema.DerivedLabelMetricLabelMappingLabelNameMappings {
	return intschema.DerivedLabelMetricLabelMappingLabelNameMappings{
		Filters:       sliceutil.Map(v.Filters, mappingLabelFilterFromModel),
		SourceLabel:   v.SourceLabel,
		ValueMappings: sliceutil.Map(v.ValueMappings, valueMappingFromModel),
	}
}

func valueMappingFromModel(
	v *models.MappingLabelValueMapping,
) intschema.ValueMappings {
	return intschema.ValueMappings{
		SourceValueGlobs: v.SourceValueGlobs,
		TargetValue:      v.TargetValue,
	}
}

func mappingLabelFilterFromModel(
	f *models.Configv1LabelFilter,
) intschema.DerivedLabelMetricLabelMappingLabelNameMappingsFilters {
	return intschema.DerivedLabelMetricLabelMappingLabelNameMappingsFilters{
		Name:      f.Name,
		ValueGlob: f.ValueGlob,
	}
}

func derivedLabelMetricLabelToModel(
	d *intschema.DerivedLabelMetricLabel,
) *models.DerivedLabelMetricLabel {
	if d == nil {
		return nil
	}
	return &models.DerivedLabelMetricLabel{
		ConstructedLabel: constructedLabelToModel(d.ConstructedLabel),
		MappingLabel:     mappingLabelToModel(d.MappingLabel),
	}
}

func constructedLabelToModel(
	c *intschema.DerivedLabelMetricLabelConstructedLabel,
) *models.MetricLabelConstructedLabel {
	if c == nil {
		return nil
	}
	return &models.MetricLabelConstructedLabel{
		ValueDefinitions: sliceutil.Map(c.ValueDefinitions, valueDefinitionToModel),
	}
}

func valueDefinitionToModel(
	v intschema.DerivedLabelMetricLabelConstructedLabelValueDefinitions,
) *models.ConstructedLabelValueDefinition {
	return &models.ConstructedLabelValueDefinition{
		Filters: sliceutil.Map(v.Filters, constructedLabelFilterToModel),
		Value:   v.Value,
	}
}

func constructedLabelFilterToModel(
	f intschema.DerivedLabelMetricLabelConstructedLabelValueDefinitionsFilters,
) *models.Configv1LabelFilter {
	return &models.Configv1LabelFilter{
		Name:      f.Name,
		ValueGlob: f.ValueGlob,
	}
}

func mappingLabelToModel(
	c *intschema.DerivedLabelMetricLabelMappingLabel,
) *models.MetricLabelMappingLabel {
	if c == nil {
		return nil
	}
	return &models.MetricLabelMappingLabel{
		NameMappings:  sliceutil.Map(c.NameMappings, nameMappingToModel),
		ValueMappings: sliceutil.Map(c.ValueMappings, valueMappingToModel),
	}
}

func nameMappingToModel(
	v intschema.DerivedLabelMetricLabelMappingLabelNameMappings,
) *models.MappingLabelNameMapping {
	return &models.MappingLabelNameMapping{
		Filters:       sliceutil.Map(v.Filters, mappingLabelFilterToModel),
		SourceLabel:   v.SourceLabel,
		ValueMappings: sliceutil.Map(v.ValueMappings, valueMappingToModel),
	}
}

func valueMappingToModel(
	v intschema.ValueMappings,
) *models.MappingLabelValueMapping {
	return &models.MappingLabelValueMapping{
		SourceValueGlobs: v.SourceValueGlobs,
		TargetValue:      v.TargetValue,
	}
}

func mappingLabelFilterToModel(
	f intschema.DerivedLabelMetricLabelMappingLabelNameMappingsFilters,
) *models.Configv1LabelFilter {
	return &models.Configv1LabelFilter{
		Name:      f.Name,
		ValueGlob: f.ValueGlob,
	}
}
