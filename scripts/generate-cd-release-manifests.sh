#!/usr/bin/env bash

additional_help() {
    echo "Important info: generate-release-manifest-for-cd.sh scripts overrides several parameters and expects/uses only some of the other ones - see below."
    echo ""
    echo "                The parameters written below are overridden with these values:"
    echo "                      --template-version ${DEFAULT_VERSION}"
    echo "                      --next-version 0.0.<number-of-commits>-<short-sha-of-latest-commit>"
    echo "                      --replace-version 0.0.<number-of-commits-1>-<short-sha-of-last-but-one-commit>"
    echo ""
    echo "                Expected parameters to be passed (if needed):"
    echo "                      --project-root"
    echo "                      --embedded-repo"
    echo "                      --quay-namespace"
    echo "                      --operator-name"
    echo ""
    echo "                Variables overrides:"
    echo "                      QUAY_NAMESPACE  - If this variables is set then you don't have to use the --quay-namespace parameter."
    echo ""
    echo "Example:"
    echo "   ./scripts/generate-release-manifest-for-cd.sh -pr ../host-operator"
    echo "          - This command will (re)generate CSV, CRDs in the manifests/ directory for the host-operator project"
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

# setup version variables based on commits so they can be used for generation process
setup_version_variables_based_on_commits true

# generate manifests
check_main_and_embedded_repos_and_generate_manifests $@ --template-version ${DEFAULT_VERSION} --next-version ${NEXT_CSV_VERSION} --replace-version ${REPLACE_CSV_VERSION}
