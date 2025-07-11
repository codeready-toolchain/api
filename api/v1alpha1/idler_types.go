package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of an Idler
const (
	// IdlerTriggeredNotificationCreated is used to track the status of the notification send to a user
	// when the idler is active for the very first time in user's namespace
	IdlerTriggeredNotificationCreated ConditionType = "IdlerTriggeredNotificationCreated"

	// Status condition reasons
	IdlerUnableToEnsureIdlingReason                = "UnableToEnsureIdling"
	IdlerRunningReason                             = "Running"
	IdlerTriggeredReason                           = "IdlerRunningFirstTime"
	IdlerTriggeredNotificationCreationFailedReason = "UnableToCreateIdlerNotification"
	IdlerNoDeactivationReason                      = "NoDeactivation"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IdlerSpec defines the desired state of Idler
// +k8s:openapi-gen=true
type IdlerSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// TimeoutSeconds is the number of seconds before the running pods will be deleted
	TimeoutSeconds int32 `json:"timeoutSeconds"`
}

// IdlerStatus defines the observed state of Idler
// +k8s:openapi-gen=true
type IdlerStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Idler conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

type Pod struct {
	Name      string      `json:"name"`
	StartTime metav1.Time `json:"startTime"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Idler enables automatic idling of payloads in a user namespaces
// where the name of the Idler matches the name of the corresponding namespace.
// For example an Idler with "foo" name will be managing pods in namespace "foo".
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Timeout",type="integer",JSONPath=`.spec.timeoutSeconds`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Idler"
type Idler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdlerSpec   `json:"spec,omitempty"`
	Status IdlerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IdlerList contains a list of Idlers
type IdlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Idler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Idler{}, &IdlerList{})
}
