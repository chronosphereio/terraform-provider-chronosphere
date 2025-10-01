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

package enum

import (
	"encoding/json"
	"os"
	"testing"

	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/stretchr/testify/require"
)

func TestEnumConversions(t *testing.T) {
	// v1 -> v1
	require.Equal(t,
		configv1.Configv1LabelMatcherMatcherTypeREGEX,
		MatcherType.V1("REGEX"))

	// legacy -> v1
	require.Equal(t,
		configv1.Configv1LabelMatcherMatcherTypeEXACT,
		MatcherType.V1("EXACT_MATCHER_TYPE"))

	// alias -> v1
	require.Equal(t,
		configv1.Configv1LabelMatcherMatcherTypeEXACT,
		MatcherType.V1("EXACT"))

	// unknown -> v1
	require.Equal(t,
		configv1.Configv1LabelMatcherMatcherType("FOO_BAR_BAZ"),
		MatcherType.V1("FOO_BAR_BAZ"))

	// empty -> v1
	require.Equal(t, configv1.Configv1LabelMatcherMatcherType(""), MatcherType.V1(""))

	// empty -> v1 w/ non-empty default
	require.Equal(t, configv1.Configv1RollupRuleModeENABLED, RollupModeType.V1(""))

	// invalid legacy -> v1
	require.Equal(t, configv1.Configv1LabelMatcherMatcherType(""), MatcherType.V1("INVALID_MATCHER_TYPE"))

	// v1 -> v1
	require.Equal(t,
		configv1.PagerDutyActionSeverityERROR,
		LogScalePagerDutyActionSeverity.V1("ERROR"))

	// v1 -> v1
	require.Equal(t,
		configv1.WebhookActionHTTPMethodPOST,
		LogScaleWebhookActionHTTPMethod.V1("POST"))

	// v1 -> v1
	require.Equal(t,
		configv1.CommonPromQLMatcherTypeMatchEqual,
		PromQLMatcherType.V1("MatchEqual"))
}

func TestEnumValidateError(t *testing.T) {
	err := MatcherType.Validate("FOO_BAR_BAZ", nil)
	require.NotNil(t, err)
	require.Equal(t,
		`"FOO_BAR_BAZ" is not a valid MatcherType value; valid values: "EXACT", "REGEX"`,
		err[0].Summary)
}

func TestAllEnumsValidate(t *testing.T) {
	v1Spec := loadSwaggerSpec(t, "../pkg/configv1/swagger.json")
	unstableSpec := loadSwaggerSpec(t, "../pkg/configunstable/swagger.json")

	tests := []struct {
		legacySwaggerName string
		v1SwaggerName     string
		unstable          bool
		enum              Enum[string]
	}{
		{
			legacySwaggerName: "entitiesMatcherType",
			v1SwaggerName:     "configv1LabelMatcherMatcherType",
			enum:              MatcherType.ToStrings(),
		},
		{
			legacySwaggerName: "entitiesThresholdOp",
			v1SwaggerName:     "ConditionOp",
			enum:              ConditionOp.ToStrings(),
		},
		{
			legacySwaggerName: "apiAggregationType",
			v1SwaggerName:     "configv1AggregationType",
			enum:              AggregationType.ToStrings(),
		},
		{
			legacySwaggerName: "apiMetricType",
			v1SwaggerName:     "RollupRuleMetricType",
			enum:              MetricType.ToStrings(),
		},
		{
			legacySwaggerName: "entitiesAlertReceiverOpsGenieResponseType",
			v1SwaggerName:     "ResponderResponderType",
			enum:              OpsgenieResponderType.ToStrings(),
		},
		{
			v1SwaggerName: "ResourceAttributesFlattenMode",
			enum:          ResourceAttributesFlattenMode.ToStrings(),
		},
		{
			v1SwaggerName: "ResourceAttributesFilterMode",
			enum:          ResourceAttributesFilterMode.ToStrings(),
		},
		{
			v1SwaggerName: "PagerDutyActionSeverity",
			enum:          LogScalePagerDutyActionSeverity.ToStrings(),
		},
		{
			v1SwaggerName: "WebhookActionHTTPMethod",
			enum:          LogScaleWebhookActionHTTPMethod.ToStrings(),
		},
		{
			v1SwaggerName: "LogScaleAlertAlertType",
			enum:          LogscaleAlertType.ToStrings(),
		},
		{
			v1SwaggerName: "SLITimeSliceSize",
			enum:          SLITimeSliceSize.ToStrings(),
		},
		{
			v1SwaggerName: "ReplaceFieldReplaceMode",
			enum:          ReplaceFieldReplaceMode.ToStrings(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.enum.Name(), func(t *testing.T) {
			swaggerSpec := v1Spec
			if tt.unstable {
				swaggerSpec = unstableSpec
			}
			for _, v := range loadEnumValues(t, swaggerSpec, tt.v1SwaggerName) {
				require.Nil(t, tt.enum.Validate(v, nil))
			}
		})
	}
}

func loadSwaggerSpec(t *testing.T, filename string) *openapi2.T {
	b, err := os.ReadFile(filename)
	require.NoError(t, err)

	spec := &openapi2.T{}
	require.NoError(t, json.Unmarshal(b, spec))

	return spec
}

func loadEnumValues(t *testing.T, spec *openapi2.T, name string) []string {
	def, ok := spec.Definitions[name]
	require.True(t, ok, "enum %q not found", name)

	var result []string
	for _, v := range def.Value.Enum {
		result = append(result, v.(string))
	}
	return result
}
