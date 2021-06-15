#!/usr/bin/env bash

OLM_SETUP_FILE=scripts/olm-setup.sh
OWNER_AND_BRANCH_LOCATION=${OWNER_AND_BRANCH_LOCATION:-codeready-toolchain/api/master}

if [[ -f ${OLM_SETUP_FILE} ]]; then
    source ${OLM_SETUP_FILE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/${OWNER_AND_BRANCH_LOCATION}/${OLM_SETUP_FILE})"
    fi
fi

install_operator() {
    DISPLAYNAME=$(echo ${OPERATOR_NAME} | tr '-' ' ' | awk '{for (i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')

    GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD`
    INDEX_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${INDEX_IMAGE_NAME}:${INDEX_IMAGE_TAG:-latest}
    CATALOGSOURCE_NAME=source-${OPERATOR_NAME}-${GIT_COMMIT_ID}
    SUBSCRIPTION_NAME=subscription-${OPERATOR_NAME}-${GIT_COMMIT_ID}
    INSTALL_OBJECTS="apiVersion: operators.coreos.com/v1
kind: OperatorGroup
metadata:
  name: og-${OPERATOR_NAME}
  namespace: ${NAMESPACE}
spec:
  targetNamespaces:
  - ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: ${SUBSCRIPTION_NAME}
  namespace: ${NAMESPACE}
spec:
  channel: ${CHANNEL}
  installPlanApproval: Automatic
  name: ${OPERATOR_NAME}
  source: ${CATALOGSOURCE_NAME}
  sourceNamespace: ${NAMESPACE}
---
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: ${CATALOGSOURCE_NAME}
  namespace: ${NAMESPACE}
spec:
  sourceType: grpc
  image: ${INDEX_IMAGE}
  displayName: ${DISPLAYNAME}
  publisher: Red Hat
  updateStrategy:
    registryPoll:
      interval: 1m0s"
    echo "objects to be created in order to install operator"
    cat <<EOF | oc apply -f -
${INSTALL_OBJECTS}
EOF
}

wait_until_is_installed() {
    set -e
    NEXT_WAIT_TIME=0
    SA_NAME=${PRJ_NAME}-controller-manager
    while [[ -z `oc get sa ${SA_NAME} -n ${NAMESPACE} 2>/dev/null` ]] || [[ -z `oc get ClusterRoles | grep "^toolchain-${PRJ_NAME}\.v"` ]]; do \
        if [[ ${NEXT_WAIT_TIME} -eq 300 ]]; then \
           echo "reached timeout of waiting for ServiceAccount ${PRJ_NAME} to be available in namespace ${NAMESPACE} - see following info for debugging:"; \
           echo "================================ CatalogSource =================================="; \
           oc get catalogsource ${CATALOGSOURCE_NAME} -n ${NAMESPACE} -o yaml; \
           echo "================================ CatalogSource Pod Logs =================================="; \
           oc logs `oc get pods -l "olm.catalogSource=${CATALOGSOURCE_NAME#*/}" -n ${NAMESPACE} -o name` -n ${NAMESPACE}; \
           echo "================================ Subscription =================================="; \
           oc get subscription ${SUBSCRIPTION_NAME} -n ${NAMESPACE} -o yaml; \
           echo "================================ InstallPlans =================================="; \
           oc get installplans -n ${NAMESPACE} -o yaml; \
           exit 1; \
        fi; \
        echo "$(( NEXT_WAIT_TIME++ )). attempt of waiting for ServiceAccount ${SA_NAME} in namespace ${NAMESPACE}"; \
        sleep 1; \
    done
}

read_arguments $@
setup_variables
install_operator
wait_until_is_installed
