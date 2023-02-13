package model

type PaymentStatus string
type PaymentMethodType string
type CreditCardType string

const (
	PaymentStatus_Created    PaymentStatus = "PaymentStatus_Created"
	PaymentStatus_Processing PaymentStatus = "PaymentStatus_Processing"
	PaymentStatus_Processed  PaymentStatus = "PaymentStatus_Processed"
	PaymentStatus_Settled    PaymentStatus = "PaymentStatus_Settled"
	PaymentStatus_Completed  PaymentStatus = "PaymentStatus_Completed"
	PaymentStatus_Failed     PaymentStatus = "PaymentStatus_Failed"
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
	CardHolderName   string
	CreditCardNumber string
	ExpiryDate       string
	Cvv              string
	CreditCardType   CreditCardType
}

type PaymentMethod struct {
	PaymentMethodType       PaymentMethodType
	PaymentMethodCreditCard *PaymentMethodCreditCard
}

type Payment struct {
	Id             string
	IdempotencyKey string
	Amount         string
	Currency       string
	MerchantId     string
	PaymentMethod  *PaymentMethod
	PaymentStatus  PaymentStatus
}
