package domain

import "fmt"

// Credentials used to login
type Credentials struct {
	SessionData

	Email    string `json:"email"`
	Password string `json:"password"`
}

//NilPrincipal is an uninitialised Principal
var NilPrincipal Principal

// Principal is an authenticated entity
type Principal struct {
	UserID UserID `json:"userID, omitempty"`
}

func (p Principal) String() string {
	if p.UserID != "" {
		return fmt.Sprintf("UserID[%s]", p.UserID)
	}
	return "(none)"

}
