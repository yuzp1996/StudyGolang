package composite_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestComposite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Composite Suite")
}
