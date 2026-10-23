// Copyright 2026 Chronosphere Inc.
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
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

const cloudflareConnectionAddress = "chronosphere_cloudflare_external_connection.cloudflare"

func TestCloudflareExternalConnectionLifecycle(t *testing.T) {
	p, api := newCloudflareConnectionProject(t)
	p.Write(t, "main.tf", cloudflareConnectionConfig("Cloudflare prod", "cfat_token_v1"))
	p.Init(t)

	step := func(name string, f func(t *testing.T)) {
		if !t.Run(name, f) {
			t.FailNow()
		}
	}

	step("create sends the configured token", func(t *testing.T) {
		p.apply(t)
		assert.Equal(t, []connectionWrite{{http.MethodPost, "cloudflare-prod", "cfat_token_v1"}}, api.takeWrites())
		assert.Equal(t, "cfat_token_v1", p.stateAttr(t, "api_token"))
	})

	step("redacted read plans no changes", func(t *testing.T) {
		assert.Empty(t, p.plannedChanges(t))
	})

	step("repeated applies write nothing", func(t *testing.T) {
		p.apply(t)
		p.apply(t)
		assert.Empty(t, api.takeWrites())
		assert.Equal(t, "cfat_token_v1", p.stateAttr(t, "api_token"))
	})

	step("rotated token updates in place", func(t *testing.T) {
		p.Write(t, "main.tf", cloudflareConnectionConfig("Cloudflare prod", "cfat_token_v2"))
		changes := p.plannedChanges(t)
		require.Contains(t, changes, cloudflareConnectionAddress)
		assert.Equal(t, []string{"update"}, changes[cloudflareConnectionAddress].Actions)

		p.apply(t)
		assert.Equal(t, []connectionWrite{{http.MethodPut, "cloudflare-prod", "cfat_token_v2"}}, api.takeWrites())
		assert.Empty(t, p.plannedChanges(t))
	})

	step("rename resends the configured token", func(t *testing.T) {
		p.Write(t, "main.tf", cloudflareConnectionConfig("Cloudflare production", "cfat_token_v2"))
		p.apply(t)
		assert.Equal(t, []connectionWrite{{http.MethodPut, "cloudflare-prod", "cfat_token_v2"}}, api.takeWrites())
		assert.Equal(t, "cfat_token_v2", api.storedToken("cloudflare-prod"))
		assert.Empty(t, p.plannedChanges(t))
	})

	step("rotating and renaming together sends one update", func(t *testing.T) {
		p.Write(t, "main.tf", cloudflareConnectionConfig("Cloudflare primary", "cfat_token_v3"))
		p.apply(t)
		assert.Equal(t, []connectionWrite{{http.MethodPut, "cloudflare-prod", "cfat_token_v3"}}, api.takeWrites())
		assert.Equal(t, "Cloudflare primary", api.storedName("cloudflare-prod"))
		assert.Empty(t, p.plannedChanges(t))
	})

	step("connection deleted outside Terraform is created again", func(t *testing.T) {
		api.remove("cloudflare-prod")
		changes := p.plannedChanges(t)
		require.Contains(t, changes, cloudflareConnectionAddress)
		assert.Equal(t, []string{"create"}, changes[cloudflareConnectionAddress].Actions)

		p.apply(t)
		assert.Equal(t, []connectionWrite{{http.MethodPost, "cloudflare-prod", "cfat_token_v3"}}, api.takeWrites())
		assert.Empty(t, p.plannedChanges(t))
	})

	step("destroy deletes the connection", func(t *testing.T) {
		require.NoError(t, p.runTF("destroy", "-auto-approve", "-input=false"))
		assert.Equal(t, []connectionWrite{{http.MethodDelete, "cloudflare-prod", ""}}, api.takeWrites())
	})
}

