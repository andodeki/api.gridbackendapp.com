package serverping

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	pong = "pong" + "\n"
)

var (
	PingHttp pingHttpInterface = &pingHttp{}
)

type pingHttpInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingHttp struct{}

// Ping checks if sever is up
func (c *pingHttp) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(pong)); err != nil {
		logrus.WithError(err).Debug("Error Ping.")
		log.Println(err)

	}
}

/*
// Ping checks if sever is up
func Ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Pong" + "\n")); err != nil {
		logrus.WithError(err).Debug("Error Ping.")

	}
}
*/
