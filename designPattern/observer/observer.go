package observer

import "fmt"

type observer interface {
	getname()string
	update(string)
	display()
}

//0 号观察者
type PanelHeat struct {
	Heat string
	*WeaterDate
	Name string
}


func NewPanelHeat(weaterdate *WeaterDate)*PanelHeat{
	return &PanelHeat{Heat:"20C",WeaterDate: weaterdate,Name:"panelheat"}
}

func(panel *PanelHeat)getname()string{
	return panel.Name
}

func (panel *PanelHeat)update(data string){
	 panel.Heat = string(data)
	 panel.display()

}
func (panel *PanelHeat)display(){
	fmt.Printf("【PanelHeat】now the heat is %v\n", panel.Heat)
}

// 1号观察者
type PanelHeatOne struct {
	Heat string
	*WeaterDate
	Name string
}


func NewPanelHeatOne(weaterdate *WeaterDate)*PanelHeatOne{
	return &PanelHeatOne{Heat:"20C",WeaterDate: weaterdate,Name:"panelheatone"}
}

func(panel *PanelHeatOne)getname()string{
	return panel.Name
}

func (panel *PanelHeatOne)update(data string){
	panel.Heat = string(data)
	panel.display()

}
func (panel *PanelHeatOne)display(){
	fmt.Printf("【PanelHeatOne】now the heat is %v\n", panel.Heat)
}