//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbfsInitParameters) DeepCopyInto(out *AbfsInitParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretKey != nil {
		in, out := &in.ClientSecretKey, &out.ClientSecretKey
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretScope != nil {
		in, out := &in.ClientSecretScope, &out.ClientSecretScope
		*out = new(string)
		**out = **in
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.InitializeFileSystem != nil {
		in, out := &in.InitializeFileSystem, &out.InitializeFileSystem
		*out = new(bool)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbfsInitParameters.
func (in *AbfsInitParameters) DeepCopy() *AbfsInitParameters {
	if in == nil {
		return nil
	}
	out := new(AbfsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbfsObservation) DeepCopyInto(out *AbfsObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretKey != nil {
		in, out := &in.ClientSecretKey, &out.ClientSecretKey
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretScope != nil {
		in, out := &in.ClientSecretScope, &out.ClientSecretScope
		*out = new(string)
		**out = **in
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.InitializeFileSystem != nil {
		in, out := &in.InitializeFileSystem, &out.InitializeFileSystem
		*out = new(bool)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbfsObservation.
func (in *AbfsObservation) DeepCopy() *AbfsObservation {
	if in == nil {
		return nil
	}
	out := new(AbfsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbfsParameters) DeepCopyInto(out *AbfsParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretKey != nil {
		in, out := &in.ClientSecretKey, &out.ClientSecretKey
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretScope != nil {
		in, out := &in.ClientSecretScope, &out.ClientSecretScope
		*out = new(string)
		**out = **in
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.InitializeFileSystem != nil {
		in, out := &in.InitializeFileSystem, &out.InitializeFileSystem
		*out = new(bool)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbfsParameters.
func (in *AbfsParameters) DeepCopy() *AbfsParameters {
	if in == nil {
		return nil
	}
	out := new(AbfsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdlInitParameters) DeepCopyInto(out *AdlInitParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretKey != nil {
		in, out := &in.ClientSecretKey, &out.ClientSecretKey
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretScope != nil {
		in, out := &in.ClientSecretScope, &out.ClientSecretScope
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.SparkConfPrefix != nil {
		in, out := &in.SparkConfPrefix, &out.SparkConfPrefix
		*out = new(string)
		**out = **in
	}
	if in.StorageResourceName != nil {
		in, out := &in.StorageResourceName, &out.StorageResourceName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdlInitParameters.
func (in *AdlInitParameters) DeepCopy() *AdlInitParameters {
	if in == nil {
		return nil
	}
	out := new(AdlInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdlObservation) DeepCopyInto(out *AdlObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretKey != nil {
		in, out := &in.ClientSecretKey, &out.ClientSecretKey
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretScope != nil {
		in, out := &in.ClientSecretScope, &out.ClientSecretScope
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.SparkConfPrefix != nil {
		in, out := &in.SparkConfPrefix, &out.SparkConfPrefix
		*out = new(string)
		**out = **in
	}
	if in.StorageResourceName != nil {
		in, out := &in.StorageResourceName, &out.StorageResourceName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdlObservation.
func (in *AdlObservation) DeepCopy() *AdlObservation {
	if in == nil {
		return nil
	}
	out := new(AdlObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdlParameters) DeepCopyInto(out *AdlParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretKey != nil {
		in, out := &in.ClientSecretKey, &out.ClientSecretKey
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretScope != nil {
		in, out := &in.ClientSecretScope, &out.ClientSecretScope
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.SparkConfPrefix != nil {
		in, out := &in.SparkConfPrefix, &out.SparkConfPrefix
		*out = new(string)
		**out = **in
	}
	if in.StorageResourceName != nil {
		in, out := &in.StorageResourceName, &out.StorageResourceName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdlParameters.
func (in *AdlParameters) DeepCopy() *AdlParameters {
	if in == nil {
		return nil
	}
	out := new(AdlParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFile) DeepCopyInto(out *DbfsFile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFile.
func (in *DbfsFile) DeepCopy() *DbfsFile {
	if in == nil {
		return nil
	}
	out := new(DbfsFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbfsFile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFileInitParameters) DeepCopyInto(out *DbfsFileInitParameters) {
	*out = *in
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Md5 != nil {
		in, out := &in.Md5, &out.Md5
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFileInitParameters.
func (in *DbfsFileInitParameters) DeepCopy() *DbfsFileInitParameters {
	if in == nil {
		return nil
	}
	out := new(DbfsFileInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFileList) DeepCopyInto(out *DbfsFileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DbfsFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFileList.
func (in *DbfsFileList) DeepCopy() *DbfsFileList {
	if in == nil {
		return nil
	}
	out := new(DbfsFileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbfsFileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFileObservation) DeepCopyInto(out *DbfsFileObservation) {
	*out = *in
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.DbfsPath != nil {
		in, out := &in.DbfsPath, &out.DbfsPath
		*out = new(string)
		**out = **in
	}
	if in.FileSize != nil {
		in, out := &in.FileSize, &out.FileSize
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Md5 != nil {
		in, out := &in.Md5, &out.Md5
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFileObservation.
func (in *DbfsFileObservation) DeepCopy() *DbfsFileObservation {
	if in == nil {
		return nil
	}
	out := new(DbfsFileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFileParameters) DeepCopyInto(out *DbfsFileParameters) {
	*out = *in
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Md5 != nil {
		in, out := &in.Md5, &out.Md5
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFileParameters.
func (in *DbfsFileParameters) DeepCopy() *DbfsFileParameters {
	if in == nil {
		return nil
	}
	out := new(DbfsFileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFileSpec) DeepCopyInto(out *DbfsFileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFileSpec.
func (in *DbfsFileSpec) DeepCopy() *DbfsFileSpec {
	if in == nil {
		return nil
	}
	out := new(DbfsFileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsFileStatus) DeepCopyInto(out *DbfsFileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsFileStatus.
func (in *DbfsFileStatus) DeepCopy() *DbfsFileStatus {
	if in == nil {
		return nil
	}
	out := new(DbfsFileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *File) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileInitParameters) DeepCopyInto(out *FileInitParameters) {
	*out = *in
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Md5 != nil {
		in, out := &in.Md5, &out.Md5
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.RemoteFileModified != nil {
		in, out := &in.RemoteFileModified, &out.RemoteFileModified
		*out = new(bool)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileInitParameters.
func (in *FileInitParameters) DeepCopy() *FileInitParameters {
	if in == nil {
		return nil
	}
	out := new(FileInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileList) DeepCopyInto(out *FileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileList.
func (in *FileList) DeepCopy() *FileList {
	if in == nil {
		return nil
	}
	out := new(FileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileObservation) DeepCopyInto(out *FileObservation) {
	*out = *in
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.FileSize != nil {
		in, out := &in.FileSize, &out.FileSize
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Md5 != nil {
		in, out := &in.Md5, &out.Md5
		*out = new(string)
		**out = **in
	}
	if in.ModificationTime != nil {
		in, out := &in.ModificationTime, &out.ModificationTime
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.RemoteFileModified != nil {
		in, out := &in.RemoteFileModified, &out.RemoteFileModified
		*out = new(bool)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileObservation.
func (in *FileObservation) DeepCopy() *FileObservation {
	if in == nil {
		return nil
	}
	out := new(FileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileParameters) DeepCopyInto(out *FileParameters) {
	*out = *in
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Md5 != nil {
		in, out := &in.Md5, &out.Md5
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.RemoteFileModified != nil {
		in, out := &in.RemoteFileModified, &out.RemoteFileModified
		*out = new(bool)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileParameters.
func (in *FileParameters) DeepCopy() *FileParameters {
	if in == nil {
		return nil
	}
	out := new(FileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStatus) DeepCopyInto(out *FileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStatus.
func (in *FileStatus) DeepCopy() *FileStatus {
	if in == nil {
		return nil
	}
	out := new(FileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsInitParameters) DeepCopyInto(out *GsInitParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsInitParameters.
func (in *GsInitParameters) DeepCopy() *GsInitParameters {
	if in == nil {
		return nil
	}
	out := new(GsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsObservation) DeepCopyInto(out *GsObservation) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsObservation.
func (in *GsObservation) DeepCopy() *GsObservation {
	if in == nil {
		return nil
	}
	out := new(GsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsParameters) DeepCopyInto(out *GsParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsParameters.
func (in *GsParameters) DeepCopy() *GsParameters {
	if in == nil {
		return nil
	}
	out := new(GsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mount.
func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountInitParameters) DeepCopyInto(out *MountInitParameters) {
	*out = *in
	if in.Abfs != nil {
		in, out := &in.Abfs, &out.Abfs
		*out = make([]AbfsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Adl != nil {
		in, out := &in.Adl, &out.Adl
		*out = make([]AdlInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.ExtraConfigs != nil {
		in, out := &in.ExtraConfigs, &out.ExtraConfigs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Gs != nil {
		in, out := &in.Gs, &out.Gs
		*out = make([]GsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]S3InitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	if in.Wasb != nil {
		in, out := &in.Wasb, &out.Wasb
		*out = make([]WasbInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountInitParameters.
func (in *MountInitParameters) DeepCopy() *MountInitParameters {
	if in == nil {
		return nil
	}
	out := new(MountInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountList) DeepCopyInto(out *MountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountList.
func (in *MountList) DeepCopy() *MountList {
	if in == nil {
		return nil
	}
	out := new(MountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountObservation) DeepCopyInto(out *MountObservation) {
	*out = *in
	if in.Abfs != nil {
		in, out := &in.Abfs, &out.Abfs
		*out = make([]AbfsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Adl != nil {
		in, out := &in.Adl, &out.Adl
		*out = make([]AdlObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.ExtraConfigs != nil {
		in, out := &in.ExtraConfigs, &out.ExtraConfigs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Gs != nil {
		in, out := &in.Gs, &out.Gs
		*out = make([]GsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]S3Observation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	if in.Wasb != nil {
		in, out := &in.Wasb, &out.Wasb
		*out = make([]WasbObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountObservation.
func (in *MountObservation) DeepCopy() *MountObservation {
	if in == nil {
		return nil
	}
	out := new(MountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountParameters) DeepCopyInto(out *MountParameters) {
	*out = *in
	if in.Abfs != nil {
		in, out := &in.Abfs, &out.Abfs
		*out = make([]AbfsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Adl != nil {
		in, out := &in.Adl, &out.Adl
		*out = make([]AdlParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.ExtraConfigs != nil {
		in, out := &in.ExtraConfigs, &out.ExtraConfigs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Gs != nil {
		in, out := &in.Gs, &out.Gs
		*out = make([]GsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]S3Parameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	if in.Wasb != nil {
		in, out := &in.Wasb, &out.Wasb
		*out = make([]WasbParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountParameters.
func (in *MountParameters) DeepCopy() *MountParameters {
	if in == nil {
		return nil
	}
	out := new(MountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountSpec) DeepCopyInto(out *MountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountSpec.
func (in *MountSpec) DeepCopy() *MountSpec {
	if in == nil {
		return nil
	}
	out := new(MountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountStatus) DeepCopyInto(out *MountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountStatus.
func (in *MountStatus) DeepCopy() *MountStatus {
	if in == nil {
		return nil
	}
	out := new(MountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3InitParameters) DeepCopyInto(out *S3InitParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.InstanceProfile != nil {
		in, out := &in.InstanceProfile, &out.InstanceProfile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3InitParameters.
func (in *S3InitParameters) DeepCopy() *S3InitParameters {
	if in == nil {
		return nil
	}
	out := new(S3InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Observation) DeepCopyInto(out *S3Observation) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.InstanceProfile != nil {
		in, out := &in.InstanceProfile, &out.InstanceProfile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Observation.
func (in *S3Observation) DeepCopy() *S3Observation {
	if in == nil {
		return nil
	}
	out := new(S3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Parameters) DeepCopyInto(out *S3Parameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.InstanceProfile != nil {
		in, out := &in.InstanceProfile, &out.InstanceProfile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Parameters.
func (in *S3Parameters) DeepCopy() *S3Parameters {
	if in == nil {
		return nil
	}
	out := new(S3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasbInitParameters) DeepCopyInto(out *WasbInitParameters) {
	*out = *in
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretKey != nil {
		in, out := &in.TokenSecretKey, &out.TokenSecretKey
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretScope != nil {
		in, out := &in.TokenSecretScope, &out.TokenSecretScope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasbInitParameters.
func (in *WasbInitParameters) DeepCopy() *WasbInitParameters {
	if in == nil {
		return nil
	}
	out := new(WasbInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasbObservation) DeepCopyInto(out *WasbObservation) {
	*out = *in
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretKey != nil {
		in, out := &in.TokenSecretKey, &out.TokenSecretKey
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretScope != nil {
		in, out := &in.TokenSecretScope, &out.TokenSecretScope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasbObservation.
func (in *WasbObservation) DeepCopy() *WasbObservation {
	if in == nil {
		return nil
	}
	out := new(WasbObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasbParameters) DeepCopyInto(out *WasbParameters) {
	*out = *in
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretKey != nil {
		in, out := &in.TokenSecretKey, &out.TokenSecretKey
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretScope != nil {
		in, out := &in.TokenSecretScope, &out.TokenSecretScope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasbParameters.
func (in *WasbParameters) DeepCopy() *WasbParameters {
	if in == nil {
		return nil
	}
	out := new(WasbParameters)
	in.DeepCopyInto(out)
	return out
}
