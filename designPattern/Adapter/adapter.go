package Adapter

import "fmt"

type Adapter struct {
	TurkeyType
}

func NewAdapter(turkey TurkeyType) Adapter {
	return Adapter{turkey}
}

func (adapter Adapter) Fly() {
	fmt.Println("I am a turkey but Adapter is working, I will fly like duck")
	adapter.TurkeyType.fly()
}

func (adapter Adapter) Quack() {
	fmt.Println("I am a turkey but Adapter is working, I will quack like duck")
	adapter.TurkeyType.gobble()
}
