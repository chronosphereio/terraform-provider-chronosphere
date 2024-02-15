// Copyright 2023 Chronosphere Inc.
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
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/bucket"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/notification_policy"
	configmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/importmodel"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/localid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

type bucketOpType int

const (
	bucketOpCreate bucketOpType = iota + 1

	// Note: Update is split into Pre/Post to handle inline <--> independent
	// changes while ensuring the bucket always has a policy (for monitors in the bucket).
	// Inline to independent: The inline policy can only be deleted after the bucket
	// is updated to reference the independent policy.
	// Independent to inline: The inline policy should be created before the bucket
	// reference to the independent policy is removed.
	// To handle these cases, the update is split into:
	// UpdatePre: Only upsert the policy if required, don't delete.
	// UpdatePost: Only delete the policy if required, don't upsert.
	bucketOpUpdatePre
	bucketOpUpdatePost
	bucketOpDelete
)

// BucketDryRunCount tracks how many times dry run is run during validation for testing.
var BucketDryRunCount atomic.Int64

// resourceBucket represents a bucket and an optional default notification policy.
func resourceBucket() *schema.Resource {
	r := newGenericResource[
		*configmodels.Configv1Bucket,
		intschema.Bucket,
		*intschema.Bucket,
	](
		"bucket",
		bucketConverter{},
		generatedBucket{},
	)

	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: resourceBucketCreate,
		ReadContext:   resourceBucketRead,
		UpdateContext: resourceBucketUpdate,
		DeleteContext: resourceBucketDelete,
		Schema:        tfschema.Bucket,
		CustomizeDiff: r.ValidateDryRun(&BucketDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: resourceBucketImport,
		},
	}
}

// resourceBucketCreate creates a new bucket. If a notification policy is given, it is also created.
func resourceBucketCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = newBucketContext(ctx)
	cli := getConfigClient(meta)

	b, err := resourceBucketBuild(d)
	if err != nil {
		return diag.Errorf("unable to build bucket: %v", err)
	}
	resp, err := cli.Bucket.CreateBucket(&bucket.CreateBucketParams{
		Context: ctx,
		Body: &configmodels.Configv1CreateBucketRequest{
			Bucket: b,
		},
	})
	if err != nil {
		return diag.Errorf("unable to create bucket %s: %v", b.Slug, clienterror.Wrap(err))
	}

	d.SetId(resp.Payload.Bucket.Slug)

	// Create the notification policy for the bucket
	if err := reconcileNotificationPolicy(ctx, d, meta, bucketOpCreate); err != nil {
		// If the notification policy couldn't be created, fully roll back the create by deleting the bucket
		req := &bucket.DeleteBucketParams{
			Context: ctx,
			Slug:    d.Id(),
		}
		if _, err := cli.Bucket.DeleteBucket(req); err != nil { //nolint:staticcheck
			tflog.Warn(ctx, "error deleting bucket on reconcileNotificationPolicy error", map[string]any{
				"error": err,
			})
			// Move along and return the first error
		}

		return err
	}

	return nil
}

// resourceBucketRead reads a bucket.
func resourceBucketRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = newBucketContext(ctx)
	cfgCli := getConfigClient(meta)

	resp, err := cfgCli.Bucket.ReadBucket(&bucket.ReadBucketParams{
		Context: ctx,
		Slug:    d.Id(),
	})
	if err != nil {
		if clienterror.IsNotFound(err) {
			setResourceNotFound(d)
			return nil
		}
		return diag.Errorf("unable to read bucket: %v", clienterror.Wrap(err))
	}

	bucketSlug := d.Id()
	b, err := bucketConverter{}.fromModel(resp.Payload.Bucket)
	if err != nil {
		diag.Errorf("unable to convert bucket: %v", clienterror.Wrap(err))
	}
	b.NotificationPolicySlug = d.Get("notification_policy_slug").(string)

	// Also read the notification policy to ensure its existence
	if b.NotificationPolicySlug != "" {
		res, err := cfgCli.NotificationPolicy.ReadNotificationPolicy(&notification_policy.ReadNotificationPolicyParams{
			Context: ctx,
			Slug:    bucketSlug, // for inline policies, the slug of the policy is the same as the bucket.
		})
		if err != nil {
			if clienterror.IsNotFound(err) {
				b.NotificationPolicySlug = ""
			} else {
				return diag.Errorf("unable to get notification policy: %v", clienterror.Wrap(err))
			}
		} else {
			notificationPolicyData, err := notificationPolicyResponseToInlineData(res)
			if err != nil {
				return diag.Errorf("unable to convert notification policy: %v", err)
			}
			b.NotificationPolicyData = notificationPolicyData
		}
	}
	return b.ToResourceData(d)
}

