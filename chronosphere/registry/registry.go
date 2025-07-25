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

package registry

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// AllEntities returns all entities which are registered against the given API.
// It includes Singletons, and is used for generating CRUD bindings.
func AllEntities(api API) []Resource {
	return entities(api, true /* includeSingletons */)
}

// NamedEntities returns all entities which are registered against the given API.
// It only includes non-singleton entities and is useful for use-cases that
// don't make sense for singletons, such as generating List bindings.
func NamedEntities(api API) []Resource {
	return entities(api, false /* includeSingletons */)
}

func entities(api API, includeSingletons bool) []Resource {
	// Some of our resources share the same underlying entity (i.e. notifier),
	// so we use a map to remove duplicates.
	m := make(map[string]Resource)
	for _, r := range Resources {
		if r.API != api {
			continue
		}
		if r.SingletonID != "" && !includeSingletons {
			continue
		}
		m[r.Entity] = r
	}
	ents := maps.Keys(m)
	sort.Strings(ents)
	var res []Resource
	for _, ent := range ents {
		res = append(res, m[ent])
	}
	return res
}

// API is an API namespace.
type API string

const (
	V1       API = "v1"
	Unstable API = "unstable"
	Legacy   API = "legacy"
)

// Resource consolidates metadata for a Terraform resource.
type Resource struct {
	// Name is the name of the resource (without a "chronosphere_" prefix).
	Name string

	// Entity is the name of the entity the resource wraps.
	Entity string

	// API is the API namespace which backs the resource.
	API API

	// Schema is the tfschema which defines the resource.
	Schema map[string]*schema.Schema

	// Only set if the resource is a singleton.
	SingletonID string

	// DisableDryRun is a flag to disable dry run for a resource.
	DisableDryRun bool

	// DisableExportImport silently disables all paginated list helpers by
	// returning no results, thus preventing export-config/import-state from
	// observing any entities. Only useful when promoting an entity from
	// unstable to V1 and you need to temporarily prevent
	// export-config/import-state tests from hitting a disabled unstable API.
	DisableExportImport bool

	// UpdateUnsupported disables generateresource from creating an update helper
	// for the given resource type. This is currently designed with ServiceAccounts
	// in mind, as the resource only support CRD+list.
	UpdateUnsupported bool
}

func (r Resource) Type() string {
	return "chronosphere_" + r.Name
}

func (r Resource) validate() error {
	if r.Name == "" {
		return errors.New("Name is required")
	}
	if r.Entity == "" {
		return errors.New("Entity is required")
	}
	switch r.API {
	case V1, Unstable, Legacy:
		// valid
	default:
		return fmt.Errorf("invalid API: %q", r.API)
	}
	if r.Schema == nil {
		return errors.New("Schema is required")
	}
	if r.SingletonID != "" && r.API == Legacy {
		return errors.New("cannot set SingletonID when API=Legacy")
	}

	return nil
}

func mustValidate(rs []Resource) []Resource {
	names := make(map[string]bool)
	schemas := make(map[uintptr]bool)

	for i, r := range rs {
		if err := r.validate(); err != nil {
			panic(fmt.Errorf("resource %d is invalid: %v", i, err))
		}

		if names[r.Name] {
			panic(fmt.Errorf("resource %d Name already registered: %q", i, r.Name))
		}
		names[r.Name] = true

		sptr := reflect.ValueOf(r.Schema).Pointer()
		if schemas[sptr] {
			panic(fmt.Errorf("resource %d Schema already registered: %v", i, sptr))
		}
		schemas[sptr] = true
	}

	return rs
}

