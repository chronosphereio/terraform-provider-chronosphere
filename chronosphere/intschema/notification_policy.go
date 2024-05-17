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

type NotificationPolicy struct {
	Name                   string                       `intschema:"name,optional"`
	Slug                   string                       `intschema:"slug,optional,computed"`
	NotificationPolicyData tfid.ID                      `intschema:"notification_policy_data,optional,computed"`
	TeamId                 tfid.ID                      `intschema:"team_id,optional"`
	IsIndependent          bool                         `intschema:"is_independent,computed"`
	Route                  []NotificationRoute          `intschema:"route,optional"`
	Override               []NotificationPolicyOverride `intschema:"override,optional"`

	// Internal identifier used in the .state file, i.e. ResourceData.Id().
	// Cannot be set, else ToResourceData will panic.
	StateID string `intschema:"-"`

	// HCL-level identifier used in the .tf file. FromResourceData will always
	// leave this empty, and ToResourceData will panic if set.
	HCLID string `intschema:"-"`
}

func (o *NotificationPolicy) FromResourceData(d convertintschema.ResourceGetter) error {
	return convertintschema.FromResourceData(tfschema.NotificationPolicy, d, o)
}

func (o *NotificationPolicy) ToResourceData(d *schema.ResourceData) diag.Diagnostics {
	return convertintschema.ToResourceData(o, d)
}

func (o *NotificationPolicy) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_notification_policy", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *NotificationPolicy) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_notification_policy",
		ID:   o.HCLID,
	}.AsID()
}

type NotificationPolicyOverride struct {
	AlertLabelMatcher []Matcher           `intschema:"alert_label_matcher"`
	Route             []NotificationRoute `intschema:"route,optional"`
}
