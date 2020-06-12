#!/usr/bin/env bash

additional_help() {
    echo "Important info: recover-operator-dir.sh scripts expects/uses only one parameter:"
    echo "                      --project-root"
    echo ""
    echo "Example:"
    echo "   ./scripts/recover-operator-dir.sh   -pr ../host-operator"
    echo "          - This command will recover the operator bundle directory from the backup folder stored in /tmp/."
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
setup_variables

# bring back the original operator package directory
rm -rf ${PKG_DIR}
cp -r ${PKG_DIR_BACKUP} ${PKG_DIR}