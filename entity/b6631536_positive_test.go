package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewWithT(t)

	book := Book {
		Title : "Golang",
		Price : 50.00,
		Code : "BK123456" ,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}