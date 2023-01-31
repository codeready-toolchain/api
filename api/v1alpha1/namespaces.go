package v1alpha1

// Namespace it's used to define an end user namespace type.
// At the moment there is only one type (default)
type Namespace struct {
	// Name of the namespace
	Name string `json:"name"`
	// Type of the namespace
	Type string `json:"type"`
}
