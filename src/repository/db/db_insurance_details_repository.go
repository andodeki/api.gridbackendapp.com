package db

import (
	"context"
	"log"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/client"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/domain"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type insuranceDetailsDBRepository struct{}

type InsuranceDetailsDBRepository interface {
	CreateInsuranceDetails(ctx context.Context, insurance *domain.InsuranceDetails) error
	UpdateInsuranceDetails(ctx context.Context, insurance *domain.InsuranceDetails) error
	GetInsuranceDetailsByID(ctx context.Context, insuranceID domain.InsuranceID) (*domain.InsuranceDetails, error)
	ListInsuranceDetailsByUserID(ctx context.Context, userID domain.UserID) ([]*domain.InsuranceDetails, error)
	ListInsuranceDetailsByCarID(ctx context.Context, carID domain.CarID) ([]*domain.InsuranceDetails, error)
	DeleteInsuranceDetails(ctx context.Context, insurance *domain.InsuranceDetails) (bool, error)
}

func NewInsuranceDetailsDBRepository() InsuranceDetailsDBRepository {
	return &insuranceDetailsDBRepository{}
}

func (i *insuranceDetailsDBRepository) CreateInsuranceDetails(ctx context.Context, insurance *domain.InsuranceDetails) error {
	logger := logrus.WithField("func", "db_insurance_details_repository.go -> CreateInsuranceDetails()")
	log.Println(logger)

	rows, err := client.Conn.GetClient().NamedQueryContext(ctx, createInsuranceDetailsQuery, insurance)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return errors.Wrap(err, "Could Not Create Insurance")
	}
	rows.Next()
	if err := rows.Scan(&insurance.ID); err != nil {
		return err
	}
	return nil
}

func (i *insuranceDetailsDBRepository) UpdateInsuranceDetails(ctx context.Context, insurance *domain.InsuranceDetails) error {
	logger := logrus.WithField("func", "db_insurance_details_repository.go -> UpdateInsuranceDetails()")
	log.Println(logger)
	result, err := client.Conn.GetClient().NamedExecContext(ctx, updateInsuranceDetailsQuery, insurance)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("insurance not found")
	}
	return nil
}
func (i *insuranceDetailsDBRepository) GetInsuranceDetailsByID(ctx context.Context, insuranceID domain.InsuranceID) (*domain.InsuranceDetails, error) {
	logger := logrus.WithField("func", "db_insurance_details_repository.go -> GetInsuranceDetailsByID()")
	log.Println(logger)
	var insuranceDetails domain.InsuranceDetails
	if err := client.Conn.GetClient().GetContext(ctx, &insuranceDetails, getInsuranceDetailsByIDQuery, insuranceID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Insurance By ID")
	}
	return &insuranceDetails, nil
}
func (i *insuranceDetailsDBRepository) ListInsuranceDetailsByUserID(ctx context.Context, userID domain.UserID) ([]*domain.InsuranceDetails, error) {
	logger := logrus.WithField("func", "db_car_repository.go -> ListInsuranceDetailsByUserID()")
	log.Println(logger)

	var insuranceDetails []*domain.InsuranceDetails
	if err := client.Conn.GetClient().SelectContext(ctx, &insuranceDetails, listInsuranceDetailsByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Users Insurance")
	}
	return insuranceDetails, nil
}
func (i *insuranceDetailsDBRepository) ListInsuranceDetailsByCarID(ctx context.Context, carID domain.CarID) ([]*domain.InsuranceDetails, error) {
	logger := logrus.WithField("func", "db_car_repository.go -> ListInsuranceDetailsByCarID()")
	log.Println(logger)

	var insuranceDetails []*domain.InsuranceDetails
	if err := client.Conn.GetClient().SelectContext(ctx, &insuranceDetails, listInsuranceDetailsByCarID, carID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Cars Insurance")
	}
	return insuranceDetails, nil
}
func (i *insuranceDetailsDBRepository) DeleteInsuranceDetails(ctx context.Context, insurance *domain.InsuranceDetails) (bool, error) {
	logger := logrus.WithField("func", "db_car_repository.go -> DeleteInsuranceDetails()")
	log.Println(logger)

	result, err := client.Conn.GetClient().ExecContext(ctx, deleteInsuranceDetailsQuery, insurance)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}
