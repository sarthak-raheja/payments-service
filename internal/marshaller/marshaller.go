package marshaller

type marshaller struct{}

type Marshaller interface{}

// NewMarshaller returns an interface for the Marshaller.
// Usage: Handles the transformation from internal models to protobuf
// format for GRPC.

func NewMarshaller() Marshaller {
	return marshaller{}
}
