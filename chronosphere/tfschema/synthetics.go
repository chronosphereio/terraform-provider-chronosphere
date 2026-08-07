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

package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// SyntheticResponseTimeAssertionSchema asserts on the response time of an
// HTTP, TCP, or TLS test.
var SyntheticResponseTimeAssertionSchema = syntheticResponseTimeAssertion(true /* scoped */)

// SyntheticDNSResponseTimeAssertionSchema is the DNS variant, which takes no
// scope: a DNS test measures resolution itself, so there is no non-DNS portion
// to exclude and the API rejects the field.
var SyntheticDNSResponseTimeAssertionSchema = syntheticResponseTimeAssertion(false /* scoped */)

func syntheticResponseTimeAssertion(scoped bool) *schema.Schema {
	s := map[string]*schema.Schema{
		"operator": Enum{
			Value:       enum.SyntheticResponseTimeAssertionOperator.ToStrings(),
			Required:    true,
			Description: "Comparison applied to the measured response time. Only `LESS_THAN` is currently supported.",
		}.Schema(),
		"target_ms": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Response time threshold in milliseconds.",
		},
	}
	if scoped {
		s["scope"] = Enum{
			Value:       enum.SyntheticResponseTimeScope.ToStrings(),
			Required:    true,
			Description: "Whether DNS resolution time counts towards the measured response time: `INCLUDING_DNS` or `WITHOUT_DNS`.",
		}.Schema()
	}
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts on the total time taken by the check.",
		Elem:        &schema.Resource{Schema: s},
	}
}

var syntheticTestOneOfConfigFields = []string{"http_test", "dns_test", "tcp_test", "tls_test"}

func syntheticProtocolTest(description string, s map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeList,
		Optional:     true,
		MaxItems:     1,
		ExactlyOneOf: syntheticTestOneOfConfigFields,
		Description:  description,
		Elem:         &schema.Resource{Schema: s},
	}
}

var SyntheticTest = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the synthetic test. Can be changed after creation.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the synthetic test. Generated from `name` if omitted. Immutable after creation.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Free-form description of the synthetic test.",
	},
	"collection_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "ID of the collection that owns this synthetic test.",
	},
	"labels": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Key/value labels attached to the synthetic test. Visible in notifications and usable for routing overrides in notification policies.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"test_type": Enum{
		Value:       enum.SyntheticTestType.ToStrings(),
		Required:    true,
		Description: "Protocol exercised by the test: `HTTP`, `DNS`, `TCP`, or `TLS`. Must match the protocol block set below.",
	}.Schema(),
	"status": Enum{
		Value:       enum.SyntheticTestStatus.ToStrings(),
		Optional:    true,
		Description: "Whether the test executes on its schedule (`ENABLED`) or is suspended (`PAUSED`).",
	}.Schema(),
	"locations": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		Description: "Probe locations the test runs from. Each location executes the test independently.",
		Elem: &schema.Schema{
			Type:             schema.TypeString,
			ValidateDiagFunc: enum.SyntheticTestLocation.ToStrings().Validate,
		},
	},
	"interval_secs": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "How often the test runs, in seconds.",
	},
	"timeout_secs": {
		Type:     schema.TypeInt,
		Optional: true,
		// Computed because the server persists its default rather than
		// leaving the field unset, so an omitted value must not plan as 0.
		Computed:    true,
		Description: "Per-run timeout in seconds, between 1 and 60. Defaults to 60.",
	},
	"retry_config": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Retry behaviour applied to a failing run before it is reported as a failure.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"max_retries": {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: "Number of retries attempted after the initial failure.",
				},
				"retry_interval_ms": {
					Type:     schema.TypeInt,
					Optional: true,
					// Computed: the server persists 300 when max_retries is set.
					Computed:    true,
					Description: "Delay between retries, in milliseconds. Defaults to 300.",
				},
			},
		},
	},
	"monitor_config": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Alerting behaviour derived from the test's pass/fail results.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"failing_duration_secs": {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: "How long the test must fail continuously before a signal fires, in seconds.",
				},
				"min_failing_locations": {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: "Number of locations that must fail concurrently before a signal fires.",
				},
				"notification_policy_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "ID of the notification policy that routes signals from this test. If omitted, the owning collection's default policy applies.",
				},
				"annotations": {
					Type:        schema.TypeMap,
					Optional:    true,
					Description: "Free-form key/value pairs included in notifications generated by this test.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	"http_test": syntheticProtocolTest("HTTP request configuration. Set when `test_type` is `HTTP`.", syntheticHTTPTest),
	"dns_test":  syntheticProtocolTest("DNS resolution configuration. Set when `test_type` is `DNS`.", syntheticDNSTest),
	"tcp_test":  syntheticProtocolTest("TCP connection configuration. Set when `test_type` is `TCP`.", syntheticTCPTest),
	"tls_test":  syntheticProtocolTest("TLS handshake configuration. Set when `test_type` is `TLS`.", syntheticTLSTest),
}

