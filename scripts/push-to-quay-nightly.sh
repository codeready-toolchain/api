#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-to-quay-nightly.sh scripts overrides several parameters and expects/uses only some of the other ones - see below."
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
    echo "                      --channel"
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

# generate release manifests
GENERATE_CD_RELEASE=scripts/generate-cd-release-manifests.sh
if [[ -f ${GENERATE_CD_RELEASE} ]]; then
    source ${GENERATE_CD_RELEASE}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${GENERATE_CD_RELEASE} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${GENERATE_CD_RELEASE}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/${GENERATE_CD_RELEASE})"
    fi
fi

# push the manifests to quay application
PUSH_TO_QUAY=scripts/push-manifests-as-app.sh
if [[ -f ${PUSH_TO_QUAY} ]]; then
    source ${PUSH_TO_QUAY}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${PUSH_TO_QUAY} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${PUSH_TO_QUAY}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/${PUSH_TO_QUAY})"
    fi
fi

# recover the operator directory containing operator bundle from the backup
RECOVER=scripts/recover-operator-dir.sh
if [[ -f ${RECOVER} ]]; then
    source ${RECOVER}
else
    if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${RECOVER} ]]; then
        source ${GOPATH}/src/github.com/codeready-toolchain/api/${RECOVER}
    else
        source /dev/stdin <<< "$(curl -sSL https://raw.githubusercontent.com/codeready-toolchain/api/master/${RECOVER})"
    fi
fi
