package maptest1

import (
	"fmt"
)

func Newmap() {
	var nilmap map[string]string
	if nilmap == nil {
		fmt.Println("I am a nil map")
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
	k, v := goodmap["first"]
	fmt.Printf("k is %v, and v is %v\n", k, v)
	for k, v := range goodmap {
		fmt.Printf("for range k is %v and v is %v\n", k, v)
	}
	fmt.Printf("len of good map is %d delete element in goodmap\n", len(goodmap))
	delete(goodmap, "first")
	fmt.Printf("after delete len is %d\n", len(goodmap))

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
