package userservice

import (
	"context"
	"errors"
	"log"

	"github.com/andodeki/api.gridbackendapp.com/src/repository/db"

	usersDomain "github.com/andodeki/api.gridbackendapp.com/src/domain"

	"github.com/sirupsen/logrus"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user usersDomain.UserParameters) (*usersDomain.User, error)
	GetUserByID(ctx context.Context, userid *usersDomain.User) (*usersDomain.User, error)
	GetUserByEmail(ctx context.Context, credentials usersDomain.Credentials) (*usersDomain.User, error)
}

type usersService struct {
	dbRepo db.DBRepository
}

//NewService is a constructor
func NewUserService(dbRepo db.DBRepository) UserServiceInterface {
	return &usersService{
		dbRepo: dbRepo,
	}
}

var dbUser usersDomain.User

// CreateUser creates the user in various services
func (s *usersService) CreateUser(ctx context.Context, user usersDomain.UserParameters) (*usersDomain.User, error) {
	logger := logrus.WithField("func", "users_service.go -> CreateUser()") // Show Func Name in logs to track error faster

	if err := user.Validate(); err != nil {
		return nil, err
	}
	hashed, err := usersDomain.HashPassword(user.Password)
	if err != nil {
		logger.WithError(err).Warn("Could Not Hash Password.")
		// return
	}

	newUser := &usersDomain.User{
		Username:     user.Username,
		Email:        user.Email,
		Status:       usersDomain.StatusActive,
		PasswordHash: hashed,
	}
	if err := s.dbRepo.CreateUser(ctx, newUser); err != nil {
		logger.WithError(err).Warn("User Already Exist.")
	}

	return newUser, nil
	// return nil
}

func (s *usersService) GetUserByID(ctx context.Context, userid *usersDomain.User) (*usersDomain.User, error) {
	logger := logrus.WithField("func", "users_service.go -> GetUserByID()") // Show Func Name in logs to track error faster

	log.Printf("User Id is: %s", userid.ID)
	if len(userid.ID) == 0 {
		return nil, errors.New("User ID is Null")
	}
	passeduserID, err := s.dbRepo.GetUserByID(ctx, &userid.ID)
	if err != nil {
		logger.WithError(err).Warn("Error Getting User.")
	}
	return passeduserID, nil
}

func (s *usersService) GetUserByEmail(ctx context.Context, credentials usersDomain.Credentials) (*usersDomain.User, error) {
	logger := logrus.WithField("func", "users_service.go -> GetUserByEmail()") // Show Func Name in logs to track error faster

	hashed, err := usersDomain.HashPassword(credentials.Password)
	if err != nil {
		logger.WithError(err).Warn("Could Not Hash Password.")
		// return
	}
	dao := &usersDomain.User{
		Email:        credentials.Email,
		PasswordHash: hashed,
	}
	if err := dao.CheckPassword(credentials.Password); err != nil {
		logger.WithError(err).Warn("Error Login in.")
		// resterrors.WriteError(w, http.StatusConflict, "Invalid credentials.", nil)
		// return
	}
	passeduserEmail, err := s.dbRepo.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		logger.WithError(err).Warn("Error Getting User.")
	}
	return passeduserEmail, nil
}
