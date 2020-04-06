package decorator


type Coffee interface {
	cost()int
	getDescription()string
}



type NormalCoffee struct {
	Description string
}

func (coffee NormalCoffee)cost()int{
	return 10
}
func (coffee NormalCoffee)getDescription()string{
	return coffee.Description
}

func GetNormalCoffee()NormalCoffee  {
	return NormalCoffee{
		Description:"It's NormalCoffee ",
	}

}

type DarkCoffee struct {
	Description string
}

func (coffee DarkCoffee)cost()int{
	return 15
}
func (coffee DarkCoffee)getDescription()string{
	return coffee.Description
}

func GetDarkCoffee()DarkCoffee  {
	return DarkCoffee{
		Description:"It's DarkCoffee ",
	}

}


