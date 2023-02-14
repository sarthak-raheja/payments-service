package utils

import (
	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
)

const (
	obfuscatedCvv    = "xxx"
	obfuscatedExpiry = "XX-XX"
	obfuscatedCard   = "XXXX-XXXX-XXXX-"
)

func ObfuscateSensitiveFields(payment *api.Payment) *api.Payment {

	pm := payment.GetPaymentMethod()

	pmType := pm.GetPaymentMethodType()

	if pmType == api.PaymentMethodType_PaymentMethodType_CREDITCARD {
		ccApi := pm.GetPaymentMethodCreditCardDetails()

		creditCardDetails := obfuscateCreditCardPaymentMethod(ccApi)

		pm = &api.PaymentMethod{
			PaymentMethodType: api.PaymentMethodType_PaymentMethodType_CREDITCARD,
			PaymentMethodsDetails: &api.PaymentMethod_PaymentMethodCreditCardDetails{
				PaymentMethodCreditCardDetails: creditCardDetails,
			},
		}
	}

	payment.PaymentMethod = pm
	return payment

}

func obfuscateCreditCardPaymentMethod(cc *api.PaymentMethodCreditCardDetails) *api.PaymentMethodCreditCardDetails {

	obfuscatedCreditCardNumber := obfuscateCreditCardNumber(cc.GetCreditCardNumber())
	return &api.PaymentMethodCreditCardDetails{
		CardHolderName:   cc.GetCardHolderName(),
		CreditCardNumber: obfuscatedCreditCardNumber,
		ExpiryDate:       obfuscatedExpiry,
		Cvv:              obfuscatedCvv,
		CreditCardType:   cc.GetCreditCardType(),
	}
}

func obfuscateCreditCardNumber(ccNumber string) string {
	ccLen := len(ccNumber) - 4

	creditCardLast := ccNumber[ccLen:]
	return obfuscatedCard + creditCardLast
}
