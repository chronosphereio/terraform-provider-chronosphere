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
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/convertintschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/apiclients"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/clienterror"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/tfresource"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/shared/pkg/container/set"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// internalSchema is implemented by all intschema types.
type internalSchema interface {
	FromResourceData(d convertintschema.ResourceGetter) error
	ToResourceData(d *schema.ResourceData) diag.Diagnostics
}

// internalSchemaPtr defines a pointer to an intschema struct value, where
// the pointer implements internalSchema.
type internalSchemaPtr[SV any] interface {
	*SV
	internalSchema
}

// newInternalSchema creates a new pointer type of SV, where *SV implements
// internalSchema.
func newInternalSchema[SV any, S internalSchemaPtr[SV]]() S {
	var v SV
	return S(&v)
}

// modelConverter provides a two-way mapping of a internal Terraform schema to
// its respective Swagger model.
//
// This is intended to be implemented manually.
type modelConverter[M any, S internalSchema] interface {
	// toModel maps a Terraform resource into a Swagger model.
	toModel(S) (M, error)

	// fromModel maps a Swagger model into a Terraform resource.
	fromModel(M) (S, error)
}

// resourceCRUD provides CRUD methods against a Swagger model.
//
// This is intended to be implemented by generated code (see tools/generateresources).
type resourceCRUD[M any] interface {
	create(ctx context.Context, clients apiclients.Clients, m M, dryRun bool) (slug string, err error)
	read(ctx context.Context, clients apiclients.Clients, slug string) (M, error)
	update(ctx context.Context, clients apiclients.Clients, m M, params updateParams) error
	delete(ctx context.Context, clients apiclients.Clients, slug string) error
}

type updateParams struct {
	createIfMissing bool
	dryRun          bool
}

// resourceFields provides field getters against a Swagger model, as Swagger
// does not generate any getters.
//
// This is intended to be implemented by generated code (see tools/generateresources).
type resourceFields[M any] interface {
	slugOf(M) string
}

// generatedResource joins the interfaces implemented by tools/generateresources.
type generatedResource[M any] interface {
	resourceCRUD[M]
	resourceFields[M]
}

// genericResource implements Terraform CRUD bindings for any resource.
type genericResource[M any, SV any, S internalSchemaPtr[SV]] struct {
	name      string
	converter modelConverter[M, S]
	crud      resourceCRUD[M]
	fields    resourceFields[M]
}

func newGenericResource[M any, SV any, S internalSchemaPtr[SV]](
	name string,
	c modelConverter[M, S],
	g generatedResource[M],
) genericResource[M, SV, S] {
	return genericResource[M, SV, S]{
		name:      name,
		converter: c,
		crud:      g,
		fields:    g,
	}
}

// CreateContext implements schema.CreateContextFunc.
func (r genericResource[M, SV, S]) CreateContext(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	ctx = tfresource.NewContext(ctx, r.name)

	clients := meta.(apiclients.Clients)

	s := newInternalSchema[SV, S]()
	if err := s.FromResourceData(d); err != nil {
		return diag.Errorf(err.Error())
	}

	m, err := r.converter.toModel(s)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	slug, err := r.crud.create(ctx, clients, m, false /* dryRun */)
	if err != nil {
		return diag.Errorf("unable to create %s: %v", r.name, clienterror.Wrap(err))
	}
	d.SetId(slug)

	return nil
}

// ReadContext implements schema.ReadContextFunc.
func (r genericResource[M, SV, S]) ReadContext(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	clients := meta.(apiclients.Clients)

	m, err := r.crud.read(ctx, clients, d.Id())
	if err != nil {
		if clienterror.IsNotFound(err) {
			setResourceNotFound(d)
			return nil
		}
		return diag.Errorf("unable to read %s: %v", r.name, clienterror.Wrap(err))
	}

	d.SetId(r.fields.slugOf(m))

	s, err := r.converter.fromModel(m)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	type normalizer interface {
		normalize(schemaCfg, serverCfg S)
	}
	if nr, ok := r.converter.(normalizer); ok {
		// Normalize the server-read value against the value set in config.
		configured := newInternalSchema[SV, S]()
		if err := configured.FromResourceData(d); err != nil {
			return diag.Errorf("cannot read config from schema: %v", err)
		}
		nr.normalize(configured, s)
	}

	return s.ToResourceData(d)
}

