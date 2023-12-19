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

// APIServicesGetter has a method to return a APIServiceInterface.
// A group's client should implement this interface.
type APIServicesGetter interface {
	APIServices() APIServiceInterface
}

// APIServiceInterface has methods to work with APIService resources.
type APIServiceInterface interface {
	Create(ctx context.Context, aPIService *v3.APIService, opts v1.CreateOptions) (*v3.APIService, error)
	Update(ctx context.Context, aPIService *v3.APIService, opts v1.UpdateOptions) (*v3.APIService, error)
	UpdateStatus(ctx context.Context, aPIService *v3.APIService, opts v1.UpdateOptions) (*v3.APIService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.APIService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.APIServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.APIService, err error)
	APIServiceExpansion
}

// aPIServices implements APIServiceInterface
type aPIServices struct {
	client rest.Interface
}

// newAPIServices returns a APIServices
func newAPIServices(c *ManagementV3Client) *aPIServices {
	return &aPIServices{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIService, and returns the corresponding aPIService object, and an error if there is any.
func (c *aPIServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.APIService, err error) {
	result = &v3.APIService{}
	err = c.client.Get().
		Resource("apiservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIServices that match those selectors.
func (c *aPIServices) List(ctx context.Context, opts v1.ListOptions) (result *v3.APIServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.APIServiceList{}
	err = c.client.Get().
		Resource("apiservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIServices.
func (c *aPIServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apiservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIService and creates it.  Returns the server's representation of the aPIService, and an error, if there is any.
func (c *aPIServices) Create(ctx context.Context, aPIService *v3.APIService, opts v1.CreateOptions) (result *v3.APIService, err error) {
	result = &v3.APIService{}
	err = c.client.Post().
		Resource("apiservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIService and updates it. Returns the server's representation of the aPIService, and an error, if there is any.
func (c *aPIServices) Update(ctx context.Context, aPIService *v3.APIService, opts v1.UpdateOptions) (result *v3.APIService, err error) {
	result = &v3.APIService{}
	err = c.client.Put().
		Resource("apiservices").
		Name(aPIService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIServices) UpdateStatus(ctx context.Context, aPIService *v3.APIService, opts v1.UpdateOptions) (result *v3.APIService, err error) {
	result = &v3.APIService{}
	err = c.client.Put().
		Resource("apiservices").
		Name(aPIService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIService and deletes it. Returns an error if one occurs.
func (c *aPIServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apiservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apiservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIService.
func (c *aPIServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.APIService, err error) {
	result = &v3.APIService{}
	err = c.client.Patch(pt).
		Resource("apiservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
