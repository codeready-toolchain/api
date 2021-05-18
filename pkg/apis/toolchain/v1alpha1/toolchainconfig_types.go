package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ToolchainConfigSpec contains all configuration for host and member operators
// +k8s:openapi-gen=true
type ToolchainConfigSpec struct {
	// Contains all host operator configuration
	// +optional
	Host Host `json:"host,omitempty"`

	// Contains all member operator configurations for all member clusters
	// +optional
	Members Members `json:"members,omitempty"`
}

// Host contains all configuration parameters of the host operator
// +k8s:openapi-gen=true
type Host struct {
	// Keeps parameters necessary for automatic approval
	// +optional
	AutomaticApproval AutomaticApprovalCfg `json:"automaticApproval,omitempty"`

	// Keeps parameters concerned with user deactivation
	// +optional
	Deactivation DeactivationCfg `json:"deactivation,omitempty"`
}

// Members contains all configuration for member operators
// +k8s:openapi-gen=true
type Members struct {
	// Defines default configuration to be applied to all member clusters
	// +optional
	Default MemberOperatorConfigSpec `json:"default,omitempty"`

	// A map of cluster-specific member operator configurations indexed by member toolchaincluster name
	// +optional
	// +mapType=atomic
	SpecificPerMemberCluster map[string]MemberOperatorConfigSpec `json:"specificPerMemberCluster,omitempty"`
}

// Defines all parameters necessary for automatic approval
// +k8s:openapi-gen=true
type AutomaticApprovalCfg struct {
	// Defines if the automatic approval is enabled or not
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	// Contains threshold (in percentage of usage) that defines when the automatic approval should be stopped
	// +optional
	ResourceCapacityThreshold ResourceCapacityThresholdCfg `json:"resourceCapacityThreshold,omitempty"`

	// Defines the maximal number of users to be allowed for automatic approval.
	// When the number is reached, then the automatic approval is stopped.
	// +optional
	MaxNumberOfUsers MaxNumberOfUsersCfg `json:"maxNumberOfUsers,omitempty"`
}

// Contains default capacity threshold as well as specific ones for particular member clusters
// +k8s:openapi-gen=true
type ResourceCapacityThresholdCfg struct {
	// It is the default capacity threshold (in percentage of usage) to be used for all member clusters if no special threshold is defined
	// +optional
	DefaultThreshold *int `json:"defaultThreshold,omitempty"`

	// Contains a map of specific capacity thresholds (in percentage of usage) for particular member clusters mapped by their names
	// +optional
	// +mapType=atomic
	SpecificPerMemberCluster map[string]int `json:"specificPerMemberCluster,omitempty"`
}

// Contains maximal number of users to be provisioned automatically in the system overall as well as
// max number of users automatically provisioned per member cluster
// +k8s:openapi-gen=true
type MaxNumberOfUsersCfg struct {
	// It is the maximal number of users provisioned in the system overall - equals to max number of MasterUserRecords in host cluster
	// +optional
	Overall *int `json:"overall,omitempty"`

	// Contains a map of maximal number of users provisioned per member cluster mapped by the cluster name
	// - equals to max number of UserAccounts in member cluster
	// +optional
	// +mapType=atomic
	SpecificPerMemberCluster map[string]int `json:"specificPerMemberCluster,omitempty"`
}

type DeactivationCfg struct {

	// DeactivatingNotificationDays is the number of days after a pre-deactivating notification is sent that actual
	// deactivation occurs.  If this parameter is set to zero, then there will be no delay
	// +optional
	DeactivatingNotificationDays *int `json:"deactivatingNotificationDays,omitempty"`
}

// ToolchainConfigStatus defines the observed state of ToolchainConfig
// +k8s:openapi-gen=true
type ToolchainConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainConfig keeps all configuration parameters needed for host and member operators
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=toolchainconfigs,scope=Namespaced
// +kubebuilder:printcolumn:name="AutomaticApproval",type="boolean",JSONPath=`.spec.host.automaticApproval.enabled`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Toolchain Operator Config"
type ToolchainConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToolchainConfigSpec   `json:"spec,omitempty"`
	Status ToolchainConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainConfigList contains a list of ToolchainConfig
type ToolchainConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolchainConfig{}, &ToolchainConfigList{})
}
