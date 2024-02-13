package chronosphere

import (
	"context"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/service_account"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ServiceAccountFromModel converts a service account to the intschema representation
func ServiceAccountFromModel(m *models.Configv1ServiceAccount) (*intschema.ServiceAccount, error) {
	return serviceAccountConverter{}.fromModel(m), nil
}

// NOTE: Converting the service account resource to use genericResource has several issues:
//   - normally we don't set anything as part of create other than the ID, so the main issue
//     is that we need to set token as part of create, but genericResource only sets ID on create
//     object in the create response
//   - the token is not returned in the read response, so we may unintentionally unset the value
//   - service accounts dont support udpate
func resourceServiceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceAccountCreate,
		ReadContext:   resourceServiceAccountRead,
		DeleteContext: resourceServiceAccountDelete,
		Schema:        tfschema.ServiceAccount,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceServiceAccountCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "service_account")
	cli := getConfigClient(meta)

	serviceAccount, err := resourceServiceAccountBuild(d)
	if err != nil {
		return diag.Errorf("unable to build service account: %v", err)
	}

	resp, err := cli.ServiceAccount.CreateServiceAccount(&service_account.CreateServiceAccountParams{
		Context: ctx,
		Body: &models.Configv1CreateServiceAccountRequest{
			ServiceAccount: serviceAccount,
		},
	})
	if err != nil {
		return diag.Errorf("unable to create service account: %v", clienterror.Wrap(err))
	}

	// We set the token here instead of waiting for the subsequent read because the token
	// is only ever returned in the create response.
	err = d.Set("token", resp.Payload.ServiceAccount.Token)
	if err != nil {
		return diag.Errorf("unable to set service account token: %v", clienterror.Wrap(err))
	}

	d.SetId(resp.Payload.ServiceAccount.Slug)

	return nil
}

func resourceServiceAccountRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "service_account")
	cli := getConfigClient(meta)
	resp, err := cli.ServiceAccount.ReadServiceAccount(&service_account.ReadServiceAccountParams{
		Context: ctx,
		Slug:    d.Id(),
	})
	if err != nil {
		if clienterror.IsNotFound(err) {
			setResourceNotFound(d)
			return nil
		}
		return diag.Errorf("unable to read service account: %v", clienterror.Wrap(err))
	}

	serviceAccount := resp.Payload.ServiceAccount
	// The token does not come back from the server when we read a service account.
	// Therefore we have to set it to the value we already have in the tf state.
	serviceAccount.Token = d.Get("token").(string)
	o := serviceAccountConverter{}.fromModel(serviceAccount)
	if err := o.ToResourceData(d); err != nil {
		return err
	}

	d.SetId(serviceAccount.Slug)

	return nil
}

// resourceServiceAccountDelete deletes a service account.
func resourceServiceAccountDelete(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, "service_account")
	cli := getConfigClient(meta)

	if _, err := cli.ServiceAccount.DeleteServiceAccount(&service_account.DeleteServiceAccountParams{
		Slug:    d.Id(),
		Context: ctx,
	}); clienterror.IsNotFound(err) {
		// Already deleted on the server, treat as success.
	} else if err != nil {
		return diag.Errorf("unable to delete service account: %v", clienterror.Wrap(err))
	}

	d.SetId("")

	return nil
}

func resourceServiceAccountBuild(d ResourceGetter) (*models.Configv1ServiceAccount, error) {
	m := &intschema.ServiceAccount{}
	if err := m.FromResourceData(d); err != nil {
		return nil, err
	}

	sa, err := serviceAccountConverter{}.toModel(m)
	if err != nil {
		return nil, err
	}

	return sa, nil
}

type serviceAccountConverter struct{}

func (serviceAccountConverter) toModel(s *intschema.ServiceAccount) (*models.Configv1ServiceAccount, error) {
	m := &models.Configv1ServiceAccount{
		Name:         s.Name,
		Slug:         s.Slug,
		Email:        s.Email,
		Unrestricted: s.Unrestricted,
	}

	if s.Restriction != nil {
		restriction := &models.ServiceAccountMetricsRestriction{
			Permission: enum.Permission.V1(s.Restriction.Permission),
			Labels:     s.Restriction.Labels,
		}

		m.MetricsRestriction = restriction
	}

	return m, nil
}

func (serviceAccountConverter) fromModel(m *models.Configv1ServiceAccount) *intschema.ServiceAccount {
	o := &intschema.ServiceAccount{
		Name:         m.Name,
		Slug:         m.Slug,
		Token:        m.Token,
		Email:        m.Email,
		Unrestricted: m.Unrestricted,
	}

	if perm := m.MetricsRestriction; perm != nil {
		o.Restriction = &intschema.ServiceAccountRestriction{
			Permission: string(perm.Permission),
			Labels:     perm.Labels,
		}
	}
	return o
}
