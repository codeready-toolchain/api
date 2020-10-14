module github.com/codeready-toolchain/api

go 1.14

require (
	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-openapi/spec v0.19.3
	github.com/gobuffalo/flect v0.2.1 // indirect
	github.com/golangci/golangci-lint v1.31.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	// using latest commit from 'github.com/openshift/api@release-4.5'
	github.com/openshift/api v0.0.0-20200821140346-b94c46af3f2b
	github.com/spf13/cobra v1.0.0 // indirect
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/code-generator v0.18.3
	k8s.io/gengo v0.0.0-20200114144118-36b2048a9120
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/controller-tools v0.3.0
)
