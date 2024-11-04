//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimAutoscaler) DeepCopyInto(out *PersistentVolumeClaimAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimAutoscaler.
func (in *PersistentVolumeClaimAutoscaler) DeepCopy() *PersistentVolumeClaimAutoscaler {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PersistentVolumeClaimAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimAutoscalerList) DeepCopyInto(out *PersistentVolumeClaimAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PersistentVolumeClaimAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimAutoscalerList.
func (in *PersistentVolumeClaimAutoscalerList) DeepCopy() *PersistentVolumeClaimAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PersistentVolumeClaimAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimAutoscalerSpec) DeepCopyInto(out *PersistentVolumeClaimAutoscalerSpec) {
	*out = *in
	out.MaxCapacity = in.MaxCapacity.DeepCopy()
	out.ScaleTargetRef = in.ScaleTargetRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimAutoscalerSpec.
func (in *PersistentVolumeClaimAutoscalerSpec) DeepCopy() *PersistentVolumeClaimAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimAutoscalerStatus) DeepCopyInto(out *PersistentVolumeClaimAutoscalerStatus) {
	*out = *in
	in.LastCheck.DeepCopyInto(&out.LastCheck)
	in.NextCheck.DeepCopyInto(&out.NextCheck)
	out.PrevSize = in.PrevSize.DeepCopy()
	out.NewSize = in.NewSize.DeepCopy()
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimAutoscalerStatus.
func (in *PersistentVolumeClaimAutoscalerStatus) DeepCopy() *PersistentVolumeClaimAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}
