// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterUserRecord) DeepCopyInto(out *MasterUserRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterUserRecord.
func (in *MasterUserRecord) DeepCopy() *MasterUserRecord {
	if in == nil {
		return nil
	}
	out := new(MasterUserRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MasterUserRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterUserRecordList) DeepCopyInto(out *MasterUserRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MasterUserRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterUserRecordList.
func (in *MasterUserRecordList) DeepCopy() *MasterUserRecordList {
	if in == nil {
		return nil
	}
	out := new(MasterUserRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MasterUserRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterUserRecordSpec) DeepCopyInto(out *MasterUserRecordSpec) {
	*out = *in
	if in.UserAccounts != nil {
		in, out := &in.UserAccounts, &out.UserAccounts
		*out = make([]UserAccountEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterUserRecordSpec.
func (in *MasterUserRecordSpec) DeepCopy() *MasterUserRecordSpec {
	if in == nil {
		return nil
	}
	out := new(MasterUserRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterUserRecordStatus) DeepCopyInto(out *MasterUserRecordStatus) {
	*out = *in
	if in.UserAccounts != nil {
		in, out := &in.UserAccounts, &out.UserAccounts
		*out = make([]UserAccountStatusEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterUserRecordStatus.
func (in *MasterUserRecordStatus) DeepCopy() *MasterUserRecordStatus {
	if in == nil {
		return nil
	}
	out := new(MasterUserRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSTemplateSet) DeepCopyInto(out *NSTemplateSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSTemplateSet.
func (in *NSTemplateSet) DeepCopy() *NSTemplateSet {
	if in == nil {
		return nil
	}
	out := new(NSTemplateSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NSTemplateSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSTemplateSetList) DeepCopyInto(out *NSTemplateSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NSTemplateSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSTemplateSetList.
func (in *NSTemplateSetList) DeepCopy() *NSTemplateSetList {
	if in == nil {
		return nil
	}
	out := new(NSTemplateSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NSTemplateSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSTemplateSetSpec) DeepCopyInto(out *NSTemplateSetSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]Namespace, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSTemplateSetSpec.
func (in *NSTemplateSetSpec) DeepCopy() *NSTemplateSetSpec {
	if in == nil {
		return nil
	}
	out := new(NSTemplateSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSTemplateSetStatus) DeepCopyInto(out *NSTemplateSetStatus) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]NamespaceStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSTemplateSetStatus.
func (in *NSTemplateSetStatus) DeepCopy() *NSTemplateSetStatus {
	if in == nil {
		return nil
	}
	out := new(NSTemplateSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceStatus) DeepCopyInto(out *NamespaceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceStatus.
func (in *NamespaceStatus) DeepCopy() *NamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAccount) DeepCopyInto(out *UserAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAccount.
func (in *UserAccount) DeepCopy() *UserAccount {
	if in == nil {
		return nil
	}
	out := new(UserAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAccountEmbedded) DeepCopyInto(out *UserAccountEmbedded) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAccountEmbedded.
func (in *UserAccountEmbedded) DeepCopy() *UserAccountEmbedded {
	if in == nil {
		return nil
	}
	out := new(UserAccountEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAccountList) DeepCopyInto(out *UserAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAccountList.
func (in *UserAccountList) DeepCopy() *UserAccountList {
	if in == nil {
		return nil
	}
	out := new(UserAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAccountSpec) DeepCopyInto(out *UserAccountSpec) {
	*out = *in
	in.NSTemplateSet.DeepCopyInto(&out.NSTemplateSet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAccountSpec.
func (in *UserAccountSpec) DeepCopy() *UserAccountSpec {
	if in == nil {
		return nil
	}
	out := new(UserAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAccountStatus) DeepCopyInto(out *UserAccountStatus) {
	*out = *in
	in.NSTemplateSetStatus.DeepCopyInto(&out.NSTemplateSetStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAccountStatus.
func (in *UserAccountStatus) DeepCopy() *UserAccountStatus {
	if in == nil {
		return nil
	}
	out := new(UserAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAccountStatusEmbedded) DeepCopyInto(out *UserAccountStatusEmbedded) {
	*out = *in
	in.UserAccountStatus.DeepCopyInto(&out.UserAccountStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAccountStatusEmbedded.
func (in *UserAccountStatusEmbedded) DeepCopy() *UserAccountStatusEmbedded {
	if in == nil {
		return nil
	}
	out := new(UserAccountStatusEmbedded)
	in.DeepCopyInto(out)
	return out
}
