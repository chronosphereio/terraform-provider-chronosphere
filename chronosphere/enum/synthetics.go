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
	configunstable "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

var SyntheticTestType = newEnum("SyntheticTestType", []value[configunstable.SyntheticTestTestType]{
	{v1: configunstable.SyntheticTestTestTypeHTTP, alias: "HTTP"},
	{v1: configunstable.SyntheticTestTestTypeDNS, alias: "DNS"},
	{v1: configunstable.SyntheticTestTestTypeTCP, alias: "TCP"},
	{v1: configunstable.SyntheticTestTestTypeTLS, alias: "TLS"},
})

var SyntheticTestStatus = newEnum("SyntheticTestStatus", []value[configunstable.ConfigunstableSyntheticTestStatus]{
	{v1: configunstable.ConfigunstableSyntheticTestStatusENABLED, alias: "ENABLED"},
	{v1: configunstable.ConfigunstableSyntheticTestStatusPAUSED, alias: "PAUSED"},
})

var SyntheticTestLocation = newEnum("SyntheticTestLocation", []value[configunstable.SyntheticTestTestLocation]{
	{v1: configunstable.SyntheticTestTestLocationGCPUSOREGON, alias: "GCP_US_OREGON"},
	{v1: configunstable.SyntheticTestTestLocationGCPUSVIRGINIA, alias: "GCP_US_VIRGINIA"},
})

var SyntheticHTTPMethod = newEnum("SyntheticHTTPMethod", []value[configunstable.HTTPTestConfigHTTPMethod]{
	{v1: configunstable.HTTPTestConfigHTTPMethodGET, alias: "GET"},
	{v1: configunstable.HTTPTestConfigHTTPMethodPOST, alias: "POST"},
})

// SyntheticHTTPVersion aliases away the wire values' redundant HTTP_VERSION_ prefix.
var SyntheticHTTPVersion = newEnum("SyntheticHTTPVersion", []value[configunstable.HTTPTestConfigHTTPVersion]{
	{v1: configunstable.HTTPTestConfigHTTPVersionHTTPVERSIONHTTP11, alias: "HTTP_1_1"},
	{v1: configunstable.HTTPTestConfigHTTPVersionHTTPVERSIONHTTP2, alias: "HTTP_2"},
})

// SyntheticHTTPContentType aliases away the wire values' redundant CONTENT_TYPE_ prefix.
var SyntheticHTTPContentType = newEnum("SyntheticHTTPContentType", []value[configunstable.HTTPTestConfigContentType]{
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONJSON, alias: "APPLICATION_JSON"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONOCTETSTREAM, alias: "APPLICATION_OCTET_STREAM"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPEAPPLICATIONXWWWFORMURLENCODED, alias: "APPLICATION_X_WWW_FORM_URLENCODED"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPEGRAPHQL, alias: "GRAPHQL"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPEMULTIPARTFORMDATA, alias: "MULTIPART_FORM_DATA"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPETEXTHTML, alias: "TEXT_HTML"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPETEXTPLAIN, alias: "TEXT_PLAIN"},
	{v1: configunstable.HTTPTestConfigContentTypeCONTENTTYPETEXTXML, alias: "TEXT_XML"},
})

// SyntheticOAuth2TokenAuthMethod aliases away the wire values' redundant OTAM_ prefix.
var SyntheticOAuth2TokenAuthMethod = newEnum("SyntheticOAuth2TokenAuthMethod", []value[configunstable.ConfigunstableOAuth2TokenAuthMethod]{
	{v1: configunstable.ConfigunstableOAuth2TokenAuthMethodOTAMBASICAUTHHEADER, alias: "BASIC_AUTH_HEADER"},
	{v1: configunstable.ConfigunstableOAuth2TokenAuthMethodOTAMREQUESTBODY, alias: "REQUEST_BODY"},
})

