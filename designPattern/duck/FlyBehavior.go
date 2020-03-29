package duck

import "fmt"

type FlyBehavior interface {
	fly()
}
type FlyWithoutWings struct {}

func (FlyWithoutWings)fly() {
	fmt.Println("how niubi am I, I can fly without Wings")
}

type FlyWithAirPlane struct {
}
func (FlyWithAirPlane)fly() {
	fmt.Println("Ai ,  I just  can fly with airplane")
}