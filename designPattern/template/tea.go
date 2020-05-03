package template

import (
	"bufio"
	"fmt"
	"os"
)



type Tea struct {
	CaffeineBeverage
	Name string
}

func NewTea(name string, base CaffeineBeverage) Tea {
	return Tea{base, name}
}

func (tea Tea) AddCodiments() {
	fmt.Println("add codiments for tea ...")
}

func (tea Tea) Brew() {
	fmt.Println("brew beverage for tea ...")
}

func (tea Tea) BoilWater() {
	fmt.Println("boil water  for tea ...")
}

func (tea Tea) NeedCodiments()bool {
	fmt.Printf("hello do you need codiments?\n")
	reader := bufio.NewReader(os.Stdin)
	text,_:=reader.ReadString('\n')
	if text == "yes"{
		fmt.Printf("I will add codiment for you, Thanks\n")
		return true
	}
	fmt.Printf("OK, No codiment will added!\n")
	return false
}
