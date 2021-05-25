# current groupname and version of the operators'API
API_GROUPNAME=toolchain
API_FULL_GROUPNAME=toolchain.dev.openshift.com
API_VERSION:=v1alpha1

# how to dispatch the CRD files per repository (space-separated lists)
# !!! IMPORTANT !!! - when there is a new CRD added or an existing one removed or renamed, don't forget to change it also here: https://github.com/codeready-toolchain/toolchain-common/blob/master/scripts/add-cluster.sh#L54-L69
HOST_CLUSTER_CRDS:=masteruserrecords nstemplatetiers usersignups registrationservices bannedusers changetierrequests notifications tiertemplates templateupdaterequests toolchainstatuses toolchainclusters hostoperatorconfigs toolchainconfigs
MEMBER_CLUSTER_CRDS:=useraccounts nstemplatesets memberstatuses idlers toolchainclusters memberoperatorconfigs

PATH_TO_CRD_BASES=config/crd/bases

.PHONY: generate
## Generate deepcopy, openapi and CRD files after the API was modified
generate: vendor generate-deepcopy-and-crds generate-openapi dispatch-crds generate-csv copy-reg-service-template
	
.PHONY: generate-deepcopy-and-crds
generate-deepcopy-and-crds: remove-config
	@echo "Re-generating the deepcopy go file & the Toolchain CRD files... "
	$(Q)go run $(shell pwd)/vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go \
	crd:trivialVersions=true,preserveUnknownFields=false \
	object paths="./..." output:crd:artifacts:config=config/crd/bases
	
.PHONY: generate-openapi
generate-openapi:
	@echo "re-generating the openapi go file..."
	$(Q)go run $(shell pwd)/vendor/k8s.io/kube-openapi/cmd/openapi-gen/openapi-gen.go \
	--input-dirs ./api/$(API_VERSION)/ \
	--output-package github.com/codeready-toolchain/api/api/$(API_VERSION) \
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
host_repo_status := $(shell cd ../host-operator && git status -s | grep -v ${PATH_TO_CRD_BASES})
member_repo_status := $(shell cd ../member-operator && git status -s | grep -v ${PATH_TO_CRD_BASES})

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

.PHONY: dispatch-crds
dispatch-crds: vendor prepare-host-operator prepare-member-operator
	@echo "Dispatching CRD files in the 'host-operator' and 'member-operator' repositories..."
    # Dispatching CRD files to operator repositories
	@for crd in $(HOST_CLUSTER_CRDS) ; do \
		cp ${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml ../host-operator/deploy/crds/$(API_FULL_GROUPNAME)_$${crd}.yaml ; \
	done
	@for crd in $(MEMBER_CLUSTER_CRDS) ; do \
		cp ${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml ../member-operator/deploy/crds/$(API_FULL_GROUPNAME)_$${crd}.yaml ; \
	done
	# Now let's remove the CRDs from config/crd/bases directory
	@for crd in $(HOST_CLUSTER_CRDS) $(MEMBER_CLUSTER_CRDS) ; do \
		rm ${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml 2>/dev/null || true; \
	done
	@if [[ "$$(ls -A ${PATH_TO_CRD_BASES}/*yaml)" ]]; then \
	    echo "ERROR: some CRD files were not dispatched: $$(ls -A ${PATH_TO_CRD_BASES}/*yaml)"; \
	    echo "Please update this Makefile accordingly."; \
	    exit 1; \
	fi

.PHONY: copy-reg-service-template
copy-reg-service-template:
	cp ../registration-service/deploy/registration-service.yaml ../host-operator/deploy/registration-service/registration-service.yaml
