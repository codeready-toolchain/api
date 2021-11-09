package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ToolchainEventSpec defines the parameters for a Toolchain event, such as a training session or workshop. Users
// may register for the event by using the event's unique activation code
//
// +k8s:openapi-gen=true
type ToolchainEventSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The timestamp from which users may register via this event's activation code
	StartTime metav1.Time `json:"startTime"`

	// The timestamp after which users may no longer register via this event's activation code
	EndTime metav1.Time `json:"endTime"`

	// An optional description that may be provided describing the purpose of the event
	// +optional
	Description string `json:"description,omitempty"`

	// The maximum number of attendees
	MaxAttendees int `json:"maxAttendees"`

	// The tier to assign to users registering for the event.  This must be the valid name of an nstemplatetier resource.
	// +optional
	Tier string `json:"tier,omitempty"`

	// The unique activation code for the event
	ActivationCode string `json:"activationCode"`

	// If true, best effort is made to provision all attendees of the event on the same cluster
	// +optional
	PreferSameCluster bool `json:"preferSameCluster"`
}

// ToolchainEventStatus defines the observed state of ToolchainEvent
// +k8s:openapi-gen=true
type ToolchainEventStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current ToolchainEventStatus conditions
	// Supported condition types:
	// Complete
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	ActivationCount int `json:"activationCount"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ToolchainEvent registers a toolchain event in the CodeReady Toolchain
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="StartTime",type="string",JSONPath=`.spec.startTime`
// +kubebuilder:printcolumn:name="EndTime",type="string",JSONPath=`.spec.endTime`
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=`.spec.description`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Toolchain Event"
type ToolchainEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToolchainEventSpec   `json:"spec,omitempty"`
	Status ToolchainEventStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ToolchainEventList contains a list of ToolchainEvent
type ToolchainEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolchainEvent{}, &ToolchainEventList{})
}
