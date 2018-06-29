// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidSDK) DeepCopyInto(out *AndroidSDK) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidSDK.
func (in *AndroidSDK) DeepCopy() *AndroidSDK {
	if in == nil {
		return nil
	}
	out := new(AndroidSDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AndroidSDK) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidSDKList) DeepCopyInto(out *AndroidSDKList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AndroidSDK, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidSDKList.
func (in *AndroidSDKList) DeepCopy() *AndroidSDKList {
	if in == nil {
		return nil
	}
	out := new(AndroidSDKList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AndroidSDKList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidSDKSpec) DeepCopyInto(out *AndroidSDKSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidSDKSpec.
func (in *AndroidSDKSpec) DeepCopy() *AndroidSDKSpec {
	if in == nil {
		return nil
	}
	out := new(AndroidSDKSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidSDKStatus) DeepCopyInto(out *AndroidSDKStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidSDKStatus.
func (in *AndroidSDKStatus) DeepCopy() *AndroidSDKStatus {
	if in == nil {
		return nil
	}
	out := new(AndroidSDKStatus)
	in.DeepCopyInto(out)
	return out
}