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

package v1beta2

import (
	"context"
	"time"

	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	scheme "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnginesGetter has a method to return a EngineInterface.
// A group's client should implement this interface.
type EnginesGetter interface {
	Engines(namespace string) EngineInterface
}

// EngineInterface has methods to work with Engine resources.
type EngineInterface interface {
	Create(ctx context.Context, engine *v1beta2.Engine, opts v1.CreateOptions) (*v1beta2.Engine, error)
	Update(ctx context.Context, engine *v1beta2.Engine, opts v1.UpdateOptions) (*v1beta2.Engine, error)
	UpdateStatus(ctx context.Context, engine *v1beta2.Engine, opts v1.UpdateOptions) (*v1beta2.Engine, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.Engine, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.EngineList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.Engine, err error)
	EngineExpansion
}

// engines implements EngineInterface
type engines struct {
	client rest.Interface
	ns     string
}

// newEngines returns a Engines
func newEngines(c *LonghornV1beta2Client, namespace string) *engines {
	return &engines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the engine, and returns the corresponding engine object, and an error if there is any.
func (c *engines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.Engine, err error) {
	result = &v1beta2.Engine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("engines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Engines that match those selectors.
func (c *engines) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.EngineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.EngineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested engines.
func (c *engines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a engine and creates it.  Returns the server's representation of the engine, and an error, if there is any.
func (c *engines) Create(ctx context.Context, engine *v1beta2.Engine, opts v1.CreateOptions) (result *v1beta2.Engine, err error) {
	result = &v1beta2.Engine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(engine).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a engine and updates it. Returns the server's representation of the engine, and an error, if there is any.
func (c *engines) Update(ctx context.Context, engine *v1beta2.Engine, opts v1.UpdateOptions) (result *v1beta2.Engine, err error) {
	result = &v1beta2.Engine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("engines").
		Name(engine.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(engine).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *engines) UpdateStatus(ctx context.Context, engine *v1beta2.Engine, opts v1.UpdateOptions) (result *v1beta2.Engine, err error) {
	result = &v1beta2.Engine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("engines").
		Name(engine.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(engine).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the engine and deletes it. Returns an error if one occurs.
func (c *engines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("engines").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *engines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched engine.
func (c *engines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.Engine, err error) {
	result = &v1beta2.Engine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("engines").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
