package factory

//PizzaIngredientFacotry  intreface for Pizza Ingredient Factory
type PizzaIngredientFacotry interface {
	createsalt() string
	createsugger() string
}

type BeijingIngredientFacotry struct {
	saltname   string
	suggername string
}

func (factory BeijingIngredientFacotry) createsalt() string {
	return factory.saltname
}

func (factory BeijingIngredientFacotry) createsugger() string {
	return factory.suggername
}

type IngredientFacotry struct {
	saltname   string
	suggername string
}

func (factory IngredientFacotry) createsalt() string {
	return factory.saltname
}

func (factory IngredientFacotry) createsugger() string {
	return factory.suggername
}
