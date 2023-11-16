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

package v3

import (
	"context"
	"time"

	scheme "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterAlertGroupsGetter has a method to return a ClusterAlertGroupInterface.
// A group's client should implement this interface.
type ClusterAlertGroupsGetter interface {
	ClusterAlertGroups(namespace string) ClusterAlertGroupInterface
}

// ClusterAlertGroupInterface has methods to work with ClusterAlertGroup resources.
type ClusterAlertGroupInterface interface {
	Create(ctx context.Context, clusterAlertGroup *v3.ClusterAlertGroup, opts v1.CreateOptions) (*v3.ClusterAlertGroup, error)
	Update(ctx context.Context, clusterAlertGroup *v3.ClusterAlertGroup, opts v1.UpdateOptions) (*v3.ClusterAlertGroup, error)
	UpdateStatus(ctx context.Context, clusterAlertGroup *v3.ClusterAlertGroup, opts v1.UpdateOptions) (*v3.ClusterAlertGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ClusterAlertGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ClusterAlertGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterAlertGroup, err error)
	ClusterAlertGroupExpansion
}

// clusterAlertGroups implements ClusterAlertGroupInterface
type clusterAlertGroups struct {
	client rest.Interface
	ns     string
}

// newClusterAlertGroups returns a ClusterAlertGroups
func newClusterAlertGroups(c *ManagementV3Client, namespace string) *clusterAlertGroups {
	return &clusterAlertGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterAlertGroup, and returns the corresponding clusterAlertGroup object, and an error if there is any.
func (c *clusterAlertGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterAlertGroup, err error) {
	result = &v3.ClusterAlertGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterAlertGroups that match those selectors.
func (c *clusterAlertGroups) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterAlertGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.ClusterAlertGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterAlertGroups.
func (c *clusterAlertGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterAlertGroup and creates it.  Returns the server's representation of the clusterAlertGroup, and an error, if there is any.
func (c *clusterAlertGroups) Create(ctx context.Context, clusterAlertGroup *v3.ClusterAlertGroup, opts v1.CreateOptions) (result *v3.ClusterAlertGroup, err error) {
	result = &v3.ClusterAlertGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAlertGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterAlertGroup and updates it. Returns the server's representation of the clusterAlertGroup, and an error, if there is any.
func (c *clusterAlertGroups) Update(ctx context.Context, clusterAlertGroup *v3.ClusterAlertGroup, opts v1.UpdateOptions) (result *v3.ClusterAlertGroup, err error) {
	result = &v3.ClusterAlertGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		Name(clusterAlertGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAlertGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterAlertGroups) UpdateStatus(ctx context.Context, clusterAlertGroup *v3.ClusterAlertGroup, opts v1.UpdateOptions) (result *v3.ClusterAlertGroup, err error) {
	result = &v3.ClusterAlertGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		Name(clusterAlertGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAlertGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterAlertGroup and deletes it. Returns an error if one occurs.
func (c *clusterAlertGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterAlertGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusteralertgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterAlertGroup.
func (c *clusterAlertGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterAlertGroup, err error) {
	result = &v3.ClusterAlertGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusteralertgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
