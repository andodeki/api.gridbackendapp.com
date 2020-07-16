package db

import (
	"context"
	"log"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/client"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/domain"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type paymentDBRepository struct{}

type PaymentDBRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) error
	UpdatePayment(ctx context.Context, payment *domain.Payment) error
	GetPaymentByID(ctx context.Context, paymentID domain.PaymentID) (*domain.Payment, error)
	ListPaymentByUserID(ctx context.Context, userID domain.UserID) ([]*domain.Payment, error)
	DeletePayment(ctx context.Context, payment *domain.Payment) (bool, error)
}

func NewPaymentDBRepository() PaymentDBRepository {
	return &paymentDBRepository{}
}

func (p *paymentDBRepository) CreatePayment(ctx context.Context, payment *domain.Payment) error {

	logger := logrus.WithField("func", "db_payment_repository.go -> CreatePayment()")
	log.Println(logger)

	rows, err := client.Conn.GetClient().NamedQueryContext(ctx, createPaymentQuery, payment)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return errors.Wrap(err, "Could Not Create Payment")
	}
	rows.Next()
	if err := rows.Scan(&payment.ID); err != nil {
		return err
	}
	return nil
}
func (p *paymentDBRepository) UpdatePayment(ctx context.Context, payment *domain.Payment) error {
	logger := logrus.WithField("func", "db_payment_repository.go -> UpdatePayment()")
	log.Println(logger)
	result, err := client.Conn.GetClient().NamedExecContext(ctx, updatePaymentQuery, payment)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("payment not found")
	}
	return nil
}
func (p *paymentDBRepository) GetPaymentByID(ctx context.Context, paymentID domain.PaymentID) (*domain.Payment, error) {

	logger := logrus.WithField("func", "db_payment_repository.go -> GetPaymentByID()")
	log.Println(logger)
	var payment domain.Payment
	if err := client.Conn.GetClient().GetContext(ctx, &payment, getPaymentByIDQuery, paymentID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Payment By ID")
	}
	return &payment, nil
}
func (p *paymentDBRepository) ListPaymentByUserID(ctx context.Context, userID domain.UserID) ([]*domain.Payment, error) {
	logger := logrus.WithField("func", "db_payment_repository.go -> ListPaymentByUserID()")
	log.Println(logger)

	var payment []*domain.Payment
	if err := client.Conn.GetClient().SelectContext(ctx, &payment, listPaymentByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Users Payment")
	}
	return payment, nil
}

func (p *paymentDBRepository) DeletePayment(ctx context.Context, payment *domain.Payment) (bool, error) {
	logger := logrus.WithField("func", "db_payment_repository.go -> DeletePayment()")
	log.Println(logger)

	result, err := client.Conn.GetClient().ExecContext(ctx, deletePaymentQuery, payment)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}
