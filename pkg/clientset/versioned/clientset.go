/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	catalogv1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/catalog.cattle.io/v1"
	clusterv1alpha4 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/cluster.x-k8s.io/v1alpha4"
	harvesterhciv1beta1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/harvesterhci.io/v1beta1"
	k8scnicncfiov1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	kubevirtv1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/kubevirt.io/v1"
	longhornv1beta2 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/longhorn.io/v1beta2"
	managementv3 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/management.cattle.io/v3"
	monitoringv1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/monitoring.coreos.com/v1"
	networkingv1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/networking.k8s.io/v1"
	snapshotv1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/snapshot.storage.k8s.io/v1"
	storagev1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/storage.k8s.io/v1"
	upgradev1 "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/typed/upgrade.cattle.io/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	lbv1 "github.com/harvester/harvester-load-balancer/pkg/generated/clientset/versioned/typed/loadbalancer.harvesterhci.io/v1beta1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CatalogV1() catalogv1.CatalogV1Interface
	ClusterV1alpha4() clusterv1alpha4.ClusterV1alpha4Interface
	HarvesterhciV1beta1() harvesterhciv1beta1.HarvesterhciV1beta1Interface
	K8sCniCncfIoV1() k8scnicncfiov1.K8sCniCncfIoV1Interface
	KubevirtV1() kubevirtv1.KubevirtV1Interface
	LonghornV1beta2() longhornv1beta2.LonghornV1beta2Interface
	ManagementV3() managementv3.ManagementV3Interface
	MonitoringV1() monitoringv1.MonitoringV1Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
	SnapshotV1() snapshotv1.SnapshotV1Interface
	StorageV1() storagev1.StorageV1Interface
	UpgradeV1() upgradev1.UpgradeV1Interface
	CoreV1() corev1.CoreV1Interface
	LoadbalancerV1beta1() lbv1.LoadbalancerV1beta1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	catalogV1           *catalogv1.CatalogV1Client
	clusterV1alpha4     *clusterv1alpha4.ClusterV1alpha4Client
	harvesterhciV1beta1 *harvesterhciv1beta1.HarvesterhciV1beta1Client
	k8sCniCncfIoV1      *k8scnicncfiov1.K8sCniCncfIoV1Client
	kubevirtV1          *kubevirtv1.KubevirtV1Client
	longhornV1beta2     *longhornv1beta2.LonghornV1beta2Client
	managementV3        *managementv3.ManagementV3Client
	monitoringV1        *monitoringv1.MonitoringV1Client
	networkingV1        *networkingv1.NetworkingV1Client
	snapshotV1          *snapshotv1.SnapshotV1Client
	storageV1           *storagev1.StorageV1Client
	upgradeV1           *upgradev1.UpgradeV1Client
	corev1							*corev1.CoreV1Client
	lbv1beta1 				 *lbv1.LoadbalancerV1beta1Client
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	if c == nil {
		return nil
	}
	return c.corev1
}

// CatalogV1 retrieves the CatalogV1Client
func (c *Clientset) CatalogV1() catalogv1.CatalogV1Interface {
	return c.catalogV1
}

// ClusterV1alpha4 retrieves the ClusterV1alpha4Client
func (c *Clientset) ClusterV1alpha4() clusterv1alpha4.ClusterV1alpha4Interface {
	return c.clusterV1alpha4
}

// HarvesterhciV1beta1 retrieves the HarvesterhciV1beta1Client
func (c *Clientset) HarvesterhciV1beta1() harvesterhciv1beta1.HarvesterhciV1beta1Interface {
	return c.harvesterhciV1beta1
}

// K8sCniCncfIoV1 retrieves the K8sCniCncfIoV1Client
func (c *Clientset) K8sCniCncfIoV1() k8scnicncfiov1.K8sCniCncfIoV1Interface {
	return c.k8sCniCncfIoV1
}

// KubevirtV1 retrieves the KubevirtV1Client
func (c *Clientset) KubevirtV1() kubevirtv1.KubevirtV1Interface {
	return c.kubevirtV1
}

// LonghornV1beta2 retrieves the LonghornV1beta2Client
func (c *Clientset) LonghornV1beta2() longhornv1beta2.LonghornV1beta2Interface {
	return c.longhornV1beta2
}

// ManagementV3 retrieves the ManagementV3Client
func (c *Clientset) ManagementV3() managementv3.ManagementV3Interface {
	return c.managementV3
}

// MonitoringV1 retrieves the MonitoringV1Client
func (c *Clientset) MonitoringV1() monitoringv1.MonitoringV1Interface {
	return c.monitoringV1
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return c.networkingV1
}

// SnapshotV1 retrieves the SnapshotV1Client
func (c *Clientset) SnapshotV1() snapshotv1.SnapshotV1Interface {
	return c.snapshotV1
}

// StorageV1 retrieves the StorageV1Client
func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
	return c.storageV1
}

// UpgradeV1 retrieves the UpgradeV1Client
func (c *Clientset) UpgradeV1() upgradev1.UpgradeV1Interface {
	return c.upgradeV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// LoadBalancerV1beta1 retrieves the LoadBalancerV1beta1Client
func (c *Clientset) LoadbalancerV1beta1() lbv1.LoadbalancerV1beta1Interface {
	return c.lbv1beta1
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.catalogV1, err = catalogv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.clusterV1alpha4, err = clusterv1alpha4.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.harvesterhciV1beta1, err = harvesterhciv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.k8sCniCncfIoV1, err = k8scnicncfiov1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kubevirtV1, err = kubevirtv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.longhornV1beta2, err = longhornv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.managementV3, err = managementv3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.monitoringV1, err = monitoringv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.snapshotV1, err = snapshotv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1, err = storagev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.upgradeV1, err = upgradev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.corev1, err = corev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.lbv1beta1, err = lbv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.catalogV1 = catalogv1.New(c)
	cs.clusterV1alpha4 = clusterv1alpha4.New(c)
	cs.harvesterhciV1beta1 = harvesterhciv1beta1.New(c)
	cs.k8sCniCncfIoV1 = k8scnicncfiov1.New(c)
	cs.kubevirtV1 = kubevirtv1.New(c)
	cs.longhornV1beta2 = longhornv1beta2.New(c)
	cs.managementV3 = managementv3.New(c)
	cs.monitoringV1 = monitoringv1.New(c)
	cs.networkingV1 = networkingv1.New(c)
	cs.snapshotV1 = snapshotv1.New(c)
	cs.storageV1 = storagev1.New(c)
	cs.upgradeV1 = upgradev1.New(c)
	cs.corev1 = corev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