var SyntheticBodyAssertionOperator = newEnum("SyntheticBodyAssertionOperator", []value[configunstable.SyntheticTestBodyAssertionOperator]{
	{v1: configunstable.SyntheticTestBodyAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestBodyAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configunstable.SyntheticTestBodyAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configunstable.SyntheticTestBodyAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configunstable.SyntheticTestBodyAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestBodyAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticBodyHashAlgorithm = newEnum("SyntheticBodyHashAlgorithm", []value[configunstable.BodyHashAssertionAlgorithm]{
	{v1: configunstable.BodyHashAssertionAlgorithmMD5, alias: "MD5"},
	{v1: configunstable.BodyHashAssertionAlgorithmSHA1, alias: "SHA1"},
	{v1: configunstable.BodyHashAssertionAlgorithmSHA256, alias: "SHA256"},
})

var SyntheticBodyJSONPathAssertionOperator = newEnum("SyntheticBodyJSONPathAssertionOperator", []value[configunstable.SyntheticTestBodyJSONPathAssertionOperator]{
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configunstable.SyntheticTestBodyJSONPathAssertionOperatorISUNDEFINED, alias: "IS_UNDEFINED"},
})

var SyntheticBodyJSONPathMatchType = newEnum("SyntheticBodyJSONPathMatchType", []value[configunstable.BodyJSONPathAssertionMatchType]{
	{v1: configunstable.BodyJSONPathAssertionMatchTypeFIRSTELEMENT, alias: "FIRST_ELEMENT"},
	{v1: configunstable.BodyJSONPathAssertionMatchTypeEVERYELEMENT, alias: "EVERY_ELEMENT"},
	{v1: configunstable.BodyJSONPathAssertionMatchTypeATLEASTONEELEMENT, alias: "AT_LEAST_ONE_ELEMENT"},
	{v1: configunstable.BodyJSONPathAssertionMatchTypeSERIALIZATION, alias: "SERIALIZATION"},
})

var SyntheticBodyJSONSchemaDraft = newEnum("SyntheticBodyJSONSchemaDraft", []value[configunstable.BodyJSONSchemaAssertionSchemaDraft]{
	{v1: configunstable.BodyJSONSchemaAssertionSchemaDraftDRAFT06, alias: "DRAFT_06"},
	{v1: configunstable.BodyJSONSchemaAssertionSchemaDraftDRAFT07, alias: "DRAFT_07"},
})

var SyntheticBodyXPathAssertionOperator = newEnum("SyntheticBodyXPathAssertionOperator", []value[configunstable.SyntheticTestBodyXPathAssertionOperator]{
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configunstable.SyntheticTestBodyXPathAssertionOperatorISUNDEFINED, alias: "IS_UNDEFINED"},
})

var SyntheticHeaderAssertionOperator = newEnum("SyntheticHeaderAssertionOperator", []value[configunstable.SyntheticTestHeaderAssertionOperator]{
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorGREATERTHANOREQUAL, alias: "GREATER_THAN_OR_EQUAL"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorLESSTHANOREQUAL, alias: "LESS_THAN_OR_EQUAL"},
	{v1: configunstable.SyntheticTestHeaderAssertionOperatorNOTEXISTS, alias: "NOT_EXISTS"},
})

