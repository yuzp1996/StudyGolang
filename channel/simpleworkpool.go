package channel

import (
	"fmt"
	"sync"
	"time"
)

var job = make(chan int,2)
var result = make(chan int,2)

func take(done chan bool){
	for re := range result{
		fmt.Printf("re is %d \n", re)
	}
	//2.这里读完了 才允许下一步 结束主进程
	done <- true
}

func add(numOfJobs int){
	for i:=0; i<numOfJobs; i++{
		job <- i
	}
	//不用写入了 所以就关了吧  否则会panic
	close(job)
}

func createWorkerPool(numOfWorker int){
	var wg sync.WaitGroup
	for i:=0; i<numOfWorker; i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//1.这里结束 才算是真正的结束  所以close result  使take可以结束
	close(result)
}

func worker(wg *sync.WaitGroup){
	for jo := range job{
		dojob(jo)
	}
	wg.Done()
}

func dojob(jo int){
	time.Sleep(1*time.Second)
	result <- jo
}

func MainFunc(){
	starttime := time.Now()

	go add(10)
	done := make(chan bool)
	go take(done)
	createWorkerPool(5)
	<-done
	endtime := time.Now()
	difftime := endtime.Sub(starttime)
	fmt.Println( "total time is ",difftime.Seconds())
}
