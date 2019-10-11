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
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UserSignupSpec defines the desired state of UserSignup
// +k8s:openapi-gen=true
type UserSignupSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// UserID is the user ID from RHD Identity Provider token (“sub” claim)
	UserID string `json:"userID"`

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

	// The DNS-1123-compliant username.  This username is generated by the registration app and may differ from the
	// Username, as a limited character set is available for naming (see RFC1123).  If the username contains characters
	// which are disqualified, the username is transformed into a compliant name instead.
	// For example, johnsmith@redhat.com -> johnsmith-at-redhat-com
	CompliantUsername string `json:compliantUserName`
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
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserSignup is the Schema for the usersignup API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="User ID",type="string",JSONPath=".spec.userID",priority=1
// +kubebuilder:printcolumn:name="Username",type="string",JSONPath=".spec.username",priority=1
// +kubebuilder:printcolumn:name="TargetCluster",type="string",JSONPath=".spec.targetCluster",priority=1
// +kubebuilder:printcolumn:name="Complete",type="string",JSONPath=".status.conditions[?(@.type=="Complete")].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=="Complete")].reason"
// +kubebuilder:printcolumn:name="Approved",type="string",JSONPath=".status.conditions[?(@.type=="Approved")].status",priority=1
// +kubebuilder:printcolumn:name="ApprovedBy",type="string",JSONPath=".status.conditions[?(@.type=="Approved")].reason",priority=1
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
