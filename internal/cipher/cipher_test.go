package cipher_test

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/sarthakraheja/payments-service/internal/cipher"
	"github.com/sarthakraheja/payments-service/internal/model"

	"github.com/stretchr/testify/suite"
)

type CipherTestSuite struct {
	suite.Suite

	cipher cipher.Cipher
}

func NewCipherTestSuite(t *testing.T) CipherTestSuite {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	cipher := cipher.NewCipher(key)

	return CipherTestSuite{
		cipher: cipher,
	}
}

func TestCipherTestSuite(t *testing.T) {
	cipherTestSuire := NewCipherTestSuite(t)

	suite.Run(t, &cipherTestSuire)
}

func (c *CipherTestSuite) Test_Cipher() {
	inputString := "Hello, World!"
	input := []byte(inputString)

	encrypted, err := c.cipher.Encrypt(input)
	c.Assert().Nil(err)

	decrypted, err := c.cipher.Decrypt(encrypted)
	c.Assert().Nil(err)

	decryptedString := string(decrypted)

	c.Assert().Equal(inputString, decryptedString)

	expectedPm := &model.PaymentMethod{
		PaymentMethodType: model.PaymentMethodType_CreditCard,
		PaymentMethodCreditCard: &model.PaymentMethodCreditCard{
			CardHolderName:   "hello",
			CreditCardNumber: "1234123412341234",
			ExpiryDate:       "12-04",
			Cvv:              "222",
			CreditCardType:   model.CreditCardType_Amex,
		},
	}

	encodedPM, err := json.Marshal(expectedPm)
	c.Assert().Nil(err)

	encryptedPm, err := c.cipher.Encrypt(encodedPM)
	c.Assert().Nil(err)

	decryptedPm, err := c.cipher.Decrypt(encryptedPm)
	c.Assert().Nil(err)

	outputPm := &model.PaymentMethod{}
	json.Unmarshal(decryptedPm, outputPm)

	c.Assert().Equal(outputPm, expectedPm)
}
