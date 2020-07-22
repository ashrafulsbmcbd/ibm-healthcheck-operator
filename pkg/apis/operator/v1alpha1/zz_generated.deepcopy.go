// +build !ignore_autogenerated

//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthService) DeepCopyInto(out *HealthService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthService.
func (in *HealthService) DeepCopy() *HealthService {
	if in == nil {
		return nil
	}
	out := new(HealthService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthServiceList) DeepCopyInto(out *HealthServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthServiceList.
func (in *HealthServiceList) DeepCopy() *HealthServiceList {
	if in == nil {
		return nil
	}
	out := new(HealthServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthServiceSpec) DeepCopyInto(out *HealthServiceSpec) {
	*out = *in
	in.Memcached.DeepCopyInto(&out.Memcached)
	in.HealthService.DeepCopyInto(&out.HealthService)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthServiceSpec.
func (in *HealthServiceSpec) DeepCopy() *HealthServiceSpec {
	if in == nil {
		return nil
	}
	out := new(HealthServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthServiceSpecHealthService) DeepCopyInto(out *HealthServiceSpecHealthService) {
	*out = *in
	out.Image = in.Image
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthServiceSpecHealthService.
func (in *HealthServiceSpecHealthService) DeepCopy() *HealthServiceSpecHealthService {
	if in == nil {
		return nil
	}
	out := new(HealthServiceSpecHealthService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthServiceSpecMemcached) DeepCopyInto(out *HealthServiceSpecMemcached) {
	*out = *in
	out.Image = in.Image
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthServiceSpecMemcached.
func (in *HealthServiceSpecMemcached) DeepCopy() *HealthServiceSpecMemcached {
	if in == nil {
		return nil
	}
	out := new(HealthServiceSpecMemcached)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthServiceStatus) DeepCopyInto(out *HealthServiceStatus) {
	*out = *in
	if in.MemcachedNodes != nil {
		in, out := &in.MemcachedNodes, &out.MemcachedNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HealthCheckNodes != nil {
		in, out := &in.HealthCheckNodes, &out.HealthCheckNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthServiceStatus.
func (in *HealthServiceStatus) DeepCopy() *HealthServiceStatus {
	if in == nil {
		return nil
	}
	out := new(HealthServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGather) DeepCopyInto(out *MustGather) {
	*out = *in
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGather.
func (in *MustGather) DeepCopy() *MustGather {
	if in == nil {
		return nil
	}
	out := new(MustGather)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherJob) DeepCopyInto(out *MustGatherJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherJob.
func (in *MustGatherJob) DeepCopy() *MustGatherJob {
	if in == nil {
		return nil
	}
	out := new(MustGatherJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MustGatherJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherJobList) DeepCopyInto(out *MustGatherJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MustGatherJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherJobList.
func (in *MustGatherJobList) DeepCopy() *MustGatherJobList {
	if in == nil {
		return nil
	}
	out := new(MustGatherJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MustGatherJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherJobSpec) DeepCopyInto(out *MustGatherJobSpec) {
	*out = *in
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherJobSpec.
func (in *MustGatherJobSpec) DeepCopy() *MustGatherJobSpec {
	if in == nil {
		return nil
	}
	out := new(MustGatherJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherJobStatus) DeepCopyInto(out *MustGatherJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherJobStatus.
func (in *MustGatherJobStatus) DeepCopy() *MustGatherJobStatus {
	if in == nil {
		return nil
	}
	out := new(MustGatherJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherService) DeepCopyInto(out *MustGatherService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherService.
func (in *MustGatherService) DeepCopy() *MustGatherService {
	if in == nil {
		return nil
	}
	out := new(MustGatherService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MustGatherService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherServiceList) DeepCopyInto(out *MustGatherServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MustGatherService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherServiceList.
func (in *MustGatherServiceList) DeepCopy() *MustGatherServiceList {
	if in == nil {
		return nil
	}
	out := new(MustGatherServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MustGatherServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherServiceSpec) DeepCopyInto(out *MustGatherServiceSpec) {
	*out = *in
	in.MustGather.DeepCopyInto(&out.MustGather)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherServiceSpec.
func (in *MustGatherServiceSpec) DeepCopy() *MustGatherServiceSpec {
	if in == nil {
		return nil
	}
	out := new(MustGatherServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MustGatherServiceStatus) DeepCopyInto(out *MustGatherServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MustGatherServiceStatus.
func (in *MustGatherServiceStatus) DeepCopy() *MustGatherServiceStatus {
	if in == nil {
		return nil
	}
	out := new(MustGatherServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}
