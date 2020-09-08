#!/usr/bin/env bash

additional_help() {
    echo "Important info: push-bundle-and-index-image.sh scripts use only some parameters, so use only these to specify necessary values:"
    echo "                      --project-root"
    echo "                      --embedded-repo"
    echo "                      --main-repo"
    echo "                      --quay-namespace"
    echo "                      --index-image"
    echo "                      --image-builder"
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
    echo "   ./scripts/push-bundle-and-index-image.sh -pr ../host-operator --image-builder podman --index-image hosted-toolchain-index"
    echo "          - This command generates bundle image (using 'staging' channel) for the host-operator project, adds it to the index image 'hosted-toolchain-index'"
    echo "            and pushes them to a repository in quay namespace defined by either \"\${QUAY_NAMESPACE}\" variable or --quay-namespace parameter."
}

handle_missing_version_in_combined_repo() {
    # split the version to get the number of commits and commit hashes as separated params (ie. from 0.0.217-105-commit-6c9926d-64af1be to 0.0.217 105 commit 6c9926d 64af1be)
    SPLIT_REPLACE_VERSION=(${REPLACE_VERSION_SUFFIX//-/ })
    # check if the number of split parts in the array is 5 to handle the missing version only for cases with embedded or main repos (host-operator + registration-service)
    if [[ ${#SPLIT_REPLACE_VERSION[@]} -eq 5 ]]; then
        # now we need to get the latest version that was added to the index - to do that we use "opm index export" command that pulls the index and all bundles that are added to the index
        # however, we don't want to pull all bundle images (it could take long time and consume a lot of resources) so we redirect the output of the command to a file
        # and then we will wait until opm unpacks the index image and prints out the list of bundle images that it is going to pull.
        # We expect that the last added bundle is the first in in the list
        echo "going to check the latest bundle image in the index image ..."
        EXPORT_OUTPUT_FILE=${TEMP_DIR}/export-output
        cd ${TEMP_DIR}
            opm index export --index=${FROM_INDEX_IMAGE} -c=${IMAGE_BUILDER} -o=${OPERATOR_NAME} >${EXPORT_OUTPUT_FILE} 2>&1 &
        cd ${CURRENT_DIR}
        # read the file containing the output and check if it contains "Preparing to pull bundles" clause. The timeout is 1 minute
        LAST_VERSION_IN_INDEX=""
        while [[ -z ${LAST_VERSION_IN_INDEX} ]] && [[ ${NEXT_WAIT_TIME} -lt 60 ]]; do
            # if the output contains "Preparing to pull bundles" clause then pick the first bundle image and take only the version suffix
            # ie. from:
            #     ... Preparing to pull bundles [\"quay.io/matousjobanek/host-operator-bundle:0.0.217-105-commit-25cbcfd-64af1be\" \"quay.io ...
            # take only 0.0.217-105-commit-25cbcfd-64af1be
            LAST_VERSION_IN_INDEX=$(grep "Preparing to pull bundles" ${EXPORT_OUTPUT_FILE} | sed  's|^.*\[\\"[^"]*bundle:\([^"]*\)\\".*$|\1|')
            echo "$(( NEXT_WAIT_TIME++ )). attempt of waiting for the latest bundle image in the index"
            sleep 1
        done
        # kill the opm index export process
        echo "the wait finished - killing the opm process"
        kill %1 || true

        # check if the latest bundle image version was found or not
        if [[ -n ${LAST_VERSION_IN_INDEX} ]]; then
            echo "latest bundle image version in index was found: ${LAST_VERSION_IN_INDEX}"

            # if it was found then compare if the expected replace version is the same as the latest one
            if [[ ${REPLACE_VERSION_SUFFIX} != ${LAST_VERSION_IN_INDEX} ]]; then
                echo "the replaces version \"${REPLACE_VERSION_SUFFIX}\" is NOT the same as the latest version in index \"${LAST_VERSION_IN_INDEX}\""

                # if the expected replace version is not the same as the latest one then let's check if new version that is going to be added is same as the latest one
                # if it is, then it means that the bundle image with that version was already added
                if [[ ${CURRENT_VERSION} != ${LAST_VERSION_IN_INDEX} ]]; then
                    echo "the next new version \"${CURRENT_VERSION}\" is NOT the same as the latest version in index \"${LAST_VERSION_IN_INDEX}\""

                    # if it's not the same, then it means that there is some version missing
                    # in this "if" statement we check if the latest image version has something in common with the one in "replaces" clause
                    if [[ -n $(echo ${LAST_VERSION_IN_INDEX} | grep "${SPLIT_REPLACE_VERSION[0]}-[0-9]*-commit-${SPLIT_REPLACE_VERSION[3]}.*") ]] || \
                    [[ -n $(echo ${LAST_VERSION_IN_INDEX} | grep "0.0.[0-9]*-${SPLIT_REPLACE_VERSION[1]}-commit-[^-]*-${SPLIT_REPLACE_VERSION[4]}") ]]; then
                        echo "we found something in common with the latest bundle image version in the index, so we will use the version \"${LAST_VERSION_IN_INDEX}\" for the replacement"

                        # if it has something in common then it would mean that there is most like only one version missing
                        # in that case let's replace the "replaces" clause with the latest image version and use that one
                        TEMP_CSV_REPLACE="${TEMP_DIR}/${PRJ_NAME}_${CURRENT_VERSION}_csv_replace.yaml"
                        sed "s/replaces: ${REPLACE_VERSION}$/replaces: ${OPERATOR_NAME}.v${LAST_VERSION_IN_INDEX}/" ${CSV_LOCATION} > ${TEMP_CSV_REPLACE}
                        mv ${TEMP_CSV_REPLACE} ${CSV_LOCATION}
                    else
                        echo "we didn't find anything in common in the the latest bundle image version, but we will continue and hope that it will magically work"
                    fi
                else
                    echo "the next new version \"${CURRENT_VERSION}\" IS the same as the latest version in index \"${LAST_VERSION_IN_INDEX}\""
                    echo "this means that the version has been already added to the index"
                    echo "exiting the CD process..."
                    exit 0
                fi
             else
                echo "the replaces version \"${REPLACE_VERSION_SUFFIX}\" IS the same as the latest version in index \"${LAST_VERSION_IN_INDEX}\""
                echo "that means that no version is missing in the index"
             fi
        else
            echo "the latest bundle image wasn't found in the opm output - we will continue and hope that everything will work fine"
            echo "see the opm output:"
            cat ${EXPORT_OUTPUT_FILE}
        fi
    else
        echo "the split replace version doesn't have 5 parts (${SPLIT_REPLACE_VERSION[@]}) so it means that it's not a combined operator (host-operator + registration-service)"
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

# if the main repo is specified then reconfigure the variables so the project root points to the temp directory
if [[ -n "${MAIN_REPO_URL}"  ]]; then
    REPO_NAME_WITH_GIT=$(basename $(echo ${MAIN_REPO_URL}))
    OTHER_REPO_PATH=${OTHER_REPO_ROOT_DIR}/${REPO_NAME_WITH_GIT%.*}
    read_arguments $@ -pr ${OTHER_REPO_PATH}
fi

# retrieve the current version
CSV_LOCATION="${BUNDLE_DIR}/*clusterserviceversion.yaml"
CURRENT_VERSION=`grep "^  version: " ${CSV_LOCATION} | awk '{print $2}'`

# set the image names variables
BUNDLE_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${PRJ_NAME}-bundle:${CURRENT_VERSION}
INDEX_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${INDEX_IMAGE_NAME}:latest
FROM_INDEX_IMAGE="${INDEX_IMAGE}"

if [[ "${INDEX_PER_COMMIT}" == "true" ]]; then
    INDEX_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${INDEX_IMAGE_NAME}:${CURRENT_VERSION}
fi

# get the current version that is in the "replaces" clause of the CSV
REPLACE_VERSION=`grep "^  replaces: " ${CSV_LOCATION} | awk '{print $2}'`
if [[ -n ${REPLACE_VERSION} ]]; then
    # get only the suffix from the replace version (ie. from toolchain-host-operator.v0.0.217-105-commit-6c9926d-64af1be get only 0.0.217-105-commit-6c9926d-64af1be)
    REPLACE_VERSION_SUFFIX=`echo ${REPLACE_VERSION} | sed -e "s/^.*operator.v//"`

    if [[ "${INDEX_PER_COMMIT}" == "true" ]]; then
        FROM_INDEX_IMAGE="quay.io/${QUAY_NAMESPACE_TO_PUSH}/${INDEX_IMAGE_NAME}:${REPLACE_VERSION_SUFFIX}"
    fi

    handle_missing_version_in_combined_repo
elif [[ "${INDEX_PER_COMMIT}" == "true" ]]; then
        FROM_INDEX_IMAGE=""
fi

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

if [[ -n ${FROM_INDEX_IMAGE} ]] && [[ `${IMAGE_BUILDER} pull ${FROM_INDEX_IMAGE}` ]]; then
    opm index add --bundles ${BUNDLE_IMAGE} --build-tool ${IMAGE_BUILDER} --tag ${INDEX_IMAGE} --from-index ${FROM_INDEX_IMAGE} ${PULL_TOOL_PARAM}
else
    opm index add --bundles ${BUNDLE_IMAGE} --build-tool ${IMAGE_BUILDER} --tag ${INDEX_IMAGE} ${PULL_TOOL_PARAM}
fi

if [[ ${IMAGE_BUILDER} == "buildah" ]]; then
    ${IMAGE_BUILDER} push ${INDEX_IMAGE} docker://${INDEX_IMAGE}
else
    ${IMAGE_BUILDER} push ${INDEX_IMAGE}
fi