#!/bin/bash
TMP_DIR=/tmp/
BASE_REPO_PATH=$(mktemp -d ${TMP_DIR}replace-verify.XXX)
GH_BASE_URL_KS=https://github.com/kubesaw/
GH_BASE_URL_CRT=https://github.com/codeready-toolchain/
declare -a REPOS=("${GH_BASE_URL_KS}ksctl" "${GH_BASE_URL_CRT}host-operator" "${GH_BASE_URL_CRT}member-operator" "${GH_BASE_URL_CRT}registration-service" "${GH_BASE_URL_CRT}toolchain-e2e" "${GH_BASE_URL_CRT}toolchain-common")
C_PATH=${PWD}
ERROR_REPO_LIST=()
ERROR_FILE_LIST=()
STD_OUT_FILE_LIST=()
GO_LINT_REGEX="[\s\w.\/]*:[0-9]*:[0-9]*:[\w\s)(*.\`]*"
ERROR_REGEX="Error[:]*" #unit test or any other failure we log goes into stdoutput, hence making that regex too to fetch the error
FAIL_REGEX="FAIL[:]*" #unit test or any other failure we log goes into stdoutput, hence making that regex too to fetch the error

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
    std_out_file=$(mktemp ${BASE_REPO_PATH}/$(basename ${repo})-output.XXX)
    echo "std output file : ${std_out_file}"
    echo "Cloning repo in /tmp"
    git clone --depth=1 ${repo} ${repo_path}
    echo "Repo cloned successfully"
    cd ${repo_path}
    make pre-verify 2> >(tee ${err_file})
    rc=$?
    if [ ${rc} -ne 0 ]; then
        ERROR_REPO_LIST+="$(basename ${repo}) "
        ERROR_FILE_LIST+="${err_file}  "
        continue
    fi
    echo "Initiating 'go mod replace' of current api version in dependent repos"
    go mod edit -replace github.com/codeready-toolchain/api=${C_PATH}
        make verify-dependencies 2> >(tee ${err_file}) 1> >(tee ${std_out_file})
    rc=$?
    if [ ${rc} -ne 0 ]; then
    ERROR_REPO_LIST+="$(basename ${repo}) " 
    ERROR_FILE_LIST+="${err_file}  "
    STD_OUT_FILE_LIST+="${std_out_file} "
    fi
    echo                                                          
    echo =========================================================================================
    echo                                                           
done
echo                "Summary"
if [ ${#ERROR_REPO_LIST[@]} -ne 0 ]; then
    echo "Below are the repos with error: "
    for error_repo_name in ${ERROR_REPO_LIST[*]}
    do
        for error_file_name in ${ERROR_FILE_LIST[*]}
        do
            if [[ ${error_file_name} =~ ${error_repo_name} ]]; then 
                echo                                                          
                echo =========================================================================================
                echo 
                echo                       "${error_repo_name} has the following errors "
                echo                                                          
                echo =========================================================================================
                echo 
                cat "${error_file_name}"
            fi
        done
        for std_out_file_name in ${STD_OUT_FILE_LIST[*]}
            do
                if [[ ${std_out_file_name} =~ ${error_repo_name} ]]; then 
                    cat "${std_out_file_name}" | grep -E ${GO_LINT_REGEX}|${ERROR_REGEX}|${FAIL_REGEX}
                fi
        done                                             
    done
    exit 1
else
    echo "No errors detected"
fi