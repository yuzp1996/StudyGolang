package main

import (
	"fmt"
	"geometry/array"
	"geometry/channel"
	"geometry/firstfunction"
	"geometry/function"
	"geometry/inheritance"
	"geometry/maptest"
	"geometry/method"
	"geometry/playInterface"
	"geometry/playdefer"
	"geometry/playerror"
	"geometry/playpanic"
	"geometry/playstring"
	"geometry/pointer"
	"geometry/polymorphism"
	"geometry/readfromfile"
	"geometry/rectangle"
	"geometry/reflectplay"
	"geometry/sort"
	"geometry/structasclassemployee"
)

var rectLen, rectWidth float64 = 1, 2

const (
	Separator = "--------------------"
)

func init() {
	if rectLen < 0 {
		fmt.Println("bad it should bigger than 0")
	} else if rectWidth > 0 {
		fmt.Println("better thing happend")
	}

	fmt.Println("init main")

	if num := 10; num%2 == 0 {
		fmt.Printf("num %d is Even \n", num)
	}
	//fmt.Println("num out of if-esle is %d", num)

}
func main() {
	fmt.Println("Geometrical shape properties")
	if rectLen > 0 {
		fmt.Println("right")
	}
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))

	fmt.Println("now let's learn for")
	for i := 1; i <= 10; i++ {
		if i > 9 {
			fmt.Println("bigger rhan 9")
			break
		} else if i == 4 {
			continue
		} else {
			fmt.Printf("i is %d\n", i)
		}
	}
	var i int = 1
	for i < 10 {
		fmt.Printf("i is %d \n", i)
		i++
	}
	//for {
	//  fmt.Println("I love XYH")
	//}

	switch figer := 1; figer {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
	num := 70
	switch {
	case num < 100:
		fmt.Println("lower than 100")
		fallthrough
	case num < 79:
		fmt.Println("lower than 79")
		fallthrough
	case num < 71:
		fmt.Println("lower than 71")
	}

	fmt.Println(Separator)

	array.Printarrays([2][3]int{{1, 2, 3}, {4, 5, 6}})

	array.Testarray("Now do array\n-----------------------------\n")
	testArray := [5]string{"chinadog", "dog"}
	fmt.Printf("testArray is %v \n", testArray)
	array.Changearray(testArray)
	fmt.Println("testArray is ", testArray)

	slicetest := []int{1, 2, 3, 4}
	fmt.Printf("slice can work and now it is %v \n", slicetest)
	array.SlicewillChange(slicetest)
	fmt.Printf("slice can work and last it is %v\n", slicetest)

	function.Multipama(1, 2, 2, 2, 2, 3, 3, 1)
	function.Multipama(2, 2, 2, 2, 2, 3, 3, 1)
	nums := []int{1, 2, 3, 4, 5}
	//语法糖
	function.Multipama(1, nums...)

	names := []string{"yuzhipeng", "like", "you"}
	function.Change(names...)
	fmt.Println(names)

	NewSection("MAP")
	maptest1.Newmap()
	NewSection("String")
	playstring.PrintString("abcdefg")
	fmt.Println(playstring.ChangeString([]rune("string"), 0))

	playstring.CheckforJira("aaaa")
	NewSection("Pointer")
	pointer.Pointertest()
	val := 1

	pointer.ShowPointerChange(&val)
	fmt.Printf("val now is %v\n", val)
	pointerarray := [3]int{5, 4, 3}
	fmt.Printf("before chagne pointerarray is %v \n", pointerarray)
	pointer.ChangeArray(&pointerarray)
	fmt.Printf("after changearray now pointerarray is %v\n", pointerarray)
	pointer.ChangewithSlice(pointerarray[:])
	fmt.Printf("after slice chagne is %v\n", pointerarray)
	quicksort.QuickSort()
	test := []int{1,2,3,4,5} // golang 的也是最后一个不显示
	fmt.Printf("test [0:1] is %v and test [1:] is %v", test[0:1], test[1:5])
	NewSection("Method")
	method.Method()
	method.ChangePersonInfo()
	method.UseMyInt()

	NewSection("Interface")
	trainee := playInterface.Trainee{
		Salary:10,
	}
	worker := playInterface.Worker{
		Salary:10,
		Welfare:0,
	}
	some := playInterface.Somethingelse{
		Salary:10,
		Welfare:4,
	}
	playInterface.Calculate([]playInterface.TotalSalary{trainee,worker,some})
	playInterface.WhoAmI("good")
	playInterface.WhoAmI(123)
	playInterface.WhoAmI(struct {
		name string
	}{
		name:"zpyu",
	})
	playInterface.Assert(3)
    playInterface.FindType(12)
	playInterface.FindType("string")
	playInterface.FindType(struct{
		name string
	}{
		name: "yuzhipeng",
	})
	playInterface.FindType(playInterface.Person{
		Name:"yuzhipeng",
		Age:12,
	})


	NewSection("Thread")
	//thread.MainThread()
	//thread.Maingorount()

	NewSection("Channel")
	channel.ChanTest()
	channel.MainChannel()
	channel.ChanCount()
	//channel.Chanfor()
	//channel.ChanRange()
	//channel.BufferChannel()
	//channel.FullChannel()
	//channel.Mainprocess()
	//channel.MainWorkPool()
	//channel.MainFunc()
	//playselect.MainSelect()
	//playselect.MainMutex()
	//playselect.Mainselectchannel()


	e := structasclassemployee.New("Yu", "zhipeng", 100,5,)
	e.LeavesRemaining()

	author := inheritance.Author{
		"yu","zhipeng",24,
	}
	post := inheritance.Post{
		"goodtitle","very good content",author,
	}
	web:=inheritance.Website{
		Posts:[]inheritance.Post{post},
	}
	web.Contents()
	polymorphism.MainPoly()
	playdefer.MainDefer()

	playerror.Openfile()
	playerror.OpenWeb()
	err := playerror.ErrorNew("yu zhipeng create an Error")
	fmt.Println(err)
	playerror.MainCoustmerError()
	playpanic.Mainplaypanic()
	NewSection("firstfunction")
	firstfunction.Mainfirstfunc()
	firstfunction.Closure()
	firstfunction.SelectStudent()
	NewSection("NewReflect")
	reflectplay.Mainreflect()
	reflectplay.TryReflect()
	NewSection("FileHandling")
	//readfromfile.Readfile()
	readfromfile.Readfrombuffer()
}

func NewSection(name string) {
	fmt.Printf("%v \n New Section is %s \n%v \n", Separator, name, Separator)
}
