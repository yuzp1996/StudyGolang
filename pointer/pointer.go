package pointer

import (
	"fmt"
)

func Pointertest() {
	b := 225
	var a *int = &b
	fmt.Printf("a is %v and type  is %T\n", a, a)

	var c *int
	if c == nil {
		fmt.Println("c is nil")
		c = &b
		fmt.Println("after c = &b  c is ", c)
		//解引用
		fmt.Printf("解引用:\nvalue of c is *c :%v\n", *c)
		fmt.Printf("address of c is c: %v\n", c)
	}
	fmt.Printf("b now is %v\n", b)
	*c++
	fmt.Printf("after I exec *c++ b is %v\n", b)
}



 func ShowPointerChange(val *int){
 	fmt.Printf("val is %v \n", *val)
 	*val++
 }


func ChangeArray(items *[3]int){
	(*items)[1] = 1
	items[0] = 0
}



func ChangewithSlice(sliceitem []int){
	sliceitem[0] = 100
}