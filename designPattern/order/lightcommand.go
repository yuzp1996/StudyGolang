package order

import "fmt"

type Light struct {
}

func (light Light) on() {
	fmt.Println("Light on")
}

func (light Light) off() {
	fmt.Println("Light off")
}

type LightOnCommand struct {
	Light
}

func (command LightOnCommand) execute() {
	command.Light.on()
}

type LightOffCommand struct {
	Light
}

func (command LightOffCommand) execute() {
	command.Light.off()
}
