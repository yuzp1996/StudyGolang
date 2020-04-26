package template

import "fmt"

type Waiter struct {
	Name string
}

func NewWaiter(name string) *Waiter {
	return &Waiter{Name: name}
}

func (waiter Waiter) MakeDrink(beverage BaseDrink) {
	fmt.Printf("Hello I am waiter %s, I am happy server for you \n", waiter.Name)
	beverage.BoilWater()
	beverage.Brew()
	beverage.AddCodiments()
}
