package composite

import "fmt"

type MeunItem struct {
	Name string
	Description string
	Price int
	MeunComponent
}

func NewMeunItem(name string, description string, price int, meunComponent MeunComponent) *MeunItem {
	return &MeunItem{Name: name, Description: description, Price: price, MeunComponent: meunComponent}
}


func (item MeunItem)Print(){
	fmt.Printf("This is meunItem %s - %s, Price is %v\n", item.Name,item.Description, item.Price)
}