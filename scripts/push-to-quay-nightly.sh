#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-to-quay-nightly.sh scripts overrides all the parameters but \"--project-root\" and \"--operator-name\", so use only these two to specify necessary values."
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
NEXT_CSV_VERSION="0.0.$(git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-list --count HEAD)-commit-${GIT_COMMIT_ID}"
REPLACE_CSV_VERSION="0.0.$(git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-list --count HEAD^)-commit-${PREVIOUS_GIT_COMMIT_ID}"

QUAY_NAMESPACE=${QUAY_NAMESPACE:codeready-toolchain}

# generate manifests
generate_manifests $@ --channel nightly --template-version ${DEFAULT_VERSION} --next-version ${NEXT_CSV_VERSION} --replace-version ${REPLACE_CSV_VERSION}

# push manifests to quay
DIR_TO_PUSH=${PKG_DIR}
push_to_quay

# bring back the original operator package directory
rm -rf ${PKG_DIR}
cp -r ${PKG_DIR_BACKUP} ${PKG_DIR}