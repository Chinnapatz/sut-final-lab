package entity_test

import (
	"testing"

	"github.com/Chinnapatz/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomer(t *testing.T)  {
	g := NewWithT(t)
	t.Run(`case incorrect age`, func (t *testing.T)  {
		customer := entity.Customers {
			Name: "best",
			Age: "12.5",
			CustomerID: "CM12345678",
		}
		ok,err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Age is not interger"))
	})
	
}