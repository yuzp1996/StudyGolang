package template

import "fmt"

type Waiter struct {
	Name string
}

func NewWaiter(name string) *Waiter {
	return &Waiter{Name: name}
}

func (waiter Waiter) MakeDrink(propress DrinkProgress) {
	fmt.Printf("Hello I am waiter %s, I am happy server for you \n", waiter.Name)
	propress.PrepareRecipe()
}
