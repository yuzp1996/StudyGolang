package method

import (
	"fmt"
)

type square struct {
	Side int
}

type rectangle struct {
	Width int
	Longth int
}

func (s square)Area() (area int) {
	area = s.Side * s.Side
	return
}

func (r rectangle)Area()(area int){ //front is alias  after is type
	area = r.Longth * r.Width
	return
}

func Area(Recangle rectangle)(area int){
	area = Recangle.Width * Recangle.Longth
	return
}


func Method(){
	Square := square{
		Side : 5,
	}
	Rectangle := rectangle{
		Width:5,
		Longth:6,
	}
    fmt.Printf("Square area is %d \n ", Square.Area())
	fmt.Printf("Rectangle area is %d\n ", Rectangle.Area())

	fmt.Printf("Area(Rectangle) is %d\n",Area(Rectangle))
	r := &(Rectangle)
	//fmt.Printf("Area(*rectangle) is %d", Area(r))
	fmt.Printf("r.Area is %d \n",r.Area())
	fmt.Printf("Rectangel id %d \n", Rectangle.Area())
	//函数说用什么就用什么 方法没事 不管声明的是指针还是值 都能用


}

type Person struct {
	Name string
	Age int
}

func (p Person)ChangeName(name string){
	p.Name = name
	fmt.Printf("yuzhipeng's Name should be %s\n", name)
}
func (p *Person)ChangeAge(age int){
	p.Age = age
	fmt.Printf("yuzhipeng's Age should be %d\n", age)
}

func (p Person)UseValChange(name string)(person Person){
	person = p
	person.Name = name
	return
}

func ChangePersonInfo(){
	person := Person{
		Name:"yuzhipeng",
		Age:24,
	}
	fmt.Printf("Before Change Age yuzhipeng is %v \n", person)
	(&person).ChangeAge(25)
	fmt.Printf("After *Change Ager yuzhipeng is %v \n", person)

	fmt.Printf("Before Change Name yuzhipeng is %v \n", person)
	person.ChangeName("yuzhipengshuai")
	fmt.Printf("After Change Name yuzhipeng is %v \n", person)

	fmt.Printf("Before VAL Change Name yuzhipeng is %v \n", person)
	person = person.UseValChange("yuzhipenghao")
	fmt.Printf("After VAL Change Name yuzhipeng is %v \n", person)


}


type MyInt int

func (myint MyInt)Add(addint MyInt) MyInt{
	return myint + addint
}


func UseMyInt(){
	Valone := MyInt(3)
	Valtwo := MyInt(4)
	final := Valone.Add(Valtwo)
	fmt.Printf("finall is %d", final)
}