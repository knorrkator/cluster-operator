// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImages) DeepCopyInto(out *ContainerImages) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerImages.
func (in *ContainerImages) DeepCopy() *ContainerImages {
	if in == nil {
		return nil
	}
	out := new(ContainerImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealth) DeepCopyInto(out *NodeHealth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealth.
func (in *NodeHealth) DeepCopy() *NodeHealth {
	if in == nil {
		return nil
	}
	out := new(NodeHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSCSI) DeepCopyInto(out *StorageOSCSI) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSCSI.
func (in *StorageOSCSI) DeepCopy() *StorageOSCSI {
	if in == nil {
		return nil
	}
	out := new(StorageOSCSI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSCluster) DeepCopyInto(out *StorageOSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSCluster.
func (in *StorageOSCluster) DeepCopy() *StorageOSCluster {
	if in == nil {
		return nil
	}
	out := new(StorageOSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageOSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSClusterList) DeepCopyInto(out *StorageOSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageOSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSClusterList.
func (in *StorageOSClusterList) DeepCopy() *StorageOSClusterList {
	if in == nil {
		return nil
	}
	out := new(StorageOSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageOSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSIngress) DeepCopyInto(out *StorageOSIngress) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSIngress.
func (in *StorageOSIngress) DeepCopy() *StorageOSIngress {
	if in == nil {
		return nil
	}
	out := new(StorageOSIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSService) DeepCopyInto(out *StorageOSService) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSService.
func (in *StorageOSService) DeepCopy() *StorageOSService {
	if in == nil {
		return nil
	}
	out := new(StorageOSService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSServiceStatus) DeepCopyInto(out *StorageOSServiceStatus) {
	*out = *in
	if in.NodeHealthStatus != nil {
		in, out := &in.NodeHealthStatus, &out.NodeHealthStatus
		*out = make(map[string]NodeHealth, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSServiceStatus.
func (in *StorageOSServiceStatus) DeepCopy() *StorageOSServiceStatus {
	if in == nil {
		return nil
	}
	out := new(StorageOSServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageOSSpec) DeepCopyInto(out *StorageOSSpec) {
	*out = *in
	out.CSI = in.CSI
	in.Service.DeepCopyInto(&out.Service)
	in.Ingress.DeepCopyInto(&out.Ingress)
	out.Images = in.Images
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageOSSpec.
func (in *StorageOSSpec) DeepCopy() *StorageOSSpec {
	if in == nil {
		return nil
	}
	out := new(StorageOSSpec)
	in.DeepCopyInto(out)
	return out
}
