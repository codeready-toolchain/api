package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NSTemplateTierSpec defines the desired state of NSTemplateTier
// +k8s:openapi-gen=true
type NSTemplateTierSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The namespace templates
	// +listType=map
	// +listMapKey=type
	Namespaces []NSTemplateTierNamespace `json:"namespaces"`

	// the cluster resources template (for cluster-wide quotas, etc.)
	// +optional
	ClusterResources *NSTemplateTierClusterResources `json:"clusterResources,omitempty"`
}

// NSTemplateTierNamespace the namespace definition in an NSTemplateTier resource
type NSTemplateTierNamespace struct {
	// TemplateRef The name of the TierTemplate resource which exists in the host cluster and which contains the template to use
	TemplateRef string `json:"templateRef"`
}

// NSTemplateTierClusterResources defines the cluster-scoped resources associated with a given user
type NSTemplateTierClusterResources struct {
	// TemplateRef The name of the TierTemplate resource which exists in the host cluster and which contains the template to use
	TemplateRef string `json:"templateRef"`
}

// NSTemplateTierStatus defines the observed state of NSTemplateTier
// +k8s:openapi-gen=true
type NSTemplateTierStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current NSTemplateTier conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateTier configures user environment via templates used for namespaces the user has access to
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:shortName=tier
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Namespace Template Tier"
type NSTemplateTier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NSTemplateTierSpec   `json:"spec,omitempty"`
	Status NSTemplateTierStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSTemplateTierList contains a list of NSTemplateTier
type NSTemplateTierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NSTemplateTier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NSTemplateTier{}, &NSTemplateTierList{})
}
