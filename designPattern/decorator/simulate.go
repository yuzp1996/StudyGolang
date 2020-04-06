package decorator

import "fmt"

func Simulate(){
	Coffee := GetNormalCoffee()
	coffee := AddSalt(Coffee)
	coffee = AddSugger(coffee)
	fmt.Printf("coffeewithsalt and sugger cost is 『%v』\n",coffee.cost())
	fmt.Printf("coffeewithsalt and sugger descritpion is 『%v』\n", coffee.getDescription())


	darkCoffee := GetDarkCoffee()
	firstcoffee := AddSugger(darkCoffee)
	firstcoffee = AddSalt(firstcoffee)
	fmt.Printf("darkcoffee and sugger and salt is 『%v』\n",firstcoffee.cost())
	fmt.Printf("darkcoffee and sugger and salt descritpion is 『%v』\n",firstcoffee.getDescription())

	halfcoffee := HalfSize(GetDarkCoffee())
	halfcoffee = AddSalt(halfcoffee)
	fmt.Printf("halfcoffee salt is 『%v』\n",halfcoffee.cost())
	fmt.Printf("halfcoffee salt descritpion is 『%v』\n",halfcoffee.getDescription())

}