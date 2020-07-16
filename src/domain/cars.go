package domain

import (
	"errors"
	"time"
)

type CarID string

var NilCarID CarID

type Car struct {
	ID                CarID      `json:"id,omitempty" db:"car_id"`
	UserID            *UserID    `json:"user_id,omitempty" db:"user_id"`
	CarValue          *int64     `json:"car_value,omitempty" db:"car_value"`
	CarMake           *string    `json:"car_make,omitempty" db:"car_make"`
	CarModel          *string    `json:"car_model,omitempty" db:"car_model"`
	YearOfManufacture *int32     `json:"year_of_manufacture,omitempty" db:"year_of_manufacture"`
	CarUse            *string    `json:"car_use,omitempty" db:"car_use"`
	PolicyPeriod      *time.Time `json:"policy_period,omitempty" db:"policy_period"`
	CreatedAt         *time.Time `json:"-,omitempty" db:"created_at"`
	UpdatedAt         *time.Time `json:"-,omitempty" db:"updated_at"`
	DeletedAt         *time.Time `json:"-,omitempty" db:"deleted_at"`
}

// Validate validate User structs fields
func (car *Car) Validate() error {
	if car.UserID == nil || (car.UserID != nil && len(*car.UserID) == 0) {
		return errors.New("UserID is Required")
	}
	if car.CarValue == nil {
		return errors.New("Car Value Is Required")
	}
	if car.CarMake == nil || (car.CarMake != nil && len(*car.CarMake) == 0) {
		return errors.New("Car Make Is Required")
	}
	if car.CarModel == nil || (car.CarModel != nil && len(*car.CarModel) == 0) {
		return errors.New("Car Model Is Required")
	}
	if car.YearOfManufacture == nil {
		return errors.New("Year Of Manufacture Is Required")
	}
	if car.CarUse == nil || (car.CarUse != nil && len(*car.CarUse) == 0) {
		return errors.New("Car Use Is Required")
	}
	if car.PolicyPeriod == nil || car.PolicyPeriod != nil {
		return errors.New("Policy Period Is Required")
	}

	return nil
}
