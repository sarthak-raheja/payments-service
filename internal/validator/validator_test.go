package validator_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/validator"

	"github.com/stretchr/testify/suite"
)

type ValidatorTestSuite struct {
	suite.Suite

	validator validator.Validator
}

func NewValidatorTestSuite(t *testing.T) ValidatorTestSuite {
	validator := validator.NewValidator()

	return ValidatorTestSuite{
		validator: validator,
	}
}

func TestHandlerTestSuite(t *testing.T) {
	validatorTestSuite := NewValidatorTestSuite(t)

	suite.Run(t, &validatorTestSuite)
}

type ValidatorTestCases struct {
	input         api.CreatePaymentRequest
	expectedError error
}

func setupTestCases() []ValidatorTestCases {
	var testCases []ValidatorTestCases

	validCreatePaymentRequest := api.CreatePaymentRequest{
		IdempotencyKey: "123",
		Amount:         "100.00",
		Currency:       "USD",
		MerchantId:     "123",
		PaymentMethod: &api.PaymentMethod{
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
		},
	}

	testCases = append(testCases, ValidatorTestCases{
		input:         validCreatePaymentRequest,
		expectedError: nil,
	})

	invalidAmountReq := validCreatePaymentRequest
	invalidAmountReq.Amount = "-12"

	testCases = append(testCases, ValidatorTestCases{
		input:         invalidAmountReq,
		expectedError: fmt.Errorf("payment amount is invalid"),
	})

	invalidCurrencyReq := validCreatePaymentRequest
	invalidCurrencyReq.Currency = "btc"

	testCases = append(testCases, ValidatorTestCases{
		input:         invalidCurrencyReq,
		expectedError: fmt.Errorf("currency not supported"),
	})

	return testCases
}

func (validatorTest *ValidatorTestSuite) Test_Validator() {
	testCases := setupTestCases()

	for _, v := range testCases {
		err := validatorTest.validator.ValidateCreatePayment(context.Background(), &v.input)

		if v.expectedError == nil {
			validatorTest.Assert().Nil(err)
		}

		validatorTest.Assert().Equal(v.expectedError, err)
	}
}
