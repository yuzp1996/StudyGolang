package channel

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("write %d to chan ch \n", i)
		ch <- i
	}
	close(ch)
}

func BufferChannel() {
	ch := make(chan int, 5)
	go write(ch)
	time.Sleep(time.Millisecond * 2000)
	for v := range ch {
		fmt.Printf("read from ch and value is %d\n", v)
		time.Sleep(1 * time.Second)
	}
	//for{   // I can get nil(0) from channel ch 从一个关闭的channel中读取的总是零值 如果锁没有关闭 就会有锁的错误
	//	//	testv, _ := <-ch
	//	//	fmt.Printf("v is %v",testv)
	//	//}

}

func LenandCap() {
	ch := make(chan string, 2)
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
	ch <- "yuzhipeng"
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
	ch <- "xuyahui"
	fmt.Println(<-ch)
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
}
