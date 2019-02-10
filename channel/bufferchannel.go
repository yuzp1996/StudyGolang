package channel

import (
	"fmt"
	"time"
)

func write(ch chan int){
	for i:=0; i<10; i++{
		fmt.Printf("write %d to chan ch \n", i)
		ch <- i
	}
	close(ch)
}


func BufferChannel(){
	ch := make(chan int,5)
	go write(ch)
	time.Sleep(time.Millisecond * 2000)
	for v := range ch{
        fmt.Printf("read from ch and value is %d\n", v)
		time.Sleep(1 * time.Second)
	}
	//for{   // I can get nil(0) from channel ch
	//	testv, _ := <-ch
	//	fmt.Printf("v is %v",testv)
	//}

}

func FullChannel(){
	ch := make(chan string,2)
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
	ch <- "yuzhipeng"
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
	ch <- "xuyahui"
	fmt.Println(<-ch)
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("len is %d, and cap is %d \n", len(ch), cap(ch))
}
