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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// StringFilterMatchType is a wrapper of models.TraceSearchFilterStringFilterMatchType with support of user friendly values.
type StringFilterMatchType string

// Supported matcher types.
const (
	StringFilterMatchTypeExactModel         = StringFilterMatchType(models.StringFilterStringFilterMatchTypeEXACT)
	StringFilterMatchTypeRegexModel         = StringFilterMatchType(models.StringFilterStringFilterMatchTypeREGEX)
	StringFilterMatchTypeExactNegationModel = StringFilterMatchType(models.StringFilterStringFilterMatchTypeEXACTNEGATION)
	StringFilterMatchTypeRegexNegationModel = StringFilterMatchType(models.StringFilterStringFilterMatchTypeREGEXNEGATION)
	StringFilterMatchTypeInModel            = StringFilterMatchType(models.StringFilterStringFilterMatchTypeIN)
	StringFilterMatchTypeNotInModel         = StringFilterMatchType(models.StringFilterStringFilterMatchTypeNOTIN)
)

var (
	StringFilterMatchTypeExact         = StringFilterMatchType(strings.ToLower(string(StringFilterMatchTypeExactModel)))
	StringFilterMatchTypeRegex         = StringFilterMatchType(strings.ToLower(string(StringFilterMatchTypeRegexModel)))
	StringFilterMatchTypeExactNegation = StringFilterMatchType(strings.ToLower(string(StringFilterMatchTypeExactNegationModel)))
	StringFilterMatchTypeRegexNegation = StringFilterMatchType(strings.ToLower(string(StringFilterMatchTypeRegexNegationModel)))
	StringFilterMatchTypeIn            = StringFilterMatchType(strings.ToLower(string(StringFilterMatchTypeInModel)))
	StringFilterMatchTypeNotIn         = StringFilterMatchType(strings.ToLower(string(StringFilterMatchTypeNotInModel)))
)

var modelFromStringFilterMatchType = map[StringFilterMatchType]models.StringFilterStringFilterMatchType{
	StringFilterMatchTypeExact:         models.StringFilterStringFilterMatchTypeEXACT,
	StringFilterMatchTypeRegex:         models.StringFilterStringFilterMatchTypeREGEX,
	StringFilterMatchTypeExactNegation: models.StringFilterStringFilterMatchTypeEXACTNEGATION,
	StringFilterMatchTypeRegexNegation: models.StringFilterStringFilterMatchTypeREGEXNEGATION,
	StringFilterMatchTypeIn:            models.StringFilterStringFilterMatchTypeIN,
	StringFilterMatchTypeNotIn:         models.StringFilterStringFilterMatchTypeNOTIN,

	StringFilterMatchTypeExactModel:         models.StringFilterStringFilterMatchTypeEXACT,
	StringFilterMatchTypeRegexModel:         models.StringFilterStringFilterMatchTypeREGEX,
	StringFilterMatchTypeExactNegationModel: models.StringFilterStringFilterMatchTypeEXACTNEGATION,
	StringFilterMatchTypeRegexNegationModel: models.StringFilterStringFilterMatchTypeREGEXNEGATION,
	StringFilterMatchTypeInModel:            models.StringFilterStringFilterMatchTypeIN,
	StringFilterMatchTypeNotInModel:         models.StringFilterStringFilterMatchTypeNOTIN,
}

var stringFilterMatchTypeFromModel = map[models.StringFilterStringFilterMatchType]StringFilterMatchType{
	models.StringFilterStringFilterMatchTypeEXACT:         StringFilterMatchTypeExact,
	models.StringFilterStringFilterMatchTypeREGEX:         StringFilterMatchTypeRegex,
	models.StringFilterStringFilterMatchTypeEXACTNEGATION: StringFilterMatchTypeExactNegation,
	models.StringFilterStringFilterMatchTypeREGEXNEGATION: StringFilterMatchTypeRegexNegation,
	models.StringFilterStringFilterMatchTypeIN:            StringFilterMatchTypeIn,
	models.StringFilterStringFilterMatchTypeNOTIN:         StringFilterMatchTypeNotIn,
}

// ValidateStringFilterMatchType validates the raw matcher type value.
func ValidateStringFilterMatchType(raw string) error {
	_, ok := modelFromStringFilterMatchType[StringFilterMatchType(raw)]
	if ok {
		return nil
	}
	return fmt.Errorf("invalid match: %s", raw)
}

// NewStringFilterMatchType creates a new matcher type,
func NewStringFilterMatchType(raw string) (StringFilterMatchType, error) {
	if err := ValidateStringFilterMatchType(raw); err != nil {
		return "", err
	}
	return StringFilterMatchType(raw), nil
}

// Model returns the model value of the matcher type.
func (mt StringFilterMatchType) Model() models.StringFilterStringFilterMatchType {
	res, ok := modelFromStringFilterMatchType[mt]
	if ok {
		return res
	}
	return ""
}

// StringFilterMatchTypeFromModel returns the human readable string for a matcher type enum
func StringFilterMatchTypeFromModel(m models.StringFilterStringFilterMatchType) StringFilterMatchType {
	res, ok := stringFilterMatchTypeFromModel[m]
	if !ok {
		return ""
	}
	return res
}
