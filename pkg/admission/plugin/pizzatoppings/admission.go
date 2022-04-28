package pizzatoppings

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"

	"pizza-apiserver/pkg/admission/custominitializer"
	"pizza-apiserver/pkg/apis/restaurant"
	informers "pizza-apiserver/pkg/generated/informers/externalversions"
	listers "pizza-apiserver/pkg/generated/listers/restaurant/v1alpha1"
)

func Register(plugins *admission.Plugins) {
	plugins.Register("PizzaToppings", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

type PizzaToppingsPlugin struct {
	*admission.Handler
	toppingLister listers.ToppingLister
}

var _ = custominitializer.WantsRestaurantInformerFactory(&PizzaToppingsPlugin{})
var _ = admission.ValidationInterface(&PizzaToppingsPlugin{})

func (d *PizzaToppingsPlugin) Validate(ctx context.Context, a admission.Attributes, _ admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != restaurant.Kind("Pizza") {
		return nil
	}

	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	obj := a.GetOldObject()
	pizza := obj.(*restaurant.Pizza)
	for _, top := range pizza.Spec.Toppings {
		if _, err := d.toppingLister.Get(top.Name); err != nil && errors.IsNotFound(err) {
			return admission.NewForbidden(
				a,
				fmt.Errorf("unknown topping: %s", top.Name),
			)
		}
	}

	return nil
}

func (d *PizzaToppingsPlugin) SetRestaurantInformerFactory(f informers.SharedInformerFactory) {
	d.toppingLister = f.Autobusi().V1alpha1().Toppings().Lister()
	d.SetReadyFunc(f.Autobusi().V1alpha1().Toppings().Informer().HasSynced)
}

func (d *PizzaToppingsPlugin) ValidateInitialization() error {
	if d.toppingLister == nil {
		return fmt.Errorf("missing policy lister")
	}
	return nil
}

func New() (*PizzaToppingsPlugin, error) {
	return &PizzaToppingsPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
