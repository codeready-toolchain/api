#!/usr/bin/env bash

additional_help() {
    echo "create-release-bundle.sh creates an operator bundle of a given version in <project-root>/manifests directory"
    echo "Required parameters:"
    echo "              \"--project-root\"   to specify the root of the project"
    echo "              \"--next-version\"   to specify the version the operator bundle should be generated for"
    echo "Optional parameters:"
    echo "              \"--operator-name\"  to specify the name of the operator"
    echo "              \"--quay-namespace\" to specify quay namespace the operator bundle should be pushed to"
    echo "              \"--embedded-repo\"  to specify embedded repo that should be deployed through the main one"
    echo ""
    echo "Example:"
    echo "   ./scripts/create-release-bundle.sh -pr ../toolchain-operator/ -on codeready-toolchain-operator --next-version 0.1.0"
    echo "          - This command will generate CSV, CRDs and package info for version 0.1.0 in the project toolchain-operator,"
    echo "            it will use codeready-toolchain-operator as a name of the operator."
    echo "            The resulting manifest will be placed in ../toolchain-operator/manifests/0.1.0/ directory"

}

# use the olm-setup as the source
OLM_SETUP_FILE=scripts/olm-setup.sh
if [[ -f ${OLM_SETUP_FILE} ]]; then
    source ${OLM_SETUP_FILE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/{OLM_SETUP_FILE})"
    fi
fi
# read argument to get project root dir
read_arguments $@

# take the latest version to be replaced
if [[ -d ${MANIFESTS_DIR} ]]; then
    LAST_VERSION=`basename $(ls -d ${MANIFESTS_DIR}/*/ | sort | tail -1)`

    if [[ "${LAST_VERSION}" != "${NEXT_CSV_VERSION}" ]]; then
        REPLACE_LAST_VERSION_PARAM="--replace-version ${LAST_VERSION}"

    elif [[ `ls -d ${MANIFESTS_DIR}/*/ | wc -l` -ge 2 ]]; then
        LAST_VERSION=`basename $(ls -d ${MANIFESTS_DIR}/*/ | sort | tail -2 | head -n 1)`
        REPLACE_LAST_VERSION_PARAM="--replace-version ${LAST_VERSION}"
    fi
fi

GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short origin/master`
# generate manifests
check_main_and_embedded_repos_and_generate_manifests --template-version ${DEFAULT_VERSION} ${REPLACE_LAST_VERSION_PARAM}

copy_manifests_to_versioned_dir_and_adjust_package_file

# delete the default bundle directory that is used as a template
rm -rf ${BUNDLE_DIR}

# copy everything from package dir to manifests
if [[ ! -d ${MANIFESTS_DIR} ]]; then
    mkdir -p ${MANIFESTS_DIR}
fi
cp -r ${PKG_DIR}/* ${MANIFESTS_DIR}/

# bring back the original operator package directory
rm -rf ${PKG_DIR}
cp -r ${PKG_DIR_BACKUP} ${PKG_DIR}

# verify the manifests
operator-courier --verbose verify ${MANIFESTS_DIR}

# Vars for deploy hack
generate_deploy_hack -crds ${MANIFESTS_DIR}/${NEXT_CSV_VERSION} -csvs ${MANIFESTS_DIR}/ -pf ${MANIFESTS_DIR}/${OPERATOR_NAME}.package.yaml -hd ${TEMP_DIR}/hack_deploy_${OPERATOR_NAME}_${NEXT_CSV_VERSION} -on ${OPERATOR_NAME}