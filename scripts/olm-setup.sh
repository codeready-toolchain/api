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

    replace_with_sed "${CSV_SED_REPLACE}" "${CSV_DIR}/*clusterserviceversion.yaml"

    echo "-> Bundle generated."
}

replace_with_sed() {
    TMP_CSV="/tmp/${OPERATOR_NAME}_${NEXT_CSV_VERSION}_replace-file"
    sed -e "$1" $2 > ${TMP_CSV}
    cat ${TMP_CSV} > $2
    rm -rf ${TMP_CSV}
}

generate_hack() {
    # Name and display name vars for CatalogSource
    HACK_DIR=${PRJ_ROOT_DIR}/hack
    echo "## Generating files for easy deployment and installation of project '${PRJ_NAME}' into ${HACK_DIR} ..."

    DISPLAYNAME=$(echo ${OPERATOR_NAME} | tr '-' ' ' | awk '{for (i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')

    # Create hack directory if is missing
    if [[ ! -d ${HACK_DIR} ]]; then
        mkdir ${HACK_DIR}
    fi

    # CatalogSource and ConfigMap for easy deployment
    echo "# This file was autogenerated by github.com/codeready-toolchain/api/olm-catalog.sh'
---
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: source-${OPERATOR_NAME}
  namespace: openshift-marketplace
spec:
  configMap: cm-${OPERATOR_NAME}
  displayName: $DISPLAYNAME
  publisher: Red Hat
  sourceType: internal
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: cm-${OPERATOR_NAME}
  namespace: openshift-marketplace
data:
  customResourceDefinitions: |-
$(for crd in `ls ${CRDS_DIR}/*.yaml`; do cat ${crd} | indent_list; done)
  clusterServiceVersions: |-
$(cat ${CSV_DIR}/*clusterserviceversion.yaml | indent_list | sed -e 's|^ *$||g')
  packages: |
$(cat ${PKG_FILE} | indent_list "packageName")" > ${HACK_DIR}/deploy_csv.yaml


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

indent_list() {
    local INDENT="      "
    sed -e "s/^/${INDENT}/;1s/^${INDENT}/${INDENT:0:${#INDENT}-2}- /"
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