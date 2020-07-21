package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/andodeki/api.gridbackendapp.com/src/app/router"
	"github.com/andodeki/api.gridbackendapp.com/src/client"
	"github.com/andodeki/api.gridbackendapp.com/src/helper/logger"
	// "github.com/andodeki/api.griffins.com/src/app/router"
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

	wg := sync.WaitGroup{}
	detectSignal := checkStopOSSignals(&wg)
	for !(*detectSignal) {
		if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
			os.Exit(-1)
		}
	}

	logger.Info("Exit signal triggered, writing data... ")
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	// defer cancel()
	// newServer.Shutdown(ctx)
	wg.Wait()
	logger.Info("Exiting program...")

	// log.Fatal(newServer.ListenAndServe())

	// go func() {
	// 	if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatal(err)
	// 		os.Exit(-1)
	// 	}
	// }()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c
	// logrus.Debug("Bye")
	// logger.Info("Shutting down...")
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	// defer cancel()
	// newServer.Shutdown(ctx)
}

func checkStopOSSignals(wg *sync.WaitGroup) *bool {
	Signal := false
	go func(s *bool) {
		wg.Add(1)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		logger.Info("Exit signals received... ")
		*s = true
		wg.Done()
	}(&Signal)
	return &Signal
}
