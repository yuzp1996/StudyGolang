package duck

type Duck struct {
	flybehavior FlyBehavior
	quackBehavior QuackBehavior
	kind string
	name string
}

func (duck Duck)Fly(){
	duck.flybehavior.fly()
}

func (duck Duck)Quack(){
	duck.quackBehavior.quack()
}

func NewFatherDuck()Duck{
	return Duck{
		kind:"father",
		name:"father duck",
		flybehavior:FlyWithoutWings{},
		quackBehavior:QuackLikeMan{},
	}
}