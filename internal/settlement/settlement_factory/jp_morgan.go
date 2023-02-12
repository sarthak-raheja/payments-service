package settlement_factory

import (
	"github.com/sarthakraheja/payments-service/internal/model"
)

type jpMorgan struct{}

func (jp *jpMorgan) AuthorizePayment(*model.Payment) error {
	return nil
}

func (jp *jpMorgan) CapturePayment(*model.Payment) error {
	return nil
}
