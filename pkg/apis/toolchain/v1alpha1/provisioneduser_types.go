package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProvisionedUserSpec defines the desired state of ProvisionedUser
// +k8s:openapi-gen=true
type ProvisionedUserSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The cluster in which the user exists
	Cluster string `json:"cluster"`
}

// ProvisionedUserStatus defines the observed state of ProvisionedUser
// +k8s:openapi-gen=true
type ProvisionedUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProvisionedUser is the Schema for the provisionedusers API
// ProvisionedUser resource represents an already-provisioned user, and is used to track all users across all clusters.
// It should only be created on the host cluster; the complete set of all ProvisionedUser resources contained within
// the host cluster in their entirety represent the user registry
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ProvisionedUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProvisionedUserSpec   `json:"spec,omitempty"`
	Status ProvisionedUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProvisionedUserList contains a list of ProvisionedUser
type ProvisionedUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProvisionedUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProvisionedUser{}, &ProvisionedUserList{})
}
