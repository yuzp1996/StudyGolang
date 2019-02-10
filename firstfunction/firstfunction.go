package firstfunction

import (
	"fmt"
)

func Mainfirstfunc()  {
	a := func(){
		fmt.Println("I am firstfunc")
	}
	a()
	func(name string){fmt.Println("Immediately call", name)}("zpyuname")


	typefunction()

	// as pamar
	var hiherordervar = func(a, b int)int{
		return a+b
	}
	hiher_order(1,2,hiherordervar)


	// as return
	f := returnfunc()
	fmt.Printf("f is %d\n",f(4,2))
	}

type add func(a,b int)int

func typefunction(){
	var a add = func(a,b int)int{
		return a + b
	}
	q := a(4,6)
	fmt.Printf("q is %d\n", q)
}


func hiher_order(first, second int, a func(a,b int)int){
	fmt.Printf("a(a, b ) is %d\n", a(first,second))
}


func returnfunc()(func(a, b int)int){
	f := func(a,b int)int{
		return a+b
	}
	return f
}