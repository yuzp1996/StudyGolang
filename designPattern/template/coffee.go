package template

import "fmt"

type Coffee struct {
	Name string
}

func NewCoffee(name string) CaffeineBeverage {
	return *NewCaffeineBeverage(Coffee{Name: name})
}

func (coffee Coffee) AddCodiments() {
	fmt.Println("add codiments for coffee ...")
}
func (coffee Coffee) Brew() {
	fmt.Println("brew beverage for coffee...")
}
