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

type SlackAlertNotifier struct {
	Name                  string                     `intschema:"name"`
	Slug                  string                     `intschema:"slug,optional,computed"`
	ApiUrl                string                     `intschema:"api_url"`
	Channel               string                     `intschema:"channel"`
	Action                []SlackAlertNotifierAction `intschema:"action,optional"`
	BasicAuthPassword     string                     `intschema:"basic_auth_password,optional"`
	BasicAuthUsername     string                     `intschema:"basic_auth_username,optional"`
	BearerToken           string                     `intschema:"bearer_token,optional"`
	CallbackId            string                     `intschema:"callback_id,optional"`
	Color                 string                     `intschema:"color,optional"`
	Fallback              string                     `intschema:"fallback,optional"`
	Fields                []SlackAlertNotifierFields `intschema:"fields,optional"`
	Footer                string                     `intschema:"footer,optional"`
	IconEmoji             string                     `intschema:"icon_emoji,optional"`
	IconUrl               string                     `intschema:"icon_url,optional"`
	ImageUrl              string                     `intschema:"image_url,optional"`
	LinkNames             bool                       `intschema:"link_names,optional"`
	MrkdwnIn              []string                   `intschema:"mrkdwn_in,optional"`
	Pretext               string                     `intschema:"pretext,optional"`
	ProxyUrl              string                     `intschema:"proxy_url,optional"`
	SendResolved          bool                       `intschema:"send_resolved,optional,default:true"`
	ShortFields           bool                       `intschema:"short_fields,optional"`
	Text                  string                     `intschema:"text,optional"`
	ThumbUrl              string                     `intschema:"thumb_url,optional"`
	Title                 string                     `intschema:"title,optional"`
	TitleLink             string                     `intschema:"title_link,optional"`
	TlsInsecureSkipVerify bool                       `intschema:"tls_insecure_skip_verify,optional"`
	Username              string                     `intschema:"username,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *SlackAlertNotifier) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.SlackAlertNotifier, d, o)
}

func (o *SlackAlertNotifier) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *SlackAlertNotifier) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_slack_alert_notifier", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *SlackAlertNotifier) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_slack_alert_notifier",
		ID:   o.HCLID,
	}.AsID()
}

type SlackAlertNotifierFields struct {
	Short bool   `intschema:"short,optional"`
	Title string `intschema:"title,optional"`
	Value string `intschema:"value,optional"`
}

type SlackAlertNotifierAction struct {
	Name                     string `intschema:"name,optional"`
	ActionConfirmDismissText string `intschema:"action_confirm_dismiss_text,optional"`
	ActionConfirmOkText      string `intschema:"action_confirm_ok_text,optional"`
	ActionConfirmText        string `intschema:"action_confirm_text,optional"`
	ActionConfirmTile        string `intschema:"action_confirm_tile,optional"`
	Style                    string `intschema:"style,optional"`
	Text                     string `intschema:"text,optional"`
	Type                     string `intschema:"type,optional"`
	Url                      string `intschema:"url,optional"`
	Value                    string `intschema:"value,optional"`
}
