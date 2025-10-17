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
			rule.EmitMetrics = convertEmitMetricsToModel(r.EmitMetrics)
			rule.ReplaceField = convertReplaceFieldToModel(r.ReplaceField)

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
					rule.DropField.ParentPath = &intschema.LogControlConfigFieldPath{
						Selector: r.DropField.ParentPath.Selector,
					}
				}
			}

			rule.EmitMetrics = convertEmitMetricsFromModel(r.EmitMetrics)
			rule.ReplaceField = convertReplaceFieldFromModel(r.ReplaceField)

			return rule
		}),
	}, nil
}

func convertEmitMetricsToModel(em *intschema.LogControlConfigRulesEmitMetrics) *models.LogControlRuleEmitMetrics {
	if em == nil {
		return nil
	}

	result := &models.LogControlRuleEmitMetrics{
		DropLog: em.DropLog,
		Mode:    models.EmitMetricsMetricMode(em.Mode),
		Name:    em.Name,
	}

	if em.Counter != nil {
		result.Counter = &models.EmitMetricsCounter{}
		if em.Counter.Value != nil {
			result.Counter.Value = &models.Configv1LogFieldPath{
				Selector: em.Counter.Value.Selector,
			}
		}
	}

	if em.Gauge != nil {
		result.Gauge = &models.EmitMetricsGauge{
			AggregationType: models.GaugeAggregationType(em.Gauge.AggregationType),
		}
		if em.Gauge.Value != nil {
			result.Gauge.Value = &models.Configv1LogFieldPath{
				Selector: em.Gauge.Value.Selector,
			}
		}
	}

	if em.Histogram != nil {
		result.Histogram = &models.EmitMetricsHistogram{}
		if em.Histogram.Value != nil {
			result.Histogram.Value = &models.Configv1LogFieldPath{
				Selector: em.Histogram.Value.Selector,
			}
		}
	}

	result.Labels = sliceutil.Map(em.Labels, func(l intschema.LogControlConfigRulesEmitMetricsLabels) *models.LogControlRuleEmitMetricsLabel {
		label := &models.LogControlRuleEmitMetricsLabel{
			Key: l.Key,
		}
		if l.Value != nil {
			label.Value = &models.Configv1LogFieldPath{
				Selector: l.Value.Selector,
			}
		}
		return label
	})

	return result
}

func convertReplaceFieldToModel(rf *intschema.LogControlConfigRulesReplaceField) *models.LogControlRuleReplaceField {
	if rf == nil {
		return nil
	}

	result := &models.LogControlRuleReplaceField{
		ReplaceAll:   rf.ReplaceAll,
		ReplaceMode:  models.ReplaceFieldReplaceMode(rf.ReplaceMode),
		ReplaceRegex: rf.ReplaceRegex,
	}

	if rf.Field != nil {
		result.Field = &models.Configv1LogFieldPath{
			Selector: rf.Field.Selector,
		}
	}

	if rf.MappedValue != nil {
		result.MappedValue = &models.ReplaceFieldMappedValue{
			DefaultValue: rf.MappedValue.DefaultValue,
			UseDefault:   rf.MappedValue.UseDefault,
		}
		result.MappedValue.Pairs = sliceutil.Map(rf.MappedValue.Pairs, func(p intschema.LogControlConfigRulesReplaceFieldMappedValuePairs) *models.MappedValueReplacePair {
			return &models.MappedValueReplacePair{
				Key:   p.Key,
				Value: p.Value,
			}
		})

	}

	if rf.StaticValue != nil {
		result.StaticValue = &models.ReplaceFieldStaticValue{
			Value: rf.StaticValue.Value,
		}
	}

	return result
}

func convertEmitMetricsFromModel(em *models.LogControlRuleEmitMetrics) *intschema.LogControlConfigRulesEmitMetrics {
	if em == nil {
		return nil
	}

	result := &intschema.LogControlConfigRulesEmitMetrics{
		DropLog: em.DropLog,
		Mode:    string(em.Mode),
		Name:    em.Name,
	}

	if em.Counter != nil {
		result.Counter = &intschema.LogControlConfigRulesEmitMetricsCounter{}
		if em.Counter.Value != nil {
			result.Counter.Value = &intschema.LogControlConfigFieldPath{
				Selector: em.Counter.Value.Selector,
			}
		}
	}

	if em.Gauge != nil {
		result.Gauge = &intschema.LogControlConfigRulesEmitMetricsGauge{
			AggregationType: string(em.Gauge.AggregationType),
		}
		if em.Gauge.Value != nil {
			result.Gauge.Value = &intschema.LogControlConfigFieldPath{
				Selector: em.Gauge.Value.Selector,
			}
		}
	}

	if em.Histogram != nil {
		result.Histogram = &intschema.LogControlConfigRulesEmitMetricsHistogram{}
		if em.Histogram.Value != nil {
			result.Histogram.Value = &intschema.LogControlConfigFieldPath{
				Selector: em.Histogram.Value.Selector,
			}
		}
	}

	result.Labels = sliceutil.Map(em.Labels, func(l *models.LogControlRuleEmitMetricsLabel) intschema.LogControlConfigRulesEmitMetricsLabels {
		label := intschema.LogControlConfigRulesEmitMetricsLabels{
			Key: l.Key,
		}
		if l.Value != nil {
			label.Value = &intschema.LogControlConfigFieldPath{
				Selector: l.Value.Selector,
			}
		}
		return label
	})

	return result
}

func convertReplaceFieldFromModel(rf *models.LogControlRuleReplaceField) *intschema.LogControlConfigRulesReplaceField {
	if rf == nil {
		return nil
	}

	result := &intschema.LogControlConfigRulesReplaceField{
		ReplaceAll:   rf.ReplaceAll,
		ReplaceMode:  string(rf.ReplaceMode),
		ReplaceRegex: rf.ReplaceRegex,
	}

	if rf.Field != nil {
		result.Field = &intschema.LogControlConfigFieldPath{
			Selector: rf.Field.Selector,
		}
	}

	if rf.MappedValue != nil {
		result.MappedValue = &intschema.LogControlConfigRulesReplaceFieldMappedValue{
			DefaultValue: rf.MappedValue.DefaultValue,
			UseDefault:   rf.MappedValue.UseDefault,
		}
		result.MappedValue.Pairs = sliceutil.Map(rf.MappedValue.Pairs, func(p *models.MappedValueReplacePair) intschema.LogControlConfigRulesReplaceFieldMappedValuePairs {
			return intschema.LogControlConfigRulesReplaceFieldMappedValuePairs{
				Key:   p.Key,
				Value: p.Value,
			}
		})
	}

	if rf.StaticValue != nil {
		result.StaticValue = &intschema.LogControlConfigRulesReplaceFieldStaticValue{
			Value: rf.StaticValue.Value,
		}
	}

	return result
}
