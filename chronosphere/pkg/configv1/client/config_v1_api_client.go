// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/bucket"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/classic_dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/collection"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/dataset"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/derived_label"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/derived_metric"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/drop_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/gcp_metrics_integration"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/grafana_dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/mapping_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/monitor"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/muting_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/notification_policy"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/notifier"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/recording_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/resource_pools"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/rollup_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/service"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/service_account"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/team"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_behavior_config"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_jaeger_remote_sampling_strategy"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_metrics_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_tail_sampling_rules"
)

// Default config v1 API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new config v1 API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ConfigV1API {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new config v1 API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ConfigV1API {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new config v1 API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ConfigV1API {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ConfigV1API)
	cli.Transport = transport
	cli.Bucket = bucket.New(transport, formats)
	cli.ClassicDashboard = classic_dashboard.New(transport, formats)
	cli.Collection = collection.New(transport, formats)
	cli.Dashboard = dashboard.New(transport, formats)
	cli.Dataset = dataset.New(transport, formats)
	cli.DerivedLabel = derived_label.New(transport, formats)
	cli.DerivedMetric = derived_metric.New(transport, formats)
	cli.DropRule = drop_rule.New(transport, formats)
	cli.GcpMetricsIntegration = gcp_metrics_integration.New(transport, formats)
	cli.GrafanaDashboard = grafana_dashboard.New(transport, formats)
	cli.MappingRule = mapping_rule.New(transport, formats)
	cli.Monitor = monitor.New(transport, formats)
	cli.MutingRule = muting_rule.New(transport, formats)
	cli.NotificationPolicy = notification_policy.New(transport, formats)
	cli.Notifier = notifier.New(transport, formats)
	cli.RecordingRule = recording_rule.New(transport, formats)
	cli.ResourcePools = resource_pools.New(transport, formats)
	cli.RollupRule = rollup_rule.New(transport, formats)
	cli.Service = service.New(transport, formats)
	cli.ServiceAccount = service_account.New(transport, formats)
	cli.Team = team.New(transport, formats)
	cli.TraceBehaviorConfig = trace_behavior_config.New(transport, formats)
	cli.TraceJaegerRemoteSamplingStrategy = trace_jaeger_remote_sampling_strategy.New(transport, formats)
	cli.TraceMetricsRule = trace_metrics_rule.New(transport, formats)
	cli.TraceTailSamplingRules = trace_tail_sampling_rules.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// ConfigV1API is a client for config v1 API
type ConfigV1API struct {
	Bucket bucket.ClientService

	ClassicDashboard classic_dashboard.ClientService

	Collection collection.ClientService

	Dashboard dashboard.ClientService

	Dataset dataset.ClientService

	DerivedLabel derived_label.ClientService

	DerivedMetric derived_metric.ClientService

	DropRule drop_rule.ClientService

	GcpMetricsIntegration gcp_metrics_integration.ClientService

	GrafanaDashboard grafana_dashboard.ClientService

	MappingRule mapping_rule.ClientService

	Monitor monitor.ClientService

	MutingRule muting_rule.ClientService

	NotificationPolicy notification_policy.ClientService

	Notifier notifier.ClientService

	RecordingRule recording_rule.ClientService

	ResourcePools resource_pools.ClientService

	RollupRule rollup_rule.ClientService

	Service service.ClientService

	ServiceAccount service_account.ClientService

	Team team.ClientService

	TraceBehaviorConfig trace_behavior_config.ClientService

	TraceJaegerRemoteSamplingStrategy trace_jaeger_remote_sampling_strategy.ClientService

	TraceMetricsRule trace_metrics_rule.ClientService

	TraceTailSamplingRules trace_tail_sampling_rules.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ConfigV1API) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Bucket.SetTransport(transport)
	c.ClassicDashboard.SetTransport(transport)
	c.Collection.SetTransport(transport)
	c.Dashboard.SetTransport(transport)
	c.Dataset.SetTransport(transport)
	c.DerivedLabel.SetTransport(transport)
	c.DerivedMetric.SetTransport(transport)
	c.DropRule.SetTransport(transport)
	c.GcpMetricsIntegration.SetTransport(transport)
	c.GrafanaDashboard.SetTransport(transport)
	c.MappingRule.SetTransport(transport)
	c.Monitor.SetTransport(transport)
	c.MutingRule.SetTransport(transport)
	c.NotificationPolicy.SetTransport(transport)
	c.Notifier.SetTransport(transport)
	c.RecordingRule.SetTransport(transport)
	c.ResourcePools.SetTransport(transport)
	c.RollupRule.SetTransport(transport)
	c.Service.SetTransport(transport)
	c.ServiceAccount.SetTransport(transport)
	c.Team.SetTransport(transport)
	c.TraceBehaviorConfig.SetTransport(transport)
	c.TraceJaegerRemoteSamplingStrategy.SetTransport(transport)
	c.TraceMetricsRule.SetTransport(transport)
	c.TraceTailSamplingRules.SetTransport(transport)
}
