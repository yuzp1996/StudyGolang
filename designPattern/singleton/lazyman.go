package singleton

import "fmt"

type CEO struct {
	name string
}

var ceo *CEO

// GetLazyCEO get ceo instance with lazy pattern
func GetLazyCEO() *CEO {
	if ceo == nil {
		fmt.Println("create a ceo now")
		ceo = new(CEO)
		ceo.name = "gougou"
	}
	fmt.Printf("my name %v\n", ceo.name)
	return ceo
}
