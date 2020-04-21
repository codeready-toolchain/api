package v1alpha1

import (
	templatev1 "github.com/openshift/api/template/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TierTemplateSpec defines the desired state of TierTemplate
// +k8s:openapi-gen=true
type TierTemplateSpec struct {
	
	// The type of the template. For example: "code", "dev", "stage" or "cluster"
	Type string `json:"type"`

	// The revision of the corresponding template
	Revision string `json:"revision"`

	// Template the OpenShift Template to be used to provision a user's namespace or cluster-wide resources
	Template templatev1.Template `json:"template"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TierTemplate is the Schema for the tiertemplates API
// +kubebuilder:resource:path=tiertemplates,scope=Namespaced
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="Revision",type="string",JSONPath=`.spec.revision`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Template Tier"
type TierTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TierTemplateSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TierTemplateList contains a list of TierTemplate
type TierTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TierTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TierTemplate{}, &TierTemplateList{})
}
