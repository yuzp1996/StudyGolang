package duck_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDuck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Duck Suite")
}