var SyntheticResponseTimeAssertionOperator = newEnum("SyntheticResponseTimeAssertionOperator", []value[configunstable.SyntheticTestResponseTimeAssertionOperator]{
	{v1: configunstable.SyntheticTestResponseTimeAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
})

var SyntheticResponseTimeScope = newEnum("SyntheticResponseTimeScope", []value[configunstable.ResponseTimeAssertionScope]{
	{v1: configunstable.ResponseTimeAssertionScopeINCLUDINGDNS, alias: "INCLUDING_DNS"},
	{v1: configunstable.ResponseTimeAssertionScopeWITHOUTDNS, alias: "WITHOUT_DNS"},
})

var SyntheticStatusCodeAssertionOperator = newEnum("SyntheticStatusCodeAssertionOperator", []value[configunstable.SyntheticTestStatusCodeAssertionOperator]{
	{v1: configunstable.SyntheticTestStatusCodeAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestStatusCodeAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configunstable.SyntheticTestStatusCodeAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestStatusCodeAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticDNSRecordAssertionOperator = newEnum("SyntheticDNSRecordAssertionOperator", []value[configunstable.SyntheticTestDNSRecordAssertionOperator]{
	{v1: configunstable.SyntheticTestDNSRecordAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestDNSRecordAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configunstable.SyntheticTestDNSRecordAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestDNSRecordAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticDNSRecordMatchScope = newEnum("SyntheticDNSRecordMatchScope", []value[configunstable.DNSRecordAssertionMatchScope]{
	{v1: configunstable.DNSRecordAssertionMatchScopeATLEASTONE, alias: "AT_LEAST_ONE"},
	{v1: configunstable.DNSRecordAssertionMatchScopeEVERYAVAILABLE, alias: "EVERY_AVAILABLE"},
})

var SyntheticDNSRecordType = newEnum("SyntheticDNSRecordType", []value[configunstable.DNSRecordAssertionRecordType]{
	{v1: configunstable.DNSRecordAssertionRecordTypeA, alias: "A"},
	{v1: configunstable.DNSRecordAssertionRecordTypeAAAA, alias: "AAAA"},
	{v1: configunstable.DNSRecordAssertionRecordTypeCNAME, alias: "CNAME"},
	{v1: configunstable.DNSRecordAssertionRecordTypeMX, alias: "MX"},
	{v1: configunstable.DNSRecordAssertionRecordTypeNS, alias: "NS"},
	{v1: configunstable.DNSRecordAssertionRecordTypeSOA, alias: "SOA"},
	{v1: configunstable.DNSRecordAssertionRecordTypeSRV, alias: "SRV"},
	{v1: configunstable.DNSRecordAssertionRecordTypeTXT, alias: "TXT"},
})

var SyntheticConnectionAssertionOperator = newEnum("SyntheticConnectionAssertionOperator", []value[configunstable.SyntheticTestConnectionAssertionOperator]{
	{v1: configunstable.SyntheticTestConnectionAssertionOperatorEQUALS, alias: "EQUALS"},
})

var SyntheticConnectionStatus = newEnum("SyntheticConnectionStatus", []value[configunstable.SyntheticTestConnectionAssertionStatus]{
	{v1: configunstable.SyntheticTestConnectionAssertionStatusESTABLISHED, alias: "ESTABLISHED"},
	{v1: configunstable.SyntheticTestConnectionAssertionStatusREFUSED, alias: "REFUSED"},
	{v1: configunstable.SyntheticTestConnectionAssertionStatusTIMEOUT, alias: "TIMEOUT"},
})

var SyntheticNetworkHopsAssertionOperator = newEnum("SyntheticNetworkHopsAssertionOperator", []value[configunstable.SyntheticTestNetworkHopsAssertionOperator]{
	{v1: configunstable.SyntheticTestNetworkHopsAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestNetworkHopsAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configunstable.SyntheticTestNetworkHopsAssertionOperatorGREATERTHANOREQUAL, alias: "GREATER_THAN_OR_EQUAL"},
	{v1: configunstable.SyntheticTestNetworkHopsAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configunstable.SyntheticTestNetworkHopsAssertionOperatorLESSTHANOREQUAL, alias: "LESS_THAN_OR_EQUAL"},
})

var SyntheticCertificateAssertionOperator = newEnum("SyntheticCertificateAssertionOperator", []value[configunstable.SyntheticTestCertificateAssertionOperator]{
	{v1: configunstable.SyntheticTestCertificateAssertionOperatorEXPIRESINLESSTHANDAYS, alias: "EXPIRES_IN_LESS_THAN_DAYS"},
	{v1: configunstable.SyntheticTestCertificateAssertionOperatorEXPIRESINMORETHANDAYS, alias: "EXPIRES_IN_MORE_THAN_DAYS"},
})

var SyntheticCertificatePropertyAssertionOperator = newEnum("SyntheticCertificatePropertyAssertionOperator", []value[configunstable.SyntheticTestCertificatePropertyAssertionOperator]{
	{v1: configunstable.SyntheticTestCertificatePropertyAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestCertificatePropertyAssertionOperatorNOTEQUALS, alias: "NOT_EQUALS"},
	{v1: configunstable.SyntheticTestCertificatePropertyAssertionOperatorCONTAINS, alias: "CONTAINS"},
	{v1: configunstable.SyntheticTestCertificatePropertyAssertionOperatorNOTCONTAINS, alias: "NOT_CONTAINS"},
	{v1: configunstable.SyntheticTestCertificatePropertyAssertionOperatorREGEXEQUALS, alias: "REGEX_EQUALS"},
	{v1: configunstable.SyntheticTestCertificatePropertyAssertionOperatorREGEXNOTEQUALS, alias: "REGEX_NOT_EQUALS"},
})

var SyntheticTLSVersionAssertionOperator = newEnum("SyntheticTLSVersionAssertionOperator", []value[configunstable.SyntheticTestTLSVersionAssertionOperator]{
	{v1: configunstable.SyntheticTestTLSVersionAssertionOperatorEQUALS, alias: "EQUALS"},
	{v1: configunstable.SyntheticTestTLSVersionAssertionOperatorGREATERTHAN, alias: "GREATER_THAN"},
	{v1: configunstable.SyntheticTestTLSVersionAssertionOperatorGREATERTHANOREQUAL, alias: "GREATER_THAN_OR_EQUAL"},
	{v1: configunstable.SyntheticTestTLSVersionAssertionOperatorLESSTHAN, alias: "LESS_THAN"},
	{v1: configunstable.SyntheticTestTLSVersionAssertionOperatorLESSTHANOREQUAL, alias: "LESS_THAN_OR_EQUAL"},
})

var SyntheticTLSVersionBound = newEnum("SyntheticTLSVersionBound", []value[configunstable.TLSVersionAssertionBound]{
	{v1: configunstable.TLSVersionAssertionBoundMIN, alias: "MIN"},
	{v1: configunstable.TLSVersionAssertionBoundMAX, alias: "MAX"},
})

var SyntheticTLSVersion = newEnum("SyntheticTLSVersion", []value[configunstable.TLSVersionAssertionVersion]{
	{v1: configunstable.TLSVersionAssertionVersionTLS10, alias: "TLS_1_0"},
	{v1: configunstable.TLSVersionAssertionVersionTLS11, alias: "TLS_1_1"},
	{v1: configunstable.TLSVersionAssertionVersionTLS12, alias: "TLS_1_2"},
	{v1: configunstable.TLSVersionAssertionVersionTLS13, alias: "TLS_1_3"},
})
