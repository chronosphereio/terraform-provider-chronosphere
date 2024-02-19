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

package overridecreate

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

var (
	ErrExistsOutsideTerraform = errors.New("entity already exists on server and is not managed in " +
		"terraform state. set the override_create property to import it or this request will fail")
	ErrDoesNotExistOutsideTerraform = errors.New("override_create set to import_or_fail. no resource present " +
		"to import from server")
)

// Field is the Terraform field name which, if included in a schema,
// will enable override_create handling.
const Field = "override_create"

// Action indicates what create action to take.
type Action int

const (
	AlwaysCreate Action = iota + 1
	ImportIfExists
	ImportOrFail
)

var actionToString = map[Action]string{
	AlwaysCreate: "", // default
	// ImportIfExists When Terraform attempts to create the resource, we check for any
	// existing singleton resources on the server. If one is found, it will be imported into
	// terraform state and updated with the config defined in terraform. Otherwise, the Create API
	// call will be used.
	ImportIfExists: "import_if_exists",
	// ImportOrFail is an option for the override_create flag.
	// When Terraform attempts to create the resource, instead of the Create API call, we look for
	// an existing singleton resource. If one is found, it will be imported into terraform state
	// and updated with the config defined in terraform. If not, the call will fail, failing the
	// Terraform apply.
	ImportOrFail: "import_or_fail",
}

var actionFromString = invertMap(actionToString)

func (a Action) String() string {
	return actionToString[a]
}

// Schema returns the Terraform schema of the OverrideCreate field.
func Schema() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		ValidateDiagFunc: validate,
		Optional:         true,
	}
}

func validate(v any, _ cty.Path) diag.Diagnostics {
	s, ok := v.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}

	_, err := ParseAction(s)
	if err != nil {
		return diag.Errorf("%v", err)
	}
	return nil
}

func ParseAction(input string) (Action, error) {
	action, ok := actionFromString[input]
	if !ok {
		return 0, fmt.Errorf("invalid override_create value: %q; must be one of %v", input, maps.Keys(actionFromString))
	}
	return action, nil
}

func invertMap(m map[Action]string) map[string]Action {
	n := make(map[string]Action, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

type existsOnServerFunc func(context.Context, string, interface{}) (bool, error)

// CustomizeDiff defines custom diffs for resources using the override_create property to encourage proper
// use of singleton resources in TF. `additionalLogFields` can optionally be supplied to output additional
// resource-specific fields in logs triggered by this operation - minimally, the values  of `override_create` and the
// provided resourceIDField will always be included in logs, even if no additional fields are specified with this
// parameter.
func CustomizeDiff(
	resourceIDField string,
	existsOnServer existsOnServerFunc,
	additionalLogFields ...string,
) func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
		if d.Id() != "" {
			// already in terraform state.
			return nil
		}
		// Not in terraform state.

		// NB: It's ok if these are accidentally duplicated by the value passed in by `additionalLogFields` as these
		// will become map keys. Set them here to ensure they are included.
		additionalLogFields = append(additionalLogFields, Field)
		additionalLogFields = append(additionalLogFields, resourceIDField)

		resourceFieldMap := make(map[string]any, len(additionalLogFields))
		for _, key := range additionalLogFields {
			if value, ok := d.Get(key).(string); ok {
				resourceFieldMap[key] = value
			}
		}

		overrideCreateAction := AlwaysCreate
		if overrideCreate, ok := resourceFieldMap[Field].(string); ok {
			var err error
			overrideCreateAction, err = ParseAction(overrideCreate)
			if err != nil {
				return err
			}
		}

		resourceID := ""
		if r, ok := resourceFieldMap[resourceIDField].(string); ok {
			resourceID = r
		}

		switch overrideCreateAction {
		case ImportOrFail:
			tflog.Warn(
				ctx, "override_create set. will import state from the server and update it, or fail if none exists",
				resourceFieldMap,
			)

			if exists, err := existsOnServer(ctx, resourceID, meta); err == nil && !exists {
				return fmt.Errorf(ErrDoesNotExistOutsideTerraform.Error())
			}
		case ImportIfExists:
			tflog.Warn(
				ctx,
				"override_create set. If server already has corresponding entity not yet managed "+
					"by terraform state, read it and update it. Otherwise if it doesn't exist, create it",
				resourceFieldMap,
			)
		default:
			// Best effort to see if a there's already an entity on the server. If there is, then fail with a
			// message about using override create to import it.
			if exists, err := existsOnServer(ctx, resourceID, meta); err == nil && exists {
				return fmt.Errorf(ErrExistsOutsideTerraform.Error())
			}
		}
		return nil
	}
}
