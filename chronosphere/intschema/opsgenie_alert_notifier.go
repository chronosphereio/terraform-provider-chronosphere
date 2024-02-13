// Code generated by go generate; DO NOT EDIT.
package intschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/convertintschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
)

var _ tfid.ID // Always use tfid for simplified import generation.

type OpsgenieAlertNotifier struct {
	Name                  string                           `intschema:"name"`
	Slug                  string                           `intschema:"slug,optional,computed"`
	ApiKey                string                           `intschema:"api_key"`
	ApiUrl                string                           `intschema:"api_url,optional"`
	BasicAuthPassword     string                           `intschema:"basic_auth_password,optional"`
	BasicAuthUsername     string                           `intschema:"basic_auth_username,optional"`
	BearerToken           string                           `intschema:"bearer_token,optional"`
	Description           string                           `intschema:"description,optional"`
	Details               map[string]string                `intschema:"details,optional"`
	Message               string                           `intschema:"message,optional"`
	Note                  string                           `intschema:"note,optional"`
	Priority              string                           `intschema:"priority,optional"`
	ProxyUrl              string                           `intschema:"proxy_url,optional"`
	Responder             []OpsgenieAlertNotifierResponder `intschema:"responder,optional"`
	SendResolved          bool                             `intschema:"send_resolved,optional,default:true"`
	Source                string                           `intschema:"source,optional"`
	Tags                  []string                         `intschema:"tags,optional"`
	TlsInsecureSkipVerify bool                             `intschema:"tls_insecure_skip_verify,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *OpsgenieAlertNotifier) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.OpsgenieAlertNotifier, d, o)
}

func (o *OpsgenieAlertNotifier) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *OpsgenieAlertNotifier) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_opsgenie_alert_notifier", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *OpsgenieAlertNotifier) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_opsgenie_alert_notifier",
		ID:   o.HCLID,
	}.AsID()
}

type OpsgenieAlertNotifierResponder struct {
	Name     string `intschema:"name,optional"`
	Type     string `intschema:"type"`
	Id       string `intschema:"id,optional"`
	Username string `intschema:"username,optional"`
}
