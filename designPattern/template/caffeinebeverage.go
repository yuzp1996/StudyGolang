package template

import "fmt"

type Drink interface {
	Brew()
	AddCodiments()
}

type BaseDrink interface {
	BoilWater()
	Drink
}

type CaffeineBeverage struct {
}

func NewCaffeineBeverage() *CaffeineBeverage {
	return &CaffeineBeverage{}
}

func (beverate CaffeineBeverage) BoilWater() {
	fmt.Println("water boiling ...")
}
func (beverate CaffeineBeverage) Brew() {
	fmt.Println("brew normal ...")
}
func (beverate CaffeineBeverage) AddCodiments() {
	fmt.Println("add codiments ...")
}
