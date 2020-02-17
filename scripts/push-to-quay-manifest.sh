#!/usr/bin/env bash

additional_help() {
    echo "push-to-quay-manifest.sh pushes the latest version of operator bundle from <project-root>/manifest directory to quay.io"
    echo "Required parameters:"
    echo "              \"--project-root\"   to specify the root of the project"
    echo "Optional parameters:"
    echo "              \"--operator-name\"  to specify the name of the operator"
    echo ""
    echo "Example:"
    echo "   ./scripts/push-to-quay-manifest.sh -pr ../toolchain-operator -on codeready-toolchain-operator"
    echo "          - This command will push the latest version of the operator bundle of codeready-toolchain-operator from ../toolchain-operator/manifests directory to quay.io"

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

# setup version
NEXT_CSV_VERSION=`basename $(ls -d ${MANIFESTS_DIR}/*/ | sort | tail -1)`

# read final arguments and setup vars
read_arguments $@ --channel alpha --next-version ${NEXT_CSV_VERSION}
setup_variables

# push manifests to quay
DIR_TO_PUSH=${MANIFESTS_DIR}
push_to_quay
