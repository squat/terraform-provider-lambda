export GO111MODULE=on
.PHONY: clean deps fmt generate lint test unit

SHELL := /usr/bin/env bash
ARCH ?= amd64
OS ?= linux
ALL_ARCH := amd64 arm arm64
DOCKER_ARCH := "amd64" "arm v7" "arm64 v8"
BINS := $(addprefix bin/$(ARCH)/,terraform-provider-lambda)
PROJECT := terraform-provider-lambda
PKG := github.com/squat/$(PROJECT)

TAG := $(shell git describe --abbrev=0 --tags HEAD 2>/dev/null)
COMMIT := $(shell git rev-parse HEAD)
VERSION := $(COMMIT)
ifneq ($(TAG),)
    ifeq ($(COMMIT), $(shell git rev-list -n1 $(TAG)))
        VERSION := $(TAG)
    endif
endif
DIRTY := $(shell test -z "$$(git diff --shortstat 2>/dev/null)" || echo -dirty)
VERSION := $(VERSION)$(DIRTY)
LD_FLAGS := -ldflags "-X main.Version=$(VERSION) -extldflags -static"
GO_FILES := $(shell find . -name '*.go')
ifeq ($(GO_FILES),)
GO_FILES := main.go
endif
GO_PKGS := $(shell go list ./...)
SPEAKEASY_FILES := go.mod go.sum $(GO_FILES)
DOCS := $(shell find docs -type f -name '*.md')
ifeq ($(DOCS),)
DOCS := docs/index.md
endif
GENERATED := $(SPEAKEASY_FILES) $(DOCS)

STATICCHECK_BINARY := go run honnef.co/go/tools/cmd/staticcheck@2023.1.6
SPEAKEASY_BINARY := bin/speakeasy

GO_VERSION ?= 1.21.3
BUILD_IMAGE ?= golang:$(GO_VERSION)-alpine

build: $(BINS)

build-%:
	@$(MAKE) --no-print-directory ARCH=$* build

all-build: $(addprefix build-, $(ALL_ARCH))

CONTAINERIZE_BUILD ?= true
BUILD_PREFIX :=
BUILD_SUFIX :=
ifeq ($(CONTAINERIZE_BUILD), true)
	BUILD_PREFIX := docker run --rm \
	    -u $$(id -u):$$(id -g) \
	    -v $$(pwd):/src \
	    -w /src \
	    --entrypoint '' \
	    $(BUILD_IMAGE) \
	    /bin/sh -c ' \
	        GOCACHE=$$(pwd)/.cache
	BUILD_SUFIX := '
endif

$(BINS): $(GO_FILES) go.mod
	@mkdir -p bin/$(ARCH)
	@echo "building: $@"
	@$(BUILD_PREFIX) \
	        GOARCH=$(ARCH) \
	        GOOS=linux \
		CGO_ENABLED=0 \
		go build -o $@ \
		    $(LD_FLAGS) \
		    . \
	$(BUILD_SUFIX)

fmt:
	@echo $(GO_PKGS)
	gofmt -w -s $(GO_FILES)

lint:
	@echo 'go vet $(GO_PKGS)'
	@vet_res=$$(GO111MODULE=on go vet $(GO_PKGS) 2>&1); if [ -n "$$vet_res" ]; then \
		echo ""; \
		echo "Go vet found issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review:"; \
		echo "$$vet_res"; \
		exit 1; \
	fi
	@echo '$(STATICCHECK_BINARY) $(GO_PKGS)'
	@lint_res=$$($(STATICCHECK_BINARY) $(GO_PKGS)); if [ -n "$$lint_res" ]; then \
		echo ""; \
		echo "Staticcheck found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review:"; \
		echo "$$lint_res"; \
		exit 1; \
	fi
	@echo 'gofmt -d -s $(GO_FILES)'
	@fmt_res=$$(gofmt -d -s $(GO_FILES)); if [ -n "$$fmt_res" ]; then \
		echo ""; \
		echo "Gofmt found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review:"; \
		echo "$$fmt_res"; \
		exit 1; \
	fi

unit:
	go test --race ./...

test: lint unit

bin-clean:
	rm -rf bin

deps:
	go get go@$(GO_VERSION) toolchain@go$(GO_VERSION)
	go mod tidy

$(SPEAKEASY_BINARY):
	mkdir -p $(@D)
	cd $(@D) && curl https://github.com/speakeasy-api/speakeasy/releases/download/v1.402.0/speakeasy_$(OS)_$(ARCH).zip -L -o speakeasy.zip && unzip -o speakeasy.zip $(@F) && rm speakeasy.zip ; chmod +x $(@F)

lambda.yaml: overlay.yaml $(SPEAKEASY_BINARY)
	$(SPEAKEASY_BINARY) overlay apply --overlay overlay.yaml --schema <(curl https://cloud.lambdalabs.com/static/api/v1/openapi.yaml) > $@

$(SPEAKEASY_FILES) &: lambda.yaml $(SPEAKEASY_BINARY)
	$(SPEAKEASY_BINARY) generate sdk --lang terraform --schema $< --out .
	$(MAKE) fmt
	$(MAKE) deps

$(DOCS) &: $(SPEAKEASY_FILES)
	go generate
	sed -i 's/lambda Provider/Lambda Provider/' docs/index.md

generate: $(GENERATED)

-include lambda.mk
