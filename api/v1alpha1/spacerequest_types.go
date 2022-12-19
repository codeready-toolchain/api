package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpaceRequestSpec defines the desired state of Space
// +k8s:openapi-gen=true
type SpaceRequestSpec struct {
	// TierName is a required property introduced to retain the name of the tier
	// for which this Space is provisioned
	// If not set then the tier name will be set automatically
	// +optional
	TierName string `json:"tierName,omitempty"`

	// TargetClusterLabels one or more labels that define a set of clusters
	// where the Space can be provisioned.
	// +optional
	TargetClusterLabels map[string]string `json:"targetClusterLabels,omitempty"`
}

// SpaceRequestStatus defines the observed state of Space
// +k8s:openapi-gen=true
type SpaceRequestStatus struct {

	// TargetClusterURL The API URL of the cluster where Space is currently provisioned
	// Can be empty if provisioning did not start or failed
	// To be used to de-provision the NSTemplateSet if the Spec.TargetCluster is either changed or removed
	// +optional
	TargetClusterURL string `json:"targetClusterURL,omitempty"`

	// The status of user accounts in the member clusters which belong to this MasterUserRecord
	// +listType=atomic
	NamespaceAccess []NamespaceAccess `json:"namespaceAccess,omitempty"`

	// Conditions is an array of current Master User Record conditions
	// Supported condition types:
	// Provisioning, UserAccountNotReady and Ready
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// NamespaceAccess defines the name of the namespace and the secret reference to access it
type NamespaceAccess struct {
	// Name is the corresponding name of the provisioned namespace
	Name string `json:"name"`
	// SecretRef is the name of the secret with a SA token that has admin-like
	// (or whatever we set in the tier template) permissions in the namespace
	SecretRef string `json:"secretRef"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SpaceRequest is the Schema for the space request API
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Tier",type="string",JSONPath=`.spec.tierName`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="SpaceRequest"
type SpaceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceRequestSpec   `json:"spec,omitempty"`
	Status SpaceRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpaceRequestList contains a list of SpaceRequests
type SpaceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceRequest{}, &SpaceRequestList{})
}
