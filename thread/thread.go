package thread

import (
	"fmt"
	"time"
)


func SayHello(){
    fmt.Println("Hello world")
}

func MainThread(){
	go SayHello()
	fmt.Println("Main thread working")
	time.Sleep(10 * time.Second)
}

func gorountint(){
	for i := 1; i <= 5; i++{
		fmt.Println(i)
		time.Sleep(250 * time.Millisecond)
	}
}
func gorountstring(){
	for i := 'a'; i < 'e'; i++{
		fmt.Printf("%c\n",i)
		time.Sleep(400 * time.Millisecond)
	}
}

func Maingorount(){
	go gorountint()
    go gorountstring()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main thread")
}




