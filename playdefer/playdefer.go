package playdefer

import (
	"fmt"
)

func FinalFunc(biggestnum int) {
	fmt.Printf("Finall function  and biggest num is %d\n", biggestnum)
}

func Biggest(nums []int) (num int) {
	num = nums[0]
	for _, v := range nums {
		if v > num {
			num = v
		}
	}
	return num

}

func finalend() {
	fmt.Println("everything ends")
}

// 就是个栈  越在前面 就越后面执行
// 也就是defer是在最后执行的 然后defer也有执行顺序 那就是从下往上执行
func MainDefer() {
	defer finalend()
	nums := []int{12, 32, 34, 44, 53, 253, 234, 33}
	num := Biggest(nums)
	FinalFunc(num)

	person := Person{
		"Yuzhipeng",
	}
	person.Name = "zpyu"
	//这里输出的PersonName 是zpyu  不是yuzhipeng 在defer执行的时候 就确定了参数值了
	defer person.PersonName()
	person.Name = "yuzhipeng"
	fmt.Println("Welcome")
	reverstirng()
}

type Person struct {
	Name string
}

func (person Person) PersonName() {
	fmt.Println(person.Name)
}

//翻转字符串 多好
func reverstirng() {
	name := "zpyu"
	for _, v := range []rune(name) {
		defer fmt.Printf("%c\n", v)
	}
}
