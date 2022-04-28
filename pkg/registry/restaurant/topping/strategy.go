package topping

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	"pizza-apiserver/pkg/apis/restaurant"
)

func NewStrategy(typer runtime.ObjectTyper) toppingStrategy {
	return toppingStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*restaurant.Topping)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Topping")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

func MatchTopping(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func SelectableFields(obj *restaurant.Topping) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type toppingStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (toppingStrategy) NamespaceScoped() bool {
	return false
}

func (toppingStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (toppingStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (toppingStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (toppingStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (toppingStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (toppingStrategy) Canonicalize(obj runtime.Object) {
}

func (toppingStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (toppingStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return []string{}
}

func (toppingStrategy) WarningsOnUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) []string {
	return []string{}
}
