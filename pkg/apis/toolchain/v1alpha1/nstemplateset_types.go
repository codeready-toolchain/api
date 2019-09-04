package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Namespaces []Namespace `json:"namespaces"`
}

type Namespace struct {

	// The type of the namespace. For example: ide|cicd|stage|default
	Type string `json:"type"`

	// The revision of the corresponding template
	Revision string `json:"revision"`

	// Optional field. Used to specify a custom template
	Template string `json:"template,omitempty"`
}

// NSTemplateSetStatus defines the observed state of NSTemplateSet
// +k8s:openapi-gen=true
type NSTemplateSetStatus struct {
	// Conditions is an array of current NSTemplateSet conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateSet is the Schema for the nstemplatesets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Tier Name",type="string",JSONPath=".spec.tierName"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=="Ready")].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=="Ready")].reason"
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
