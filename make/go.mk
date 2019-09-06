# enable Go modules
GO111MODULE?=on
export GO111MODULE

# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}

.PHONY: build
## Build
build: $(shell find . -path ./vendor -prune -o -name '*.go' -print)
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    go build github.com/codeready-toolchain/api/pkg/apis/

.PHONY: vendor
vendor: 
	$(Q)go mod vendor
