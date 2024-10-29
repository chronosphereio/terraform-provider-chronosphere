// Copyright 2024 Chronosphere Inc.
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
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/localid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// NotificationPolicyDryRunCount tracks how many times dry run is run during validation for testing.
var NotificationPolicyDryRunCount atomic.Int64

func resourceNotificationPolicy() *schema.Resource {
	independent := newIndependentNotificationPolicy()
	npr := &notificationPolicyResourceMeta{
		validateDryRunOptions: independent.ValidateDryRunOptions,
	}
	r := &schema.Resource{
		CreateContext: npInlineOrIndependent("create", resourceInlineNotificationPolicyCreate, independent.CreateContext),
		ReadContext:   npInlineOrIndependent("read", resourceInlineNotificationPolicyRead, independent.ReadContext),
		UpdateContext: npInlineOrIndependent("update", resourceInlineNotificationPolicyUpdate, independent.UpdateContext),
		DeleteContext: npInlineOrIndependent("delete", resourceInlineNotificationPolicyDelete, independent.DeleteContext),
		CustomizeDiff: npr.resourceNotificationPolicyCustomizeDiff,
		Schema:        tfschema.NotificationPolicy,
		Importer: &schema.ResourceImporter{
			StateContext: npr.resourceNotificationPolicyImport,
		},
	}
	npr.readContext = r.ReadContext
	return r
}

// npStableOrUnstable only runs the unstable code path if the target
// notification policy is independent, as the unstable code path does not handle
// inline policies.
func npInlineOrIndependent[F resourceFunc](funcName string, inlineFunc, independentFunc F) F {
	return func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
		np := &intschema.NotificationPolicy{}
		if err := np.FromResourceData(d); err != nil {
			return diag.Errorf(err.Error())
		}
		if isNotificationPolicyIndependent(np) {
			tflog.Info(ctx, "calling independent policy resource", map[string]any{
				"funcName": funcName,
				"policy":   np,
			})
			return independentFunc(ctx, d, meta)
		}
		tflog.Info(ctx, "calling inline policy resource", map[string]any{
			"funcName": funcName,
			"policy":   np,
		})
		return inlineFunc(ctx, d, meta)
	}
}

func resourceInlineNotificationPolicyCreate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	if err := setInlinePolicyData(d); err != nil {
		return err
	}

	d.SetId(localid.NewLocalID())
	return nil
}

func resourceInlineNotificationPolicyRead(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	tflog.Info(ctx, "not querying inline notification policy from server", map[string]any{
		"id": d.Id(),
	})

	if err := setInlinePolicyData(d); err != nil {
		return err
	}

	return nil
}

func resourceInlineNotificationPolicyUpdate(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	if err := setInlinePolicyData(d); err != nil {
		return err
	}

	return nil
}

func resourceInlineNotificationPolicyDelete(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	// Since inline policies don't manage a server-side object, no API calls are needed.
	d.SetId("")
	return nil
}

// notificationPolicyResourceMeta is used for methods that need access to the schema functions.
type notificationPolicyResourceMeta struct {
	validateDryRunOptions func(dryRunCounter *atomic.Int64, opts ValidateDryRunOpts[*configmodels.Configv1NotificationPolicy]) schema.CustomizeDiffFunc
	readContext           schema.ReadContextFunc
}

// resourceNotificationPolicyImport imports a notification policy.
// The ID accepted by the import can be either:
//   - base64-encoded JSON blob that represents an importmodel.NotificationPolicy.
//     This is a deprecated case still used by customers where buckets inline a notification policy definition.
//   - slug of the notification policy.
func (npr *notificationPolicyResourceMeta) resourceNotificationPolicyImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	policyModel, err := importPolicyModel(d.Id())
	if err != nil {
		// If we can't decode the policy model, we assume it's an independent policy (e.g., Id is the policy slug).
		tflog.Info(ctx, "importing notification policy by slug since import id is not an encoded JSON config", map[string]any{
			"slug":        d.Id(),
			"decodeError": err,
		})
	} else {
		// Inline policy, set route and override (only fields expected in an inline policy)
		// so we can generate a notification_policy_data.
		// We need to set these from the import ID since we don't have access to the resource config
		// as part of `d` at import time.
		var errors diag.Diagnostics
		errors = setKey(errors, d, "route", policyModel["route"])
		errors = setKey(errors, d, "override", policyModel["override"])
		d.SetId(localid.NewImportedID())
		if errors.HasError() {
			return nil, diagError(errors)
		}
	}

	if err := npr.readContext(ctx, d, meta); err.HasError() {
		return nil, diagError(err)
	}

	return []*schema.ResourceData{d}, nil
}

