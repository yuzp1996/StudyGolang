package Struct

import (
	"encoding/json"
	"fmt"
)

func Deduplica() {
	//var TestMap = map[string]struct{}{}
	//TestMap["first"] = struct{}{}
	//TestMap["second"] = struct{}{}
	//TestMap["zpyu"] = struct{}{}
	//TestMap["zpyu"] = struct{}{}
	//fmt.Printf("%v", TestMap)

	result := getinterface()

	fmt.Printf("result is %v", result)
	if v, ok := result.(map[string]interface{}); ok {
		fmt.Printf("\nv is %v and ok is %v", v, ok)
	}

}

func getinterface() interface{} {
	value := `{"type":"select"}`

	test := teststruct{}
	json.Unmarshal([]byte(value), &test)

	fmt.Printf("result is %v\n", test)

	return test
}

type teststruct struct {
	Type string `json:"type"`
}