var syntheticHTTPTest = map[string]*schema.Schema{
	"url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "URL requested by the test.",
	},
	"method": Enum{
		Value:       enum.SyntheticHTTPMethod.ToStrings(),
		Required:    true,
		Description: "HTTP method used for the request: `GET` or `POST`.",
	}.Schema(),
	"http_version": Enum{
		Value:       enum.SyntheticHTTPVersion.ToStrings(),
		Optional:    true,
		Description: "HTTP version negotiated for the request: `HTTP_1_1` or `HTTP_2`.",
	}.Schema(),
	"content_type": Enum{
		Value:       enum.SyntheticHTTPContentType.ToStrings(),
		Optional:    true,
		Description: "Content type sent with `request_body`, e.g. `APPLICATION_JSON` or `TEXT_PLAIN`.",
	}.Schema(),
	"request_body": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Request body sent with the request. Encoded on the wire; supply it as plain text.",
	},
	"query_params": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Query parameters appended to `url`.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"headers": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Headers sent with the request.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Header name.",
				},
				"value": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Header value. May contain a `{{VAR_NAME}}` reference to a synthetic global variable.",
				},
			},
		},
	},
	"cookies": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Cookies sent with the request.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Cookie name.",
				},
				"value": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Cookie value. May contain a `{{VAR_NAME}}` reference to a synthetic global variable.",
				},
			},
		},
	},
	"follow_redirects": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Whether the test follows HTTP redirects.",
	},
	"max_redirects": {
		Type:     schema.TypeInt,
		Optional: true,
		// Computed: the server persists 10 when follow_redirects is set, and
		// forces 0 when it is not.
		Computed:    true,
		Description: "Maximum number of redirects followed when `follow_redirects` is set. Defaults to 10.",
	},
	"allow_insecure_tls": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Whether to accept TLS certificates that fail verification.",
	},
	"max_response_body_bytes": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Maximum number of response body bytes read and retained.",
	},
	"do_not_save_response_body_on_failure": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Whether to omit the response body from stored results when a run fails. Set this when the body may contain sensitive data.",
	},
	"authentication": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Credentials presented with the request. Set exactly one of the nested blocks.",
		Elem: &schema.Resource{
			Schema: syntheticHTTPAuth,
		},
	},
	"assertions": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Conditions evaluated against the response. Every assertion must pass for the run to pass. Set exactly one nested block per entry.",
		Elem: &schema.Resource{
			Schema: syntheticHTTPAssertion,
		},
	},
}

