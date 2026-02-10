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

package transport

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	httpruntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/buildinfo"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/cliutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	xswagger "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/swagger"
)

const (
	apiURLFormat   = "https://%s.chronosphere.io"
	apiTokenHeader = "api-token"

	// clientLogFileEnvVar is an env var of a filename that, if set, will result in the
	// full Chronosphere client request and response calls being logged to that file.
	clientLogFileEnvVar = "CHRONOSPHERE_TF_CLIENT_LOG_FILE"

	// OverrideTestUserAgentEnv is an environment variable which allows us to
	// dynamically override the Terraform user-agent. This should only be used
	// during testing.
	OverrideTestUserAgentEnv = "CHRONOSPHERE_OVERRIDE_TEST_USER_AGENT"
)

type Component string

var (
	ComponentTerraformProvider         Component = "chrono-terraform-provider"
	ComponentTerraformUnstableProvider Component = "chrono-terraform-unstable-provider"
	ComponentConfigSyncCmd             Component = "chrono-terraform-configsync"
)

var clientLogger *logger

func init() {
	clientLogFile := os.Getenv(clientLogFileEnvVar)
	if clientLogFile != "" {
		f, err := os.Create(clientLogFile)
		if err != nil {
			panic(fmt.Errorf("unable to open client log file %s: %v", clientLogFile, err))
		}

		clientLogger = &logger{w: f}
	}
}

type Params struct {
	Component       Component
	Org             string
	APIToken        string
	Mount           string // optional
	EntityNamespace string // optional
}

func New(p Params) (*httptransport.Runtime, error) {
	if p.Component == "" {
		return nil, errors.New("component is required")
	}
	if p.Org == "" {
		return nil, errors.New("organization is required")
	}
	if p.APIToken == "" {
		return nil, errors.New("api token is required")
	}

	rawURL := fmt.Sprintf(apiURLFormat+p.Mount, p.Org)
	apiURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse API URL: %w", err)
	}

	transport := httptransport.New(apiURL.Host, apiURL.Path, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth(apiTokenHeader, "header", p.APIToken)
	transport.Consumers[httpruntime.JSONMime] = xswagger.JSONConsumer()
	transport.Consumers[httpruntime.HTMLMime] = xswagger.TextConsumer()
	transport.Consumers[httpruntime.TextMime] = xswagger.TextConsumer()
	transport.Consumers["*/*"] = xswagger.TextConsumer() // backup, default consumer.

	if clientLogger != nil {
		transport.SetDebug(true)
		transport.SetLogger(clientLogger)
	}

	userAgent := fmt.Sprintf("%s/%s-%s (%s; %s; %s)",
		p.Component, buildinfo.Version, buildinfo.SHA, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	transport.Transport = xswagger.WithRequestIDTrailerTransport(&customHeaderRoundTripper{
		rt:              transport.Transport,
		userAgent:       userAgent,
		entityNamespace: p.EntityNamespace,
		actor:           os.Getenv(cliutil.ActorEnvVar),
	})

	return transport, nil
}

// customHeaderRoundTripper wraps an http.RoundTripper, setting the User-Agent header.
type customHeaderRoundTripper struct {
	rt              http.RoundTripper
	userAgent       string
	entityNamespace string
	actor           string
}

// RoundTrip adds the User-Agent header to the request
func (h *customHeaderRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	userAgent := h.userAgent
	if override := os.Getenv(OverrideTestUserAgentEnv); override != "" {
		userAgent = "terraform-override-" + override
	}
	r.Header.Set("User-Agent", userAgent)

	var (
		resource string
		found    bool
	)
	if resource, found = tfresource.FromContext(r.Context()); found {
		r.Header.Set("Chrono-Terraform-Resource", strings.ToLower(resource))
	}

	if h.entityNamespace != "" {
		r.Header.Set("Chrono-Entity-Namespace", h.entityNamespace)
	}

	if h.actor != "" {
		r.Header.Set("Chronosphere-Actor", h.actor)
	}

	tflog.Debug(r.Context(), "http request", map[string]any{
		"url":              r.URL.String(),
		"method":           r.Method,
		"user_agent":       userAgent,
		"tf_resource":      resource,
		"entity_namespace": h.entityNamespace,
	})
	return h.rt.RoundTrip(r)
}

var _ http.RoundTripper = &customHeaderRoundTripper{}

type logger struct {
	w io.Writer
}

func (l *logger) Printf(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(l.w, format, args...)
}

func (l *logger) Debugf(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(l.w, format, args...)
}
