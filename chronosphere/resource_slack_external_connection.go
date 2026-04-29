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

func SlackExternalConnectionFromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.SlackExternalConnection, error) {
	return slackExternalConnectionConverter{}.fromModel(m)
}

func resourceSlackExternalConnection() *schema.Resource {
	r := newGenericResource(
		"slack_external_connection",
		slackExternalConnectionConverter{},
		generatedExternalConnection{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.SlackExternalConnection,
		CustomizeDiff: r.ValidateDryRun(&SlackExternalConnectionDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// SlackExternalConnectionDryRunCount tracks how many times dry run is run during validation for testing.
var SlackExternalConnectionDryRunCount atomic.Int64

type slackExternalConnectionConverter struct{}

func (slackExternalConnectionConverter) toModel(
	n *intschema.SlackExternalConnection,
) (*models.Configv1ExternalConnection, error) {
	return &models.Configv1ExternalConnection{
		Name: n.Name,
		Slug: n.Slug,
		Slack: &models.Configv1ExternalConnectionSlackConfig{
			APIURL: n.ApiUrl,
			Token:  n.Token,
		},
	}, nil
}

func (slackExternalConnectionConverter) fromModel(
	m *models.Configv1ExternalConnection,
) (*intschema.SlackExternalConnection, error) {
	s := m.Slack
	if s == nil {
		return &intschema.SlackExternalConnection{
			Name: m.Name + " (not a slack external connection)",
			Slug: m.Slug,
		}, nil
	}
	return &intschema.SlackExternalConnection{
		Name:   m.Name,
		Slug:   m.Slug,
		ApiUrl: s.APIURL,
		Token:  s.Token,
	}, nil
}

func (slackExternalConnectionConverter) normalize(config, server *intschema.SlackExternalConnection) {
	// Only Token is redacted by the API; ApiUrl is returned normally.
	server.Token = config.Token
}
