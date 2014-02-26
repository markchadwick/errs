package errs

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	"testing"
)

func TestErrs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errs Suite")
}

func makeError(name string) error {
	return New(name)
}

func nameError() error {
	return Wrap(makeError("bad name"))
}

var _ = Describe("Errs", func() {

	It("should wrap an error", func() {
		err := Wrap(nameError())
		stack := strings.Split(err.Error(), "\n")
		Expect(stack).To(HaveLen(4))
		Expect(stack[0]).To(Equal("bad name"))
	})
})
