package repository

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sarthakraheja/payments-service/internal/cipher"
	"github.com/sarthakraheja/payments-service/internal/model"
)

type repository struct {
	db     *sql.DB
	cipher cipher.Cipher
}

type Repository interface {
	GetPayment(paymentId string) (*model.Payment, error)
	CreatePayment(payment *model.Payment) (*model.Payment, error)
	UpdatePaymentStatus(paymentId string, status model.PaymentStatus) error
}

func NewRepository(db *sql.DB, cipher cipher.Cipher) Repository {
	return &repository{
		db:     db,
		cipher: cipher,
	}
}

func (r *repository) GetPayment(paymentId string) (*model.Payment, error) {

	paymentIdTable, _ := strconv.Atoi(paymentId)

	sqlStatement := fmt.Sprintf(`SELECT * FROM payments where ID=%v`, paymentIdTable)
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("could not find payment")
	}

	var payment *model.Payment
	// fetch table
	for rows.Next() {
		var (
			id             string
			idempotencyKey string
			merchantId     string
			amount         string
			currency       string
			paymentStatus  string
			paymentMethod  string
			createdAt      string
			updatedAt      string
		)

		if err := rows.Scan(&id, &idempotencyKey, &merchantId, &amount, &currency, &paymentStatus, &paymentMethod, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		encryptedPm, _ := hex.DecodeString(paymentMethod)
		decryptedPm, err := r.cipher.Decrypt(encryptedPm)
		if err != nil {
			return nil, err
		}

		pm := &model.PaymentMethod{}

		json.Unmarshal(decryptedPm, pm)

		payment = &model.Payment{
			Id:             id,
			IdempotencyKey: idempotencyKey,
			MerchantId:     merchantId,
			Amount:         amount,
			Currency:       currency,
			PaymentMethod:  pm,
			PaymentStatus:  model.PaymentStatus(paymentStatus),
		}
	}
	if payment == nil {
		return nil, fmt.Errorf("Could not find payment with ID:%v", paymentId)
	}

	return payment, nil
}

func (r *repository) CreatePayment(payment *model.Payment) (*model.Payment, error) {

	jsonPm, err := json.Marshal(payment.PaymentMethod)
	if err != nil {
		return nil, err
	}
	encryptedPm, err := r.cipher.Encrypt(jsonPm)
	if err != nil {
		return nil, err
	}

	hexPm := hex.EncodeToString(encryptedPm)

	sqlStatement := fmt.Sprintf(`INSERT INTO payments (idempotency_key,merchant_id, amount,currency,payment_status,payment_method) VALUES (%v,%v,%v,'%v','%v','%v') RETURNING id`, payment.IdempotencyKey, payment.MerchantId, payment.Amount, payment.Currency, payment.PaymentStatus, hexPm)

	id := 0
	err = r.db.QueryRow(sqlStatement).Scan(&id)
	if err != nil {
		return nil, err
	}

	payment.Id = fmt.Sprintf("%v", id)

	return payment, nil
}

func (r *repository) UpdatePaymentStatus(paymentId string, paymentStatus model.PaymentStatus) error {
	paymentIdTable, _ := strconv.Atoi(paymentId)
	sqlStatement := fmt.Sprintf(`UPDATE payments SET payment_status='%v' WHERE ID=%v`, paymentStatus, paymentIdTable)

	_, err := r.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}
