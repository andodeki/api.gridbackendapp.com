package db

import (
	"context"
	"log"

	"github.com/andodeki/api.gridbackendapp.com/src/client"
	"github.com/andodeki/api.gridbackendapp.com/src/domain"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewRepository is a constructor that will create an object that represent the users.Repository interface
func NewCarRepository() CarDBRepository {
	return &carDBRepository{}
}

type carDBRepository struct{}

type CarDBRepository interface {
	CreateCar(ctx context.Context, car *domain.Car) error
	UpdateCar(ctx context.Context, car *domain.Car) error
	GetCarByID(ctx context.Context, carID domain.CarID) (*domain.Car, error)
	ListCarsByUserID(ctx context.Context, userID domain.UserID) ([]*domain.Car, error)
	DeleteCar(ctx context.Context, car *domain.Car) (bool, error)
}

func (c *carDBRepository) CreateCar(ctx context.Context, car *domain.Car) error {
	logger := logrus.WithField("func", "db_car_repository.go -> CreateCar()")
	log.Println(logger)

	rows, err := client.Conn.GetClient().NamedQueryContext(ctx, createCarQuery, car)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return errors.Wrap(err, "Could Not Create Car")
	}
	rows.Next()
	if err := rows.Scan(&car.ID); err != nil {
		return err
	}
	return nil
}

func (c *carDBRepository) UpdateCar(ctx context.Context, car *domain.Car) error {
	logger := logrus.WithField("func", "db_car_repository.go -> UpdateCar()")
	log.Println(logger)
	result, err := client.Conn.GetClient().NamedExecContext(ctx, updateCarQuery, car)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("car not found")
	}
	return nil
}

func (c *carDBRepository) GetCarByID(ctx context.Context, carID domain.CarID) (*domain.Car, error) {
	logger := logrus.WithField("func", "db_car_repository.go -> GetCarByID()")
	log.Println(logger)
	var car domain.Car
	if err := client.Conn.GetClient().GetContext(ctx, &car, getCarByIDQuery, carID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Car By ID")
	}
	return &car, nil
}

func (c *carDBRepository) ListCarsByUserID(ctx context.Context, userID domain.UserID) ([]*domain.Car, error) {
	logger := logrus.WithField("func", "db_car_repository.go -> ListCarsByUserID()")
	log.Println(logger)

	var cars []*domain.Car
	if err := client.Conn.GetClient().SelectContext(ctx, &cars, listCarsByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "Cant Get Users Cars")
	}
	return cars, nil
}
func (c *carDBRepository) DeleteCar(ctx context.Context, car *domain.Car) (bool, error) {
	logger := logrus.WithField("func", "db_car_repository.go -> ListCarsByUserID()")
	log.Println(logger)

	result, err := client.Conn.GetClient().ExecContext(ctx, deleteCarQuery, car)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}
