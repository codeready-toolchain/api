package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of a MasterUserRecord
const (
	// MasterUserRecordProvisioning means the Master User Record is being provisioned
	MasterUserRecordProvisioning ConditionType = "Provisioning"
	// MasterUserRecordUserAccountNotReady means the User Account failed to be provisioned
	MasterUserRecordUserAccountNotReady ConditionType = "UserAccountNotReady"
	// MasterUserRecordReady means the Master User Record provisioning succeeded
	MasterUserRecordReady ConditionType = "Ready"

	MasterUserRecordUserIDLabelKey = "toolchain.dev.openshift.com/user-id"
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

	// If set to true then the corresponding user has been banned from logging in and accessing their resources
	// +optional
	Banned bool `json:"banned,omitempty"`

	// If set to true then the corresponding UserAccount should be deleted
	// "false" is assumed by default
	// +optional
	Deprovisioned bool `json:"deprovisioned,omitempty"`

	// The list of user accounts in the member clusters which belong to this MasterUserRecord
	// +listType
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
	// Supported condition types:
	// Provisioning, UserAccountNotReady and Ready
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// The status of user accounts in the member clusters which belong to this MasterUserRecord
	// +listType
	UserAccounts []UserAccountStatusEmbedded `json:"userAccounts,omitempty"`
}

type UserAccountStatusEmbedded struct {

	// Cluster is the cluster in which the user exists
	Cluster Cluster `json:"cluster"`

	// SyncIndex is used for checking if there is needed some MasterUserRecord <-> UserAccount
	// synchronization for this specific UserAccount in the specific member cluster
	SyncIndex string `json:"syncIndex"`

	// Inherits the status from the corresponding UserAccount status
	UserAccountStatus `json:",inline"`
}

type Cluster struct {
	// Name is the name of the corresponding KubeFedCluster resource
	Name string `json:"name"`

	// APIEndpoint is the API Endpoint of the cluster
	APIEndpoint string `json:"apiEndpoint"`

	// ConsoleURL is the web console URL of the cluster
	ConsoleURL string `json:"consoleURL"`

	// CheDashboardURL is the Che Dashboard URL of the cluster if Che is installed
	// +optional
	CheDashboardURL string `json:"cheDashboardURL,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterUserRecord is the Schema for the masteruserrecords API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:shortName=mur
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=`.spec.userAccounts[].targetCluster`
// +kubebuilder:printcolumn:name="Banned",type="string",JSONPath=`.spec.banned`,priority=1
// +kubebuilder:printcolumn:name="Disabled",type="string",JSONPath=`.spec.disabled`,priority=1
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

func init() {
	SchemeBuilder.Register(&MasterUserRecord{}, &MasterUserRecordList{})
}
