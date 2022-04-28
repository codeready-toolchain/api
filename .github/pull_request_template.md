## Description
A few sentences describing the overall goals of the pull request's commits.

## Checks
1. Did you run `make generate` target? **yes/no**

2. Did `make generate` change anything in other projects (host-operator, member-operator)? **yes/no**

3. In case of **new** CRD, did you the following? **yes/no**
    - make/generate.mk in this repository
    - `resources/setup/roles/host.yaml` in the sandbox-sre repository
    - `PROJECT` file: https://github.com/codeready-toolchain/host-operator/blob/master/PROJECT
    - `CSV` file: https://github.com/codeready-toolchain/host-operator/blob/master/config/manifests/bases/host-operator.clusterserviceversion.yaml

4. In case other projects are changed, please provides PR links.
    - host-operator: https://github.com/codeready-toolchain/host-operator/pull/#
    - member-operator: https://github.com/codeready-toolchain/member-operator/pull/#
