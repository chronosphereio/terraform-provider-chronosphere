DEV_PROVIDER_VERSION = 0.0.1-dev
BINARY=terraform-provider-chronosphere_v$(DEV_PROVIDER_VERSION)
TESTPKGS?=$$(go list ./... | grep -v 'vendor' | grep -v 'scenario')
TOOLS_BIN=$(abspath ./bin)
OS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(shell uname -m | sed "s/x86_64/amd64/" | sed "s/aarch64/arm64/")
OS_ARCH=${OS}_${ARCH}
SHELL=/bin/bash -o pipefail

SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
REPO_DIR := $(abspath $(SELF_DIR))

BUILD                     := $(abspath ./bin)
GO_BUILD_LDFLAGS_CMD      := $(abspath ./scripts/go-build-ldflags.sh)
GO_BUILD_LDFLAGS          := $(shell $(GO_BUILD_LDFLAGS_CMD))
GO_BUILD_COMMON_ENV       := CGO_ENABLED=0
# TODO: this was causing issues with releases. May not be a problem in recent go versions
GOFLAGS=-buildvcs=false
INSTRUMENT_PACKAGE        := github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/buildinfo
GIT_VERSION               := $(shell git describe --tags --abbrev=0 2>/dev/null || echo unknown)

GO_GENERATE=PATH=$(TOOLS_BIN):$(PATH) go generate

GO_RELEASER_RELEASE_ARGS  ?= --clean
GO_RELEASER_WORKING_DIR   := /go/src/github.com/chronosphere/terraform-provider-chronosphere

SNAPSHOT_VERSION=$(shell ./scripts/next_version.sh || echo unknown)
SNAPSHOT_VERSION_NUMBER=$(subst v,,$(SNAPSHOT_VERSION))
LATEST_GIT_COMMIT=$(shell git rev-parse --short HEAD)
SNAPSHOT_BINARY=terraform-provider-chronosphere_${SNAPSHOT_VERSION}-SNAPSHOT-${LATEST_GIT_COMMIT}
LOCAL_SNAPSHOT_VERSION_DIR=${SNAPSHOT_VERSION_NUMBER}-SNAPSHOT-${LATEST_GIT_COMMIT}
SNAPSHOT_GO_BUILD_LDFLAGS_CMD      := $(abspath ./scripts/go-build-ldflags.sh --snapshot)
SNAPSHOT_GO_BUILD_LDFLAGS          := $(shell $(GO_BUILD_LDFLAGS_CMD))

INTERNAL_TOOLS := \
  chronosphere/generateresources \
  chronosphere/pagination/generatepagination \
  chronosphere/intschema/generateintschema

GO_TEST   := $(or $(GO_TEST),go test)

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
	$(GO_GENERATE) ./chronosphere ./chronosphere/intschema ./chronosphere/pagination

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
	@[ -n "${SWAGGER_PATH}" ] || (echo "SWAGGER_PATH must be set, please set it and rerun this command"; exit 1)
	rm -rf chronosphere/pkg/configunstable/{client,models}
	cp ${SWAGGER_PATH}/unstable_config_swagger.json chronosphere/pkg/configunstable/swagger.json

	rm -rf chronosphere/pkg/configv1/{client,models}
	cp ${SWAGGER_PATH}/v1_config_swagger.json chronosphere/pkg/configv1/swagger.json

	$(GO_GENERATE) ./chronosphere/pkg/configunstable/...
	$(GO_GENERATE) ./chronosphere/pkg/configv1/...

.PHONY: debug
debug:
	go run  . -debug

.PHONY: build
build:
	@echo building with git version ${GIT_VERSION}
	mkdir -p bin
	$(GO_BUILD_COMMON_ENV) go build $(GOFLAGS) -ldflags '$(GO_BUILD_LDFLAGS)' -o $(BUILD)/${BINARY}

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
	@terraform -version | grep ${OS_ARCH} >/dev/null || (echo "The terraform binary doesn't match OS/architecture ${OS_ARCH}"; exit 1)

.PHONY: test
test:
	$(GO_TEST) $(TESTARGS) -v -timeout=30s -parallel=4 $(TESTPKGS)

.PHONY: test-ci
test-ci: install-tools
	make GO_TEST=./scripts/gotestsum.sh test | tee test.log | go run ./scripts/json2test

.PHONY: version
version:
	@echo "git: $(GIT_VERSION)"
	@echo "dev: $(DEV_PROVIDER_VERSION)"
	@echo "snapshot: $(shell ./scripts/next_version.sh || echo unknown)-SNAPSHOT"

