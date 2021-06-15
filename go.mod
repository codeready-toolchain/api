module github.com/codeready-toolchain/api

go 1.16

require (
	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	// using latest commit from 'github.com/openshift/api@release-4.7'
	github.com/openshift/api v0.0.0-20210428205234-a8389931bee7
	golang.org/x/tools v0.0.0-20200812195022-5ae4c3c160a0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/code-generator v0.20.2
	k8s.io/gengo v0.0.0-20201113003025-83324d819ded
	k8s.io/klog/v2 v2.8.0 // indirect
	k8s.io/kube-openapi v0.0.0-20210527164424-3c818078ee3d
	sigs.k8s.io/controller-runtime v0.8.3
	sigs.k8s.io/controller-tools v0.5.0
)
