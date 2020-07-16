package resterrors

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

// RestErr provide standard structure to handle errors between services
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// GenericError represent the error structure the Generic Error
type GenericError struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}
type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func NewError(msg string) error {
	return errors.New(msg)
}

// HandleErr is a func
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// NewBadRequestError a fucntion that displays bad request errors
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError a function that displays bad request errors
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalSeverError a function that displays bad request errors
func NewInternalSeverError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

// WriteError return a JSON error message and HTTP error status code
func WriteError(w http.ResponseWriter, code int, message string, data interface{}) {
	response := GenericError{
		Error: message,
		Code:  code,
		Data:  data,
	}
	WriteJSON(w, code, response)
}

// WriteJSON returns a JSON data and HTTP status code
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.WithError(err).Warn("Error Writing Response")
	}
}
