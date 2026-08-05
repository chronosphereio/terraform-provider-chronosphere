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
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/shared/pkg/container/set"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

func TestSyntheticTestHTTPConversion(t *testing.T) {
	in := &intschema.SyntheticTest{
		Name:         "checkout availability",
		Slug:         "checkout-availability",
		CollectionId: tfid.Slug("web"),
		TestType:     "HTTP",
		Status:       "ENABLED",
		Locations:    []string{"GCP_US_OREGON"},
		IntervalSecs: 60,
		TimeoutSecs:  30,
		Labels:       map[string]string{"team": "web"},
		RetryConfig: &intschema.SyntheticTestRetryConfig{
			MaxRetries:      2,
			RetryIntervalMs: 500,
		},
		MonitorConfig: &intschema.SyntheticTestMonitorConfig{
			FailingDurationSecs:  300,
			MinFailingLocations:  1,
			NotificationPolicyId: tfid.Slug("web-oncall"),
		},
		HttpTest: &intschema.SyntheticTestHttpTest{
			Url:         "https://example.com/checkout",
			Method:      "POST",
			HttpVersion: "HTTP_2",
			ContentType: "APPLICATION_JSON",
			RequestBody: `{"probe":true}`,
			Headers: []intschema.SyntheticTestHttpTestHeaders{
				{Name: "X-Probe", Value: "1"},
			},
			Assertions: []intschema.SyntheticTestHttpTestAssertions{
				{
					StatusCodeAssertion: &intschema.SyntheticTestHttpTestAssertionsStatusCodeAssertion{
						Operator: "EQUALS",
						Target:   "200",
					},
				},
				{
					ResponseTimeAssertion: &intschema.SyntheticResponseTimeAssertion{
						Operator: "LESS_THAN",
						TargetMs: 1500,
						Scope:    "WITHOUT_DNS",
					},
				},
			},
		},
	}

	m, err := syntheticTestConverter{}.toModel(in)
	require.NoError(t, err)

	assert.Equal(t, "web", m.CollectionSlug)
	assert.Equal(t, "web-oncall", m.MonitorConfig.NotificationPolicySlug)
	assert.Equal(t, models.SyntheticTestTestTypeHTTP, m.TestType)
	assert.Equal(t, []models.SyntheticTestTestLocation{models.SyntheticTestTestLocationGCPUSOREGON}, m.Locations)
	assert.Equal(t, models.HTTPTestConfigHTTPVersionHTTPVERSIONHTTP2, m.HTTPTest.HTTPVersion)
	assert.Equal(t, models.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONJSON, m.HTTPTest.ContentType)
	assert.Equal(t, `{"probe":true}`, string(m.HTTPTest.RequestBody))
	require.Len(t, m.HTTPTest.Assertions, 2)
	assert.Equal(t, "200", m.HTTPTest.Assertions[0].StatusCodeAssertion.Target)
	assert.Equal(t, int32(1500), m.HTTPTest.Assertions[1].ResponseTimeAssertion.TargetMs)

	out, err := syntheticTestConverter{}.fromModel(m)
	require.NoError(t, err)
	assert.Equal(t, in, out)
}

