package validation

import (
	"pizza-apiserver/pkg/apis/restaurant"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidatePizza(f *restaurant.Pizza) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidatePizzaSpec(&f.Spec, field.NewPath("spec"))...)
	return allErrs
}

func ValidatePizzaSpec(s *restaurant.PizzaSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	prevNames := map[string]bool{}
	for i := range s.Toppings {
		if s.Toppings[i].Quantity <= 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("topping").Index(i), s.Toppings[i].Quantity, "cannot be negative or zero"))
		}
		if len(s.Toppings[i].Name) == 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("toppings").Index(i).Child("name"), s.Toppings[i].Name, "cannot be empty"))
		} else {
			if prevNames[s.Toppings[i].Name] {
				allErrs = append(allErrs, field.Invalid(fldPath.Child("toppings").Index(i).Child("name"), s.Toppings[i].Name, "must be unique"))
			}
			prevNames[s.Toppings[i].Name] = true
		}
	}

	return allErrs
}
