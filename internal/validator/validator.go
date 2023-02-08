package validator

type validator struct{}

type Validator interface{}

func NewValidator() Validator {
	return &validator{}
}
