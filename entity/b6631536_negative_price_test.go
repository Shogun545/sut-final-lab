package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativePrice(t *testing.T) {
	g := NewWithT(t)

	book := Book {
		Title : "Golang",
		Price : 4.00,
		Code : "BK123456" ,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
}