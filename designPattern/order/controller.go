package order

type Controller struct {
	slot Command
}

func (controller Controller) SetCommand(command Command) {
	controller.slot = command
}

func (controller Controller) PressButton(command Command) {
	controller.slot.execute()
}
