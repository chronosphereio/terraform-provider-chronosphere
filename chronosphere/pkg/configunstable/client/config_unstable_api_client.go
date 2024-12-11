// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/link_template"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/log_allocation_config"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/log_control_config"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/noop_entity"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/object_discovery_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/saved_trace_search"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/service"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/slo"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/sync_prometheus"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/trace_behavior"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/trace_jaeger_remote_sampling_strategy"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/trace_tail_sampling_rules"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/client/trace_top_tag_config"
)

// Default config unstable API HTTP client.
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

// NewHTTPClient creates a new config unstable API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ConfigUnstableAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new config unstable API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ConfigUnstableAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new config unstable API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ConfigUnstableAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ConfigUnstableAPI)
	cli.Transport = transport
	cli.Dashboard = dashboard.New(transport, formats)
	cli.LinkTemplate = link_template.New(transport, formats)
	cli.LogAllocationConfig = log_allocation_config.New(transport, formats)
	cli.LogControlConfig = log_control_config.New(transport, formats)
	cli.NoopEntity = noop_entity.New(transport, formats)
	cli.ObjectDiscoveryRule = object_discovery_rule.New(transport, formats)
	cli.SavedTraceSearch = saved_trace_search.New(transport, formats)
	cli.Service = service.New(transport, formats)
	cli.SLO = slo.New(transport, formats)
	cli.SyncPrometheus = sync_prometheus.New(transport, formats)
	cli.TraceBehavior = trace_behavior.New(transport, formats)
	cli.TraceJaegerRemoteSamplingStrategy = trace_jaeger_remote_sampling_strategy.New(transport, formats)
	cli.TraceTailSamplingRules = trace_tail_sampling_rules.New(transport, formats)
	cli.TraceTopTagConfig = trace_top_tag_config.New(transport, formats)
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

// ConfigUnstableAPI is a client for config unstable API
type ConfigUnstableAPI struct {
	Dashboard dashboard.ClientService

	LinkTemplate link_template.ClientService

	LogAllocationConfig log_allocation_config.ClientService

	LogControlConfig log_control_config.ClientService

	NoopEntity noop_entity.ClientService

	ObjectDiscoveryRule object_discovery_rule.ClientService

	SavedTraceSearch saved_trace_search.ClientService

	Service service.ClientService

	SLO slo.ClientService

	SyncPrometheus sync_prometheus.ClientService

	TraceBehavior trace_behavior.ClientService

	TraceJaegerRemoteSamplingStrategy trace_jaeger_remote_sampling_strategy.ClientService

	TraceTailSamplingRules trace_tail_sampling_rules.ClientService

	TraceTopTagConfig trace_top_tag_config.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ConfigUnstableAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Dashboard.SetTransport(transport)
	c.LinkTemplate.SetTransport(transport)
	c.LogAllocationConfig.SetTransport(transport)
	c.LogControlConfig.SetTransport(transport)
	c.NoopEntity.SetTransport(transport)
	c.ObjectDiscoveryRule.SetTransport(transport)
	c.SavedTraceSearch.SetTransport(transport)
	c.Service.SetTransport(transport)
	c.SLO.SetTransport(transport)
	c.SyncPrometheus.SetTransport(transport)
	c.TraceBehavior.SetTransport(transport)
	c.TraceJaegerRemoteSamplingStrategy.SetTransport(transport)
	c.TraceTailSamplingRules.SetTransport(transport)
	c.TraceTopTagConfig.SetTransport(transport)
}
