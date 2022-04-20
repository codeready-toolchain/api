package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UserTier contains user-specific configuration
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="User Tier"
type UserTier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserTierSpec   `json:"spec,omitempty"`
	Status UserTierStatus `json:"status,omitempty"`
}

// UserTierSpec defines the desired state of UserTier
// +k8s:openapi-gen=true
type UserTierSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// the period (in days) after which users within the tier will be deactivated
	// +optional
	DeactivationTimeoutDays int `json:"deactivationTimeoutDays,omitempty"`
}

// UserTierStatus defines the observed state of UserTier
// +k8s:openapi-gen=true
type UserTierStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current UserTier conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// UserTierHistory a track record of an update
type UserTierHistory struct {
	// StartTime is the time when the UserTier was updated
	UpdateTime metav1.Time `json:"startTime"`
	// Hash the hash matching of the UserTier spec
	Hash string `json:"hash"`
}

//+kubebuilder:object:root=true

// UserTierList contains a list of UserTier
type UserTierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserTier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserTier{}, &UserTierList{})
}
