package unmarshaller

import (
	"fmt"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type unmarshaller struct{}

type Unmarshaller interface {
	UnmarshallPaymentMethod(*api.PaymentMethod) (*models.PaymentMethod, error)
}

// NewUnmarshaller returns an interface for the Unmarshaller.
// Usage: Handles the transformation from protobuf to internal models
// format for GRPC.
func NewUnmarshaller() Unmarshaller {
	return unmarshaller{}
}

func (u *unmarshaller) unmarshallCreditCardPaymentMethod(creditCardPm *api.PaymentMethodCreditCardDetails) (*models.PaymentMethodCreditCard, error) {
	if creditCardPm == nil {
		return nil, fmt.Errorf("unable to unmarshall nil credit card details")
	}

	return &model.CreditCardDetails{}, nil

}

func (u *unmarshaller) UnmarshallPaymentMethod(pm *api.PaymentMethod) (*models.PaymentMethod, error) {
	var paymentMethod *model.PaymentMethod
	if pm == nil {
		return nil, fmt.Errorf("unable to unmarshall nil payment method")
	}

	pmType, err := pm.GetPaymentMethodType()
	if err != nil {
		return nil, fmt.Errorf("unable to determine payment method type")
	}

	if pmType == api.PaymentMethodType_CREDITCARD {
		creditCardDetails, err := u.unmarshallCreditCardPaymentMethod(PaymentMethodCreditCardDetails)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshall credit card details")
		}

		paymentMethod.PaymentMethodType = model.PaymentMethodType_CreditCard
		paymentMethod.PaymentMethodCreditCard = creditCardDetails
	}

	return paymentMethod, nil
}
