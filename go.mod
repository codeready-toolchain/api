module github.com/codeready-toolchain/api

go 1.16

require (
	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	// using latest commit from 'github.com/openshift/api@release-4.9'
	github.com/openshift/api v0.0.0-20211028023115-7224b732cc14
	k8s.io/api v0.22.7
	k8s.io/apimachinery v0.22.7
	k8s.io/code-generator v0.22.7
	k8s.io/gengo v0.0.0-20201214224949-b6c5ce23f027
	k8s.io/kube-openapi v0.0.0-20211109043538-20434351676c
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/controller-runtime v0.10.3
	sigs.k8s.io/controller-tools v0.7.0
)
