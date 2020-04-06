package factory

import "fmt"

//Pizza interface of pizza
type Pizza interface {
	prepare()
	bake()
}

// PizzaFactory interface of PizzaFactory
type PizzaFactory interface {
	Createpizza(string) Pizza
}

//CheesePizza one of the pizza
type CheesePizza struct {
	name string
}

func (pizza CheesePizza) prepare() {
	fmt.Printf("pizza %s is prepared\n", pizza.name)
}
func (pizza CheesePizza) bake() {
	fmt.Printf("pizza %s is bake\n", pizza.name)
}
