package settlement_factory

import (
	bankSimulator "github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type chase struct {
	acquiringBankClient bankSimulator.AcquiringBankServiceClient
}

func (chase *chase) AuthorizePayment(*model.Payment) error {
	return nil
}

func (chase *chase) CapturePayment(*model.Payment) error {
	return nil
}

func NewChaseAcquiringBank() AcquringBank {
	return &chase{}
}
