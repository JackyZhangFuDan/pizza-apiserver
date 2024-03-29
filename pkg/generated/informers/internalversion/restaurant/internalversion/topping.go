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

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	restaurant "pizza-apiserver/pkg/apis/restaurant"
	clientsetinternalversion "pizza-apiserver/pkg/generated/clientset/internalversion"
	internalinterfaces "pizza-apiserver/pkg/generated/informers/internalversion/internalinterfaces"
	internalversion "pizza-apiserver/pkg/generated/listers/restaurant/internalversion"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ToppingInformer provides access to a shared informer and lister for
// Toppings.
type ToppingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ToppingLister
}

type toppingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewToppingInformer constructs a new informer for Topping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewToppingInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredToppingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredToppingInformer constructs a new informer for Topping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredToppingInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Autobusi().Toppings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Autobusi().Toppings().Watch(context.TODO(), options)
			},
		},
		&restaurant.Topping{},
		resyncPeriod,
		indexers,
	)
}

func (f *toppingInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredToppingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *toppingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&restaurant.Topping{}, f.defaultInformer)
}

func (f *toppingInformer) Lister() internalversion.ToppingLister {
	return internalversion.NewToppingLister(f.Informer().GetIndexer())
}