var Resources = mustValidate([]Resource{
	{
		Name:   "blackhole_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.BlackholeAlertNotifier,
	},
	{
		Name:   "bucket",
		Entity: "Bucket",
		API:    V1,
		Schema: tfschema.Bucket,
	},
	{
		Name:   "collection",
		Entity: "Collection",
		API:    V1,
		Schema: tfschema.Collection,
	},
	{
		Name:   "dashboard",
		Entity: "Dashboard",
		API:    V1,
		Schema: tfschema.Dashboard,
	},
	{
		Name:   "dataset",
		Entity: "Dataset",
		API:    V1,
		Schema: tfschema.Dataset,
	},
	{
		Name:   "derived_label",
		Entity: "DerivedLabel",
		API:    V1,
		Schema: tfschema.DerivedLabel,
	},
	{
		Name:   "derived_metric",
		Entity: "DerivedMetric",
		API:    V1,
		Schema: tfschema.DerivedMetric,
	},
	{
		Name:   "drop_rule",
		Entity: "DropRule",
		API:    V1,
		Schema: tfschema.DropRule,
	},
	{
		Name:   "email_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.EmailAlertNotifier,
	},
	{
		Name:   "classic_dashboard",
		Entity: "GrafanaDashboard",
		API:    V1,
		Schema: tfschema.ClassicDashboard,
	},
	{
		Name:   "gcp_metrics_integration",
		Entity: "GcpMetricsIntegration",
		API:    V1,
		Schema: tfschema.GcpMetricsIntegration,
	},
	{
		Name:   "mapping_rule",
		Entity: "MappingRule",
		API:    V1,
		Schema: tfschema.MappingRule,
	},
	{
		Name:   "monitor",
		Entity: "Monitor",
		API:    V1,
		Schema: tfschema.Monitor,
	},
	{
		Name:   "notification_policy",
		Entity: "NotificationPolicy",
		API:    V1,
		Schema: tfschema.NotificationPolicy,
		// N.B. Notification Policies explicitly don't support ownership transfers.
	},
	{
		Name:   "opsgenie_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.OpsgenieAlertNotifier,
	},
	{
		Name:        "otel_metrics_ingestion",
		Entity:      "OtelMetricsIngestion",
		API:         V1,
		Schema:      tfschema.OtelMetricsIngestion,
		SingletonID: "otel_metrics_ingestion_singleton",
	},
	{
		Name:   "pagerduty_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.PagerdutyAlertNotifier,
	},
	{
		Name:   "recording_rule",
		Entity: "RecordingRule",
		API:    V1,
		Schema: tfschema.RecordingRule,
	},
	{
		Name:        "resource_pools_config",
		Entity:      "ResourcePools",
		API:         V1,
		Schema:      tfschema.ResourcePoolsConfig,
		SingletonID: "resource_pool_singleton",
	},
	{
		Name:   "rollup_rule",
		Entity: "RollupRule",
		API:    V1,
		Schema: tfschema.RollupRule,
	},
	{
		Name:              "service_account",
		Entity:            "ServiceAccount",
		API:               V1,
		Schema:            tfschema.ServiceAccount,
		UpdateUnsupported: true,
	},
	{
		Name:   "slack_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.SlackAlertNotifier,
	},
	{
		Name:   "slo",
		Entity: "SLO",
		API:    V1,
		Schema: tfschema.Slo,
	},
	{
		Name:   "team",
		Entity: "Team",
		API:    V1,
		Schema: tfschema.Team,
	},
	{
		Name:   "trace_metrics_rule",
		Entity: "TraceMetricsRule",
		API:    V1,
		Schema: tfschema.TraceMetricsRule,
	},
	{
		Name:   "trace_jaeger_remote_sampling_strategy",
		Entity: "TraceJaegerRemoteSamplingStrategy",
		API:    V1,
		Schema: tfschema.TraceJaegerRemoteSamplingStrategy,
	},
	{
		Name:   "victorops_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.VictoropsAlertNotifier,
	},
	{
		Name:   "webhook_alert_notifier",
		Entity: "Notifier",
		API:    V1,
		Schema: tfschema.WebhookAlertNotifier,
	},
	{
		Name:        "trace_tail_sampling_rules",
		Entity:      "TraceTailSamplingRules",
		API:         V1,
		Schema:      tfschema.TraceTailSamplingRules,
		SingletonID: "trace_tail_sampling_singleton",
	},
	{
		Name:   "logscale_alert",
		Entity: "LogScaleAlert",
		API:    V1,
		Schema: tfschema.LogscaleAlert,
	},
	{
		Name:   "logscale_action",
		Entity: "LogScaleAction",
		API:    V1,
		Schema: tfschema.LogscaleAction,
	},
	{
		Name:        "log_allocation_config",
		Entity:      "LogAllocationConfig",
		API:         V1,
		Schema:      tfschema.LogAllocationConfig,
		SingletonID: "log_allocation_config_singleton",
	},
	{
		Name:        "log_ingest_config",
		Entity:      "LogIngestConfig",
		API:         V1,
		Schema:      tfschema.LogIngestConfig,
		SingletonID: "log_ingest_config_singleton",
	},
	{
		Name:        "consumption_config",
		Entity:      "ConsumptionConfig",
		API:         Unstable,
		Schema:      tfschema.ConsumptionConfig,
		SingletonID: "consumption_config_singleton",
	},
})
