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

// LogIngestConfigFromModel maps an API model into an intschema model.
func LogIngestConfigFromModel(m *models.Configv1LogIngestConfig) (*intschema.LogIngestConfig, error) {
	return (logIngestConfigConverter{}).fromModel(m)
}

func resourceLogIngestConfig() *schema.Resource {
	r := newGenericResource[
		*models.Configv1LogIngestConfig,
		intschema.LogIngestConfig,
		*intschema.LogIngestConfig,
	](
		"log_ingest_config",
		logIngestConfigConverter{},
		generatedLogIngestConfig{},
	)

	return &schema.Resource{
		Schema:        tfschema.LogIngestConfig,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&LogIngestConfigDryRunCount),
		SchemaVersion: 1,
		Description:   "Config configuring log ingestion in Chronosphere.",
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// LogIngestConfigDryRunCount tracks how many times dry run is run during validation for testing.
var LogIngestConfigDryRunCount atomic.Int64

type logIngestConfigConverter struct{}

func (logIngestConfigConverter) toModel(
	m *intschema.LogIngestConfig,
) (*models.Configv1LogIngestConfig, error) {
	result := &models.Configv1LogIngestConfig{
		PlaintextParsers: sliceutil.Map(m.PlaintextParser, func(p intschema.LogIngestConfigPlaintextParser) *models.LogIngestConfigPlaintextParser {
			return &models.LogIngestConfigPlaintextParser{
				Name:         p.Name,
				Mode:         models.LogIngestConfigPlaintextParserMode(p.Mode),
				Parser:       convertLogParserToModel(p.Parser),
				KeepOriginal: p.KeepOriginal,
			}
		}),
		FieldParsers: sliceutil.Map(m.FieldParser, func(p intschema.LogIngestConfigFieldParser) *models.LogIngestConfigLogFieldParser {
			fp := &models.LogIngestConfigLogFieldParser{
				Mode:   models.LogIngestConfigLogFieldParserMode(p.Mode),
				Source: &models.Configv1LogFieldPath{Selector: p.Source.Selector},
				Parser: convertLogParserToModel(p.Parser),
			}
			if p.Destination != nil {
				fp.Destination = &models.Configv1LogFieldPath{Selector: p.Destination.Selector}
			}
			return fp
		}),
	}

	result.FieldNormalization = convertFieldNormalizationToModel(m.FieldNormalization)

	return result, nil
}

func convertLogParserToModel(p *intschema.LogParser) *models.LogIngestConfigLogParser {
	if p == nil {
		return nil
	}
	result := &models.LogIngestConfigLogParser{
		ParserType: models.LogParserParserType(p.ParserType),
	}

	if p.RegexParser != nil {
		result.RegexParser = &models.LogParserRegexParser{
			Regex: p.RegexParser.Regex,
		}
	}

	if p.KeyValueParser != nil {
		result.KeyValueParser = &models.LogParserKeyValueParser{
			PairSeparator: p.KeyValueParser.PairSeparator,
			Delimiter:     p.KeyValueParser.Delimiter,
			TrimSet:       p.KeyValueParser.TrimSet,
		}
	}

	return result
}

func (logIngestConfigConverter) fromModel(
	m *models.Configv1LogIngestConfig,
) (*intschema.LogIngestConfig, error) {
	result := &intschema.LogIngestConfig{
		PlaintextParser: sliceutil.Map(m.PlaintextParsers, func(p *models.LogIngestConfigPlaintextParser) intschema.LogIngestConfigPlaintextParser {
			return intschema.LogIngestConfigPlaintextParser{
				Name:         p.Name,
				Mode:         string(p.Mode),
				Parser:       convertLogParserFromModel(p.Parser),
				KeepOriginal: p.KeepOriginal,
			}
		}),
		FieldParser: sliceutil.Map(m.FieldParsers, func(p *models.LogIngestConfigLogFieldParser) intschema.LogIngestConfigFieldParser {
			fp := intschema.LogIngestConfigFieldParser{
				Mode: string(p.Mode),
				Source: &intschema.LogFieldPath{
					Selector: p.Source.Selector,
				},
				Parser: convertLogParserFromModel(p.Parser),
			}
			if p.Destination != nil {
				fp.Destination = &intschema.LogFieldPath{
					Selector: p.Destination.Selector,
				}
			}
			return fp
		}),
	}

	if m.FieldNormalization != nil {
		result.FieldNormalization = convertFieldNormalizationFromModel(m.FieldNormalization)
	}

	return result, nil
}

func convertLogParserFromModel(p *models.LogIngestConfigLogParser) *intschema.LogParser {
	result := &intschema.LogParser{
		ParserType: string(p.ParserType),
	}

	if p.RegexParser != nil {
		result.RegexParser = &intschema.RegexParser{
			Regex: p.RegexParser.Regex,
		}
	}

	if p.KeyValueParser != nil {
		result.KeyValueParser = &intschema.KeyValueParser{
			PairSeparator: p.KeyValueParser.PairSeparator,
			Delimiter:     p.KeyValueParser.Delimiter,
			TrimSet:       p.KeyValueParser.TrimSet,
		}
	}

	return result
}

func convertFieldNormalizationToModel(fn *intschema.LogIngestConfigFieldNormalization) *models.LogIngestConfigFieldNormalization {
	if fn == nil {
		return nil
	}

	result := &models.LogIngestConfigFieldNormalization{}

	result.CustomFieldNormalization = sliceutil.Map(fn.CustomFieldNormalization, func(n intschema.LogIngestConfigNamedStringNormalization) *models.LogIngestConfigNamedStringNormalization {
		return &models.LogIngestConfigNamedStringNormalization{
			Normalization: convertStringNormalizationToModel(n.Normalization),
			Target:        n.Target,
		}
	})

	result.Message = convertStringNormalizationToModel(fn.Message)

	if fn.PrimaryKey != nil {
		result.PrimaryKey = &models.LogIngestConfigNamedStringNormalization{
			Normalization: convertStringNormalizationToModel(fn.PrimaryKey.Normalization),
			Target:        fn.PrimaryKey.Target,
		}
	}

	result.Severity = convertStringNormalizationToModel(fn.Severity)

	if fn.Timestamp != nil {
		result.Timestamp = &models.LogIngestConfigTimestampNormalization{
			Source: sliceutil.Map(fn.Timestamp.Source, func(fp intschema.LogFieldPath) *models.Configv1LogFieldPath {
				return &models.Configv1LogFieldPath{Selector: fp.Selector}
			}),
		}
	}

	return result
}

func convertStringNormalizationToModel(sn *intschema.LogIngestConfigStringNormalization) *models.LogIngestConfigStringNormalization {
	if sn == nil {
		return nil
	}

	result := &models.LogIngestConfigStringNormalization{
		DefaultValue:     sn.DefaultValue,
		SanitizePatterns: sn.SanitizePatterns,
		ValueMap:         sn.ValueMap,
	}

	result.Source = sliceutil.Map(sn.Source, func(fp intschema.LogFieldPath) *models.Configv1LogFieldPath {
		return &models.Configv1LogFieldPath{Selector: fp.Selector}
	})

	return result
}

func convertFieldNormalizationFromModel(fn *models.LogIngestConfigFieldNormalization) *intschema.LogIngestConfigFieldNormalization {
	if fn == nil {
		return nil
	}

	result := &intschema.LogIngestConfigFieldNormalization{}

	result.CustomFieldNormalization = sliceutil.Map(fn.CustomFieldNormalization, func(n *models.LogIngestConfigNamedStringNormalization) intschema.LogIngestConfigNamedStringNormalization {
		return intschema.LogIngestConfigNamedStringNormalization{
			Normalization: convertStringNormalizationFromModel(n.Normalization),
			Target:        n.Target,
		}
	})

	result.Message = convertStringNormalizationFromModel(fn.Message)

	if fn.PrimaryKey != nil {
		result.PrimaryKey = &intschema.LogIngestConfigNamedStringNormalization{
			Normalization: convertStringNormalizationFromModel(fn.PrimaryKey.Normalization),
			Target:        fn.PrimaryKey.Target,
		}
	}

	result.Severity = convertStringNormalizationFromModel(fn.Severity)

	if fn.Timestamp != nil {
		result.Timestamp = &intschema.LogIngestConfigFieldNormalizationTimestamp{
			Source: sliceutil.Map(fn.Timestamp.Source, func(fp *models.Configv1LogFieldPath) intschema.LogFieldPath {
				return intschema.LogFieldPath{Selector: fp.Selector}
			}),
		}
	}

	return result
}

func convertStringNormalizationFromModel(sn *models.LogIngestConfigStringNormalization) *intschema.LogIngestConfigStringNormalization {
	if sn == nil {
		return nil
	}

	result := &intschema.LogIngestConfigStringNormalization{
		DefaultValue:     sn.DefaultValue,
		SanitizePatterns: sn.SanitizePatterns,
		ValueMap:         sn.ValueMap,
	}

	result.Source = sliceutil.Map(sn.Source, func(fp *models.Configv1LogFieldPath) intschema.LogFieldPath {
		return intschema.LogFieldPath{Selector: fp.Selector}
	})

	return result
}
