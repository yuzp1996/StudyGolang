package template

import "fmt"

type Drink interface {
	Brew()
	AddCodiments()
}

type DrinkProgress interface {
	PrepareRecipe()
}

type CaffeineBeverage struct {
	Drink
}

func NewCaffeineBeverage(drink Drink) *CaffeineBeverage {
	return &CaffeineBeverage{Drink: drink}
}

func (beverage CaffeineBeverage) PrepareRecipe() {
	beverage.BoilWater()
	beverage.Brew()
	beverage.Drink.AddCodiments()
	fmt.Printf("CaffeineBeverage is ready, enjoy yourself...\n\n\n")
}

func (beverate CaffeineBeverage) BoilWater() {
	fmt.Println("water boiling ...")
}
func (beverate CaffeineBeverage) Brew() {
	fmt.Println("brew normal ...")
}
