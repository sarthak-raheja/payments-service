package processor

import (
	"github.com/google/uuid"
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/repository"
	"github.com/sarthakraheja/payments-service/internal/settlement/settlement_router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type processor struct {
	repo   repository.Repository
	router settlement_router.AcquiringBankRouter
}

type ProcessPaymentRequest struct {
	IdempotencyKey string
	PaymentMethod  *model.PaymentMethod
	Amount         string
	Currency       string
	MerchantId     string
}

type ProcessPaymentResponse struct {
	Payment *model.Payment
}

type Processor interface {
	ProcessPayment(*ProcessPaymentRequest) (*ProcessPaymentResponse, error)
}

func NewProcessor(repo repository.Repository, settlementRouter settlement_router.AcquiringBankRouter) Processor {
	return &processor{
		repo:   repo,
		router: settlementRouter,
	}
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
	paymentId := uuid.NewString()

	payment := &model.Payment{
		Id:             paymentId,
		IdempotencyKey: req.IdempotencyKey,
		Amount:         req.Amount,
		Currency:       req.Currency,
		PaymentMethod:  req.PaymentMethod,
	}

	// payment, err := p.repo.CreatePayment(payment)
	//  if err != nil {
	//  	return nil, err
	//  }

	return payment, nil
}
