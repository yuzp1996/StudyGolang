package factory

import "fmt"

// PizzaStore create and sale pizza
type PizzaStore struct {
	Factory PizzaFactory
}

func (pizzastore PizzaStore) orderPizza(pizzatype string) Pizza {
	fmt.Println("This is order progress,It is  commom")
	fmt.Println("someone order pizza, and the info will be copy to the master company money mangement system")
	pizza := pizzastore.Factory.Createpizza(pizzatype)
	return pizza
}

// NewPizzaStore create a new pizza store
func NewPizzaStore(factory PizzaFactory) PizzaStore {
	return PizzaStore{
		Factory: factory,
	}
}
