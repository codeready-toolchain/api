package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HostOperatorConfigSpec contains all configuration parameters of the host operator
// +k8s:openapi-gen=true
type HostOperatorConfigSpec struct {
	// Keeps parameters necessary for automatic approval
	// +optional
	AutomaticApproval AutomaticApproval `json:"automaticApproval,omitempty"`
}

// Defines all parameters necessary for automatic approval
// +k8s:openapi-gen=true
type AutomaticApproval struct {
	// Defines if the automatic approval is activated or not
	Activated bool `json:"activated"`

	// Contains threshold (in percentage of usage) that defines when the automatic approval should be stopped
	// +optional
	ResourceCapacityThreshold ResourceCapacityThreshold `json:"resourceCapacityThreshold,omitempty"`

	// Defines the maximal number of users to be allowed for automatic approval.
	// When the number is reached, then the automatic approval is stopped.
	// +optional
	MaxNumberOfUsers MaxNumberOfUsers `json:"maxNumberOfUsers,omitempty"`
}

// Contains default capacity threshold as well as specific ones for particular member clusters
// +k8s:openapi-gen=true
type ResourceCapacityThreshold struct {
	// It is the default capacity threshold (in percentage of usage) to be used for all member clusters if no special threshold is defined
	DefaultThreshold int `json:"defaultThreshold"`

	// Contains a map of specific capacity thresholds (in percentage of usage) for particular member clusters mapped by their names
	// +optional
	// +mapType=atomic
	SpecificPerMemberCluster map[string]int `json:"specificPerMemberCluster,omitempty"`
}

// Contains maximal number of users to be provisioned automatically in the system overall as well as
// max number of users automatically provisioned per member cluster
// +k8s:openapi-gen=true
type MaxNumberOfUsers struct {
	// It is the maximal number of users provisioned in the system overall - equals to max number of MasterUserRecords in host cluster
	Overall int `json:"overall"`

	// Contains a map of maximal number of users provisioned per member cluster mapped by the cluster name
	// - equals to max number of UserAccounts in member cluster
	// +optional
	// +mapType=atomic
	SpecificPerMemberCluster map[string]int `json:"specificPerMemberCluster,omitempty"`
}

// HostOperatorConfigStatus defines the observed state of HostOperatorConfig
// +k8s:openapi-gen=true
type HostOperatorConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostOperatorConfig keeps all configuration parameters needed in host operator
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=hostoperatorconfigs,scope=Namespaced
// +kubebuilder:printcolumn:name="AutomaticApproval",type="boolean",JSONPath=`.spec.automaticApproval.activated`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Host Operator Config"
type HostOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostOperatorConfigSpec   `json:"spec,omitempty"`
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
