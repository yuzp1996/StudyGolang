package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once
var bestceo *CEO

func GetOnceDoCEO() *CEO {
	once.Do(func() {
		fmt.Println("do once ceo is creating")
		bestceo = new(CEO)
		bestceo.name = "bestCEO"
	})
	return bestceo
}
