package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeCode(t *testing.T) {
	g := NewWithT(t)

	book := Book {
		Title : "Golang",
		Price : 50.00,
		Code : "Bj123456" ,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
}