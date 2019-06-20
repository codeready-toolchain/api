package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UserAccountSpec defines the desired state of UserAccount
// +k8s:openapi-gen=true
type UserAccountSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// UserID is the user ID from RHD Identity Provider token (“sub” claim)
	// Is to be used to create Identity and UserIdentityMapping resources
	UserID string `json:"userID"`

	// The namespace limit name
	NSLimit string `json:"nsLimit"`

	// Namespace template set
	NSTemplateSet NSTemplateSetSpec `json:"nsTemplateSet"`
}

// UserAccountStatus defines the observed state of UserAccount
// +k8s:openapi-gen=true
type UserAccountStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Observed status. For example: provisioning|provisioned
	Status string `json:"status,omitempty"`

	// The error message in case of failed status
	Error string `json:"error,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserAccount is the Schema for the useraccounts API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type UserAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserAccountSpec   `json:"spec,omitempty"`
	Status UserAccountStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserAccountList contains a list of UserAccount
type UserAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAccount `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserAccount{}, &UserAccountList{})
}
