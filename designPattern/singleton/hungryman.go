package singleton

import "fmt"

var hungryceo *CEO

func init() {
	fmt.Println("hungry ceo is creating")
	hungryceo = new(CEO)
	hungryceo.name = "hungryceo"
}

// GetHungryCEO
func GetHungryCEO() *CEO {
	fmt.Printf("my name %v\n", hungryceo.name)

	return hungryceo
}
