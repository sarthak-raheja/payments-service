package processor

import (
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/repository"
	"github.com/sarthakraheja/payments-service/internal/settlement/settlement_router"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type processor struct {
	repo   repository.Repository
	router settlement_router.AcquiringBankRouter
}

func (p *processor) ProcessPayment(req *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	// Register Payment in the database
	payment, err := p.registerPayment(req)
	if err != nil {
		return nil, err
	}

	// Get appropriate acquring Bank from the Router
	acquiringBank, err := p.router.GetAcquiringBank(payment)
	if err != nil {
		return nil, err
	}

	// Authorize Payment
	err = acquiringBank.AuthorizePayment(payment)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "could not authorize payment")
	}

	// Capture Payment
	err = acquiringBank.CapturePayment(payment)
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, "could not capture payment")
	}

	return &ProcessPaymentResponse{
		Payment: payment,
	}, nil
}

// registerPayment registers the payment in the Database
func (p *processor) registerPayment(req *ProcessPaymentRequest) (*model.Payment, error) {
	payment := &model.Payment{
		IdempotencyKey: req.IdempotencyKey,
		Amount:         req.Amount,
		Currency:       req.Currency,
		MerchantId:     req.MerchantId,
		PaymentMethod:  req.PaymentMethod,
		PaymentStatus:  model.PaymentStatus_Created,
	}

	payment, err := p.repo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
