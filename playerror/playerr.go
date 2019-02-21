package playerror

import (
	"fmt"
	"net"
	"os"
)

func Openfile(){
	file, err := os.Open("./zpyutest.txt")
	//在这获取他的类型 多好
	if err, ok := err.(*os.PathError); ok{
		fmt.Println(err.Path)
		return
	}
	fmt.Println("Open succeeful %v", file.Name())
}

func OpenWeb(){
	addr, err := net.LookupAddr("www.baidu.com")
	//可以去查看源码 然后看看对应的是什么错误 然后进行捕捉  查看有什么方法
	//在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
	//s := i.(int) //get the underlying int value from i
	//v, ok := i.(T)
	//这里是类型断言
	if err, ok := err.(*net.DNSError); ok{
		if err.Temporary(){
			fmt.Println("tempory  error")
		}else if err.Timeout(){
			fmt.Println("timeout")
		}else{
			fmt.Println("generic error", err)
		}
		return
	}
	fmt.Println(addr)
}