# current groupname and version of the operators'API
API_GROUPNAME=toolchain
API_FULL_GROUPNAME=toolchain.dev.openshift.com
API_VERSION:=v1alpha1

# how to dispatch the CRD files per repository (space-separated lists)
HOST_CLUSTER_CRDS:=masteruserrecord nstemplatetier usersignup registrationservice banneduser
MEMBER_CLUSTER_CRDS:=useraccount nstemplateset

.PHONY: generate
## Generate deepcopy, openapi and CRD files after the API was modified
generate: vendor generate-deepcopy generate-openapi generate-crds generate-csv copy-reg-service-template
	
.PHONY: generate-deepcopy
generate-deepcopy:
	@echo "re-generating the deepcopy go file..."
	$(Q)go run $(shell pwd)/vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go \
	--input-dirs ./pkg/apis/$(API_GROUPNAME)/$(API_VERSION)/ -O zz_generated.deepcopy \
	--bounding-dirs github.com/codeready-toolchain/api/pkg/apis "$(API_GROUPNAME):$(API_VERSION)" \
	--go-header-file=make/go-header.txt
	
.PHONY: generate-openapi
generate-openapi:
	@echo "re-generating the openapi go file..."
	$(Q)go run $(shell pwd)/vendor/k8s.io/kube-openapi/cmd/openapi-gen/openapi-gen.go \
	--input-dirs ./pkg/apis/$(API_GROUPNAME)/$(API_VERSION)/ \
	--output-package github.com/codeready-toolchain/api/pkg/apis/$(API_GROUPNAME)/$(API_VERSION) \
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

PHONY: generate-csv
generate-csv:
	./scripts/olm-catalog-generate.sh -pr ../host-operator
	./scripts/olm-catalog-generate.sh -pr ../member-operator

PHONY: prepare-host-operator
prepare-host-operator: ../host-operator
ifdef host_repo_status
	@echo "The local '../host-operator' repository has pending changes. Please stash them or commit them, first."
	@exit 1
endif
ifneq ($(wildcard ../host-operator/deploy/crds/*.yaml),)
	@-find ../host-operator/deploy/crds -type f -not -name "*kubefed*" | xargs rm
else
	@-mkdir -p ../host-operator/deploy/crds
endif

PHONY: prepare-member-operator
prepare-member-operator: ../member-operator
ifdef member_repo_status
	@echo "The local '../member-operator' repository has pending changes. Please stash them or commit them, first."
	@exit 1
endif
ifneq ($(wildcard ../member-operator/deploy/crds/*.yaml),)
	@-find ../member-operator/deploy/crds -type f -not -name "*kubefed*" | xargs rm
else
	@-mkdir -p ../member-operator/deploy/crds
endif

.PHONY: generate-crds
generate-crds: vendor prepare-host-operator prepare-member-operator generate-kubefed-crd
	@echo "Re-generating the Toolchain CRD files..."
	$(Q)go run $(shell pwd)/vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd:trivialVersions=true \
	paths=./pkg/apis/... output:crd:dir=deploy/crds output:stdout
	@echo "Dispatching CRD files in the 'host-operator' and 'member-operator' repositories..."
    # When dispatching CRD files we delete two first lines of CRDs ("\n----\n") to make a single manifest file out of the original multiple manifest file
    # Also we remove the line with 'type: object' from validation.openAPIV3Schema.properties path because it's incompatible with kube 1.11 which is used by minishift
	@for crd in $(HOST_CLUSTER_CRDS) ; do \
		sed -e '1,2d' -e '/^      type: object/d' deploy/crds/$(API_FULL_GROUPNAME)_$${crd}s.yaml > ../host-operator/deploy/crds/$(API_GROUPNAME)_$(API_VERSION)_$${crd}_crd.yaml ; \
		rm deploy/crds/$(API_FULL_GROUPNAME)_$${crd}s.yaml; \
	done
	@for crd in $(MEMBER_CLUSTER_CRDS) ; do \
		sed -e '1,2d' -e '/^      type: object/d' deploy/crds/$(API_FULL_GROUPNAME)_$${crd}s.yaml > ../member-operator/deploy/crds/$(API_GROUPNAME)_$(API_VERSION)_$${crd}_crd.yaml ; \
		rm deploy/crds/$(API_FULL_GROUPNAME)_$${crd}s.yaml; \
	done
ifneq ($(wildcard deploy/crds/*.yaml),)
	@echo "ERROR: some CRD files were not dispatched: $(wildcard deploy/crds/*.yaml)"
	@echo "Please update this Makefile accordingly."
	@exit 1
endif

.PHONY: generate-kubefed-crd
generate-kubefed-crd: vendor
	@echo "Re-generating the KubeFed CRD..."
	$(Q)go run $(shell pwd)/vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd:trivialVersions=true \
	paths=./vendor/sigs.k8s.io/kubefed/pkg/apis/core/v1beta1/... output:crd:dir=deploy/crds/kubefed/row output:stdout
    # Delete two first lines of the CRD ("\n----\n") to make a single manifest file out of the original multiple manifest file
    # Also remove the line with 'type: object' from validation.openAPIV3Schema.properties path because it's incompatible with kube 1.11 which is used by minishift
	@sed -e '1,2d' -e '/^      type: object/d' deploy/crds/kubefed/row/core.kubefed.io_kubefedclusters.yaml > deploy/crds/kubefed/core.kubefed.io_kubefedclusters.yaml
	@echo "Generating bindata and dispatching it in the 'toolchain-common' repository..."
	@go install github.com/go-bindata/go-bindata/...
	@$(GOPATH)/bin/go-bindata -pkg cluster -o ../toolchain-common/pkg/cluster/kubefedcluster_assets.go -nocompress -prefix deploy/crds/kubefed deploy/crds/kubefed
	@rm -rf deploy/crds/kubefed

.PHONY: copy-reg-service-template
copy-reg-service-template:
	cp ../registration-service/deploy/registration-service.yaml ../host-operator/deploy/registration-service/registration-service.yaml