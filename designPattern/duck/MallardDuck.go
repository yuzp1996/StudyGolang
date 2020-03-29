package duck

type MallardDuck struct {
	Duck
	DisplayName string
	Habbit string
}

func (mallardduck *MallardDuck)setflybehavior(behavior FlyBehavior){
	mallardduck.flybehavior = behavior

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
}

