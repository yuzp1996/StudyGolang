package observer

import "fmt"

func Simulate(){

	subject := NewWeaterDate()

	panel := NewPanelHeat(&subject)
	panelone := NewPanelHeatOne(&subject)
	//other panel

	subject.registryobserver(panel)
	subject.registryobserver(panelone)
	subject.setrealtimeData("20C")

	fmt.Println("Delete panel panel")
	subject.deleteobserver(panel)
	subject.setrealtimeData("30C")



}