# release publishes release artifacts
.PHONY: release
release:
ifneq ($(SKIP_RELEASE_BRANCH_VALIDATION),true)
	@[ ! -z "${RELEASE_VERSION}" ] || (echo "Set RELEASE_VERSION to the expected release version (e.g., ${GIT_VERSION})"; exit 1)
	@[ "${GIT_VERSION}" = "${RELEASE_VERSION}" ] || (echo "RELEASE_VERSION (${RELEASE_VERSION}) does not match checked out git version (${GIT_VERSION})"; exit 1)
	@[ "$$(git status --porcelain)" = "" ] || (echo "Must release from a clean git repo"; exit 1)
endif
	@echo Releasing new version
	GO_BUILD_LDFLAGS="$(GO_BUILD_LDFLAGS)" \
		INSTRUMENT_PACKAGE=$(INSTRUMENT_PACKAGE) \
		GO_RELEASER_DOCKER_IMAGE=$(GO_RELEASER_DOCKER_IMAGE) \
		GO_RELEASER_RELEASE_ARGS="$(GO_RELEASER_RELEASE_ARGS)" \
		GO_RELEASER_WORKING_DIR=$(GO_RELEASER_WORKING_DIR) \
		SSH_AUTH_SOCK=$(SSH_AUTH_SOCK) \
		GIT_VERSION=$(GIT_VERSION) \
		./scripts/run_goreleaser.sh ${GO_RELEASER_RELEASE_ARGS}

.PHONY: snapshot
.IGNORE: snapshot # ignore build errors to make sure the git tag is removed.
snapshot:
	@echo "Building snapshot version ${SNAPSHOT_VERSION}-SNAPSHOT with goreleaser"
	git tag ${SNAPSHOT_VERSION}
	# --snapshot mode allows building artifacts w/o release tag present and w/ publishing mode disabled
	# useful when we want to test whether we can build binaries, but not publish yet.
	GO_BUILD_LDFLAGS="$(GO_BUILD_LDFLAGS)" \
		INSTRUMENT_PACKAGE=$(INSTRUMENT_PACKAGE) \
		GO_RELEASER_DOCKER_IMAGE=$(GO_RELEASER_DOCKER_IMAGE) \
		GO_RELEASER_RELEASE_ARGS="$(GO_RELEASER_RELEASE_ARGS)" \
		GO_RELEASER_WORKING_DIR=$(GO_RELEASER_WORKING_DIR) \
		SSH_AUTH_SOCK=$(SSH_AUTH_SOCK) \
		./scripts/run_goreleaser.sh --snapshot --skip=publish --skip=validate ${GO_RELEASER_RELEASE_ARGS}
	# remove the tag after building the snapshot
	git tag -d ${SNAPSHOT_VERSION}

.PHONY: install-snapshot
install-snapshot: snapshot verify-terraform-arch
	@echo "Installing snapshot version: ${SNAPSHOT_VERSION}-SNAPSHOT-${LATEST_GIT_COMMIT}"
	@echo "Installing snapshot binary: ${SNAPSHOT_BINARY}"
	mkdir -p ~/.terraform.d/plugins/${OS_ARCH}
	cp ./dist/terraform-provider-chronosphere_${OS_ARCH}/${SNAPSHOT_BINARY} ~/.terraform.d/plugins/${OS_ARCH}
	mkdir -p ~/.terraform.d/plugins/local/chronosphereio/chronosphere/${LOCAL_SNAPSHOT_VERSION_DIR}/${OS_ARCH}/
	cp ./dist/terraform-provider-chronosphere_${OS_ARCH}/${SNAPSHOT_BINARY} ~/.terraform.d/plugins/local/chronosphereio/chronosphere/${LOCAL_SNAPSHOT_VERSION_DIR}/${OS_ARCH}
	@echo "Installed snapshot version: ${SNAPSHOT_VERSION}-SNAPSHOT-${LATEST_GIT_COMMIT}"
	@echo ""
	@echo "reference this in your terraform code like so:"
	@echo ""
	@echo "provider \"chronosphere\" {"
	@echo "  source = \"local/chronosphereio/chronosphere\""
	@echo "  version = \"${SNAPSHOT_VERSION_NUMBER}-SNAPSHOT-${LATEST_GIT_COMMIT}\""
	@echo "}"
.PHONY: fmt
fmt:
	go fmt $(shell go list ./...)
	# TODO: enable once examples are added
	# terraform fmt examples/

.PHONY: lint
lint: install-tools
	@echo "--- :golang: linting code"
	GOFLAGS=$(GOFLAGS) $(TOOLS_BIN)/golangci-lint run -D gosec -D unused -E gofmt -E goimports -E gofumpt
