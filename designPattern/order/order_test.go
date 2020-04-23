package order_test

import (
	"StudyGolang/designPattern/order"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Order", func() {

	It("simulate controller", func() {
		controller := order.Controller{}
		onCommand := order.NewLightOnCommand(order.Light{})
		controller.SetCommand(onCommand)
		controller.PressButton()

		offCommand := order.NewLightOffCommand(order.Light{})
		controller.SetCommand(offCommand)
		controller.PressButton()

	})

})
