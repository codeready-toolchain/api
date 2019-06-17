package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MasterUserRecordSpec defines the desired state of MasterUserRecord
// +k8s:openapi-gen=true
type MasterUserRecordSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Desired state of the user record: approved|banned|deactivated
	State string `json:"state,omitempty"`

	// The list of user accounts in the member clusters which belong to this MasterUserRecord
	UserAccounts []UserAccountEmbedded `json:"userAccounts,omitempty"`
}

type UserAccountEmbedded struct {

	// The cluster in which the user exists
	TargetCluster string `json:"targetCluster"`

	// The resource version of the corresponding UserAccount
	ResourceVersion string `json:"resourceVersion"`

	// The spec of the corresponding UserAccount
	Spec UserAccountSpec `json:"spec"`
}

// MasterUserRecordStatus defines the observed state of MasterUserRecord
// +k8s:openapi-gen=true
type MasterUserRecordStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Observed status. For example: provisioning|provisioned
	Status string `json:"status,omitempty"`

	// The status of user accounts in the member clusters which belong to this MasterUserRecord
	UserAccounts []UserAccountStatusEmbedded `json:"userAccounts,omitempty"`
}

type UserAccountStatusEmbedded struct {

	// The cluster in which the user exists
	TargetCluster string `json:"targetCluster"`

	// The resource version of the corresponding UserAccount
	ResourceVersion string `json:"resourceVersion"`

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

func init() {
	SchemeBuilder.Register(&MasterUserRecord{}, &MasterUserRecordList{})
}
