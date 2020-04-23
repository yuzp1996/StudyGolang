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

func NewLightOnCommand(light Light) LightOnCommand {
	return LightOnCommand{light}
}

func (command LightOnCommand) execute() {
	command.Light.on()
}

type LightOffCommand struct {
	Light
}

func NewLightOffCommand(light Light) LightOffCommand {
	return LightOffCommand{light}
}

func (command LightOffCommand) execute() {
	command.Light.off()
}
