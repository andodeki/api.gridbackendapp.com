package domain

import (
	"errors"
	"time"
)

type InsuranceID string

var NilInsuranceID InsuranceID

/*
insurance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_id UUID NOT NULL REFERENCES cars,
    insurance_type TEXT NOT NULL,
    basic_cover JSON NOT NULL,
    add_ons JSON NOT NULL,
    start_date DATETIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
*/
type InsuranceDetails struct {
	ID            InsuranceID `json:"id,omitempty" db:"insurance_id"`
	UserID        *UserID     `json:"user_id,omitempty" db:"user_id"`
	CarID         *CarID      `json:"car_id,omitempty" db:"car_id"`
	InsuranceType *string     `json:"insurance_type,omitempty" db:"insurance_type"`
	BasicCover    *string     `json:"basic_cover,omitempty" db:"basic_cover"`
	AddOns        *int32      `json:"add_ons,omitempty" db:"add_ons"`
	StartDate     *time.Time  `json:"start_date,omitempty" db:"start_date"`
	CreatedAt     *time.Time  `json:"-,omitempty" db:"created_at"`
	UpdatedAt     *time.Time  `json:"-,omitempty" db:"updated_at"`
	DeletedAt     *time.Time  `json:"-,omitempty" db:"deleted_at"`
}

// Validate validate User structs fields
func (i *InsuranceDetails) Validate() error {
	if i.UserID == nil || (i.UserID != nil && len(*i.UserID) == 0) {
		return errors.New("UserID is Required")
	}
	if i.CarID == nil {
		return errors.New("CarID Is Required")
	}
	if i.InsuranceType == nil || (i.InsuranceType != nil && len(*i.InsuranceType) == 0) {
		return errors.New("Insurance Type Is Required")
	}
	if i.BasicCover == nil || (i.BasicCover != nil && len(*i.BasicCover) == 0) {
		return errors.New("Basic Cover Is Required")
	}
	if i.StartDate == nil {
		return errors.New("Start Date Is Required")
	}

	return nil
}
