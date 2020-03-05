package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// These are valid conditions of a ChangeTierRequest

	// ChangeTierRequestComplete means that the change of the tier is complete
	ChangeTierRequestComplete ConditionType = "Complete"

	// MurNameLabelKey stores the name of MasterUserRecord the tier was changed for
	MurNameLabelKey = LabelKeyPrefix + "murname"

	// Status condition reasons
	ChangeTierRequestChangedReason     = "Changed"
	ChangeTierRequestChangeFiledReason = "ChangeFailed"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChangeTierRequestSpec defines the desired state of ChangeTierRequest
// +k8s:openapi-gen=true
type ChangeTierRequestSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The cluster to define UserAccount whose tier should be changed.
	// Optional. If not set then update all the UserAccounts in the MasterUserRecord.
	// +optional
	TargetCluster string `json:"targetCluster,omitempty"`

	// The murName is a name of MUR/UserAccount whose tier should be changed.
	MurName string `json:"murName"`

	// The tier name the tier should be changed to.
	TierName string `json:"tierName"`
}

// ChangeTierRequestStatus defines the observed state of ChangeTierRequest
// +k8s:openapi-gen=true
type ChangeTierRequestStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current ChangeTierRequest conditions
	// Supported condition types:
	// Complete
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChangeTierRequest requests for changing tier for MUR/UserAccount
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="MurName",type="string",JSONPath=`.spec.murName`
// +kubebuilder:printcolumn:name="TierName",type="string",JSONPath=`.spec.tierName`
// +kubebuilder:printcolumn:name="TargetCluster",type="string",JSONPath=`.spec.targetCluster`,priority=1
// +kubebuilder:printcolumn:name="Complete",type="string",JSONPath=`.status.conditions[?(@.type=="Complete")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Complete")].reason`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Change Tier Request"
type ChangeTierRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChangeTierRequestSpec   `json:"spec,omitempty"`
	Status ChangeTierRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChangeTierRequestList contains a list of ChangeTierRequest
type ChangeTierRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChangeTierRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChangeTierRequest{}, &ChangeTierRequestList{})
}
