package server

import (
	"context"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
)

func (s *server) CreatePayment(ctx context.Context, req *api.CreatePaymentRequest) (*api.CreatePaymentResponse, error) {
	err := s.validator.ValidateCreatePayment(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *server) GetPayment(ctx context.Context, req *api.GetPaymentRequest) (*api.GetPaymentResponse, error) {
	return nil, nil
}
