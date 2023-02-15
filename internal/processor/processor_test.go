package processor_test

import (
	"fmt"
	"testing"

	"github.com/sarthakraheja/payments-service/internal/model"
	"github.com/sarthakraheja/payments-service/internal/processor"
	"github.com/sarthakraheja/payments-service/test/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ProcessorTestSuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	repo   *mocks.MockRepository
	router *mocks.MockAcquiringBankRouter

	processor processor.Processor
}

func NewProcessorTestSuite(t *testing.T) ProcessorTestSuite {
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockRepository(ctrl)
	router := mocks.NewMockAcquiringBankRouter(ctrl)

	processor := processor.NewProcessor(repo, router)

	return ProcessorTestSuite{
		ctrl:      ctrl,
		repo:      repo,
		router:    router,
		processor: processor,
	}
}

func TestProcessorTestSuite(t *testing.T) {
	processorTestSuite := NewProcessorTestSuite(t)

	suite.Run(t, &processorTestSuite)
}

func (p *ProcessorTestSuite) Test_ProcessPayment() {

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

	req := &processor.ProcessPaymentRequest{
		IdempotencyKey: "2",
		Amount:         "100",
		Currency:       "USD",
		MerchantId:     "3",
		PaymentMethod:  &pm,
	}

	p.repo.EXPECT().CreatePayment(gomock.Any()).Return(nil, fmt.Errorf("duplicate idem"))

	resp, err := p.processor.ProcessPayment(req)
	p.Assert().NotNil(err)
	p.Assert().Nil(resp)
}

func (p *ProcessorTestSuite) Test_AcquiringBankFailure() {

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

	req := &processor.ProcessPaymentRequest{
		IdempotencyKey: "2",
		Amount:         "100",
		Currency:       "USD",
		MerchantId:     "3",
		PaymentMethod:  &pm,
	}

	respPayment := &model.Payment{
		Id:             "1",
		IdempotencyKey: "2",
		Amount:         "100",
		Currency:       "USD",
		MerchantId:     "3",
		PaymentMethod:  &pm,
		PaymentStatus:  model.PaymentStatus_Created,
	}

	p.repo.EXPECT().CreatePayment(gomock.Any()).Return(respPayment, nil)
	p.router.EXPECT().GetAcquiringBank(gomock.Any()).Return(nil, fmt.Errorf("could not find acquiring bank"))

	resp, err := p.processor.ProcessPayment(req)

	p.Assert().NotNil(err)
	p.Assert().Nil(resp)
}

func (p *ProcessorTestSuite) Test_GetPayment() {

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

	respPayment := &model.Payment{
		Id:             "1",
		IdempotencyKey: "2",
		Amount:         "100",
		Currency:       "USD",
		MerchantId:     "3",
		PaymentMethod:  &pm,
		PaymentStatus:  model.PaymentStatus_Created,
	}

	expectedResponce := &processor.GetPaymentResponse{
		Payment: respPayment,
	}
	req := &processor.GetPaymentRequest{
		Id:         "1",
		MerchantId: "3",
	}

	p.repo.EXPECT().GetPayment(gomock.Any()).Return(respPayment, nil)
	res, err := p.processor.GetPayment(req)

	p.Assert().Nil(err)
	p.Assert().Equal(res, expectedResponce)

}
