package factory

//SimplePizzaFactory can create pizza
type SimplePizzaFactory struct {
}

//NewSimplePizzaFactory return a new NewSimplePizzaFactory
func NewSimplePizzaFactory() SimplePizzaFactory {
	return SimplePizzaFactory{}
}

//Createpizza create pizza
func (SimplePizzaFactory) Createpizza(pizzatype string) Pizza {
	if pizzatype == "cheese" {
		return CheesePizza{
			name: "cheese",
		}
	}
	return CheesePizza{
		name: "defaultcheese",
	}
}

// BeijingPizzaFactory represent the beijing pizza factory
type BeijingPizzaFactory struct {
}

// NewBeijingPizzaFactory start beijingpizza factory
func NewBeijingPizzaFactory() BeijingPizzaFactory {
	return BeijingPizzaFactory{}
}

//Createpizza create pizza
func (BeijingPizzaFactory) Createpizza(pizzatype string) Pizza {
	if pizzatype == "beijingcheese" {
		return CheesePizza{
			name: "beijingcheese",
		}
	}
	return CheesePizza{
		name: "defaultbeijingcheese",
	}
}
