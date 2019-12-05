package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of a UserSignup
const (
	// UserSignupApproved reflects whether the signup request has been approved or not
	UserSignupApproved ConditionType = "Approved"
	// UserSignupComplete means provisioning is complete
	UserSignupComplete ConditionType = "Complete"

	// UserSignupUserEmailAnnotationKey is used for the usersignup email annotations key
	UserSignupUserEmailAnnotationKey = "toolchain.dev.openshift.com/user-email"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UserSignupSpec defines the desired state of UserSignup
// +k8s:openapi-gen=true
type UserSignupSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The cluster in which the user is provisioned in
	// If not set then the target cluster will be picked automatically
	// +optional
	TargetCluster string `json:"targetCluster,omitempty"`

	// If Approved set to 'true' then the user has been manually approved
	// If not set then the user is subject of auto-approval (if enabled)
	// +optional
	Approved bool `json:"approved,omitempty"`

	// The user's username, obtained from the identity provider.
	Username string `json:"username"`
}

// UserSignupStatus defines the observed state of UserSignup
// +k8s:openapi-gen=true
type UserSignupStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current UserSignup conditions
	// Supported condition types:
	// PendingApproval, Provisioning, Complete
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// CompliantUsername is used to store the transformed, DNS-1123 compliant username
	// +optional
	CompliantUsername string `json:"compliantUsername,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserSignup is the Schema for the usersignup API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="User ID",type="string",JSONPath=`.spec.userID`,priority=1
// +kubebuilder:printcolumn:name="Username",type="string",JSONPath=`.spec.username`
// +kubebuilder:printcolumn:name="TargetCluster",type="string",JSONPath=`.spec.targetCluster`,priority=1
// +kubebuilder:printcolumn:name="Complete",type="string",JSONPath=`.status.conditions[?(@.type=="Complete")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Complete")].reason`
// +kubebuilder:printcolumn:name="Approved",type="string",JSONPath=`.status.conditions[?(@.type=="Approved")].status`,priority=1
// +kubebuilder:printcolumn:name="ApprovedBy",type="string",JSONPath=`.status.conditions[?(@.type=="Approved")].reason`,priority=1
// +kubebuilder:printcolumn:name="CompliantUsername",type="string",JSONPath=`.status.compliantUsername`
// +kubebuilder:printcolumn:name="Email",type="string",JSONPath=`.metadata.annotations.toolchain\.dev\.openshift\.com/user-email`
type UserSignup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSignupSpec   `json:"spec,omitempty"`
	Status UserSignupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserSignupList contains a list of UserSignup
type UserSignupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserSignup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserSignup{}, &UserSignupList{})
}
