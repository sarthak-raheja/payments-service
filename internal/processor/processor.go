package processor

import (
	"github.com/google/uuid"
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/repository"
	settlement_router "github.com/sarthakraheja/payments-service/internal/settlement/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type processor struct {
	repo   repository.Repository
	router settlement_router.AcquiringBankRouter
}

type ProcessPaymentRequest struct {
	idempotencyKey string
	paymentMethod  *model.PaymentMethod
	amount         string
	currency       string
	merchantId     string
}

type ProcessPaymentResponse struct {
	Payment *model.Payment
}

type Processor interface {
	ProcessPayment(*ProcessPaymentRequest) (*ProcessPaymentResponse, error)
}

func NewProcessor(repo repository.Repository) Processor {
	return &processor{
		repo: repo,
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
		return nil, grpc.Errorf(codes.Unimplemented, "could not determine acquring bank")
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
		Payment: nil,
	}, nil
}

// registerPayment registers the payment in the Database
func (p *processor) registerPayment(req *ProcessPaymentRequest) (*model.Payment, error) {
	paymentId := uuid.NewString()
	eventIdem := uuid.NewString()

	paymentEvent := &model.PaymentEvent{
		PaymentId:        paymentId,
		IdempotencyKey:   eventIdem,
		Timestamp:        *timestamppb.Now(),
		PaymentEventType: model.PaymentEventType_Created,
	}

	payment := &model.Payment{
		Id:             uuid.NewString(),
		IdempotencyKey: req.idempotencyKey,
		Amount:         req.amount,
		Currency:       req.amount,
		PaymentEvent:   []*model.PaymentEvent{paymentEvent},
	}

	payment, err := p.repo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
