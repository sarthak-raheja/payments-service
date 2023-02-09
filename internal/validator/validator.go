package validator

import (
	"context"
	"strconv"
	"strings"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	GBP = "gbp"
	USD = "usd"
	EUR = "eur"
)

type validator struct{}

type Validator interface {
	ValidateCreatePayment(ctx context.Context, req *api.CreatePaymentRequest) error
}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) ValidateCreatePayment(ctx context.Context, req *api.CreatePaymentRequest) error {

	currency := req.GetCurrency()
	if currency == "" {
		return grpc.Errorf(codes.InvalidArgument, "currency not provided")
	}

	err := v.ValidateCurrency(currency)
	if err != nil {
		return grpc.Errorf(codes.InvalidArgument, "currency not supported: %v", currency)
	}

	amount, err := strconv.ParseFloat(req.GetAmount(), 64)
	if err != nil {
		return grpc.Errorf(codes.InvalidArgument, "unable to parse amount: %v", req.GetAmount())
	}

	if amount <= 0 {
		return grpc.Errorf(codes.InvalidArgument, "payment amount is invalid")
	}

	return nil
}

func (v *validator) ValidateCurrency(currency string) error {
	supportedCurrency := map[string]bool{
		USD: true,
		GBP: true,
		EUR: true,
	}

	_, found := supportedCurrency[strings.ToLower(currency)]
	if !found {
		return grpc.Errorf(codes.InvalidArgument, "currency not supported: %v", currency)
	}
	return nil
}
