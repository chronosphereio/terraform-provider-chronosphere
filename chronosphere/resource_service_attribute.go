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
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// ServiceAttributeFromModel maps an API model into an intschema model.
func ServiceAttributeFromModel(m *models.Configv1ServiceAttribute) (*intschema.ServiceAttribute, error) {
	return (serviceAttributeConverter{}).fromModel(m)
}

func resourceServiceAttribute() *schema.Resource {
	r := newGenericResource(
		"service_attribute",
		serviceAttributeConverter{},
		generatedServiceAttribute{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.ServiceAttribute,
		CustomizeDiff: r.ValidateDryRun(&ServiceAttributeDryRunCount),
		Description:   "Service attributes configuration for associating metadata with services",
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// ServiceAttributeDryRunCount tracks how many times dry run is run during validation for testing.
var ServiceAttributeDryRunCount atomic.Int64

type serviceAttributeConverter struct{}

func (serviceAttributeConverter) toModel(
	r *intschema.ServiceAttribute,
) (*models.Configv1ServiceAttribute, error) {
	if r == nil {
		return nil, fmt.Errorf("service attribute cannot be nil")
	}
	return &models.Configv1ServiceAttribute{
		ServiceSlug:            r.ServiceSlug,
		Name:                   r.Name,
		Description:            r.Description,
		TeamSlug:               r.TeamId.Slug(),
		NotificationPolicySlug: r.NotificationPolicyId.Slug(),
	}, nil
}

func (serviceAttributeConverter) fromModel(
	m *models.Configv1ServiceAttribute,
) (*intschema.ServiceAttribute, error) {
	if m == nil {
		return nil, fmt.Errorf("service attribute model cannot be nil")
	}
	return &intschema.ServiceAttribute{
		ServiceSlug:          m.ServiceSlug,
		Name:                 m.Name,
		Description:          m.Description,
		TeamId:               tfid.Slug(m.TeamSlug),
		NotificationPolicyId: tfid.Slug(m.NotificationPolicySlug),
	}, nil
}
