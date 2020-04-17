package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NotificationTemplateSpec defines the template used for generating notifications
// +k8s:openapi-gen=true
type NotificationTemplateSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Subject is the subject line (e.g. in an email) for the notification
	Subject string `json:"subject"`

	// Content defines the content of the notification
	Content string `json:"content"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationTemplate defines a notification template in the CodeReady Toolchain
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Subject",type="string",JSONPath=`.spec.subject`,priority=1
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="NotificationTemplate"
type NotificationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NotificationTemplateSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationTemplateList contains a list of NotificationTemplate
type NotificationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationTemplate{}, &NotificationTemplateList{})
}
