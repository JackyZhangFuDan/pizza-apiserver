package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Pizza struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   PizzaSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status PizzaStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PizzaSpec struct {
	// +k8s:conversion-gen=false
	Toppings []string `json:"toppings" protobuf:"bytes,1,rep,name=toppings"`
}

type PizzaStatus struct {
	Cost float64 `json:"cost,omitempty" protobuf:"bytes,1,opt,name=cost"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PizzaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Pizza `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Topping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec ToppingSpec
}

type ToppingSpec struct {
	Cost float64 `json:"cost" protobuf:"bytes,1,name=cost"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ToppingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Topping `json:"items" protobuf:"bytes,2,rep,name=items"`
}
