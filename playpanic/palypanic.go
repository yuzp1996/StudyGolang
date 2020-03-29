package playpanic

import (
	"fmt"
	"runtime/debug"
)

//panic之后 延迟函数会继续调用
func Mainplaypanic() {
	defer fmt.Println("call in Mainplaypanic")
	firatname := "yu"
	lastname := "zhipeng"
	printFullName(&firatname, &lastname)
	fmt.Println("1Normal print")
	//runtimeerror()
	//fmt.Println("Normal print")
}

func printFullName(firstname *string, lastname *string) {
	defer recoverName()
	defer fmt.Println("call in print FullName")
	if firstname == nil {
		panic("firstname is nil")
	}
	if lastname == nil {
		panic("lastname is nil")
	}
	//go panic("zpyu crash here")

	fmt.Println("name is ", *firstname, *lastname)
}

func runtimeerror() {
	//只有在延迟函数的内部，调用 recover 才有用。
	defer recoverName()
	a := []int{1, 2, 3}
	fmt.Printf("a[3] is %d", a[3])
}

//都在延迟函数调用我
func recoverName() {
	//调用 recover() 来重新获得 panic 协程的控制
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}
