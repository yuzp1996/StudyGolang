package Adapter

import "fmt"

type TurkeyType interface {
	gobble()
	fly()
}

type TurkeyInstance struct {
	Name string
}

func (turkey TurkeyInstance) fly() {
	fmt.Println("I am a turkey, I am fly")
}

func (turkey TurkeyInstance) gobble() {
	fmt.Println("I am a turkey, I am gobble")
}

func NewTurkey() TurkeyType {
	return TurkeyInstance{Name: "turkey"}
}
