package v1alpha1

import (
	"reflect"
)

func (first *NSTemplateSetSpec) CompareTo(second *NSTemplateSetSpec) bool {
	if first.TierName != second.TierName {
		return false
	}
	return compareNamespaces(first.Namespaces, second.Namespaces)
}

func compareNamespaces(namespaces1, namespaces2 []Namespace) bool {
	if len(namespaces1) != len(namespaces2) {
		return false
	}
	for _, ns1 := range namespaces1 {
		found := findNamespace(ns1, namespaces2)
		if !found {
			return false
		}
	}
	return true
}

func findNamespace(thisNs Namespace, namespaces []Namespace) bool {
	for _, ns := range namespaces {
		if reflect.DeepEqual(thisNs, ns) {
			return true
		}
	}
	return false
}
