#!/usr/bin/env bash

# Exit on error
set -e

user_help () {
    echo "Generate ClusterServiceVersion and additional deployment files for openshift-marketplace"
    echo "options:"
    echo "-pr, --project-root      Path to the root of the project the CSV should be generated for/in"
    echo "-tv, --template-version  CSV version that should be used as a base for the creation of the new version"
    echo "-nv, --next-version      Semantic version of the new CSV to be created"
    echo "-rv, --replace-version   The CSV version to be replaced by the new version (this param has to be specified even if it's same as template-version)"
    echo "-ch, --channel           Channel to be used for the CSV in the package manifest"
    echo "-on, --operator-name     Name of the operator - by default it uses toolchain-{repository_name}"
    echo "-mr, --main-repo         URL of the GH repo that should be used as the main repo (for CD). The current repo should be embedded in the main one. The operator bundle should be taken from the main repository (example of the main repo: https://github.com/codeready-toolchain/host-operator)"
    echo "-er, --embedded-repo     URL of the GH repo that should be used as the embedded repo (for CD). The repository should be embedded in the current repo. The operator bundle should be taken from the current repository (example of the embedded repo: https://github.com/codeready-toolchain/registration-service)"
    echo "-an, --allnamespaces     If set to true, then defines that the hack files should be created for AllNamespaces mode"
    echo "-qn, --quay-namespace    Specify the quay namespace the CSV should be pushed to - if not used then it uses the one stored in \"\${QUAY_NAMESPACE}\" variable"
    echo "-td, --temp-dir          Directory that should be used for storing temporal files - by default '/tmp' is used"
    echo "-ib, --image-builder     Tool to build container images - will be used by opm. One of: [docker, podman] (default "docker")"
    echo "-im, --index-image       Name of the index image where the bundle image should be added - when building & pushing operator bundle as an image."
    echo "-ic, --index-per-commit  If set to true, then the script will build&push unique index image for every release/commit."
    echo "-fr, --first-release     If set to true, then it will generate CSV without replaces clause."
    echo "-iu, --index-image-url   The whole url of the next index image (including the tag)."
    echo "-fu, --from-index-url    The whole url of the index image the script should build from (including the tag)."
    echo "-ci, --component-image   The name of the image to be used as a component of this operator."
    echo "-h,  --help              To show this help text"
    echo ""
    additional_help 2>/dev/null || true
    exit 0
}

read_arguments() {
    if [[ $# -lt 2 ]]
    then
        user_help
    fi

    while test $# -gt 0; do
           case "$1" in
                -h|--help)
                    user_help
                    ;;
                -pr|--project-root)
                    shift
                    PRJ_ROOT_DIR=$1
                    shift
                    ;;
                -tv|--template-version)
                    shift
                    TEMPLATE_CSV_VERSION=$1
                    shift
                    ;;
                -nv|--next-version)
                    shift
                    NEXT_CSV_VERSION=$1
                    shift
                    ;;
                -rv|--replace-version)
                    shift
                    REPLACE_VERSION=$1
                    shift
                    ;;
                -ch|--channel)
                    shift
                    CHANNEL=$1
                    shift
                    ;;
                -on|--operator-name)
                    shift
                    SET_OPERATOR_NAME=$1
                    shift
                    ;;
                -mr|--main-repo)
                    shift
                    MAIN_REPO_URL=$1
                    shift
                    ;;
                -er|--embedded-repo)
                    shift
                    EMBEDDED_REPO_URL=$1
                    shift
                    ;;
                -an|--allnamespaces)
                    shift
                    ALLNAMESPACES_MODE=$1
                    shift
                    ;;
                -qn|--quay-namespace)
                    shift
                    QUAY_NAMESPACE_TO_PUSH=$1
                    shift
                    ;;
                -td|--temp-dir)
                    shift
                    TEMP_DIR=$1
                    shift
                    ;;
                -ib|--image-builder)
                    shift
                    IMAGE_BUILDER=$1
                    shift
                    ;;
                -im|--index-image)
                    shift
                    INDEX_IMAGE_NAME=$1
                    shift
                    ;;
                -ic|--index-per-commit)
                    shift
                    INDEX_PER_COMMIT=$1
                    shift
                    ;;
                -fr|--first-release)
                    shift
                    FIRST_RELEASE=$1
                    shift
                    ;;
                -iu|--index-image-url)
                    shift
                    INDEX_IMAGE_URL=$1
                    shift
                    ;;
                -fu|--from-index-url)
                    shift
                    FROM_INDEX_URL=$1
                    shift
                    ;;
                -ci|--component-image)
                    shift
                    COMPONENT_IMAGE=$1
                    shift
                    ;;
                *)
                   echo "$1 is not a recognized flag!" >> /dev/stderr
                   user_help
                   exit -1
                   ;;
          esac
    done

    if [[ -z ${PRJ_ROOT_DIR} ]]; then
        echo "--project-root parameter is not specified" >> /dev/stderr
        user_help
        exit 1;
    fi

    cd ${PRJ_ROOT_DIR}
    PRJ_ROOT_DIR=${PWD}
    cd - > /dev/null

    if [[ -n "${EMBEDDED_REPO_URL}" ]] && [[ -n "${MAIN_REPO_URL}" ]]; then
        echo "you cannot specify both parameters '--main-repo' and '--embedded-repo' at the same time - use only one" >> /dev/stderr
        user_help
        exit 1
    fi

    if [[ -z ${QUAY_NAMESPACE_TO_PUSH} ]]; then
        QUAY_NAMESPACE_TO_PUSH=${QUAY_NAMESPACE:codeready-toolchain}
    fi

    MANIFESTS_DIR=${PRJ_ROOT_DIR}/manifests

    setup_variables
}

