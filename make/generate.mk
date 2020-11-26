# current groupname and version of the operators'API
API_GROUPNAME=toolchain
API_FULL_GROUPNAME=toolchain.dev.openshift.com
API_VERSION:=v1alpha1

# how to dispatch the CRD files per repository (space-separated lists)
# !!! IMPORTANT !!! - when there is a new CRD added or an existing one removed or renamed, don't forget to change it also here: https://github.com/codeready-toolchain/toolchain-common/blob/master/scripts/add-cluster.sh#L54-L69
HOST_CLUSTER_CRDS:=masteruserrecord nstemplatetier usersignup registrationservice banneduser changetierrequest notification tiertemplate templateupdaterequest toolchainstatus toolchaincluster hostoperatorconfig
MEMBER_CLUSTER_CRDS:=useraccount nstemplateset memberstatus idler toolchaincluster

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
	@-find ../host-operator/deploy/crds -type f | grep -v "cr\.yaml" | xargs rm || true
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
	@-find ../member-operator/deploy/crds -type f | grep -v "cr\.yaml" | xargs rm || true
else
	@-mkdir -p ../member-operator/deploy/crds
endif

.PHONY: generate-crds
generate-crds: vendor prepare-host-operator prepare-member-operator
	@echo "Re-generating the Toolchain CRD files..."
	rm -rf deploy/crds/
	operator-sdk generate crds
	@echo "Dispatching CRD files in the 'host-operator' and 'member-operator' repositories..."
    # Dispatching CRD files to operator repositories
	@for crd in $(HOST_CLUSTER_CRDS) ; do \
		crd_plural=$$(echo $${crd} | sed -e 's/s$$/se/')s; \
		cp deploy/crds/$(API_FULL_GROUPNAME)_$${crd_plural}_crd.yaml ../host-operator/deploy/crds/$(API_GROUPNAME)_$(API_VERSION)_$${crd}_crd.yaml ; \
	done
	@for crd in $(MEMBER_CLUSTER_CRDS) ; do \
		crd_plural=$$(echo $${crd} | sed -e 's/s$$/se/')s; \
		cp deploy/crds/$(API_FULL_GROUPNAME)_$${crd_plural}_crd.yaml ../member-operator/deploy/crds/$(API_GROUPNAME)_$(API_VERSION)_$${crd}_crd.yaml ; \
	done
	# Now let's remove the CRDs from deploy/crds directory
	@for crd in $(HOST_CLUSTER_CRDS) $(MEMBER_CLUSTER_CRDS) ; do \
		crd_plural=$$(echo $${crd} | sed -e 's/s$$/se/')s; \
		rm deploy/crds/$(API_FULL_GROUPNAME)_$${crd_plural}_crd.yaml 2>/dev/null || true; \
	done
ifneq ($(wildcard deploy/crds/*.yaml),)
	@echo "ERROR: some CRD files were not dispatched: $(wildcard deploy/crds/*.yaml)"
	@echo "Please update this Makefile accordingly."
	@exit 1
endif

.PHONY: copy-reg-service-template
copy-reg-service-template:
	cp ../registration-service/deploy/registration-service.yaml ../host-operator/deploy/registration-service/registration-service.yaml
