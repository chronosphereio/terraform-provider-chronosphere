// Copyright 2023 Chronosphere Inc.
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/bucket"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// DataBucketFromModel maps an API model to an intschema model.
func DataBucketFromModel(b *models.Configv1Bucket) *intschema.DataBucket {
	return &intschema.DataBucket{
		Name:        b.Name,
		Slug:        b.Slug,
		Description: b.Description,
		Labels:      b.Labels,
	}
}

// dataSourceBucket creates a schema for a bucket data source.
func dataSourceBucket() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBucketRead,
		Schema:      tfschema.DataBucket,
	}
}

// dataSourceBucketRead reads a bucket data source.
func dataSourceBucketRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	cli := getConfigClient(meta)

	b := &intschema.DataBucket{}
	if err := b.FromResourceData(d); err != nil {
		return diag.FromErr(err)
	}

	bucket, diags := findBucket(ctx, cli, b)
	if diags != nil {
		return diags
	}

	b = DataBucketFromModel(bucket)
	d.SetId(bucket.Slug)
	return b.ToResourceData(d)
}

func findBucket(ctx context.Context, cli *client.ConfigV1API, b *intschema.DataBucket) (*models.Configv1Bucket, diag.Diagnostics) {
	p := &bucket.ListBucketsParams{
		Context: ctx,
	}

	// Only one will be set, but prefer slug since it's immutable.
	if b.Slug != "" {
		p.Slugs = []string{b.Slug}
	} else if b.Name != "" {
		p.Names = []string{b.Name}
	}

	resp, err := cli.Bucket.ListBuckets(p)
	if err != nil {
		return nil, diag.Errorf("unable to list buckets: %v", clienterror.Wrap(err))
	}

	if numBuckets := len(resp.GetPayload().Buckets); numBuckets != 1 {
		field, value := "name", b.Name
		if b.Slug != "" {
			field, value = "slug", b.Slug
		}

		if numBuckets == 0 {
			return nil, diag.Errorf("bucket with %s `%s` not found", field, value)
		}
		return nil, diag.Errorf("found %v buckets matching %s `%s`", numBuckets, field, value)
	}

	return resp.GetPayload().Buckets[0], nil
}