# Default version var - it has to be out of the function to make it available in help text
DEFAULT_VERSION=0.0.1

setup_variables() {
    # Version vars
    NEXT_CSV_VERSION=${NEXT_CSV_VERSION:-${DEFAULT_VERSION}}

    # Channel to be used
    CHANNEL=${CHANNEL:alpha}

    # Temporal directory
    TEMP_DIR=${TEMP_DIR:-/tmp}
    if [[ "${TEMP_DIR}" != "/tmp" ]]; then
        mkdir -p ${TEMP_DIR} || true
    fi
    OTHER_REPO_ROOT_DIR=${TEMP_DIR}/cd/other-repo

    # Image builder
    IMAGE_BUILDER=${IMAGE_BUILDER:-"docker"}

    # Files and directories related vars
    PRJ_NAME=`basename ${PRJ_ROOT_DIR}`
    OPERATOR_NAME=${SET_OPERATOR_NAME:-toolchain-${PRJ_NAME}}
    CRDS_DIR=${PRJ_ROOT_DIR}/deploy/crds
    PKG_DIR=${PRJ_ROOT_DIR}/deploy/olm-catalog/${OPERATOR_NAME}
    PKG_FILE=${PKG_DIR}/${OPERATOR_NAME}.package.yaml
    BUNDLE_DIR=${PKG_DIR}/manifests
    PKG_DIR_BACKUP=${TEMP_DIR}/deploy_olm-catalog_${PRJ_NAME}_backup
    CURRENT_DIR=${PWD}

    export GO111MODULE=on
}

