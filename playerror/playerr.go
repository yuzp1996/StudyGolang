package playerror

import (
	"fmt"
	"net"
	"os"
)

func Openfile(){
	file, err := os.Open("./zpyutest.txt")
	if err, ok := err.(*os.PathError); ok{
		fmt.Println(err.Path)
		return
	}
	fmt.Println("Open succeeful %v", file.Name())
}

func OpenWeb(){
	addr, err := net.LookupAddr("baidu")
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