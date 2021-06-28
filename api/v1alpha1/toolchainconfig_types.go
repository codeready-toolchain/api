package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of a ToolchainConfig
const (
	ToolchainConfigSyncComplete ConditionType = "SyncComplete"

	// Status condition reasons
	// ToolchainConfigSyncedReason when the MemberOperatorConfigs were successfully synced to the member clusters
	ToolchainConfigSyncedReason = "Synced"
	// ToolchainConfigSyncFailedReason when there were failures while syncing MemberOperatorConfigs to the member clusters
	ToolchainConfigSyncFailedReason = "SyncFailed"
)

// ToolchainConfigSpec contains all configuration for host and member operators
// +k8s:openapi-gen=true
type ToolchainConfigSpec struct {
	// Contains all host operator configuration
	// +optional
	Host HostConfig `json:"host,omitempty"`

	// Contains all member operator configurations for all member clusters
	// +optional
	Members Members `json:"members,omitempty"`
}

// HostConfig contains all configuration parameters of the host operator
// +k8s:openapi-gen=true
type HostConfig struct {
	ChangeTierRequest ch

	TemplateUpdateRequestMaxPoolSize tu

	// Keeps parameters necessary for automatic approval
	// +optional
	AutomaticApproval AutomaticApprovalConfig `json:"automaticApproval,omitempty"`

	// Keeps parameters concerned with user deactivation
	// +optional
	Deactivation DeactivationConfig `json:"deactivation,omitempty"`

	// Keeps parameters concerned with metrics
	// +optional
	Metrics MetricsConfig `json:"metrics,omitempty"`

	// Keeps parameters concerned with notifications
	// +optional
	Notifications NotificationsConfig `json:"notifications,omitempty"`

	// Keeps parameters necessary for the registration service
	// +optional
	RegistrationService RegistrationServiceConfig `json:"registrationService,omitempty"`
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
type AutomaticApprovalConfig struct {
	// Defines if the automatic approval is enabled or not
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

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
type MaxNumberOfUsers struct {
	// It is the maximal number of users provisioned in the system overall - equals to max number of MasterUserRecords in host cluster
	// +optional
	Overall *int `json:"overall,omitempty"`

	// Contains a map of maximal number of users provisioned per member cluster mapped by the cluster name
	// - equals to max number of UserAccounts in member cluster
	// +optional
	// +mapType=atomic
	SpecificPerMemberCluster map[string]int `json:"specificPerMemberCluster,omitempty"`
}

type DeactivationConfig struct {

	// DeactivatingNotificationDays is the number of days after a pre-deactivating notification is sent that actual
	// deactivation occurs.  If this parameter is set to zero, then there will be no delay
	// +optional
	DeactivatingNotificationDays *int `json:"deactivatingNotificationDays,omitempty"`
}

type ToolchainSecret struct {

	// Reference is the name of the secret resource to look up
	// +optional
	Ref *string `json:"ref,omitempty"`
}

type MetricsConfig struct {

	// ForceSynchronization is a flag used to trigger synchronization of the metrics
	// based on the resources rather than on the content of `ToolchainStatus.status.metrics`
	// +optional
	ForceSynchronization *bool `json:"forceSynchronization,omitempty"`
}

type NotificationsConfig struct {

	// NotificationDeliveryService is notification delivery service to use for notifications
	// +optional
	NotificationDeliveryService *string `json:"notificationDeliveryService,omitempty"`

	// DurationBeforeNotificationDeletion is notification delivery service to use for notifications
	// +optional
	DurationBeforeNotificationDeletion *string `json:"durationBeforeNotificationDeletion,omitempty"`

	// The administrator email address for system notifications
	// +optional
	AdminEmail *string `json:"adminEmail,omitempty"`

	// Defines all secrets related to notification configuration
	// +optional
	Secret NotificationSecret `json:"secret,omitempty"`
}

// Defines all secrets related to notification configuration
// +k8s:openapi-gen=true
type NotificationSecret struct {
	// The reference to the secret that is expected to contain the keys below
	// +optional
	ToolchainSecret `json:",inline"`

	// The key for the host operator mailgun domain used for creating an instance of mailgun
	// +optional
	MailgunDomain *string `json:"mailgunDomain,omitempty"`

	// The key for the host operator mailgun api key used for creating an instance of mailgun
	// +optional
	MailgunAPIKey *string `json:"mailgunAPIKey,omitempty"`

	// The key for the host operator mailgun senders email
	// +optional
	MailgunSenderEmail *string `json:"mailgunSenderEmail,omitempty"`

	// The key for the reply-to email address that will be set in sent notifications
	// +optional
	MailgunReplyToEmail *string `json:"mailgunReplyToEmail,omitempty"`
}

type RegistrationServiceConfig struct {

	// RegistrationServiceURL is the URL used to a ccess the registration service
	// +optional
	RegistrationServiceURL *string `json:"registrationServiceURL,omitempty"`
}

// ToolchainConfigStatus defines the observed state of ToolchainConfig
// +k8s:openapi-gen=true
type ToolchainConfigStatus struct {

	// SyncErrors is a map of sync errors indexed by toolchaincluster name that indicates whether
	// an attempt to sync configuration to a member cluster failed
	// +optional
	// +mapType=atomic
	SyncErrors map[string]string `json:"syncErrors,omitempty"`

	// Conditions is an array of the current ToolchainConfig conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

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

//+kubebuilder:object:root=true

// ToolchainConfigList contains a list of ToolchainConfig
type ToolchainConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolchainConfig{}, &ToolchainConfigList{})
}
