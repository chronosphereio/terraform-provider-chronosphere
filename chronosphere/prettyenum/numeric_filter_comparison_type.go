// Copyright 2023 Chronosphere Inc.
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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// NumericFilterComparisonType is a wrapper of models.TraceSearchFilterNumericFilterComparisonType with support of user friendly values.
type NumericFilterComparisonType string

// Supported comparison types.
const (
	NumericFilterComparisonTypeEqualModel              = NumericFilterComparisonType(models.NumericFilterComparisonTypeEQUAL)
	NumericFilterComparisonTypeNotEqualModel           = NumericFilterComparisonType(models.NumericFilterComparisonTypeNOTEQUAL)
	NumericFilterComparisonTypeGreaterThanModel        = NumericFilterComparisonType(models.NumericFilterComparisonTypeGREATERTHAN)
	NumericFilterComparisonTypeGreaterThanOrEqualModel = NumericFilterComparisonType(models.NumericFilterComparisonTypeGREATERTHANOREQUAL)
	NumericFilterComparisonTypeLessThanModel           = NumericFilterComparisonType(models.NumericFilterComparisonTypeLESSTHAN)
	NumericFilterComparisonTypeLessThanOrEqualModel    = NumericFilterComparisonType(models.NumericFilterComparisonTypeLESSTHANOREQUAL)
)

var (
	NumericFilterComparisonTypeEqualLower              = NumericFilterComparisonType(strings.ToLower(string(NumericFilterComparisonTypeEqualModel)))
	NumericFilterComparisonTypeNotEqualLower           = NumericFilterComparisonType(strings.ToLower(string(NumericFilterComparisonTypeNotEqualModel)))
	NumericFilterComparisonTypeGreaterThanLower        = NumericFilterComparisonType(strings.ToLower(string(NumericFilterComparisonTypeGreaterThanModel)))
	NumericFilterComparisonTypeGreaterThanOrEqualLower = NumericFilterComparisonType(strings.ToLower(string(NumericFilterComparisonTypeGreaterThanOrEqualModel)))
	NumericFilterComparisonTypeLessThanLower           = NumericFilterComparisonType(strings.ToLower(string(NumericFilterComparisonTypeLessThanModel)))
	NumericFilterComparisonTypeLessThanOrEqualLower    = NumericFilterComparisonType(strings.ToLower(string(NumericFilterComparisonTypeLessThanOrEqualModel)))
)

var modelFromNumericFilterComparisonType = map[NumericFilterComparisonType]models.NumericFilterComparisonType{
	NumericFilterComparisonTypeEqualModel:              models.NumericFilterComparisonTypeEQUAL,
	NumericFilterComparisonTypeEqualLower:              models.NumericFilterComparisonTypeEQUAL,
	NumericFilterComparisonTypeNotEqualModel:           models.NumericFilterComparisonTypeNOTEQUAL,
	NumericFilterComparisonTypeNotEqualLower:           models.NumericFilterComparisonTypeNOTEQUAL,
	NumericFilterComparisonTypeGreaterThanModel:        models.NumericFilterComparisonTypeGREATERTHAN,
	NumericFilterComparisonTypeGreaterThanLower:        models.NumericFilterComparisonTypeGREATERTHAN,
	NumericFilterComparisonTypeGreaterThanOrEqualModel: models.NumericFilterComparisonTypeGREATERTHANOREQUAL,
	NumericFilterComparisonTypeGreaterThanOrEqualLower: models.NumericFilterComparisonTypeGREATERTHANOREQUAL,
	NumericFilterComparisonTypeLessThanModel:           models.NumericFilterComparisonTypeLESSTHAN,
	NumericFilterComparisonTypeLessThanLower:           models.NumericFilterComparisonTypeLESSTHAN,
	NumericFilterComparisonTypeLessThanOrEqualModel:    models.NumericFilterComparisonTypeLESSTHANOREQUAL,
	NumericFilterComparisonTypeLessThanOrEqualLower:    models.NumericFilterComparisonTypeLESSTHANOREQUAL,
}

// ValidateNumericFilterComparisonType validates the raw comparison type value.
func ValidateNumericFilterComparisonType(raw string) error {
	_, ok := modelFromNumericFilterComparisonType[NumericFilterComparisonType(raw)]
	if ok {
		return nil
	}
	return fmt.Errorf("invalid match: %s", raw)
}

// NewNumericFilterComparisonType creates a new comparison type.
func NewNumericFilterComparisonType(raw string) (NumericFilterComparisonType, error) {
	if err := ValidateNumericFilterComparisonType(raw); err != nil {
		return "", err
	}
	return NumericFilterComparisonType(raw), nil
}

// Model returns the model value of the comparison type.
func (ct NumericFilterComparisonType) Model() models.NumericFilterComparisonType {
	res, ok := modelFromNumericFilterComparisonType[ct]
	if ok {
		return res
	}
	return ""
}
