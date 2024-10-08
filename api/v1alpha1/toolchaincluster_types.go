package v1alpha1

// Most of the code was copied from the KubeFedCluster CRD of the KubeFed project https://github.com/kubernetes-sigs/kubefed

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of a cluster.
const (
	// ToolchainClusterOffline means the cluster is temporarily down or not reachable
	ToolchainClusterOffline ConditionType = "Offline"

	ToolchainClusterClusterReadyReason        = "ClusterReady"
	ToolchainClusterClusterNotReadyReason     = "ClusterNotReady"
	ToolchainClusterClusterNotReachableReason = "ClusterNotReachable"
	ToolchainClusterClusterReachableReason    = "ClusterReachable"

	// ToolchainClusterLabel is the label on the Secret containing the credentials to connect
	// to the cluster represented by the ToolchainCluster object.
	ToolchainClusterLabel = LabelKeyPrefix + "toolchain-cluster"
)

// ToolchainClusterSpec defines the desired state of ToolchainCluster
// +k8s:openapi-gen=true
type ToolchainClusterSpec struct {
	// The API endpoint of the member cluster. This can be a hostname,
	// hostname:port, IP or IP:port.
	//
	// Be aware that this is kept in the spec only for compatibility reasons
	// and doesn't serve any purpose. Use the Status.APIEndpoint instead.
	//
	// Deprecated: This is not used for anything.
	// +optional
	APIEndpoint string `json:"apiEndpoint,omitempty"`

	// Name of the secret containing the kubeconfig required to connect
	// to the cluster.
	SecretRef LocalSecretReference `json:"secretRef"`
}

// LocalSecretReference is a reference to a secret within the enclosing
// namespace.
// +k8s:openapi-gen=true
type LocalSecretReference struct {
	// Name of a secret within the enclosing
	// namespace
	Name string `json:"name"`
}

// ToolchainClusterStatus contains information about the current status of a
// cluster updated periodically by cluster controller.
// +k8s:openapi-gen=true
type ToolchainClusterStatus struct {
	// APIEndpoint is the API endpoint of the remote cluster. This can be a hostname,
	// hostname:port, IP or IP:port.
	// +optional
	APIEndpoint string `json:"apiEndpoint"`

	// OperatorNamespace is the namespace in which the operator runs in the remote cluster
	// +optional
	OperatorNamespace string `json:"operatorNamespace"`

	// Conditions is an array of current cluster conditions.
	// +listType=atomic
	Conditions []Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ToolchainCluster configures Toolchain to be aware of a Kubernetes
// cluster and encapsulates the details necessary to communicate with
// the cluster.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:path=toolchainclusters
// +kubebuilder:printcolumn:name=age,type=date,JSONPath=.metadata.creationTimestamp
// +kubebuilder:printcolumn:name=ready,type=string,JSONPath=.status.conditions[?(@.type=='Ready')].status
// +kubebuilder:subresource:status
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Toolchain Cluster"
type ToolchainCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ToolchainClusterSpec `json:"spec"`
	// +optional
	Status ToolchainClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ToolchainClusterList contains a list of ToolchainCluster
type ToolchainClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolchainCluster{}, &ToolchainClusterList{})
}
