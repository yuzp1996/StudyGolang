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
	ingredientfactory PizzaIngredientFacotry
	name              string
}

func (pizza CheesePizza) prepare() {
	fmt.Printf("CheesePizza use salt %v use sugger %v\n", pizza.ingredientfactory.createsalt(), pizza.ingredientfactory.createsugger())
	fmt.Printf("Cheese pizza %s is prepared\n", pizza.name)
}
func (pizza CheesePizza) bake() {
	fmt.Printf("Cheese pizza %s is bake\n", pizza.name)
}

//BeijingCheesePizza one of the pizza
type BeijingCheesePizza struct {
	ingredientfactory PizzaIngredientFacotry
	name              string
}

func (pizza BeijingCheesePizza) prepare() {
	fmt.Printf("BeijingPizza use salt %v use sugger %v\n", pizza.ingredientfactory.createsalt(), pizza.ingredientfactory.createsugger())
	fmt.Printf("BeijingCheese pizza %s is prepared\n", pizza.name)
}
func (pizza BeijingCheesePizza) bake() {
	fmt.Printf("BeijingCheese pizza %s is bake\n", pizza.name)
}
