package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SpaceLabelKey is used to label the SpaceBinding with the name of the Space it is bound to
	SpaceLabelKey = LabelKeyPrefix + "space"

	// MurLabelKey is used to label the SpaceBinding with the name of the MasterUserRecord it belongs to
	MurLabelKey = LabelKeyPrefix + "mur"
)

// SpaceBindingSpec defines the desired state of SpaceBinding
// +k8s:openapi-gen=true
type SpaceBindingSpec struct {

	// The MurName is a name of the MasterUserRecord this SpaceBinding belongs to.
	MurName string `json:"murName"`

	// The SpaceName is a name of the Space this SpaceBinding is bound to.
	SpaceName string `json:"spaceName"`

	// The SpaceRole is a name of the SpaceRole that is granted to the user for the Space. For example: admin, view, ...
	SpaceRole string `json:"spaceRole"`
}

// SpaceBindingStatus defines the observed state of SpaceBinding
// +k8s:openapi-gen=true
type SpaceBindingStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// SpaceBinding is the Schema for the spacebindings API which defines relationship between Spaces and MasterUserRecords
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="MUR",type="string",JSONPath=`.spec.murName`
// +kubebuilder:printcolumn:name="Space",type="string",JSONPath=`.spec.spaceName`
// +kubebuilder:printcolumn:name="SpaceRole",type="string",JSONPath=`.spec.spaceRole`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="SpaceBinding"
type SpaceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceBindingSpec   `json:"spec,omitempty"`
	Status SpaceBindingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpaceBindingList contains a list of SpaceBinding
// +k8s:openapi-gen=true
type SpaceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceBinding{}, &SpaceBindingList{})
}
