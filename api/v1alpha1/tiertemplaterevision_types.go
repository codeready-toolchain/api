package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TierTemplateRevision is the Schema for the tiertemplaterevisions API
// +kubebuilder:resource:path=tiertemplaterevisions,scope=Namespaced
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=`.spec.type`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Template Tier Revision"
type TierTemplateRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TierTemplateSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// TierTemplateRevisionList contains a list of TierTemplateRevisions
type TierTemplateRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TierTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TierTemplateRevision{}, &TierTemplateRevisionList{})
}
