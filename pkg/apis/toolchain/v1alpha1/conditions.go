package v1alpha1

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ConditionType string

const (
	// ConditionReady specifies that the resource is ready
	ConditionReady ConditionType = "Ready"

	// Status reasons
	provisioningReason = "Provisioning"
	provisionedReason  = "Provisioned"
	disabledReason     = "Disabled"
	terminatingReason  = "Terminating"
	updatingReason     = "Updating"

	// Condition types
	deletionError = "DeletionError"
)

// These are valid status condition reasons for Toolchain status
const (
	// overall status condition reasons
	ToolchainStatusReasonAllComponentsReady = "AllComponentsReady"
	ToolchainStatusReasonComponentsNotReady = "ComponentsNotReady"

	// deployment reasons
	ToolchainStatusReasonDeploymentReady    = "DeploymentReady"
	ToolchainStatusReasonDeploymentNotReady = "DeploymentNotReady"
	ToolchainStatusReasonDeploymentNotFound = "DeploymentNotFound"

	// kubefed reasons
	ToolchainStatusReasonClusterConnectionNotFound              = "KubefedNotFound"
	ToolchainStatusReasonClusterConnectionLastProbeTimeExceeded = "KubefedLastProbeTimeExceeded"

	// registration service reasons
	ToolchainStatusReasonRegServiceReady              = "RegServiceReady"
	ToolchainStatusReasonRegServiceNotReady           = "RegServiceNotReady"
	ToolchainStatusReasonRegServiceResourceNotFound   = "RegServiceResourceNotFound"
	ToolchainStatusReasonRegServiceDeploymentNotFound = "RegServiceDeploymentNotFound"

	// member status reasons
	ToolchainStatusReasonMemberStatusNoClustersFound = "NoMemberClustersFound"
	ToolchainStatusReasonMemberStatusNotFound        = "MemberStatusNotFound"
)

type Condition struct {
	// Type of condition
	Type ConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status apiv1.ConditionStatus `json:"status"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
	// Last time the condition was updated
	// +optional
	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`
}
