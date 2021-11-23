package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid status condition reasons of a UserAccount
const (
	// Status condition reasons
	SpaceUnableToCreateNSTemplateSetReason = "UnableToCreateNSTemplateSet"
	SpaceProvisioningReason                = provisioningReason
	SpaceProvisionedReason                 = provisionedReason
	SpaceTerminatingReason                 = terminatingReason
	SpaceNSTemplateSetUpdateFailedReason   = "NSTemplateSetUpdateFailed"
)

// SpaceSpec defines the desired state of Space
type SpaceSpec struct {

	// TargetCluster The cluster in which this Space is provisioned
	// If not set then the target cluster will be picked automatically
	// +optional
	TargetCluster string `json:"targetCluster,omitempty"`

	// TierName is a required property introduced to retain the name of the tier
	// for which this Space is provisioned
	TierName string `json:"tierName"`
}

// SpaceStatus defines the observed state of Space
type SpaceStatus struct {

	// Conditions is an array of current User Account conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Space is the Schema for the spaces API
type Space struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceSpec   `json:"spec,omitempty"`
	Status SpaceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpaceList contains a list of Space
type SpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Space `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Space{}, &SpaceList{})
}
