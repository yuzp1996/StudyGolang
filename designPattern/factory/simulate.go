package factory

import "fmt"

//Simulate simulate the pizzastore
func Simulate() {
	// Now different store use the same order system but use the different system to make pizza
	// Use different factory
	// pizzastore -> pizzafactory -> pizza
	fmt.Println("Pizza Store ")
	// 这里的factory是一个接口， 这也印证了 针对接口编程而不是针对实现编程
	pizzastore := NewPizzaStore(NewSimplePizzaFactory(IngredientFacotry{"saltname", "suggername"}))

	defaultpizza := pizzastore.orderPizza("whaterver")
	defaultpizza.prepare()
	defaultpizza.bake()

	cheesepizza := pizzastore.orderPizza("cheese")
	cheesepizza.prepare()
	cheesepizza.bake()

	fmt.Println("Bejing Pizza Store")

	beijingpizzastore := NewPizzaStore(NewBeijingPizzaFactory(BeijingIngredientFacotry{"beijingsaltname", "beijing suggername"}))

	defaultbeijingpizza := beijingpizzastore.orderPizza("whaterver")
	defaultbeijingpizza.prepare()
	defaultbeijingpizza.bake()

	cheesebeijingpizza := beijingpizzastore.orderPizza("beijingcheese")
	cheesebeijingpizza.prepare()
	cheesebeijingpizza.bake()

}
