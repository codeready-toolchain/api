package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MasterUserRecordConditionType string

// These are valid conditions of a MasterUserRecord
const (
	// MasterUserRecordProvisioning means the Master User Record is being provisioned
	MasterUserRecordProvisioning MasterUserRecordConditionType = "Provisioning"
	// MasterUserRecordUserAccountNotReady means the User Account failed to be provisioned
	MasterUserRecordUserAccountNotReady MasterUserRecordConditionType = "UserAccountNotReady"
	// MasterUserRecordReady means the Master User Record failed to be provisioned
	MasterUserRecordReady MasterUserRecordConditionType = "Ready"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MasterUserRecordSpec defines the desired state of MasterUserRecord
// +k8s:openapi-gen=true
type MasterUserRecordSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// UserID is the user ID from RHD Identity Provider token (“sub” claim)
	UserID string `json:"userID"`

	// If set to true then the corresponding user should not be able to login (but the underlying UserAccounts still exists)
	// "false" is assumed by default
	// +optional
	Disabled bool `json:"disabled,omitempty"`

	// If set to true then the corresponding UserAccount should be deleted
	// "false" is assumed by default
	// +optional
	Deprovisioned bool `json:"deprovisioned,omitempty"`

	// The list of user accounts in the member clusters which belong to this MasterUserRecord
	UserAccounts []UserAccountEmbedded `json:"userAccounts,omitempty"`
}

type UserAccountEmbedded struct {

	// The cluster in which the user exists
	TargetCluster string `json:"targetCluster"`

	// SyncIndex is to be updated by UserAccount Controller
	// when the member needs to trigger MasterUserRecord <-> UserAccount synchronization
	SyncIndex string `json:"syncIndex"`

	// The spec of the corresponding UserAccount
	Spec UserAccountSpec `json:"spec"`
}

// MasterUserRecordStatus defines the observed state of MasterUserRecord
// +k8s:openapi-gen=true
type MasterUserRecordStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Master User Record conditions
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []MasterUserRecordCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// The status of user accounts in the member clusters which belong to this MasterUserRecord
	UserAccounts []UserAccountStatusEmbedded `json:"userAccounts,omitempty"`
}

type UserAccountStatusEmbedded struct {

	// The cluster in which the user exists
	TargetCluster string `json:"targetCluster"`

	// SyncIndex is used for checking if there is needed some MasterUserRecord <-> UserAccount
	// synchronization for this specific UserAccount in the specific member cluster
	SyncIndex string `json:"syncIndex"`

	// Inherits the status from the corresponding UserAccount status
	UserAccountStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterUserRecord is the Schema for the masteruserrecords API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type MasterUserRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MasterUserRecordSpec   `json:"spec,omitempty"`
	Status MasterUserRecordStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterUserRecordList contains a list of MasterUserRecord
type MasterUserRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MasterUserRecord `json:"items"`
}

// MasterUserRecordCondition describes current state of a MasterUserRecord
type MasterUserRecordCondition struct {
	Condition `json:",inline"`
	// Type of MasterUserRecord condition, Provisioning, UserAccountNotReady or Ready
	Type MasterUserRecordConditionType `json:"type"`
}

func init() {
	SchemeBuilder.Register(&MasterUserRecord{}, &MasterUserRecordList{})
}