var syntheticHTTPAuth = map[string]*schema.Schema{
	"basic_auth": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "HTTP basic authentication.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Basic auth username.",
				},
				"password_wo": WriteOnlySecret{
					Name:        "password",
					Required:    true,
					Description: "Basic auth password. May contain a `{{VAR_NAME}}` reference to a synthetic global variable.",
				}.Schema(),
				"password_wo_version": SecretVersionSchema("password"),
			},
		},
	},
	"api_token_auth": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Static API token sent as a header.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Header name carrying the token, e.g. `Authorization` or `X-API-Key`.",
				},
				"token_wo": WriteOnlySecret{
					Name:        "token",
					Required:    true,
					Description: "Full header value, e.g. `Bearer <secret>`. Treated as a secret in its entirety.",
				}.Schema(),
				"token_wo_version": SecretVersionSchema("token"),
			},
		},
	},
	"client_certificate": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Client certificate presented during the TLS handshake (mTLS).",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"certificate": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "PEM leaf certificate plus any intermediates.",
				},
				"private_key_wo": WriteOnlySecret{
					Name:        "private_key",
					Required:    true,
					Description: "Matching unencrypted PKCS#8/PKCS#1/SEC1 PEM key.",
				}.Schema(),
				"private_key_wo_version": SecretVersionSchema("private_key"),
			},
		},
	},
	"oauth2_client_credentials": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "OAuth 2.0 client-credentials grant.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"client_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "OAuth 2.0 client identifier.",
				},
				"client_secret_wo": WriteOnlySecret{
					Name:        "client_secret",
					Required:    true,
					Description: "OAuth 2.0 client secret.",
				}.Schema(),
				"client_secret_wo_version": SecretVersionSchema("client_secret"),
				"common":                   SyntheticOAuth2CommonSchema,
			},
		},
	},
	"oauth2_resource_owner_password": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "OAuth 2.0 resource-owner password-credentials grant.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Resource owner username.",
				},
				"password_wo": WriteOnlySecret{
					Name:        "password",
					Required:    true,
					Description: "Resource owner password.",
				}.Schema(),
				"password_wo_version": SecretVersionSchema("password"),
				"client_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "OAuth 2.0 client identifier. Set when the token endpoint also authenticates the client.",
				},
				"client_secret_wo": WriteOnlySecret{
					Name:        "client_secret",
					Description: "OAuth 2.0 client secret. Set when the token endpoint also authenticates the client.",
				}.Schema(),
				"client_secret_wo_version": SecretVersionSchema("client_secret"),
				"common":                   SyntheticOAuth2CommonSchema,
			},
		},
	},
}

var SyntheticOAuth2CommonSchema = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Token endpoint parameters.",
	Elem: &schema.Resource{
		Schema: syntheticOAuth2Common,
	},
}

var syntheticOAuth2Common = map[string]*schema.Schema{
	"access_token_url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Token endpoint the grant is exchanged against.",
	},
	"token_auth_method": Enum{
		Value:       enum.SyntheticOAuth2TokenAuthMethod.ToStrings(),
		Optional:    true,
		Description: "How client credentials are sent to the token endpoint: `BASIC_AUTH_HEADER` or `REQUEST_BODY`. Defaults to `BASIC_AUTH_HEADER`.",
	}.Schema(),
	"audience": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "`audience` parameter sent to the token endpoint.",
	},
	"resource": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "`resource` parameter sent to the token endpoint.",
	},
	"scopes": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Scopes requested from the token endpoint.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

