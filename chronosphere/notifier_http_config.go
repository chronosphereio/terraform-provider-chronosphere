// Copyright 2023 Chronosphere Inc.
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
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

type notifierHTTPConfig struct {
	basicAuthUsername     string
	basicAuthPassword     string
	tlsInsecureSkipVerify bool
	bearerToken           string
	proxyURL              string
}

func (c notifierHTTPConfig) toModel() *configv1.NotifierHTTPConfig {
	m := &configv1.NotifierHTTPConfig{
		BearerToken: c.bearerToken,
		ProxyURL:    c.proxyURL,
		TLSConfig: &configv1.HTTPConfigTLSConfig{
			InsecureSkipVerify: c.tlsInsecureSkipVerify,
		},
	}
	if c.basicAuthUsername != "" && c.basicAuthPassword != "" {
		m.BasicAuth = &configv1.HTTPConfigBasicAuth{
			Username: c.basicAuthUsername,
			Password: c.basicAuthPassword,
		}
	}
	return m
}

func notifierHTTPConfigFromModel(m *configv1.NotifierHTTPConfig) notifierHTTPConfig {
	c := notifierHTTPConfig{}
	if m.BasicAuth != nil {
		c.basicAuthUsername = m.BasicAuth.Username
		c.basicAuthPassword = m.BasicAuth.Password
	}
	if m.TLSConfig != nil {
		c.tlsInsecureSkipVerify = m.TLSConfig.InsecureSkipVerify
	}
	c.bearerToken = m.BearerToken
	c.proxyURL = m.ProxyURL
	return c
}
