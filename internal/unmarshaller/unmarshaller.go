package unmarshaller

import (
	"fmt"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
)

func (u *unmarshaller) unmarshallCreditCardPaymentMethod(creditCardPm *api.PaymentMethodCreditCardDetails) (*model.PaymentMethodCreditCard, error) {
	if creditCardPm == nil {
		return nil, fmt.Errorf("unable to unmarshall nil credit card details")
	}

	creditCardType := u.resolveCreditCardType(creditCardPm.GetCreditCardType())

	return &model.PaymentMethodCreditCard{
		CardHolderName:   creditCardPm.GetCardHolderName(),
		CreditCardNumber: creditCardPm.GetCreditCardNumber(),
		ExpiryDate:       creditCardPm.GetExpiryDate(),
		Cvv:              creditCardPm.GetCvv(),
		CreditCardType:   creditCardType,
	}, nil

}

func (u *unmarshaller) resolveCreditCardType(creditCardType api.CreditCardType) model.CreditCardType {
	switch creditCardType {
	case api.CreditCardType_CreditCardType_AMEX:
		return model.CreditCardType_Amex
	case api.CreditCardType_CreditCardType_VISA:
		return model.CreditCardType_Visa
	case api.CreditCardType_CreditCardType_MASTERCARD:
		return model.CreditCardType_MasterCard
	}
	return model.CreditCardType("")
}

func (u *unmarshaller) UnmarshallPaymentMethod(pm *api.PaymentMethod) (*model.PaymentMethod, error) {
	paymentMethod := &model.PaymentMethod{}
	if pm == nil {
		return nil, fmt.Errorf("unable to unmarshall nil payment method")
	}

	pmType := pm.GetPaymentMethodType()

	if pmType == api.PaymentMethodType_PaymentMethodType_CREDITCARD {
		ccApi := pm.GetPaymentMethodCreditCardDetails()

		creditCardDetails, err := u.unmarshallCreditCardPaymentMethod(ccApi)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshall credit card details")
		}

		paymentMethod.PaymentMethodType = model.PaymentMethodType_CreditCard
		paymentMethod.PaymentMethodCreditCard = creditCardDetails
	}

	return paymentMethod, nil
}
