package Struct

import "fmt"

func Deduplica(){
	var TestMap = map[string]struct{}{}
	TestMap["first"] = struct{}{}
	TestMap["second"] = struct{}{}
	TestMap["zpyu"] = struct{}{}
	TestMap["zpyu"] = struct{}{}
	fmt.Printf("%v", TestMap)
}
