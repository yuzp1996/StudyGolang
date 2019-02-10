package playpanic

import (
	"fmt"
	"runtime/debug"
)

func printFullName(firstname *string, lastname *string){
	defer recoverName()
	defer fmt.Println("call in print FullName")
	if firstname == nil{
		panic("firstname is nil")
	}
	if lastname == nil{
		panic("lastname is nil")
	}
	//go panic("zpyu crash here")

	fmt.Println("name is ",*firstname,*lastname)
}


func Mainplaypanic(){
	defer fmt.Println("call in Mainplaypanic")
	firatname := "yu"
	lastname := "zhipeng"
	printFullName(&firatname,&lastname)
	fmt.Println("1Normal print")
	runtimeerror()
	fmt.Println("Normal print")
}


func runtimeerror(){
	defer recoverName();
	a := []int{1,2,3}
	fmt.Printf("a[3] is %d", a[3])
}

func recoverName(){
	if r := recover(); r != nil{
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}