func importPolicyModel(id string) (map[string]any, error) {
	marshalled, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return nil, err
	}

	var policyModel map[string]any
	if err := json.Unmarshal(marshalled, &policyModel); err != nil {
		return nil, err
	}

	return policyModel, nil
}

func (npr *notificationPolicyResourceMeta) resourceNotificationPolicyCustomizeDiff(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
	// The policy name is used to determine if a policy is inline or independent. Inline policies must not have a name;
	// independent policies must have a name. See https://stackoverflowteams.com/c/chronosphere/questions/648 for details.
	//
	// The is_independent field is set to true if the policy is given a name, and false if its name is removed.
	// This field is marked ForceNew, so if it changes, the resource is recreated.
	//
	// This is done in this diff function rather than marking the name field ForceNew to to avoid unnecessary
	// deletes-and-recreates when an independent policy is renamed.
	hasName := stringAttrLikelyDefined(diff, "name")
	oldName, _ := diff.GetChange("name")
	hadName := oldName.(string) != ""
	if hadName != hasName {
		tflog.Info(ctx, "updating `is_independent` field to ForceNew notification policy", map[string]interface{}{
			"nameBefore":     hadName,
			"nameNow":        hasName,
			"nameChanged":    diff.HasChange("name"),
			"is_independent": hasName,
		})
		if err := diff.SetNew("is_independent", hasName); err != nil {
			return err
		}
	}

	// This forces a diff to notification_policy_data if the route list, or ownership changes.
	// This is not the default terraform behavior, which usually only computes fields on resource creation.
	if changedKeys := diff.GetChangedKeysPrefix(""); len(changedKeys) > 0 {
		tflog.Debug(ctx, "Notification policy change detected, update notification_policy_data", map[string]any{
			"changedKeys": changedKeys,
		})
		if err := diff.SetNewComputed("notification_policy_data"); err != nil {
			return err
		}
	}

	var policy intschema.NotificationPolicy
	if err := policy.FromResourceData(diff); err != nil {
		return err
	}

	// At least one-of route, or override must be set.
	if len(policy.Route) == 0 && len(policy.Override) == 0 {
		return errors.New("specify at least one `route` or `override`")
	}

	// We prevent duplicate routes with the same severity.
	bySeverity := make(map[string]bool)
	for _, o := range policy.Route {
		if _, ok := bySeverity[o.Severity]; ok {
			return fmt.Errorf("duplicate route with severity=%v", o.Severity)
		}

		bySeverity[o.Severity] = true
	}

	dryRunOpts := ValidateDryRunOpts[*configmodels.Configv1NotificationPolicy]{
		// Note: We skip notifier fields in route and override as they're
		// within a set, which setUnknownReferences does not support.
		// When there's an unknown notifier slug, the list of notifiers is empty.
		// Since the server accepts empty lists, we don't need to set dummy values for dry-run
		// validation to work properly.
		SetUnknownReferencesSkip: []string{
			"route.[].notifiers.[]",
			"override.[].route.[].notifiers.[]",
		},

		ModifyAPIModel: func(apiPolicy *configmodels.Configv1NotificationPolicy) {
			if !isNotificationPolicyIndependentForCustomizeDiff(diff, policy) {
				// If the policy is inline, we want to populate required fields that are inherited from
				// the bucket normally (name, bucket_slug). Otherwise, the API rejects the missing fields.
				apiPolicy.BucketSlug = dummyRef.Slug()
				apiPolicy.Name = dummyRef.Slug()

				// Inline policies don't have a slug, but use a dummy slug since it's required for Update.
				apiPolicy.Slug = dummyRef.Slug()
			}
		},
	}

	if err := npr.validateDryRunOptions(&NotificationPolicyDryRunCount, dryRunOpts)(ctx, diff, meta); err != nil {
		return err
	}

	return nil
}

// stringAttrLikelyDefined returns true if an attribute in a diff:
// 1. Is set to a value in Terraform config (not state only, so computed values return false), AND
// 2. That value is likely to have a value that is not an empty string.
//
// Regarding 2, it is not always known at diff time what the value may be since it could reference another
// resource field.
func stringAttrLikelyDefined(diff *schema.ResourceDiff, name string) bool {
	ctyVal := diff.GetRawConfig().GetAttr(name)
	if ctyVal.IsNull() {
		return false
	}

	// This is the case of the field referencing an attribute of another resource.
	// Example: team_id = chronosphere_team.everyone.id
	// This code assumes that whatever attribute is being referenced is ultimately a non-empty string.
	if !ctyVal.IsKnown() {
		return true
	}

	return ctyVal.AsString() != ""
}

