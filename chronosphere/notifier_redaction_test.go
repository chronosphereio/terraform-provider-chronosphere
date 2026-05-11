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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
)

// These tests cover the Read-path normalize methods on alert notifier
// converters. The Config API redacts secret fields by replacing their value
// with the literal "**REDACTED**" sentinel (chronosphereio/monorepo#85581);
// without normalization, terraform refresh would persist the sentinel into
// state and produce a diff on the next plan.

func TestPreserveRedactedSecret(t *testing.T) {
	t.Run("sentinel replaced with config value", func(t *testing.T) {
		s := notifierRedactedSecret
		preserveRedactedSecret(&s, "configured")
		require.Equal(t, "configured", s)
	})

	t.Run("real server value flows through", func(t *testing.T) {
		s := "rotated-server-side"
		preserveRedactedSecret(&s, "configured")
		require.Equal(t, "rotated-server-side", s)
	})

	t.Run("sentinel with unset config stays unset", func(t *testing.T) {
		s := notifierRedactedSecret
		preserveRedactedSecret(&s, "")
		require.Equal(t, "", s)
	})

	t.Run("empty server value flows through", func(t *testing.T) {
		s := ""
		preserveRedactedSecret(&s, "configured")
		require.Equal(t, "", s)
	})
}

func TestWebhookAlertNotifierNormalize(t *testing.T) {
	t.Run("sentinel preserved with config", func(t *testing.T) {
		config := &intschema.WebhookAlertNotifier{
			BasicAuthPassword: "configured-password",
			BearerToken:       "configured-bearer",
		}
		server := &intschema.WebhookAlertNotifier{
			BasicAuthPassword: notifierRedactedSecret,
			BearerToken:       notifierRedactedSecret,
		}
		webhookAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "configured-password", server.BasicAuthPassword)
		require.Equal(t, "configured-bearer", server.BearerToken)
	})

	t.Run("real server value flows through", func(t *testing.T) {
		config := &intschema.WebhookAlertNotifier{
			BasicAuthPassword: "configured-password",
		}
		server := &intschema.WebhookAlertNotifier{
			BasicAuthPassword: "rotated-out-of-band",
		}
		webhookAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "rotated-out-of-band", server.BasicAuthPassword)
	})
}

func TestSlackAlertNotifierNormalize(t *testing.T) {
	t.Run("sentinel preserved with config", func(t *testing.T) {
		config := &intschema.SlackAlertNotifier{
			BasicAuthPassword: "configured-password",
			BearerToken:       "configured-bearer",
		}
		server := &intschema.SlackAlertNotifier{
			BasicAuthPassword: notifierRedactedSecret,
			BearerToken:       notifierRedactedSecret,
		}
		slackAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "configured-password", server.BasicAuthPassword)
		require.Equal(t, "configured-bearer", server.BearerToken)
	})

	t.Run("real server value flows through", func(t *testing.T) {
		config := &intschema.SlackAlertNotifier{BearerToken: "configured-bearer"}
		server := &intschema.SlackAlertNotifier{BearerToken: "rotated-bearer"}
		slackAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "rotated-bearer", server.BearerToken)
	})

	t.Run("api_url is not normalized", func(t *testing.T) {
		config := &intschema.SlackAlertNotifier{ApiUrl: "https://configured.example/hook"}
		server := &intschema.SlackAlertNotifier{ApiUrl: "https://rotated.example/hook"}
		slackAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "https://rotated.example/hook", server.ApiUrl)
	})
}

func TestPagerdutyAlertNotifierNormalize(t *testing.T) {
	t.Run("sentinel preserved with config", func(t *testing.T) {
		config := &intschema.PagerdutyAlertNotifier{
			BasicAuthPassword: "configured-password",
			BearerToken:       "configured-bearer",
			ServiceKey:        "configured-service-key",
			RoutingKey:        "configured-routing-key",
		}
		server := &intschema.PagerdutyAlertNotifier{
			BasicAuthPassword: notifierRedactedSecret,
			BearerToken:       notifierRedactedSecret,
			ServiceKey:        notifierRedactedSecret,
			RoutingKey:        notifierRedactedSecret,
		}
		pagerdutyAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "configured-password", server.BasicAuthPassword)
		require.Equal(t, "configured-bearer", server.BearerToken)
		require.Equal(t, "configured-service-key", server.ServiceKey)
		require.Equal(t, "configured-routing-key", server.RoutingKey)
	})

	t.Run("real server value flows through", func(t *testing.T) {
		config := &intschema.PagerdutyAlertNotifier{
			ServiceKey: "configured-service-key",
			RoutingKey: "configured-routing-key",
		}
		server := &intschema.PagerdutyAlertNotifier{
			ServiceKey: "rotated-service-key",
			RoutingKey: "rotated-routing-key",
		}
		pagerdutyAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "rotated-service-key", server.ServiceKey)
		require.Equal(t, "rotated-routing-key", server.RoutingKey)
	})
}

func TestOpsgenieAlertNotifierNormalize(t *testing.T) {
	t.Run("sentinel preserved with config", func(t *testing.T) {
		config := &intschema.OpsgenieAlertNotifier{
			ApiKey:            "configured-api-key",
			BasicAuthPassword: "configured-password",
			BearerToken:       "configured-bearer",
		}
		server := &intschema.OpsgenieAlertNotifier{
			ApiKey:            notifierRedactedSecret,
			BasicAuthPassword: notifierRedactedSecret,
			BearerToken:       notifierRedactedSecret,
		}
		opsgenieAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "configured-api-key", server.ApiKey)
		require.Equal(t, "configured-password", server.BasicAuthPassword)
		require.Equal(t, "configured-bearer", server.BearerToken)
	})

	t.Run("real server value flows through", func(t *testing.T) {
		config := &intschema.OpsgenieAlertNotifier{ApiKey: "configured-api-key"}
		server := &intschema.OpsgenieAlertNotifier{ApiKey: "rotated-api-key"}
		opsgenieAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "rotated-api-key", server.ApiKey)
	})
}

func TestVictoropsAlertNotifierNormalize(t *testing.T) {
	t.Run("sentinel preserved with config", func(t *testing.T) {
		config := &intschema.VictoropsAlertNotifier{
			ApiKey:            "configured-api-key",
			RoutingKey:        "configured-routing-key",
			BasicAuthPassword: "configured-password",
			BearerToken:       "configured-bearer",
		}
		server := &intschema.VictoropsAlertNotifier{
			ApiKey:            notifierRedactedSecret,
			RoutingKey:        notifierRedactedSecret,
			BasicAuthPassword: notifierRedactedSecret,
			BearerToken:       notifierRedactedSecret,
		}
		victoropsAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "configured-api-key", server.ApiKey)
		require.Equal(t, "configured-routing-key", server.RoutingKey)
		require.Equal(t, "configured-password", server.BasicAuthPassword)
		require.Equal(t, "configured-bearer", server.BearerToken)
	})

	t.Run("real server value flows through", func(t *testing.T) {
		config := &intschema.VictoropsAlertNotifier{
			ApiKey:     "configured-api-key",
			RoutingKey: "configured-routing-key",
		}
		server := &intschema.VictoropsAlertNotifier{
			ApiKey:     "rotated-api-key",
			RoutingKey: "rotated-routing-key",
		}
		victoropsAlertNotifierConverter{}.normalize(config, server)
		require.Equal(t, "rotated-api-key", server.ApiKey)
		require.Equal(t, "rotated-routing-key", server.RoutingKey)
	})
}
