package marshaller

import (
	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type marshaller struct{}

func (m *marshaller) MarshalPayment(payment *model.Payment) (*api.Payment, error) {
	paymentPb := &api.Payment{}

	paymentPb.Id = payment.Id
	paymentPb.IdempotencyKey = payment.IdempotencyKey
	paymentPb.Amount = payment.Amount
	paymentPb.Currency = payment.Currency
	paymentPb.MerchantId = payment.MerchantId
	paymentPb.PaymentStatus = m.resolvePaymentStatus(payment.PaymentStatus)

	pm, err := m.marshalPaymentMethod(payment.PaymentMethod)
	if err != nil {
		return nil, err
	}
	paymentPb.PaymentMethod = &pm
	return paymentPb, nil

}

func (m *marshaller) marshalPaymentMethod(pm *model.PaymentMethod) (api.PaymentMethod, error) {
	pmType := pm.PaymentMethodType
	pmPb := api.PaymentMethod{}

	switch pmType {
	case model.PaymentMethodType_CreditCard:
		cc := pm.PaymentMethodCreditCard

		ccType := m.resolveCreditCardType(cc.CreditCardType)

		pmPb = api.PaymentMethod{
			PaymentMethodType: api.PaymentMethodType_PaymentMethodType_CREDITCARD,
			PaymentMethodsDetails: &api.PaymentMethod_PaymentMethodCreditCardDetails{
				PaymentMethodCreditCardDetails: &api.PaymentMethodCreditCardDetails{
					CardHolderName:   cc.CardHolderName,
					CreditCardNumber: cc.CreditCardNumber,
					ExpiryDate:       cc.ExpiryDate,
					Cvv:              cc.Cvv,
					CreditCardType:   ccType,
				},
			},
		}

	}
	return pmPb, nil
}

func (m *marshaller) resolveCreditCardType(creditCardType model.CreditCardType) api.CreditCardType {
	switch creditCardType {
	case model.CreditCardType_Amex:
		return api.CreditCardType_CreditCardType_AMEX
	case model.CreditCardType_Visa:
		return api.CreditCardType_CreditCardType_VISA
	case model.CreditCardType_MasterCard:
		return api.CreditCardType_CreditCardType_MASTERCARD
	}
	return api.CreditCardType_CreditCardType_UNSPECIFIED
}

func (m *marshaller) resolvePaymentStatus(paymentStatus model.PaymentStatus) api.PaymentStatus {
	switch paymentStatus {
	case model.PaymentStatus_Created:
		return api.PaymentStatus_PaymentStatus_CREATED
	case model.PaymentStatus_Processing:
		return api.PaymentStatus_PaymentStatus_PROCESSING
	case model.PaymentStatus_Processed:
		return api.PaymentStatus_PaymentStatus_PROCESSED
	case model.PaymentStatus_Completed:
		return api.PaymentStatus_PaymentStatus_COMPLETED
	case model.PaymentStatus_Failed:
		return api.PaymentStatus_PaymentStatus_FAILED
	default:
	}

	return api.PaymentStatus_PaymentStatus_UNSPECIFIED
}
