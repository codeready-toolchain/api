module github.com/codeready-toolchain/api

go 1.13

require (
	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-openapi/spec v0.19.3
	// using 'github.com/openshift/api@release-4.4'
	github.com/openshift/api v0.0.0-20200414152312-3e8f22fb0b56
	github.com/stretchr/testify v1.4.0
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/code-generator v0.17.4
	k8s.io/gengo v0.0.0-20190822140433-26a664648505
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
	sigs.k8s.io/controller-runtime v0.5.2
	sigs.k8s.io/controller-tools v0.2.8
	sigs.k8s.io/kubefed v0.3.0
)
