package entity_test

import (
	"testing"

	"github.com/Chinnapatz/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerCorrect(t *testing.T)  {
	g := NewWithT(t)
	
	t.Run(`case correct`, func (t *testing.T)  {
		customer := entity.Customers {
			Name: "best",
			Age: "12",
			CustomerID: "CM12345678",
		}
		ok,err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

		
	})
	
}