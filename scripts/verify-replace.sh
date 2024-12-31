#!/bin/bash
TMP_DIR=/tmp/
BASE_REPO_PATH=$(mktemp -d ${TMP_DIR}replace-verify.XXX)
GH_BASE_URL_KS=https://github.com/kubesaw/
GH_BASE_URL_CRT=https://github.com/codeready-toolchain/
declare -a REPOS=("${GH_BASE_URL_KS}ksctl" "${GH_BASE_URL_CRT}host-operator" "${GH_BASE_URL_CRT}member-operator" "${GH_BASE_URL_CRT}registration-service" "${GH_BASE_URL_CRT}toolchain-e2e" "${GH_BASE_URL_CRT}toolchain-common")
C_PATH=${PWD}
ERRORLIST=()

echo Initiating verify-replace on dependent repos
for repo in "${REPOS[@]}"
do
    echo =========================================================================================
    echo  
    echo                        "$(basename ${repo})"
    echo                                                                     
    echo =========================================================================================
    repo_path=${BASE_REPO_PATH}/$(basename ${repo})
    echo "Cloning repo in /tmp"
    git clone --depth=1 ${repo} ${repo_path}
    echo "Repo cloned successfully"
    cd ${repo_path}
    if ! make pre-verify; then
        ERRORLIST+="($(basename ${repo}))"
        continue
    fi
    echo "Initiating 'go mod replace' of current api version in dependent repos"
    go mod edit -replace github.com/codeready-toolchain/api=${C_PATH}
    make verify-dependencies || ERRORLIST+="($(basename ${repo}))"
    echo                                                          
    echo =========================================================================================
    echo 
done
if [ ${#ERRORLIST[@]} -ne 0 ]; then
    echo "Below are the repos with error: "
    for e in ${ERRORLIST[*]}
    do
        echo "${e}"
    done
    exit 1
else
    echo "No errors detected"
fi