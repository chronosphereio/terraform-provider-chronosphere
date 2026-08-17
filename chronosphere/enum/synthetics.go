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

package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

var SyntheticTestType = newEnum("SyntheticTestType", []value[configv1.SyntheticTestTestType]{
	{v1: configv1.SyntheticTestTestTypeHTTP, alias: "HTTP"},
	{v1: configv1.SyntheticTestTestTypeDNS, alias: "DNS"},
	{v1: configv1.SyntheticTestTestTypeTCP, alias: "TCP"},
	{v1: configv1.SyntheticTestTestTypeTLS, alias: "TLS"},
})

var SyntheticTestStatus = newEnum("SyntheticTestStatus", []value[configv1.SyntheticTestTestStatus]{
	{v1: configv1.SyntheticTestTestStatusENABLED, alias: "ENABLED"},
	{v1: configv1.SyntheticTestTestStatusPAUSED, alias: "PAUSED"},
})

var SyntheticTestLocation = newEnum("SyntheticTestLocation", []value[configv1.SyntheticTestTestLocation]{
	{v1: configv1.SyntheticTestTestLocationGCPUSOREGON, alias: "GCP_US_OREGON"},
	{v1: configv1.SyntheticTestTestLocationGCPUSVIRGINIA, alias: "GCP_US_VIRGINIA"},
})

var SyntheticHTTPMethod = newEnum("SyntheticHTTPMethod", []value[configv1.HTTPTestConfigHTTPMethod]{
	{v1: configv1.HTTPTestConfigHTTPMethodGET, alias: "GET"},
	{v1: configv1.HTTPTestConfigHTTPMethodPOST, alias: "POST"},
})

// SyntheticHTTPVersion aliases away the wire values' redundant HTTP_VERSION_ prefix.
var SyntheticHTTPVersion = newEnum("SyntheticHTTPVersion", []value[configv1.HTTPTestConfigHTTPVersion]{
	{v1: configv1.HTTPTestConfigHTTPVersionHTTPVERSIONHTTP11, alias: "HTTP_1_1"},
	{v1: configv1.HTTPTestConfigHTTPVersionHTTPVERSIONHTTP2, alias: "HTTP_2"},
})

// SyntheticHTTPContentType aliases away the wire values' redundant CONTENT_TYPE_ prefix.
var SyntheticHTTPContentType = newEnum("SyntheticHTTPContentType", []value[configv1.HTTPTestConfigContentType]{
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONJSON, alias: "APPLICATION_JSON"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONOCTETSTREAM, alias: "APPLICATION_OCTET_STREAM"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONXWWWFORMURLENCODED, alias: "APPLICATION_X_WWW_FORM_URLENCODED"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPEGRAPHQL, alias: "GRAPHQL"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPEMULTIPARTFORMDATA, alias: "MULTIPART_FORM_DATA"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPETEXTHTML, alias: "TEXT_HTML"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPETEXTPLAIN, alias: "TEXT_PLAIN"},
	{v1: configv1.HTTPTestConfigContentTypeCONTENTTYPETEXTXML, alias: "TEXT_XML"},
})

// SyntheticOAuth2TokenAuthMethod aliases away the wire values' redundant OTAM_ prefix.
var SyntheticOAuth2TokenAuthMethod = newEnum("SyntheticOAuth2TokenAuthMethod", []value[configv1.Configv1OAuth2TokenAuthMethod]{
	{v1: configv1.Configv1OAuth2TokenAuthMethodOTAMBASICAUTHHEADER, alias: "BASIC_AUTH_HEADER"},
	{v1: configv1.Configv1OAuth2TokenAuthMethodOTAMREQUESTBODY, alias: "REQUEST_BODY"},
})