generate_bundle() {
    # Generate CSV
    if [[ -n "${TEMPLATE_CSV_VERSION}" ]]; then
        FROM_VERSION_PARAM="--from-version ${TEMPLATE_CSV_VERSION}"
    fi
    if [[ -n "${CHANNEL}" ]]; then
        CHANNEL_PARAM="--csv-channel ${CHANNEL}"
    fi

    echo "## Generating operator bundle of project '${PRJ_NAME}' ..."

    # check if pkg/apis/toolchain/v1alpha1/ folder is available, if yes then run "operator-sdk generate csv" without pointing to specific dir as sources of api types
    if [[ -d "${PRJ_ROOT_DIR}/pkg/apis/toolchain/v1alpha1" ]]; then
        echo "  - running 'operator-sdk generate csv' using the local api types"
        cd ${PRJ_ROOT_DIR}
        operator-sdk generate csv --verbose --output-dir ${PKG_DIR} --csv-version ${NEXT_CSV_VERSION} --update-crds --operator-name ${OPERATOR_NAME} ${FROM_VERSION_PARAM} ${CHANNEL_PARAM}
        cd ${CURRENT_DIR}
    else
        # We have to run operator-sdk generate from the codeready-toolchain/api repo so it can reach the api source code to scan annotations
        # So, we either use local codeready-toolchain/api repo to or clone the repo from GitHub

        # check if the script directory is api repository directory - contains ../cmd/manager/main.go
        # if it is, then copy the directory to the temporary one, if not then clone the repo there
        SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
        if [[ -f ${SCRIPT_DIR}/../cmd/manager/main.go ]]; then
            echo "  - using local codeready-toolchain/api repo from ${SCRIPT_DIR}"
            API_REPO_DIR=${SCRIPT_DIR}/..
        else
            GENERATE_BUNDLE_TMP_DIR="${TEMP_DIR}/generate_bundle"
            API_TMP_DIR="${GENERATE_BUNDLE_TMP_DIR}/api"
            rm -rf ${GENERATE_BUNDLE_TMP_DIR} > /dev/null || true
            mkdir -p ${GENERATE_BUNDLE_TMP_DIR}
            echo "  - cloning codeready-toolchain/api repo to ${API_TMP_DIR}"
            git clone https://github.com/codeready-toolchain/api.git ${API_TMP_DIR}
            API_REPO_DIR=${API_TMP_DIR}
        fi

        cd ${API_REPO_DIR}
        echo "  - running 'operator-sdk generate csv' command inside of the codeready-toolchain/api directory '${API_REPO_DIR}'"
        operator-sdk generate csv --verbose --apis-dir api/ --output-dir ${PKG_DIR} --deploy-dir ${PRJ_ROOT_DIR}/deploy --csv-version ${NEXT_CSV_VERSION} --update-crds --operator-name ${OPERATOR_NAME} ${FROM_VERSION_PARAM} ${CHANNEL_PARAM}
        cd ${CURRENT_DIR}
    fi

    if [[ -n "${REPLACE_VERSION}" ]]; then
        NEW_REPLACE_CLAUSE="replaces: ${OPERATOR_NAME}.v${REPLACE_VERSION}"
    fi

    CURRENT_REPLACE_CLAUSE=`grep "replaces:" ${BUNDLE_DIR}/*clusterserviceversion.yaml || true`
    if [[ -n "${CURRENT_REPLACE_CLAUSE}" ]]; then
        if [[ -n "${TEMPLATE_CSV_VERSION}" ]]; then
            CSV_SED_REPLACE+=";s/replaces: ${OPERATOR_NAME}.v${TEMPLATE_CSV_VERSION}/${NEW_REPLACE_CLAUSE}/"
        else
            CSV_SED_REPLACE+=";s/replaces: ${OPERATOR_NAME}.*$/${NEW_REPLACE_CLAUSE}/"
        fi
    elif [[ -n "${NEW_REPLACE_CLAUSE}" ]]; then
            CSV_SED_REPLACE+=";s/  version: ${NEXT_CSV_VERSION}/  ${NEW_REPLACE_CLAUSE}\n  version: ${NEXT_CSV_VERSION}/"
    fi

    if [[ -n "${IMAGE_IN_CSV}" ]]; then
        IMAGE_IN_CSV_DIGEST_FORMAT=`get_digest_format ${IMAGE_IN_CSV}`
        CSV_SED_REPLACE+=";s|REPLACE_IMAGE|${IMAGE_IN_CSV_DIGEST_FORMAT}|g;s|REPLACE_CREATED_AT|$(date -u +%FT%TZ)|g;"
    fi
    if [[ -n "${EMBEDDED_REPO_IMAGE}" ]]; then
        EMBEDDED_REPO_IMAGE_DIGEST_FORMAT=`get_digest_format ${EMBEDDED_REPO_IMAGE}`
        CSV_SED_REPLACE+=";s|${EMBEDDED_REPO_REPLACEMENT}|${EMBEDDED_REPO_IMAGE_DIGEST_FORMAT}|g;"
    fi
    if [[ -n ${COMPONENT_IMAGE} ]]; then
        COMPONENT_IMAGE_URL=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${COMPONENT_IMAGE}:${GIT_COMMIT_ID}
        COMPONENT_IMAGE_DIGEST_FORMAT=`get_digest_format ${COMPONENT_IMAGE_URL}`
        COMPONENT_IMAGE_REPLACEMENT=REPLACE_$(echo ${COMPONENT_IMAGE} | awk '{ print toupper($0) }' | tr '-' '_')_IMAGE
        CSV_SED_REPLACE+=";s|${COMPONENT_IMAGE_REPLACEMENT}|${COMPONENT_IMAGE_DIGEST_FORMAT}|g;"
    fi
    if [[ "${CHANNEL}" == "nightly" ]]; then
        CSV_SED_REPLACE+=";s|  annotations:|  annotations:\n    olm.skipRange: '<${NEXT_CSV_VERSION}'|g;"
    fi
    CSV_LOCATION=${BUNDLE_DIR}/*clusterserviceversion.yaml
    replace_with_sed "${CSV_SED_REPLACE}" "${CSV_LOCATION}"

    if [[ -n "${IMAGE_IN_CSV}" ]]; then
        CONFIG_ENV_FILE=${PRJ_ROOT_DIR}/deploy/env/prod.yaml

        echo "enriching ${CSV_LOCATION} by params defined in ${CONFIG_ENV_FILE}"
        enrich-by-envs-from-yaml ${CSV_LOCATION} ${CONFIG_ENV_FILE}
    fi

    echo "-> Bundle generated."
}

get_digest_format() {
    IMG=$1
    IMG_LOC=`echo ${IMG} | cut -d: -f1`

    IMG_ORG=`echo ${IMG_LOC} | awk -F/ '{print $2}'`
    IMG_NAME=`echo ${IMG_LOC} | awk -F/ '{print $3}'`
    IMG_TAG=`echo ${IMG} | cut -d: -f2`

    echo "Getting digest of the image ${IMG}" >> /dev/stderr

    while [[ -z ${IMG_DIGEST} || "${IMG_DIGEST}" == "null" ]]; do
		if [[ ${NEXT_WAIT_TIME} -eq 10 ]]; then
		   echo " the digest of the image ${IMG} wasn't found" >> /dev/stderr
		   exit 1
		fi
		echo -n "." >> /dev/stderr
		(( NEXT_WAIT_TIME++ ))
		sleep 1
		IMG_DIGEST=`curl https://quay.io/api/v1/repository/${IMG_ORG}/${IMG_NAME} 2>/dev/null | jq -r ".tags.\"${IMG_TAG}\".manifest_digest"`
	done
    echo " found: ${IMG_DIGEST}" >> /dev/stderr

    echo ${IMG_LOC}@${IMG_DIGEST}
}


enrich-by-envs-from-yaml() {
    ENRICHED_CSV="${TEMP_DIR}/${OPERATOR_NAME}_${NEXT_CSV_VERSION}-enriched-file"

    ENRICH_BY_ENVS_FROM_YAML=scripts/enrich-by-envs-from-yaml.sh
    if [[ -f ${ENRICH_BY_ENVS_FROM_YAML} ]]; then
        ${ENRICH_BY_ENVS_FROM_YAML} $@ > ${ENRICHED_CSV}
    else
        if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${ENRICH_BY_ENVS_FROM_YAML} ]]; then
            ${GOPATH}/src/github.com/codeready-toolchain/api/${ENRICH_BY_ENVS_FROM_YAML} $@ > ${ENRICHED_CSV}
        else
            curl -sSL  https://raw.githubusercontent.com/codeready-toolchain/api/master/${ENRICH_BY_ENVS_FROM_YAML} | bash -s -- $@ > ${ENRICHED_CSV}
        fi
    fi
    cat ${ENRICHED_CSV} > $1
}

replace_with_sed() {
    TMP_CSV="${TEMP_DIR}/${OPERATOR_NAME}_${NEXT_CSV_VERSION}_replace-file"
    sed -e "$1" $2 > ${TMP_CSV}
    cat ${TMP_CSV} > $2
    rm -rf ${TMP_CSV}
}

generate_hack() {
    # Vars for deploy hack
    HACK_DIR=${PRJ_ROOT_DIR}/hack
    echo "## Generating files for easy deployment and installation of project '${PRJ_NAME}' into ${HACK_DIR} ..."

    generate_deploy_hack -crds ${CRDS_DIR} -csvs ${BUNDLE_DIR} -pf ${PKG_FILE} -hd ${HACK_DIR} -on ${OPERATOR_NAME}

    echo "# This file was autogenerated by github.com/codeready-toolchain/api/olm-catalog.sh'" > ${HACK_DIR}/install_operator.yaml
    if [[ "${ALLNAMESPACES_MODE}" != "true" ]]; then
        SUBSCRIPTION_NS="REPLACE_NAMESPACE"
        echo "---
apiVersion: operators.coreos.com/v1
kind: OperatorGroup
metadata:
  name: og-${OPERATOR_NAME}
  namespace: REPLACE_NAMESPACE
spec:
  targetNamespaces:
  - REPLACE_NAMESPACE
---" >> ${HACK_DIR}/install_operator.yaml
    else
        SUBSCRIPTION_NS="openshift-operators"
    fi

    echo "apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: subscription-${OPERATOR_NAME}
  namespace: ${SUBSCRIPTION_NS}
spec:
  channel: alpha
  installPlanApproval: Automatic
  name: ${OPERATOR_NAME}
  source: source-${OPERATOR_NAME}
  sourceNamespace: openshift-marketplace
  startingCSV: ${OPERATOR_NAME}.v0.0.1" >> ${HACK_DIR}/install_operator.yaml

  echo "-> Hack files generated."
}

generate_deploy_hack() {
    GENERATE_DEPLOY_HACK_FILE=scripts/generate-deploy-hack.sh
    if [[ -f ${GENERATE_DEPLOY_HACK_FILE} ]]; then
        ${GENERATE_DEPLOY_HACK_FILE} $@
    else
        if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${GENERATE_DEPLOY_HACK_FILE} ]]; then
            ${GOPATH}/src/github.com/codeready-toolchain/api/${GENERATE_DEPLOY_HACK_FILE} $@
        else
            curl -sSL  https://raw.githubusercontent.com/codeready-toolchain/api/master/${GENERATE_DEPLOY_HACK_FILE} | bash -s -- $@
        fi
    fi
}

# it takes one boolean parameter - if the other repo (either embedded or main one) should be cloned or not
setup_version_variables_based_on_commits() {
    # setup version and commit variables for the current repo
    GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD`
    PREVIOUS_GIT_COMMIT_ID=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-parse --short HEAD^`
    NUMBER_OF_COMMITS=`git --git-dir=${PRJ_ROOT_DIR}/.git --work-tree=${PRJ_ROOT_DIR} rev-list --count HEAD`

    # check if there is main repo or inner repo specified
    if [[ -n "${MAIN_REPO_URL}${EMBEDDED_REPO_URL}" ]]; then
        if [[ "true" == "$1" ]]; then
            # if there is, then clone the latest version of the repo to ${TEMP_DIR} dir
            if [[ -d ${OTHER_REPO_ROOT_DIR} ]]; then
                rm -rf ${OTHER_REPO_ROOT_DIR}
            fi
            mkdir -p ${OTHER_REPO_ROOT_DIR}
            git -C ${OTHER_REPO_ROOT_DIR} clone ${MAIN_REPO_URL}${EMBEDDED_REPO_URL}
        fi
        OTHER_REPO_PATH=${OTHER_REPO_ROOT_DIR}/`basename -s .git $(echo ${MAIN_REPO_URL}${EMBEDDED_REPO_URL})`

        # and set version and comit variables also for this repo
        OTHER_REPO_GIT_COMMIT_ID=`git --git-dir=${OTHER_REPO_PATH}/.git --work-tree=${OTHER_REPO_PATH} rev-parse --short HEAD`
        OTHER_REPO_NUMBER_OF_COMMITS=`git --git-dir=${OTHER_REPO_PATH}/.git --work-tree=${OTHER_REPO_PATH} rev-list --count HEAD`

        if [[ -n "${MAIN_REPO_URL}"  ]]; then
            # the other repo is main, so the number of commits and commit ID should be specified as the first one
            NEXT_CSV_VERSION="0.0.${OTHER_REPO_NUMBER_OF_COMMITS}-${NUMBER_OF_COMMITS}-commit-${OTHER_REPO_GIT_COMMIT_ID}-${GIT_COMMIT_ID}"
            REPLACE_CSV_VERSION="0.0.${OTHER_REPO_NUMBER_OF_COMMITS}-$((${NUMBER_OF_COMMITS}-1))-commit-${OTHER_REPO_GIT_COMMIT_ID}-${PREVIOUS_GIT_COMMIT_ID}"
        else
            # the other repo is inner, so the number of commits and commit ID should be specified as the second one
            NEXT_CSV_VERSION="0.0.${NUMBER_OF_COMMITS}-${OTHER_REPO_NUMBER_OF_COMMITS}-commit-${GIT_COMMIT_ID}-${OTHER_REPO_GIT_COMMIT_ID}"
            REPLACE_CSV_VERSION="0.0.$((${NUMBER_OF_COMMITS}-1))-${OTHER_REPO_NUMBER_OF_COMMITS}-commit-${PREVIOUS_GIT_COMMIT_ID}-${OTHER_REPO_GIT_COMMIT_ID}"
        fi
    else
        # there is no other repo specified - use the basic version format
        NEXT_CSV_VERSION="0.0.${NUMBER_OF_COMMITS}-commit-${GIT_COMMIT_ID}"
        REPLACE_CSV_VERSION="0.0.$((${NUMBER_OF_COMMITS}-1))-commit-${PREVIOUS_GIT_COMMIT_ID}"
    fi
}

check_main_and_embedded_repos_and_generate_manifests() {
    #read arguments and setup variables
    read_arguments $@
    setup_variables

    IMAGE_IN_CSV=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${PRJ_NAME}:${GIT_COMMIT_ID}
    # check if there is main repo or inner repo specified
    if [[ -n ${MAIN_REPO_URL}${EMBEDDED_REPO_URL}  ]] && [[ -n ${OTHER_REPO_GIT_COMMIT_ID} ]]; then

        OTHER_REPO_NAME=`basename -s .git $(echo ${MAIN_REPO_URL}${EMBEDDED_REPO_URL})`

        if [[ -n "${MAIN_REPO_URL}"  ]]; then
            IMAGE_IN_CSV=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${OTHER_REPO_NAME}:${OTHER_REPO_GIT_COMMIT_ID}

            EMBEDDED_REPO_REPLACEMENT=REPLACE_$(echo ${PRJ_NAME} | awk '{ print toupper($0) }' | tr '-' '_')_IMAGE
            EMBEDDED_REPO_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${PRJ_NAME}:${GIT_COMMIT_ID}
            generate_manifests $@ -pr ${OTHER_REPO_PATH}
        else
            EMBEDDED_REPO_REPLACEMENT=REPLACE_$(echo ${OTHER_REPO_NAME} | awk '{ print toupper($0) }' | tr '-' '_')_IMAGE
            EMBEDDED_REPO_IMAGE=quay.io/${QUAY_NAMESPACE_TO_PUSH}/${OTHER_REPO_NAME}:${OTHER_REPO_GIT_COMMIT_ID}
            generate_manifests $@
        fi
    else
        generate_manifests $@
    fi
}

generate_manifests() {
    #read arguments and setup variables
    read_arguments $@
    setup_variables

    # create backup of the current operator package directory
    if [[ -d ${PKG_DIR_BACKUP} ]]; then
        rm -rf ${PKG_DIR_BACKUP}
    fi
    cp -r ${PKG_DIR} ${PKG_DIR_BACKUP}

    # copy everything from manifests dir to package dir
    if [[ -d ${MANIFESTS_DIR} ]]; then
        cp -r ${MANIFESTS_DIR}/* ${PKG_DIR}/
    fi

    # generate the bundle
    generate_bundle
}
