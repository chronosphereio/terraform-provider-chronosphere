DEV_PROVIDER_VERSION = 0.0.1-dev
BINARY=terraform-provider-chronosphere_v$(PROVIDER_VERSION)
TESTPKGS?=$$(go list ./... | grep -v 'vendor' | grep -v 'scenario')
TOOLS_BIN=$(abspath ./bin)
OS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(shell uname -m | sed "s/x86_64/amd64/" | sed "s/aarch64/arm64/")
OS_ARCH=${OS}_${ARCH}
SHELL=/bin/bash -o pipefail

SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
REPO_DIR := $(abspath $(SELF_DIR))

# LD Flags
GIT_REVISION              := $(shell git rev-parse --short HEAD)
GIT_BRANCH                := $(shell git rev-parse --abbrev-ref HEAD)
GIT_VERSION               := $(shell git describe --tags 2>/dev/null | egrep '^v.+' || echo unknown)
BUILD_DATE                := $(shell date -u  +"%Y-%m-%dT%H:%M:%SZ") # Use RFC-3339 date format
BUILD_TS_UNIX             := $(shell date '+%s') # second since epoch
INSTRUMENT_PACKAGE        := github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/buildvar

# FIXME(codyg): Our Go binaries are built with -buildvcs=false due to some
# strange buildkite issue where Go would detect a VCS but then fail to retrieve
# any information: "error obtaining VCS status: exit status 128". We gave up on
# root causing the issue to unblock the Go 1.18 upgrade.
GOFLAGS=-buildvcs=false

GO_GENERATE=PATH=$(TOOLS_BIN):$(PATH) go generate

BUILD                     := $(abspath ./bin)

GO_RELEASER_DOCKER_IMAGE  := gcr.io/chronosphere-dev/goreleaser:v1.21.2-gpg-agent
GO_RELEASER_RELEASE_ARGS  ?= --rm-dist
GO_RELEASER_WORKING_DIR   := /go/src/github.com/chronosphere/terraform-provider-chronosphere

INTERNAL_TOOLS := \
  chronosphere/generateresources \
  chronosphere/pagination/generatepagination \
  chronosphere/intschema/generateintschema

GO_TEST   := $(or $(GO_TEST),go test)
SCENARIOS := $(or $(SCENARIOS),./scenarios/...)

SHARD_SCENARIOS = go run ./cmd/test_sharder -num-shards 2 -salt "8"

# START_RULES general
.PHONY: setup
setup:
	mkdir -p $(BUILD)

.PHONY: install-tools
install-tools:
	mkdir -p bin
	cd tools && cat tools.go | grep _ | awk -F'"' '{print $$2}' | GOBIN=$(TOOLS_BIN) xargs -tI % go install $(GOFLAGS) %
	$(foreach path, $(INTERNAL_TOOLS), \
	  GOBIN=$(TOOLS_BIN) go install $(GOFLAGS) github.com/chronosphereio/terraform-provider-chronosphere/$(path); )

.PHONY: generate
generate: install-tools
	$(GO_GENERATE) ./...

# subset of generate that skips swagger since swagger is slow to generate.
.PHONY: generate-no-swagger
generate-no-swagger: install-tools
	$(GO_GENERATE) ./chronosphere ./chronosphere.intschema ./chronosphere/pagination

.PHONY: test-generate
test-generate: generate
	@echo "--- :golang: go mod tidy"
	go mod tidy
	@echo "--- :git: Testing generate code is up to date"
	@[ -z "$$(git status --porcelain)" ] || ((set -x; git status --porcelain; git diff); echo -e "^^^ +++\nCheck git status + diff above, there are changed or untracked files"; exit 1)

.PHONY: intschema
intschema: install-tools
	$(GO_GENERATE) ./chronosphere/intschema

.PHONY: generatedresources
generatedresources: install-tools
	$(GO_GENERATE) ./chronosphere

.PHONY: pagination
pagination: install-tools
	$(GO_GENERATE) ./chronosphere/pagination

.PHONY: update-swagger
update-swagger: install-tools
	rm -rf chronosphere/pkg/configunstable/{client,models}
	rm -rf chronosphere/pkg/configv1/{client,models}

	$(GO_GENERATE) ./chronosphere/pkg/configunstable/...
	$(GO_GENERATE) ./chronosphere/pkg/configv1/...

.PHONY: debug
debug:
	go run  . -debug

