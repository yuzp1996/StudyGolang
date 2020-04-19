package singleton_test

import (
	"StudyGolang/designPattern/singleton"
	"fmt"
	. "github.com/onsi/ginkgo"
	"time"
)

var _ = Describe("Singleton", func() {
	go singleton.GetLazyCEO()
	go singleton.GetLazyCEO()
	go singleton.GetLazyCEO()
	go singleton.GetLazyCEO()
	time.Sleep(time.Second)
	fmt.Println("Lazy Get CEO done")

	go singleton.GetHungryCEO()
	go singleton.GetHungryCEO()
	go singleton.GetHungryCEO()
	go singleton.GetHungryCEO()
	time.Sleep(time.Second)
	fmt.Println("Hungry Get CEO done")

	go singleton.GetSafeLazyCEO()
	go singleton.GetSafeLazyCEO()
	go singleton.GetSafeLazyCEO()
	go singleton.GetSafeLazyCEO()
	time.Sleep(time.Second)
	fmt.Println(" Get Safe CEO done")

	go singleton.GetOnceDoCEO()
	go singleton.GetOnceDoCEO()
	go singleton.GetOnceDoCEO()
	time.Sleep(time.Second)
	fmt.Println(" Get Safe CEO done")
})
