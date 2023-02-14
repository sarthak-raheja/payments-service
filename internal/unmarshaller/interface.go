package unmarshaller

import (
	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type unmarshaller struct{}

type Unmarshaller interface {
	UnmarshallPaymentMethod(*api.PaymentMethod) (*model.PaymentMethod, error)
}

// NewUnmarshaller returns an interface for the Unmarshaller.
// Usage: Handles the transformation from protobuf to internal models
// format.
func NewUnmarshaller() Unmarshaller {
	return &unmarshaller{}
}
