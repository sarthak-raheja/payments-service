package settlement_factory

import (
	"context"

	bankSimulator "github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type jpMorgan struct {
	acquiringBankClient bankSimulator.AcquiringBankServiceClient
}

func (jp *jpMorgan) AuthorizePayment(*model.Payment) error {
	jpReq := &bankSimulator.AuthorisePaymentRequest{}

	_, err := jp.acquiringBankClient.AuthorisePayment(context.Background(), jpReq)

	return err
}

func (jp *jpMorgan) CapturePayment(*model.Payment) error {
	jpReq := &bankSimulator.CapturePaymentRequest{}

	_, err := jp.acquiringBankClient.CapturePayment(context.Background(), jpReq)

	return err
}
