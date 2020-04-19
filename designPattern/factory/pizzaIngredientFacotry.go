package factory

//PizzaIngredientFacotry  intreface for Pizza Ingredient Factory
type PizzaIngredientFacotry interface {
	getsalt() string
	getsugger() string
}

type BeijingIngredientFacotry struct {
	saltname   string
	suggername string
}

func (factory BeijingIngredientFacotry) getsalt() string {
	return factory.saltname
}

func (factory BeijingIngredientFacotry) getsugger() string {
	return factory.suggername
}

type IngredientFacotry struct {
	saltname   string
	suggername string
}

func (factory IngredientFacotry) getsalt() string {
	return factory.saltname
}

func (factory IngredientFacotry) getsugger() string {
	return factory.suggername
}
