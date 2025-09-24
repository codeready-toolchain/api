package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemberOperatorConfigSpec contains all configuration parameters of the member operator
// +k8s:openapi-gen=true
type MemberOperatorConfigSpec struct {
	// Keeps parameters concerned with authentication
	// +optional
	Auth AuthConfig `json:"auth,omitempty"`

	// Keeps parameters concerned with the autoscaler
	// +optional
	Autoscaler AutoscalerConfig `json:"autoscaler,omitempty"`

	// Keeps parameters concerned with the console
	// +optional
	Console ConsoleConfig `json:"console,omitempty"`

	// Environment specifies the member-operator environment such as prod, stage, unit-tests, e2e-tests, dev, etc
	// +optional
	Environment *string `json:"environment,omitempty"`

	// Defines the flag that determines whether User and Identity resources should be created for a UserAccount
	// +optional
	SkipUserCreation *bool `json:"skipUserCreation,omitempty"`

	// Keeps parameters concerned with member status
	// +optional
	MemberStatus MemberStatusConfig `json:"memberStatus,omitempty"`

	// Keeps parameters concerned with the toolchaincluster
	// +optional
	ToolchainCluster ToolchainClusterConfig `json:"toolchainCluster,omitempty"`

	// Keeps parameters concerned with the webhook
	// +optional
	Webhook WebhookConfig `json:"webhook,omitempty"`
}

// Defines all parameters concerned with the autoscaler
// +k8s:openapi-gen=true
type AuthConfig struct {
	// Represents the configured identity provider
	// +optional
	Idp *string `json:"idp,omitempty"`
}

// Defines all parameters concerned with the autoscaler
// +k8s:openapi-gen=true
type AutoscalerConfig struct {
	// Defines the flag that determines whether to deploy the autoscaler buffer
	// +optional
	Deploy *bool `json:"deploy,omitempty"`

	// Represents how much memory should be required by the autoscaler buffer
	// +optional
	BufferMemory *string `json:"bufferMemory,omitempty"`

	// Represents how much CPU should be required by the autoscaler buffer
	// +optional
	BufferCPU *string `json:"bufferCPU,omitempty"`

	// Represents the number of autoscaler buffer replicas to request
	// +optional
	BufferReplicas *int `json:"bufferReplicas,omitempty"`
}

// Defines all parameters concerned with the console
// +k8s:openapi-gen=true
type ConsoleConfig struct {
	// Defines the console route namespace
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	// Defines the console route name
	// +optional
	RouteName *string `json:"routeName,omitempty"`
}

// GitHubSecret defines all secrets related to GitHub authentication/integration
// +k8s:openapi-gen=true
type GitHubSecret struct {
	// The reference to the secret that is expected to contain the keys below
	// +optional
	ToolchainSecret `json:",inline"`

	// The key for the GitHub Access token in the secret values map
	// +optional
	AccessTokenKey *string `json:"accessTokenKey,omitempty"`
}

// Defines all parameters concerned with the toolchaincluster resource
// +k8s:openapi-gen=true
type ToolchainClusterConfig struct {
	// Defines the period in between health checks
	// +optional
	HealthCheckPeriod *string `json:"healthCheckPeriod,omitempty"`

	// Defines the timeout for each health check
	// +optional
	HealthCheckTimeout *string `json:"healthCheckTimeout,omitempty"`
}

// Defines all parameters concerned with the Webhook
// +k8s:openapi-gen=true
type WebhookConfig struct {
	// Defines the flag that determines whether to deploy the Webhook.
	// If the deploy flag is set to False and the Webhook was deployed previously it will be deleted by the memberoperatorconfig controller.
	// +optional
	Deploy *bool `json:"deploy,omitempty"`

	// Defines all secrets related to webhook configuration
	// +optional
	Secret *WebhookSecret `json:"secret,omitempty"`
}

// WebhookSecret defines all secrets related to webhook configuration
// +k8s:openapi-gen=true
type WebhookSecret struct {
	// The reference to the secret that is expected to contain the keys below
	// +optional
	ToolchainSecret `json:",inline"`

	// The key in the secret values map that contains a comma-separated list of SSH keys
	// +optional
	VirtualMachineAccessKey *string `json:"virtualMachineAccessKey,omitempty"`
}

// Defines all parameters concerned with member status
// +k8s:openapi-gen=true
type MemberStatusConfig struct {
	// Defines the period between refreshes of the member status
	// +optional
	RefreshPeriod *string `json:"refreshPeriod,omitempty"`

	// Defines all secrets related to GitHub authentication/integration
	// +optional
	GitHubSecret GitHubSecret `json:"gitHubSecret,omitempty"`
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
