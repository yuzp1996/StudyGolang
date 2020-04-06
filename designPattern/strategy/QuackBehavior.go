package strategy

import "fmt"

type QuackBehavior interface {
	quack()
}

type QuackLikeCar struct {}

func (QuackLikeCar)quack() {
	fmt.Println("I quack like car")
}

type QuackLikeMan struct {
}
func (QuackLikeMan)quack() {
	fmt.Println("I quack like man")
}