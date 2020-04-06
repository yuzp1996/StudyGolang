package strategy

import "fmt"

type MallardDuck struct {
	Duck
	DisplayName string
	Habbit string
}

func (mallardduck *MallardDuck)setflybehavior(behavior FlyBehavior){
	mallardduck.flybehavior = behavior

}


type RedHeadDuck struct {
	Duck
	DisplayName string
	Habbit string
}

func (readheadduck *RedHeadDuck)setflybehavior(behavior FlyBehavior){
	readheadduck.flybehavior = behavior

}


func Startegy(){

	ParentDuck := NewFatherDuck()
	SonDuck := MallardDuck{
		ParentDuck,
		"mallardDuck",
		"basketball",
	}
	SonDuck.Fly()
	SonDuck.setflybehavior(FlyWithAirPlane{})
	SonDuck.Fly()

	fmt.Println("RedHeadDuck is comming")

	SonDuck1 := RedHeadDuck{
		ParentDuck,
		"redheadDuck",
		"football",
	}
	SonDuck1.Fly()
	SonDuck1.setflybehavior(FlyWithAirPlane{})
	SonDuck1.Fly()
}

