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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	restaurant "pizza-apiserver/pkg/apis/restaurant"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeToppings implements ToppingInterface
type FakeToppings struct {
	Fake *FakeAutobusi
}

var toppingsResource = schema.GroupVersionResource{Group: "autobusi.group.pizza.restaurant", Version: "", Resource: "toppings"}

var toppingsKind = schema.GroupVersionKind{Group: "autobusi.group.pizza.restaurant", Version: "", Kind: "Topping"}

// Get takes name of the topping, and returns the corresponding topping object, and an error if there is any.
func (c *FakeToppings) Get(ctx context.Context, name string, options v1.GetOptions) (result *restaurant.Topping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(toppingsResource, name), &restaurant.Topping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*restaurant.Topping), err
}

// List takes label and field selectors, and returns the list of Toppings that match those selectors.
func (c *FakeToppings) List(ctx context.Context, opts v1.ListOptions) (result *restaurant.ToppingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(toppingsResource, toppingsKind, opts), &restaurant.ToppingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &restaurant.ToppingList{ListMeta: obj.(*restaurant.ToppingList).ListMeta}
	for _, item := range obj.(*restaurant.ToppingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested toppings.
func (c *FakeToppings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(toppingsResource, opts))
}

// Create takes the representation of a topping and creates it.  Returns the server's representation of the topping, and an error, if there is any.
func (c *FakeToppings) Create(ctx context.Context, topping *restaurant.Topping, opts v1.CreateOptions) (result *restaurant.Topping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(toppingsResource, topping), &restaurant.Topping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*restaurant.Topping), err
}

// Update takes the representation of a topping and updates it. Returns the server's representation of the topping, and an error, if there is any.
func (c *FakeToppings) Update(ctx context.Context, topping *restaurant.Topping, opts v1.UpdateOptions) (result *restaurant.Topping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(toppingsResource, topping), &restaurant.Topping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*restaurant.Topping), err
}

// Delete takes name of the topping and deletes it. Returns an error if one occurs.
func (c *FakeToppings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(toppingsResource, name, opts), &restaurant.Topping{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeToppings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(toppingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &restaurant.ToppingList{})
	return err
}

// Patch applies the patch and returns the patched topping.
func (c *FakeToppings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *restaurant.Topping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(toppingsResource, name, pt, data, subresources...), &restaurant.Topping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*restaurant.Topping), err
}