.PHONY: build
build:
	@echo building with git version ${GIT_VERSION}
	mkdir -p bin
	go build $(GOFLAGS) -o bin/${BINARY} -ldflags "-X ${INSTRUMENT_PACKAGE}.Revision=${GIT_REVISION} -X ${INSTRUMENT_PACKAGE}.Version=${GIT_VERSION}"

.PHONY: install
install: build verify-terraform-arch
	# Copy the provider to the local terraform plugin directory for terraform <= 0.12
	mkdir -p ~/.terraform.d/plugins/${OS_ARCH}
	cp bin/${BINARY} ~/.terraform.d/plugins/${OS_ARCH}
	# Copy the provider to the local terraform plugin directory for terraform >= 0.13
	mkdir -p ~/.terraform.d/plugins/local/chronosphereio/chronosphere/$(DEV_PROVIDER_VERSION)/${OS_ARCH}/
	mv bin/${BINARY} ~/.terraform.d/plugins/local/chronosphereio/chronosphere/$(DEV_PROVIDER_VERSION)/${OS_ARCH}

.PHONY: verify-terraform-arch
verify-terraform-arch:
	@terraform -version | grep ${OS_ARCH} >/dev/null || (echo "The terraform binary doesn't match OS/architecture ${OS_ARCH}. See http://go/terraform-arch for details on how to resolve this."; exit 1)

.PHONY: test
test:
	$(GO_TEST) $(TESTARGS) -v -timeout=30s -parallel=4 $(TESTPKGS)

.PHONY: test-ci-old
test-ci-old:
	make TESTARGS="-json" test | tee test.log | go run ./scripts/json2test

.PHONY: test-ci
test-ci: install-tools
	make GO_TEST=./scripts/gotestsum.sh test | tee test.log | go run ./scripts/json2test

.PHONY: version
version:
	@echo "git: $(GIT_VERSION)"
	@echo "dev: $(DEV_PROVIDER_VERSION)"

# release publishes release artifacts
.PHONY: release
release:
ifneq ($(SKIP_RELEASE_BRANCH_VALIDATION),true)
	@[ ! -z "${RELEASE_VERSION}" ] || (echo "Set RELEASE_VERSION to the expected release version (e.g., ${GIT_VERSION})"; exit 1)
	@[ "${GIT_VERSION}" = "${RELEASE_VERSION}" ] || (echo "RELEASE_VERSION (${RELEASE_VERSION}) does not match checked out git version (${GIT_VERSION})"; exit 1)
	@[ "$$(git status --porcelain)" = "" ] || (echo "Must release from a clean git repo"; exit 1)
endif
	@echo Releasing new version
	GIT_REVISION=$(GIT_REVISION) \
		GIT_BRANCH=$(GIT_BRANCH) \
		GIT_VERSION=$(GIT_VERSION) \
		BUILD_DATE=$(BUILD_DATE) \
		BUILD_TS_UNIX=$(BUILD_TS_UNIX) \
		INSTRUMENT_PACKAGE=$(INSTRUMENT_PACKAGE) \
		GO_RELEASER_DOCKER_IMAGE=$(GO_RELEASER_DOCKER_IMAGE) \
		GO_RELEASER_RELEASE_ARGS="$(GO_RELEASER_RELEASE_ARGS)" \
		GO_RELEASER_WORKING_DIR=$(GO_RELEASER_WORKING_DIR) \
		SSH_AUTH_SOCK=$(SSH_AUTH_SOCK) \
		./scripts/run_goreleaser.sh ${GO_RELEASER_RELEASE_ARGS}
	GIT_VERSION=$(GIT_VERSION) ./scripts/copy_release.py

.PHONY: release-snapshot
release-snapshot:
	@echo Building binaries with goreleaser
	# --snapshot mode allows building artifacts w/o release tag present and w/ publishing mode disabled
	# useful when we want to test whether we can build binaries, but not publish yet.
	make release SKIP_RELEASE_BRANCH_VALIDATION=true GO_RELEASER_RELEASE_ARGS="--snapshot --rm-dist --skip-publish --skip-sign"

.PHONY: fmt
fmt:
	go fmt $(shell go list ./...)
	terraform fmt examples/

.PHONY: lint
lint: install-tools
	@echo "--- :golang: linting code"
	GOFLAGS=$(GOFLAGS) $(TOOLS_BIN)/golangci-lint run -D gosec -D unused -E gofmt -E goimports -E gofumpt
