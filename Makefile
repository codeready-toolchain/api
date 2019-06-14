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
build: $(shell find . -path ./vendor -prune -o -name '*.go' -print)
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    go build github.com/codeready-toolchain/api/pkg/apis/

.PHONY: generate
## Generate deepcopy after modifying API
# generate: vendor generate-deepcopy generate-openapi generate-crds
generate: generate-crds
	
.PHONY: generate-deepcopy
generate-deepcopy:
	@echo "re-generating the deepcopy go file..."
	$(Q)go run $(shell pwd)/vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go \
	--input-dirs ./pkg/apis/toolchain/v1alpha1/ -O zz_generated.deepcopy \
	--bounding-dirs github.com/codeready-toolchain/api/pkg/apis "toolchain:v1alpha1" \
	--go-header-file=make/go-header.txt
	
.PHONY: generate-openapi
generate-openapi:
	@echo "re-generating the openapi go file..."
	$(Q)go run $(shell pwd)/vendor/k8s.io/kube-openapi/cmd/openapi-gen/openapi-gen.go \
	--input-dirs ./pkg/apis/toolchain/v1alpha1/ \
	--output-package github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1 \
	--output-file-base zz_generated.openapi \
	--go-header-file=make/go-header.txt

# make sure that that the `host-operator` and `member-operator` repositories exist locally 
# and that they don't have any pending changes (except for the CRD files). 
# The reasonning here is that when a change is made in the `api` repository, the resulting changes 
# in the `host-operator` and `member-operator` repositories can be pushed at the same time on GitHub, 
# without having to wait for some other feature or fix to be completed.
# TODO: we could even go further and checkout new branches from `master` in the 'host-operator'
# and 'member-operator` repositories and give them the name of the current branch in this repo.
# The developer would have 3 branches with the same name and could then push to GitHub at the 
# same time...
host_repo_status := $(shell cd ../host-operator && git status -s | grep -v deploy/crds)
member_repo_status := $(shell cd ../member-operator && git status -s | grep -v deploy/crds)

PHONY: host-operator
host-operator: ../host-operator
ifdef host_repo_status
	@echo "The local '../host-operator' repository has pending changes. Please stash them or commit them, first."
	@exit 1
endif
ifneq ($(wildcard ../host-operator/deploy/crds/*.yaml),)
	@-rm ../host-operator/deploy/crds/*.yaml
else
	@-mkdir -p ../host-operator/deploy/crds
endif

PHONY: member-operator
member-operator: ../member-operator
ifdef member_repo_status
	@echo "The local '../member-operator' repository has pending changes. Please stash them or commit them, first."
	@exit 1
endif
ifneq ($(wildcard ../member-operator/deploy/crds/*.yaml),)
	@-rm ../member-operator/deploy/crds/*.yaml
else
	@-mkdir -p ../member-operator/deploy/crds
endif


.PHONY: generate-crds
generate-crds: host-operator member-operator
	@echo "Re-generating the CRD files..."
	$(Q)go run $(shell pwd)/vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd \
	--domain dev.openshift.com \
	--output-dir deploy/crds
	@echo "Dispatching CRD files in the 'host-operator' and 'member-operator' repositories..."
	@mv deploy/crds/toolchain_v1alpha1_masteruserrecord.yaml ../host-operator/deploy/crds
	@mv deploy/crds/toolchain_v1alpha1_nstemplatetier.yaml ../host-operator/deploy/crds
	@mv deploy/crds/toolchain_v1alpha1_useraccount.yaml ../host-operator/deploy/crds
	@mv deploy/crds/toolchain_v1alpha1_nstemplateset.yaml ../member-operator/deploy/crds
ifneq ($(wildcard deploy/crds/*.yaml),)
	@echo "ERROR: some CRD files were not dispatched: $(wildcard deploy/crds/*.yaml)"
	@echo "Please update this Makefile accordingly."
	@exit 1
endif

.PHONY: test
## runs the tests
test:
	@echo "running the tests..."
	$(Q)go test ./...

.PHONY: clean
## Clean
clean:
	$(Q)-rm -rf ${V_FLAG} ./vendor
	$(Q)go clean ${X_FLAG} ./...

.PHONY: vendor
vendor: 
	$(Q)go mod vendor
