package firstfunction

import (
	"fmt"
)

func Mainfirstfunc()  {
	//函数赋值给变量 没有括号偶
	a := func(){
		fmt.Println("I am firstfunc")
	}
	a()

	//可以立即调用
	func(name string){fmt.Println("Immediately call", name)}("zpyuname")


	typefunction()

	//函数作为参数传递给
	var hiherordervar = func(a, b int)int{
		return a+b
	}
	//首先确定传入什么样的函数 然后再去写函数
	hiher_order(1,2,hiherordervar)


	// 函数作为返回值
	f := returnfunc()
	fmt.Printf("f is %d\n",f(4,2))
	}


//自定义函数类型
type add func(a,b int)int

func typefunction(){
	//实现上面那个函数类型
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