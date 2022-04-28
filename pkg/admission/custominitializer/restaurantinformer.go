package custominitializer

import (
	informers "pizza-apiserver/pkg/generated/informers/externalversions"

	"k8s.io/apiserver/pkg/admission"
)

type restaurantInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = restaurantInformerPluginInitializer{}

func New(informers informers.SharedInformerFactory) restaurantInformerPluginInitializer {
	return restaurantInformerPluginInitializer{
		informers: informers,
	}
}

func (i restaurantInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsRestaurantInformerFactory); ok {
		wants.SetRestaurantInformerFactory(i.informers)
	}
}
