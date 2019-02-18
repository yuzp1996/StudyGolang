package main

import (
	"StudyGolang/array"
	"StudyGolang/function"
	"StudyGolang/maptest"
	"StudyGolang/method"
	"StudyGolang/playInterface"
	"StudyGolang/playstring"
	"StudyGolang/pointer"
	"StudyGolang/rectangle"
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
	}

	}















	//
	//
	//NewSection("Thread")
	////thread.MainThread()
	////thread.Maingorount()
	//
	//NewSection("Channel")
	//channel.ChanTest()
	//channel.MainChannel()
	//channel.ChanCount()
	////channel.Chanfor()
	////channel.ChanRange()
	////channel.BufferChannel()
	////channel.FullChannel()
	////channel.Mainprocess()
	////channel.MainWorkPool()
	////channel.MainFunc()
	////playselect.MainSelect()
	////playselect.MainMutex()
	////playselect.Mainselectchannel()
	//
	//
	//e := structasclassemployee.New("Yu", "zhipeng", 100,5,)
	//e.LeavesRemaining()
	//
	//author := inheritance.Author{
	//	"yu","zhipeng",24,
	//}
	//post := inheritance.Post{
	//	"goodtitle","very good content",author,
	//}
	//web:=inheritance.Website{
	//	Posts:[]inheritance.Post{post},
	//}
	//web.Contents()
	//polymorphism.MainPoly()
	//playdefer.MainDefer()
	//
	//playerror.Openfile()
	//playerror.OpenWeb()
	//err := playerror.ErrorNew("yu zhipeng create an Error")
	//fmt.Println(err)
	//playerror.MainCoustmerError()
	//playpanic.Mainplaypanic()
	//NewSection("firstfunction")
	//firstfunction.Mainfirstfunc()
	//firstfunction.Closure()
	//firstfunction.SelectStudent()
	//NewSection("NewReflect")
	//reflectplay.Mainreflect()
	//reflectplay.TryReflect()
	//NewSection("FileHandling")
	////readfromfile.Readfile()
	////readfromfile.Readfrombuffer()
	//quicksort.QuickSort()

}

