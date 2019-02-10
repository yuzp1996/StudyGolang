package playInterface

import "fmt"

type TotalSalary interface{
	CalculateSalary() int
}

type Trainee struct {
	Salary int
}

type Worker struct {
	Salary int
	Welfare int
}

type Somethingelse struct{
	Salary int
	Welfare int
	Admin int
}

func (s Somethingelse)CalculateSalary()int{
	return s.Salary + s.Welfare + s.Admin
}

func (t Trainee)CalculateSalary() int{
	return t.Salary
}

func (w Worker)CalculateSalary()int{
	return w.Salary + 12 * w.Welfare
}

func Calculate(worker []TotalSalary){
	for _, v:= range(worker){
		fmt.Printf("Salary is %d \n", v.CalculateSalary())
	}
}


func WhoAmI(who interface{}){
	fmt.Printf(" I am %v and my type is %T\n", who, who)
}

func Assert(i interface{}){
	s, ok:= i.(int)
	if ok{
		fmt.Println(s)
	}else{
		fmt.Println("Wrong Type not int ")
	}

}

type Describer interface{
	Describe()
}

type Person struct {
    Name string
    Age int
}

func (p Person)Describe(){
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}

type Address struct {
	Country string
	Area string
}
func (a *Address)Describe(){
	fmt.Println(a.Area)
	fmt.Println(a.Country)
}

func FindType(i interface{}){
	switch v := i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case Describer:
		v.Describe()
	default:
		fmt.Println("unknow type")
	}
}









