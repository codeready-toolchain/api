package v1alpha1

// SpaceNamespace is a common type to define the information about a namespace within a Space
// Used in NSTemplateSet status and Workspace status
// +k8s:openapi-gen=true
type SpaceNamespace struct {

	// Name the name of the namespace.
	// +optional
	Name string `json:"name,omitempty"`

	// Type the type of the namespace. eg. default
	// +optional
	Type string `json:"type,omitempty"`
}
