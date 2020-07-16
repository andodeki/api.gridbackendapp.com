package carservice

import (
	"context"
	"errors"
	"log"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/domain"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/repository/db"
	"github.com/sirupsen/logrus"
)

type CarsServiceInterface interface {
	CreateCar(ctx context.Context, car domain.Car) (*domain.Car, error)
	UpdateCar(ctx context.Context, carid *domain.Car) (*domain.Car, error)
	GetCarByID(ctx context.Context, carid *domain.CarID) (*domain.Car, error)
	ListCarByID(ctx context.Context, credentials domain.Credentials) ([]*domain.Car, error)
	DeleteCar(ctx context.Context, car domain.Car) (bool, error)
}

type carsService struct {
	carsdbRepo db.CarDBRepository
}

//NewService is a constructor
func NewCarService(carsdbRepo db.CarDBRepository) CarsServiceInterface {
	return &carsService{
		carsdbRepo: carsdbRepo,
	}
}

func (c *carsService) CreateCar(ctx context.Context, car domain.Car) (*domain.Car, error) {
	logger := logrus.WithField("func", "users_service.go -> CreateUser()") // Show Func Name in logs to track error faster

	if err := car.Validate(); err != nil {
		return nil, err
	}

	newCar := &domain.Car{
		UserID:            car.UserID,
		CarValue:          car.CarValue,
		CarMake:           car.CarMake,
		CarModel:          car.CarModel,
		YearOfManufacture: car.YearOfManufacture,
		CarUse:            car.CarUse,
		PolicyPeriod:      car.PolicyPeriod,
	}

	if err := c.carsdbRepo.CreateCar(ctx, newCar); err != nil {
		logger.WithError(err).Warn("Car Already Exist.")
	}
	return nil, nil
}
func (c *carsService) UpdateCar(ctx context.Context, carDetails *domain.Car) (*domain.Car, error) {
	logger := logrus.WithField("func", "users_service.go -> CreateUser()") // Show Func Name in logs to track error faster

	carDetailsFromDB, err := c.carsdbRepo.GetCarByID(ctx, carDetails.ID)
	if err != nil {
		logger.WithError(err).Warn("could not fetch car details fron db.")
	}

	if carDetails.CarMake != nil || len(*carDetails.CarMake) != 0 {
		carDetailsFromDB.CarMake = carDetails.CarMake
	}
	if carDetails.CarValue != nil {
		carDetailsFromDB.CarValue = carDetails.CarValue
	}
	if carDetails.CarModel != nil || len(*carDetails.CarModel) != 0 {
		carDetailsFromDB.CarModel = carDetails.CarModel
	}
	if carDetails.YearOfManufacture != nil || len(*carDetails.CarMake) != 0 {
		carDetailsFromDB.YearOfManufacture = carDetails.YearOfManufacture
	}
	if carDetails.CarUse != nil || len(*carDetails.CarUse) != 0 {
		carDetailsFromDB.CarUse = carDetails.CarUse
	}
	if carDetails.PolicyPeriod != nil || carDetails.PolicyPeriod != nil {
		carDetailsFromDB.PolicyPeriod = carDetails.PolicyPeriod
	}
	if err := c.carsdbRepo.UpdateCar(ctx, carDetailsFromDB); err != nil {
		logger.WithError(err).Warn("Car Details Not Updated.")
	}
	return carDetailsFromDB, nil
}

func (c *carsService) GetCarByID(ctx context.Context, carid *domain.CarID) (*domain.Car, error) {
	logger := logrus.WithField("func", "users_service.go -> GetUserByID()") // Show Func Name in logs to track error faster

	log.Printf("User Id is: %s", carid)
	if carid == nil {
		return nil, errors.New("User ID is Null")
	}
	passedcarID, err := c.carsdbRepo.GetCarByID(ctx, *carid)
	if err != nil {
		logger.WithError(err).Warn("Error Getting User.")
	}
	return passedcarID, nil

}
func (c *carsService) ListCarByID(ctx context.Context, credentials domain.Credentials) ([]*domain.Car, error) {

	return nil, nil
}
func (c *carsService) DeleteCar(ctx context.Context, car domain.Car) (bool, error) {

	return false, nil
}
