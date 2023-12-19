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

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterApps implements MultiClusterAppInterface
type FakeMultiClusterApps struct {
	Fake *FakeManagementV3
	ns   string
}

var multiclusterappsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "multiclusterapps"}

var multiclusterappsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MultiClusterApp"}

// Get takes name of the multiClusterApp, and returns the corresponding multiClusterApp object, and an error if there is any.
func (c *FakeMultiClusterApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.MultiClusterApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterappsResource, c.ns, name), &v3.MultiClusterApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterApp), err
}

// List takes label and field selectors, and returns the list of MultiClusterApps that match those selectors.
func (c *FakeMultiClusterApps) List(ctx context.Context, opts v1.ListOptions) (result *v3.MultiClusterAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterappsResource, multiclusterappsKind, c.ns, opts), &v3.MultiClusterAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.MultiClusterAppList{ListMeta: obj.(*v3.MultiClusterAppList).ListMeta}
	for _, item := range obj.(*v3.MultiClusterAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterApps.
func (c *FakeMultiClusterApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterappsResource, c.ns, opts))

}

// Create takes the representation of a multiClusterApp and creates it.  Returns the server's representation of the multiClusterApp, and an error, if there is any.
func (c *FakeMultiClusterApps) Create(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.CreateOptions) (result *v3.MultiClusterApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterappsResource, c.ns, multiClusterApp), &v3.MultiClusterApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterApp), err
}

// Update takes the representation of a multiClusterApp and updates it. Returns the server's representation of the multiClusterApp, and an error, if there is any.
func (c *FakeMultiClusterApps) Update(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.UpdateOptions) (result *v3.MultiClusterApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterappsResource, c.ns, multiClusterApp), &v3.MultiClusterApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterApps) UpdateStatus(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.UpdateOptions) (*v3.MultiClusterApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusterappsResource, "status", c.ns, multiClusterApp), &v3.MultiClusterApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterApp), err
}

// Delete takes name of the multiClusterApp and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(multiclusterappsResource, c.ns, name, opts), &v3.MultiClusterApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterappsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.MultiClusterAppList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterApp.
func (c *FakeMultiClusterApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.MultiClusterApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterappsResource, c.ns, name, pt, data, subresources...), &v3.MultiClusterApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterApp), err
}
