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
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/otel_metrics_ingestion"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceOtelMetricsIngestion() *schema.Resource {
	return &schema.Resource{
		Schema:        tfschema.OtelMetricsIngestion,
		CreateContext: resourceOtelMetricsIngestionCreate,
		ReadContext:   resourceOtelMetricsIngestionRead,
		UpdateContext: resourceOtelMetricsIngestionUpdate,
		DeleteContext: resourceOtelMetricsIngestionDelete,
		CustomizeDiff: resourceOtelMetricsIngestionCustomizeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceOtelMetricsIngestionToModel(in *intschema.OtelMetricsIngestion) (*models.ConfigunstableOtelMetricsIngestion, error) {
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

func resourceOtelMetricsIngestionFromModel(in *models.ConfigunstableOtelMetricsIngestion) *intschema.OtelMetricsIngestion {
	out := &intschema.OtelMetricsIngestion{}
	if in.ResourceAttributes != nil {
		out.ResourceAttributes = &intschema.OtelMetricsIngestionResourceAttributes{
			FlattenMode:        string(in.ResourceAttributes.FlattenMode),
			FilterMode:         string(in.ResourceAttributes.FilterMode),
			ExcludeKeys:        in.ResourceAttributes.ExcludeKeys,
			GenerateTargetInfo: in.ResourceAttributes.GenerateTargetInfo,
		}
	}
	return out
}

func resourceOtelMetricsIngestionCreate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "otel_metrics_ingestion")
	cli := getConfigUnstableClient(meta)

	omi, err := buildOtelMetricsIngestion(d)
	if err != nil {
		return diag.Errorf("could not build OTel metrics ingestion config: %v", err)
	}
	req := &otel_metrics_ingestion.CreateOtelMetricsIngestionParams{
		Context: ctx,
		Body: &models.ConfigunstableCreateOtelMetricsIngestionRequest{
			OtelMetricsIngestion: omi,
		},
	}

	if _, err := cli.OtelMetricsIngestion.CreateOtelMetricsIngestion(req); err != nil {
		return diag.Errorf("could not create OTel metrics ingestion config: %v", err)
	}

	d.SetId(OtelMetricsIngestionID)

	return nil
}

func resourceOtelMetricsIngestionRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "otel_metrics_ingestion")
	cli := getConfigUnstableClient(meta)

	resp, err := cli.OtelMetricsIngestion.ReadOtelMetricsIngestion(&otel_metrics_ingestion.ReadOtelMetricsIngestionParams{Context: ctx})
	if clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to read OTel metrics ingestion config: %v", clienterror.Wrap(err))
	}

	s := resourceOtelMetricsIngestionFromModel(resp.Payload.OtelMetricsIngestion)
	if err := s.ToResourceData(d); err != nil {
		return err
	}
	d.SetId(OtelMetricsIngestionID)
	return nil
}

func resourceOtelMetricsIngestionUpdate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "otel_metrics_ingestion")
	cli := getConfigUnstableClient(meta)

	omi, err := buildOtelMetricsIngestion(d)
	if err != nil {
		return diag.Errorf("could not build OTel metrics ingestion config: %v", err)
	}
	req := &otel_metrics_ingestion.UpdateOtelMetricsIngestionParams{
		Context: ctx,
		Body: &models.ConfigunstableUpdateOtelMetricsIngestionRequest{
			OtelMetricsIngestion: omi,
		},
	}
	if _, err := cli.OtelMetricsIngestion.UpdateOtelMetricsIngestion(req); err != nil {
		return diag.Errorf("unable to update OTel metrics ingestion config: %v", err)
	}
	return nil
}

func resourceOtelMetricsIngestionDelete(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "otel_metrics_ingestion")
	cli := getConfigUnstableClient(meta)

	req := &otel_metrics_ingestion.DeleteOtelMetricsIngestionParams{Context: ctx}
	if _, err := cli.OtelMetricsIngestion.DeleteOtelMetricsIngestion(req); clienterror.IsNotFound(err) {
		setResourceNotFound(d)
		return nil
	} else if err != nil {
		return diag.Errorf("unable to delete OTel metrics ingestion config: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceOtelMetricsIngestionCustomizeDiff(
	_ context.Context, d *schema.ResourceDiff, meta any,
) error {
	m, err := buildOtelMetricsIngestion(d)
	if err != nil {
		return fmt.Errorf("unable to build OTel metrics ingestion config: %w", err)
	}
	return validateOtelMetricsIngestion(m)
}

func validateOtelMetricsIngestion(in *models.ConfigunstableOtelMetricsIngestion) error {
	// FIXME: Implement validation.

	return nil
}

func buildOtelMetricsIngestion(d ResourceGetter) (*models.ConfigunstableOtelMetricsIngestion, error) {
	out := &intschema.OtelMetricsIngestion{}
	if err := out.FromResourceData(d); err != nil {
		return nil, err
	}
	return resourceOtelMetricsIngestionToModel(out)
}
