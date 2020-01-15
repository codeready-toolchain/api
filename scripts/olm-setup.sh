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
    echo "-an, --allnamespaces     If set to true, then defines that the hack files should be created for AllNamespaces mode"
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
                -an|--allnamespaces)
                    shift
                    ALLNAMESPACES_MODE=$1
                    shift
                    ;;
                *)
                   echo "$1 is not a recognized flag!" >> /dev/stderr
                   user_help
                   exit -1
                   ;;
          esac
    done

    if [[ -z PRJ_ROOT_DIR ]]; then
        echo "--project-root parameter is not specified" >> /dev/stderr
        user_help
        exit 1;
    fi

    MANIFESTS_DIR=${PRJ_ROOT_DIR}/manifests
}

# Default version var - it has to be out of the function to make it available in help text
DEFAULT_VERSION=0.0.1

setup_variables() {
    # Version vars
    NEXT_CSV_VERSION=${NEXT_CSV_VERSION:-${DEFAULT_VERSION}}

    # Channel to be used
    CHANNEL=${CHANNEL:alpha}

    # Files and directories related vars
    PRJ_NAME=`basename ${PRJ_ROOT_DIR}`
    OPERATOR_NAME=${SET_OPERATOR_NAME:-toolchain-${PRJ_NAME}}
    CRDS_DIR=${PRJ_ROOT_DIR}/deploy/crds
    PKG_DIR=${PRJ_ROOT_DIR}/deploy/olm-catalog/${OPERATOR_NAME}
    PKG_FILE=${PKG_DIR}/${OPERATOR_NAME}.package.yaml
    CSV_DIR=${PKG_DIR}/${NEXT_CSV_VERSION}

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
    CURRENT_DIR=${PWD}
    cd ${PRJ_ROOT_DIR}
    operator-sdk olm-catalog gen-csv --csv-version ${NEXT_CSV_VERSION} --update-crds --operator-name ${OPERATOR_NAME} ${FROM_VERSION_PARAM} ${CHANNEL_PARAM}
    cd ${CURRENT_DIR}

    CURRENT_REPLACE_CLAUSE=`grep "replaces:" ${CSV_DIR}/*clusterserviceversion.yaml || true`
    if [[ -n "${REPLACE_VERSION}" ]]; then
        if [[ -n "${TEMPLATE_CSV_VERSION}" ]]; then
            CSV_SED_REPLACE+=";s/replaces: ${OPERATOR_NAME}.v${TEMPLATE_CSV_VERSION}/replaces: ${OPERATOR_NAME}.v${REPLACE_VERSION}/"
        else
            if [[ -n "${CURRENT_REPLACE_CLAUSE}" ]]; then
                CSV_SED_REPLACE+=";s/replaces: ${OPERATOR_NAME}.*$/replaces: ${OPERATOR_NAME}.v${REPLACE_VERSION}/"
            else
                CSV_SED_REPLACE+=";s/  version: ${NEXT_CSV_VERSION}/replaces: ${OPERATOR_NAME}.v${REPLACE_VERSION}\n  version: ${NEXT_CSV_VERSION}/"
            fi
        fi
    else
        if [[ -n "${CURRENT_REPLACE_CLAUSE}" ]]; then
            CSV_SED_REPLACE+="/${CURRENT_REPLACE_CLAUSE}$/d"
        fi
    fi
    if [[ -n "${IMAGE_IN_CSV}" ]]; then
        CSV_SED_REPLACE+=";s|REPLACE_IMAGE|${IMAGE_IN_CSV}|g;s|REPLACE_CREATED_AT|$(date -u +%FT%TZ)|g;"
    fi

    CSV_LOCATION=${CSV_DIR}/*clusterserviceversion.yaml
    replace_with_sed "${CSV_SED_REPLACE}" "${CSV_LOCATION}"

    if [[ -n "${IMAGE_IN_CSV}" ]]; then
        CONFIG_ENV_FILE=${PRJ_ROOT_DIR}/deploy/env/prod.yaml

        echo "enriching ${CSV_LOCATION} by params defined in ${CONFIG_ENV_FILE}"
        enrich-by-envs-from-yaml ${CSV_LOCATION} ${CONFIG_ENV_FILE}
    fi

    echo "-> Bundle generated."
}

enrich-by-envs-from-yaml() {
    ENRICHED_CSV="/tmp/${OPERATOR_NAME}_${NEXT_CSV_VERSION}-enriched-file"

    ENRICH_BY_ENVS_FROM_YAML=scripts/enrich-by-envs-from-yaml.sh
    if [[ -f ${ENRICH_BY_ENVS_FROM_YAML} ]]; then
        ${ENRICH_BY_ENVS_FROM_YAML} $@ > ${ENRICHED_CSV}
    else
        if [[ -f ${GOPATH}/src/github.com/codeready-toolchain/api/${ENRICH_BY_ENVS_FROM_YAML} ]]; then
            ${GOPATH}/src/github.com/codeready-toolchain/api/${ENRICH_BY_ENVS_FROM_YAML} $@ > ${ENRICHED_CSV}
        else
            curl -sSL  https://raw.githubusercontent.com/codeready-toolchain/api/master/scripts/generate-deploy-hack.sh | bash -s -- $@ > ${ENRICHED_CSV}
        fi
    fi
    cat ${ENRICHED_CSV} > $1
}

replace_with_sed() {
    TMP_CSV="/tmp/${OPERATOR_NAME}_${NEXT_CSV_VERSION}_replace-file"
    sed -e "$1" $2 > ${TMP_CSV}
    cat ${TMP_CSV} > $2
    rm -rf ${TMP_CSV}
}

generate_hack() {
    # Vars for deploy hack
    HACK_DIR=${PRJ_ROOT_DIR}/hack
    echo "## Generating files for easy deployment and installation of project '${PRJ_NAME}' into ${HACK_DIR} ..."

    generate_deploy_hack -crds ${CRDS_DIR} -csvs ${CSV_DIR} -pf ${PKG_FILE} -hd ${HACK_DIR} -on ${OPERATOR_NAME}

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
            curl -sSL  https://raw.githubusercontent.com/codeready-toolchain/api/master/scripts/generate-deploy-hack.sh | bash -s -- $@
        fi
    fi
}

generate_manifests() {
    #read arguments and setup variables
    read_arguments $@
    setup_variables

    # setup additional variables for pushing images
    IMAGE_IN_CSV=quay.io/${QUAY_NAMESPACE}/${PRJ_NAME}:${GIT_COMMIT_ID}

    # create backup of the current operator package directory
    PKG_DIR_BACKUP=/tmp/deploy_olm-catalog_${PRJ_NAME}_backup
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

push_to_quay() {
    RELEASE_BACKUP_DIR="/tmp/${OPERATOR_NAME}_${NEXT_CSV_VERSION}_${CHANNEL}"

    echo "## Pushing the OperatorHub package '${OPERATOR_NAME}' to the Quay.io '${QUAY_NAMESPACE}' organization ..."

    echo " - Copy package to backup folder: ${RELEASE_BACKUP_DIR}"

    rm -rf "${RELEASE_BACKUP_DIR}" > /dev/null 2>&1
    cp -r "${DIR_TO_PUSH}" ${RELEASE_BACKUP_DIR}

    echo " - Push flattened files to Quay.io namespace '${QUAY_NAMESPACE}' as version ${NEXT_CSV_VERSION}"

    if [[ -z ${QUAY_AUTH_TOKEN} ]]; then
        QUAY_AUTH_TOKEN=`cat ~/.docker/config.json | jq -r '.auths["quay.io"].auth'`
    fi

    operator-courier --verbose push ${RELEASE_BACKUP_DIR} "${QUAY_NAMESPACE}" "${OPERATOR_NAME}" "${NEXT_CSV_VERSION}" "basic ${QUAY_AUTH_TOKEN}"

    echo "-> Operator bundle pushed."
}