package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of a MasterUserRecord
const (
	// UserProvisionRequestPendingApproval means the request is pending approval
	UserProvisionRequestPendingApproval ConditionType = "PendingApproval"
	// UserProvisionRequestProvisioning means the user is being provisioned
	UserProvisionRequestProvisioning ConditionType = "Provisioning"
	// UserProvisionRequestComplete means provisioning is complete
	UserProvisionRequestComplete ConditionType = "Complete"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UserProvisionRequestSpec defines the desired state of UserProvisionRequest
// +k8s:openapi-gen=true
type UserProvisionRequestSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// UserID is the user ID from RHD Identity Provider token (“sub” claim)
	UserID string `json:"userID"`

	// The cluster in which the user is provisioned in
	// If not set then the target cluster will be picked automatically
	// +optional
	TargetCluster string `json:"targetCluster,omitempty"`

	// If Approved set to 'true' then the user has been manually approved
	// If not set then the user is subject of auto-approval (if enabled)
	// +optional
	Approved bool `json:"approved,omitempty"`
}

// UserProvisionRequestStatus defines the observed state of UserProvisionRequest
// +k8s:openapi-gen=true
type UserProvisionRequestStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current UserProvisionRequest conditions
	// Supported condition types:
	// PendingApproval, Provisioning, Complete
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserProvisionRequest is the Schema for the userprovisionrequests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type UserProvisionRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserProvisionRequestSpec   `json:"spec,omitempty"`
	Status UserProvisionRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserProvisionRequestList contains a list of UserProvisionRequest
type UserProvisionRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProvisionRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserProvisionRequest{}, &UserProvisionRequestList{})
}
