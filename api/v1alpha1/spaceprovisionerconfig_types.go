package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	SpaceProvisionerConfigToolchainClusterNotFoundReason  = "ToolchainClusterNotFound"
	SpaceProvisionerConfigToolchainClusterNotReadyReason  = "ToolchainClusterNotReady"
	SpaceProvisionerConfigInsufficientCapacityReason      = "InsufficientCapacity"
	SpaceProvisionerConfigFailedToDetermineCapacityReason = "FailedToDetermineCapacity"
	SpaceProvisionerConfigValidReason                     = "AllChecksPassed"
	SpaceProvisionerConfigDisabledReason                  = "Disabled"
)

// +k8s:openapi-gen=true
type SpaceProvisionerConfigSpec struct {
	// PlacementRoles is the list of roles, or flavors, that the provisioner possesses that influence
	// the space scheduling decisions.
	// +optional
	// +listType=set
	PlacementRoles []string `json:"placementRoles,omitempty"`

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
	// MaxNumberOfSpaces is the maximum number of spaces that can be provisioned to the referenced cluster.
	//
	// 0 or undefined value means no limit.
	//
	// +kubebuilder:validation:Minimum=0
	// +optional
	MaxNumberOfSpaces uint `json:"maxNumberOfSpaces,omitempty"`
	// MaxMemoryUtilizationPercent is the maximum memory utilization of the cluster to permit provisioning
	// new spaces to it.
	//
	// 0 or undefined value means no limit.
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +optional
	MaxMemoryUtilizationPercent uint `json:"maxMemoryUtilizationPercent,omitempty"`
}

// ConsumedCapacity describes the capacity of the cluster consumed by the spaces
// currently provisioned to it.
type ConsumedCapacity struct {
	// MemoryUsagePercentPerNodeRole is the percent of the memory used per node role (eg. worker, master)
	MemoryUsagePercentPerNodeRole map[string]int `json:"memoryUsagePercentPerNode"`

	// SpaceCount is the number of spaces currently deployed to the cluster
	SpaceCount int `json:"spaceCount"`
}

// +k8s:openapi-gen=true
type SpaceProvisionerConfigStatus struct {
	// ConsumedCapacity reflects the runtime state of the cluster and the capacity it currently consumes.
	// Nil if the consumed capacity is not known
	// +optional
	ConsumedCapacity *ConsumedCapacity `json:"consumedCapacity,omitempty"`

	// Conditions describes the state of the configuration (its validity).
	// The only known condition type is "Ready". The SpaceProvisionerConfig is ready when the following is true:
	//    * the referenced ToolchainCluster object exists and is itself ready
	//    * the consumed capacity doesn't breach the thresholds defined in the spec
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// SpaceProvisionerConfig is the configuration of space provisioning in the member clusters.
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
