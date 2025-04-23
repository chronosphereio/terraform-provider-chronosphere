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
		Parsers: sliceutil.Map(m.Parser, func(p intschema.LogIngestConfigParser) *models.Configv1LogParser {
			return &models.Configv1LogParser{
				Name:  p.Name,
				Regex: p.Regex,
			}
		}),
	}, nil
}

func (logIngestConfigConverter) fromModel(
	m *models.Configv1LogIngestConfig,
) (*intschema.LogIngestConfig, error) {
	return &intschema.LogIngestConfig{
		Parser: sliceutil.Map(m.Parsers, func(p *models.Configv1LogParser) intschema.LogIngestConfigParser {
			return intschema.LogIngestConfigParser{
				Name:  p.Name,
				Regex: p.Regex,
			}
		}),
	}, nil
}
