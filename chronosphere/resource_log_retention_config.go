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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// LogRetentionConfigFromModel maps an API model into an intschema model.
func LogRetentionConfigFromModel(m *models.Configv1LogRetentionConfig) (*intschema.LogRetentionConfig, error) {
	return (logRetentionConfigConverter{}).fromModel(m)
}

func resourceLogRetentionConfig() *schema.Resource {
	r := newGenericResource(
		"log_retention_config",
		logRetentionConfigConverter{},
		generatedLogRetentionConfig{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.LogRetentionConfig,
		CustomizeDiff: r.ValidateDryRun(&LogRetentionConfigDryRunCount),
		Description:   "Configures long-term log retention for logs matching a filter.",
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// LogRetentionConfigDryRunCount tracks how many times dry run is run during validation for testing.
var LogRetentionConfigDryRunCount atomic.Int64

type logRetentionConfigConverter struct{}

func (logRetentionConfigConverter) toModel(
	m *intschema.LogRetentionConfig,
) (*models.Configv1LogRetentionConfig, error) {
	return &models.Configv1LogRetentionConfig{
		Name:          m.Name,
		Slug:          m.Slug,
		Filter:        m.Filter,
		Mode:          enum.LogRetentionConfigMode.V1(m.Mode),
		RetentionDays: m.RetentionDays,
	}, nil
}

func (logRetentionConfigConverter) fromModel(
	m *models.Configv1LogRetentionConfig,
) (*intschema.LogRetentionConfig, error) {
	return &intschema.LogRetentionConfig{
		Name:          m.Name,
		Slug:          m.Slug,
		Filter:        m.Filter,
		Mode:          string(m.Mode),
		RetentionDays: m.RetentionDays,
	}, nil
}
