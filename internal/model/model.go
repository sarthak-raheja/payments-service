package model

type PaymentStatus string
type PaymentMethodType string
type CreditCardType string

const (
	PaymentStatus_Created    PaymentStatus = "created"
	PaymentStatus_Processing PaymentStatus = "processing"
	PaymentStatus_Processed  PaymentStatus = "processed"
	PaymentStatus_Settled    PaymentStatus = "settled"
	PaymentStatus_Completed  PaymentStatus = "completed"
	PaymentStatus_Failed     PaymentStatus = "failed"
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
	CardHolderName   string         `json:"card_holder_name"`
	CreditCardNumber string         `json:"credit_card_number"`
	ExpiryDate       string         `json:"expiry_date"`
	Cvv              string         `json:"cvv"`
	CreditCardType   CreditCardType `json:"credit_card_type"`
}

type PaymentMethod struct {
	PaymentMethodType       PaymentMethodType        `json:"payment_method_type"`
	PaymentMethodCreditCard *PaymentMethodCreditCard `json:"payment_method_credit_card,omitempty"`
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
