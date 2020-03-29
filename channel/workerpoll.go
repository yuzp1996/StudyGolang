package channel

//import (
//	"fmt"
//	"sync"
//	"time"
//	"math/rand"
//)
//
//func process(i int ,wg *sync.WaitGroup){
//	fmt.Println("start goroutinue", i)
//	time.Sleep(2 * time.Second)
//	fmt.Println("end gogroutinue ", i)
//	wg.Done()   //2 deADD
//}
//
//func Mainprocess(){
//	num := 3
//	var wg sync.WaitGroup
//	for i := 0; i < num; i++{
//		wg.Add(1)  // 1 ADD
//		go process(i, &wg)
//	}
//
//	wg.Wait()  //3 Wait
//	fmt.Println("Main process end")
//}
//
//
//
//
//// true workpoll
//
//
//type Job struct{
//	id int
//	randomno int
//}
//
//type Result struct {
//	job Job
//	sumofdigits int
//}
//
//var jobs = make(chan Job, 10)
//var results = make(chan Result, 10)
//
//
//func digits(number int)int{
//	sum := 0
//	no := number
//	for no != 0{
//		digst := no%10
//		sum += digst
//		no /= 10
//	}
//	time.Sleep(time.Second * 2)
//	return sum
//}
//
//func worker(wg *sync.WaitGroup){
//	for job := range jobs{
//       output := Result{job, digits(job.randomno)}
//       results <- output
//	}
//	wg.Done()
//}
//
//func createWorkerPool(noOfWorker int){
//	var wg sync.WaitGroup // 相当于是一个锁  会等待所有worker执行结束
//	for i := 0; i < noOfWorker; i++{
//		wg.Add(1)
//		//启动相应数量的worker 并且等待他们完成
//		go worker(&wg)
//	}
//	wg.Wait()
//	close(results)
//}
//
//
//func allocate(noOfJobs int){
//	//发送任务  往jobs缓冲信道扔任务
//	for i := 0; i< noOfJobs; i++{
//		randomno := rand.Intn(999)
//		jobs <- Job{i, randomno}
//	}
//	close(jobs)
//}
//
//func result(done chan bool){
//	for result := range results{
//		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
//	}
//	done <- true
//}
//
//func MainWorkPool(){
//	startTime := time.Now()
//	noOfJobs := 100
//	//分配任务 往jobs中
//	go allocate(noOfJobs)
//	done := make(chan bool)
//
//	//读取任务结果 从result中
//	go result(done)
//	noOfWorker := 10
//	//从jobs中拿数据  并且处理完塞到result中  //启动相应数量的worker 并且等待他们完成
//	createWorkerPool(noOfWorker)
//	<-done
//	endTime := time.Now()
//	diff := endTime.Sub(startTime)
//	fmt.Println("total time is ",diff.Seconds()," seconds")
//}
