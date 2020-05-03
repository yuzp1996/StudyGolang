package template_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/yuzp1996/StudyGolang/designPattern/template"
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
