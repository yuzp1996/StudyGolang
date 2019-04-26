package maptest1

import (
	"fmt"
)

func Newmap() {
	var nilmap map[string]string
	if nilmap == nil {
		fmt.Println("I am a nil map")
		//map 必须使用 make 函数初始化。
		nilmap = make(map[string]string)
	}

	if nilmap != nil {
		fmt.Println("nilmap is not nil")
	}
	nilmap["first"] = "first"
	nilmap["second"] = "second"
	fmt.Printf("nilmap is %v\n", nilmap)

	goodmap := map[string]string{
		"first":  "first",
		"second": "second",
	}
	fmt.Printf("not nil map is %v\n", goodmap)
	fmt.Printf("get fiirst is %s", goodmap["fiirst"])

	fmt.Println(" if a key is exsit")

	//值是否存在
	value, ok := goodmap["first"]
	fmt.Printf("k is %v, and v is %v\n", value, ok)
	for k, v := range goodmap {
		fmt.Printf("for range k is %v and v is %v\n", k, v)
	}
	fmt.Printf("len of good map is %d delete element in goodmap\n", len(goodmap))
	delete(goodmap, "first")
	fmt.Printf("after delete len is %d\n", len(goodmap))

	//map是引用类型  所以如果map1赋给map2的话 map2改变的话 map1也会改变
	Salay := map[string]int{
		"yuzhipeng":   12000,
		"zhangjiajie": 13000,
	}
	SalayCopy := Salay
	fmt.Println("Salaycopy up your salay")
	SalayCopy["yuzhipeng"] = 13000
	fmt.Printf("Salay now is %v\n", Salay)

	fmt.Println("I want to know if two map [Salay]&[Salay] is same")
	same := false

//比较map是否相等
	for k, v := range Salay {
		if SalayCopy[k] == v {
			same = true
			break
		}
	}
	if same {
		fmt.Println("yes 是相同的")
	} else {
		fmt.Println("no it is different")
	}

}


func MapTesthowtogetkey(){
	person1 := Testpersonstruct{"zpyu",10}
	person2 := Testpersonstruct{"jingtao",13}
	personstructMap := map[string]Testpersonstruct{}
	personstructMap[person1.Name] = person1
	personstructMap[person2.Name] = person2
	fmt.Println(personstructMap)
	fmt.Println(personstructMap["zpyu"])
	for _, personmapelement := range personstructMap{

		fmt.Println(personmapelement)
	}


}

type Testpersonstruct struct {
	Name string
	Age int
}