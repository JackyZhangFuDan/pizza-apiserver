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

package internalversion

import (
	"context"
	restaurant "pizza-apiserver/pkg/apis/restaurant"
	scheme "pizza-apiserver/pkg/generated/clientset/internalversion/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ToppingsGetter has a method to return a ToppingInterface.
// A group's client should implement this interface.
type ToppingsGetter interface {
	Toppings() ToppingInterface
}

// ToppingInterface has methods to work with Topping resources.
type ToppingInterface interface {
	Create(ctx context.Context, topping *restaurant.Topping, opts v1.CreateOptions) (*restaurant.Topping, error)
	Update(ctx context.Context, topping *restaurant.Topping, opts v1.UpdateOptions) (*restaurant.Topping, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*restaurant.Topping, error)
	List(ctx context.Context, opts v1.ListOptions) (*restaurant.ToppingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *restaurant.Topping, err error)
	ToppingExpansion
}

// toppings implements ToppingInterface
type toppings struct {
	client rest.Interface
}

// newToppings returns a Toppings
func newToppings(c *AutobusiClient) *toppings {
	return &toppings{
		client: c.RESTClient(),
	}
}

// Get takes name of the topping, and returns the corresponding topping object, and an error if there is any.
func (c *toppings) Get(ctx context.Context, name string, options v1.GetOptions) (result *restaurant.Topping, err error) {
	result = &restaurant.Topping{}
	err = c.client.Get().
		Resource("toppings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Toppings that match those selectors.
func (c *toppings) List(ctx context.Context, opts v1.ListOptions) (result *restaurant.ToppingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &restaurant.ToppingList{}
	err = c.client.Get().
		Resource("toppings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested toppings.
func (c *toppings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("toppings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a topping and creates it.  Returns the server's representation of the topping, and an error, if there is any.
func (c *toppings) Create(ctx context.Context, topping *restaurant.Topping, opts v1.CreateOptions) (result *restaurant.Topping, err error) {
	result = &restaurant.Topping{}
	err = c.client.Post().
		Resource("toppings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(topping).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a topping and updates it. Returns the server's representation of the topping, and an error, if there is any.
func (c *toppings) Update(ctx context.Context, topping *restaurant.Topping, opts v1.UpdateOptions) (result *restaurant.Topping, err error) {
	result = &restaurant.Topping{}
	err = c.client.Put().
		Resource("toppings").
		Name(topping.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(topping).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the topping and deletes it. Returns an error if one occurs.
func (c *toppings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("toppings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *toppings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("toppings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched topping.
func (c *toppings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *restaurant.Topping, err error) {
	result = &restaurant.Topping{}
	err = c.client.Patch(pt).
		Resource("toppings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