var syntheticHTTPAssertion = map[string]*schema.Schema{
	"status_code_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts on the HTTP response status code.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"operator": Enum{
					Value:       enum.SyntheticStatusCodeAssertionOperator.ToStrings(),
					Required:    true,
					Description: "Comparison applied to the status code.",
				}.Schema(),
				"target": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Status code, or regular expression when `operator` is a regex comparison.",
				},
			},
		},
	},
	"header_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts on a response header.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Header name to evaluate.",
				},
				"operator": Enum{
					Value:       enum.SyntheticHeaderAssertionOperator.ToStrings(),
					Required:    true,
					Description: "Comparison applied to the header value.",
				}.Schema(),
				"target": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Value compared against the header. Omit when `operator` is `NOT_EXISTS`.",
				},
			},
		},
	},
	"body_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts on the raw response body.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"operator": Enum{
					Value:       enum.SyntheticBodyAssertionOperator.ToStrings(),
					Required:    true,
					Description: "Comparison applied to the response body.",
				}.Schema(),
				"target": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Value compared against the body, or regular expression when `operator` is a regex comparison.",
				},
			},
		},
	},
	"body_hash_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts that the response body hashes to a known digest.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"algorithm": Enum{
					Value:       enum.SyntheticBodyHashAlgorithm.ToStrings(),
					Required:    true,
					Description: "Digest algorithm: `MD5`, `SHA1`, or `SHA256`.",
				}.Schema(),
				"target": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Expected digest of the response body.",
				},
			},
		},
	},
	"body_json_path_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts on a JSONPath expression evaluated against the response body.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"json_path": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "JSONPath expression selecting the value to evaluate.",
				},
				"operator": Enum{
					Value:       enum.SyntheticBodyJSONPathAssertionOperator.ToStrings(),
					Required:    true,
					Description: "Comparison applied to the selected value.",
				}.Schema(),
				"target": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Value compared against the selection. Omit when `operator` is `IS_UNDEFINED`.",
				},
				"match_type": Enum{
					Value:       enum.SyntheticBodyJSONPathMatchType.ToStrings(),
					Optional:    true,
					Description: "How a multi-element selection is evaluated: `FIRST_ELEMENT`, `EVERY_ELEMENT`, `AT_LEAST_ONE_ELEMENT`, or `SERIALIZATION`.",
				}.Schema(),
			},
		},
	},
	"body_json_schema_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts that the response body validates against a JSON schema.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"schema": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "JSON schema document the response body is validated against.",
				},
				"draft": Enum{
					Value:       enum.SyntheticBodyJSONSchemaDraft.ToStrings(),
					Optional:    true,
					Description: "JSON schema draft the document is written against: `DRAFT_06` or `DRAFT_07`.",
				}.Schema(),
			},
		},
	},
	"body_xpath_assertion": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Asserts on an XPath expression evaluated against the response body.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"xpath": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "XPath expression selecting the value to evaluate.",
				},
				"operator": Enum{
					Value:       enum.SyntheticBodyXPathAssertionOperator.ToStrings(),
					Required:    true,
					Description: "Comparison applied to the selected value.",
				}.Schema(),
				"target": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Value compared against the selection. Omit when `operator` is `IS_UNDEFINED`.",
				},
			},
		},
	},
	"response_time_assertion": SyntheticResponseTimeAssertionSchema,
}

var syntheticDNSTest = map[string]*schema.Schema{
	"domain": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Domain name resolved by the test.",
	},
	"dns_server": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "DNS server queried. Defaults to the probe's resolver.",
	},
	"dns_server_port": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Port on `dns_server` to query.",
	},
	"assertions": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Conditions evaluated against the DNS response. Every assertion must pass for the run to pass. Set exactly one nested block per entry.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dns_record_assertion": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Asserts on the returned DNS records.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"record_type": Enum{
								Value:       enum.SyntheticDNSRecordType.ToStrings(),
								Required:    true,
								Description: "Record type evaluated, e.g. `A`, `AAAA`, `CNAME`, `MX`.",
							}.Schema(),
							"operator": Enum{
								Value:       enum.SyntheticDNSRecordAssertionOperator.ToStrings(),
								Required:    true,
								Description: "Comparison applied to the record values.",
							}.Schema(),
							"target": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Value compared against the records, or regular expression when `operator` is a regex comparison.",
							},
							"match_scope": Enum{
								Value:       enum.SyntheticDNSRecordMatchScope.ToStrings(),
								Required:    true,
								Description: "Whether every returned record must match (`EVERY_AVAILABLE`) or at least one (`AT_LEAST_ONE`).",
							}.Schema(),
						},
					},
				},
				"response_time_assertion": SyntheticDNSResponseTimeAssertionSchema,
			},
		},
	},
}

