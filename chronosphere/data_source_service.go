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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/service"
	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataServiceFromModel maps an API model to an intschema model.
func DataServiceFromModel(s *configmodels.Configv1Service) *intschema.DataService {
	return &intschema.DataService{
		Slug:        s.Slug,
		Name:        s.Name,
		Description: s.Description,
	}
}

// datasourceService creates a schema for a service data source.
func dataSourceService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServiceRead,
		Schema:      tfschema.DataService,
	}
}

// datasourceServiceRead reads a service data source.
func dataSourceServiceRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	cli := getConfigClient(meta)

	ds := &intschema.DataService{}
	if err := ds.FromResourceData(d); err != nil {
		return diag.FromErr(err)
	}

	svc, err := lookupService(ctx, cli.Service, ds)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(CollectionTypeSlugToID(configmodels.Configv1CollectionReferenceTypeSERVICE, svc.Slug))
	ds = DataServiceFromModel(svc)
	return ds.ToResourceData(d)
}

func lookupService(ctx context.Context, cli configv1.ClientService, s *intschema.DataService) (*configmodels.Configv1Service, error) {
	resp, err := cli.ReadService(&configv1.ReadServiceParams{
		Context: ctx,
		Slug:    s.Slug,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetPayload().Service, nil
}
