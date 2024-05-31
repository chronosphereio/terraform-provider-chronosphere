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

type LogscaleAction struct {
	Name                   string                                `intschema:"name"`
	Slug                   string                                `intschema:"slug,optional,computed"`
	Repository             string                                `intschema:"repository"`
	EmailAction            *LogscaleActionEmailAction            `intschema:"email_action,optional,list_encoded_object"`
	HumioAction            *LogscaleActionHumioAction            `intschema:"humio_action,optional,list_encoded_object"`
	OpsGenieAction         *LogscaleActionOpsGenieAction         `intschema:"ops_genie_action,optional,list_encoded_object"`
	PagerDutyAction        *LogscaleActionPagerDutyAction        `intschema:"pager_duty_action,optional,list_encoded_object"`
	SlackAction            *LogscaleActionSlackAction            `intschema:"slack_action,optional,list_encoded_object"`
	SlackPostMessageAction *LogscaleActionSlackPostMessageAction `intschema:"slack_post_message_action,optional,list_encoded_object"`
	UploadFileAction       *LogscaleActionUploadFileAction       `intschema:"upload_file_action,optional,list_encoded_object"`
	VictorOpsAction        *LogscaleActionVictorOpsAction        `intschema:"victor_ops_action,optional,list_encoded_object"`
	WebhookAction          *LogscaleActionWebhookAction          `intschema:"webhook_action,optional,list_encoded_object"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *LogscaleAction) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.LogscaleAction, d, o)
}

func (o *LogscaleAction) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *LogscaleAction) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_logscale_action", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *LogscaleAction) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_logscale_action",
		ID:   o.HCLID,
	}.AsID()
}

type LogscaleActionWebhookAction struct {
	Method       string            `intschema:"method"`
	Url          string            `intschema:"url"`
	BodyTemplate string            `intschema:"body_template,optional"`
	Headers      map[string]string `intschema:"headers,optional"`
	IgnoreSsl    bool              `intschema:"ignore_ssl,optional"`
	UseProxy     bool              `intschema:"use_proxy,optional"`
}

type LogscaleActionVictorOpsAction struct {
	MessageType string `intschema:"message_type"`
	NotifyUrl   string `intschema:"notify_url"`
	UseProxy    bool   `intschema:"use_proxy,optional"`
}

type LogscaleActionUploadFileAction struct {
	FileName string `intschema:"file_name"`
}

type LogscaleActionSlackPostMessageAction struct {
	ApiToken string            `intschema:"api_token"`
	Channels []string          `intschema:"channels"`
	Fields   map[string]string `intschema:"fields,optional"`
	UseProxy bool              `intschema:"use_proxy,optional"`
}

type LogscaleActionSlackAction struct {
	Url      string            `intschema:"url"`
	Fields   map[string]string `intschema:"fields,optional"`
	UseProxy bool              `intschema:"use_proxy,optional"`
}

type LogscaleActionPagerDutyAction struct {
	RoutingKey string `intschema:"routing_key"`
	Severity   string `intschema:"severity"`
	UseProxy   bool   `intschema:"use_proxy,optional"`
}

type LogscaleActionOpsGenieAction struct {
	ApiUrl      string `intschema:"api_url"`
	OpsGenieKey string `intschema:"ops_genie_key"`
	UseProxy    bool   `intschema:"use_proxy,optional"`
}

type LogscaleActionHumioAction struct {
	IngestToken string `intschema:"ingest_token"`
}

type LogscaleActionEmailAction struct {
	Recipients      []string `intschema:"recipients"`
	AttachCsv       bool     `intschema:"attach_csv,optional"`
	BodyTemplate    string   `intschema:"body_template,optional"`
	SubjectTemplate string   `intschema:"subject_template,optional"`
	UseProxy        bool     `intschema:"use_proxy,optional"`
}