// UpdateContext implements schema.UpdateContextFunc.
func (r genericResource[M, SV, S]) UpdateContext(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	clients := meta.(apiclients.Clients)

	s := newInternalSchema[SV, S]()
	if err := s.FromResourceData(d); err != nil {
		return diag.Errorf(err.Error())
	}

	m, err := r.converter.toModel(s)
	if err != nil {
		return diag.Errorf(err.Error())
	}
	slug := r.fields.slugOf(m)
	if slug != d.Id() {
		return diag.Errorf(
			"cannot change slug of existing %s: current slug %q, new slug %q",
			r.name, d.Id(), slug)
	}

	if err := r.crud.update(ctx, clients, m, updateParams{}); err != nil {
		return diag.Errorf("unable to update %s: %v", r.name, clienterror.Wrap(err))
	}

	return nil
}

// DeleteContext implements schema.DeleteContextFunc.
func (r genericResource[M, SV, S]) DeleteContext(
	ctx context.Context, d *schema.ResourceData, meta any,
) diag.Diagnostics {
	clients := meta.(apiclients.Clients)

	if err := r.crud.delete(ctx, clients, d.Id()); clienterror.IsNotFound(err) {
		// Already deleted on the server, treat as success.
	} else if err != nil {
		return diag.Errorf("unable to delete %s: %v", r.name, clienterror.Wrap(err))
	}
	d.SetId("")
	return nil
}

// ValidateDryRun implements schema.CustomizeDiffFunc, and performs a dry-run validation.
func (r genericResource[M, SV, S]) ValidateDryRun(dryRunCounter *atomic.Int64) schema.CustomizeDiffFunc {
	return r.ValidateDryRunOptions(dryRunCounter, ValidateDryRunOpts[M]{})
}

// ValidateDryRunOpts is optional parameters for ValidateDryRun.
type ValidateDryRunOpts[M any] struct {
	// SetUnknownReferencesSkip is a set of fields to skip when setting unknown references.
	SetUnknownReferencesSkip []string

	// DryRunDefaults are default values to set during dry-run if the field has a value in config
	// but not in the "plan" ResourceData. This typically happens due to unknown values.
	// E.g., interpolating an ID field for a resource that hasn't been created yet.
	DryRunDefaults map[string]any

	// ModifyAPIModel is used to modify an API model before making the dry-run API call.
	ModifyAPIModel func(M)
}

// ValidateDryRunOptions is the same as ValidateDryRun but supports additional options, see ValidateDryRunOpts.
func (r genericResource[M, SV, S]) ValidateDryRunOptions(dryRunCounter *atomic.Int64, opts ValidateDryRunOpts[M]) schema.CustomizeDiffFunc {
	return func(ctx context.Context, diff *schema.ResourceDiff, meta any) error {
		clients := meta.(apiclients.Clients)

		if clients.DisableDryRun || skipDryRun(diff) {
			return nil
		}
		// Increment dry run count for testing.
		dryRunCounter.Add(1)

		logParams := map[string]any{
			"resourceType": r.name,
			"id":           diff.Id(),
		}
		if slug, ok := diff.Get("slug").(string); ok {
			logParams["slug"] = slug
		}
		tflog.Info(ctx, "running dry run validation", logParams)

		s := newInternalSchema[SV, S]()
		if err := s.FromResourceData(diff); err != nil {
			return err
		}

		setUnknown(s, setUnknownParams{
			rawConfig:      diff.GetRawConfig(),
			skipIDs:        set.New(opts.SetUnknownReferencesSkip...),
			dryRunDefaults: opts.DryRunDefaults,
		})

		m, err := r.converter.toModel(s)
		if err != nil {
			return err
		}

		if opts.ModifyAPIModel != nil {
			opts.ModifyAPIModel(m)
		}

		if diff.Id() == "" {
			_, err = r.crud.create(ctx, clients, m, true /* dryRun */)
		} else {
			err = r.crud.update(ctx, clients, m, updateParams{dryRun: true})
		}
		if err != nil && clienterror.IsEntityValidationFailed(err) {
			return fmt.Errorf("dry run validation failed: %w", err)
		}
		return nil
	}
}
