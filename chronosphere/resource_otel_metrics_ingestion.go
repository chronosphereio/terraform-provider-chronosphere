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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceOtelMetricsIngestion() *schema.Resource {
	r := newGenericResource(
		"otel_metrics_ingestion",
		otelMetricsIngestionConverter{},
		generatedUnstableOtelMetricsIngestion{},
	)
	return &schema.Resource{
		Schema:        tfschema.OtelMetricsIngestion,
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		CustomizeDiff: r.ValidateDryRun(&OtelMetricsIngestionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// OtelMetricsIngestionDryRunCount tracks how many times dry run is run during validation for testing.
var OtelMetricsIngestionDryRunCount atomic.Int64

type otelMetricsIngestionConverter struct{}

func (otelMetricsIngestionConverter) toModel(in *intschema.OtelMetricsIngestion) (*models.ConfigunstableOtelMetricsIngestion, error) {
	out := &models.ConfigunstableOtelMetricsIngestion{}
	if in.ResourceAttributes != nil {
		out.ResourceAttributes = &models.OtelMetricsIngestionResourceAttributes{
			FlattenMode:        enum.ResourceAttributesFlattenMode.V1(in.ResourceAttributes.FlattenMode),
			FilterMode:         enum.ResourceAttributesFilterMode.V1(in.ResourceAttributes.FilterMode),
			ExcludeKeys:        in.ResourceAttributes.ExcludeKeys,
			GenerateTargetInfo: in.ResourceAttributes.GenerateTargetInfo,
		}
	}
	return out, nil
}

func (otelMetricsIngestionConverter) fromModel(in *models.ConfigunstableOtelMetricsIngestion) (*intschema.OtelMetricsIngestion, error) {
	out := &intschema.OtelMetricsIngestion{}
	if in.ResourceAttributes != nil {
		out.ResourceAttributes = &intschema.OtelMetricsIngestionResourceAttributes{
			FlattenMode:        string(in.ResourceAttributes.FlattenMode),
			FilterMode:         string(in.ResourceAttributes.FilterMode),
			ExcludeKeys:        in.ResourceAttributes.ExcludeKeys,
			GenerateTargetInfo: in.ResourceAttributes.GenerateTargetInfo,
		}
	}
	return out, nil
}
