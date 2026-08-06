# enable Go modules
GO111MODULE?=on
export GO111MODULE
GO?=go

# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}

.PHONY: format-go-code
## Formats any go file that does not match formatting defined by gofmt
format-go-code:
# The + tells find to batch multiple found files into a single gofmt invocation (like xargs),
# which is much faster than the alternative \;, which runs gofmt once per file. Removing it
# would be a syntax error — find -exec requires either + or \; as a terminator.
	$(Q)find . -name '*.go' -not -path '*/vendor/*' -not -path '*/.git/*' -exec gofmt -s -l -w {} +

.PHONY: build
## Build
build:
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    $(GO) build github.com/codeready-toolchain/api/api/v1alpha1/

.PHONY: verify-replace-run
verify-replace-run: 
	./scripts/verify-replace.sh;

