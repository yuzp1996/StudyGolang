package composite

import "fmt"

type Meun struct {
	Name string
	Description string
	Item []MeunComponentInterface
	MeunComponent
}

func NewMeun(name string, description string,  meunComponent MeunComponent) *Meun {
	return &Meun{Name: name, Description: description,MeunComponent: meunComponent}
}

func (m *Meun) Add(item MeunComponentInterface) {
	m.Items = append(m.Items,item)
}

func (m *Meun)GetChild(index int)MeunComponentInterface{
	return m.Items[index]
}


func (meun Meun)Print()  {
	fmt.Printf("This is Menu Named %s and Description is %s\n", meun.Name,meun.Description)
	for _,item := range meun.Items{
		item.Print()
	}
	fmt.Println()
}
