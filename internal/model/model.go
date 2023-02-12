package model

import "google.golang.org/protobuf/types/known/timestamppb"

type PaymentEventType string
type PaymentMethodType string
type CreditCardType string

const (
	PaymentEventType_Created    PaymentEventType = "PaymentEventType_Created"
	PaymentEventType_Processing PaymentEventType = "PaymentEventType_Processing"
	PaymentEventType_Processed  PaymentEventType = "PaymentEventType_Processed"
	PaymentEventType_Settled    PaymentEventType = "PaymentEventType_Settled"
	PaymentEventType_Completed  PaymentEventType = "PaymentEvenetType_Completed"
	PaymentEventType_Failed     PaymentEventType = "PaymentEventType_Failed"
)

const (
	PaymentMethodType_CreditCard PaymentMethodType = "PaymentMethodType_CreditCard"
)

const (
	CreditCardType_Amex       CreditCardType = "CreditCardType_Amex"
	CreditCardType_Visa       CreditCardType = "CreditCardType_Visa"
	CreditCardType_MasterCard CreditCardType = "CreditCardType_MasterCard"
)

type PaymentMethodCreditCard struct {
	cardHolderName   string
	creditCardNumber string
	expiryDate       string
	cvv              string
	CreditCardType   CreditCardType
}

type PaymentMethod struct {
	PaymentMethodType       PaymentMethodType
	PaymentMethodCreditCard PaymentMethodCreditCard
}

type Payment struct {
	Id             string
	IdempotencyKey string
	Amount         string
	Currency       string
	MerchantId     string
	PaymentEvent   []*PaymentEvent
	PaymentMethod  *PaymentMethod
}

type PaymentEvent struct {
	PaymentId        string
	IdempotencyKey   string
	Timestamp        timestamppb.Timestamp
	PaymentEventType PaymentEventType
}
