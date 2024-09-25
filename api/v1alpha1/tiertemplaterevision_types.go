package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

	Spec TierTemplateRevisionSpec `json:"spec,omitempty"`
}

// TierTemplateRevisionSpec defines the desired state of TierTemplateRevision
// +k8s:openapi-gen=true
type TierTemplateRevisionSpec struct {
	// TemplateObjects contains list of Unstructured Objects that can be parsed at runtime and will be applied as part of the tier provisioning.
	// The template parameters values will be defined in the NSTemplateTier CRD.
	// +optional
	// +listType=atomic
	// +kubebuilder:pruning:PreserveUnknownFields
	TemplateObjects []runtime.RawExtension `json:"templateObjects,omitempty" protobuf:"bytes,3,opt,name=templateObjects"`
}

//+kubebuilder:object:root=true

// TierTemplateRevisionList contains a list of TierTemplateRevisions
type TierTemplateRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TierTemplateRevision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TierTemplateRevision{}, &TierTemplateRevisionList{})
}
