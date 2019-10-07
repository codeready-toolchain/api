package v1alpha1

import (
	templatev1 "github.com/openshift/api/template/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NSTemplateTierSpec defines the desired state of NSTemplateTier
// +k8s:openapi-gen=true
type NSTemplateTierSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The namespace templates
	// +listType
	Namespaces []NSTemplateTierNamespace `json:"namespaces"`
}

// NSTemplateTierNamespace the namespace definition in an NSTemplateTier resource
type NSTemplateTierNamespace struct {

	// The type of the namespace. For example: ide|cicd|stage|default
	Type string `json:"type"`

	// The revision of the corresponding template
	Revision string `json:"revision"`

	// Template contains an OpenShift Template to be used for namespace provisioning
	Template templatev1.Template `json:"template"`
}

// NSTemplateTierStatus defines the observed state of NSTemplateTier
// +k8s:openapi-gen=true
type NSTemplateTierStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateTier is the Schema for the nstemplatetiers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type NSTemplateTier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NSTemplateTierSpec   `json:"spec,omitempty"`
	Status NSTemplateTierStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateTierList contains a list of NSTemplateTier
type NSTemplateTierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NSTemplateTier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NSTemplateTier{}, &NSTemplateTierList{})
}
