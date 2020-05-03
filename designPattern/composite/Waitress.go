package composite

import "fmt"

type Waitress struct {
	Name string
	Meun []MeunComponentInterface
}

func NewWaitress(name string, meun []MeunComponentInterface) *Waitress {
	return &Waitress{Name: name, Meun: meun}
}

func(waitress *Waitress)AddMeun(meun MeunComponentInterface){
	waitress.Meun = append(waitress.Meun,meun)
}

func(waitress Waitress)PrintMeun(){
	fmt.Printf("Hello I am %v, Now the meun is below\n",waitress.Name)
	for _,meun := range waitress.Meun{
		meun.Print()
	}
}
