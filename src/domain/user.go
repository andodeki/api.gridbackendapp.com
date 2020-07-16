package domain

import (
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	StatusActive = "active"
)

//UserID is an identifier for User
type UserID string

type User struct {
	ID           UserID     `json:"id,omitempty" db:"user_id"`
	Username     string     `json:"username" db:"user_name"`
	Email        string     `json:"email" db:"email"`
	PasswordHash []byte     `json:"-" db:"password_hash"`
	Status       string     `json:"-" db:"status"`
	CreatedAt    *time.Time `json:"-" db:"created_at"`
	UpdatedAt    *time.Time `json:"-" db:"updated_at"`
	DeletedAt    *time.Time `json:"-" db:"deleted_at"`
}

type UserParameters struct {
	User
	SessionData
	Password string `json:"password"`
}

// Validate validate User structs fields
func (user *User) Validate() error {
	user.Username = strings.TrimSpace(strings.Title(user.Username))
	// user.LastName = strings.TrimSpace(strings.Title(user.LastName))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" || (user.Email != "" && len(user.Email) == 0) {
		return errors.New("Invalid Email Address")
	}
	// return nil
	// user.PasswordHash = strings.TrimSpace(user.PasswordHash)
	// if user.PasswordHash == "" {
	// 	return resterrors.NewBadRequestError("Invalid Password")
	// }
	return nil
}

// SetPassword updates a user's password
func (user *User) SetPassword(password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	user.PasswordHash = hash
	return nil
}

//CheckPassword verifies user's password
func (user *User) CheckPassword(password string) error {
	if hex.EncodeToString(user.PasswordHash) != "" && len(user.PasswordHash) == 0 {
		return errors.New("Password Not Set")
	}
	return bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
}

//HashPassword hashes a users's raw password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
