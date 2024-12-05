//go:build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1beta1"
	cluster_apiapiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetwork) DeepCopyInto(out *ClusterNetwork) {
	*out = *in
	if in.PrivateCluster != nil {
		in, out := &in.PrivateCluster, &out.PrivateCluster
		*out = new(PrivateCluster)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetwork.
func (in *ClusterNetwork) DeepCopy() *ClusterNetwork {
	if in == nil {
		return nil
	}
	out := new(ClusterNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedCluster) DeepCopyInto(out *GCPManagedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedCluster.
func (in *GCPManagedCluster) DeepCopy() *GCPManagedCluster {
	if in == nil {
		return nil
	}
	out := new(GCPManagedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPManagedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedClusterList) DeepCopyInto(out *GCPManagedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCPManagedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedClusterList.
func (in *GCPManagedClusterList) DeepCopy() *GCPManagedClusterList {
	if in == nil {
		return nil
	}
	out := new(GCPManagedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPManagedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedClusterSpec) DeepCopyInto(out *GCPManagedClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	in.Network.DeepCopyInto(&out.Network)
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(apiv1beta1.Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceManagerTags != nil {
		in, out := &in.ResourceManagerTags, &out.ResourceManagerTags
		*out = make(apiv1beta1.ResourceManagerTags, len(*in))
		copy(*out, *in)
	}
	if in.CredentialsRef != nil {
		in, out := &in.CredentialsRef, &out.CredentialsRef
		*out = new(apiv1beta1.ObjectReference)
		**out = **in
	}
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedClusterSpec.
func (in *GCPManagedClusterSpec) DeepCopy() *GCPManagedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(GCPManagedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedClusterStatus) DeepCopyInto(out *GCPManagedClusterStatus) {
	*out = *in
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(cluster_apiapiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Network.DeepCopyInto(&out.Network)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedClusterStatus.
func (in *GCPManagedClusterStatus) DeepCopy() *GCPManagedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GCPManagedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedControlPlane) DeepCopyInto(out *GCPManagedControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedControlPlane.
func (in *GCPManagedControlPlane) DeepCopy() *GCPManagedControlPlane {
	if in == nil {
		return nil
	}
	out := new(GCPManagedControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPManagedControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedControlPlaneList) DeepCopyInto(out *GCPManagedControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCPManagedControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedControlPlaneList.
func (in *GCPManagedControlPlaneList) DeepCopy() *GCPManagedControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(GCPManagedControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPManagedControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedControlPlaneSpec) DeepCopyInto(out *GCPManagedControlPlaneSpec) {
	*out = *in
	if in.ClusterNetwork != nil {
		in, out := &in.ClusterNetwork, &out.ClusterNetwork
		*out = new(ClusterNetwork)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(LoggingConfig)
		**out = **in
	}
	if in.MonitoringConfig != nil {
		in, out := &in.MonitoringConfig, &out.MonitoringConfig
		*out = new(MonitoringConfig)
		**out = **in
	}
	if in.ClusterIpv4Cidr != nil {
		in, out := &in.ClusterIpv4Cidr, &out.ClusterIpv4Cidr
		*out = new(string)
		**out = **in
	}
	if in.IPAllocationPolicy != nil {
		in, out := &in.IPAllocationPolicy, &out.IPAllocationPolicy
		*out = new(IPAllocationPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ReleaseChannel != nil {
		in, out := &in.ReleaseChannel, &out.ReleaseChannel
		*out = new(ReleaseChannel)
		**out = **in
	}
	if in.ControlPlaneVersion != nil {
		in, out := &in.ControlPlaneVersion, &out.ControlPlaneVersion
		*out = new(string)
		**out = **in
	}
	out.Endpoint = in.Endpoint
	if in.MasterAuthorizedNetworksConfig != nil {
		in, out := &in.MasterAuthorizedNetworksConfig, &out.MasterAuthorizedNetworksConfig
		*out = new(MasterAuthorizedNetworksConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkPolicy != nil {
		in, out := &in.NetworkPolicy, &out.NetworkPolicy
		*out = new(NetworkPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedControlPlaneSpec.
func (in *GCPManagedControlPlaneSpec) DeepCopy() *GCPManagedControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(GCPManagedControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedControlPlaneStatus) DeepCopyInto(out *GCPManagedControlPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedControlPlaneStatus.
func (in *GCPManagedControlPlaneStatus) DeepCopy() *GCPManagedControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(GCPManagedControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedMachinePool) DeepCopyInto(out *GCPManagedMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedMachinePool.
func (in *GCPManagedMachinePool) DeepCopy() *GCPManagedMachinePool {
	if in == nil {
		return nil
	}
	out := new(GCPManagedMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPManagedMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedMachinePoolList) DeepCopyInto(out *GCPManagedMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCPManagedMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedMachinePoolList.
func (in *GCPManagedMachinePoolList) DeepCopy() *GCPManagedMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(GCPManagedMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPManagedMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedMachinePoolSpec) DeepCopyInto(out *GCPManagedMachinePoolSpec) {
	*out = *in
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int32)
		**out = **in
	}
	if in.LocalSsdCount != nil {
		in, out := &in.LocalSsdCount, &out.LocalSsdCount
		*out = new(int32)
		**out = **in
	}
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(NodePoolAutoScaling)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeLocations != nil {
		in, out := &in.NodeLocations, &out.NodeLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImageType != nil {
		in, out := &in.ImageType, &out.ImageType
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(DiskType)
		**out = **in
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int64)
		**out = **in
	}
	if in.MaxPodsPerNode != nil {
		in, out := &in.MaxPodsPerNode, &out.MaxPodsPerNode
		*out = new(int64)
		**out = **in
	}
	in.NodeNetwork.DeepCopyInto(&out.NodeNetwork)
	in.NodeSecurity.DeepCopyInto(&out.NodeSecurity)
	if in.KubernetesLabels != nil {
		in, out := &in.KubernetesLabels, &out.KubernetesLabels
		*out = make(apiv1beta1.Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubernetesTaints != nil {
		in, out := &in.KubernetesTaints, &out.KubernetesTaints
		*out = make(Taints, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(apiv1beta1.Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(NodePoolManagement)
		**out = **in
	}
	if in.LinuxNodeConfig != nil {
		in, out := &in.LinuxNodeConfig, &out.LinuxNodeConfig
		*out = new(LinuxNodeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedMachinePoolSpec.
func (in *GCPManagedMachinePoolSpec) DeepCopy() *GCPManagedMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(GCPManagedMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPManagedMachinePoolStatus) DeepCopyInto(out *GCPManagedMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPManagedMachinePoolStatus.
func (in *GCPManagedMachinePoolStatus) DeepCopy() *GCPManagedMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(GCPManagedMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocationPolicy) DeepCopyInto(out *IPAllocationPolicy) {
	*out = *in
	if in.UseIPAliases != nil {
		in, out := &in.UseIPAliases, &out.UseIPAliases
		*out = new(bool)
		**out = **in
	}
	if in.ClusterSecondaryRangeName != nil {
		in, out := &in.ClusterSecondaryRangeName, &out.ClusterSecondaryRangeName
		*out = new(string)
		**out = **in
	}
	if in.ServicesSecondaryRangeName != nil {
		in, out := &in.ServicesSecondaryRangeName, &out.ServicesSecondaryRangeName
		*out = new(string)
		**out = **in
	}
	if in.ClusterIpv4CidrBlock != nil {
		in, out := &in.ClusterIpv4CidrBlock, &out.ClusterIpv4CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.ServicesIpv4CidrBlock != nil {
		in, out := &in.ServicesIpv4CidrBlock, &out.ServicesIpv4CidrBlock
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocationPolicy.
func (in *IPAllocationPolicy) DeepCopy() *IPAllocationPolicy {
	if in == nil {
		return nil
	}
	out := new(IPAllocationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxNodeConfig) DeepCopyInto(out *LinuxNodeConfig) {
	*out = *in
	if in.Sysctls != nil {
		in, out := &in.Sysctls, &out.Sysctls
		*out = make([]SysctlConfig, len(*in))
		copy(*out, *in)
	}
	if in.CgroupMode != nil {
		in, out := &in.CgroupMode, &out.CgroupMode
		*out = new(ManagedNodePoolCgroupMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxNodeConfig.
func (in *LinuxNodeConfig) DeepCopy() *LinuxNodeConfig {
	if in == nil {
		return nil
	}
	out := new(LinuxNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfig) DeepCopyInto(out *LoggingConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfig.
func (in *LoggingConfig) DeepCopy() *LoggingConfig {
	if in == nil {
		return nil
	}
	out := new(LoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterAuthorizedNetworksConfig) DeepCopyInto(out *MasterAuthorizedNetworksConfig) {
	*out = *in
	if in.CidrBlocks != nil {
		in, out := &in.CidrBlocks, &out.CidrBlocks
		*out = make([]*MasterAuthorizedNetworksConfigCidrBlock, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MasterAuthorizedNetworksConfigCidrBlock)
				**out = **in
			}
		}
	}
	if in.GcpPublicCidrsAccessEnabled != nil {
		in, out := &in.GcpPublicCidrsAccessEnabled, &out.GcpPublicCidrsAccessEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterAuthorizedNetworksConfig.
func (in *MasterAuthorizedNetworksConfig) DeepCopy() *MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := new(MasterAuthorizedNetworksConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterAuthorizedNetworksConfigCidrBlock) DeepCopyInto(out *MasterAuthorizedNetworksConfigCidrBlock) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterAuthorizedNetworksConfigCidrBlock.
func (in *MasterAuthorizedNetworksConfigCidrBlock) DeepCopy() *MasterAuthorizedNetworksConfigCidrBlock {
	if in == nil {
		return nil
	}
	out := new(MasterAuthorizedNetworksConfigCidrBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfig) DeepCopyInto(out *MonitoringConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfig.
func (in *MonitoringConfig) DeepCopy() *MonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicy) DeepCopyInto(out *NetworkPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicy.
func (in *NetworkPolicy) DeepCopy() *NetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetworkConfig) DeepCopyInto(out *NodeNetworkConfig) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreatePodRange != nil {
		in, out := &in.CreatePodRange, &out.CreatePodRange
		*out = new(bool)
		**out = **in
	}
	if in.PodRangeName != nil {
		in, out := &in.PodRangeName, &out.PodRangeName
		*out = new(string)
		**out = **in
	}
	if in.PodRangeCidrBlock != nil {
		in, out := &in.PodRangeCidrBlock, &out.PodRangeCidrBlock
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetworkConfig.
func (in *NodeNetworkConfig) DeepCopy() *NodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NodeNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolAutoScaling) DeepCopyInto(out *NodePoolAutoScaling) {
	*out = *in
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int32)
		**out = **in
	}
	if in.EnableAutoscaling != nil {
		in, out := &in.EnableAutoscaling, &out.EnableAutoscaling
		*out = new(bool)
		**out = **in
	}
	if in.LocationPolicy != nil {
		in, out := &in.LocationPolicy, &out.LocationPolicy
		*out = new(ManagedNodePoolLocationPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolAutoScaling.
func (in *NodePoolAutoScaling) DeepCopy() *NodePoolAutoScaling {
	if in == nil {
		return nil
	}
	out := new(NodePoolAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolManagement) DeepCopyInto(out *NodePoolManagement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolManagement.
func (in *NodePoolManagement) DeepCopy() *NodePoolManagement {
	if in == nil {
		return nil
	}
	out := new(NodePoolManagement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSecurityConfig) DeepCopyInto(out *NodeSecurityConfig) {
	*out = *in
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	if in.SandboxType != nil {
		in, out := &in.SandboxType, &out.SandboxType
		*out = new(string)
		**out = **in
	}
	if in.EnableSecureBoot != nil {
		in, out := &in.EnableSecureBoot, &out.EnableSecureBoot
		*out = new(bool)
		**out = **in
	}
	if in.EnableIntegrityMonitoring != nil {
		in, out := &in.EnableIntegrityMonitoring, &out.EnableIntegrityMonitoring
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSecurityConfig.
func (in *NodeSecurityConfig) DeepCopy() *NodeSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(NodeSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCluster) DeepCopyInto(out *PrivateCluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCluster.
func (in *PrivateCluster) DeepCopy() *PrivateCluster {
	if in == nil {
		return nil
	}
	out := new(PrivateCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountConfig) DeepCopyInto(out *ServiceAccountConfig) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountConfig.
func (in *ServiceAccountConfig) DeepCopy() *ServiceAccountConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SysctlConfig) DeepCopyInto(out *SysctlConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SysctlConfig.
func (in *SysctlConfig) DeepCopy() *SysctlConfig {
	if in == nil {
		return nil
	}
	out := new(SysctlConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Taint) DeepCopyInto(out *Taint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taint.
func (in *Taint) DeepCopy() *Taint {
	if in == nil {
		return nil
	}
	out := new(Taint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Taints) DeepCopyInto(out *Taints) {
	{
		in := &in
		*out = make(Taints, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taints.
func (in Taints) DeepCopy() Taints {
	if in == nil {
		return nil
	}
	out := new(Taints)
	in.DeepCopyInto(out)
	return *out
}
