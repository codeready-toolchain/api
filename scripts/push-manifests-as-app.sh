#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-manifests-as-app.sh scripts use only some parameters, so use only these to specify necessary values:"
    echo "                      --project-root"
    echo "                      --embedded-repo"
    echo "                      --main-repo"
    echo "                      --quay-namespace"
    echo "                      --operator-name"
    echo "                      --channel"
    echo ""
    echo "                Variables overrides:"
    echo "                      QUAY_NAMESPACE  - If this variables is set then you don't have to use the --quay-namespace parameter."
    echo ""
    echo "                Optional variables:"
    echo "                      QUAY_AUTH_TOKEN - Quay authentication token to be used for pushing to the quay namespace. If not set, then it's taken from ~/.docker/config.json file."
    echo ""
    echo "Example:"
    echo "   ./scripts/push-manifests-as-app.sh -pr ../host-operator"
    echo "          - This command will copy manifests to versioned directory, modify the package file and push it all to quay namespace"
    echo "            defined by either \"\${QUAY_NAMESPACE}\" variable or --quay-namespace parameter."
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

# read arguments to get project root dir
read_arguments $@

# setup version variables based on commits so they can be used for pushing it to quay
setup_version_variables_based_on_commits false

# if the main repo is specified then reconfigure the variables so the project root points to the temp directory
if [[ -n "${MAIN_REPO_URL}"  ]]; then
    read_arguments $@ -pr ${OTHER_REPO_PATH}
fi

# copy manifestst to versioned directory and adjust the package file so it can be pushed to quay application
copy_manifests_to_versioned_dir_and_adjust_package_file

# push manifests to quay
DIR_TO_PUSH=${PKG_DIR}
push_manifests_as_app_to_quay