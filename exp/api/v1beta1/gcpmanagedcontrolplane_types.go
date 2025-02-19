/*
Copyright 2022 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const (
	// ManagedControlPlaneFinalizer allows Reconcile to clean up GCP resources associated with the GCPManagedControlPlane before
	// removing it from the apiserver.
	ManagedControlPlaneFinalizer = "gcpmanagedcontrolplane.infrastructure.cluster.x-k8s.io"
)

type PrivateCluster struct {
	// EnablePrivateEndpoint: Whether the master's internal IP
	// address is used as the cluster endpoint.
	// +optional
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint,omitempty"`
	// EnablePrivateNodes: Whether nodes have internal IP
	// addresses only. If enabled, all nodes are given only RFC
	// 1918 private addresses and communicate with the master via
	// private networking.
	// +optional
	EnablePrivateNodes bool `json:"enablePrivateNodes,omitempty"`
	// ControlPlaneCidrBlock is the IP range in CIDR notation to use for the hosted master network. This range must not
	// overlap with any other ranges in use within the cluster's network. Honored when enabled is true.
	// +optional
	ControlPlaneCidrBlock string `json:"controlPlaneCidrBlock,omitempty"`
}

// ClusterNetwork define the cluster network.
type ClusterNetwork struct {
	// PrivateCluster defines the private cluster spec.
	// +optional
	PrivateCluster *PrivateCluster `json:"privateCluster,omitempty"`
}

// LoggingConfig defines the logging on Cluster.
type LoggingConfig struct {
	// SystemComponents enables the system component logging.
	// +optional
	SystemComponents bool `json:"systemComponents,omitempty"`
	// Workloads enables the Workloads logging.
	// +optional
	Workloads bool `json:"workloads,omitempty"`
}

// MonitoringConfig defines the monitoring on Cluster.
type MonitoringConfig struct {
	// EnableManagedPrometheus Enable Google Cloud Managed Service for Prometheus in the cluster.
	// +optional
	EnableManagedPrometheus bool `json:"enableManagedPrometheus,omitempty"`
}

// GCPManagedControlPlaneSpec defines the desired state of GCPManagedControlPlane.
type GCPManagedControlPlaneSpec struct {
	// ClusterName allows you to specify the name of the GKE cluster.
	// If you don't specify a name then a default name will be created
	// based on the namespace and name of the managed control plane.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`
	// ClusterNetwork define the cluster network.
	// +optional
	ClusterNetwork *ClusterNetwork `json:"clusterNetwork,omitempty"`
	// LoggingConfig defines the logging on Cluster.
	// +optional
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
	// MonitoringConfig defines the monitoring on Cluster.
	// +optional
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`
	// Project is the name of the project to deploy the cluster to.
	Project string `json:"project"`
	// Location represents the location (region or zone) in which the GKE cluster
	// will be created.
	Location string `json:"location"`
	// ClusterIpv4Cidr is the IP address range of the container pods in the GKE cluster, in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `10.96.0.0/14`).
	// If not specified then one will be automatically chosen.
	// If this field is specified then IPAllocationPolicy.ClusterIpv4CidrBlock should be left blank.
	// +optional
	ClusterIpv4Cidr *string `json:"clusterIpv4Cidr,omitempty"`
	// IPAllocationPolicy represents configuration options for GKE cluster IP allocation.
	// If not specified then GKE default values will be used.
	// +optional
	IPAllocationPolicy *IPAllocationPolicy `json:"ipAllocationPolicy,omitempty"`
	// EnableAutopilot indicates whether to enable autopilot for this GKE cluster.
	// +optional
	EnableAutopilot bool `json:"enableAutopilot"`
	// ReleaseChannel represents the release channel of the GKE cluster.
	// +optional
	ReleaseChannel *ReleaseChannel `json:"releaseChannel,omitempty"`
	// ControlPlaneVersion represents the control plane version of the GKE cluster.
	// If not specified, the default version currently supported by GKE will be
	// used.
	// +optional
	ControlPlaneVersion *string `json:"controlPlaneVersion,omitempty"`
	// Endpoint represents the endpoint used to communicate with the control plane.
	// +optional
	Endpoint clusterv1.APIEndpoint `json:"endpoint"`
	// MasterAuthorizedNetworksConfig represents configuration options for master authorized networks feature of the GKE cluster.
	// This feature is disabled if this field is not specified.
	// +optional
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"master_authorized_networks_config,omitempty"`
	// NetworkPolicy represents configuration options for NetworkPolicy feature of the GKE cluster.
	// This feature is disabled if this field is not specified.
	// +optional
	NetworkPolicy *NetworkPolicy `json:"networkPolicy,omitempty"`
}

// GCPManagedControlPlaneStatus defines the observed state of GCPManagedControlPlane.
type GCPManagedControlPlaneStatus struct {
	// Ready denotes that the GCPManagedControlPlane API Server is ready to
	// receive requests.
	// +kubebuilder:default=false
	Ready bool `json:"ready"`

	// Initialized is true when the control plane is available for initial contact.
	// This may occur before the control plane is fully ready.
	// +optional
	Initialized bool `json:"initialized,omitempty"`

	// Conditions specifies the conditions for the managed control plane
	Conditions clusterv1.Conditions `json:"conditions,omitempty"`

	// CurrentVersion shows the current version of the GKE control plane.
	// +optional
	CurrentVersion string `json:"currentVersion,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=gcpmanagedcontrolplanes,scope=Namespaced,categories=cluster-api,shortName=gcpmcp
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".metadata.labels.cluster\\.x-k8s\\.io/cluster-name",description="Cluster to which this GCPManagedControlPlane belongs"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Control plane is ready"
// +kubebuilder:printcolumn:name="CurrentVersion",type="string",JSONPath=".status.currentVersion",description="The current Kubernetes version"
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".spec.endpoint",description="API Endpoint",priority=1

// GCPManagedControlPlane is the Schema for the gcpmanagedcontrolplanes API.
type GCPManagedControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPManagedControlPlaneSpec   `json:"spec,omitempty"`
	Status GCPManagedControlPlaneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GCPManagedControlPlaneList contains a list of GCPManagedControlPlane.
type GCPManagedControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPManagedControlPlane `json:"items"`
}

// ReleaseChannel is the release channel of the GKE cluster
// +kubebuilder:validation:Enum=rapid;regular;stable;extended
type ReleaseChannel string

const (
	// Rapid release channel.
	Rapid ReleaseChannel = "rapid"
	// Regular release channel.
	Regular ReleaseChannel = "regular"
	// Stable release channel.
	Stable ReleaseChannel = "stable"
	// Extended release channel.
	Extended ReleaseChannel = "extended"
)

// IPAllocationPolicy represents configuration options for GKE cluster IP allocation.
type IPAllocationPolicy struct {
	// UseIPAliases represents whether alias IPs will be used for pod IPs in the cluster.
	// If unspecified will default to false.
	// +optional
	UseIPAliases *bool `json:"useIPAliases,omitempty"`
	// ClusterSecondaryRangeName represents the name of the secondary range to be used for the GKE cluster CIDR block.
	// The range will be used for pod IP addresses and must be an existing secondary range associated with the cluster subnetwork.
	// This field is only applicable when use_ip_aliases is set to true.
	// +optional
	ClusterSecondaryRangeName *string `json:"clusterSecondaryRangeName,omitempty"`
	// ServicesSecondaryRangeName represents the name of the secondary range to be used for the services CIDR block.
	// The range will be used for service ClusterIPs and must be an existing secondary range associated with the cluster subnetwork.
	// This field is only applicable when use_ip_aliases is set to true.
	// +optional
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName,omitempty"`
	// ClusterIpv4CidrBlock represents the IP address range for the GKE cluster pod IPs. If this field is set, then
	// GCPManagedControlPlaneSpec.ClusterIpv4Cidr must be left blank.
	// This field is only applicable when use_ip_aliases is set to true.
	// If not specified the range will be chosen with the default size.
	// +optional
	ClusterIpv4CidrBlock *string `json:"clusterIpv4CidrBlock,omitempty"`
	// ServicesIpv4CidrBlock represents the IP address range for services IPs in the GKE cluster.
	// This field is only applicable when use_ip_aliases is set to true.
	// If not specified the range will be chosen with the default size.
	// +optional
	ServicesIpv4CidrBlock *string `json:"servicesIpv4CidrBlock,omitempty"`
}

// MasterAuthorizedNetworksConfig contains configuration options for the master authorized networks feature.
// Enabled master authorized networks will disallow all external traffic to access
// Kubernetes master through HTTPS except traffic from the given CIDR blocks,
// Google Compute Engine Public IPs and Google Prod IPs.
type MasterAuthorizedNetworksConfig struct {
	// cidr_blocks define up to 50 external networks that could access
	// Kubernetes master through HTTPS.
	// +optional
	CidrBlocks []*MasterAuthorizedNetworksConfigCidrBlock `json:"cidr_blocks,omitempty"`
	// Whether master is accessible via Google Compute Engine Public IP addresses.
	// +optional
	GcpPublicCidrsAccessEnabled *bool `json:"gcp_public_cidrs_access_enabled,omitempty"`
}

// MasterAuthorizedNetworksConfigCidrBlock contains an optional name and one CIDR block.
type MasterAuthorizedNetworksConfigCidrBlock struct {
	// display_name is an field for users to identify CIDR blocks.
	DisplayName string `json:"display_name,omitempty"`
	// cidr_block must be specified in CIDR notation.
	// +kubebuilder:validation:Pattern=`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}(?:\/([0-9]|[1-2][0-9]|3[0-2]))?$|^([a-fA-F0-9:]+:+)+[a-fA-F0-9]+\/[0-9]{1,3}$`
	CidrBlock string `json:"cidr_block,omitempty"`
}

// NetworkPolicy represents configuration options for NetworkPolicy feature of the GKE cluster.
type NetworkPolicy struct {
	// The selected network policy provider.
	// +kubebuilder:validation:Enum=calico
	// +optional
	Provider string `json:"provider,omitempty"`
}

// GetConditions returns the control planes conditions.
func (r *GCPManagedControlPlane) GetConditions() clusterv1.Conditions {
	return r.Status.Conditions
}

// SetConditions sets the status conditions for the GCPManagedControlPlane.
func (r *GCPManagedControlPlane) SetConditions(conditions clusterv1.Conditions) {
	r.Status.Conditions = conditions
}

func init() {
	SchemeBuilder.Register(&GCPManagedControlPlane{}, &GCPManagedControlPlaneList{})
}