var syntheticTCPTest = map[string]*schema.Schema{
	"host": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Host the test connects to.",
	},
	"port": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Port the test connects to.",
	},
	"assertions": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Conditions evaluated against the connection attempt. Every assertion must pass for the run to pass. Set exactly one nested block per entry.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"connection_assertion": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Asserts on the outcome of the connection attempt.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"operator": Enum{
								Value:       enum.SyntheticConnectionAssertionOperator.ToStrings(),
								Required:    true,
								Description: "Comparison applied to the connection outcome. Only `EQUALS` is currently supported.",
							}.Schema(),
							"target": Enum{
								Value:       enum.SyntheticConnectionStatus.ToStrings(),
								Required:    true,
								Description: "Expected outcome: `ESTABLISHED`, `REFUSED`, or `TIMEOUT`.",
							}.Schema(),
						},
					},
				},
				"network_hops_assertion": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Asserts on the number of network hops traversed.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"operator": Enum{
								Value:       enum.SyntheticNetworkHopsAssertionOperator.ToStrings(),
								Required:    true,
								Description: "Comparison applied to the hop count.",
							}.Schema(),
							"target": {
								Type:        schema.TypeInt,
								Required:    true,
								Description: "Hop count compared against.",
							},
						},
					},
				},
				"response_time_assertion": SyntheticResponseTimeAssertionSchema,
			},
		},
	},
}

var syntheticTLSTest = map[string]*schema.Schema{
	"host": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Host the test performs the TLS handshake against.",
	},
	"port": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Port the test performs the TLS handshake against.",
	},
	"server_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "SNI server name sent in the handshake. Defaults to `host`.",
	},
	"accept_self_signed": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Whether a self-signed certificate is accepted.",
	},
	"fail_on_incomplete_chain": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Whether a certificate chain missing intermediates fails the run.",
	},
	"assertions": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Conditions evaluated against the handshake. Every assertion must pass for the run to pass. Set exactly one nested block per entry.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"certificate_assertion": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Asserts on the certificate's remaining validity.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"operator": Enum{
								Value:       enum.SyntheticCertificateAssertionOperator.ToStrings(),
								Required:    true,
								Description: "Comparison applied to the expiry: `EXPIRES_IN_LESS_THAN_DAYS` or `EXPIRES_IN_MORE_THAN_DAYS`.",
							}.Schema(),
							"target_days": {
								Type:        schema.TypeInt,
								Required:    true,
								Description: "Number of days compared against.",
							},
						},
					},
				},
				"cert_property_assertion": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Asserts on a named property of the presented certificate.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"property": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Certificate property evaluated, e.g. `issuer` or `subject`.",
							},
							"operator": Enum{
								Value:       enum.SyntheticCertificatePropertyAssertionOperator.ToStrings(),
								Required:    true,
								Description: "Comparison applied to the property value.",
							}.Schema(),
							"target": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Value compared against the property, or regular expression when `operator` is a regex comparison.",
							},
						},
					},
				},
				"tls_version_assertion": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Asserts on the negotiated TLS version.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"bound": Enum{
								Value:       enum.SyntheticTLSVersionBound.ToStrings(),
								Required:    true,
								Description: "Whether `target` is the minimum (`MIN`) or maximum (`MAX`) acceptable version.",
							}.Schema(),
							"operator": Enum{
								Value:       enum.SyntheticTLSVersionAssertionOperator.ToStrings(),
								Required:    true,
								Description: "Comparison applied to the negotiated version.",
							}.Schema(),
							"target": Enum{
								Value:       enum.SyntheticTLSVersion.ToStrings(),
								Required:    true,
								Description: "TLS version compared against: `TLS_1_0` through `TLS_1_3`.",
							}.Schema(),
						},
					},
				},
				"response_time_assertion": SyntheticResponseTimeAssertionSchema,
			},
		},
	},
}
