package Adapter_test

import (
	"StudyGolang/designPattern/Adapter"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Adapter", func() {
	It("adapter will work, turkey will work like duck", func() {
		turkey := Adapter.NewTurkey()
		duckadapter := Adapter.NewAdapter(turkey)
		duckadapter.Fly()
		duckadapter.Quack()

	})

})
