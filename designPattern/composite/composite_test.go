package composite_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/yuzp1996/StudyGolang/designPattern/composite"
)

var _ = Describe("Composite", func() {
	It("simulate Waitress", func() {
		basecomponent := composite.NewMeunComponent("nothing","nothing",0)

		mainmeun := composite.NewMeun("MainMeun","MainMeun will contains all meuns",*basecomponent)


		gouzi := composite.NewMeunItem("果子","早餐果子",2,*basecomponent)
		youtiaoi := composite.NewMeunItem("油条","好吃的油条",1,*basecomponent)
		breakfastmeun := composite.NewMeun("breakfastmeun","It is breakfastmeun",*basecomponent)
		breakfastmeun.Add(gouzi)
		breakfastmeun.Add(youtiaoi)

		rice := composite.NewMeunItem("米饭","好吃的米饭",2,*basecomponent)
		noodles := composite.NewMeunItem("面条","好吃的面条",1,*basecomponent)
		lunchmeun := composite.NewMeun("lunchmeun","It is lunchmeun",*basecomponent)
		lunchmeun.Add(rice)
		lunchmeun.Add(noodles)

		desertmenu := composite.NewMeun("desertmeun","desertmeun for lunchmeun", *basecomponent)
		cookies := composite.NewMeunItem("cookies","delicous cookies",5,*basecomponent)
		desertmenu.Add(cookies)
		lunchmeun.Add(desertmenu)



		broccoli := composite.NewMeunItem("西兰花","好吃的西兰花",2,*basecomponent)
		chicken_breast := composite.NewMeunItem("鸡胸肉","好吃的鸡胸肉",1,*basecomponent)
		dinermeun := composite.NewMeun("dinermeun","It is dinermeun",*basecomponent)
		dinermeun.Add(broccoli)
		dinermeun.Add(chicken_breast)

		mainmeun.Add(breakfastmeun)
		mainmeun.Add(lunchmeun)
		mainmeun.Add(dinermeun)


		Waitress := composite.NewWaitress("Marry",[]composite.MeunComponentInterface{})
		Waitress.AddMeun(mainmeun)
		Waitress.PrintMeun()

	})
})
