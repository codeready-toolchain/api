// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecord":           schema_pkg_apis_toolchain_v1alpha1_MasterUserRecord(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordSpec":       schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordStatus":     schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSet":              schema_pkg_apis_toolchain_v1alpha1_NSTemplateSet(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec":          schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetStatus":        schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTier":             schema_pkg_apis_toolchain_v1alpha1_NSTemplateTier(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierSpec":         schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierStatus":       schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccount":                schema_pkg_apis_toolchain_v1alpha1_UserAccount(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpec":            schema_pkg_apis_toolchain_v1alpha1_UserAccountSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatus":          schema_pkg_apis_toolchain_v1alpha1_UserAccountStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequest":       schema_pkg_apis_toolchain_v1alpha1_UserProvisionRequest(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequestSpec":   schema_pkg_apis_toolchain_v1alpha1_UserProvisionRequestSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequestStatus": schema_pkg_apis_toolchain_v1alpha1_UserProvisionRequestStatus(ref),
	}
}

func schema_pkg_apis_toolchain_v1alpha1_MasterUserRecord(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MasterUserRecord is the Schema for the masteruserrecords API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MasterUserRecordSpec defines the desired state of MasterUserRecord",
				Properties: map[string]spec.Schema{
					"userID": {
						SchemaProps: spec.SchemaProps{
							Description: "UserID is the user ID from RHD Identity Provider token (“sub” claim)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"disabled": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding user should not be able to login (but the underlying UserAccounts still exists) \"false\" is assumed by default",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"deprovisioned": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding UserAccount should be deleted \"false\" is assumed by default",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"userAccounts": {
						SchemaProps: spec.SchemaProps{
							Description: "The list of user accounts in the member clusters which belong to this MasterUserRecord",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountEmbedded"),
									},
								},
							},
						},
					},
				},
				Required: []string{"userID"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountEmbedded"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MasterUserRecordStatus defines the observed state of MasterUserRecord",
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current Master User Record conditions Supported condition types: Provisioning, UserAccountNotReady and Ready",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
					"userAccounts": {
						SchemaProps: spec.SchemaProps{
							Description: "The status of user accounts in the member clusters which belong to this MasterUserRecord",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatusEmbedded"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatusEmbedded"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateSet is the Schema for the nstemplatesets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateSetSpec defines the desired state of NSTemplateSet",
				Properties: map[string]spec.Schema{
					"tierName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the tier represented by this template set",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespaces": {
						SchemaProps: spec.SchemaProps{
							Description: "The namespace templates",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Namespace"),
									},
								},
							},
						},
					},
				},
				Required: []string{"tierName", "namespaces"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Namespace"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateSetStatus defines the observed state of NSTemplateSet",
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "String representation of the overall observed status. For example: provisioning|provisioned|updating",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespaces": {
						SchemaProps: spec.SchemaProps{
							Description: "The detailed namespace statuses",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NamespaceStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NamespaceStatus"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateTier(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateTier is the Schema for the nstemplatetiers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateTierSpec defines the desired state of NSTemplateTier",
				Properties: map[string]spec.Schema{
					"namespaces": {
						SchemaProps: spec.SchemaProps{
							Description: "The namespace templates",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Namespace"),
									},
								},
							},
						},
					},
				},
				Required: []string{"namespaces"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Namespace"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateTierStatus defines the observed state of NSTemplateTier",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccount(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccount is the Schema for the useraccounts API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccountSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccountSpec defines the desired state of UserAccount",
				Properties: map[string]spec.Schema{
					"userID": {
						SchemaProps: spec.SchemaProps{
							Description: "UserID is the user ID from RHD Identity Provider token (“sub” claim) Is to be used to create Identity and UserIdentityMapping resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"disabled": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding user should not be able to login \"false\" is assumed by default",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"nsLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "The namespace limit name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nsTemplateSet": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace template set",
							Ref:         ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"),
						},
					},
				},
				Required: []string{"userID", "nsLimit", "nsTemplateSet"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccountStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccountStatus defines the observed state of UserAccount",
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current User Account conditions Supported condition types: Provisioning, UserNotReady, IdentityNotReady, UserIdentityMappingNotReady, NSTemplateSetNotReady and Ready",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserProvisionRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserProvisionRequest is the Schema for the userprovisionrequests API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequestSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequestStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequestSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserProvisionRequestStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserProvisionRequestSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserProvisionRequestSpec defines the desired state of UserProvisionRequest",
				Properties: map[string]spec.Schema{
					"userID": {
						SchemaProps: spec.SchemaProps{
							Description: "UserID is the user ID from RHD Identity Provider token (“sub” claim)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"targetCluster": {
						SchemaProps: spec.SchemaProps{
							Description: "The cluster in which the user is provisioned in If not set then the target cluster will be picked automatically",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"approved": {
						SchemaProps: spec.SchemaProps{
							Description: "If Approved set to 'true' then the user has been manually approved\n\tIf not set then the user is subject of auto-approval (if enabled)",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"userID"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserProvisionRequestStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserProvisionRequestStatus defines the observed state of UserProvisionRequest",
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current UserProvisionRequest conditions Supported condition types: PendingApproval, Provisioning, Complete",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}
