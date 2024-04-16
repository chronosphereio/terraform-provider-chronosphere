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
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/stretchr/testify/require"
)

// Sanity test that the provider schema can be loaded by TF.
func TestProviderLoadSchema(t *testing.T) {
	p := NewProject(t)

	p.Write(t, "main.tf", `
	resource "chronosphere_collection" "c" {
		name = "C"
	}
	`)
	p.Init(t)
	p.Plan(t)
}

// Project represents a Terraform project.
type Project struct {
	dir         string
	providerCfg string
}

// NewProject creates a new Terraform project directory with a basic provider configuration.
func NewProject(t testing.TB) *Project {
	dir, err := os.MkdirTemp(t.TempDir(), "tf-work")
	require.NoError(t, err)

	require.NoError(t, os.WriteFile(filepath.Join(dir, "provider.tf"), []byte(`
	`), 0o666))

	p := &Project{
		dir:         dir,
		providerCfg: startProvider(t),
	}
	p.Write(t, "provider.tf", `
	terraform {
		required_providers {
			chronosphere = {
				version = "0.0.1-dev"
				source  = "local/chronosphereio/chronosphere"
			}
		}
	}

	provider "chronosphere" {
		org = "test"
		api_token = "test"
	}
	`)
	return p
}

// Write adds the given file to the project.
func (p *Project) Write(t testing.TB, name, contents string) {
	filename := filepath.Join(p.dir, name)
	require.NoError(t, os.WriteFile(filename, []byte(contents), 0o666), "write file %q", name)
}

// Init runs "terraform init".
func (p *Project) Init(t testing.TB) {
	require.NoError(t, p.runTF("init"), "terraform init failed")
}

// Plan runs "terraform plan".
func (p *Project) Plan(t testing.TB) {
	require.NoError(t, p.runTF("plan"), "terraform plan failed")
}

func (p *Project) runTF(args ...string) error {
	cmd := exec.Command("terraform", args...)
	cmd.Dir = p.dir
	cmd.Env = append(os.Environ(),
		"TF_REATTACH_PROVIDERS="+string(p.providerCfg),
	)

	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func startProvider(t testing.TB) string {
	// Disable dry-run validation so Plan doesn't make API calls.
	t.Setenv("CHRONOSPHERE_DRY_RUN_VALIDATION_DISABLED", "1")

	// Disable logging, defaults to TRACE which is very noisy.
	t.Setenv("TF_LOG", "ERROR")
	t.Setenv("TF_LOG_SDK_PROTO", "ERROR")
	t.Setenv("TF_LOG_SDK", "ERROR")
	t.Setenv("TF_LOG_PROVIDER_CHRONOSPHERE", "ERROR")

	reattachCh := make(chan *goplugin.ReattachConfig)
	closeCh := make(chan struct{})
	t.Cleanup(func() {
		close(closeCh)
	})

	opts := &plugin.ServeOpts{
		NoLogOutputOverride: true,
		Logger:              hclog.NewNullLogger(),

		ProviderFunc: Provider,
		ProviderAddr: LocalName,
		TestConfig: &goplugin.ServeTestConfig{
			Context:          context.Background(),
			ReattachConfigCh: reattachCh,
			CloseCh:          closeCh,
		},
	}

	reattachCfg, _, err := plugin.DebugServe(context.Background(), opts)
	require.NoError(t, err)
	reattachJSON, err := json.Marshal(map[string]plugin.ReattachConfig{
		LocalName: reattachCfg,
	})
	require.NoError(t, err)
	return string(reattachJSON)
}
