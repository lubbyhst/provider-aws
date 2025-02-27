//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSecretParameters) DeepCopyInto(out *CustomSecretParameters) {
	*out = *in
	if in.KMSKeyIDRef != nil {
		in, out := &in.KMSKeyIDRef, &out.KMSKeyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyIDSelector != nil {
		in, out := &in.KMSKeyIDSelector, &out.KMSKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StringSecretRef != nil {
		in, out := &in.StringSecretRef, &out.StringSecretRef
		*out = new(SecretReference)
		(*in).DeepCopyInto(*out)
	}
	if in.BinarySecretRef != nil {
		in, out := &in.BinarySecretRef, &out.BinarySecretRef
		*out = new(SecretReference)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceDeleteWithoutRecovery != nil {
		in, out := &in.ForceDeleteWithoutRecovery, &out.ForceDeleteWithoutRecovery
		*out = new(bool)
		**out = **in
	}
	if in.RecoveryWindowInDays != nil {
		in, out := &in.RecoveryWindowInDays, &out.RecoveryWindowInDays
		*out = new(int64)
		**out = **in
	}
	if in.ResourcePolicy != nil {
		in, out := &in.ResourcePolicy, &out.ResourcePolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSecretParameters.
func (in *CustomSecretParameters) DeepCopy() *CustomSecretParameters {
	if in == nil {
		return nil
	}
	out := new(CustomSecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaRegionType) DeepCopyInto(out *ReplicaRegionType) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaRegionType.
func (in *ReplicaRegionType) DeepCopy() *ReplicaRegionType {
	if in == nil {
		return nil
	}
	out := new(ReplicaRegionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationStatusType) DeepCopyInto(out *ReplicationStatusType) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = (*in).DeepCopy()
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationStatusType.
func (in *ReplicationStatusType) DeepCopy() *ReplicationStatusType {
	if in == nil {
		return nil
	}
	out := new(ReplicationStatusType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationRulesType) DeepCopyInto(out *RotationRulesType) {
	*out = *in
	if in.AutomaticallyAfterDays != nil {
		in, out := &in.AutomaticallyAfterDays, &out.AutomaticallyAfterDays
		*out = new(int64)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.ScheduleExpression != nil {
		in, out := &in.ScheduleExpression, &out.ScheduleExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationRulesType.
func (in *RotationRulesType) DeepCopy() *RotationRulesType {
	if in == nil {
		return nil
	}
	out := new(RotationRulesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Secret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretList) DeepCopyInto(out *SecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretList.
func (in *SecretList) DeepCopy() *SecretList {
	if in == nil {
		return nil
	}
	out := new(SecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretListEntry) DeepCopyInto(out *SecretListEntry) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = (*in).DeepCopy()
	}
	if in.DeletedDate != nil {
		in, out := &in.DeletedDate, &out.DeletedDate
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = (*in).DeepCopy()
	}
	if in.LastChangedDate != nil {
		in, out := &in.LastChangedDate, &out.LastChangedDate
		*out = (*in).DeepCopy()
	}
	if in.LastRotatedDate != nil {
		in, out := &in.LastRotatedDate, &out.LastRotatedDate
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwningService != nil {
		in, out := &in.OwningService, &out.OwningService
		*out = new(string)
		**out = **in
	}
	if in.PrimaryRegion != nil {
		in, out := &in.PrimaryRegion, &out.PrimaryRegion
		*out = new(string)
		**out = **in
	}
	if in.RotationEnabled != nil {
		in, out := &in.RotationEnabled, &out.RotationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RotationLambdaARN != nil {
		in, out := &in.RotationLambdaARN, &out.RotationLambdaARN
		*out = new(string)
		**out = **in
	}
	if in.RotationRules != nil {
		in, out := &in.RotationRules, &out.RotationRules
		*out = new(RotationRulesType)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretVersionsToStages != nil {
		in, out := &in.SecretVersionsToStages, &out.SecretVersionsToStages
		*out = make(map[string][]*string, len(*in))
		for key, val := range *in {
			var outVal []*string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]*string, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(string)
						**out = **in
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretListEntry.
func (in *SecretListEntry) DeepCopy() *SecretListEntry {
	if in == nil {
		return nil
	}
	out := new(SecretListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretObservation) DeepCopyInto(out *SecretObservation) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.ReplicationStatus != nil {
		in, out := &in.ReplicationStatus, &out.ReplicationStatus
		*out = make([]*ReplicationStatusType, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicationStatusType)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretObservation.
func (in *SecretObservation) DeepCopy() *SecretObservation {
	if in == nil {
		return nil
	}
	out := new(SecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretParameters) DeepCopyInto(out *SecretParameters) {
	*out = *in
	if in.AddReplicaRegions != nil {
		in, out := &in.AddReplicaRegions, &out.AddReplicaRegions
		*out = make([]*ReplicaRegionType, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicaRegionType)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceOverwriteReplicaSecret != nil {
		in, out := &in.ForceOverwriteReplicaSecret, &out.ForceOverwriteReplicaSecret
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomSecretParameters.DeepCopyInto(&out.CustomSecretParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretParameters.
func (in *SecretParameters) DeepCopy() *SecretParameters {
	if in == nil {
		return nil
	}
	out := new(SecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatus) DeepCopyInto(out *SecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatus.
func (in *SecretStatus) DeepCopy() *SecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionsListEntry) DeepCopyInto(out *SecretVersionsListEntry) {
	*out = *in
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = (*in).DeepCopy()
	}
	if in.VersionStages != nil {
		in, out := &in.VersionStages, &out.VersionStages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionsListEntry.
func (in *SecretVersionsListEntry) DeepCopy() *SecretVersionsListEntry {
	if in == nil {
		return nil
	}
	out := new(SecretVersionsListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