// resourceBucketUpdate updates a bucket.
// It will also create/update/delete the bucket's notification policy if added/changed/removed.
func resourceBucketUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = newBucketContext(ctx)
	cli := getConfigClient(meta)

	b, err := resourceBucketBuild(d)
	if err != nil {
		return diag.Errorf("unable to build bucket: %v", err)
	}

	if err := reconcileNotificationPolicy(ctx, d, meta, bucketOpUpdatePre); err != nil {
		return err
	}

	if _, err := cli.Bucket.UpdateBucket(&bucket.UpdateBucketParams{
		Context: ctx,
		Slug:    b.Slug,
		Body: bucket.UpdateBucketBody{
			Bucket: b,
		},
	}); err != nil {
		return diag.Errorf("unable to update bucket: %v", clienterror.Wrap(err))
	}

	if err := reconcileNotificationPolicy(ctx, d, meta, bucketOpUpdatePost); err != nil {
		return err
	}

	return nil
}

// resourceBucketDelete deletes a bucket.
func resourceBucketDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = newBucketContext(ctx)
	cli := getConfigClient(meta)

	// Delete the notification policy, if set
	if err := reconcileNotificationPolicy(ctx, d, meta, bucketOpDelete); err != nil {
		return err
	}

	if _, err := cli.Bucket.DeleteBucket(&bucket.DeleteBucketParams{
		Context: ctx,
		Slug:    d.Id(),
	}); clienterror.IsNotFound(err) {
		// Already deleted on the server, treat as success.
	} else if err != nil {
		return diag.Errorf("unable to delete bucket: %v", clienterror.Wrap(err))
	}

	d.SetId("")

	return nil
}

// resourceBucketImport imports a bucket along with its inline notification policy.
func resourceBucketImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	bucketModel, err := importmodel.ParseBucket(d.Id())
	if err != nil {
		return nil, errors.Wrap(err, "invalid id")
	}
	d.SetId(bucketModel.Slug)

	if bucketModel.NotificationPolicySlug != "" {
		if err := d.Set("notification_policy_slug", bucketModel.NotificationPolicySlug); err != nil {
			return nil, err
		}
		// notification_policy_data will be set based on the slug in resourceBucketRead.
	}

	if err := resourceBucketRead(ctx, d, meta); err.HasError() {
		return nil, diagError(err)
	}

	return []*schema.ResourceData{d}, nil
}

// reconcileNotificationPolicy creates, updates, or deletes a notification policy based on a bucket's policy data.
// Handled cases:
// New bucket,      policy,         policy inline     -> upsert notification policy
// Existing bucket, policy added,   policy inline     -> upsert notification policy
// Existing bucket, policy changed, policy inline     -> upsert notification policy
// Existing bucket, policy removed, policy was inline -> delete policy unless independently managed by chronosphere_notification_policy
// Deleted bucket,  policy,         policy was inline -> delete policy unless independently managed by chronosphere_notification_policy
func reconcileNotificationPolicy(
	ctx context.Context, d *schema.ResourceData, meta any, bucketOp bucketOpType,
) diag.Diagnostics {
	ctx = newBucketContext(ctx)
	configCli := getConfigClient(meta)

	bucketSlug := d.Id()

	bucketPolicySlug := d.Get("notification_policy_slug").(string) // Set if bucket currently managed a policy

	policy, independent, err := policyModelFromBucket(d) // The inlined policy definition, which might be nil if field is unset
	if err != nil {
		return diag.FromErr(err)
	}
	if independent {
		// Can only inline policies that are not independent
		return diag.Errorf("cannot use notification_policy_data of owned policy, use a notification policy without team_id")
	}

	// If policy wasn't set before and still isn't, no-op
	if bucketPolicySlug == "" && policy == nil {
		return nil
	}

	// If bucket owned a policy, possibly try to delete the policy
	// Inline policies should not be deleted before the bucket is updated (UpdatePre)
	// since it could leave a bucket without a policy, even if the user has a default policy set.
	if bucketPolicySlug != "" && bucketOp != bucketOpUpdatePre {
		var deletionReason string
		switch {
		case bucketOp == bucketOpDelete && policy != nil:
			deletionReason = "bucket managing policy is being deleted"
		case policy == nil:
			deletionReason = "policy definition no longer inlined in bucket"
		default:
		}

		if deletionReason != "" {
			tflog.Info(ctx, "deleting notification policy", map[string]any{
				"slug":          bucketPolicySlug,
				"reason":        deletionReason,
				"bucketOp":      bucketOp,
				"policyInlined": policy != nil,
			})

			if _, err := configCli.NotificationPolicy.DeleteNotificationPolicy(&notification_policy.DeleteNotificationPolicyParams{
				Context: ctx,
				Slug:    bucketPolicySlug,
			}); policyIndependenceDeleteConflict(err) {
				tflog.Info(ctx, "not deleting independent notification policy", map[string]any{
					"slug": bucketPolicySlug,
				})
			} else if clienterror.IsNotFound(err) {
				// Already deleted on the server, treat as success.
			} else if err != nil {
				return diag.Errorf("unable to delete notification policy with slug %s: %v", bucketPolicySlug, clienterror.Wrap(err))
			}

			if err := d.Set("notification_policy_slug", ""); err != nil {
				return diag.FromErr(err)
			}
			return nil
		}
	}

	// The only relevant case left after this point is upserting the policy to be managed by this bucket.
	// If policy is nil, no updates are necessary
	if policy == nil {
		return nil
	}
	// Bucket deletions don't require policy updates. Any deletions would have happened above.
	if bucketOp == bucketOpDelete {
		return nil
	}
	// For updates, UpdatePre handles upsert. UpdatePost only needs to handle deletions so we can return early.
	if bucketOp == bucketOpUpdatePost {
		return nil
	}

	// Since the policy is owned by the bucket, use the same name and slug as the bucket on upserts
	policy.BucketSlug = bucketSlug
	policy.Slug = bucketPolicySlug // bucketPolicySlug will be empty for new policy, set for updated policy
	policy.Name = d.Get("name").(string)

	// Upsert the notification policy if it doesn't exist or if the policy changed
	if bucketPolicySlug == "" || d.HasChange("notification_policy_data") {
		tflog.Info(ctx, "upserting notification policy", map[string]any{
			"slug": policy.Slug,
			"new":  bucketPolicySlug == "",
		})

		npSlug := ""
		if policy.Slug == "" {
			resp, err := configCli.NotificationPolicy.CreateNotificationPolicy(&notification_policy.CreateNotificationPolicyParams{
				Context: ctx,
				Body: &configmodels.Configv1CreateNotificationPolicyRequest{
					NotificationPolicy: policy,
				},
			})
			if err != nil {
				return diag.Errorf("unable to create bucket notification policy `%s`: %v", policy.Name, clienterror.Wrap(err))
			}
			npSlug = resp.Payload.NotificationPolicy.Slug
		} else {
			resp, err := configCli.NotificationPolicy.UpdateNotificationPolicy(&notification_policy.UpdateNotificationPolicyParams{
				Context: ctx,
				Slug:    policy.Slug,
				Body: notification_policy.UpdateNotificationPolicyBody{
					NotificationPolicy: policy,
				},
			})
			if err != nil {
				return diag.Errorf("unable to update bucket notification policy `%s`: %v", policy.Name, clienterror.Wrap(err))
			}
			npSlug = resp.Payload.NotificationPolicy.Slug
		}

		if err := d.Set("notification_policy_slug", npSlug); err != nil {
			return diag.FromErr(err)
		}

		return nil
	}

	return nil
}

