package playselect

import (
	"fmt"
	"time"
)

func Server1(output1 chan string) {
	time.Sleep(time.Second * 3)
	output1 <- "Channel1 is here"
}

func Server2(output2 chan string) {
	time.Sleep(time.Second * 2)
	output2 <- "Channel2 is here"
}

func MainSelect() {
	output1 := make(chan string)
	output2 := make(chan string)

	go Server1(output1)
	go Server2(output2)
	select {
	case s1 := <-output1:
		fmt.Printf("s1 is %s\n", s1)
	case s2 := <-output2:
		fmt.Printf("s2 is %s\n", s2)
		//default:
		//	fmt.Printf("no server can response in time")
	}
}

func sendsignal(signal chan string) {
	sendstring := "successful signal"
	time.Sleep(10 * time.Second)
	signal <- sendstring

}

func ForandSelect() {
	signal := make(chan string)
	go sendsignal(signal)
	for {
		select {
		case result := <-signal:
			fmt.Println(result)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("watting")

		}
	}
}
