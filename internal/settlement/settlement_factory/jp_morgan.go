package settlement_factory

import (
	bankSimulator "github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type jpMorgan struct {
	acquiringBankClient bankSimulator.AcquiringBankServiceClient
}

func (jp *jpMorgan) AuthorizePayment(*model.Payment) error {
	return nil
}

func (jp *jpMorgan) CapturePayment(*model.Payment) error {
	return nil
}
