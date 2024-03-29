package install

import (
	"pizza-apiserver/pkg/apis/restaurant"
	"pizza-apiserver/pkg/apis/restaurant/v1alpha1"
	"pizza-apiserver/pkg/apis/restaurant/v1beta1"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(restaurant.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must((v1beta1.AddToScheme(scheme)))
	utilruntime.Must(scheme.SetVersionPriority(v1beta1.SchemeGroupVersion, v1alpha1.SchemeGroupVersion))
}
