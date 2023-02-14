package marshaller

import (
	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type Marshaller interface {
	MarshalPayment(*model.Payment) (*api.Payment, error)
}

// NewMarshaller returns an interface for the Marshaller.
// Usage: Handles the transformation from internal models to protobuf
// format for GRPC.
func NewMarshaller() Marshaller {
	return &marshaller{}
}
