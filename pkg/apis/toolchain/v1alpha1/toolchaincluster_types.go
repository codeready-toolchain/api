package v1alpha1

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Most of the code was copied from the KubeFedCluster CRD of the KubeFed project https://github.com/kubernetes-sigs/kubefed

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ToolchainClusterConditionType string

// These are valid conditions of a cluster.
const (
	// ToolchainClusterReady means the cluster is ready to accept workloads.
	ToolchainClusterReady ToolchainClusterConditionType = "Ready"
	// ClusterOffline means the cluster is temporarily down or not reachable
	ToolchainClusterOffline ToolchainClusterConditionType = "Offline"

	ToolchainClusterClusterReadyReason        = "ClusterReady"
	ToolchainClusterClusterNotReadyReason     = "ClusterNotReady"
	ToolchainClusterClusterNotReachableReason = "ClusterNotReachable"
	ToolchainClusterClusterReachableReason    = "ClusterReachable"
)

type TLSValidation string

const (
	TLSAll            TLSValidation = "*"
	TLSSubjectName    TLSValidation = "SubjectName"
	TLSValidityPeriod TLSValidation = "ValidityPeriod"
)

// ToolchainClusterSpec defines the desired state of ToolchainCluster
// +k8s:openapi-gen=true
type ToolchainClusterSpec struct {
	// The API endpoint of the member cluster. This can be a hostname,
	// hostname:port, IP or IP:port.
	APIEndpoint string `json:"apiEndpoint"`

	// CABundle contains the certificate authority information.
	// +optional
	CABundle string `json:"caBundle,omitempty"`

	// Name of the secret containing the token required to access the
	// member cluster. The secret needs to exist in the same namespace
	// as the control plane and should have a "token" key.
	SecretRef LocalSecretReference `json:"secretRef"`

	// DisabledTLSValidations defines a list of checks to ignore when validating
	// the TLS connection to the member cluster.  This can be any of *, SubjectName, or ValidityPeriod.
	// If * is specified, it is expected to be the only option in list.
	// +optional
	// +listType=set
	DisabledTLSValidations []TLSValidation `json:"disabledTLSValidations,omitempty"`
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
	// Conditions is an array of current cluster conditions.
	// +listType=atomic
	Conditions []ToolchainClusterCondition `json:"conditions"`
	// Zones are the names of availability zones in which the nodes of the cluster exist, e.g. 'us-east1-a'.
	// +optional
	// +listType=set
	Zones []string `json:"zones,omitempty"`
	// Region is the name of the region in which all of the nodes in the cluster exist.  e.g. 'us-east1'.
	// +optional
	Region *string `json:"region,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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

// ToolchainClusterCondition describes current state of a cluster.
// +k8s:openapi-gen=true
type ToolchainClusterCondition struct {
	// Type of cluster condition, Ready or Offline.
	Type ToolchainClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status apiv1.ConditionStatus `json:"status"`
	// Last time the condition was checked.
	LastProbeTime metav1.Time `json:"lastProbeTime"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainClusterList contains a list of ToolchainCluster
type ToolchainClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ToolchainCluster{}, &ToolchainClusterList{})
}
