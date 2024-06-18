package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// ConditionReconciled is a status condition type that indicates whether the
	// CR has been successfully reconciled
	ConditionReconciled ConditionType = "Reconciled"

	// ReconciledReasonComplete indicates the CR was successfully reconciled
	ReconciledReasonComplete = "ReconcileComplete"
	// ReconciledReasonError indicates an error was encountered while
	// reconciling the CR
	ReconciledReasonError = "ReconcileError"
)

type ExportSpec struct {
	// Method download or git. This defines which process
	// to use for exporting objects from a cluster
	Method string `json:"method"`
	// Branch within the git repository
	Branch string `json:"branch,omitempty"`
	// Git repository which will be cloned and updated
	Repo string `json:"repo,omitempty"`
	// Email used to specify the user who performed the git commit
	Email string `json:"email,omitempty"`
	// Predefined secret that contains an SSH key that will
	// be used for git cloning and pushing
	Secret string `json:"secret,omitempty"`
	// Set automatically by the webhook to dictate who will
	// run the export process
	User string `json:"user,omitempty"`
	// Set automatically by the webhook to dictate the
	// group
	Group []string `json:"group,omitempty"`
	// Filter Export by labels
	LabelSelector string `json:"labelSelector,omitempty"`
}

// ExportStatus defines the observed state of Export
type ExportStatus struct {
	// Condition set by controller to signify the export completed
	// successfully and the route is available
	Completed  bool        `json:"completed,omitempty"`
	Extracted  bool        `json:"extracted,omitempty"`
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// Route that is defined by the controller to specify the
	// location of the zip file
	Route string `json:"route,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Export is the Schema for the exports API
type Export struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExportSpec   `json:"spec,omitempty"`
	Status ExportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExportList contains a list of Export
type ExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Export `json:"items"`
}

func init() {
	//SchemeBuilder.Register(&Export{}, &ExportList{})
}
