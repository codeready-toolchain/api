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

echo Initiating verify-replace on dependent repos
for repo in "${REPOS[@]}"
do
    echo =========================================================================================
    echo  
    echo                        "$(basename ${repo})"
    echo                                                                     
    echo =========================================================================================                                            
    repo_path=${BASE_REPO_PATH}/$(basename ${repo})
    err_file=$(mktemp ${BASE_REPO_PATH}/$(basename ${repo})-error.XXX)
    echo "error output file : ${err_file}"
    stdout_file=$(mktemp ${BASE_REPO_PATH}/$(basename ${repo})-output.XXX)
    echo "std output file : ${stdout_file}"
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
    make verify-dependencies 2> >(tee ${err_file}) 1> >(tee ${stdout_file})
    rc=$?
    if [ ${rc} -ne 0 ]; then
    ERRORREPOLIST+="$(basename ${repo}) " 
    ERRORFILELIST+="${err_file}  "
    STDOUTFILELIST+="${stdout_file} "
    fi
    echo                                                          
    echo =========================================================================================
    echo                                                           
done
echo                "Summary"
if [ ${#ERRORREPOLIST[@]} -ne 0 ]; then
    echo "Below are the repos with error: "
    for errorreponame in ${ERRORREPOLIST[*]}
    do
        for errorfilename in ${ERRORFILELIST[*]}
        do
            if [[ ${errorfilename} =~ ${errorreponame} ]]; then 
                echo                                                          
                echo =========================================================================================
                echo 
                echo                       "${errorreponame} has the following errors "
                echo                                                          
                echo =========================================================================================
                echo 
                cat "${errorfilename}"
            else
            # Since golint check is the only check for which we parse the error from the standard-ouput
            # and is checked at last after all the other checks are done. But if there is any error in the previous checks,
            # the script won't go ahead to check golint, and hence this check is seperated and put into else
            # Meaning if the other checks have passed then only it wil proceed to check golint and we would parse 
            # golint errors from std files if any.
                for stdoutfilename in ${STDOUTFILELIST[*]}
                do
                    if [[ ${stdoutfilename} =~ ${errorreponame} ]]; then 
                        cat "${stdoutfilename}" | grep -E ${GOLINTREGEX}
                    fi
                done                                            
            fi
        done
    done
    exit 1
else
    echo "No errors detected"
fi