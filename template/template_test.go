package template_test

import (
	"StudyGolang/template"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Template", func() {
	template.ExampleTemplate()
	template.ExampleTemplate_block()

})
