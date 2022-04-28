package restaurant

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Pizza struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   PizzaSpec
	Status PizzaStatus
}

type PizzaSpec struct {
	Toppings []PizzaTopping
}

type PizzaTopping struct {
	Name     string
	Quantity int
}

type PizzaStatus struct {
	Cost float64
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PizzaList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Pizza
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Topping struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec ToppingSpec
}

type ToppingSpec struct {
	Cost float64
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ToppingList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Topping
}
