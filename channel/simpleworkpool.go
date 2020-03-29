package channel

import (
	"fmt"
	"sync"
	"time"
)

var job = make(chan int, 2)
var result = make(chan int, 2)

func take(done chan bool) {
	for re := range result {
		fmt.Printf("re is %d \n", re)
	}
	//2.这里读完了 才允许下一步 结束主进程
	done <- true
}

func add(numOfJobs int) {
	for i := 0; i < numOfJobs; i++ {
		job <- i
	}
	//不用写入了 所以就关了吧  否则会panic
	close(job)
}

func createWorkerPool(numOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//1.这里结束 才算是真正的结束  所以close result  使take可以结束
	close(result)
}

func worker(wg *sync.WaitGroup) {
	for jo := range job {
		dojob(jo)
	}
	wg.Done()
}

func dojob(jo int) {
	time.Sleep(1 * time.Second)
	result <- jo
}

func MainFunc() {
	starttime := time.Now()

	// 往job中塞任务
	go add(40)

	// 从result中取结果 并且给done信息 结束阻塞
	done := make(chan bool)
	go take(done)

	//启动一定数量的goroutinue 的worker去做job【也就是把值从job拿到result中】
	createWorkerPool(10) // 这个不能跟上面的互换 因为这个是要把值塞到result中的  所以需要有个goroutinue从result中取数据 没有的话就会锁住
	//而且你还把这里锁死了  只有执行完了才能进行下一步 要是跟上面换的话肯定不行呗

	<-done
	endtime := time.Now()
	difftime := endtime.Sub(starttime)
	fmt.Println("total time is ", difftime.Seconds())
}
