package domain

// CREATE TABLE payments(
//     payment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//     user_id UUID NOT NULL REFERENCES users,
//     insurance_id UUID NOT NULL REFERENCES insurance_details,
//     car_id UUID NOT NULL REFERENCES cars,
//     payment_frequency TEXT NOT NULL,
//     min_deposit INTEGER NOT NULL,
//     monthly_repayment INTEGER NOT NULL,
//     total_interest INTEGER NOT NULL,
//     payment_mode TEXT NOT NULL,
//     currency TEXT,
//     status TEXT NOT NULL,
//     payment_at TIMESTAMP NOT NULL DEFAULT NOW(),
//     updated_at TIMESTAMP NOT NULL DEFAULT NOW()

// );
import (
	"errors"
	"time"
)

const (
	StatusPaid = "paid"
)

type PaymentID string

var NilPaymentID PaymentID

type Payment struct {
	ID               PaymentID    `json:"id,omitempty" db:"insurance_id"`
	UserID           *UserID      `json:"user_id,omitempty" db:"user_id"`
	CarID            *CarID       `json:"car_id,omitempty" db:"car_id"`
	InsuranceID      *InsuranceID `json:"insurance_id,omitempty" db:"insurance_id"`
	PaymentFrequency *string      `json:"payment_frequency,omitempty" db:"payment_frequency"`
	MinDeposit       *int32       `json:"min_deposit,omitempty" db:"min_deposit"`
	RefundAmount     *int32       `json:"refund_amount,omitempty" db:"min_deposit"`
	MonthlyRepayment *int32       `json:"monthly_repayment,omitempty" db:"refund_amount"`
	TotalInterest    *int32       `json:"total_interest,omitempty" db:"total_interest"`
	PaymentMode      *string      `json:"payment_mode,omitempty" db:"payment_mode"`
	Currency         *string      `json:"currency,omitempty" db:"currency"`
	Status           *string      `json:"-,omitempty" db:"status"`
	CreatedAt        *time.Time   `json:"-,omitempty" db:"created_at"`
	UpdatedAt        *time.Time   `json:"-,omitempty" db:"updated_at"`
	DeletedAt        *time.Time   `json:"-,omitempty" db:"deleted_at"`
}

// Validate validate User structs fields
func (i *Payment) Validate() error {
	if i.UserID == nil || (i.UserID != nil && len(*i.UserID) == 0) {
		return errors.New("UserID is Required")
	}
	if i.CarID == nil {
		return errors.New("CarID Is Required")
	}
	if i.InsuranceID == nil {
		return errors.New("InsuranceID Is Required")
	}
	if i.PaymentFrequency == nil || (i.PaymentFrequency != nil && len(*i.PaymentFrequency) == 0) {
		return errors.New("Payment Frequency Is Required")
	}
	if i.MinDeposit == nil {
		return errors.New("Minimun Deposit Is Required")
	}
	if i.MonthlyRepayment == nil {
		return errors.New("Monthly Repayment Is Required")
	}
	if i.TotalInterest == nil {
		return errors.New("Total Interest Is Required")
	}
	if i.PaymentMode == nil || (i.PaymentMode != nil && len(*i.PaymentMode) == 0) {
		return errors.New("Payment Mode Is Required")
	}
	return nil
}
