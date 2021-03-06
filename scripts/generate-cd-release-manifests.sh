#!/usr/bin/env bash

# use the olm-setup as the source
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

# read argument to get project root dir
read_arguments $@
set -ex
echo "arguments read"

# setup version variables based on commits so they can be used for generation process
setup_version_variables_based_on_commits

echo "versions configured"

# generate manifests
check_main_and_embedded_repos_and_generate_manifests $@ --next-version ${NEXT_CSV_VERSION}
