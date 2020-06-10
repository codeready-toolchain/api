package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NSTemplateTierNameLabelKey stores the name of NSTemplateTier that this TemplateUpdateRequest is related to
const NSTemplateTierNameLabelKey = LabelKeyPrefix + "nstemplatetier"

// TemplateUpdateRequestSpec defines the desired state of TemplateUpdateRequest
// It contains the new TemplateRefs to
type TemplateUpdateRequestSpec struct {
	TierName           string           `json:"tierName"`
	NSTemplateTierSpec `json:",inline"` // TODO: factorize this struct with NSTemplateTierSpec and NSTemplateSetSpec
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TemplateUpdateRequest is the Schema for the templateupdaterequests API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=templateupdaterequests,scope=Namespaced
type TemplateUpdateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TemplateUpdateRequestSpec `json:"spec,omitempty"`
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
