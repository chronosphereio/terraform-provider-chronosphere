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
	"os"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/cliutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/apiclients"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/transport"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/unstable"
)

// LocalName is the provider name that is used for local installation.
const LocalName = "local/chronosphereio/chronosphere"

// Provider returns the chronosphere provider.
func Provider() *schema.Provider {
	allResources := map[string]*schema.Resource{
		"chronosphere_classic_dashboard":                     resourceClassicDashboard(),
		"chronosphere_bucket":                                resourceBucket(),
		"chronosphere_notification_policy":                   resourceNotificationPolicy(),
		"chronosphere_monitor":                               resourceMonitor(),
		"chronosphere_pagerduty_alert_notifier":              resourcePagerdutyAlertNotifier(),
		"chronosphere_slack_alert_notifier":                  resourceSlackAlertNotifier(),
		"chronosphere_webhook_alert_notifier":                resourceWebhookAlertNotifier(),
		"chronosphere_opsgenie_alert_notifier":               resourceOpsGenieAlertNotifier(),
		"chronosphere_email_alert_notifier":                  resourceEmailAlertNotifier(),
		"chronosphere_victorops_alert_notifier":              resourceVictorOpsAlertNotifier(),
		"chronosphere_blackhole_alert_notifier":              resourceBlackHoleAlertNotifier(),
		"chronosphere_gcp_metrics_integration":               resourceGcpMetricsIntegration(),
		"chronosphere_recording_rule":                        resourceRecordingRule(),
		"chronosphere_rollup_rule":                           resourceRollupRule(),
		"chronosphere_drop_rule":                             resourceDropRule(),
		"chronosphere_mapping_rule":                          resourceMappingRule(),
		"chronosphere_team":                                  resourceTeam(),
		"chronosphere_collection":                            resourceCollection(),
		"chronosphere_derived_metric":                        resourceDerivedMetric(),
		"chronosphere_resource_pools_config":                 resourceResourcePoolsConfig(),
		"chronosphere_dashboard":                             resourceDashboard(),
		"chronosphere_trace_metrics_rule":                    resourceTraceMetricsRule(),
		"chronosphere_trace_jaeger_remote_sampling_strategy": resourceTraceJRSStrategy(),
		"chronosphere_trace_tail_sampling_rules":             resourceTraceTailSamplingRules(),
		"chronosphere_service_account":                       resourceServiceAccount(),
		"chronosphere_derived_label":                         resourceDerivedLabel(),
		"chronosphere_dataset":                               resourceDataset(),
	}

	// Apply common CRUD wrappers to all resources.
	composeMutationsReadResources(allResources)
	mutexProtectedResources(allResources)

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"org": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{cliutil.OrgEnvVar, cliutil.OrgNameEnvVar}, nil),
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(cliutil.ApiTokenEnvVar, nil),
			},
			"unstable": {
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: func() (any, error) {
					return os.Getenv(unstable.Env) == "1", nil
				},
			},
			"entity_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(cliutil.EntityNamespaceEnvVar, ""),
			},
		},
		ResourcesMap: allResources,
		DataSourcesMap: map[string]*schema.Resource{
			"chronosphere_bucket":     dataSourceBucket(),
			"chronosphere_collection": dataSourceCollection(),
			"chronosphere_service":    dataSourceService(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func getConfigClient(meta any) *configv1.Client {
	return meta.(apiclients.Clients).ConfigV1
}

func getConfigUnstableClient(meta any) *configunstable.Client {
	return meta.(apiclients.Clients).ConfigUnstable
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	org := d.Get("org").(string)
	apiToken := d.Get("api_token").(string)
	entityNamespace := d.Get("entity_namespace").(string)
	unstable.Set(ctx, d.Get("unstable").(bool))

	configUnstableClient, err := configunstable.NewClient(transport.ComponentTerraformUnstableProvider, org, apiToken, entityNamespace)
	if err != nil {
		return nil, diag.Errorf("unable to construct Chronosphere unstable config client: %v", err)
	}

	configClient, err := configv1.NewClient(transport.ComponentTerraformProvider, org, apiToken, entityNamespace)
	if err != nil {
		return nil, diag.Errorf("unable to construct Chronosphere config v1 client: %v", err)
	}

	return apiclients.Clients{
		ConfigUnstable: configUnstableClient,
		ConfigV1:       configClient,
	}, nil
}

// composeMutationsReadResources updates all the given resources' CreateContext
// and UpdateContext to ReadContext after the mutation operation.
// The primary purpose of this wrapper (vs inlining the calls) is to mark
// only Update failures using ResourceData.Partial, while not treating
// Read failures as partial state updates.
// See internal ticket 28275 for context
func composeMutationsReadResources(resources map[string]*schema.Resource) {
	for _, r := range resources {
		create := r.CreateContext
		update := r.UpdateContext
		read := r.ReadContext

		r.CreateContext = func(ctx context.Context, d *schema.ResourceData, meta any) (retErr diag.Diagnostics) {
			if err := create(ctx, d, meta); err != nil {
				return err
			}

			return read(ctx, d, meta)
		}

		// Update is optional for resources that only support ForceNew for all attributes.
		if update != nil {
			r.UpdateContext = func(ctx context.Context, d *schema.ResourceData, meta any) (retErr diag.Diagnostics) {
				if err := update(ctx, d, meta); err != nil {
					// When an update fails, even though an error is returned, Terraform updates the state file.
					// This causes an unnecessary "object changed outside of Terraform"
					// Issue + workaround: https://github.com/hashicorp/terraform-plugin-sdk/issues/476
					d.Partial(true)
					return err
				}

				return read(ctx, d, meta)
			}
		}
	}
}

// mutexProtectedResources updates all the given resources, intercepting each
// create, update, and delete resource function to first acquire a provider-scoped mutex.
// This guards against the server being extremely sensitive to even low levels of concurrent request
// to change alert-related config.
func mutexProtectedResources(resources map[string]*schema.Resource) {
	if allowConcurrentMutations() {
		return
	}

	var mu sync.Mutex

	// wrap the function with the mutex above
	locked := func(fn func(context.Context, *schema.ResourceData, any) diag.Diagnostics) func(context.Context, *schema.ResourceData, any) diag.Diagnostics {
		return func(ctx context.Context, data *schema.ResourceData, i any) diag.Diagnostics {
			mu.Lock()
			defer mu.Unlock()

			return fn(ctx, data, i)
		}
	}

	for _, resource := range resources {
		resource.CreateContext = locked(resource.CreateContext)
		// Update is optional for resources that only support ForceNew for all attributes.
		if resource.UpdateContext != nil {
			resource.UpdateContext = locked(resource.UpdateContext)
		}
		resource.DeleteContext = locked(resource.DeleteContext)
	}
}

// allowConcurrentMutations determines whether to enable concurrency when the CONCURRENT_MUTATIONS=1
// environment variable is set.
func allowConcurrentMutations() bool {
	return os.Getenv("CONCURRENT_MUTATIONS") == "1"
}

type resourceFunc interface {
	~func(context.Context, *schema.ResourceData, any) diag.Diagnostics
}
