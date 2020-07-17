package Adapter

import "fmt"

type DuckAdapter struct {
	TurkeyType
}

func NewAdapter(turkey TurkeyType) DuckAdapter {
	return DuckAdapter{turkey}
}

func (adapter DuckAdapter) Fly() {
	fmt.Println("I am a turkey but DuckAdapter is working, I will fly like duck")
	adapter.TurkeyType.fly()
}

func (adapter DuckAdapter) Quack() {
	fmt.Println("I am a turkey but DuckAdapter is working, I will quack like duck")
	adapter.TurkeyType.gobble()
}
