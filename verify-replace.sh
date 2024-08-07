#!/bin/bash
TMP_DIR=/tmp/
BASE_REPO_PATH=$(mktemp -d ${TMP_DIR}replace-verify.XXX)
GH_BASE_URL_KS=https://github.com/kubesaw/
GH_BASE_URL_CRT=https://github.com/codeready-toolchain/
GH_KSCTL=${GH_BASE_URL_KS}ksctl
GH_HOST=${GH_BASE_URL_CRT}host-operator
GH_MEMBER=${GH_BASE_URL_CRT}member-operator
GH_REGSVC=${GH_BASE_URL_CRT}registration-service
GH_E2E=${GH_BASE_URL_CRT}toolchain-e2e
GH_TC=${GH_BASE_URL_CRT}toolchain-common
C_PATH=${PWD}
ERRORLIST=()

echo Initiating verify-replace on dependent repos
for repo in ${GH_HOST} ${GH_REGSVC} ${GH_KSCTL} ${GH_MEMBER} ${GH_E2E} ${GH_TC}
do
    REPO_PATH=$BASE_REPO_PATH/$(basename $repo)
    echo Cloning repo in /tmp
    git clone --depth=1 $repo $REPO_PATH
    echo Repo cloned successfully
    cd $REPO_PATH
    echo Initiating 'go mod replace' of current toolchain common version in dependent repos
    go mod edit -replace github.com/codeready-toolchain/api=$C_PATH
    make verify-dependencies || ERRORLIST+=($(basename $repo))
done
if [ ${#ERRORLIST[@]} -ne 0 ]; then
    echo Below are the repos with error:
    for e in ${ERRORLIST[*]}
    do
        echo $e
    done
else
    echo No errors detected
fi