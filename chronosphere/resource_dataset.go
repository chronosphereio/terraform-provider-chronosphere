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
	"github.com/pkg/errors"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	v1models "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/prettyenum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceDataset() *schema.Resource {
	r := newGenericResource[
		*models.ConfigunstableDataset,
		intschema.Dataset,
		*intschema.Dataset,
	]("dataset",
		datasetConverter{},
		generatedUnstableDataset{})

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.Dataset,
		CustomizeDiff: r.ValidateDryRun(&DatasetDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// DatasetDryRunCount tracks how many times dry run is run during validation for testing.
var DatasetDryRunCount atomic.Int64

type datasetConverter struct{}

func (d datasetConverter) toModel(s *intschema.Dataset) (*models.ConfigunstableDataset, error) {
	return &models.ConfigunstableDataset{
		Configuration: datasetConfigurationToModel(s.Configuration),
		Description:   s.Description,
		Name:          s.Name,
		Slug:          s.Slug,
	}, nil
}

func datasetConfigurationToModel(s intschema.DatasetConfiguration) *models.DatasetDatasetConfiguration {
	return &models.DatasetDatasetConfiguration{
		TraceDataset: traceDatasetToModel(s.TraceDataset),
		Type:         prettyenum.DatasetType(s.Type).Model(),
	}
}

func traceDatasetToModel(dataset *intschema.DatasetConfigurationTraceDataset) *models.ConfigunstableTraceDataset {
	if dataset == nil {
		return nil
	}
	configv1Criteria := traceSearchFilterToModel(dataset.MatchCriteria)
	criteria := unstableSearchFilterFromV1(configv1Criteria)
	return &models.ConfigunstableTraceDataset{
		MatchCriteria: criteria,
	}
}

func unstableSearchFilterFromV1(criteria *v1models.Configv1TraceSearchFilter) *models.Configv1TraceSearchFilter {
	if criteria == nil {
		return nil
	}
	serialized, err := criteria.MarshalBinary()
	if err != nil {
		return nil
	}
	result := &models.Configv1TraceSearchFilter{}
	err = result.UnmarshalBinary(serialized)
	if err != nil {
		return nil
	}
	return result
}

func (d datasetConverter) fromModel(m *models.ConfigunstableDataset) (*intschema.Dataset, error) {
	dsConfig, err := datasetConfigurationFromModel(m.Configuration)
	if err != nil {
		return nil, err
	}
	return &intschema.Dataset{
		Name:          m.Name,
		Slug:          m.Slug,
		Configuration: dsConfig,
		Description:   m.Description,
	}, nil
}

func datasetConfigurationFromModel(m *models.DatasetDatasetConfiguration) (intschema.DatasetConfiguration, error) {
	dType := prettyenum.DatasetTypeFromModel(m.Type)
	switch dType {
	case prettyenum.DatasetDatasetTypeTracesModel:
		if m.TraceDataset == nil {
			return intschema.DatasetConfiguration{}, errors.New("when type = %s, trace_dataset must be provided")
		}
		tds, err := traceDatasetFromModel(m.TraceDataset)
		if err != nil {
			return intschema.DatasetConfiguration{}, err
		}
		return intschema.DatasetConfiguration{
			Type:         string(dType),
			TraceDataset: tds,
		}, nil
	default:
		return intschema.DatasetConfiguration{}, errors.Errorf("unsupported dataset type '%s'", dType)
	}
}

func traceDatasetFromModel(m *models.ConfigunstableTraceDataset) (*intschema.DatasetConfigurationTraceDataset, error) {
	configV1, err := v1SearchFilterFromUnstable(m.MatchCriteria)
	if err != nil {
		return nil, err
	}
	return &intschema.DatasetConfigurationTraceDataset{
		MatchCriteria: traceSearchFilterFromModel(configV1),
	}, nil
}

func v1SearchFilterFromUnstable(criteria *models.Configv1TraceSearchFilter) (*v1models.Configv1TraceSearchFilter, error) {
	if criteria == nil {
		return nil, nil
	}
	serialized, err := criteria.MarshalBinary()
	if err != nil {
		return nil, err
	}
	result := &v1models.Configv1TraceSearchFilter{}
	err = result.UnmarshalBinary(serialized)
	if err != nil {
		return nil, err
	}
	return result, nil
}
