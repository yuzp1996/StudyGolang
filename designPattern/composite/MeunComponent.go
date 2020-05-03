package composite

import "fmt"

// MeunComponent represent the iterface of MeunComponent like Meun and MeunItem
type MeunComponentInterface interface {
	Add(MeunComponentInterface)
	GetChild(int)MeunComponentInterface

	GetName()
	GetDescription()
	GetPrice()

	Print()
}

type MeunComponent struct {
	Name string
	Description string
	Price int
	Items []MeunComponentInterface
}

func NewMeunComponent(name string, description string, price int) *MeunComponent {
	return &MeunComponent{Name: name, Description: description, Price: price}
}

func (m MeunComponent) Add(MeunComponentInterface) {
	fmt.Println("Not Support Method")
}

func (m MeunComponent) GetChild(int)MeunComponentInterface {
	fmt.Println("Not Support Method")
	return nil
}

func (m MeunComponent) GetName() {
	fmt.Println("Not Support Method")
}

func (m MeunComponent) GetDescription() {
	fmt.Println("Not Support Method")
}

func (m MeunComponent) GetPrice() {
	fmt.Println("Not Support Method")
}

func (m MeunComponent) Print() {
	fmt.Println("Not Support Method")
}




