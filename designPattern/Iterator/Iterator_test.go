package Iterator_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/yuzp1996/StudyGolang/designPattern/Iterator"
)

var _ = Describe("Iterator", func() {
	breakfastmeun := Iterator.NewBreakfastMeun()
	lunchmeun := Iterator.NewLunchMeun()
	waitress := Iterator.NewWaitress("Marry")
	waitress.SetTaskMeun(breakfastmeun,lunchmeun)
	waitress.PrintMeuns()
})
