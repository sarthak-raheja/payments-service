package settlement_router

import (
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/settlement/settlement_factory"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AcquiringBankRouter interface {
	GetAcquiringBank(*model.Payment) (settlement_factory.AcquringBank, error)
}
type router struct {
	acquringBankFactory settlement_factory.AcquringBankFactory
}

func NewAcquiringBankRouter(acquringBankFactory settlement_factory.AcquringBankFactory) AcquiringBankRouter {
	return &router{
		acquringBankFactory: acquringBankFactory,
	}
}

func (r *router) GetAcquiringBank(payment *model.Payment) (settlement_factory.AcquringBank, error) {
	pm := payment.PaymentMethod
	if pm == nil {
		return nil, nil
	}

	switch pm.PaymentMethodType {
	case model.PaymentMethodType_CreditCard:
		cc := pm.PaymentMethodCreditCard
		switch cc.CreditCardType {
		case model.CreditCardType_Amex:
			return r.acquringBankFactory.BuildJPMorganAcquiringBank(), nil
		case model.CreditCardType_MasterCard:
			return r.acquringBankFactory.BuildChaseAcquringBank(), nil
		case model.CreditCardType_Visa:
			return r.acquringBankFactory.BuildChaseAcquringBank(), nil
		default:
			return nil, grpc.Errorf(codes.Unimplemented, "could not determine acquring bank")
		}
	default:
		return nil, grpc.Errorf(codes.Unimplemented, "could not determine acquiring bank")
	}
}
