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

	It("should pass through a nil error", func() {
		err := Wrap(nil)
		Expect(err).To(BeNil())
	})

	It("should handle wrapping a nil error", func() {
		var e0 *Err = nil
		Expect(e0).To(BeNil())

		var e1 = Wrap(e0)
		Expect(e1).To(BeNil())
	})
})
