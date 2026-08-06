# current groupname and version of the operators'API
API_GROUPNAME=toolchain
API_FULL_GROUPNAME=toolchain.dev.openshift.com
API_VERSION:=v1alpha1

PROJECT_DIR := $(shell pwd)

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)


## Tool Binaries
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
OPENAPI_GEN ?= $(LOCALBIN)/openapi-gen
CRD_REF_DOCS ?= $(LOCALBIN)/crd-ref-docs
PATH_TO_CRD_BASES=config/crd/bases

$(CONTROLLER_GEN): go.mod | $(LOCALBIN) ## install controller-gen locally if necessary. Version is pinned in go.mod.
	GOBIN=$(LOCALBIN) $(GO) install sigs.k8s.io/controller-tools/cmd/controller-gen

$(OPENAPI_GEN): go.mod | $(LOCALBIN) ## install openapi-gen locally if necessary. Version is pinned in go.mod.
	GOBIN=$(LOCALBIN) $(GO) install k8s.io/kube-openapi/cmd/openapi-gen

$(CRD_REF_DOCS): go.mod | $(LOCALBIN) ## install crd-ref-docs locally if necessary. Version is pinned in go.mod.
	GOBIN=$(LOCALBIN) $(GO) install github.com/elastic/crd-ref-docs

.PHONY: manifests
manifests: $(CONTROLLER_GEN) ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./api/..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: generate-object generate-crd gen-crd-ref-docs generate-openapi dispatch-crds ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.

.PHONY: generate-object
generate-object: $(CONTROLLER_GEN) ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object paths="./api/..."

.PHONY: generate-crd
generate-crd: $(CONTROLLER_GEN) remove-config ## Generate CRD manifests.
	$(CONTROLLER_GEN) crd paths="./api/..." output:crd:artifacts:config=config/crd/bases

.PHONY: gen-crd-ref-docs
gen-crd-ref-docs: $(CRD_REF_DOCS)
	@echo "Re-generating the api doc ref: ./api/$(API_VERSION)/docs/apiref.adoc "
	$(CRD_REF_DOCS) --source-path ./api/$(API_VERSION) --config ./crdrefdocs.config.yaml --output-path ./api/$(API_VERSION)/docs/apiref.adoc

.PHONY: generate-openapi
generate-openapi: $(OPENAPI_GEN)
	@echo "re-generating the openapi go file..."
	$(OPENAPI_GEN) ./api/$(API_VERSION)/ \
	--output-pkg github.com/codeready-toolchain/api/api/$(API_VERSION) \
	--output-file zz_generated.openapi.go \
	--output-dir ./api/$(API_VERSION) \
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

# how to dispatch the CRD files per repository (space-separated lists)
HOST_CLUSTER_CRDS:=masteruserrecords nstemplatetiers usersignups bannedusers notifications spaces spacebindings socialevents tiertemplates tiertemplaterevisions toolchainstatuses toolchainclusters toolchainconfigs usertiers proxyplugins spacerequests spacebindingrequests spaceprovisionerconfigs
MEMBER_CLUSTER_CRDS:=useraccounts nstemplatesets memberstatuses idlers toolchainclusters memberoperatorconfigs spacerequests workspaces spacebindingrequests

.PHONY: dispatch-crds
dispatch-crds: prepare-host-operator prepare-member-operator
	@echo "Dispatching CRD files in the 'host-operator' and 'member-operator' repositories..."
    # Dispatching CRD files to operator repositories
	@for crd in $(HOST_CLUSTER_CRDS) ; do \
		cp ${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml ../host-operator/${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml ; \
	done
	@for crd in $(MEMBER_CLUSTER_CRDS) ; do \
		cp ${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml ../member-operator/${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml ; \
	done
	# Now let's remove the CRDs from config/crd/bases directory
	@for crd in $(HOST_CLUSTER_CRDS) $(MEMBER_CLUSTER_CRDS) ; do \
		rm ${PATH_TO_CRD_BASES}/$(API_FULL_GROUPNAME)_$${crd}.yaml 2>/dev/null || true; \
	done
	@if [[ "$$(ls -A ${PATH_TO_CRD_BASES}/*yaml || true)" ]]; then \
	    echo "ERROR: some CRD files were not dispatched: $$(ls -A ${PATH_TO_CRD_BASES}/*yaml)"; \
	    echo "Please update this Makefile accordingly."; \
	    exit 1; \
	fi
	@echo "Dispatch successfully finished \o/"

PHONY: prepare-host-operator
prepare-host-operator: ../host-operator
ifdef host_repo_status
	@echo "The local '../host-operator' repository has pending changes. Please stash them or commit them, first."
	@exit 1
endif
ifneq ($(wildcard ../host-operator/${PATH_TO_CRD_BASES}/*.yaml),)
	@-find ../host-operator/${PATH_TO_CRD_BASES} -type f | grep -v "cr\.yaml" | xargs rm || true
else
	@-mkdir -p ../host-operator/${PATH_TO_CRD_BASES}
endif

PHONY: prepare-member-operator
prepare-member-operator: ../member-operator
ifdef member_repo_status
	@echo "The local '../member-operator' repository has pending changes. Please stash them or commit them, first."
	@exit 1
endif
ifneq ($(wildcard ../member-operator/${PATH_TO_CRD_BASES}/*.yaml),)
	@-find ../member-operator/${PATH_TO_CRD_BASES} -type f | grep -v "cr\.yaml" | xargs rm || true
else
	@-mkdir -p ../member-operator/${PATH_TO_CRD_BASES}
endif