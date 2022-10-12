package custominitializer

import (
	informers "pizza-apiserver/pkg/generated/informers/externalversions"

	"k8s.io/apiserver/pkg/admission"
)

// 需要plugin去实现这个接口，从而保证可以接收informerfactory；
// 其实接口内容自己定义的，不必拘泥于一下定义，这里只是copy了kube api server的做法
type WantsRestaurantInformerFactory interface {
	SetRestaurantInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
