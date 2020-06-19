package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// These are valid conditions of a Notification

	// NotificationDeletionError indicates that the notification failed to be deleted
	NotificationDeletionError ConditionType = deletionError

	// NotificationSent reflects whether the notification has been sent to the user
	NotificationSent ConditionType = "Sent"

	// Status condition reasons
	NotificationSentReason          = "Sent"
	NotificationDeletionErrorReason = "UnableToDeleteNotification"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NotificationSpec defines the desired state of Notification
// +k8s:openapi-gen=true
type NotificationSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// UserID is the user ID from RHD Identity Provider token (“sub” claim).  The UserID is used by
	// the notification service (i.e. the NotificationController) to lookup the UserSignup resource for the user,
	// and extract from it the values required to generate the notification content and to deliver the notification
	UserID string `json:"userID"`

	// Template is the name of the NotificationTemplate resource that will be used to generate the notification
	Template string `json:"template"`
}

// NotificationStatus defines the observed state of Notification
// +k8s:openapi-gen=true
type NotificationStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Notification conditions
	// Supported condition types:
	// Delivered
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Notification registers a notification in the CodeReady Toolchain
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="User ID",type="string",JSONPath=`.spec.userID`,priority=1
// +kubebuilder:printcolumn:name="Delivered",type="string",JSONPath=`.status.conditions[?(@.type=="Delivered")].status`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Notification"
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotificationSpec   `json:"spec,omitempty"`
	Status NotificationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationList contains a list of Notification
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
