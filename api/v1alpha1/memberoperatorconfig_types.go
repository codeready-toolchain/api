package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemberOperatorConfigSpec contains all configuration parameters of the member operator
// +k8s:openapi-gen=true
type MemberOperatorConfigSpec struct {
	// Keeps parameters concerned with authentication
	// +optional
	Auth AuthConfig `json:"authConfig,omitempty"`

	// Keeps parameters concerned with the autoscaler
	// +optional
	Autoscaler AutoscalerConfig `json:"autoscalerConfig,omitempty"`

	// Keeps parameters concerned with Che/CRW
	// +optional
	Che CheConfig `json:"che,omitempty"`

	// Keeps parameters concerned with the console
	// +optional
	Console ConsoleConfig `json:"console,omitempty"`

	// Keeps parameters concerned with member status
	// +optional
	MemberStatus MemberStatusConfig `json:"memberStatus,omitempty"`

	// Keeps parameters concerned with the toolchaincluster
	// +optional
	ToolchainCluster ToolchainClusterConfig `json:"toolchainClusterConfig,omitempty"`

	// Keeps parameters concerned with the webhook
	// +optional
	Webhook WebhookConfig `json:"webhook,omitempty"`
}

// Defines all parameters concerned with the autoscaler
// +k8s:openapi-gen=true
type AuthConfig struct {
	// Represents the configured identity provider
	IdP *string `json:"idP,omitempty"`
}

// Defines all parameters concerned with the autoscaler
// +k8s:openapi-gen=true
type AutoscalerConfig struct {
	// Defines the flag that determines whether to deploy the autoscaler buffer
	Deploy *bool `json:"deploy,omitempty"`

	// Represents how much memory should be required by the autoscaler buffer
	BufferMemory *string `json:"bufferMemory,omitempty"`

	// Represents the number of autoscaler buffer pods
	BufferReplicas *int `json:"bufferReplicas,omitempty"`
}

// Defines all parameters concerned with Che
// +k8s:openapi-gen=true
type CheConfig struct {
	// Defines the Che/CRW Keycloak route name
	KeycloakRouteName *string `json:"keycloakRouteName,omitempty"`

	// Defines the Che/CRW route name
	RouteName *string `json:"routeName,omitempty"`

	// Defines the Che/CRW operator namespace
	Namespace *string `json:"namespace,omitempty"`

	// Defines a flag that indicates whether the Che/CRW operator is required to be installed on the cluster. May be used in monitoring.
	Required *bool `json:"required,omitempty"`

	// Defines a flag to turn the Che user deletion logic on/off
	UserDeletionEnabled *bool `json:"userDeletionEnabled,omitempty"`

	Secret CheSecret `json:"cheSecret,omitempty"`
}

type CheSecret struct {
	ToolchainSecret `json:",inline"`

	CheAdminUsernameKey *string `json:"cheAdminUsernameKey,omitempty"`
	CheAdminPasswordKey *string `json:"cheAdminPasswordKey,omitempty"`
}

// Defines all parameters concerned with the console
// +k8s:openapi-gen=true
type ConsoleConfig struct {
	// Defines the console route namespace
	Namespace *string `json:"namespace,omitempty"`

	// Defines the console route name
	RouteName *string `json:"routeName,omitempty"`
}

// Defines all parameters concerned with the console
// +k8s:openapi-gen=true
type ToolchainClusterConfig struct {
	// Defines the period in between health checks
	HealthCheckPeriod *string `json:"healthCheckPeriod,omitempty"`

	// Defines the timeout for each health check
	HealthCheckTimeout *string `json:"healthCheckTimeout,omitempty"`
}

// Defines all parameters concerned with the Webhook
// +k8s:openapi-gen=true
type WebhookConfig struct {
	// Defines the flag that determines whether to deploy the Webhook
	Deploy *bool `json:"deploy,omitempty"`

	// Defines the Webhook image
	Image *string `json:"image,omitempty"`
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
