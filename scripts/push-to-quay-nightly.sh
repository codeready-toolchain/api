#!/usr/bin/env bash


push_to_quay() {
    TMP_FLATTEN_DIR="/tmp/${OPERATOR_NAME}_${GIT_COMMIT_ID}_flatten"

    echo "## Pushing the OperatorHub package '${OPERATOR_NAME}' to the Quay.io '${QUAY_NAMESPACE}' organization ..."

    echo " - Flatten package to temporary folder: ${TMP_FLATTEN_DIR}"

    rm -Rf "${TMP_FLATTEN_DIRREPLACE_VERSION}" > /dev/null 2>&1
    mkdir -p "${TMP_FLATTEN_DIR}"
    operator-courier flatten "${PKG_DIR}" ${TMP_FLATTEN_DIR}

    echo " - Push flattened files to Quay.io namespace '${QUAY_NAMESPACE}' as version ${NEXT_CSV_VERSION}"

    if [[ -z ${QUAY_AUTH_TOKEN} ]]; then
        QUAY_AUTH_TOKEN=`cat ~/.docker/config.json | jq -r '.auths["quay.io"].auth'`
    fi

    operator-courier push ${TMP_FLATTEN_DIR} "${QUAY_NAMESPACE}" "${OPERATOR_NAME}" "${NEXT_CSV_VERSION}" "basic ${QUAY_AUTH_TOKEN}"

    echo "-> Operator bundle pushed."
}

GIT_COMMIT_ID=`git rev-parse --short HEAD`
PREVIOUS_GIT_COMMIT_ID=`git rev-parse --short HEAD^`

NEXT_CSV_VERSION="0.0.$(git rev-list --count HEAD)-${GIT_COMMIT_ID}"
CURRENT_CSV_VERSION="0.0.$(git rev-list --count HEAD^)-${PREVIOUS_GIT_COMMIT_ID}"

OLM_SETUP_FILE=./scripts/olm-setup.sh

if [[ -f ${OLM_SETUP_FILE} ]]; then
    source ${OLM_SETUP_FILE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/scripts/olm-setup.sh)"
    fi
fi

read_arguments $@ --channel nightly
setup_variables

QUAY_NAMESPACE=${QUAY_NAMESPACE:codeready-toolchain}
IMAGE=quay.io/${QUAY_NAMESPACE}/${PRJ_NAME}:${GIT_COMMIT_ID}

generate_bundle
push_to_quay

