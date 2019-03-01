package cart_test

import (
	. "StudyGolang/ginkgo/cart"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {
	Context("initially", func() {
		carts := Cart{}

		It("has 0 items", func() {
			Expect(carts.TotalUniqueItems()).Should(BeZero())
		})
		It("has 0 units", func() {
			Expect(carts.TotalUnits()).Should(BeZero())
		})
		Specify("the total amount is 0.00", func() {
			Expect(carts.TotalUniqueItems()).Should(BeZero())
		})
	})
	Context("When a new item is added", func() {
		item1 := Item{Name: "test"}
		carts := Cart{}
		carts.AddItem(item1)
		Specify("has 1 more unique item than it had earlier", func() {
			Expect(carts.Totalmount()).Should(Equal(1))
		})
		It("has 1 more unit than it had earlier", func() {})
		Specify("total amount increases by item price", func() {})
	})

})
