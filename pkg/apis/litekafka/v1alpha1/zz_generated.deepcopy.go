// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaCluster) DeepCopyInto(out *KafkaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaCluster.
func (in *KafkaCluster) DeepCopy() *KafkaCluster {
	if in == nil {
		return nil
	}
	out := new(KafkaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterList) DeepCopyInto(out *KafkaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterList.
func (in *KafkaClusterList) DeepCopy() *KafkaClusterList {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterSpec) DeepCopyInto(out *KafkaClusterSpec) {
	*out = *in
	out.ContainerPort = in.ContainerPort
	out.ServicePort = in.ServicePort
	out.Options = in.Options
	out.Zookeeper = in.Zookeeper
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterSpec.
func (in *KafkaClusterSpec) DeepCopy() *KafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterStatus) DeepCopyInto(out *KafkaClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterStatus.
func (in *KafkaClusterStatus) DeepCopy() *KafkaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaOptions) DeepCopyInto(out *KafkaOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaOptions.
func (in *KafkaOptions) DeepCopy() *KafkaOptions {
	if in == nil {
		return nil
	}
	out := new(KafkaOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperSpec) DeepCopyInto(out *ZookeeperSpec) {
	*out = *in
	out.Port = in.Port
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperSpec.
func (in *ZookeeperSpec) DeepCopy() *ZookeeperSpec {
	if in == nil {
		return nil
	}
	out := new(ZookeeperSpec)
	in.DeepCopyInto(out)
	return out
}
