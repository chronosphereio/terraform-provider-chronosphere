package prettyenum

import (
	"fmt"
	"strings"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// SpanFilterMatchType is a wrapper of models.TraceSearchFilterSpanFilterMatchType with support of user friendly values.
type SpanFilterMatchType string

// Supported matcher types.
const (
	SpanFilterMatchTypeIncludeModel = SpanFilterMatchType(models.SpanFilterSpanFilterMatchTypeINCLUDE)
	SpanFilterMatchTypeExcludeModel = SpanFilterMatchType(models.SpanFilterSpanFilterMatchTypeEXCLUDE)
)

var (
	SpanFilterMatchTypeInclude = SpanFilterMatchType(strings.ToLower(string(SpanFilterMatchTypeIncludeModel)))
	SpanFilterMatchTypeExclude = SpanFilterMatchType(strings.ToLower(string(SpanFilterMatchTypeExcludeModel)))
)

var modelFromSpanFilterMatchType = map[SpanFilterMatchType]models.SpanFilterSpanFilterMatchType{
	SpanFilterMatchTypeInclude: models.SpanFilterSpanFilterMatchTypeINCLUDE,
	SpanFilterMatchTypeExclude: models.SpanFilterSpanFilterMatchTypeEXCLUDE,

	SpanFilterMatchTypeIncludeModel: models.SpanFilterSpanFilterMatchTypeINCLUDE,
	SpanFilterMatchTypeExcludeModel: models.SpanFilterSpanFilterMatchTypeEXCLUDE,
}

var spanFilterMatchTypeFromModel = map[models.SpanFilterSpanFilterMatchType]SpanFilterMatchType{
	models.SpanFilterSpanFilterMatchTypeINCLUDE: SpanFilterMatchTypeInclude,
	models.SpanFilterSpanFilterMatchTypeEXCLUDE: SpanFilterMatchTypeExclude,
}

// ValidateSpanFilterMatchType validates the raw matcher type value.
func ValidateSpanFilterMatchType(raw string) error {
	_, ok := modelFromSpanFilterMatchType[SpanFilterMatchType(raw)]
	if ok {
		return nil
	}
	return fmt.Errorf("invalid match_type: %s", raw)
}

// NewSpanFilterMatchType creates a new matcher type,
func NewSpanFilterMatchType(raw string) (SpanFilterMatchType, error) {
	if err := ValidateSpanFilterMatchType(raw); err != nil {
		return "", err
	}
	return SpanFilterMatchType(raw), nil
}

// Model returns the model value of the matcher type.
func (mt SpanFilterMatchType) Model() models.SpanFilterSpanFilterMatchType {
	res, ok := modelFromSpanFilterMatchType[mt]
	if ok {
		return res
	}
	return ""
}

// SpanFilterMatchTypeFromModel returns the human readable string for a matcher type enum
func SpanFilterMatchTypeFromModel(m models.SpanFilterSpanFilterMatchType) SpanFilterMatchType {
	res, ok := spanFilterMatchTypeFromModel[m]
	if !ok {
		return ""
	}
	return res
}
