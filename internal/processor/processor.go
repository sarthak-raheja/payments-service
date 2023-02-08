package processor

type processor struct{}

type Processor interface{}

func NewProcessor() Processor {
	return processor{}
}
