#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-to-quay-nightly.sh scripts overrides all the parameters but \"--project-root\", \"--embedded-repo\", \"--quay-namespace\" and \"--operator-name\", so use only these to specify necessary values."
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

# if the main repo is specified then reconfigure the variables so the project root points to the temp directory
if [[ -n "${MAIN_REPO_URL}"  ]]; then
    OTHER_REPO_PATH=${OTHER_REPO_ROOT_DIR}/`basename -s .git $(echo ${MAIN_REPO_URL})`
    read_arguments $@ -pr ${OTHER_REPO_PATH}
fi

# retrieve the current version
CURRENT_VERSION=`grep "^  version: " ${BUNDLE_DIR}/*clusterserviceversion.yaml | awk '{print $2}'`

# set the image names variables
BUNDLE_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${PRJ_NAME}-bundle:${CURRENT_VERSION}
INDEX_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${INDEX_IMAGE}:latest

# replace the channels in both the bundle.Dockerfile and annotation.yaml file
sed "s/\(channel.*: \)\".*\"/\1\"${CHANNEL}\"/" ${PKG_DIR}/metadata/annotations.yaml > ${TEMP_DIR}/${PRJ_NAME}_${CURRENT_VERSION}_metadata_annotations.yaml
mv ${TEMP_DIR}/${PRJ_NAME}_${CURRENT_VERSION}_metadata_annotations.yaml ${PKG_DIR}/metadata/annotations.yaml
sed "s/\(channel.*=\).*$/\1${CHANNEL}/" ${PKG_DIR}/bundle.Dockerfile > ${TEMP_DIR}/${PRJ_NAME}_${CURRENT_VERSION}_bundle.Dockerfile
mv ${TEMP_DIR}/${PRJ_NAME}_${CURRENT_VERSION}_bundle.Dockerfile ${PKG_DIR}/bundle.Dockerfile

# build and push the bundle image
if [[ ${IMAGE_BUILDER} == "buildah" ]]; then
    ${IMAGE_BUILDER} bud --layers -f ${PKG_DIR}/bundle.Dockerfile -t ${BUNDLE_IMAGE} ${PKG_DIR}/.
    ${IMAGE_BUILDER} push ${BUNDLE_IMAGE} docker://${BUNDLE_IMAGE}
else
    ${IMAGE_BUILDER} build -f ${PKG_DIR}/bundle.Dockerfile -t ${BUNDLE_IMAGE} ${PKG_DIR}/.
    ${IMAGE_BUILDER} push ${BUNDLE_IMAGE}
fi

# add manifests to the bundle image
cd ${PKG_DIR}
opm alpha bundle build --image-builder ${IMAGE_BUILDER} --directory ./manifests/ -t ${BUNDLE_IMAGE} -p ${OPERATOR_NAME} -c ${CHANNEL} -e ${CHANNEL}

if [[ ${IMAGE_BUILDER} == "buildah" ]]; then
    ${IMAGE_BUILDER} push ${BUNDLE_IMAGE} docker://${BUNDLE_IMAGE}
else
    ${IMAGE_BUILDER} push ${BUNDLE_IMAGE}
fi
cd ${CURRENT_DIR}

if [[ ${IMAGE_BUILDER} == "podman" ]]; then
    PULL_TOOL_PARAM="--pull-tool podman"
fi

opm index add --bundles ${BUNDLE_IMAGE} --build-tool ${IMAGE_BUILDER} --tag ${INDEX_IMAGE} --from-index ${INDEX_IMAGE} ${PULL_TOOL_PARAM}
if [[ ${IMAGE_BUILDER} == "buildah" ]]; then
    ${IMAGE_BUILDER} push ${INDEX_IMAGE} docker://${INDEX_IMAGE}
else
    ${IMAGE_BUILDER} push ${INDEX_IMAGE}
fi