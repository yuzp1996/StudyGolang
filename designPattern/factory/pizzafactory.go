package factory

//SimplePizzaFactory can create pizza
type SimplePizzaFactory struct {
	PizzaIngredientFacotry
}

//NewSimplePizzaFactory return a new NewSimplePizzaFactory
func NewSimplePizzaFactory(facotry PizzaIngredientFacotry) SimplePizzaFactory {
	return SimplePizzaFactory{facotry}
}

//Createpizza create pizza
func (facotry SimplePizzaFactory) Createpizza(pizzatype string) Pizza {
	if pizzatype == "cheese" {
		return CheesePizza{
			ingredientfactory: facotry.PizzaIngredientFacotry,
			name:              "cheese",
		}
	}
	return CheesePizza{
		ingredientfactory: facotry.PizzaIngredientFacotry,
		name:              "defaultcheese",
	}
}

// BeijingPizzaFactory represent the beijing pizza factory
type BeijingPizzaFactory struct {
	PizzaIngredientFacotry
}

// NewBeijingPizzaFactory start beijingpizza factory
func NewBeijingPizzaFactory(facotry PizzaIngredientFacotry) BeijingPizzaFactory {
	return BeijingPizzaFactory{facotry}
}

// Createpizza create pizza
func (factory BeijingPizzaFactory) Createpizza(pizzatype string) Pizza {
	if pizzatype == "beijingcheese" {
		return BeijingCheesePizza{
			ingredientfactory: factory,
			name:              "beijingcheese",
		}
	}
	return BeijingCheesePizza{
		ingredientfactory: factory,
		name:              "defaultbeijingcheese",
	}
}