// The API never returns the token, so an import cannot know it; the first apply
// after an import sets it once and later plans are clean.
func TestCloudflareExternalConnectionImport(t *testing.T) {
	p, api := newCloudflareConnectionProject(t)
	api.seed(&models.Configv1ExternalConnection{
		Name:       "Cloudflare prod",
		Slug:       "cloudflare-prod",
		Cloudflare: &models.ExternalConnectionCloudflareConfig{APIToken: "cfat_hub_token"},
	})
	p.Write(t, "main.tf", cloudflareConnectionConfig("Cloudflare prod", "cfat_hub_token"))
	p.Init(t)

	require.NoError(t, p.runTF("import", "-input=false", cloudflareConnectionAddress, "cloudflare-prod"))
	assert.Empty(t, p.stateAttr(t, "api_token"))

	changes := p.plannedChanges(t)
	require.Contains(t, changes, cloudflareConnectionAddress)
	change := changes[cloudflareConnectionAddress]
	assert.Equal(t, []string{"update"}, change.Actions)
	assert.Equal(t, "", change.Before["api_token"])
	assert.Equal(t, "cfat_hub_token", change.After["api_token"])
	assert.Equal(t, change.Before["name"], change.After["name"])

	p.apply(t)
	assert.Equal(t, []connectionWrite{{http.MethodPut, "cloudflare-prod", "cfat_hub_token"}}, api.takeWrites())
	assert.Empty(t, p.plannedChanges(t))
}

func TestCloudflareExternalConnectionTokenUnknownAtPlan(t *testing.T) {
	p, api := newCloudflareConnectionProject(t)
	p.Write(t, "main.tf", `
resource "terraform_data" "token" {
  input = "cfut_generated"
}

resource "chronosphere_cloudflare_external_connection" "cloudflare" {
  name      = "Cloudflare generated"
  api_token = terraform_data.token.output
}
`)
	p.Init(t)

	p.apply(t)
	// Terraform plans the resource again during apply, once the token is known.
	assert.Equal(t, []string{"dry_run_unknown_token", "cfut_generated"}, api.takeDryRunTokens())
	assert.Equal(t, []connectionWrite{{http.MethodPost, "cloudflare-generated", "cfut_generated"}}, api.takeWrites())
	assert.Equal(t, "cloudflare-generated", p.stateAttr(t, "slug"))
	assert.Empty(t, p.plannedChanges(t))
}

func TestCloudflareExternalConnectionNormalize(t *testing.T) {
	tests := []struct {
		name        string
		serverToken string
		wantToken   string
	}{
		{name: "redacted token keeps the configured one", serverToken: "**REDACTED**", wantToken: "cfat_configured"},
		{name: "token returned in full is kept so drift is planned", serverToken: "cfat_from_server", wantToken: "cfat_from_server"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &intschema.CloudflareExternalConnection{Name: "Cloudflare", ApiToken: "cfat_configured"}
			server := &intschema.CloudflareExternalConnection{Name: "Cloudflare", ApiToken: tt.serverToken}

			cloudflareExternalConnectionConverter{}.normalize(config, server)

			assert.Equal(t, tt.wantToken, server.ApiToken)
		})
	}
}

func cloudflareConnectionConfig(name, token string) string {
	return fmt.Sprintf(`
resource "chronosphere_cloudflare_external_connection" "cloudflare" {
  name      = %q
  slug      = "cloudflare-prod"
  api_token = %q
}
`, name, token)
}

func newCloudflareConnectionProject(t *testing.T) (*Project, *fakeConnectionAPI) {
	api := newFakeConnectionAPI()

	// The provider runs in this process and always dials https://<org>.chronosphere.io.
	original := http.DefaultTransport
	http.DefaultTransport = api
	t.Cleanup(func() { http.DefaultTransport = original })

	p := NewProject(t)
	p.Write(t, "provider_config.tf", `
terraform {
  required_providers {
    chronosphere = {
      source = "local/chronosphereio/chronosphere"
    }
  }
}

provider "chronosphere" {
  org            = "test"
  api_token      = "test"
  disable_dryrun = false
}
`)
	return p, api
}

