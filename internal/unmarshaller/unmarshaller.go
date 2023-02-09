package unmarshaller

type unmarshaller struct{}

type Unmarshaller interface {
	// UnmarshallCreatePaymentRequest(*api.CreatePaymentRequest)
}

// NewUnmarshaller returns an interface for the Unmarshaller.
// Usage: Handles the transformation from protobuf to internal models
// format for GRPC.
func NewUnmarshaller() Unmarshaller {
	return unmarshaller{}
}

// func (u *unmarshaller) UnmarshallCreatePaymentRequest() {

// }
