package singleton

import (
	"fmt"
	"sync"
)

var safeceo *CEO
var mux sync.Mutex

// GetSafeLazyCEO get ceo instance with lazy pattern and in the safe way
func GetSafeLazyCEO() *CEO {
	if safeceo == nil {
		mux.Lock()
		defer mux.Unlock()
		if safeceo == nil {
			fmt.Println("create a safeceo now")
			safeceo = new(CEO)
			safeceo.name = "safegougou"
		}
	}
	fmt.Printf("my name %v\n", safeceo.name)
	return safeceo
}