var SyntheticBodyAssertionOperator = newEnum("SyntheticBodyAssertionOperator", []value[configv1.SyntheticTestBodyAssertionOperator]{
	{v1: configv1.SyntheticTestBodyAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestBodyAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configv1.SyntheticTestBodyAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configv1.SyntheticTestBodyAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configv1.SyntheticTestBodyAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestBodyAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticBodyHashAlgorithm = newEnum("SyntheticBodyHashAlgorithm", []value[configv1.BodyHashAssertionAlgorithm]{
	{v1: configv1.BodyHashAssertionAlgorithmMD5, alias: "MD5"},
	{v1: configv1.BodyHashAssertionAlgorithmSHA1, alias: "SHA1"},
	{v1: configv1.BodyHashAssertionAlgorithmSHA256, alias: "SHA256"},
})

var SyntheticBodyJSONPathAssertionOperator = newEnum("SyntheticBodyJSONPathAssertionOperator", []value[configv1.SyntheticTestBodyJSONPathAssertionOperator]{
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configv1.SyntheticTestBodyJSONPathAssertionOperatorISUNDEFINED, alias: "IS_UNDEFINED"},
})

var SyntheticBodyJSONPathMatchType = newEnum("SyntheticBodyJSONPathMatchType", []value[configv1.BodyJSONPathAssertionMatchType]{
	{v1: configv1.BodyJSONPathAssertionMatchTypeFIRSTELEMENT, alias: "FIRST_ELEMENT"},
	{v1: configv1.BodyJSONPathAssertionMatchTypeEVERYELEMENT, alias: "EVERY_ELEMENT"},
	{v1: configv1.BodyJSONPathAssertionMatchTypeATLEASTONEELEMENT, alias: "AT_LEAST_ONE_ELEMENT"},
	{v1: configv1.BodyJSONPathAssertionMatchTypeSERIALIZATION, alias: "SERIALIZATION"},
})

var SyntheticBodyJSONSchemaDraft = newEnum("SyntheticBodyJSONSchemaDraft", []value[configv1.BodyJSONSchemaAssertionSchemaDraft]{
	{v1: configv1.BodyJSONSchemaAssertionSchemaDraftDRAFT06, alias: "DRAFT_06"},
	{v1: configv1.BodyJSONSchemaAssertionSchemaDraftDRAFT07, alias: "DRAFT_07"},
})

var SyntheticBodyXPathAssertionOperator = newEnum("SyntheticBodyXPathAssertionOperator", []value[configv1.SyntheticTestBodyXPathAssertionOperator]{
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configv1.SyntheticTestBodyXPathAssertionOperatorISUNDEFINED, alias: "IS_UNDEFINED"},
})

var SyntheticHeaderAssertionOperator = newEnum("SyntheticHeaderAssertionOperator", []value[configv1.SyntheticTestHeaderAssertionOperator]{
	{v1: configv1.SyntheticTestHeaderAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorGREATERTHANOREQUAL, alias: "GREATER_THAN_OR_EQUAL"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorLESSTHANOREQUAL, alias: "LESS_THAN_OR_EQUAL"},
	{v1: configv1.SyntheticTestHeaderAssertionOperatorNOTEXISTS, alias: "NOT_EXISTS"},
})

var SyntheticResponseTimeAssertionOperator = newEnum("SyntheticResponseTimeAssertionOperator", []value[configv1.SyntheticTestResponseTimeAssertionOperator]{
	{v1: configv1.SyntheticTestResponseTimeAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
})

var SyntheticResponseTimeScope = newEnum("SyntheticResponseTimeScope", []value[configv1.ResponseTimeAssertionScope]{
	{v1: configv1.ResponseTimeAssertionScopeINCLUDINGDNS, alias: "INCLUDING_DNS"},
	{v1: configv1.ResponseTimeAssertionScopeWITHOUTDNS, alias: "WITHOUT_DNS"},
})

var SyntheticStatusCodeAssertionOperator = newEnum("SyntheticStatusCodeAssertionOperator", []value[configv1.SyntheticTestStatusCodeAssertionOperator]{
	{v1: configv1.SyntheticTestStatusCodeAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestStatusCodeAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configv1.SyntheticTestStatusCodeAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestStatusCodeAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticDNSRecordAssertionOperator = newEnum("SyntheticDNSRecordAssertionOperator", []value[configv1.SyntheticTestDNSRecordAssertionOperator]{
	{v1: configv1.SyntheticTestDNSRecordAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestDNSRecordAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configv1.SyntheticTestDNSRecordAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestDNSRecordAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticDNSRecordMatchScope = newEnum("SyntheticDNSRecordMatchScope", []value[configv1.DNSRecordAssertionMatchScope]{
	{v1: configv1.DNSRecordAssertionMatchScopeATLEASTONE, alias: "AT_LEAST_ONE"},
	{v1: configv1.DNSRecordAssertionMatchScopeEVERYAVAILABLE, alias: "EVERY_AVAILABLE"},
})

var SyntheticDNSRecordType = newEnum("SyntheticDNSRecordType", []value[configv1.DNSRecordAssertionRecordType]{
	{v1: configv1.DNSRecordAssertionRecordTypeA, alias: "A"},
	{v1: configv1.DNSRecordAssertionRecordTypeAAAA, alias: "AAAA"},
	{v1: configv1.DNSRecordAssertionRecordTypeCNAME, alias: "CNAME"},
	{v1: configv1.DNSRecordAssertionRecordTypeMX, alias: "MX"},
	{v1: configv1.DNSRecordAssertionRecordTypeNS, alias: "NS"},
	{v1: configv1.DNSRecordAssertionRecordTypeSOA, alias: "SOA"},
	{v1: configv1.DNSRecordAssertionRecordTypeSRV, alias: "SRV"},
	{v1: configv1.DNSRecordAssertionRecordTypeTXT, alias: "TXT"},
})

var SyntheticConnectionAssertionOperator = newEnum("SyntheticConnectionAssertionOperator", []value[configv1.SyntheticTestConnectionAssertionOperator]{
	{v1: configv1.SyntheticTestConnectionAssertionOperatorEQUALS, alias: "EQUALS"},
})

var SyntheticConnectionStatus = newEnum("SyntheticConnectionStatus", []value[configv1.ConnectionAssertionConnectionStatus]{
	{v1: configv1.ConnectionAssertionConnectionStatusESTABLISHED, alias: "ESTABLISHED"},
	{v1: configv1.ConnectionAssertionConnectionStatusREFUSED, alias: "REFUSED"},
	{v1: configv1.ConnectionAssertionConnectionStatusTIMEOUT, alias: "TIMEOUT"},
})

var SyntheticNetworkHopsAssertionOperator = newEnum("SyntheticNetworkHopsAssertionOperator", []value[configv1.SyntheticTestNetworkHopsAssertionOperator]{
	{v1: configv1.SyntheticTestNetworkHopsAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestNetworkHopsAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configv1.SyntheticTestNetworkHopsAssertionOperatorGREATERTHANOREQUAL, alias: "GREATER_THAN_OR_EQUAL"},
	{v1: configv1.SyntheticTestNetworkHopsAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configv1.SyntheticTestNetworkHopsAssertionOperatorLESSTHANOREQUAL, alias: "LESS_THAN_OR_EQUAL"},
})

var SyntheticCertificateAssertionOperator = newEnum("SyntheticCertificateAssertionOperator", []value[configv1.SyntheticTestCertificateAssertionOperator]{
	{v1: configv1.SyntheticTestCertificateAssertionOperatorEXPIRESINLESSTHANDAYS, alias: "EXPIRES_IN_LESS_THAN_DAYS"},
	{v1: configv1.SyntheticTestCertificateAssertionOperatorEXPIRESINMORETHANDAYS, alias: "EXPIRES_IN_MORE_THAN_DAYS"},
})

var SyntheticCertificatePropertyAssertionOperator = newEnum("SyntheticCertificatePropertyAssertionOperator", []value[configv1.SyntheticTestCertificatePropertyAssertionOperator]{
	{v1: configv1.SyntheticTestCertificatePropertyAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestCertificatePropertyAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configv1.SyntheticTestCertificatePropertyAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configv1.SyntheticTestCertificatePropertyAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configv1.SyntheticTestCertificatePropertyAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configv1.SyntheticTestCertificatePropertyAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticTLSVersionAssertionOperator = newEnum("SyntheticTLSVersionAssertionOperator", []value[configv1.SyntheticTestTLSVersionAssertionOperator]{
	{v1: configv1.SyntheticTestTLSVersionAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configv1.SyntheticTestTLSVersionAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configv1.SyntheticTestTLSVersionAssertionOperatorGREATERTHANOREQUAL, alias: "GREATER_THAN_OR_EQUAL"},
	{v1: configv1.SyntheticTestTLSVersionAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configv1.SyntheticTestTLSVersionAssertionOperatorLESSTHANOREQUAL, alias: "LESS_THAN_OR_EQUAL"},
})

var SyntheticTLSVersionBound = newEnum("SyntheticTLSVersionBound", []value[configv1.TLSVersionAssertionBound]{
	{v1: configv1.TLSVersionAssertionBoundMIN, alias: "MIN"},
	{v1: configv1.TLSVersionAssertionBoundMAX, alias: "MAX"},
})

var SyntheticTLSVersion = newEnum("SyntheticTLSVersion", []value[configv1.TLSVersionAssertionVersion]{
	{v1: configv1.TLSVersionAssertionVersionTLS10, alias: "TLS_1_0"},
	{v1: configv1.TLSVersionAssertionVersionTLS11, alias: "TLS_1_1"},
	{v1: configv1.TLSVersionAssertionVersionTLS12, alias: "TLS_1_2"},
	{v1: configv1.TLSVersionAssertionVersionTLS13, alias: "TLS_1_3"},
})
