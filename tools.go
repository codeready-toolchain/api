//go:build tools
// +build tools

package tools

import (
	// Code generators built at runtime.
	_ "github.com/elastic/crd-ref-docs"
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
