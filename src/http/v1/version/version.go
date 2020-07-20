package version

import (
	"encoding/json"
	"net/http"

	"github.com/andodeki/api.gridbackendapp.com/src/config/config"

	"github.com/sirupsen/logrus"
	// "github.com/gin-gonic/gin"
)

// API for returning version
// when server starts we set version and then use use it if neccessary

// ServerVersion represents the server version
type ServerVersion struct {
	Version string `json:"version"`
}

// Mashalled JSON
var versionJSON []byte

var (
	VersionHttp versionHttpInterface = &versionHttp{}
)

type versionHttpInterface interface {
	VersionHandler(w http.ResponseWriter, _ *http.Request)
}

type versionHttp struct{}

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})

	if err != nil {
		panic(err)
	}

}

// VersionHandler serves version information
func (v *versionHttp) VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error Writing Version.")

	}
}

/*
// Ping checks if sever is up
func Ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Pong")); err != nil {
		logrus.WithError(err).Debug("Error Writing Version.")

	}
}
*/
