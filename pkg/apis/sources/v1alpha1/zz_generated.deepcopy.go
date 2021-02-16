// +build !ignore_autogenerated

/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VAuthSpec) DeepCopyInto(out *VAuthSpec) {
	*out = *in
	in.Address.DeepCopyInto(&out.Address)
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VAuthSpec.
func (in *VAuthSpec) DeepCopy() *VAuthSpec {
	if in == nil {
		return nil
	}
	out := new(VAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCheckpointSpec) DeepCopyInto(out *VCheckpointSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCheckpointSpec.
func (in *VCheckpointSpec) DeepCopy() *VCheckpointSpec {
	if in == nil {
		return nil
	}
	out := new(VCheckpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereBinding) DeepCopyInto(out *VSphereBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereBinding.
func (in *VSphereBinding) DeepCopy() *VSphereBinding {
	if in == nil {
		return nil
	}
	out := new(VSphereBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereBindingList) DeepCopyInto(out *VSphereBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSphereBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereBindingList.
func (in *VSphereBindingList) DeepCopy() *VSphereBindingList {
	if in == nil {
		return nil
	}
	out := new(VSphereBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereBindingSpec) DeepCopyInto(out *VSphereBindingSpec) {
	*out = *in
	in.BindingSpec.DeepCopyInto(&out.BindingSpec)
	in.VAuthSpec.DeepCopyInto(&out.VAuthSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereBindingSpec.
func (in *VSphereBindingSpec) DeepCopy() *VSphereBindingSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereBindingStatus) DeepCopyInto(out *VSphereBindingStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereBindingStatus.
func (in *VSphereBindingStatus) DeepCopy() *VSphereBindingStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereSource) DeepCopyInto(out *VSphereSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereSource.
func (in *VSphereSource) DeepCopy() *VSphereSource {
	if in == nil {
		return nil
	}
	out := new(VSphereSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereSourceList) DeepCopyInto(out *VSphereSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSphereSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereSourceList.
func (in *VSphereSourceList) DeepCopy() *VSphereSourceList {
	if in == nil {
		return nil
	}
	out := new(VSphereSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereSourceSpec) DeepCopyInto(out *VSphereSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	in.VAuthSpec.DeepCopyInto(&out.VAuthSpec)
	out.CheckpointConfig = in.CheckpointConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereSourceSpec.
func (in *VSphereSourceSpec) DeepCopy() *VSphereSourceSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereSourceStatus) DeepCopyInto(out *VSphereSourceStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereSourceStatus.
func (in *VSphereSourceStatus) DeepCopy() *VSphereSourceStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereSourceStatus)
	in.DeepCopyInto(out)
	return out
}
