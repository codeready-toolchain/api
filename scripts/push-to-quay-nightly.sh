#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-to-quay-nightly.sh scripts overrides all the parameters but \"--project-root\", so use only that one to specify the root of the project."
    echo "                The parameters are overridden with these values:"
    echo "                      --channel nightly"
    echo "                      --template-version ${DEFAULT_VERSION}"
    echo "                      --next-version 0.0.<number-of-commits>-<short-sha-of-latest-commit>"
    echo "                      --replace-version 0.0.<number-of-commits-1>-<short-sha-of-last-but-one-commit>"
    echo ""
    echo "                Required variables:"
    echo "                      QUAY_NAMESPACE  - Quay namespace the operator bundle should be pushed to."
    echo ""
    echo "                Optional variables:"
    echo "                      QUAY_AUTH_TOKEN - Quay authentication token to be used for pushing to the quay namespace. If not set, then it's taken from ~/.docker/config.json file."
    echo ""
    echo "Example:"
    echo "   ./scripts/push-to-quay-nightly.sh -pr ../host-operator"
    echo "          - This command will generate CSV, CRDs and package info with the values defined above for the host-operator project"
    echo "            and pushes it to quay namespace defined by \"\${QUAY_NAMESPACE}\" variable."

}

push_to_quay() {
    TMP_FLATTEN_DIR="/tmp/${OPERATOR_NAME}_${NEXT_CSV_VERSION}_flatten"

    echo "## Pushing the OperatorHub package '${OPERATOR_NAME}' to the Quay.io '${QUAY_NAMESPACE}' organization ..."

    echo " - Flatten package to temporary folder: ${TMP_FLATTEN_DIR}"

    rm -Rf "${TMP_FLATTEN_DIR}" > /dev/null 2>&1
    mkdir -p "${TMP_FLATTEN_DIR}"
    operator-courier flatten "${PKG_DIR}" ${TMP_FLATTEN_DIR}

    echo " - Push flattened files to Quay.io namespace '${QUAY_NAMESPACE}' as version ${NEXT_CSV_VERSION}"

    if [[ -z ${QUAY_AUTH_TOKEN} ]]; then
        QUAY_AUTH_TOKEN=`cat ~/.docker/config.json | jq -r '.auths["quay.io"].auth'`
    fi

    operator-courier push ${TMP_FLATTEN_DIR} "${QUAY_NAMESPACE}" "${OPERATOR_NAME}" "${NEXT_CSV_VERSION}" "basic ${QUAY_AUTH_TOKEN}"

    echo "-> Operator bundle pushed."
}

# use the olm-setup as the source
OLM_SETUP_FILE=scripts/olm-setup.sh
if [[ -f ${OLM_SETUP_FILE} ]]; then
    source ${OLM_SETUP_FILE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/scripts/olm-setup.sh)"
    fi
fi
# read argument to get project root dir
read_arguments $@

# setup version and commit variables
GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD`
PREVIOUS_GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD^`
NEXT_CSV_VERSION="0.0.$(git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-list --count HEAD)-${GIT_COMMIT_ID}"
REPLACE_CSV_VERSION="0.0.$(git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-list --count HEAD^)-${PREVIOUS_GIT_COMMIT_ID}"

#read arguments one more time with the versions set
read_arguments $@ --channel nightly --template-version ${DEFAULT_VERSION} --next-version ${NEXT_CSV_VERSION} --replace-version ${REPLACE_CSV_VERSION}
setup_variables

# setup additional variables for pushing images
QUAY_NAMESPACE=${QUAY_NAMESPACE:codeready-toolchain}
IMAGE=quay.io/${QUAY_NAMESPACE}/${PRJ_NAME}:${GIT_COMMIT_ID}

# create backup of the current operator package directory
PKG_DIR_BACKUP=/tmp/deploy_olm-catalog_${PRJ_NAME}_backup
if [[ -d ${PKG_DIR_BACKUP} ]]; then
    rm -rf ${PKG_DIR_BACKUP}
fi
cp -r ${PKG_DIR} ${PKG_DIR_BACKUP}

# generate the bundle and push it to quay
generate_bundle
push_to_quay

# bring back the original operator package directory
rm -rf ${PKG_DIR}
cp -r ${PKG_DIR_BACKUP} ${PKG_DIR}