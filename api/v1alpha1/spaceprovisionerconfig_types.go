package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	SpaceProvisionerConfigValidConditionType             = "Valid"
	SpaceProvisionerConfigToolchainClusterNotFoundReason = "ToolchainClusterNotFound"
	SpaceProvisionerConfigValidReason                    = "AllChecksPassed"
)

// +k8s:openapi-gen=true
type SpaceProvisionerConfigSpec struct {
	// PlacementRoles is the list of roles, or flavors, that the provisioner possesses that influence
	// the space scheduling decisions.
	// +optional
	PlacementRoles *[]string `json:"placementRoles,omitempty"`

	// ToolchainCluster is the name of the ToolchainCluster CR of the member cluster that this config is for.
	ToolchainCluster string `json:"toolchainCluster"`

	// Enabled specifies whether the member cluster is enabled (and therefore can hold spaces) or not.
	// +optional
	// +kubebuilder:default=false
	Enabled bool `json:"enabled"`

	// CapacityThresholds specifies the max capacities allowed in this provisioner
	// +optional
	CapacityThresholds SpaceProvisionerCapacityThresholds `json:"capacityThresholds"`
}

// SpaceProvisionerCapacityThresholds defines the capacity thresholds of the space provisioner
// +k8s:openapi-gen=true
type SpaceProvisionerCapacityThresholds struct {
	// +kubebuilder:validation:Minimum=0
	MaxNumberOfSpaces uint `json:"maxNumberOfSpaces"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	MaxMemoryUtilizationPercent uint `json:"maxMemoryUtilizationPercent"`
}

// +k8s:openapi-gen=true
type SpaceProvisionerConfigStatus struct {
	// Conditions describes the state of the configuration (its validity).
	// The only known condition type is "Valid".
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// SpaceProvisionerConfig is the configuration of space provisioning in the member clusters.
//
// Note that these objects are currently NOT used anywhere.
//
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=`.spec.toolchainCluster`
// +kubebuilder:printcolumn:name="Enabled",type="boolean",JSONPath=`.spec.enabled`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="SpaceProvisionerConfig"
type SpaceProvisionerConfig struct {
	Spec              SpaceProvisionerConfigSpec   `json:"spec,omitempty"`
	Status            SpaceProvisionerConfigStatus `json:"status,omitempty"`
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

//+kubebuilder:object:root=true

// SpaceProvisionerConfigList contains a list of SpaceProvisionerConfig
// +k8s:openapi-gen=true
type SpaceProvisionerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceProvisionerConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceProvisionerConfig{}, &SpaceProvisionerConfigList{})
}
