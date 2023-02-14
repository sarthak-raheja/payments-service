package processor

import (
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/repository"
	"github.com/sarthakraheja/payments-service/internal/settlement/settlement_router"
)

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

type GetPaymentRequest struct {
	Id         string
	MerchantId string
}

type GetPaymentResponse struct {
	Payment *model.Payment
}

type Processor interface {
	ProcessPayment(*ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	GetPayment(*GetPaymentRequest) (*GetPaymentResponse, error)
}

func NewProcessor(repo repository.Repository, settlementRouter settlement_router.AcquiringBankRouter) Processor {
	return &processor{
		repo:   repo,
		router: settlementRouter,
	}
}