func setInlinePolicyData(d *schema.ResourceData) diag.Diagnostics {
	policy, diagErr := expandNotificationPolicyRaw(d)
	if diagErr != nil {
		return diagErr
	}

	npData, err := json.Marshal(policy)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("notification_policy_data", string(npData)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func expandNotificationPolicyRaw(
	d *schema.ResourceData,
) (*NotificationPolicyData, diag.Diagnostics) {
	p := &intschema.NotificationPolicy{}
	if err := p.FromResourceData(d); err != nil {
		return nil, diag.Errorf(err.Error())
	}
	m, err := expandNotificationPolicy(p)
	if err != nil {
		return nil, diag.Errorf(err.Error())
	}
	return m, nil
}

// expandNotificationPolicy converts a notification policy resource to the corresponding API model type.
func expandNotificationPolicy(p *intschema.NotificationPolicy) (*NotificationPolicyData, error) {
	notifiers, err := expandNotificationPolicyRoutes(p.Route)
	if err != nil {
		return nil, err
	}
	routes := &configmodels.NotificationPolicyRoutes{
		Defaults: notifiers,
	}

	for _, o := range p.Override {
		notifiers, err := expandNotificationPolicyRoutes(o.Route)
		if err != nil {
			return nil, err
		}
		routes.Overrides = append(routes.Overrides, &configmodels.NotificationPolicyRoutesOverride{
			AlertLabelMatchers: sliceutil.Map(o.AlertLabelMatcher, expandMatcherSchema),
			Notifiers:          notifiers,
		})
	}

	policyRoutes, err := routesFromModel(routes)
	if err != nil {
		return nil, err
	}

	return &NotificationPolicyData{
		Routes: policyRoutes,
	}, nil
}

func expandNotificationPolicyRoutes(routes []intschema.NotificationRoute) (*configmodels.RoutesSeverityNotifiers, error) {
	// In this method we want to add a key for the severity if any route with the
	// severity exists. It doesn't have to have any notifiers in it.
	out := &configmodels.RoutesSeverityNotifiers{}
	for _, r := range routes {
		notifierList := &configmodels.RoutesNotifierList{}
		for _, notifierID := range r.Notifiers {
			notifierList.NotifierSlugs = append(notifierList.NotifierSlugs, notifierID.Slug())
		}

		duration, err := ParseOptionalDuration(r.RepeatInterval)
		if err != nil {
			panic(fmt.Sprintf("invalid repeat interval %q", r.RepeatInterval))
		}
		notifierList.RepeatIntervalSecs = int32(duration.Seconds())

		notifierList.GroupBy = notificationRouteGroupByToModel(r.GroupBy)

		if r.Severity == "warn" {
			if out.Warn != nil {
				return nil, fmt.Errorf("duplicate route with severity=warn")
			}
			out.Warn = notifierList
		} else if r.Severity == "critical" {
			if out.Critical != nil {
				return nil, fmt.Errorf("duplicate route with severity=critical")
			}
			out.Critical = notifierList
		} else {
			panic(fmt.Sprintf("unknown severity: %s", r.Severity))
		}
	}

	return out, nil
}

func isNotificationPolicyIndependent(p *intschema.NotificationPolicy) bool {
	if p.StateID != "" {
		// If policy exists, determine independence based on ID in the TF state
		return !localid.IsLocalID(p.StateID)
	}

	// Otherwise, policy is independent if name is set in the TF definition.
	return p.Name != ""
}

// isNotificationPolicyIndependentForCustomizeDiff is a version of isNotificationPolicyIndependent which is safe
// to call in CustomizeDiff when local resource references may not be populated yet.
func isNotificationPolicyIndependentForCustomizeDiff(
	diff *schema.ResourceDiff,
	p intschema.NotificationPolicy,
) bool {
	if p.StateID != "" && localid.IsLocalID(p.StateID) {
		// If policy exists, determine independence based on ID in the TF state
		return false
	}

	return !isRawAttributeNull(diff, "team_id")
}

func expandMatcherSchema(m intschema.Matcher) *configmodels.Configv1LabelMatcher {
	return &configmodels.Configv1LabelMatcher{
		Name:  m.Name,
		Type:  enum.MatcherType.V1(m.Type),
		Value: m.Value,
	}
}
