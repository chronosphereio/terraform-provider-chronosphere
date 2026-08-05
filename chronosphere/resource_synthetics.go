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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

func resourceSyntheticTest() *schema.Resource {
	r := newGenericResource(
		"synthetic_test",
		syntheticTestConverter{},
		generatedUnstableSyntheticTest{},
	)
	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Description:   "A synthetic test that probes an endpoint from Chronosphere-operated locations and alerts on the result. " + unstableAPIWarning,
		Schema:        tfschema.SyntheticTest,
		CustomizeDiff: r.ValidateDryRunOptions(&SyntheticTestDryRunCount, ValidateDryRunOpts[*models.ConfigunstableSyntheticTest]{
			SetUnknownReferencesSkip: syntheticTestDryRunSkipIDs,
			ModifyAPIModel: func(m *models.ConfigunstableSyntheticTest) {
				// Skipped above, so a policy created by this same apply arrives
				// empty; the API requires one, so stand in a placeholder rather
				// than fail the plan on a reference that will resolve.
				if mc := m.MonitorConfig; mc != nil && mc.NotificationPolicySlug == "" {
					mc.NotificationPolicySlug = dryRunUnknownRef.Slug()
				}
			},
		}),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// Dry run does not support tfids nested in lists, and monitor_config is a
// list-encoded object.
var syntheticTestDryRunSkipIDs = []string{"monitor_config.[0].notification_policy_id"}

// SyntheticTestDryRunCount tracks how many times dry run is run during validation for testing.
var SyntheticTestDryRunCount atomic.Int64

type syntheticTestConverter struct{}

func (syntheticTestConverter) toModel(
	t *intschema.SyntheticTest,
) (*models.ConfigunstableSyntheticTest, error) {
	collSlug, collRef := unstableCollectionRefFromID(t.CollectionId.Slug())
	m := &models.ConfigunstableSyntheticTest{
		Name:           t.Name,
		Slug:           t.Slug,
		Description:    t.Description,
		CollectionSlug: collSlug,
		Collection:     collRef,
		Labels:         t.Labels,
		TestType:       enum.SyntheticTestType.V1(t.TestType),
		Status:         enum.SyntheticTestStatus.V1(t.Status),
		IntervalSecs:   int32(t.IntervalSecs),
		TimeoutSecs:    int32(t.TimeoutSecs),
		Locations:      sliceutil.Map(t.Locations, enum.SyntheticTestLocation.V1),
	}
	if c := t.RetryConfig; c != nil {
		m.RetryConfig = &models.SyntheticTestRetryConfig{
			MaxRetries:      int32(c.MaxRetries),
			RetryIntervalMs: int32(c.RetryIntervalMs),
		}
	}
	if c := t.MonitorConfig; c != nil {
		m.MonitorConfig = &models.SyntheticTestMonitorConfig{
			FailingDurationSecs:    int32(c.FailingDurationSecs),
			MinFailingLocations:    int32(c.MinFailingLocations),
			NotificationPolicySlug: c.NotificationPolicyId.Slug(),
			Annotations:            c.Annotations,
		}
	}
	m.HTTPTest = syntheticHTTPTestToModel(t.HttpTest)
	m.DNSTest = syntheticDNSTestToModel(t.DnsTest)
	m.TCPTest = syntheticTCPTestToModel(t.TcpTest)
	m.TLSTest = syntheticTLSTestToModel(t.TlsTest)
	return m, nil
}

func (syntheticTestConverter) fromModel(
	m *models.ConfigunstableSyntheticTest,
) (*intschema.SyntheticTest, error) {
	t := &intschema.SyntheticTest{
		Name:         m.Name,
		Slug:         m.Slug,
		Description:  m.Description,
		CollectionId: tfid.Slug(unstableCollectionIDFromRef(m.CollectionSlug, m.Collection)),
		Labels:       m.Labels,
		TestType:     enum.SyntheticTestType.Alias(m.TestType),
		Status:       enum.SyntheticTestStatus.Alias(m.Status),
		IntervalSecs: int64(m.IntervalSecs),
		TimeoutSecs:  int64(m.TimeoutSecs),
		Locations:    sliceutil.Map(m.Locations, enum.SyntheticTestLocation.Alias),
	}
	if c := m.RetryConfig; c != nil {
		t.RetryConfig = &intschema.SyntheticTestRetryConfig{
			MaxRetries:      int64(c.MaxRetries),
			RetryIntervalMs: int64(c.RetryIntervalMs),
		}
	}
	if c := m.MonitorConfig; c != nil {
		t.MonitorConfig = &intschema.SyntheticTestMonitorConfig{
			FailingDurationSecs:  int64(c.FailingDurationSecs),
			MinFailingLocations:  int64(c.MinFailingLocations),
			NotificationPolicyId: tfid.Slug(c.NotificationPolicySlug),
			Annotations:          c.Annotations,
		}
	}
	t.HttpTest = syntheticHTTPTestFromModel(m.HTTPTest)
	t.DnsTest = syntheticDNSTestFromModel(m.DNSTest)
	t.TcpTest = syntheticTCPTestFromModel(m.TCPTest)
	t.TlsTest = syntheticTLSTestFromModel(m.TLSTest)
	return t, nil
}

func (syntheticTestConverter) normalize(config, server *intschema.SyntheticTest) {
	if config.HttpTest == nil || server.HttpTest == nil {
		return
	}
	c, s := config.HttpTest.Authentication, server.HttpTest.Authentication
	if c == nil || s == nil {
		return
	}
	if c.BasicAuth != nil && s.BasicAuth != nil {
		s.BasicAuth.PasswordWoVersion = c.BasicAuth.PasswordWoVersion
	}
	if c.ApiTokenAuth != nil && s.ApiTokenAuth != nil {
		s.ApiTokenAuth.TokenWoVersion = c.ApiTokenAuth.TokenWoVersion
	}
	if c.ClientCertificate != nil && s.ClientCertificate != nil {
		s.ClientCertificate.PrivateKeyWoVersion = c.ClientCertificate.PrivateKeyWoVersion
	}
	if c.Oauth2ClientCredentials != nil && s.Oauth2ClientCredentials != nil {
		s.Oauth2ClientCredentials.ClientSecretWoVersion = c.Oauth2ClientCredentials.ClientSecretWoVersion
	}
	if c.Oauth2ResourceOwnerPassword != nil && s.Oauth2ResourceOwnerPassword != nil {
		s.Oauth2ResourceOwnerPassword.PasswordWoVersion = c.Oauth2ResourceOwnerPassword.PasswordWoVersion
		s.Oauth2ResourceOwnerPassword.ClientSecretWoVersion = c.Oauth2ResourceOwnerPassword.ClientSecretWoVersion
	}
}

func unstableCollectionRefFromID(id string) (slug string, ref *models.Configv1CollectionReference) {
	collType, slug, ok := CollectionTypeSlugFromID(id)
	if !ok {
		return id, nil
	}
	return "", &models.Configv1CollectionReference{
		Type: models.Configv1CollectionReferenceType(collType),
		Slug: slug,
	}
}

func unstableCollectionIDFromRef(slug string, ref *models.Configv1CollectionReference) string {
	if ref == nil {
		return slug
	}
	return CollectionTypeSlugToID(configmodels.Configv1CollectionReferenceType(ref.Type), ref.Slug)
}

func syntheticHTTPTestToModel(t *intschema.SyntheticTestHttpTest) *models.SyntheticTestHTTPTestConfig {
	if t == nil {
		return nil
	}
	return &models.SyntheticTestHTTPTestConfig{
		URL:                            t.Url,
		Method:                         enum.SyntheticHTTPMethod.V1(t.Method),
		HTTPVersion:                    enum.SyntheticHTTPVersion.V1(t.HttpVersion),
		ContentType:                    enum.SyntheticHTTPContentType.V1(t.ContentType),
		RequestBody:                    []byte(t.RequestBody),
		QueryParams:                    t.QueryParams,
		FollowRedirects:                t.FollowRedirects,
		MaxRedirects:                   int32(t.MaxRedirects),
		AllowInsecureTLS:               t.AllowInsecureTls,
		MaxResponseBodyBytes:           t.MaxResponseBodyBytes,
		DoNotSaveResponseBodyOnFailure: t.DoNotSaveResponseBodyOnFailure,
		Authentication:                 syntheticHTTPAuthToModel(t.Authentication),
		Headers: sliceutil.Map(t.Headers, func(h intschema.SyntheticTestHttpTestHeaders) *models.HTTPTestConfigHeader {
			return &models.HTTPTestConfigHeader{Name: h.Name, Value: h.Value}
		}),
		Cookies: sliceutil.Map(t.Cookies, func(c intschema.SyntheticTestHttpTestCookies) *models.HTTPTestConfigCookie {
			return &models.HTTPTestConfigCookie{Name: c.Name, Value: c.Value}
		}),
		Assertions: sliceutil.Map(t.Assertions, syntheticHTTPAssertionToModel),
	}
}

func syntheticHTTPTestFromModel(m *models.SyntheticTestHTTPTestConfig) *intschema.SyntheticTestHttpTest {
	if m == nil {
		return nil
	}
	t := &intschema.SyntheticTestHttpTest{
		Url:                            m.URL,
		Method:                         enum.SyntheticHTTPMethod.Alias(m.Method),
		HttpVersion:                    enum.SyntheticHTTPVersion.Alias(m.HTTPVersion),
		ContentType:                    enum.SyntheticHTTPContentType.Alias(m.ContentType),
		RequestBody:                    string(m.RequestBody),
		QueryParams:                    m.QueryParams,
		FollowRedirects:                m.FollowRedirects,
		MaxRedirects:                   int64(m.MaxRedirects),
		AllowInsecureTls:               m.AllowInsecureTLS,
		MaxResponseBodyBytes:           m.MaxResponseBodyBytes,
		DoNotSaveResponseBodyOnFailure: m.DoNotSaveResponseBodyOnFailure,
		Authentication:                 syntheticHTTPAuthFromModel(m.Authentication),
		Headers: sliceutil.Map(m.Headers, func(h *models.HTTPTestConfigHeader) intschema.SyntheticTestHttpTestHeaders {
			return intschema.SyntheticTestHttpTestHeaders{Name: h.Name, Value: h.Value}
		}),
		Cookies: sliceutil.Map(m.Cookies, func(c *models.HTTPTestConfigCookie) intschema.SyntheticTestHttpTestCookies {
			return intschema.SyntheticTestHttpTestCookies{Name: c.Name, Value: c.Value}
		}),
	}
	for _, a := range m.Assertions {
		if a == nil {
			continue
		}
		t.Assertions = append(t.Assertions, syntheticHTTPAssertionFromModel(a))
	}
	return t
}

func syntheticHTTPAuthToModel(a *intschema.SyntheticTestHttpTestAuthentication) *models.HTTPTestConfigHTTPAuth {
	if a == nil {
		return nil
	}
	m := &models.HTTPTestConfigHTTPAuth{}
	if v := a.BasicAuth; v != nil {
		m.BasicAuth = &models.ConfigunstableBasicAuth{
			Username: v.Username,
			Password: v.PasswordWo,
		}
	}
	if v := a.ApiTokenAuth; v != nil {
		m.APITokenAuth = &models.ConfigunstableAPITokenAuth{
			Key:   v.Key,
			Token: v.TokenWo,
		}
	}
	if v := a.ClientCertificate; v != nil {
		m.ClientCertificate = &models.ConfigunstableClientCertificate{
			Certificate: v.Certificate,
			PrivateKey:  v.PrivateKeyWo,
		}
	}
	if v := a.Oauth2ClientCredentials; v != nil {
		m.Oauth2ClientCredentials = &models.ConfigunstableOAuth2ClientCredentials{
			ClientID:     v.ClientId,
			ClientSecret: v.ClientSecretWo,
			Common:       syntheticOAuth2CommonToModel(v.Common),
		}
	}
	if v := a.Oauth2ResourceOwnerPassword; v != nil {
		m.Oauth2ResourceOwnerPassword = &models.ConfigunstableOAuth2ResourceOwnerPassword{
			Username:     v.Username,
			Password:     v.PasswordWo,
			ClientID:     v.ClientId,
			ClientSecret: v.ClientSecretWo,
			Common:       syntheticOAuth2CommonToModel(v.Common),
		}
	}
	return m
}

func syntheticHTTPAuthFromModel(m *models.HTTPTestConfigHTTPAuth) *intschema.SyntheticTestHttpTestAuthentication {
	if m == nil {
		return nil
	}
	a := &intschema.SyntheticTestHttpTestAuthentication{}
	if v := m.BasicAuth; v != nil {
		a.BasicAuth = &intschema.SyntheticTestHttpTestAuthenticationBasicAuth{
			Username: v.Username,
		}
	}
	if v := m.APITokenAuth; v != nil {
		a.ApiTokenAuth = &intschema.SyntheticTestHttpTestAuthenticationApiTokenAuth{
			Key: v.Key,
		}
	}
	if v := m.ClientCertificate; v != nil {
		a.ClientCertificate = &intschema.SyntheticTestHttpTestAuthenticationClientCertificate{
			Certificate: v.Certificate,
		}
	}
	if v := m.Oauth2ClientCredentials; v != nil {
		a.Oauth2ClientCredentials = &intschema.SyntheticTestHttpTestAuthenticationOauth2ClientCredentials{
			ClientId: v.ClientID,
			Common:   syntheticOAuth2CommonFromModel(v.Common),
		}
	}
	if v := m.Oauth2ResourceOwnerPassword; v != nil {
		a.Oauth2ResourceOwnerPassword = &intschema.SyntheticTestHttpTestAuthenticationOauth2ResourceOwnerPassword{
			Username: v.Username,
			ClientId: v.ClientID,
			Common:   syntheticOAuth2CommonFromModel(v.Common),
		}
	}
	return a
}

func syntheticOAuth2CommonToModel(
	c *intschema.SyntheticOAuth2Common,
) *models.ConfigunstableOAuth2Common {
	if c == nil {
		return nil
	}
	return &models.ConfigunstableOAuth2Common{
		AccessTokenURL:  c.AccessTokenUrl,
		TokenAuthMethod: enum.SyntheticOAuth2TokenAuthMethod.V1(c.TokenAuthMethod),
		Audience:        c.Audience,
		Resource:        c.Resource,
		Scopes:          c.Scopes,
	}
}

func syntheticOAuth2CommonFromModel(
	m *models.ConfigunstableOAuth2Common,
) *intschema.SyntheticOAuth2Common {
	if m == nil {
		return nil
	}
	return &intschema.SyntheticOAuth2Common{
		AccessTokenUrl:  m.AccessTokenURL,
		TokenAuthMethod: enum.SyntheticOAuth2TokenAuthMethod.Alias(m.TokenAuthMethod),
		Audience:        m.Audience,
		Resource:        m.Resource,
		Scopes:          m.Scopes,
	}
}

func syntheticHTTPAssertionToModel(
	a intschema.SyntheticTestHttpTestAssertions,
) *models.SyntheticTestHTTPTestConfigAssertion {
	m := &models.SyntheticTestHTTPTestConfigAssertion{
		ResponseTimeAssertion: syntheticResponseTimeAssertionToModel(a.ResponseTimeAssertion),
	}
	if v := a.StatusCodeAssertion; v != nil {
		m.StatusCodeAssertion = &models.SyntheticTestStatusCodeAssertion{
			Operator: enum.SyntheticStatusCodeAssertionOperator.V1(v.Operator),
			Target:   v.Target,
		}
	}
	if v := a.HeaderAssertion; v != nil {
		m.HeaderAssertion = &models.SyntheticTestHeaderAssertion{
			Name:     v.Name,
			Operator: enum.SyntheticHeaderAssertionOperator.V1(v.Operator),
			Target:   v.Target,
		}
	}
	if v := a.BodyAssertion; v != nil {
		m.BodyAssertion = &models.SyntheticTestBodyAssertion{
			Operator: enum.SyntheticBodyAssertionOperator.V1(v.Operator),
			Target:   v.Target,
		}
	}
	if v := a.BodyHashAssertion; v != nil {
		m.BodyHashAssertion = &models.SyntheticTestBodyHashAssertion{
			Algorithm: enum.SyntheticBodyHashAlgorithm.V1(v.Algorithm),
			Target:    v.Target,
		}
	}
	if v := a.BodyJsonPathAssertion; v != nil {
		m.BodyJSONPathAssertion = &models.SyntheticTestBodyJSONPathAssertion{
			JSONPath:  v.JsonPath,
			Operator:  enum.SyntheticBodyJSONPathAssertionOperator.V1(v.Operator),
			Target:    v.Target,
			MatchType: enum.SyntheticBodyJSONPathMatchType.V1(v.MatchType),
		}
	}
	if v := a.BodyJsonSchemaAssertion; v != nil {
		m.BodyJSONSchemaAssertion = &models.SyntheticTestBodyJSONSchemaAssertion{
			Schema: v.Schema,
			Draft:  enum.SyntheticBodyJSONSchemaDraft.V1(v.Draft),
		}
	}
	if v := a.BodyXpathAssertion; v != nil {
		m.BodyXpathAssertion = &models.SyntheticTestBodyXPathAssertion{
			Xpath:    v.Xpath,
			Operator: enum.SyntheticBodyXPathAssertionOperator.V1(v.Operator),
			Target:   v.Target,
		}
	}
	return m
}

func syntheticHTTPAssertionFromModel(
	m *models.SyntheticTestHTTPTestConfigAssertion,
) intschema.SyntheticTestHttpTestAssertions {
	a := intschema.SyntheticTestHttpTestAssertions{
		ResponseTimeAssertion: syntheticResponseTimeAssertionFromModel(m.ResponseTimeAssertion),
	}
	if v := m.StatusCodeAssertion; v != nil {
		a.StatusCodeAssertion = &intschema.SyntheticTestHttpTestAssertionsStatusCodeAssertion{
			Operator: enum.SyntheticStatusCodeAssertionOperator.Alias(v.Operator),
			Target:   v.Target,
		}
	}
	if v := m.HeaderAssertion; v != nil {
		a.HeaderAssertion = &intschema.SyntheticTestHttpTestAssertionsHeaderAssertion{
			Name:     v.Name,
			Operator: enum.SyntheticHeaderAssertionOperator.Alias(v.Operator),
			Target:   v.Target,
		}
	}
	if v := m.BodyAssertion; v != nil {
		a.BodyAssertion = &intschema.SyntheticTestHttpTestAssertionsBodyAssertion{
			Operator: enum.SyntheticBodyAssertionOperator.Alias(v.Operator),
			Target:   v.Target,
		}
	}
	if v := m.BodyHashAssertion; v != nil {
		a.BodyHashAssertion = &intschema.SyntheticTestHttpTestAssertionsBodyHashAssertion{
			Algorithm: enum.SyntheticBodyHashAlgorithm.Alias(v.Algorithm),
			Target:    v.Target,
		}
	}
	if v := m.BodyJSONPathAssertion; v != nil {
		a.BodyJsonPathAssertion = &intschema.SyntheticTestHttpTestAssertionsBodyJsonPathAssertion{
			JsonPath:  v.JSONPath,
			Operator:  enum.SyntheticBodyJSONPathAssertionOperator.Alias(v.Operator),
			Target:    v.Target,
			MatchType: enum.SyntheticBodyJSONPathMatchType.Alias(v.MatchType),
		}
	}
	if v := m.BodyJSONSchemaAssertion; v != nil {
		a.BodyJsonSchemaAssertion = &intschema.SyntheticTestHttpTestAssertionsBodyJsonSchemaAssertion{
			Schema: v.Schema,
			Draft:  enum.SyntheticBodyJSONSchemaDraft.Alias(v.Draft),
		}
	}
	if v := m.BodyXpathAssertion; v != nil {
		a.BodyXpathAssertion = &intschema.SyntheticTestHttpTestAssertionsBodyXpathAssertion{
			Xpath:    v.Xpath,
			Operator: enum.SyntheticBodyXPathAssertionOperator.Alias(v.Operator),
			Target:   v.Target,
		}
	}
	return a
}

func syntheticDNSTestToModel(t *intschema.SyntheticTestDnsTest) *models.SyntheticTestDNSTestConfig {
	if t == nil {
		return nil
	}
	m := &models.SyntheticTestDNSTestConfig{
		Domain:        t.Domain,
		DNSServer:     t.DnsServer,
		DNSServerPort: t.DnsServerPort,
	}
	for _, a := range t.Assertions {
		assertion := &models.SyntheticTestDNSTestConfigAssertion{
			ResponseTimeAssertion: syntheticDNSResponseTimeAssertionToModel(a.ResponseTimeAssertion),
		}
		if v := a.DnsRecordAssertion; v != nil {
			assertion.DNSRecordAssertion = &models.SyntheticTestDNSRecordAssertion{
				RecordType: enum.SyntheticDNSRecordType.V1(v.RecordType),
				Operator:   enum.SyntheticDNSRecordAssertionOperator.V1(v.Operator),
				Target:     v.Target,
				MatchScope: enum.SyntheticDNSRecordMatchScope.V1(v.MatchScope),
			}
		}
		m.Assertions = append(m.Assertions, assertion)
	}
	return m
}

func syntheticDNSTestFromModel(m *models.SyntheticTestDNSTestConfig) *intschema.SyntheticTestDnsTest {
	if m == nil {
		return nil
	}
	t := &intschema.SyntheticTestDnsTest{
		Domain:        m.Domain,
		DnsServer:     m.DNSServer,
		DnsServerPort: m.DNSServerPort,
	}
	for _, a := range m.Assertions {
		if a == nil {
			continue
		}
		assertion := intschema.SyntheticTestDnsTestAssertions{
			ResponseTimeAssertion: syntheticDNSResponseTimeAssertionFromModel(a.ResponseTimeAssertion),
		}
		if v := a.DNSRecordAssertion; v != nil {
			assertion.DnsRecordAssertion = &intschema.SyntheticTestDnsTestAssertionsDnsRecordAssertion{
				RecordType: enum.SyntheticDNSRecordType.Alias(v.RecordType),
				Operator:   enum.SyntheticDNSRecordAssertionOperator.Alias(v.Operator),
				Target:     v.Target,
				MatchScope: enum.SyntheticDNSRecordMatchScope.Alias(v.MatchScope),
			}
		}
		t.Assertions = append(t.Assertions, assertion)
	}
	return t
}

func syntheticTCPTestToModel(t *intschema.SyntheticTestTcpTest) *models.SyntheticTestTCPTestConfig {
	if t == nil {
		return nil
	}
	m := &models.SyntheticTestTCPTestConfig{
		Host: t.Host,
		Port: t.Port,
	}
	for _, a := range t.Assertions {
		assertion := &models.SyntheticTestTCPTestConfigAssertion{
			ResponseTimeAssertion: syntheticResponseTimeAssertionToModel(a.ResponseTimeAssertion),
		}
		if v := a.ConnectionAssertion; v != nil {
			assertion.ConnectionAssertion = &models.SyntheticTestConnectionAssertion{
				Operator: enum.SyntheticConnectionAssertionOperator.V1(v.Operator),
				Target:   enum.SyntheticConnectionStatus.V1(v.Target),
			}
		}
		if v := a.NetworkHopsAssertion; v != nil {
			assertion.NetworkHopsAssertion = &models.SyntheticTestNetworkHopsAssertion{
				Operator: enum.SyntheticNetworkHopsAssertionOperator.V1(v.Operator),
				Target:   int32(v.Target),
			}
		}
		m.Assertions = append(m.Assertions, assertion)
	}
	return m
}

func syntheticTCPTestFromModel(m *models.SyntheticTestTCPTestConfig) *intschema.SyntheticTestTcpTest {
	if m == nil {
		return nil
	}
	t := &intschema.SyntheticTestTcpTest{
		Host: m.Host,
		Port: m.Port,
	}
	for _, a := range m.Assertions {
		if a == nil {
			continue
		}
		assertion := intschema.SyntheticTestTcpTestAssertions{
			ResponseTimeAssertion: syntheticResponseTimeAssertionFromModel(a.ResponseTimeAssertion),
		}
		if v := a.ConnectionAssertion; v != nil {
			assertion.ConnectionAssertion = &intschema.SyntheticTestTcpTestAssertionsConnectionAssertion{
				Operator: enum.SyntheticConnectionAssertionOperator.Alias(v.Operator),
				Target:   enum.SyntheticConnectionStatus.Alias(v.Target),
			}
		}
		if v := a.NetworkHopsAssertion; v != nil {
			assertion.NetworkHopsAssertion = &intschema.SyntheticTestTcpTestAssertionsNetworkHopsAssertion{
				Operator: enum.SyntheticNetworkHopsAssertionOperator.Alias(v.Operator),
				Target:   int64(v.Target),
			}
		}
		t.Assertions = append(t.Assertions, assertion)
	}
	return t
}

func syntheticTLSTestToModel(t *intschema.SyntheticTestTlsTest) *models.SyntheticTestTLSTestConfig {
	if t == nil {
		return nil
	}
	m := &models.SyntheticTestTLSTestConfig{
		Host:                  t.Host,
		Port:                  t.Port,
		ServerName:            t.ServerName,
		AcceptSelfSigned:      t.AcceptSelfSigned,
		FailOnIncompleteChain: t.FailOnIncompleteChain,
	}
	for _, a := range t.Assertions {
		assertion := &models.SyntheticTestTLSTestConfigAssertion{
			ResponseTimeAssertion: syntheticResponseTimeAssertionToModel(a.ResponseTimeAssertion),
		}
		if v := a.CertificateAssertion; v != nil {
			assertion.CertificateAssertion = &models.SyntheticTestCertificateAssertion{
				Operator:   enum.SyntheticCertificateAssertionOperator.V1(v.Operator),
				TargetDays: int32(v.TargetDays),
			}
		}
		if v := a.CertPropertyAssertion; v != nil {
			assertion.CertPropertyAssertion = &models.SyntheticTestCertificatePropertyAssertion{
				Property: v.Property,
				Operator: enum.SyntheticCertificatePropertyAssertionOperator.V1(v.Operator),
				Target:   v.Target,
			}
		}
		if v := a.TlsVersionAssertion; v != nil {
			assertion.TLSVersionAssertion = &models.SyntheticTestTLSVersionAssertion{
				Bound:    enum.SyntheticTLSVersionBound.V1(v.Bound),
				Operator: enum.SyntheticTLSVersionAssertionOperator.V1(v.Operator),
				Target:   enum.SyntheticTLSVersion.V1(v.Target),
			}
		}
		m.Assertions = append(m.Assertions, assertion)
	}
	return m
}

func syntheticTLSTestFromModel(m *models.SyntheticTestTLSTestConfig) *intschema.SyntheticTestTlsTest {
	if m == nil {
		return nil
	}
	t := &intschema.SyntheticTestTlsTest{
		Host:                  m.Host,
		Port:                  m.Port,
		ServerName:            m.ServerName,
		AcceptSelfSigned:      m.AcceptSelfSigned,
		FailOnIncompleteChain: m.FailOnIncompleteChain,
	}
	for _, a := range m.Assertions {
		if a == nil {
			continue
		}
		assertion := intschema.SyntheticTestTlsTestAssertions{
			ResponseTimeAssertion: syntheticResponseTimeAssertionFromModel(a.ResponseTimeAssertion),
		}
		if v := a.CertificateAssertion; v != nil {
			assertion.CertificateAssertion = &intschema.SyntheticTestTlsTestAssertionsCertificateAssertion{
				Operator:   enum.SyntheticCertificateAssertionOperator.Alias(v.Operator),
				TargetDays: int64(v.TargetDays),
			}
		}
		if v := a.CertPropertyAssertion; v != nil {
			assertion.CertPropertyAssertion = &intschema.SyntheticTestTlsTestAssertionsCertPropertyAssertion{
				Property: v.Property,
				Operator: enum.SyntheticCertificatePropertyAssertionOperator.Alias(v.Operator),
				Target:   v.Target,
			}
		}
		if v := a.TLSVersionAssertion; v != nil {
			assertion.TlsVersionAssertion = &intschema.SyntheticTestTlsTestAssertionsTlsVersionAssertion{
				Bound:    enum.SyntheticTLSVersionBound.Alias(v.Bound),
				Operator: enum.SyntheticTLSVersionAssertionOperator.Alias(v.Operator),
				Target:   enum.SyntheticTLSVersion.Alias(v.Target),
			}
		}
		t.Assertions = append(t.Assertions, assertion)
	}
	return t
}

func syntheticResponseTimeAssertionToModel(
	a *intschema.SyntheticResponseTimeAssertion,
) *models.SyntheticTestResponseTimeAssertion {
	if a == nil {
		return nil
	}
	return &models.SyntheticTestResponseTimeAssertion{
		Operator: enum.SyntheticResponseTimeAssertionOperator.V1(a.Operator),
		TargetMs: int32(a.TargetMs),
		Scope:    enum.SyntheticResponseTimeScope.V1(a.Scope),
	}
}

func syntheticResponseTimeAssertionFromModel(
	m *models.SyntheticTestResponseTimeAssertion,
) *intschema.SyntheticResponseTimeAssertion {
	if m == nil {
		return nil
	}
	return &intschema.SyntheticResponseTimeAssertion{
		Operator: enum.SyntheticResponseTimeAssertionOperator.Alias(m.Operator),
		TargetMs: int64(m.TargetMs),
		Scope:    enum.SyntheticResponseTimeScope.Alias(m.Scope),
	}
}

func syntheticDNSResponseTimeAssertionToModel(
	a *intschema.SyntheticTestDnsTestAssertionsResponseTimeAssertion,
) *models.SyntheticTestResponseTimeAssertion {
	if a == nil {
		return nil
	}
	return &models.SyntheticTestResponseTimeAssertion{
		Operator: enum.SyntheticResponseTimeAssertionOperator.V1(a.Operator),
		TargetMs: int32(a.TargetMs),
	}
}

func syntheticDNSResponseTimeAssertionFromModel(
	m *models.SyntheticTestResponseTimeAssertion,
) *intschema.SyntheticTestDnsTestAssertionsResponseTimeAssertion {
	if m == nil {
		return nil
	}
	return &intschema.SyntheticTestDnsTestAssertionsResponseTimeAssertion{
		Operator: enum.SyntheticResponseTimeAssertionOperator.Alias(m.Operator),
		TargetMs: int64(m.TargetMs),
	}
}
