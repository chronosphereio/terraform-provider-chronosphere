// Copyright 2026 Chronosphere Inc.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
)

func TestSlackExternalConnectionNormalizeKeepsConfiguredToken(t *testing.T) {
	config := &intschema.SlackExternalConnection{Token: "xoxb-configured"}
	server := &intschema.SlackExternalConnection{Token: "**REDACTED**"}

	slackExternalConnectionConverter{}.normalize(config, server)

	assert.Equal(t, "xoxb-configured", server.Token)
}

func TestSlackExternalConnectionUnknownTokenSkipsDryRun(t *testing.T) {
	config := terraform.NewResourceConfigRaw(map[string]any{
		"name":  "Slack generated",
		"token": terraformUnknownValue,
	})
	before := SlackExternalConnectionDryRunCount.Load()

	_, err := resourceSlackExternalConnection().Diff(context.Background(), nil, config, nil)

	require.NoError(t, err)
	assert.Equal(t, before, SlackExternalConnectionDryRunCount.Load())
}
