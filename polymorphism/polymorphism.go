package polymorphism

import (
	"fmt"
)

type Zpyu struct {
	Sex string
	Salary int
 	Welfare int
	Source string
}

type Yhxu struct {
	Sex string
 	Salary int
	Source string
}

type Jpxi struct{
	Sex string
	Salary int
	Welfare int
	Political int
	Source string
}


type Income interface{
	calculate() int
	source() string
}

func(zpyu Zpyu)calculate()int{
	return zpyu.Salary + zpyu.Welfare
}

func (zpyu Zpyu)source()string{
	return zpyu.Source
}


func(yhxu Yhxu)calculate()int{
	return yhxu.Salary
}

func (yhxu Yhxu)source() string{
	return yhxu.Source
}


func (jpxi Jpxi)calculate()int{
	return jpxi.Salary + jpxi.Welfare + jpxi.Political
}
func (jpxi Jpxi)source()string{
	return jpxi.Source
}

//这里的参数就是接口类型的
func printSalarySourceandTotalSalary(income []Income){
	totalsalary := 0
	for _, in := range income{
		fmt.Printf("source is 『%s』And salary is %d\n", in.source(),in.calculate())
		totalsalary += in.calculate()
	}
	fmt.Printf("totalsalary is %d", totalsalary)
}

func MainPoly(){
	zpyu1 := Zpyu{
		"male",
		12000,
		1000,
		"Salary and Welfare",
	}
	yhxu := Yhxu{
		"female",
		5000,
		"Salary",
	}
	jpxi := Jpxi{
		"male",
		15000,
		5000,
		50000,
		"Salary、Welfare and Political",
	}
	//这里也做成了一个类型 你看这里面都是这个
	incomer := []Income{
		zpyu1,
		yhxu,
		jpxi,
	}
	printSalarySourceandTotalSalary(incomer)
}





