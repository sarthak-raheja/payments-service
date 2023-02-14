package unmarshaller_test

import (
	"testing"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/unmarshaller"

	"github.com/stretchr/testify/suite"
)

type UnmarshallerTestSuite struct {
	suite.Suite

	unmarshaller unmarshaller.Unmarshaller
}

func NewUnmarshallerTestSuite(t *testing.T) UnmarshallerTestSuite {
	unmarshaller := unmarshaller.NewUnmarshaller()

	return UnmarshallerTestSuite{
		unmarshaller: unmarshaller,
	}
}

type UnmarshallerTestCases struct {
	input    api.PaymentMethod
	expected model.PaymentMethod
}

func TestHandlerTestSuite(t *testing.T) {
	unmarshallerTestSuite := NewUnmarshallerTestSuite(t)

	suite.Run(t, &unmarshallerTestSuite)
}

func setupTestCases() []UnmarshallerTestCases {
	var testCases []UnmarshallerTestCases

	pm := api.PaymentMethod{
		PaymentMethodType: api.PaymentMethodType_PaymentMethodType_CREDITCARD,
		PaymentMethodsDetails: &api.PaymentMethod_PaymentMethodCreditCardDetails{
			PaymentMethodCreditCardDetails: &api.PaymentMethodCreditCardDetails{
				CardHolderName:   "hello",
				CreditCardNumber: "1234123412341234",
				ExpiryDate:       "12-04",
				Cvv:              "222",
				CreditCardType:   api.CreditCardType_CreditCardType_AMEX,
			},
		},
	}

	expectedPm := model.PaymentMethod{
		PaymentMethodType: model.PaymentMethodType_CreditCard,
		PaymentMethodCreditCard: &model.PaymentMethodCreditCard{
			CardHolderName:   "hello",
			CreditCardNumber: "1234123412341234",
			ExpiryDate:       "12-04",
			Cvv:              "222",
			CreditCardType:   model.CreditCardType_Amex,
		},
	}

	testCases = append(testCases, UnmarshallerTestCases{
		input:    pm,
		expected: expectedPm,
	})
	return testCases
}

func (unmarshallerTest *UnmarshallerTestSuite) Test_Unmarshaller() {
	testCases := setupTestCases()

	for _, v := range testCases {

		resp, err := unmarshallerTest.unmarshaller.UnmarshallPaymentMethod(&v.input)

		unmarshallerTest.Assert().Nil(err)
		unmarshallerTest.Assert().Equal(*resp, v.expected)

	}
}
