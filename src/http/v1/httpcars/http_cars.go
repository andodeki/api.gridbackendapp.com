package httpcars

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/domain"
	resterrors "github.com/andodeki/code/HA/api.gridbackendapp.com/src/helper/utils/rest_errors"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/repository/db"
	carService "github.com/andodeki/code/HA/api.gridbackendapp.com/src/services/carservice"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CarHandler interface {
	CreateCar(w http.ResponseWriter, r *http.Request)
	UpdateCar(w http.ResponseWriter, r *http.Request)
	GetCarByID(w http.ResponseWriter, r *http.Request)
	ListCarsByUserID(w http.ResponseWriter, r *http.Request)
	DeleteCar(w http.ResponseWriter, r *http.Request)
}

type carHandler struct {
	carservice carService.CarsServiceInterface
}

func NewCarHandler(service carService.CarsServiceInterface) CarHandler {
	return &carHandler{
		carservice: service,
	}
}

func (handler *carHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "http_user.go -> Create()") // Show Func Name in logs to track error faster
	log.Println(logger)

	vars := mux.Vars(r)
	userID := domain.UserID(vars["userID"])
	var car domain.Car
	//Load Parameters
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		logger.WithError(err).Warn("Could Not Decode Parameters")
		resterrors.WriteError(w, http.StatusBadRequest, "Could Not Decode Parameters", map[string]string{
			"error": err.Error(),
		})
		// return

		car.UserID = &userID

		ctx := r.Context()
		createdCar, err := handler.carservice.CreateCar(ctx, car)
		if err == db.ErrUserExist {
			logger.WithError(err).Warn("Car Already Exist.")
			resterrors.WriteError(w, http.StatusConflict, "Car Already Exist.", nil)
		} else if err != nil {
			// logger.WithError(err).Warn("Error Creating User.")
			resterrors.WriteError(w, http.StatusConflict, "Error Creating Car Details.", nil)
		}
		createdGetCar, err := handler.carservice.GetCarByID(ctx, &createdCar.ID)
		if err != nil {
			logger.WithError(err).Warn("Error Getting Car Details.")
			resterrors.WriteError(w, http.StatusConflict, "Error Getting Car Details.", nil)
			return
		}
		logger.WithField("carID", createdGetCar.ID).Info("Car Details Created")
		resterrors.WriteJSON(w, http.StatusCreated, createdGetCar)

	}
}
func (c *carHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "http_user.go -> UpdateCar()") // Show Func Name in logs to track error faster
	log.Println(logger)
	vars := mux.Vars(r)
	userID := domain.UserID(vars["userID"])
	var carRequest domain.Car

	//Load Parameters
	if err := json.NewDecoder(r.Body).Decode(&carRequest); err != nil {
		logger.WithError(err).Warn("Could Not Decode Parameters")
		resterrors.WriteError(w, http.StatusBadRequest, "Could Not Decode Parameters", map[string]string{
			"error": err.Error(),
		})
		// return

		carRequest.UserID = &userID
		ctx := r.Context()

		updatedCar, err := c.carservice.UpdateCar(ctx, &carRequest)
		if err != nil {
			logger.WithError(err).Warn("Error Updating Car Details.")
			resterrors.WriteError(w, http.StatusConflict, "Error Updating Car Details.", nil)
			return
		}
		logger.WithField("carID", updatedCar.ID).Info("Car Details Updated")
		resterrors.WriteJSON(w, http.StatusCreated, updatedCar)
	}
}
func (c *carHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "http_user.go -> GetCarByID()") // Show Func Name in logs to track error faster
	log.Println(logger)
}
func (c *carHandler) ListCarsByUserID(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "http_user.go -> ListCarsByUserID()") // Show Func Name in logs to track error faster
	log.Println(logger)
}
func (c *carHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "http_user.go -> DeleteCar()") // Show Func Name in logs to track error faster
	log.Println(logger)
}
