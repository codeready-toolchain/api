package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RegistrationServiceSpec defines the desired state of RegistrationService
// +k8s:openapi-gen=true
type RegistrationServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// The image identifies which image of the registration service should be used for a deployment
	Image string `json:"image"`

	// The number of replicas of the deployed registration service
	// +optional
	Replicas int `json:"replicas,omitempty"`

	// The environment identifies which mode the registration service should be running in - prod, stage, e2e-tests, dev, etc.
	// +optional
	Environment string `json:"environment,omitempty"`

	// The AuthClient contains all necessary information about the auth client
	AuthClient AuthClient `json:"authClient,omitempty"`
}

type AuthClient struct {
	// The LibraryUrl identifies the auth library location
	LibraryUrl string `json:"libraryUrl,omitempty"`

	// The Config contains the auth config
	Config string `json:"config,omitempty"`

	// The PublicKeysUrl identifies the public keys location
	PublicKeysUrl string `json:"publicKeysUrl,omitempty"`
}

// RegistrationServiceStatus defines the observed state of RegistrationService
// +k8s:openapi-gen=true
type RegistrationServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Registration Service deployment conditions
	// Supported condition types:
	// Deploying, and Deployed
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RegistrationService is the Schema for the registrationservices API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=registrationservices,scope=Namespaced
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
