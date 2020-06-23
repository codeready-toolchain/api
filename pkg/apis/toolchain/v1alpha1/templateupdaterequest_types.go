package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NSTemplateTierNameLabelKey stores the name of NSTemplateTier that this TemplateUpdateRequest is related to
const NSTemplateTierNameLabelKey = LabelKeyPrefix + "nstemplatetier"

// These are valid conditions of a TemplateUpdateRequest
const (
	// TemplateUpdateRequestComplete when the MasterUserRecord has been updated (via the TemplateUpdateRequest)
	// (for the Template Update Request, "complete" makes more sense than the usual "ready" condition type)
	TemplateUpdateRequestComplete ConditionType = "Complete"
)

// TemplateUpdateRequestSpec defines the desired state of TemplateUpdateRequest
// It contains the new TemplateRefs to use in the MasterUserRecords
// +k8s:openapi-gen=true
type TemplateUpdateRequestSpec struct {
	// The name of the tier to be updated
	TierName string `json:"tierName"`

	// The namespace templates
	// +listType=atomic
	Namespaces []NSTemplateTierNamespace `json:"namespaces"`

	// the cluster resources template (for cluster-wide quotas, etc.)
	// +optional
	ClusterResources *NSTemplateTierClusterResources `json:"clusterResources,omitempty"`
}

// TemplateUpdateRequestStatus defines the observed state of TemplateUpdateRequest
// +k8s:openapi-gen=true
type TemplateUpdateRequestStatus struct {
	// Conditions is an array of current TemplateUpdateRequest conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TemplateUpdateRequest is the Schema for the templateupdaterequests API
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=templateupdaterequests,scope=Namespaced
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Template UpdateRequest"
type TemplateUpdateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TemplateUpdateRequestSpec `json:"spec,omitempty"`

	Status TemplateUpdateRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TemplateUpdateRequestList contains a list of TemplateUpdateRequest
type TemplateUpdateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TemplateUpdateRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TemplateUpdateRequest{}, &TemplateUpdateRequestList{})
}
