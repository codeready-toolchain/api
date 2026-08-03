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
PATH_TO_CRD_BASES=config/crd/bases
CONTROLLER_TOOLS_VERSION ?= v0.18.0
OPERATOR_SDK_VERSION ?= v1.42.0

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	$(call go-install-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen,$(CONTROLLER_TOOLS_VERSION))

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./api/..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object paths="./api/..."

CRD_REF_DOCS = $(PROJECT_DIR)/bin/crd-ref-docs
crd-ref-docs: ## Download crd-ref-docs locally if necessary.
	GOBIN=$(PROJECT_DIR)/bin $(GO) install github.com/elastic/crd-ref-docs@latest

.PHONY: gen-crd-ref-docs
gen-crd-ref-docs: crd-ref-docs
	@echo "Re-generating the api doc ref: ./api/$(API_VERSION)/docs/apiref.adoc "
	$(CRD_REF_DOCS) --source-path ./api/$(API_VERSION) --config ./crdrefdocs.config.yaml --output-path ./api/$(API_VERSION)/docs/apiref.adoc

OPENAPI_GEN = $(PROJECT_DIR)/bin/openapi-gen
OPENAPI_GEN_VERSION ?= master

openapi-gen: ## Download openapi-gen locally if necessary.
	$(call go-install-tool,$(OPENAPI_GEN),k8s.io/kube-openapi/cmd/openapi-gen,$(OPENAPI_GEN_VERSION))

.PHONY: generate-openapi
generate-openapi: openapi-gen
	@echo "re-generating the openapi go file..."
	$(OPENAPI_GEN) ./api/$(API_VERSION)/ \
	--output-pkg github.com/codeready-toolchain/api/api/$(API_VERSION) \
	--output-file zz_generated.openapi.go \
	--output-dir ./api/$(API_VERSION) \
	--go-header-file=make/go-header.txt



# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f $(1) || true ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv $(1) $(1)-$(3) ;\
} ;\
ln -sf $(1)-$(3) $(1)
endef


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