# ToolChain API

[![Go Report Card](https://goreportcard.com/badge/github.com/codeready-toolchain/api)](https://goreportcard.com/report/github.com/codeready-toolchain/api)
[![GoDoc](https://godoc.org/github.com/codeready-toolchain/api?status.png)](https://godoc.org/github.com/codeready-toolchain/api)

For the API reference docs go [here](api/v1alpha1/docs/apiref.adoc)

## Prerequisites

* Go version 1.24.x (1.24.4 or higher) - download for your development environment [here](https://golang.org/dl/).

CodeReady ToolChain API is built using [Go modules](https://github.com/golang/go/wiki/Modules).

## Modifying the API Types

The API types are defined in the `api/v1alpha1/*_types.go` files. After modifying these files, you must regenerate the derived files (deepcopy, OpenAPI, CRD manifests, and API reference docs) by running:

```sh
make generate
```

This command runs the following steps:

1. **Generate deepcopy and CRDs** — uses [controller-gen](https://github.com/kubernetes-sigs/controller-tools) to regenerate `zz_generated.deepcopy.go` and the CRD manifests in `config/crd/bases/`.
2. **Generate OpenAPI** — uses [openapi-gen](https://github.com/kubernetes/kube-openapi) to regenerate `zz_generated.openapi.go`.
3. **Generate API reference docs** — uses [crd-ref-docs](https://github.com/elastic/crd-ref-docs) to regenerate `api/v1alpha1/docs/apiref.adoc`.
4. **Dispatch CRDs** — copies the generated CRD `.yaml` files to the `host-operator` and `member-operator` repositories.

> **Note:** The CRD dispatch step assumes the `host-operator` and `member-operator` repositories have been checked out alongside this repository and that *they are in a clean state*, meaning that they have no pending changes besides previous versions of the CRD files.

> **Note:** After running `make generate`, you are expected to create PRs in this repository as well as in the `host-operator` and `member-operator` repositories where the CRD changes were propagated. Please do not mix other code changes with CRD changes — it is always preferred to promote CRD changes separately for easier PR review.
