package template

import "fmt"

type Coffee struct {
	CaffeineBeverage
	Name string
}

func NewCoffee(name string, base CaffeineBeverage) Coffee {
	return Coffee{base, name}
}

func (coffee Coffee) AddCodiments() {
	fmt.Println("add codiments for coffee ...")
}
func (coffee Coffee) Brew() {
	fmt.Println("brew beverage for coffee...")
}
