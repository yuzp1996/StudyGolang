package playdefer

import (
	"fmt"
)

func FinalFunc(biggestnum int){
	fmt.Printf("Finall function  and biggest num is %d\n", biggestnum)
}

func Biggest(nums []int)(num int){
    num = nums[0]
    for _, v := range(nums){
		if v > num{
			num = v
		}
	}
    return num

}

func finalend(){
	fmt.Println("everything ends")
}
func MainDefer(){
	defer finalend()
	nums := []int{12,32,34,44,53,253,234,33}
	num := Biggest(nums)
	FinalFunc(num)

	person := Person{
		"Yuzhipeng",
	}
	person.Name = "zpyu"
	defer person.PersonName()
	person.Name = "yuzhipeng"
	fmt.Println("Welcome")
	reverstirng()
}


type Person struct{
	Name string
}

func (person Person)PersonName(){
	fmt.Println(person.Name)
}



func reverstirng(){
	name := "zpyu"
	for _, v := range []rune(name){
		defer fmt.Printf("%c\n",v)
	}
}
