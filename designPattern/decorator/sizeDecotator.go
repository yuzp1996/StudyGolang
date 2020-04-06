package decorator

type HalfSizeCoffee struct {
	Coffee
}

func (halfcoffee HalfSizeCoffee)cost()int  {
	return halfcoffee.Coffee.cost()/2
}

func(halfcoffee HalfSizeCoffee)getDescription()string{
	return "Half of "+halfcoffee.Coffee.getDescription()
}


func HalfSize(coffee Coffee)Coffee{
	return  HalfSizeCoffee{
		Coffee:coffee,
	}
}