func policyModelFromBucket(d *schema.ResourceData) (policy *configmodels.Configv1NotificationPolicy, independent bool, err error) {
	policyDataStr := d.Get("notification_policy_data").(string)
	if policyDataStr == "" {
		return nil, false, nil
	}

	if policyDataStr == tfschema.IndependentNotificationPolicyData {
		return nil, true, nil
	}

	policyData, err := unmarshalNotificationPolicyData(policyDataStr)
	if err != nil {
		return nil, false, fmt.Errorf("notification_policy_data is not valid: %w", err)
	}

	policy, err = policyData.ToModel()
	if err != nil {
		return nil, false, fmt.Errorf("notification_policy_data is not valid: %w", err)
	}

	return policy, false, nil
}

// policyIndependenceDeleteConflict returns true if the given error represents a failure to delete a notification policy
// due to mismatched expectation of it being independently managed by chronosphere_notification_policy.
func policyIndependenceDeleteConflict(err error) bool {
	err = clienterror.Wrap(err)

	var clientErr *clienterror.Error
	return errors.As(err, &clientErr) &&
		clientErr.Code == http.StatusBadRequest &&
		strings.Contains(clientErr.Message, "terraform_independent_lifecycle")
}

func resourceBucketBuild(d *schema.ResourceData) (*configmodels.Configv1Bucket, error) {
	b := &intschema.Bucket{}
	if err := b.FromResourceData(d); err != nil {
		return nil, err
	}
	return bucketConverter{}.toModel(b)
}

type bucketConverter struct{}

func (bucketConverter) toModel(
	b *intschema.Bucket,
) (*configmodels.Configv1Bucket, error) {
	if localid.IsLocalID(b.NotificationPolicyId.Slug()) {
		return nil, fmt.Errorf("notification_policy_id must reference a notification policy with name")
	}

	return &configmodels.Configv1Bucket{
		Slug:        b.Slug,
		Name:        b.Name,
		Description: b.Description,
		Labels:      b.Labels,
		TeamSlug:    b.TeamId.Slug(),
		// NB: b.NotificationPolicySlug is used for inline policies internally, so
		// it doesn't map to the API's notification policy field
		NotificationPolicySlug: b.NotificationPolicyId.Slug(),
	}, nil
}

func (bucketConverter) fromModel(
	b *configmodels.Configv1Bucket,
) (*intschema.Bucket, error) {
	return &intschema.Bucket{
		Name:                 b.Name,
		Slug:                 b.Slug,
		Description:          b.Description,
		Labels:               b.Labels,
		TeamId:               tfid.Slug(b.TeamSlug),
		NotificationPolicyId: tfid.Slug(b.NotificationPolicySlug),
	}, nil
}

func newBucketContext(ctx context.Context) context.Context {
	return tfresource.NewContext(ctx, "bucket")
}
