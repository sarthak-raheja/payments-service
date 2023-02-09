package repository

import (
	"database/sql"

	"github.com/sarthakraheja/payments-service/internal/model"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	GetPayment(paymentId string) (*model.Payment, error)
	CreatePayment(payment *model.Payment) (*model.Payment, error)
	UpdatePayment(payment *model.Payment) (*model.Payment, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetPayment(paymentId string) (*model.Payment, error) {
	return nil, nil
}

func (r *repository) CreatePayment(payment *model.Payment) (*model.Payment, error) {
	return nil, nil
}

func (r *repository) UpdatePayment(payment *model.Payment) (*model.Payment, error) {
	return nil, nil
}