func (p *Project) apply(t testing.TB) {
	require.NoError(t, p.runTF("apply", "-auto-approve", "-input=false"), "terraform apply failed")
}

type plannedChange struct {
	Actions []string       `json:"actions"`
	Before  map[string]any `json:"before"`
	After   map[string]any `json:"after"`
}

func (p *Project) plannedChanges(t testing.TB) map[string]plannedChange {
	require.NoError(t, p.runTF("plan", "-input=false", "-out=tfplan"), "terraform plan failed")
	out := p.output(t, "show", "-json", "tfplan")

	var plan struct {
		ResourceChanges []struct {
			Address string        `json:"address"`
			Change  plannedChange `json:"change"`
		} `json:"resource_changes"`
	}
	require.NoError(t, json.Unmarshal(out, &plan))

	changes := make(map[string]plannedChange)
	for _, rc := range plan.ResourceChanges {
		if len(rc.Change.Actions) == 1 && rc.Change.Actions[0] == "no-op" {
			continue
		}
		changes[rc.Address] = rc.Change
	}
	return changes
}

func (p *Project) stateAttr(t testing.TB, attr string) any {
	var state struct {
		Values struct {
			RootModule struct {
				Resources []struct {
					Address string         `json:"address"`
					Values  map[string]any `json:"values"`
				} `json:"resources"`
			} `json:"root_module"`
		} `json:"values"`
	}
	require.NoError(t, json.Unmarshal(p.output(t, "show", "-json"), &state))
	for _, r := range state.Values.RootModule.Resources {
		if r.Address == cloudflareConnectionAddress {
			return r.Values[attr]
		}
	}
	require.Failf(t, "resource not in state", "%s", cloudflareConnectionAddress)
	return nil
}

func (p *Project) output(t testing.TB, args ...string) []byte {
	cmd := exec.Command("terraform", args...)
	cmd.Dir = p.dir
	cmd.Env = append(os.Environ(), "TF_REATTACH_PROVIDERS="+p.providerCfg)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	require.NoError(t, err, "terraform %s failed", strings.Join(args, " "))
	return out
}

const fakeEntityValidationFailedCode = 400008

type connectionWrite struct {
	method string
	slug   string
	token  string
}

type fakeConnectionAPI struct {
	handler      http.Handler
	mu           sync.Mutex
	connections  map[string]*models.Configv1ExternalConnection
	writes       []connectionWrite
	dryRunTokens []string
}

func newFakeConnectionAPI() *fakeConnectionAPI {
	f := &fakeConnectionAPI{connections: make(map[string]*models.Configv1ExternalConnection)}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/config/external-connections", f.create)
	mux.HandleFunc("GET /api/v1/config/external-connections/{slug}", f.read)
	mux.HandleFunc("PUT /api/v1/config/external-connections/{slug}", f.update)
	mux.HandleFunc("DELETE /api/v1/config/external-connections/{slug}", f.delete)
	f.handler = mux
	return f
}

func (f *fakeConnectionAPI) RoundTrip(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	f.handler.ServeHTTP(rec, req)
	return rec.Result(), nil
}

func (f *fakeConnectionAPI) seed(c *models.Configv1ExternalConnection) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.connections[c.Slug] = c
}

func (f *fakeConnectionAPI) takeWrites() []connectionWrite {
	f.mu.Lock()
	defer f.mu.Unlock()
	writes := f.writes
	f.writes = nil
	return writes
}

func (f *fakeConnectionAPI) takeDryRunTokens() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	tokens := f.dryRunTokens
	f.dryRunTokens = nil
	return tokens
}

func (f *fakeConnectionAPI) storedToken(slug string) string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.connections[slug].Cloudflare.APIToken
}

func (f *fakeConnectionAPI) storedName(slug string) string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.connections[slug].Name
}

