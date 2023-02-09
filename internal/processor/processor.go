package processor

import (
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/repository"
)

type processor struct {
	repo repository.Repository
}

type ProcessPaymentRequest struct {
	idempotencyKey string
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

	return nil, nil
}
