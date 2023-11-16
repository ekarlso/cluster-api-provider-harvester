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

// TemplatesGetter has a method to return a TemplateInterface.
// A group's client should implement this interface.
type TemplatesGetter interface {
	Templates() TemplateInterface
}

// TemplateInterface has methods to work with Template resources.
type TemplateInterface interface {
	Create(ctx context.Context, template *v3.Template, opts v1.CreateOptions) (*v3.Template, error)
	Update(ctx context.Context, template *v3.Template, opts v1.UpdateOptions) (*v3.Template, error)
	UpdateStatus(ctx context.Context, template *v3.Template, opts v1.UpdateOptions) (*v3.Template, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.Template, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.TemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.Template, err error)
	TemplateExpansion
}

// templates implements TemplateInterface
type templates struct {
	client rest.Interface
}

// newTemplates returns a Templates
func newTemplates(c *ManagementV3Client) *templates {
	return &templates{
		client: c.RESTClient(),
	}
}

// Get takes name of the template, and returns the corresponding template object, and an error if there is any.
func (c *templates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.Template, err error) {
	result = &v3.Template{}
	err = c.client.Get().
		Resource("templates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Templates that match those selectors.
func (c *templates) List(ctx context.Context, opts v1.ListOptions) (result *v3.TemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.TemplateList{}
	err = c.client.Get().
		Resource("templates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested templates.
func (c *templates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("templates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a template and creates it.  Returns the server's representation of the template, and an error, if there is any.
func (c *templates) Create(ctx context.Context, template *v3.Template, opts v1.CreateOptions) (result *v3.Template, err error) {
	result = &v3.Template{}
	err = c.client.Post().
		Resource("templates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(template).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a template and updates it. Returns the server's representation of the template, and an error, if there is any.
func (c *templates) Update(ctx context.Context, template *v3.Template, opts v1.UpdateOptions) (result *v3.Template, err error) {
	result = &v3.Template{}
	err = c.client.Put().
		Resource("templates").
		Name(template.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(template).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *templates) UpdateStatus(ctx context.Context, template *v3.Template, opts v1.UpdateOptions) (result *v3.Template, err error) {
	result = &v3.Template{}
	err = c.client.Put().
		Resource("templates").
		Name(template.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(template).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the template and deletes it. Returns an error if one occurs.
func (c *templates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("templates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *templates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("templates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched template.
func (c *templates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.Template, err error) {
	result = &v3.Template{}
	err = c.client.Patch(pt).
		Resource("templates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
