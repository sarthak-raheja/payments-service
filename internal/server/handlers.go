package server

import (
	"context"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/processor"
)

func (s *server) CreatePayment(ctx context.Context, req *api.CreatePaymentRequest) (*api.CreatePaymentResponse, error) {
	err := s.validator.ValidateCreatePayment(ctx, req)
	if err != nil {
		return nil, err
	}

	processorPaymentRequest := &processor.ProcessPaymentRequest{
		IdempotencyKey: "123",
		Amount:         "100.00",
		Currency:       "USD",
		MerchantId:     "123",
		PaymentMethod: &model.PaymentMethod{
			PaymentMethodType: model.PaymentMethodType_CreditCard,
			PaymentMethodCreditCard: &model.PaymentMethodCreditCard{
				CardHolderName:   "hello",
				CreditCardNumber: "123232",
				ExpiryDate:       "12-04",
				Cvv:              "222",
				CreditCardType:   model.CreditCardType_Amex,
			},
		},
	}

	_, err = s.processor.ProcessPayment(processorPaymentRequest)
	if err != nil {
		return nil, err
	}
	return &api.CreatePaymentResponse{}, nil
}

func (s *server) GetPayment(ctx context.Context, req *api.GetPaymentRequest) (*api.GetPaymentResponse, error) {
	return nil, nil
}
