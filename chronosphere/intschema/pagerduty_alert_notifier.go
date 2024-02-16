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

type PagerdutyAlertNotifier struct {
	Name                  string                        `intschema:"name"`
	Slug                  string                        `intschema:"slug,optional,computed"`
	Severity              string                        `intschema:"severity"`
	Url                   string                        `intschema:"url"`
	BasicAuthPassword     string                        `intschema:"basic_auth_password,optional"`
	BasicAuthUsername     string                        `intschema:"basic_auth_username,optional"`
	BearerToken           string                        `intschema:"bearer_token,optional"`
	Class                 string                        `intschema:"class,optional"`
	Client                string                        `intschema:"client,optional"`
	ClientUrl             string                        `intschema:"client_url,optional"`
	Component             string                        `intschema:"component,optional"`
	Description           string                        `intschema:"description,optional"`
	Details               map[string]string             `intschema:"details,optional"`
	Group                 string                        `intschema:"group,optional"`
	Image                 []PagerdutyAlertNotifierImage `intschema:"image,optional"`
	Link                  []PagerdutyAlertNotifierLink  `intschema:"link,optional"`
	ProxyUrl              string                        `intschema:"proxy_url,optional"`
	RoutingKey            string                        `intschema:"routing_key,optional"`
	SendResolved          bool                          `intschema:"send_resolved,optional,default:true"`
	ServiceKey            string                        `intschema:"service_key,optional"`
	TlsInsecureSkipVerify bool                          `intschema:"tls_insecure_skip_verify,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *PagerdutyAlertNotifier) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.PagerdutyAlertNotifier, d, o)
}

func (o *PagerdutyAlertNotifier) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *PagerdutyAlertNotifier) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_pagerduty_alert_notifier", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *PagerdutyAlertNotifier) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_pagerduty_alert_notifier",
		ID:   o.HCLID,
	}.AsID()
}

type PagerdutyAlertNotifierLink struct {
	Href string `intschema:"href"`
	Text string `intschema:"text,optional"`
}

type PagerdutyAlertNotifierImage struct {
	Src  string `intschema:"src"`
	Alt  string `intschema:"alt,optional"`
	Href string `intschema:"href,optional"`
}