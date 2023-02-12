package settlement_factory

import (
	"github.com/sarthakraheja/payments-service/internal/model"
)

type chase struct{}

func (chase *chase) AuthorizePayment(*model.Payment) error {
	return nil
}

func (chase *chase) CapturePayment(*model.Payment) error {
	return nil
}

func NewChaseAcquiringBank() AcquringBank {
	return &chase{}
}
