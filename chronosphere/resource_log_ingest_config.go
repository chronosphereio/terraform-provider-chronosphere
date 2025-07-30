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
	return &models.Configv1LogIngestConfig{
		PlaintextParsers: sliceutil.Map(m.PlaintextParser, func(p intschema.LogIngestConfigPlaintextParser) *models.Configv1PlaintextParser {
			return &models.Configv1PlaintextParser{
				Name:         p.Name,
				Mode:         models.Configv1PlaintextParserMode(p.Mode),
				Parser:       convertLogParserToModel(p.Parser),
				DropOriginal: p.DropOriginal,
			}
		}),
		FieldParsers: sliceutil.Map(m.FieldParser, func(p intschema.LogIngestConfigFieldParser) *models.Configv1LogFieldParser {
			return &models.Configv1LogFieldParser{
				Mode:        models.Configv1LogFieldParserMode(p.Mode),
				Source:      &models.Configv1LogFieldPath{Selector: p.Source.Selector},
				Destination: &models.Configv1LogFieldPath{Selector: p.Destination.Selector},
				Parser:      convertLogParserToModel(p.Parser),
			}
		}),
	}, nil
}

func convertLogParserToModel(p *intschema.LogParser) *models.Configv1LogParser {
	if p == nil {
		return nil
	}
	result := &models.Configv1LogParser{
		ParserType: models.LogParserParserType(p.ParserType),
	}

	if p.JsonParser != nil {
		result.JSONParser = struct{}{}
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
	return &intschema.LogIngestConfig{
		PlaintextParser: sliceutil.Map(m.PlaintextParsers, func(p *models.Configv1PlaintextParser) intschema.LogIngestConfigPlaintextParser {
			return intschema.LogIngestConfigPlaintextParser{
				Name:         p.Name,
				Mode:         string(p.Mode),
				Parser:       convertLogParserFromModel(p.Parser),
				DropOriginal: p.DropOriginal,
			}
		}),
		FieldParser: sliceutil.Map(m.FieldParsers, func(p *models.Configv1LogFieldParser) intschema.LogIngestConfigFieldParser {
			return intschema.LogIngestConfigFieldParser{
				Mode: string(p.Mode),
				Source: &intschema.LogFieldPath{
					Selector: p.Source.Selector,
				},
				Destination: &intschema.LogFieldPath{
					Selector: p.Destination.Selector,
				},
				Parser: convertLogParserFromModel(p.Parser),
			}
		}),
	}, nil
}

func convertLogParserFromModel(p *models.Configv1LogParser) *intschema.LogParser {
	result := &intschema.LogParser{
		ParserType: string(p.ParserType),
	}

	if p.JSONParser != nil {
		result.JsonParser = &intschema.JSONParser{}
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
