# enable Go modules
GO111MODULE?=on
export GO111MODULE
GO?=go

# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}

.PHONY: build
## Build
build:
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    $(GO) build github.com/codeready-toolchain/api/api/v1alpha1/

TMP_DIR = /tmp/
BASE_REPO_PATH = $(shell mktemp -d ${TMP_DIR}crt-verify.XXX)
GH_BASE_URL_KS = https://github.com/kubesaw/
GH_BASE_URL_CRT = https://github.com/codeready-toolchain/
GH_KSCTL = $(GH_BASE_URL_KS)ksctl
GH_HOST = $(GH_BASE_URL_CRT)host-operator
GH_MEMBER = $(GH_BASE_URL_CRT)member-operator
GH_REGSVC = $(GH_BASE_URL_CRT)registration-service
GH_E2E = $(GH_BASE_URL_CRT)toolchain-e2e
GH_TC = $(GH_BASE_URL_CRT)toolchain-common

.PHONY: verify-replace-run
verify-replace-run: generate
	$(eval C_PATH = $(PWD))\
	$(foreach repo,${GH_HOST} ${GH_MEMBER} ${GH_REGSVC} ${GH_E2E} ${GH_TC} ${GH_KSCTL},\
	$(eval REPO_PATH = ${BASE_REPO_PATH}/$(shell basename $(repo))) \
	git clone --depth=1 $(repo) ${REPO_PATH}; \
	cd ${REPO_PATH}; \
	if [[ "${GH_HOST}" == "$(repo)" ]]; then \
		$(MAKE) generate ; \
	elif [[ "${GH_MEMBER}" == "$(repo)" ]]; then\
		$(MAKE) generate-assets ; \
	fi; \
	go mod edit -replace github.com/codeready-toolchain/api=${C_PATH}; \
	$(MAKE) verify-dependencies; \
	)
	
