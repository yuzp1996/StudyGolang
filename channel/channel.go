package channel

import (
	"fmt"
	"time"
)

func ChanTest() {
	var ChanOne chan int
	if ChanOne == nil {
		fmt.Printf("ChanOne is nill\n")
		//信道和map与切片一致  都需要make来初始化
		ChanOne = make(chan (int))
		fmt.Printf("ChanOne Type is %T\n", ChanOne)
	}
}

func Channel(done chan bool) {
	fmt.Println("Channel job will done in 4s")
	time.Sleep(time.Millisecond * 4000)
	done <- true
}

func ChannelControlleGorou() {
	done := make(chan bool)
	fmt.Println("MainChannel ready to call Channel")
	go Channel(done)
	result := <-done
	fmt.Printf("MainChannel have called Channel and result is %t\n", result)
}

func ChanCount() {
	//简短声明chan是最好的方法
	chanadd := make(chan int)
	dechanadd := make(chan int)
	number := 1000
	go Add(number, chanadd)
	go deAdd(number, dechanadd)

	//这里其实就是阻塞的关键
	add, deadd := <-chanadd, <-dechanadd
	fmt.Printf("add is %d, deadd is %d\n", add, deadd)

	//单向信道和双向信道互相转换 这里  在下面的协程中是唯送信道 主协程中是双向信道
	sendonlychan := make(chan int)
	go sendData(sendonlychan) //go的关键字可是不能丢呀 协程协程呀 跟信道一起使用
	num := <-sendonlychan
	fmt.Println("num is ", num)

}

func Add(number int, chanadd chan int) {
	fmt.Println("wait for add It will compleat in 4s")
	time.Sleep(4000 * time.Millisecond)
	chanadd <- number + number

}

func deAdd(number int, dechanadd chan int) {
	dechanadd <- number - 1
}

func sendData(chandata chan<- int) {
	fmt.Printf("calculate will take 3 second")
	time.Sleep(3 * time.Second)
	chandata <- 10
}

//信道锁 是如果往信道里塞了数据 没人读 会锁， 如果没数据 还生去读 那就也会锁
//如果是有人读 信道关闭 但是没有数据 ，那读取的就是零值

//数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
//当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
//v, ok := <- ch
//上面的语句里，如果成功接收信道所发送的数据，那么 ok 等于 true。而如果 ok 等于 false，说明我们试图读取一个关闭的通道。

//两种读取chan信道的方式
//for 读取信道
func Chanfor() {
	forchan := make(chan int)
	go producter(forchan) // 总是忘记协程的事情呢
	for {
		v, ok := <-forchan
		fmt.Printf("ok is %v and ", ok)
		if ok == false {
			break
		}
		fmt.Printf("v is %d \n", v)
	}
}

//for range 循环用于在一个信道关闭之前，从信道接收数据。
func ChanRange() {
	forchan := make(chan int)
	go producter(forchan) // 总是忘记协程的事情呢
	for v := range forchan {
		fmt.Printf(" v is %d \n", v)
	}
}

func producter(forchan chan int) {
	for i := 0; i < 10; i++ {
		forchan <- i
		fmt.Printf("Projecter :")
		time.Sleep(1000 * time.Millisecond)
	}
	//记住关闭信道
	close(forchan)
}

//基本就是靠信道来传输值 然后用for range进行值的读取
func CalNum() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares, cubes)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