// Credentials are sent on writes but never read back into the schema struct,
// so nothing can carry them into Terraform state.
func TestSyntheticTestAuthSecretsAreWriteOnly(t *testing.T) {
	auth := &intschema.SyntheticTestHttpTestAuthentication{
		BasicAuth: &intschema.SyntheticTestHttpTestAuthenticationBasicAuth{
			Username:   "probe",
			PasswordWo: "hunter2",
		},
		ApiTokenAuth: &intschema.SyntheticTestHttpTestAuthenticationApiTokenAuth{
			Key:     "Authorization",
			TokenWo: "Bearer secret",
		},
		ClientCertificate: &intschema.SyntheticTestHttpTestAuthenticationClientCertificate{
			Certificate:  "-----BEGIN CERTIFICATE-----",
			PrivateKeyWo: "-----BEGIN PRIVATE KEY-----",
		},
		Oauth2ClientCredentials: &intschema.SyntheticTestHttpTestAuthenticationOauth2ClientCredentials{
			ClientId:       "client",
			ClientSecretWo: "cc-secret",
			Common: &intschema.SyntheticOAuth2Common{
				AccessTokenUrl:  "https://example.com/token",
				TokenAuthMethod: "REQUEST_BODY",
				Scopes:          []string{"read"},
			},
		},
		Oauth2ResourceOwnerPassword: &intschema.SyntheticTestHttpTestAuthenticationOauth2ResourceOwnerPassword{
			Username:       "probe",
			PasswordWo:     "rop-password",
			ClientId:       "client",
			ClientSecretWo: "rop-secret",
			Common: &intschema.SyntheticOAuth2Common{
				AccessTokenUrl: "https://example.com/token",
			},
		},
	}

	m := syntheticHTTPAuthToModel(auth)
	assert.Equal(t, "hunter2", m.BasicAuth.Password)
	assert.Equal(t, "Bearer secret", m.APITokenAuth.Token)
	assert.Equal(t, "-----BEGIN PRIVATE KEY-----", m.ClientCertificate.PrivateKey)
	assert.Equal(t, "cc-secret", m.Oauth2ClientCredentials.ClientSecret)
	assert.Equal(t, "rop-password", m.Oauth2ResourceOwnerPassword.Password)
	assert.Equal(t, "rop-secret", m.Oauth2ResourceOwnerPassword.ClientSecret)
	assert.Equal(t, models.ConfigunstableOAuth2TokenAuthMethodOTAMREQUESTBODY, m.Oauth2ClientCredentials.Common.TokenAuthMethod)

	// A read returns the redacted sentinel; none of it survives into the schema.
	m.BasicAuth.Password = "**REDACTED**"
	m.APITokenAuth.Token = "**REDACTED**"
	m.ClientCertificate.PrivateKey = "**REDACTED**"
	m.Oauth2ClientCredentials.ClientSecret = "**REDACTED**"
	m.Oauth2ResourceOwnerPassword.Password = "**REDACTED**"
	m.Oauth2ResourceOwnerPassword.ClientSecret = "**REDACTED**"

	got := syntheticHTTPAuthFromModel(m)
	assert.Empty(t, got.BasicAuth.PasswordWo)
	assert.Empty(t, got.ApiTokenAuth.TokenWo)
	assert.Empty(t, got.ClientCertificate.PrivateKeyWo)
	assert.Empty(t, got.Oauth2ClientCredentials.ClientSecretWo)
	assert.Empty(t, got.Oauth2ResourceOwnerPassword.PasswordWo)
	assert.Empty(t, got.Oauth2ResourceOwnerPassword.ClientSecretWo)

	// Non-secret siblings still round-trip.
	assert.Equal(t, "probe", got.BasicAuth.Username)
	assert.Equal(t, "Authorization", got.ApiTokenAuth.Key)
	assert.Equal(t, "-----BEGIN CERTIFICATE-----", got.ClientCertificate.Certificate)
	assert.Equal(t, "client", got.Oauth2ClientCredentials.ClientId)
	assert.Equal(t, "REQUEST_BODY", got.Oauth2ClientCredentials.Common.TokenAuthMethod)
	assert.Equal(t, []string{"read"}, got.Oauth2ClientCredentials.Common.Scopes)
}

func TestSyntheticTestProtocolConversions(t *testing.T) {
	t.Run("dns", func(t *testing.T) {
		in := &intschema.SyntheticTestDnsTest{
			Domain:        "example.com",
			DnsServer:     "8.8.8.8",
			DnsServerPort: 53,
			Assertions: []intschema.SyntheticTestDnsTestAssertions{{
				DnsRecordAssertion: &intschema.SyntheticTestDnsTestAssertionsDnsRecordAssertion{
					RecordType: "A",
					Operator:   "EQUALS",
					Target:     "93.184.216.34",
					MatchScope: "AT_LEAST_ONE",
				},
			}},
		}
		assert.Equal(t, in, syntheticDNSTestFromModel(syntheticDNSTestToModel(in)))
	})

	t.Run("tcp", func(t *testing.T) {
		in := &intschema.SyntheticTestTcpTest{
			Host: "example.com",
			Port: 443,
			Assertions: []intschema.SyntheticTestTcpTestAssertions{{
				ConnectionAssertion: &intschema.SyntheticTestTcpTestAssertionsConnectionAssertion{
					Operator: "EQUALS",
					Target:   "ESTABLISHED",
				},
			}, {
				NetworkHopsAssertion: &intschema.SyntheticTestTcpTestAssertionsNetworkHopsAssertion{
					Operator: "LESS_THAN",
					Target:   20,
				},
			}},
		}
		assert.Equal(t, in, syntheticTCPTestFromModel(syntheticTCPTestToModel(in)))
	})

	t.Run("tls", func(t *testing.T) {
		in := &intschema.SyntheticTestTlsTest{
			Host:                  "example.com",
			Port:                  443,
			ServerName:            "example.com",
			FailOnIncompleteChain: true,
			Assertions: []intschema.SyntheticTestTlsTestAssertions{{
				CertificateAssertion: &intschema.SyntheticTestTlsTestAssertionsCertificateAssertion{
					Operator:   "EXPIRES_IN_MORE_THAN_DAYS",
					TargetDays: 14,
				},
			}, {
				TlsVersionAssertion: &intschema.SyntheticTestTlsTestAssertionsTlsVersionAssertion{
					Bound:    "MIN",
					Operator: "GREATER_THAN_OR_EQUAL",
					Target:   "TLS_1_2",
				},
			}},
		}
		assert.Equal(t, in, syntheticTLSTestFromModel(syntheticTLSTestToModel(in)))
	})
}

// TestSyntheticTestSetUnknown covers the dry-run path that TestExamples cannot:
// startProvider disables dry run, so a nested tfid missing from
// syntheticTestDryRunSkipIDs would only panic against a real backend.
func TestSyntheticTestSetUnknown(t *testing.T) {
	raw := cty.ObjectVal(map[string]cty.Value{
		"collection_id": cty.StringVal("web"),
	})
	assert.NotPanics(t, func() {
		setUnknown(&intschema.SyntheticTest{}, setUnknownParams{
			rawConfig: raw,
			skipIDs:   set.New(syntheticTestDryRunSkipIDs...),
		})
	})
}
