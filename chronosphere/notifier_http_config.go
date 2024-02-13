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
