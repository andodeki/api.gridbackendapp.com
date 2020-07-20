package httpusers

import (
	"encoding/json"
	"net/http"

	users "github.com/andodeki/api.gridbackendapp.com/src/domain"
	resterrors "github.com/andodeki/api.gridbackendapp.com/src/helper/utils/rest_errors"
	"github.com/andodeki/api.gridbackendapp.com/src/repository/db"
	usersService "github.com/andodeki/api.gridbackendapp.com/src/services/userservice"

	"github.com/sirupsen/logrus"
)

type UserHandler interface {
	// GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service usersService.UserServiceInterface
}

// NewHandler a controller func
func NewUserHandler(service usersService.UserServiceInterface) UserHandler {
	return &userHandler{
		service: service,
	}
}

// Create creates users
func (handler *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userParameters users.UserParameters
	logger := logrus.WithField("func", "http_user.go -> Create()") // Show Func Name in logs to track error faster
	//Load Parameters
	if err := json.NewDecoder(r.Body).Decode(&userParameters); err != nil {
		logger.WithError(err).Warn("Could Not Decode Parameters")
		resterrors.WriteError(w, http.StatusBadRequest, "Could Not Decode Parameters", map[string]string{
			"error": err.Error(),
		})
		// return
	}
	logger = logger.WithFields(logrus.Fields{
		"email":    userParameters.Email,
		"username": userParameters.Username,
		"password": userParameters.Password,
	})

	// // newUser := &users.User{
	// // 	Email:        userParameters.Email,
	// // 	PasswordHash: &hashed,
	// // }

	ctx := r.Context()
	createdUser, err := handler.service.CreateUser(ctx, userParameters)
	if err == db.ErrUserExist {
		logger.WithError(err).Warn("User Already Exist.")
		resterrors.WriteError(w, http.StatusConflict, "User Already Exist.", nil)
	} else if err != nil {
		// logger.WithError(err).Warn("Error Creating User.")
		resterrors.WriteError(w, http.StatusConflict, "Error Creating User.", nil)
	}
	createdGetUser, err := handler.service.GetUserByID(ctx, createdUser)
	if err != nil {
		logger.WithError(err).Warn("Error Getting User Details.")
		resterrors.WriteError(w, http.StatusConflict, "Error Getting User Details.", nil)
		return
	}
	logger.Info("User Details Created")
	resterrors.WriteJSON(w, http.StatusCreated, createdGetUser)

	// // c.JSON(http.StatusCreated, result) //.Marshall(c.GetHeader("X-Public") == "true"))

}

// Login  provides user login
func (handler *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "http_user.go -> Login()") // Show Func Name in logs to track error faster
	var credentials users.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		logger.WithError(err).Warn("Could Not Decode Parameters")
		resterrors.WriteError(w, http.StatusBadRequest, "Could Not Decode Parameters", map[string]string{
			"error": err.Error(),
		})
		return
	}
	logger = logger.WithFields(logrus.Fields{
		"email": credentials.Email,
	})

	ctx := r.Context()
	user, err := handler.service.GetUserByEmail(ctx, credentials)
	if err != nil {
		logger.WithError(err).Warn("Error Login in.")
		resterrors.WriteError(w, http.StatusConflict, "Invalid credentials.", nil)
		return
	}

	logger.WithField("userID", user.ID).Debug("User Logged in.")
	resterrors.WriteJSON(w, http.StatusOK, user)
}
