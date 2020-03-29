package thread

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello world")
}

func MainThread() {
	go SayHello()
	fmt.Println("Main thread working")
	time.Sleep(5 * time.Second)
}

func Maingorount() {
	go gorountint()
	go gorountstring()
	//如果没有sleep的话 会立即退出 不会去执行gorount
	//在 Go 主协程中使用休眠，以便等待其他协程执行完毕。信道可用于在其他协程结束执行之前，阻塞 Go 主协程
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main thread")
}

func gorountint() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(250 * time.Millisecond)
	}
}
func gorountstring() {
	for i := 'a'; i < 'e'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(400 * time.Millisecond)
	}
}
