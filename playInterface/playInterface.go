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
	//类型断言  如果传进来的interface是int的话 就能成功提取出来 如果不是 ok就是false
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
	//类型选择
	switch v := i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case Describer:
		//可以对接口进行判断
		v.Describe()
	default:
		fmt.Println("unknow type")
	}

	fmt.Println()


	var d1 Describer
	p1 := Person{"yuzhipeng",12,}
	d1 = p1
	d1.Describe()

	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()
	fmt.Println()

	var d2 Describer
	a1 := Address{"China","HeBei"}
	//这里的赋值有问题 不能赋给值类型的Address 值类型的Address没有实现Describer接口   但是引用类型的*Address实现了这个方法
	//d2 = a1
	d2 = &a1
	d2.Describe()
}











type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	FirstName string
	LastName string
	BasicPay int
	Pf int
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.FirstName, e.LastName, (e.BasicPay + e.Pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.TotalLeaves - e.LeavesTaken
}


//接口也是可以嵌套的
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}






