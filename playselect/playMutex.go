package playselect

import (
	"fmt"
	"sync"
)


var a = 0

func increment(wg *sync.WaitGroup, mu *sync.Mutex){
	//这个就是个锁 开了那么多协程 但是在这里锁上了
	mu.Lock()
	a += 1
	mu.Unlock()
	wg.Done()
}

func MainMutex(){
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i:=0; i<100;i++{
		wg.Add(1)
		go increment(&wg, &mu)
	}
    wg.Wait()
	fmt.Printf("a is %v",a)
}




func Mainselectchannel(){
	var wg sync.WaitGroup
	ch := make(chan bool, 1) //channel必须是这种形式的  必须是make来声明
	for i:=0;i<100;i++{
		wg.Add(1)
		go chanselect(&wg, ch)
	}
	wg.Wait()
	fmt.Printf("a is %d", a)
}

func chanselect(wg *sync.WaitGroup,ch chan bool){
	ch <- true
	a += 1
	<- ch
	wg.Done()
}
