package firstfunction

import "fmt"

func Closure(){
	var a = 12
	func(){
		fmt.Printf("a is %d\n", a)
	}()
}



type Student struct{
	Name string
	Grade int
}

func SelectStudent(){
	s1 := Student{
		"ZPYU",
		66,
	}
	s2 := Student{
		"YAHUI",
		64,
	}
	studentparam := []Student{s1,s2}
	//就是把函数传过去 然后在那边进行计算 换了个地方而已
	result := filter(studentparam,func(s Student)bool{
		return s.Name=="ZPYU"
	})
	fmt.Printf("result is %v \n", result)
}

func filter(student []Student, f func(Student)bool)[]Student{
	var result []Student
	for _,v := range(student){
		if f(v){
			result = append(result,v)
		}
	}
	return result
}