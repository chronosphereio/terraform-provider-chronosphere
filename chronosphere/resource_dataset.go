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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/prettyenum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// DatasetFromModel maps an API model to an intschema model.
func DatasetFromModel(m *models.Configv1Dataset) (*intschema.Dataset, error) {
	return datasetConverter{}.fromModel(m)
}

func resourceDataset() *schema.Resource {
	r := newGenericResource("dataset",
		datasetConverter{},
		generatedDataset{})

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

func (d datasetConverter) toModel(s *intschema.Dataset) (*models.Configv1Dataset, error) {
	cfg, err := datasetConfigurationToModel(s.Configuration)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &models.Configv1Dataset{
		Configuration: cfg,
		Description:   s.Description,
		Name:          s.Name,
		Slug:          s.Slug,
	}, nil
}

func datasetConfigurationToModel(s intschema.DatasetConfiguration) (*models.DatasetDatasetConfiguration, error) {
	t := prettyenum.DatasetType(s.Type)
	cfg := &models.DatasetDatasetConfiguration{
		Type: t.Model(),
	}

	switch t {
	case prettyenum.DatasetDatasetTypeTracesModel:
		var err error
		cfg.TraceDataset, err = traceDatasetToModel(s.TraceDataset)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	case prettyenum.DatasetDatasetTypeLogsModel:
		cfg.LogDataset = logDatasetToModel(s.LogDataset)
	}

	return cfg, nil
}

func traceDatasetToModel(dataset *intschema.DatasetConfigurationTraceDataset) (*models.Configv1TraceDataset, error) {
	if dataset == nil {
		return nil, nil
	}
	filter, err := traceSearchFilterToModel(dataset.MatchCriteria)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &models.Configv1TraceDataset{
		MatchCriteria: filter,
	}, nil
}

func logDatasetToModel(dataset *intschema.DatasetConfigurationLogDataset) *models.Configv1LogDataset {
	if dataset == nil {
		return nil
	}
	return &models.Configv1LogDataset{
		MatchCriteria: &models.Configv1LogSearchFilter{
			Query: dataset.MatchCriteria.Query,
		},
	}
}

func (d datasetConverter) fromModel(m *models.Configv1Dataset) (*intschema.Dataset, error) {
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
	cfg := intschema.DatasetConfiguration{
		Type: string(dType),
	}

	switch dType {
	case prettyenum.DatasetDatasetTypeTracesModel:
		if m.TraceDataset == nil {
			return intschema.DatasetConfiguration{}, errors.Errorf("when type = %s, trace_dataset must be provided", dType)
		}
		tds, err := traceDatasetFromModel(m.TraceDataset)
		if err != nil {
			return intschema.DatasetConfiguration{}, err
		}
		cfg.TraceDataset = tds
	case prettyenum.DatasetDatasetTypeLogsModel:
		if m.LogDataset == nil {
			return intschema.DatasetConfiguration{}, errors.Errorf("when type = %s, log_dataset must be provided", dType)
		}
		cfg.LogDataset = logDatasetFromModel(m.LogDataset)
	default:
		return intschema.DatasetConfiguration{}, errors.Errorf("unsupported dataset type '%s'", dType)
	}

	return cfg, nil
}

func traceDatasetFromModel(m *models.Configv1TraceDataset) (*intschema.DatasetConfigurationTraceDataset, error) {
	return &intschema.DatasetConfigurationTraceDataset{
		MatchCriteria: traceSearchFilterFromModel(m.MatchCriteria),
	}, nil
}

func logDatasetFromModel(m *models.Configv1LogDataset) *intschema.DatasetConfigurationLogDataset {
	return &intschema.DatasetConfigurationLogDataset{
		MatchCriteria: &intschema.DatasetConfigurationLogDatasetMatchCriteria{
			Query: m.MatchCriteria.Query,
		},
	}
}
