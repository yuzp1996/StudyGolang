package factory

//Simulate simulate the pizzastore
func Simulate() {
	// Now different store use the same order system but use the different system to make pizza
	pizzastore := NewPizzaStore(NewSimplePizzaFactory())

	defaultpizza := pizzastore.orderPizza("whaterver")
	defaultpizza.prepare()
	defaultpizza.bake()

	cheesepizza := pizzastore.orderPizza("cheese")
	cheesepizza.prepare()
	cheesepizza.bake()

	beijingpizzastore := NewBeijingStore(NewBeijingPizzaFactory())

	defaultbeijingpizza := beijingpizzastore.orderPizza("whaterver")
	defaultbeijingpizza.prepare()
	defaultbeijingpizza.bake()

	cheesebeijingpizza := beijingpizzastore.orderPizza("beijingcheese")
	cheesebeijingpizza.prepare()
	cheesebeijingpizza.bake()

}
