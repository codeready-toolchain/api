#!/bin/bash
TMP_DIR=/tmp/
BASE_REPO_PATH=$(mktemp -d ${TMP_DIR}replace-verify.XXX)
GH_BASE_URL_KS=https://github.com/kubesaw/
GH_BASE_URL_CRT=https://github.com/codeready-toolchain/
declare -a REPOS=("${GH_BASE_URL_KS}ksctl" "${GH_BASE_URL_CRT}host-operator" "${GH_BASE_URL_CRT}member-operator" "${GH_BASE_URL_CRT}registration-service" "${GH_BASE_URL_CRT}toolchain-e2e" "${GH_BASE_URL_CRT}toolchain-common")
C_PATH=${PWD}
ERRORREPOLIST=()
ERRORFILELIST=()
STDOUTFILELIST=()
GOLINTREGEX="[\s\w.\/]*:[0-9]*:[0-9]*:[\w\s)(*.\`]*"
LINTERERRORFILE=$(mktemp ${TMP_DIR}LinterError.XXX)

echo Initiating verify-replace on dependent repos
for repo in "${REPOS[@]}"
do
    echo =========================================================================================
    echo  
    echo                        "$(basename ${repo})"
    echo                                                                     
    echo =========================================================================================                                            
    repo_path=${BASE_REPO_PATH}/$(basename ${repo})
    ERRFILE=$(mktemp ${TMP_DIR}$(basename ${repo})-error.XXX)
    STDOUTFILE=$(mktemp ${TMP_DIR}$(basename ${repo})-output.XXX)
    echo "Cloning repo in /tmp"
    git clone --depth=1 ${repo} ${repo_path}
    echo "Repo cloned successfully"
    cd ${repo_path}
    if ! make pre-verify; then
        ERRORREPOLIST+="$(basename ${repo}) "
        continue
    fi
    echo "Initiating 'go mod replace' of current api version in dependent repos"
    go mod edit -replace github.com/codeready-toolchain/api=${C_PATH}
    make verify-dependencies 2> ${ERRFILE} 1> ${STDOUTFILE}
    rc=$?
    if [ ${rc} -ne 0 ]; then
    ERRORREPOLIST+="$(basename ${repo}) " 
    ERRORFILELIST+="${ERRFILE}  "
    STDOUTFILELIST="${STDOUTFILE} "
    fi
    echo                                                          
    echo =========================================================================================
    echo                                   
done
echo                "Summary"
if [ ${#ERRORREPOLIST[@]} -ne 0 ]; then
    echo "Below are the repos with error: "
    for e in ${ERRORREPOLIST[*]}
    do
        for c in ${ERRORFILELIST[*]}
        do
            if [[ ${c} =~ ${e} ]]; then 
                echo                                                          
                echo =========================================================================================
                echo 
                echo                       "${e} has the following errors "
                echo                                                          
                echo =========================================================================================
                echo 
                cat "${c}"
            else
                for s in ${STDOUTFILELIST[*]}
                do
                    if [[ ${s} =~ ${e} ]]; then 
                        cat "${s}" | grep -E ${GOLINTREGEX} > LINTERERRORFILE
                        if ! [ -s "LINTERERRORFILE" ]; then
                            cat  LINTERERRORFILE   
                        fi
                    fi
                done                                            
            fi
        done
    done
    exit 1
else
    echo "No errors detected"
fi