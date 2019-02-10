package channel

import (
	"fmt"
	"time"
)



func ChanTest(){
	var ChanOne chan int
	if ChanOne == nil{
		fmt.Printf("ChanOne is nill\n")
    ChanOne = make(chan(int))
    fmt.Printf("ChanOne Type is %T\n", ChanOne)
	}
}

func Channel(done chan bool){
	fmt.Println("Channel is ready to do in 4s")
	//time.Sleep(time.Millisecond*4000)
    done <- false
}

func MainChannel(){
	done := make(chan bool)
	fmt.Println("MainChannel ready to call Channel")
	go Channel(done)
	<- done
	fmt.Println("MainChannel have called Channel")

}






func ChanCount(){
	chanadd := make(chan int)
	dechanadd := make(chan int)
	number := 1000
	go Add(number,chanadd)
	go deAdd(number,dechanadd)
	add, deadd := <-chanadd , <-dechanadd
	fmt.Printf("add is %d, deadd is %d\n", add, deadd)
	sendonlychan := make(chan int)
	go sendData(sendonlychan) //go的关键字可是不能丢呀 协程协程呀 跟信道一起使用
	num := <-sendonlychan
	fmt.Println("num is ",num)
	}

func Add(number int, chanadd chan int){
	fmt.Println("wait for add It will compleat in 4s")
	//time.Sleep(4000 * time.Millisecond)
	chanadd <- number + number

}

func deAdd(number int, dechanadd chan int){
	dechanadd <- number - 1
}

func sendData(chandata chan<- int){
	chandata <- 10
}




func producter(forchan chan int){
	for i := 0; i < 10; i++{
		forchan <- i
		fmt.Printf("Projecter :")
		time.Sleep(1000 * time.Millisecond)
	}
	close(forchan)
}


func Chanfor(){
	forchan := make(chan int)
	go producter(forchan) // 总是忘记协程的事情呢
	for{
		v, ok := <-forchan
		fmt.Printf("ok is %v and ", ok)
		if ok == false {
			break
		}
		fmt.Printf("v is %d \n", v)
	}
}

func ChanRange(){
	forchan := make(chan int)
	go producter(forchan) // 总是忘记协程的事情呢
    for v:= range forchan{
		fmt.Printf(" v is %d \n", v)
	}
}