package settlement_factory

import (
	"context"

	bankSimulator "github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type chase struct {
	acquiringBankClient bankSimulator.AcquiringBankServiceClient
}

func (chase *chase) AuthorizePayment(*model.Payment) error {
	chaseReq := &bankSimulator.AuthorisePaymentRequest{}

	_, err := chase.acquiringBankClient.AuthorisePayment(context.Background(), chaseReq)

	return err
}

func (chase *chase) CapturePayment(*model.Payment) error {
	chaseReq := &bankSimulator.CapturePaymentRequest{}

	_, err := chase.acquiringBankClient.CapturePayment(context.Background(), chaseReq)

	return err
}

func NewChaseAcquiringBank() AcquringBank {
	return &chase{}
}
