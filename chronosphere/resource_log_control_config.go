// Copyright 2025 Chronosphere Inc.
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
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// LogControlConfigFromModel maps an API model into an intschema model.
func LogControlConfigFromModel(m *models.Configv1LogControlConfig) (*intschema.LogControlConfig, error) {
	return (logControlConfigConverter{}).fromModel(m)
}

func resourceLogControlConfig() *schema.Resource {
	r := newGenericResource[
		*models.Configv1LogControlConfig,
		intschema.LogControlConfig,
		*intschema.LogControlConfig,
	](
		"log_control_config",
		logControlConfigConverter{},
		generatedLogControlConfig{},
	)

	return &schema.Resource{
		Schema:        tfschema.LogControlConfig,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&LogControlConfigDryRunCount),
		SchemaVersion: 1,
		Description:   "Config configuring log control in Chronosphere.",
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// LogControlConfigDryRunCount tracks how many times dry run is run during validation for testing.
var LogControlConfigDryRunCount atomic.Int64

type logControlConfigConverter struct{}

func (logControlConfigConverter) toModel(
	m *intschema.LogControlConfig,
) (*models.Configv1LogControlConfig, error) {
	return &models.Configv1LogControlConfig{
		Rules: sliceutil.Map(m.Rules, func(r intschema.LogControlConfigRules) *models.Configv1LogControlRule {
			rule := &models.Configv1LogControlRule{
				Name:   r.Name,
				Filter: r.Filter,
				Type:   models.Configv1LogControlRuleType(r.Type),
			}

			rule.Mode = models.Configv1LogControlRuleMode(r.Mode)

			if r.Sample != nil {
				rule.Sample = &models.LogControlRuleSample{
					Rate: r.Sample.Rate,
				}
			}

			if r.DropField != nil {
				rule.DropField = &models.LogControlRuleDropField{
					FieldRegex: r.DropField.FieldRegex,
				}
				if r.DropField.ParentPath != nil {
					rule.DropField.ParentPath = &models.Configv1LogFieldPath{
						Selector: r.DropField.ParentPath.Selector,
					}
				}
			}

			return rule
		}),
	}, nil
}

func (logControlConfigConverter) fromModel(
	m *models.Configv1LogControlConfig,
) (*intschema.LogControlConfig, error) {
	if m == nil {
		return nil, nil
	}

	return &intschema.LogControlConfig{
		Rules: sliceutil.Map(m.Rules, func(r *models.Configv1LogControlRule) intschema.LogControlConfigRules {
			rule := intschema.LogControlConfigRules{
				Name:   r.Name,
				Filter: r.Filter,
				Type:   string(r.Type),
			}

			rule.Mode = string(r.Mode)

			if r.Sample != nil {
				rule.Sample = &intschema.LogControlConfigRulesSample{
					Rate: r.Sample.Rate,
				}
			}

			if r.DropField != nil {
				rule.DropField = &intschema.LogControlConfigRulesDropField{
					FieldRegex: r.DropField.FieldRegex,
				}
				if r.DropField.ParentPath != nil {
					rule.DropField.ParentPath = &intschema.LogControlConfigRulesDropFieldParentPath{
						Selector: r.DropField.ParentPath.Selector,
					}
				}
			}

			return rule
		}),
	}, nil
}
