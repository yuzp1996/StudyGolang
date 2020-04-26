package template

import "fmt"

type Tea struct {
	CaffeineBeverage
	Name string
}

func NewTea(name string, base CaffeineBeverage) Tea {
	return Tea{base, name}
}

func (coffee Tea) AddCodiments() {
	fmt.Println("add codiments for tea ...")
}

func (coffee Tea) Brew() {
	fmt.Println("brew beverage for tea ...")
}

func (coffee Tea) BoilWater() {
	fmt.Println("boil water  for tea ...")
}
