# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

include ./make/verbose.mk
.DEFAULT_GOAL := help
include ./make/help.mk
include ./make/out.mk
include ./make/go.mk
include ./make/git.mk
include ./make/format.mk
include ./make/lint.mk

.PHONY: build
## Build
build: vendor $(shell find . -path ./vendor -prune -o -name '*.go' -print)
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    go build github.com/codeready-toolchain/api/pkg/apis/

.PHONY: generate
## Generate deepcopy after modifying API
generate: vendor
	@echo "re-generating the deepcopy files..."
	$(Q)go run $(shell pwd)/vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go \
	--input-dirs ./pkg/apis/toolchain/v1alpha1/ -O zz_generated.deepcopy \
	--bounding-dirs github.com/codeready-toolchain/api/pkg/apis "toolchain:v1alpha1" \
	--go-header-file=make/go-header.txt
	
	
.PHONY: clean
## Clean
clean:
	$(Q)-rm -rf ${V_FLAG} ./vendor
	$(Q)go clean ${X_FLAG} ./...

.PHONY: vendor
vendor: 
	$(Q)go mod vendor
