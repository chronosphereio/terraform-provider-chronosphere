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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/collection"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// DataCollectionFromModel maps an API model to an intschema model.
func DataCollectionFromModel(b *models.Configv1Collection) *intschema.DataCollection {
	return &intschema.DataCollection{
		Name:        b.Name,
		Slug:        b.Slug,
		Description: b.Description,
	}
}

// dataSourceCollection creates a schema for a collection data source.
func dataSourceCollection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCollectionRead,
		Schema:      tfschema.DataCollection,
	}
}

// dataSourceCollectionRead reads a collection data source.
func dataSourceCollectionRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	cli := getConfigClient(meta)

	b := &intschema.DataCollection{}
	if err := b.FromResourceData(d); err != nil {
		return diag.FromErr(err)
	}

	collection, diags := findCollection(ctx, cli, b)
	if diags != nil {
		return diags
	}

	b = DataCollectionFromModel(collection)
	d.SetId(collection.Slug)
	return b.ToResourceData(d)
}

func findCollection(ctx context.Context, cli *client.ConfigV1API, b *intschema.DataCollection) (*models.Configv1Collection, diag.Diagnostics) {
	resp, err := cli.Collection.ReadCollection(&collection.ReadCollectionParams{
		Context: ctx,
		Slug:    b.Slug,
	})
	if err != nil {
		return nil, diag.Errorf("unable to read collection: %v", clienterror.Wrap(err))
	}

	return resp.GetPayload().Collection, nil
}
