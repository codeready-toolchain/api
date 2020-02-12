#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-to-quay-nightly.sh scripts overrides all the parameters but \"--project-root\" and \"--operator-name\", so use only these two to specify necessary values."
    echo "                The parameters are overridden with these values:"
    echo "                      --channel nightly"
    echo "                      --template-version ${DEFAULT_VERSION}"
    echo "                      --next-version 0.0.<number-of-commits>-<short-sha-of-latest-commit>"
    echo "                      --replace-version 0.0.<number-of-commits-1>-<short-sha-of-last-but-one-commit>"
    echo ""
    echo "                Variables overrides:"
    echo "                      QUAY_NAMESPACE  - If this variables is set then you don't have to use the --quay-namespace parameter."
    echo ""
    echo "                Optional variables:"
    echo "                      QUAY_AUTH_TOKEN - Quay authentication token to be used for pushing to the quay namespace. If not set, then it's taken from ~/.docker/config.json file."
    echo ""
    echo "Example:"
    echo "   ./scripts/push-to-quay-nightly.sh -pr ../host-operator"
    echo "          - This command will generate CSV, CRDs and package info with the values defined above for the host-operator project"
    echo "            and pushes it to quay namespace defined by either \"\${QUAY_NAMESPACE}\" variable or --quay-namespace parameter."
}

setup_version_variables() {
    # setup version and commit variables for the current repo
    GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD`
    PREVIOUS_GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD^`
    NUMBER_OF_COMMITS=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-list --count HEAD`

    # check if there is main repo or inner repo specified
    if [[ -n "${MAIN_REPO_URL}${EMBEDDED_REPO_URL}" ]]; then
        # if there is, then clone the latest version of the repo to /tmp dir
        if [[ -d ${OTHER_REPO_ROOT_DIR} ]]; then
            rm -rf ${OTHER_REPO_ROOT_DIR}
        fi
        mkdir -p ${OTHER_REPO_ROOT_DIR}
        git -C ${OTHER_REPO_ROOT_DIR} clone ${MAIN_REPO_URL}${EMBEDDED_REPO_URL}
        OTHER_REPO_PATH=${OTHER_REPO_ROOT_DIR}/`basename -s .git $(echo ${MAIN_REPO_URL}${EMBEDDED_REPO_URL})`

        # and set version and comit variables also for this repo
        OTHER_REPO_GIT_COMMIT_ID=`git --git-dir=${OTHER_REPO_PATH}/.git --work-tree=${OTHER_REPO_PATH} rev-parse --short HEAD`
        OTHER_REPO_NUMBER_OF_COMMITS=`git --git-dir=${OTHER_REPO_PATH}/.git --work-tree=${OTHER_REPO_PATH} rev-list --count HEAD`

        if [[ -n "${MAIN_REPO_URL}"  ]]; then
            # the other repo is main, so the number of commits and commit ID should be specified as the first one
            NEXT_CSV_VERSION="0.0.${OTHER_REPO_NUMBER_OF_COMMITS}-${NUMBER_OF_COMMITS}-commit-${OTHER_REPO_GIT_COMMIT_ID}-${GIT_COMMIT_ID}"
            REPLACE_CSV_VERSION="0.0.${OTHER_REPO_NUMBER_OF_COMMITS}-$((${NUMBER_OF_COMMITS}-1))-commit-${OTHER_REPO_GIT_COMMIT_ID}-${PREVIOUS_GIT_COMMIT_ID}"
        else
            # the other repo is inner, so the number of commits and commit ID should be specified as the second one
            NEXT_CSV_VERSION="0.0.${NUMBER_OF_COMMITS}-${OTHER_REPO_NUMBER_OF_COMMITS}-commit-${GIT_COMMIT_ID}-${OTHER_REPO_GIT_COMMIT_ID}"
            REPLACE_CSV_VERSION="0.0.$((${NUMBER_OF_COMMITS}-1))-${OTHER_REPO_NUMBER_OF_COMMITS}-commit-${PREVIOUS_GIT_COMMIT_ID}-${OTHER_REPO_GIT_COMMIT_ID}"
        fi
    else
        # there is no other repo specified - use the basic version format
        NEXT_CSV_VERSION="0.0.${NUMBER_OF_COMMITS}-commit-${GIT_COMMIT_ID}"
        REPLACE_CSV_VERSION="0.0.$((${NUMBER_OF_COMMITS}-1))-commit-${PREVIOUS_GIT_COMMIT_ID}"
    fi
}

# use the olm-setup as the source
OLM_SETUP_FILE=scripts/olm-setup.sh
if [[ -f ${OLM_SETUP_FILE} ]]; then
    source ${OLM_SETUP_FILE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/${OLM_SETUP_FILE})"
    fi
fi
# read argument to get project root dir
read_arguments $@

# setup version and commit variables
setup_version_variables

# generate manifests
count_images_and_generate_manifests $@ --channel nightly --template-version ${DEFAULT_VERSION} --next-version ${NEXT_CSV_VERSION} --replace-version ${REPLACE_CSV_VERSION}

# push manifests to quay
DIR_TO_PUSH=${PKG_DIR}
push_to_quay

# bring back the original operator package directory
rm -rf ${PKG_DIR}
cp -r ${PKG_DIR_BACKUP} ${PKG_DIR}