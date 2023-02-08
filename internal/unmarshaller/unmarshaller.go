package marshaller

type unmarshaller struct{}

type Unmarshaller interface{}

// NewUnmarshaller returns an interface for the Unmarshaller.
// Usage: Handles the transformation from protobuf to internal models
// format for GRPC.
func NewUnmarshaller() Unmarshaller {
	return unmarshaller{}
}
