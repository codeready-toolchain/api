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
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// String representation of the overall observed status. For example: provisioning|provisioned|updating
	Status string `json:"status,omitempty"`

	// The detailed namespace statuses
	Namespaces []NamespaceStatus `json:"namespaces,omitempty"`
}

type NamespaceStatus struct {

	// The name of the namespace
	Name string `json:"name"`

	// The type of the namespace. For example: ide|cicd|stage|default
	Type string `json:"type"`

	// Observed status. For example: provisioning|provisioned|failed
	Status string `json:"status,omitempty"`

	// The error message in case of failed status
	Error string `json:"error,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateSet is the Schema for the nstemplatesets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
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
