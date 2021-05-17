package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemberOperatorConfigSpec contains all configuration parameters of the member operator
// +k8s:openapi-gen=true
type MemberOperatorConfigSpec struct {
	// Keeps parameters concerned with member status
	// +optional
	Status Status `json:"status,omitempty"`
}

// Defines all parameters concerned with member status
// +k8s:openapi-gen=true
type Status struct {
	// Defines the period between refreshes of the member status
	RefreshPeriod *string `json:"refreshPeriod"`
}

// MemberOperatorConfigStatus defines the observed state of MemberOperatorConfig
// +k8s:openapi-gen=true
type MemberOperatorConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemberOperatorConfig keeps all configuration parameters needed in member operator
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=memberoperatorconfigs,scope=Namespaced
// +kubebuilder:printcolumn:name="AutomaticApproval",type="boolean",JSONPath=`.spec.automaticApproval.enabled`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Member Operator Config"
type MemberOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemberOperatorConfigSpec   `json:"spec,omitempty"`
	Status MemberOperatorConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemberOperatorConfigList contains a list of MemberOperatorConfig
type MemberOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemberOperatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemberOperatorConfig{}, &MemberOperatorConfigList{})
}
