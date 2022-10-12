package custominitializer

import (
	informers "pizza-apiserver/pkg/generated/informers/externalversions"

	"k8s.io/apiserver/pkg/admission"
)

type restaurantInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = restaurantInformerPluginInitializer{}

// 做一个具有initialize方法的结构体实例
func New(informers informers.SharedInformerFactory) restaurantInformerPluginInitializer {
	return restaurantInformerPluginInitializer{
		informers: informers,
	}
}

// server启动时在config阶段被调用，从而把指定的plugin初始化掉
func (i restaurantInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsRestaurantInformerFactory); ok {
		wants.SetRestaurantInformerFactory(i.informers)
	}
}
