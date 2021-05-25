package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemberOperatorConfigSpec contains all configuration parameters of the member operator
// +k8s:openapi-gen=true
type MemberOperatorConfigSpec struct {
	// Keeps parameters concerned with member status
	// +optional
	MemberStatus MemberStatusConfig `json:"memberStatus,omitempty"`
}

// Defines all parameters concerned with member status
// +k8s:openapi-gen=true
type MemberStatusConfig struct {
	// Defines the period between refreshes of the member status
	RefreshPeriod *string `json:"refreshPeriod,omitempty"`
}

// MemberOperatorConfigStatus defines the observed state of MemberOperatorConfig
// +k8s:openapi-gen=true
type MemberOperatorConfigStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MemberOperatorConfig keeps all configuration parameters needed in member operator
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=memberoperatorconfigs,scope=Namespaced
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Member Operator Config"
type MemberOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemberOperatorConfigSpec   `json:"spec,omitempty"`
	Status MemberOperatorConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MemberOperatorConfigList contains a list of MemberOperatorConfig
type MemberOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemberOperatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemberOperatorConfig{}, &MemberOperatorConfigList{})
}
