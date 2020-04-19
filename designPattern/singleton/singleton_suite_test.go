package singleton_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSingleton(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Singleton Suite")
}
