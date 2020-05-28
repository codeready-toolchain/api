package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// These are valid conditions of a UserSignup

	// UserSignupApproved reflects whether the signup request has been approved or not
	UserSignupApproved ConditionType = "Approved"
	// UserSignupComplete means provisioning is complete
	UserSignupComplete ConditionType = "Complete"

	// UserSignupUserEmailAnnotationKey is used for the usersignup email annotations key
	UserSignupUserEmailAnnotationKey = LabelKeyPrefix + "user-email"

	// UserSignupUserEmailHashLabelKey is used for the usersignup email hash label key
	UserSignupUserEmailHashLabelKey = LabelKeyPrefix + "email-hash"

	// Status condition reasons
	UserSignupNoClusterAvailableReason             = "NoClusterAvailable"
	UserSignupNoTemplateTierAvailableReason        = "NoTemplateTierAvailable"
	UserSignupFailedToReadUserApprovalPolicyReason = "FailedToReadUserApprovalPolicy"
	UserSignupUnableToCreateMURReason              = "UnableToCreateMUR"
	UserSignupUnableToDeleteMURReason              = "UnableToDeleteMUR"
	UserSignupUserDeactivatingReason               = "Deactivating"
	UserSignupUserDeactivatedReason                = "Deactivated"
	UserSignupInvalidMURStateReason                = "InvalidMURState"
	UserSignupApprovedAutomaticallyReason          = "ApprovedAutomatically"
	UserSignupApprovedByAdminReason                = "ApprovedByAdmin"
	UserSignupPendingApprovalReason                = "PendingApproval"
	UserSignupUserBanningReason                    = "Banning"
	UserSignupUserBannedReason                     = "Banned"
	UserSignupFailedToReadBannedUsersReason        = "FailedToReadBannedUsers"
	UserSignupMissingUserEmailAnnotationReason     = "MissingUserEmailAnnotation"
	UserSignupMissingEmailHashLabelReason          = "MissingEmailHashLabel"
	UserSignupInvalidEmailHashLabelReason          = "InvalidEmailHashLabel"
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

	// Deactivated is used to deactivate the user.  If not set, then by default the user is active
	// +optional
	Deactivated bool `json:"deactivated,omitempty"`

	// The user's username, obtained from the identity provider.
	Username string `json:"username"`

	// The user's first name, obtained from the identity provider.
	// +optional
	GivenName string `json:"givenName,omitempty"`

	// The user's last name, obtained from the identity provider.
	// +optional
	FamilyName string `json:"familyName,omitempty"`

	// The user's company name, obtained from the identity provider.
	// +optional
	Company string `json:"company,omitempty"`
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
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// CompliantUsername is used to store the transformed, DNS-1123 compliant username
	// +optional
	CompliantUsername string `json:"compliantUsername,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserSignup registers a user in the CodeReady Toolchain
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="User ID",type="string",JSONPath=`.spec.userID`,priority=1
// +kubebuilder:printcolumn:name="Username",type="string",JSONPath=`.spec.username`
// +kubebuilder:printcolumn:name="First Name",type="string",JSONPath=`.spec.givenName`,priority=1
// +kubebuilder:printcolumn:name="Last Name",type="string",JSONPath=`.spec.familyName`,priority=1
// +kubebuilder:printcolumn:name="Company",type="string",JSONPath=`.spec.company`,priority=1
// +kubebuilder:printcolumn:name="TargetCluster",type="string",JSONPath=`.spec.targetCluster`,priority=1
// +kubebuilder:printcolumn:name="Complete",type="string",JSONPath=`.status.conditions[?(@.type=="Complete")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Complete")].reason`
// +kubebuilder:printcolumn:name="Approved",type="string",JSONPath=`.status.conditions[?(@.type=="Approved")].status`,priority=1
// +kubebuilder:printcolumn:name="ApprovedBy",type="string",JSONPath=`.status.conditions[?(@.type=="Approved")].reason`,priority=1
// +kubebuilder:printcolumn:name="Deactivated",type="string",JSONPath=`.spec.deactivated`,priority=1
// +kubebuilder:printcolumn:name="CompliantUsername",type="string",JSONPath=`.status.compliantUsername`
// +kubebuilder:printcolumn:name="Email",type="string",JSONPath=`.metadata.annotations.toolchain\.dev\.openshift\.com/user-email`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="User Signup"
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
