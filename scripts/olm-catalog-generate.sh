#!/usr/bin/env bash

OLM_SETUP_FILE=./scripts/olm-setup.sh

additional_help() {
    echo "Examples:"
    echo "   ./scripts/olm-catalog-generate.sh -pr ../host-operator"
    echo "          - This command will (re)generate CSV, CRDs and package info with a default version \"${DEFAULT_VERSION}\" for host-operator project"
    echo ""
    echo "   ./scripts/olm-catalog-generate.sh -nv 0.2.0 -pr ../host-operator"
    echo "          - This command will (re)generate CSV, CRDs and package info with the version 0.2.0 for host-operator project"
    echo ""
    echo "   ./scripts/olm-catalog-generate.sh -ch beta -tv 0.0.1 -nv 0.2.0 -rv 0.1.0 -pr ../host-operator   "
    echo "          - This command will (re)generate CSV, CRDs and package info for host-operator project with the version 0.2.0 and channel beta,"
    echo "            as a base version for the generation will use CSV 0.0.1 and will add into the new CSV a replace clause for the version 0.1.0"
}

if [[ -f ${OLM_SETUP_FILE} ]]; then
    source ${OLM_SETUP_FILE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${OLM_SETUP_FILE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/scripts/olm-setup.sh)"
    fi
fi

read_arguments $@
setup_variables
generate_bundle
generate_hack

