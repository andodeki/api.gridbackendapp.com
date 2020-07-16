package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/app/router"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/client"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/helper/logger"
	"github.com/sirupsen/logrus"
	// "github.com/andodeki/code/HA/api.griffins.com/src/app/router"
)

type Server struct {
	port string
	// Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

//Init all Vals
func (s *Server) Init(port string) {
	logger.Info("Initializing Server...")
	s.port = ":" + port
	// s.Db
}

// StartApplication is a function that start this service
func (s *Server) StartApplication() {
	r := router.NewRouter()
	r.Init()

	// postgres.Init()
	client.Init()
	logger.Info("Start Server on port " + s.port)

	newServer := &http.Server{
		Handler:           r.Router,
		Addr:              "0.0.0.0" + s.port,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		MaxHeaderBytes:    64e3,
	}

	// log.Fatal(newServer.ListenAndServe())
	go func() {
		if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
			os.Exit(-1)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logrus.Debug("Bye")
	logger.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	newServer.Shutdown(ctx)
}
