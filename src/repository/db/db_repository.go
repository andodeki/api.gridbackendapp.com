package db

import (
	"context"
	"log"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/client"

	usersDomain "github.com/andodeki/code/HA/api.gridbackendapp.com/src/domain"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewRepository is a constructor that will create an object that represent the users.Repository interface
func NewUserRepository() DBRepository {
	return &dbRepository{}
}

type dbRepository struct{}

// DBRepository is an interface
type DBRepository interface {
	CreateUser(ctx context.Context, user *usersDomain.User) error
	GetUserByID(ctx context.Context, userID *usersDomain.UserID) (*usersDomain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*usersDomain.User, error)
}

// UniqueViolation Postgres error string for a unique index violation
const UniqueViolation = "unique violation"

var ErrUserExist = errors.New("User With Email Exist")

func (r *dbRepository) CreateUser(ctx context.Context, user *usersDomain.User) error {
	logger := logrus.WithField("func", "db_repository.go -> CreateUser()")
	log.Println(logger)
	rows, err := client.Conn.GetClient().NamedQueryContext(ctx, createUserQuery, user)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code.Name() == UniqueViolation {
				if pqError.Constraint == "user_email" {
					return ErrUserExist // Error User Exist
				}
			}
		}
		return errors.Wrap(err, "Could Not Create User")
	}

	rows.Next()
	if err := rows.Scan(&user.ID); err != nil {
		return errors.Wrap(err, "Could Not Get Created User ID")
	}
	return nil
}

func (r *dbRepository) GetUserByID(ctx context.Context, userID *usersDomain.UserID) (*usersDomain.User, error) {
	var user usersDomain.User
	if err := client.Conn.GetClient().GetContext(ctx, &user, getUserByIDQuery, userID); err != nil {
		return nil, errors.Wrap(err, "Cant Get User By ID")
	}
	return &user, nil
	// return nil, nil
}

func (r *dbRepository) GetUserByEmail(ctx context.Context, email string) (*usersDomain.User, error) {
	var user usersDomain.User
	if err := client.Conn.GetClient().GetContext(ctx, &user, getUserByEmailQuery, email); err != nil {
		return nil, errors.Wrap(err, "Cant Get User By Email")
	}
	return &user, nil
}

/*
func (m *dbRepository) Fetch(ctx context.Context) (res []domain.Book, err error) {
	var books []domain.Book
	m.Conn.Find(&books)

	return books, nil
}
func (m *dbRepository) GetByID(ctx context.Context, id string) (res domain.Book, err error) {
	var book domain.Book
	m.Conn.Where("id = ?", id).First(&book)
	// if err := m.Conn.Where("id = ?", id).First(&book).Error; err != nil {
	// 	return err, nil
	// }

	return book, nil
}
*/
