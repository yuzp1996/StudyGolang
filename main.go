package main

import (
	"StudyGolang/Struct"
	"StudyGolang/array"
	"StudyGolang/channel"
	"StudyGolang/firstfunction"
	"StudyGolang/function"
	"StudyGolang/inheritance"
	"StudyGolang/maptest"
	"StudyGolang/method"
	"StudyGolang/playInterface"
	"StudyGolang/playdefer"
	"StudyGolang/playerror"
	"StudyGolang/playpanic"
	"StudyGolang/playselect"
	"StudyGolang/playstring"
	"StudyGolang/pointer"
	"StudyGolang/polymorphism"
	"StudyGolang/readfromfile"
	"StudyGolang/rectangle"
	"StudyGolang/reflectplay"
	"StudyGolang/sort"
	"StudyGolang/structasclassemployee"
	"StudyGolang/thread"
	"flag"
	"fmt"
)

var rectLen, rectWidth float64 = 1, 2

const (
	Separator = "--------------------"
)

func NewSection(name string) {
	fmt.Printf("%v \n New Section is %s \n%v \n", Separator, name, Separator)
}


func init() {
	fmt.Println("main init")
}

func main() {
	sectionflag := flag.String("Section","Function","which section will run")
	flag.Parse()
	fmt.Println(*sectionflag)
	switch *sectionflag {
	case "Function":{
		NewSection("Function")
		fmt.Printf("area of rectangle %.3f\n", rectangle.Area(rectLen, rectWidth))
		fmt.Printf("Diagonal of rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
	}
	case "IfElseSwitchFor":{
		NewSection("IfElseSwitchFor")
		if true{
			fmt.Println("gogo")
		}else if false{
			fmt.Println("false")
		}else{
			fmt.Println("good" )
		}

		for i:=0; i<10; i++{
			if i == 2{
				continue
			}else{
				fmt.Println(i)
			}
		}

		switch figer := 1; figer {
		case 1,3,5:
			fmt.Println("figer is 1")
		case 2:
			fmt.Println("figer is  2")
		default:
			fmt.Println("default")
		}

		num := 70
		switch {
		case num < 100:
			fmt.Println("lower than 100")
		case num < 79:
			fmt.Println("lower than 79")
			fallthrough
		case num < 71:
			fmt.Println("lower than 71")
		}
	}
	case "ArraySlice":{
		NewSection("ArraySlice")
		array.Printarrays([2][3]int{{1, 2, 3}, {4, 5, 6}})
		array.Testarray("")

		testArray := [5]string{"cat", "dog"}
		array.Changearray(testArray)
		fmt.Println("testArray is ", testArray)

		slicetest := []int{1, 2, 3, 4}
		fmt.Printf("slice can work and now it is %v \n", slicetest)
		array.SlicewillChange(slicetest)
		fmt.Printf("slice can work and last it is %v\n", slicetest)
	}
	case "Multipama":{
		NewSection("Multipama")
		function.Multipama(1, 2, 2, 2, 2, 3, 3, 1)
		nums := []int{1, 2, 3, 4, 5}
		//语法糖  可以传入切片
		function.Multipama(1, nums...)

		names := []string{"yuzhipeng", "like", "you"}
		function.Change(names...)
		fmt.Println(names)
	}
	case "MapAndString":{
		NewSection("MAP")
		maptest1.Newmap()

		NewSection("String")
		playstring.PrintString("abcdefg")
		fmt.Println(playstring.ChangeString([]rune("string"), 0))

		//遍历字符串
		for i,v := range([]rune("zpyu")){
			fmt.Printf("i is %v and v is %c\n ", i,v)
		}

		playstring.CheckforJira("aaaa")

	}
	case "Pointor":{
		NewSection("Pointer")
		pointer.Pointertest()
		val := 1
		pointer.ShowPointerChange(&val)
		fmt.Printf("val now is %v\n", val)


		pointerarray := [3]int{5, 4, 3}
		fmt.Printf("before chagne pointerarray is %v \n", pointerarray)
		pointer.ChangeArray(&pointerarray)
		fmt.Printf("after changearray now pointerarray is %v\n", pointerarray)

		//这种使用切片的方式是更加推荐的
		pointer.ChangewithSlice(pointerarray[:])
		fmt.Printf("after slice chagne is %v\n", pointerarray)

		test := []int{1,2,3,4,5} // golang 的也是最后一个不显示
		fmt.Printf("test [0:1] is %v and test [1:] is %v", test[0:1], test[1:5])
	}
	case "Method":{
		NewSection("Method")
		method.Method()
		method.ChangePersonInfo()
		method.UseMyInt()
	}
	case "Interface":{
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
		//Salary继承了那个接口 所以就变成了这个接口类型 所以这里会报错
		Slary := playInterface.Salary{
			Name:"zpyu",
		}
		playInterface.Calculate([]playInterface.TotalSalary{trainee,worker,some,Slary})
		fmt.Println()

		playInterface.WhoAmI("good")
		playInterface.WhoAmI(123)
		//匿名结构体 并且进行了赋值
		playInterface.WhoAmI(struct {
			name string
		}{
			name:"zpyu",
		})
		fmt.Println()


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
		//猜猜看这里为什么也是unknow type ?
		playInterface.FindType(playInterface.Address{
			"China",
			"hebei",
		})
		fmt.Println()


		e := playInterface.Employee {
			FirstName: "Naveen",
			LastName: "Ramanathan",
			BasicPay: 5000,
			Pf: 200,
			TotalLeaves: 30,
			LeavesTaken: 5,
		}
		//interface是在这里进行了使用 声明成某种接口 然后后面再去使用
		//其实不在这里进行声明也可以使用  只是为了说明一下
		var s playInterface.SalaryCalculator = e
		s.DisplaySalary()
		var l playInterface.LeaveCalculator = e
		fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
		//接口是可以嵌套的
		var Ee playInterface.EmployeeOperations = e
		Ee.DisplaySalary()
	}
	case "Thread":{
		NewSection("Thread")
		thread.MainThread()
		thread.Maingorount()
	}
	case "Channel":{
		NewSection("Channel")
		//channel.ChanTest()
		//channel.ChannelControlleGorou()
		//channel.ChanCount()
		//channel.ChanRange()
		//channel.Chanfor()
		//channel.CalNum()
		fmt.Println("Buffered Channels")
		//channel.BufferChannel()
		//channel.LenandCap()
		//channel.Mainprocess()
		//channel.MainWorkPool()
		channel.MainFunc()
	}
	case "Select":{
		NewSection("Select")
		//playselect.ForandSelect()
		//playselect.MainSelect()
		//playselect.MainMutex()
		playselect.Mainselectchannel()
	}
	case "StructAsClass":{
		NewSection("StructAsClass")
		e := structasclassemployee.New("Yu", "zhipeng", 100,5,)
		e.LeavesRemaining()
	}
	case "Inheritance":{
		NewSection("Inheritance")
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
	}
	case "Polymorphism":{
		NewSection("Polymorphism")
		polymorphism.MainPoly()
	}
	case "Defer":{
		NewSection("Defer")
		playdefer.MainDefer()
	}
	case "PlayError":{
		NewSection("PlayError")
		//playerror.Openfile()
		//playerror.OpenWeb()

		//simpleerr := errors.New("Errors: simple err")
		//fmt.Println(simpleerr)

		//err := playerror.ErrorNew(" yu zhipeng create an Error")
		//fmt.Println(err)

		playerror.MainCoustmerError()
	}
	case "Panic":{
		NewSection("Panic")
		playpanic.Mainplaypanic()
	}
	case "FirstClassFunction":{
		NewSection("firstfunction")
		firstfunction.Mainfirstfunc()
		firstfunction.Closure()
		firstfunction.SelectStudent()
	}
	case "Reflect":{
		NewSection("NewReflect")
		reflectplay.Mainreflect()
		reflectplay.TryReflect()
	}
	case "ReadFile":{
		NewSection("FileHandling")
		readfromfile.Readfile()
		readfromfile.Readfrombuffer()
	}
	case "QuickSort":{
		NewSection("FileHandling")
		quicksort.QuickSort()
	}
	case "Struct":{
		NewSection("Strut")
		Struct.Deduplica()
	}
	}






















}

