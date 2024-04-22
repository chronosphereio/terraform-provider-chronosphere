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

package prettyenum

import (
	"fmt"
	"strings"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// DatasetType is a wrapper of models.DatasetDatasetType with support for user friendly values.
type DatasetType string

// Supported dataset types.
const (
	DatasetDatasetTypeTracesModel = DatasetType(models.DatasetDatasetTypeTRACES)
)

var DatasetDatasetTypeTracesLower = DatasetType(strings.ToLower(string(DatasetDatasetTypeTracesModel)))

var modelFromDatasetType = map[DatasetType]models.DatasetDatasetType{
	DatasetDatasetTypeTracesModel: models.DatasetDatasetTypeTRACES,
	DatasetDatasetTypeTracesLower: models.DatasetDatasetTypeTRACES,
}

var datasetTypeFromModel = map[models.DatasetDatasetType]DatasetType{
	models.DatasetDatasetTypeTRACES: DatasetDatasetTypeTracesModel,
}

// ValidateDatasetDatasetType validates the raw dataset type value.
func ValidateDatasetDatasetType(raw string) error {
	_, ok := modelFromDatasetType[DatasetType(raw)]
	if ok {
		return nil
	}
	return fmt.Errorf("invalid dataset type: %s", raw)
}

// NewDatasetDatasetType creates a new dataset type.
func NewDatasetDatasetType(raw string) (DatasetType, error) {
	if err := ValidateDatasetDatasetType(raw); err != nil {
		return "", err
	}
	return DatasetType(raw), nil
}

// Model returns the model value of the dataset type.
func (ct DatasetType) Model() models.DatasetDatasetType {
	res, ok := modelFromDatasetType[ct]
	if ok {
		return res
	}
	return ""
}

// DatasetTypeFromModel returns the human-readable string for a matcher type enum.
func DatasetTypeFromModel(m models.DatasetDatasetType) DatasetType {
	res, ok := datasetTypeFromModel[m]
	if !ok {
		return ""
	}
	return res
}
