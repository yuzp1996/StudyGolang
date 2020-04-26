package template_test

import (
	"StudyGolang/designPattern/template"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Template", func() {
	It("template pattern should work well", func() {
		caffeinebeverage := template.NewCaffeineBeverage()

		coffee := template.NewCoffee("coffee", *caffeinebeverage)
		tea := template.NewTea("tea", *caffeinebeverage)
		waiter := template.NewWaiter("zpyu")
		waiter.MakeDrink(tea)
		waiter.MakeDrink(coffee)

	})

})
