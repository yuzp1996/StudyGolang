package template

import "fmt"

type Tea struct {
	Name string
}

func NewTea(name string) CaffeineBeverage {
	return *NewCaffeineBeverage(Tea{Name: name})
}

func (coffee Tea) AddCodiments() {
	fmt.Println("add codiments for tea ...")
}

func (coffee Tea) Brew() {
	fmt.Println("brew beverage for tea ...")
}
