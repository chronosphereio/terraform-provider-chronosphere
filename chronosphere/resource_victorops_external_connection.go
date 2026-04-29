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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func VictoropsExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.VictoropsExternalConnection, error) {
	return victoropsExternalConnectionConverter{}.fromModel(m)
}

func resourceVictoropsExternalConnection() *schema.Resource {
	r := newGenericResource(
		"victorops_external_connection",
		victoropsExternalConnectionConverter{},
		generatedExternalConnection{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.VictoropsExternalConnection,
		CustomizeDiff: r.ValidateDryRun(&VictorOpsExternalConnectionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// VictorOpsExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var VictorOpsExternalConnectionDryRunCount atomic.Int64

type victoropsExternalConnectionConverter struct{}

func (victoropsExternalConnectionConverter) toModel(
	n *intschema.VictoropsExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		VictorOps: &models.Configv1ExternalConnectionVictorOpsConfig{
			APIKey: n.ApiKey,
			APIURL: n.ApiUrl,
		},
	}, nil
}

func (victoropsExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.VictoropsExternalConnection, error) {
	v := m.VictorOps
	if v == nil {
		return &intschema.VictoropsExternalConnection{
			Name: m.Name + " (not a victorops external connection)",
			Slug: m.Slug,
		}, nil
	}
	return &intschema.VictoropsExternalConnection{
		Name:   m.Name,
		Slug:   m.Slug,
		ApiKey: v.APIKey,
		ApiUrl: v.APIURL,
	}, nil
}

func (victoropsExternalConnectionConverter) normalize(config, server *intschema.VictoropsExternalConnection) {
	server.ApiKey = config.ApiKey
}
