package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of a MasterUserRecord
const (
	RegistrationServiceDeployingReason       = "Deploying"
	RegistrationServiceDeployingFailedReason = "DeployingFailed"
	RegistrationServiceDeployedReason        = "Deployed"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RegistrationServiceSpec defines the desired state of RegistrationService
// +k8s:openapi-gen=true
type RegistrationServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// The environment variables are supposed to be set to registration service deployment template
	// +optional
	EnvironmentVariables map[string]string `json:"envVars,omitempty"`
}

// RegistrationServiceStatus defines the observed state of RegistrationService
// +k8s:openapi-gen=true
type RegistrationServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Registration Service deployment conditions
	// Supported condition reasons:
	// Deploying, and Deployed
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RegistrationService configures the registration service deployment
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=registrationservices,scope=Namespaced
// +kubebuilder:resource:shortName=rs
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=`.spec.envVars.IMAGE`
// +kubebuilder:printcolumn:name="Environment",type="string",JSONPath=`.spec.envVars.ENVIRONMENT`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Registration Service"
type RegistrationService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistrationServiceSpec   `json:"spec,omitempty"`
	Status RegistrationServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RegistrationServiceList contains a list of RegistrationService
type RegistrationServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistrationService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RegistrationService{}, &RegistrationServiceList{})
}