func (f *fakeConnectionAPI) remove(slug string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.connections, slug)
}

func (f *fakeConnectionAPI) create(w http.ResponseWriter, r *http.Request) {
	var req models.Configv1CreateExternalConnectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeFakeAPIError(w, http.StatusBadRequest, 0, err.Error())
		return
	}
	c := req.ExternalConnection
	if msg := invalidCloudflareConnection(c); msg != "" {
		writeFakeAPIError(w, http.StatusBadRequest, fakeEntityValidationFailedCode, msg)
		return
	}
	if c.Slug == "" {
		c.Slug = strings.ToLower(strings.ReplaceAll(c.Name, " ", "-"))
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	if req.DryRun {
		f.dryRunTokens = append(f.dryRunTokens, c.Cloudflare.APIToken)
		writeFakeConnection(w, c)
		return
	}
	if _, ok := f.connections[c.Slug]; ok {
		writeFakeAPIError(w, http.StatusConflict, 0, "external connection already exists")
		return
	}
	f.connections[c.Slug] = c
	f.writes = append(f.writes, connectionWrite{http.MethodPost, c.Slug, c.Cloudflare.APIToken})
	writeFakeConnection(w, c)
}

func (f *fakeConnectionAPI) read(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	c, ok := f.connections[r.PathValue("slug")]
	if !ok {
		writeFakeAPIError(w, http.StatusNotFound, 0, "external connection not found")
		return
	}
	writeFakeConnection(w, c)
}

func (f *fakeConnectionAPI) update(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	var body models.ConfigV1UpdateExternalConnectionBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeFakeAPIError(w, http.StatusBadRequest, 0, err.Error())
		return
	}
	c := body.ExternalConnection
	if c == nil || c.Slug != slug {
		writeFakeAPIError(w, http.StatusBadRequest, 0, "slug in body does not match path")
		return
	}
	if msg := invalidCloudflareConnection(c); msg != "" {
		writeFakeAPIError(w, http.StatusBadRequest, fakeEntityValidationFailedCode, msg)
		return
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	existing, ok := f.connections[slug]
	if !ok && !body.CreateIfMissing {
		writeFakeAPIError(w, http.StatusNotFound, 0, "external connection not found")
		return
	}
	// Like the real server, a redacted token on update keeps the stored one.
	if ok && c.Cloudflare.APIToken == notifierRedactedSecret {
		c.Cloudflare.APIToken = existing.Cloudflare.APIToken
	}
	if body.DryRun {
		f.dryRunTokens = append(f.dryRunTokens, c.Cloudflare.APIToken)
		writeFakeConnection(w, c)
		return
	}
	f.connections[slug] = c
	f.writes = append(f.writes, connectionWrite{http.MethodPut, slug, c.Cloudflare.APIToken})
	writeFakeConnection(w, c)
}

func (f *fakeConnectionAPI) delete(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.connections[slug]; !ok {
		writeFakeAPIError(w, http.StatusNotFound, 0, "external connection not found")
		return
	}
	delete(f.connections, slug)
	f.writes = append(f.writes, connectionWrite{method: http.MethodDelete, slug: slug})
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{}"))
}

func invalidCloudflareConnection(c *models.Configv1ExternalConnection) string {
	switch {
	case c == nil:
		return "external_connection: value required"
	case c.Name == "":
		return "external_connection.name: value required"
	case c.Cloudflare == nil || c.Cloudflare.APIToken == "":
		return "external_connection.cloudflare.api_token: value required"
	}
	return ""
}

func writeFakeConnection(w http.ResponseWriter, c *models.Configv1ExternalConnection) {
	redacted := *c
	redacted.Cloudflare = &models.ExternalConnectionCloudflareConfig{APIToken: notifierRedactedSecret}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"external_connection": &redacted})
}

func writeFakeAPIError(w http.ResponseWriter, status int, code int32, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(&models.APIError{Code: code, Message: msg})
}
