package custominitializer

import (
	informers "pizza-apiserver/pkg/generated/informers/externalversions"

	"k8s.io/apiserver/pkg/admission"
)

type WantsRestaurantInformerFactory interface {
	SetRestaurantInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
