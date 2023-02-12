package settlement_factory

import (
	"github.com/sarthakraheja/payments-service/internal/model"
)

type AcquringBank interface {
	AuthorizePayment(payment *model.Payment) error
	CapturePayment(*model.Payment) error
}
