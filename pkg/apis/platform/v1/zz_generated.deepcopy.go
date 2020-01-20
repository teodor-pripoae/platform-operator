// +build !ignore_autogenerated

/*
Copyright 2020 The Kubernetes Authors.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuota) DeepCopyInto(out *ClusterResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuota.
func (in *ClusterResourceQuota) DeepCopy() *ClusterResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaList) DeepCopyInto(out *ClusterResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaList.
func (in *ClusterResourceQuotaList) DeepCopy() *ClusterResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaSpec) DeepCopyInto(out *ClusterResourceQuotaSpec) {
	*out = *in
	in.Quota.DeepCopyInto(&out.Quota)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaSpec.
func (in *ClusterResourceQuotaSpec) DeepCopy() *ClusterResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaStatus) DeepCopyInto(out *ClusterResourceQuotaStatus) {
	*out = *in
	in.Total.DeepCopyInto(&out.Total)
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make(ResourceQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaStatus.
func (in *ClusterResourceQuotaStatus) DeepCopy() *ClusterResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaStatusByNamespace) DeepCopyInto(out *ResourceQuotaStatusByNamespace) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaStatusByNamespace.
func (in *ResourceQuotaStatusByNamespace) DeepCopy() *ResourceQuotaStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaStatusByNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ResourceQuotasStatusByNamespace) DeepCopyInto(out *ResourceQuotasStatusByNamespace) {
	{
		in := &in
		*out = make(ResourceQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotasStatusByNamespace.
func (in ResourceQuotasStatusByNamespace) DeepCopy() ResourceQuotasStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotasStatusByNamespace)
	in.DeepCopyInto(out)
	return *out
}
