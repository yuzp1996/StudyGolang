package Iterator

import "fmt"

type Waitress struct {
	Name string
	BreakfastMeun MeunIterator
	LunchMeun MeunIterator
}

func NewWaitress(name string)Waitress{
	return Waitress{
		Name:  name,
	}
}

func(waitress *Waitress)SetTaskMeun(breakfastmeun,lunchmeun Iterator){
	waitress.BreakfastMeun = breakfastmeun.CreateIterator()
	waitress.LunchMeun = lunchmeun.CreateIterator()
}

func(waitress *Waitress)PrintMeuns(){
	fmt.Printf("Breakfast meun is below\n", )
	waitress.PrintMeun(waitress.BreakfastMeun)
	fmt.Printf("\nlunch meun is below\n", )
	waitress.PrintMeun(waitress.LunchMeun)
}

func(waitress *Waitress)PrintMeun(iterator MeunIterator) {
	for(iterator.HasValue()){
		item := iterator.GetValue()
		fmt.Printf("%s-%s\n",item.name,item.description)
	}

}


