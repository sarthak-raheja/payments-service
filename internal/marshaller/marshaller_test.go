package marshaller_test

import (
	"testing"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/marshaller"
	"github.com/sarthakraheja/payments-service/internal/model"

	"github.com/stretchr/testify/suite"
)

type MarshallerTestSuite struct {
	suite.Suite

	marshaller marshaller.Marshaller
}

func NewMarshallerTestSuite(t *testing.T) MarshallerTestSuite {
	marshaller := marshaller.NewMarshaller()

	return MarshallerTestSuite{
		marshaller: marshaller,
	}
}

type MarshallerTestCases struct {
	input    model.Payment
	expected api.Payment
}

func TestHandlerTestSuite(t *testing.T) {
	marshallerTestSuite := NewMarshallerTestSuite(t)

	suite.Run(t, &marshallerTestSuite)
}

func setupTestCases() []MarshallerTestCases {
	var testCases []MarshallerTestCases

	expectedPm := api.PaymentMethod{
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

	expectedPayment := api.Payment{
		Id:             "1",
		IdempotencyKey: "2",
		Amount:         "100",
		Currency:       "USD",
		MerchantId:     "3",
		PaymentMethod:  &expectedPm,
		PaymentStatus:  api.PaymentStatus_PaymentStatus_CREATED,
	}

	pm := model.PaymentMethod{
		PaymentMethodType: model.PaymentMethodType_CreditCard,
		PaymentMethodCreditCard: &model.PaymentMethodCreditCard{
			CardHolderName:   "hello",
			CreditCardNumber: "1234123412341234",
			ExpiryDate:       "12-04",
			Cvv:              "222",
			CreditCardType:   model.CreditCardType_Amex,
		},
	}

	inputPayment := model.Payment{
		Id:             "1",
		IdempotencyKey: "2",
		Amount:         "100",
		Currency:       "USD",
		MerchantId:     "3",
		PaymentMethod:  &pm,
		PaymentStatus:  model.PaymentStatus_Created,
	}

	testCases = append(testCases, MarshallerTestCases{
		input:    inputPayment,
		expected: expectedPayment,
	})
	return testCases
}

func (marshallerTest *MarshallerTestSuite) Test_Unmarshaller() {
	testCases := setupTestCases()

	for _, v := range testCases {

		resp, err := marshallerTest.marshaller.MarshalPayment(&v.input)

		marshallerTest.Assert().Nil(err)
		marshallerTest.Assert().Equal(*resp, v.expected)

	}
}
