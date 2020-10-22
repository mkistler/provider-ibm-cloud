// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanHistoryItem) DeepCopyInto(out *PlanHistoryItem) {
	*out = *in
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanHistoryItem.
func (in *PlanHistoryItem) DeepCopy() *PlanHistoryItem {
	if in == nil {
		return nil
	}
	out := new(PlanHistoryItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstance) DeepCopyInto(out *ResourceInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstance.
func (in *ResourceInstance) DeepCopy() *ResourceInstance {
	if in == nil {
		return nil
	}
	out := new(ResourceInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstanceList) DeepCopyInto(out *ResourceInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstanceList.
func (in *ResourceInstanceList) DeepCopy() *ResourceInstanceList {
	if in == nil {
		return nil
	}
	out := new(ResourceInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstanceObservation) DeepCopyInto(out *ResourceInstanceObservation) {
	*out = *in
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.PlanHistory != nil {
		in, out := &in.PlanHistory, &out.PlanHistory
		*out = make([]PlanHistoryItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.DeletedAt != nil {
		in, out := &in.DeletedAt, &out.DeletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstanceObservation.
func (in *ResourceInstanceObservation) DeepCopy() *ResourceInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstanceParameters) DeepCopyInto(out *ResourceInstanceParameters) {
	*out = *in
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowCleanup != nil {
		in, out := &in.AllowCleanup, &out.AllowCleanup
		*out = new(bool)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.EntityLock != nil {
		in, out := &in.EntityLock, &out.EntityLock
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstanceParameters.
func (in *ResourceInstanceParameters) DeepCopy() *ResourceInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstanceSpec) DeepCopyInto(out *ResourceInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstanceSpec.
func (in *ResourceInstanceSpec) DeepCopy() *ResourceInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstanceStatus) DeepCopyInto(out *ResourceInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstanceStatus.
func (in *ResourceInstanceStatus) DeepCopy() *ResourceInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
