package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	OwnerLabelKey       = LabelKeyPrefix + "owner"
	TypeLabelKey        = LabelKeyPrefix + "type"
	TemplateRefLabelKey = LabelKeyPrefix + "templateref"
	TierLabelKey        = LabelKeyPrefix + "tier"
	ProviderLabelKey    = LabelKeyPrefix + "provider"
	ProviderLabelValue  = "codeready-toolchain"
)

// These are valid status condition reasons of a NSTemplateSet
const (
	NSTemplateSetProvisionedReason                       = provisionedReason
	NSTemplateSetProvisioningReason                      = provisioningReason
	NSTemplateSetUnableToProvisionReason                 = "UnableToProvision"
	NSTemplateSetUnableToProvisionNamespaceReason        = "UnableToProvisionNamespace"
	NSTemplateSetUnableToProvisionClusterResourcesReason = "UnableToProvisionClusteResources"
	NSTemplateSetTerminatingReason                       = terminatingReason
	NSTemplateSetTerminatingFailedReason                 = "UnableToTerminate"
	NSTemplateSetUpdatingReason                          = updatingReason
	NSTemplateSetUpdateFailedReason                      = "UpdateFailed"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NSTemplateSetSpec defines the desired state of NSTemplateSet
// +k8s:openapi-gen=true
type NSTemplateSetSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The name of the tier represented by this template set
	TierName string `json:"tierName"`

	// The namespace templates
	// +listType=set
	Namespaces []NSTemplateSetNamespace `json:"namespaces"`

	// the cluster resources template (for cluster-wide quotas, etc.)
	// +optional
	ClusterResources *NSTemplateSetClusterResources `json:"clusterResources,omitempty"`
}

// NSTemplateSetNamespace the namespace definition in an NSTemplateSet resource
type NSTemplateSetNamespace struct {
	//The type of the namespace. For example: ide|cicd|stage|default
	// +optional
	Type string `json:"type,omitempty"`

	// The revision of the corresponding template
	// +optional
	Revision string `json:"revision,omitempty"`

	// Optional field. Used to specify a custom template
	// +optional
	Template string `json:"template,omitempty"`

	// TemplateRef The name of the TierTemplate resource which exists in the host cluster and which contains the template to use
	TemplateRef string `json:"templateRef"`
}

// NSTemplateSetClusterResources defines the cluster-scoped resources associated with a given user
type NSTemplateSetClusterResources struct {
	// The revision of the corresponding template
	// +optional
	Revision string `json:"revision,omitempty"`

	// Template contains an OpenShift Template to be used for provisioning of cluster-scoped resources
	// +optional
	Template string `json:"template,omitempty"`

	// TemplateRef The name of the TierTemplate resource which exists in the host cluster and which contains the template to use
	TemplateRef string `json:"templateRef"`
}

// NSTemplateSetStatus defines the observed state of NSTemplateSet
// +k8s:openapi-gen=true
type NSTemplateSetStatus struct {
	// Conditions is an array of current NSTemplateSet conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=set
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateSet defines user environment via templates that are used for namespace provisioning
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Tier Name",type="string",JSONPath=`.spec.tierName`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Namespace Template Set"
type NSTemplateSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NSTemplateSetSpec   `json:"spec,omitempty"`
	Status NSTemplateSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateSetList contains a list of NSTemplateSet
type NSTemplateSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NSTemplateSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NSTemplateSet{}, &NSTemplateSetList{})
}

// CompareTo compares given second spec with this spec
func (first *NSTemplateSetSpec) CompareTo(second NSTemplateSetSpec) bool {
	if first.TierName != second.TierName {
		return false
	}
	return compareNamespaces(first.Namespaces, second.Namespaces)
}

func compareNamespaces(namespaces1, namespaces2 []NSTemplateSetNamespace) bool {
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

func findNamespace(thisNs NSTemplateSetNamespace, namespaces []NSTemplateSetNamespace) bool {
	for _, ns := range namespaces {
		if reflect.DeepEqual(thisNs, ns) {
			return true
		}
	}
	return false
}
