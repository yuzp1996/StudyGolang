package decorator

type CoffeeWithSalt struct {
	Coffee
}

func (saltcoffee CoffeeWithSalt) cost() int {
	return saltcoffee.Coffee.cost() + 1
}

func (saltcoffee CoffeeWithSalt) getDescription() string {
	return saltcoffee.Coffee.getDescription() + " with Salt"
}

func AddSalt(coffee Coffee) Coffee {
	return CoffeeWithSalt{
		Coffee: coffee,
	}
}

type CoffeeWithSugger struct {
	Coffee
}

func (suggercoffee CoffeeWithSugger) cost() int {
	return suggercoffee.Coffee.cost() + 2
}

func (suggercoffee CoffeeWithSugger) getDescription() string {
	return suggercoffee.Coffee.getDescription() + " with Sugger"
}

func AddSugger(coffee Coffee) Coffee {
	return CoffeeWithSugger{
		Coffee: coffee,
	}
}
