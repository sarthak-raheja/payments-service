package validator

import (
	"context"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
)

type validator struct{}

// Validator Interface implements validation methods for the API request
// Currently the interface is responsible for field validations but can be extended for
// additional validation as well for example: Fraud Detection, Billing address verification
type Validator interface {
	ValidateCreatePayment(ctx context.Context, req *api.CreatePaymentRequest) error
}

func NewValidator() Validator {
	return &validator{}
}
