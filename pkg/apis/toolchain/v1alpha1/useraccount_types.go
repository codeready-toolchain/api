package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid status condition reasons of a UserAccount
const (
	// Status condition reasons
	UserAccountUnableToCreateUserReason          = "UnableToCreateUser"
	UserAccountUnableToCreateIdentityReason      = "UnableToCreateIdentity"
	UserAccountUnableToCreateMappingReason       = "UnableToCreateMapping"
	UserAccountUnableToCreateNSTemplateSetReason = "UnableToCreateNSTemplateSet"
	UserAccountProvisioningReason                = provisioningReason
	UserAccountProvisionedReason                 = provisionedReason
	UserAccountDisabledReason                    = disabledReason
	UserAccountDisablingReason                   = "Disabling"
	UserAccountTerminatingReason                 = terminatingReason
	UserAccountUpdatingReason                    = updatingReason
	UserAccountNSTemplateSetUpdateFailedReason   = "NSTemplateSetUpdateFailed"
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

	// If set to true then the corresponding user should not be able to login
	// "false" is assumed by default
	// +optional
	Disabled bool `json:"disabled,omitempty"`

	// UserAccountBase contains all base spec fields
	UserAccountSpecBase `json:",inline"`
}

// UserAccountSpecBase defines the common fields between UserAccountSpec
// and UserAccountSpecEmbedded
// +k8s:openapi-gen=true
type UserAccountSpecBase struct {
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

	// Conditions is an array of current User Account conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=set
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserAccount keeps all information about user provisioned in the cluster
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="User ID",type="string",JSONPath=`.spec.userID`,priority=1
// +kubebuilder:printcolumn:name="NS Limit",type="string",JSONPath=`.spec.nsLimit`
// +kubebuilder:printcolumn:name="Tier Name",type="string",JSONPath=`.spec.nsTemplateSet.tierName`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:printcolumn:name="Disabled",type="boolean",JSONPath=`.spec.disabled`,priority=1
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="User Account"
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
