package validator

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	GBP = "gbp"
	USD = "usd"
	EUR = "eur"
)

func (v *validator) ValidateCreatePayment(ctx context.Context, req *api.CreatePaymentRequest) error {
	idem := req.GetIdempotencyKey()
	if idem == "" {
		return fmt.Errorf("could not find idem")
	}

	marchantId := req.GetMerchantId()
	if marchantId == "" {
		return fmt.Errorf("could not merchant idem")
	}

	currency := req.GetCurrency()
	if currency == "" {
		return fmt.Errorf("currency not provided")
	}

	err := v.validateCurrency(currency)
	if err != nil {
		return fmt.Errorf("currency not supported")
	}

	amount, err := strconv.ParseFloat(req.GetAmount(), 64)
	if err != nil {
		return fmt.Errorf("unable to parse amount")
	}

	pm := req.GetPaymentMethod()
	if pm == nil {
		return fmt.Errorf("could not fetch payment method")
	}

	err = v.validatePaymentMethod(pm)
	if err != nil {
		return err
	}

	if amount <= 0 {
		return fmt.Errorf("payment amount is invalid")
	}

	return nil
}

func (v *validator) validateCurrency(currency string) error {
	supportedCurrency := map[string]bool{
		USD: true,
		GBP: true,
		EUR: true,
	}

	_, found := supportedCurrency[strings.ToLower(currency)]
	if !found {
		return grpc.Errorf(codes.InvalidArgument, "currency not supported: %v", currency)
	}
	return nil
}

func (v *validator) validatePaymentMethod(pm *api.PaymentMethod) error {
	pmType := pm.GetPaymentMethodType()
	if pmType == api.PaymentMethodType_PaymentMethodType_UNSPECIFIED {
		return fmt.Errorf("unspecified payment method type")
	}

	creditCardDetails := pm.GetPaymentMethodCreditCardDetails()

	return v.validateCreditCardDetails(creditCardDetails)
}

func (v *validator) validateCreditCardDetails(cc *api.PaymentMethodCreditCardDetails) error {
	if cc == nil {
		return fmt.Errorf("could not find credit card details")
	}

	creditCardNumber := cc.GetCreditCardNumber()
	if len(creditCardNumber) < 16 {
		return fmt.Errorf("invalid credit card details: credit card number")
	}

	cvv := cc.GetCvv()

	if len(cvv) != 3 {
		return fmt.Errorf("invalid credit card details: cvv")
	}

	creditCardType := cc.GetCreditCardType()
	if creditCardType == api.CreditCardType_CreditCardType_UNSPECIFIED {
		return fmt.Errorf("invalid credit card type")
	}

	return nil
}
