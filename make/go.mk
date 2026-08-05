# enable Go modules
GO111MODULE?=on
export GO111MODULE
GO?=go

# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}

GOFORMAT_FILES := $(shell find  . -name '*.go' | grep -vEf ./make/gofmt_exclude)

.PHONY: format-go-code
## Formats any go file that does not match formatting defined by gofmt
format-go-code:
	$(Q)gofmt -s -l -w ${GOFORMAT_FILES}

.PHONY: build
## Build
build:
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    $(GO) build github.com/codeready-toolchain/api/api/v1alpha1/

.PHONY: verify-replace-run
verify-replace-run: 
	./scripts/verify-replace.sh;

