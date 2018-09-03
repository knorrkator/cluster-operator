package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterPhase string

const (
	ClusterPhaseInitial ClusterPhase = ""
	ClusterPhaseRunning              = "Running"

	DefaultNamespace = "storageos"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StorageOSClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []StorageOSCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StorageOSCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              StorageOSSpec          `json:"spec"`
	Status            StorageOSServiceStatus `json:"status,omitempty"`
}

type StorageOSSpec struct {
	Join               string           `json:"join"`
	CSI                StorageOSCSI     `json:"csi"`
	ResourceNS         string           `json:"namespace"`
	Service            StorageOSService `json:"service"`
	SecretRefName      string           `json:"secretRefName"`
	SecretRefNamespace string           `json:"secretRefNamespace"`
	SharedDir          string           `json:"sharedDir"`
	Ingress            StorageOSIngress `json:"ingress"`
}

// GetResourceNS returns the namespace where all the resources should be provisioned.
func (s StorageOSSpec) GetResourceNS() string {
	if s.ResourceNS != "" {
		return s.ResourceNS
	}
	return DefaultNamespace
}

type StorageOSCSI struct {
	Enable                       bool `json:"enable"`
	EnableProvisionCreds         bool `json:"enableProvisionCreds"`
	EnableControllerPublishCreds bool `json:"enableControllerPublishCreds"`
	EnableNodePublishCreds       bool `json:"enableNodePublishCreds"`
}

type StorageOSService struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	ExternalPort int               `json:"externalPort"`
	InternalPort int               `json:"internalPort"`
	Annotations  map[string]string `json:"annotations"`
}

type StorageOSIngress struct {
	Enable      bool              `json:"enable"`
	Hostname    string            `json:"hostname"`
	TLS         bool              `json:"tls"`
	Annotations map[string]string `json:"annotations"`
}

type StorageOSServiceStatus struct {
	Phase            ClusterPhase          `json:"phase"`
	NodeHealthStatus map[string]NodeHealth `json:"nodeHealthStatus,omitempty"`
	Nodes            []string              `json:"nodes"`
	Ready            string                `json:"ready"`
}

type NodeHealth struct {
	DirectfsInitiator string `json:"directfsInitiator"`
	Director          string `json:"director"`
	KV                string `json:"kv"`
	KVWrite           string `json:"kvWrite"`
	Nats              string `json:"nats"`
	Presentation      string `json:"presentation"`
	Rdb               string `json:"rdb"`
	Scheduler         string `json:"scheduler"`
}