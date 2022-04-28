package v1alpha1

func SetDefaults_PizzaSpec(obj *PizzaSpec) {
	if len(obj.Toppings) == 0 {
		obj.Toppings = []string{"salami", "mozzarella", "tomato"}
	}
}
