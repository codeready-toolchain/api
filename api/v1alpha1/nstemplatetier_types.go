package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NSTemplateTierSpec defines the desired state of NSTemplateTier
// +k8s:openapi-gen=true
type NSTemplateTierSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// The namespace templates
	// +listType=atomic
	Namespaces []NSTemplateTierNamespace `json:"namespaces"`

	// the cluster resources template (for cluster-wide quotas, etc.)
	// +optional
	ClusterResources *NSTemplateTierClusterResources `json:"clusterResources,omitempty"`

	// the templates to set the spaces roles, indexed by role
	// +optional
	// +mapType=atomic
	SpaceRoles map[string]NSTemplateTierSpaceRole `json:"spaceRoles,omitempty"`

	// SpaceRequestConfig stores all the configuration related to the Space Request feature
	// +optional
	SpaceRequestConfig *SpaceRequestConfig `json:"spaceRequestConfig,omitempty"`

	// Parameters is an optional array of Parameters used during the NSTemplateTier and TierTemplate creation.
	// When creating the NsTemplateTier and referenced TierTemplates, the parameters will be read from the NSTemplateTier and evaluated in all the TierTemplates referenced in the spec.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Parameters []Parameter `json:"parameters,omitempty" protobuf:"bytes,4,opt,name=parameters" patchStrategy:"merge" patchMergeKey:"name"`
}

// Parameter defines a name/value variable that is to be processed during
// TierTemplate creation.
type Parameter struct {
	// Name must be set and it can be referenced in the TierTemplate
	// items using {{.NAME}}
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Value holds the Parameter data.
	// The value replaces all occurrences of the Parameter {{.NAME}}.
	Value string `json:"value" protobuf:"bytes,4,opt,name=value"`
}

// SpaceRequestConfig contains all the configuration related to the Space Request feature
// +k8s:openapi-gen=true
type SpaceRequestConfig struct {
	// Provides the name of the Service Account whose token is to be copied
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
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

// NSTemplateTierSpaceRole the space roles definition in an NSTemplateTier resource
type NSTemplateTierSpaceRole struct {
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

	// Updates is an array of all NSTemplateTier updates
	// +optional
	// +patchMergeKey=startTime
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=startTime
	Updates []NSTemplateTierHistory `json:"updates,omitempty" patchStrategy:"merge" patchMergeKey:"startTime"`

	// Revisions is an array containing the names of the last applied TierTemplateRevision CRs.
	// Since the TierTemplate can be managed outside the host-operator,
	// we need to store the TierTemplate revisions, so we can get to the original contents to be able to properly propagate the update from the previous version of the template to the new one.
	// +optional
	// +mapType=atomic
	Revisions map[string]string `json:"revisions,omitempty"`
}

// NSTemplateTierHistory a track record of an update
type NSTemplateTierHistory struct {
	// StartTime is the time when the NSTemplateTier was updated
	StartTime metav1.Time `json:"startTime"`
	// Hash the hash matching on the templateRefs in the resource spec
	Hash string `json:"hash"`
	// CompletionTime is the time when the last MasterUserRecord was updated
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	// Failures is the number of MasterUserRecords which failed to be updated
	Failures int `json:"failures"`
	// FailedAccounts
	// +optional
	FailedAccounts []string `json:"failedAccounts,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

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

//+kubebuilder:object:root=true

// NSTemplateTierList contains a list of NSTemplateTier
type NSTemplateTierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NSTemplateTier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NSTemplateTier{}, &NSTemplateTierList{})
}
