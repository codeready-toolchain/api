package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HostOperatorConfigStatus defines the observed state of HostOperatorConfig
// +k8s:openapi-gen=true
type HostOperatorConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostOperatorConfig keeps all configuration parameters needed in host operator
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=hostoperatorconfigs,scope=Namespaced
// +kubebuilder:printcolumn:name="AutomaticApproval",type="boolean",JSONPath=`.spec.automaticApproval.enabled`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Host Operator Config"
type HostOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostConfig               `json:"spec,omitempty"`
	Status HostOperatorConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostOperatorConfigList contains a list of HostOperatorConfig
type HostOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostOperatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HostOperatorConfig{}, &HostOperatorConfigList{})
}
