package server

import (
	"context"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/processor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (s *server) CreatePayment(ctx context.Context, req *api.CreatePaymentRequest) (*api.CreatePaymentResponse, error) {
	err := s.validator.ValidateCreatePayment(ctx, req)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "error: %v", err)
	}

	pm, err := s.unmarshaller.UnmarshallPaymentMethod(req.GetPaymentMethod())
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "error: %v", err)
	}
	processorPaymentRequest := &processor.ProcessPaymentRequest{
		IdempotencyKey: req.GetIdempotencyKey(),
		Amount:         req.GetAmount(),
		Currency:       req.GetCurrency(),
		MerchantId:     req.GetMerchantId(),
		PaymentMethod:  pm,
	}

	processorResp, err := s.processor.ProcessPayment(processorPaymentRequest)
	if err != nil {
		return nil, err
	}

	payment, err := s.marshaller.MarshalPayment(processorResp.Payment)
	if err != nil {
		return nil, err
	}
	return &api.CreatePaymentResponse{
		Payment: payment,
	}, nil
}

func (s *server) GetPayment(ctx context.Context, req *api.GetPaymentRequest) (*api.GetPaymentResponse, error) {
	return nil, nil
}
